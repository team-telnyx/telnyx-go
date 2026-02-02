// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package requestconfig

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

var OAuth2Cache = newOAuth2Cache("https://api.telnyx.com/v2/oauth/token")

func newOAuth2Cache(tokenUrls ...string) map[string]*OAuth2State {
	state := make(map[string]*OAuth2State, len(tokenUrls))
	for _, url := range tokenUrls {
		state[url] = &OAuth2State{authPath: url}
	}
	return state
}

// OAuth2State represents the OAuth2 provider configuration and manages token
// caching.
type OAuth2State struct {
	authPath string
	entries  sync.Map // map[oAuth2ClientCredentials]*oAuthTokenInfo
}

type oAuth2ClientCredentials struct {
	ClientID     string
	ClientSecret string
}

type oAuthTokenInfo struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`

	expiresAt     time.Time
	reqTokenMutex sync.Mutex
}

func (tok *oAuthTokenInfo) reset() {
	tok.AccessToken = ""
	tok.ExpiresIn = 0
	tok.expiresAt = time.Time{}
}

func (auth *oAuthTokenInfo) validToken() string {
	if auth.AccessToken != "" && time.Now().Before(auth.expiresAt.Add(-10*time.Second)) {
		return auth.AccessToken
	}
	return ""
}

func (state *OAuth2State) GetToken(cfg *RequestConfig) (string, error) {
	tokenInfo := state.find(cfg)

	valid := tokenInfo.validToken()
	if valid != "" {
		return valid, nil
	}

	tokenInfo.reqTokenMutex.Lock()
	defer tokenInfo.reqTokenMutex.Unlock()

	tokenInfo.reset()

	authUrl, err := state.authUrl(cfg)
	if err != nil {
		return "", err
	}

	oAuthReq, err := http.NewRequestWithContext(cfg.Context, http.MethodPost, authUrl, strings.NewReader("grant_type=client_credentials"))
	if err != nil {
		return "", fmt.Errorf("requestconfig: failed to create OAuth2 token request: %w", err)
	}

	encoded := base64.StdEncoding.EncodeToString(fmt.Appendf(nil, "%s:%s", cfg.ClientID, cfg.ClientSecret))
	oAuthReq.Header.Set("Authorization", fmt.Sprintf("Basic %s", encoded))
	oAuthReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	handler := cfg.HTTPClient.Do
	if cfg.CustomHTTPDoer != nil {
		handler = cfg.CustomHTTPDoer.Do
	}

	oauthRes, err := handler(oAuthReq)
	if err != nil {
		return "", fmt.Errorf("requestconfig: OAuth2 token request failed: %w", err)
	}
	defer oauthRes.Body.Close()

	if oauthRes.StatusCode != http.StatusOK {
		return "", fmt.Errorf("requestconfig: OAuth2 token request returned status %d, expected %d", oauthRes.StatusCode, http.StatusOK)
	}

	contents, err := io.ReadAll(oauthRes.Body)
	if err != nil {
		return "", fmt.Errorf("requestconfig: failed to read OAuth2 token response: %w", err)
	}

	err = json.Unmarshal(contents, tokenInfo)
	if err != nil {
		return "", fmt.Errorf("requestconfig: failed to parse OAuth2 token response: %w", err)
	}

	if tokenInfo.AccessToken == "" {
		return "", fmt.Errorf("requestconfig: OAuth2 token response missing access_token")
	}

	tokenInfo.expiresAt = time.Now().Add(time.Duration(tokenInfo.ExpiresIn) * time.Second)

	return tokenInfo.AccessToken, nil
}

func (r *OAuth2State) find(cfg *RequestConfig) *oAuthTokenInfo {
	key := oAuth2ClientCredentials{cfg.ClientID, cfg.ClientSecret}
	resp, ok := r.entries.Load(key)
	if !ok {
		resp, _ = r.entries.LoadOrStore(key, new(oAuthTokenInfo))
	}

	return resp.(*oAuthTokenInfo)
}

func (state *OAuth2State) authUrl(cfg *RequestConfig) (string, error) {
	authUrl, err := cfg.BaseURL.Parse(strings.TrimLeft(state.authPath, "/"))
	if err != nil {
		err = fmt.Errorf("requestconfig: failed to parse OAuth2 token URL: %w", err)
		return "", err
	}
	return authUrl.String(), nil
}
