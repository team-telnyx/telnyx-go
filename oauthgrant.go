// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// OAuthGrantService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOAuthGrantService] method instead.
type OAuthGrantService struct {
	Options []option.RequestOption
}

// NewOAuthGrantService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOAuthGrantService(opts ...option.RequestOption) (r OAuthGrantService) {
	r = OAuthGrantService{}
	r.Options = opts
	return
}

// Retrieve a single OAuth grant by ID
func (r *OAuthGrantService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *OAuthGrantGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("oauth_grants/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a paginated list of OAuth grants for the authenticated user
func (r *OAuthGrantService) List(ctx context.Context, query OAuthGrantListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[OAuthGrant], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "oauth_grants"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve a paginated list of OAuth grants for the authenticated user
func (r *OAuthGrantService) ListAutoPaging(ctx context.Context, query OAuthGrantListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[OAuthGrant] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Revoke an OAuth grant
func (r *OAuthGrantService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *OAuthGrantDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("oauth_grants/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type OAuthGrant struct {
	// Unique identifier for the OAuth grant
	ID string `json:"id,required" format:"uuid"`
	// OAuth client identifier
	ClientID string `json:"client_id,required"`
	// Timestamp when the grant was created
	CreatedAt time.Time `json:"created_at,required" format:"date-time"`
	// Record type identifier
	//
	// Any of "oauth_grant".
	RecordType OAuthGrantRecordType `json:"record_type,required"`
	// List of granted OAuth scopes
	Scopes []string `json:"scopes,required"`
	// Timestamp when the grant was last used
	LastUsedAt time.Time `json:"last_used_at,nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ClientID    respjson.Field
		CreatedAt   respjson.Field
		RecordType  respjson.Field
		Scopes      respjson.Field
		LastUsedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthGrant) RawJSON() string { return r.JSON.raw }
func (r *OAuthGrant) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Record type identifier
type OAuthGrantRecordType string

const (
	OAuthGrantRecordTypeOAuthGrant OAuthGrantRecordType = "oauth_grant"
)

type OAuthGrantGetResponse struct {
	Data OAuthGrant `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthGrantGetResponse) RawJSON() string { return r.JSON.raw }
func (r *OAuthGrantGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthGrantDeleteResponse struct {
	Data OAuthGrant `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthGrantDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *OAuthGrantDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthGrantListParams struct {
	// Page number
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of results per page
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [OAuthGrantListParams]'s query parameters as `url.Values`.
func (r OAuthGrantListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
