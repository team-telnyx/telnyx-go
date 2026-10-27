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

// WhatsappHostedSignupService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWhatsappHostedSignupService] method instead.
type WhatsappHostedSignupService struct {
	Options []option.RequestOption
}

// NewWhatsappHostedSignupService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewWhatsappHostedSignupService(opts ...option.RequestOption) (r WhatsappHostedSignupService) {
	r = WhatsappHostedSignupService{}
	r.Options = opts
	return
}

// Create a time-limited onboarding URL on the Telnyx-hosted Embedded Signup page
// for a Tech Provider's end customer. Set `signup_type` to `coexistence` to launch
// Meta's WhatsApp Business app coexistence flow instead of standard Cloud API
// onboarding.
func (r *WhatsappHostedSignupService) New(ctx context.Context, body WhatsappHostedSignupNewParams, opts ...option.RequestOption) (res *WhatsappHostedSignupNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "whatsapp/hosted_signups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a hosted signup and its onboarding URL.
func (r *WhatsappHostedSignupService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *WhatsappHostedSignupGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("whatsapp/hosted_signups/%s", url.PathEscape(id))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List hosted signups created by the account, newest first. Expired hosted signups
// are included.
func (r *WhatsappHostedSignupService) List(ctx context.Context, query WhatsappHostedSignupListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[WhatsappHostedSignupListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "whatsapp/hosted_signups"
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

// List hosted signups created by the account, newest first. Expired hosted signups
// are included.
func (r *WhatsappHostedSignupService) ListAutoPaging(ctx context.Context, query WhatsappHostedSignupListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[WhatsappHostedSignupListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

type WhatsappHostedSignupNewResponse struct {
	Data WhatsappHostedSignupNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappHostedSignupNewResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappHostedSignupNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappHostedSignupNewResponseData struct {
	// Hosted signup ID
	ID string `json:"id"`
	// Meta App ID the hosted signup onboards customers into
	AppID     string    `json:"app_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// When the onboarding URL stops working
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// Signed token embedded in `url`
	Jwt string `json:"jwt"`
	// Any of "whatsapp_jwt".
	RecordType string `json:"record_type"`
	// Meta Partner Solution ID linked to the app. Returned only when creating a hosted
	// signup.
	SolutionID string `json:"solution_id"`
	// Hosted onboarding URL to share with the end customer
	URL string `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AppID       respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Jwt         respjson.Field
		RecordType  respjson.Field
		SolutionID  respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappHostedSignupNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *WhatsappHostedSignupNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappHostedSignupGetResponse struct {
	Data WhatsappHostedSignupGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappHostedSignupGetResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappHostedSignupGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappHostedSignupGetResponseData struct {
	// Hosted signup ID
	ID string `json:"id"`
	// Meta App ID the hosted signup onboards customers into
	AppID     string    `json:"app_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// When the onboarding URL stops working
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// Signed token embedded in `url`
	Jwt string `json:"jwt"`
	// Any of "whatsapp_jwt".
	RecordType string `json:"record_type"`
	// Meta Partner Solution ID linked to the app. Returned only when creating a hosted
	// signup.
	SolutionID string `json:"solution_id"`
	// Hosted onboarding URL to share with the end customer
	URL string `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AppID       respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Jwt         respjson.Field
		RecordType  respjson.Field
		SolutionID  respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappHostedSignupGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *WhatsappHostedSignupGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappHostedSignupListResponse struct {
	// Hosted signup ID
	ID string `json:"id"`
	// Meta App ID the hosted signup onboards customers into
	AppID     string    `json:"app_id"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// When the onboarding URL stops working
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// Signed token embedded in `url`
	Jwt string `json:"jwt"`
	// Any of "whatsapp_jwt".
	RecordType WhatsappHostedSignupListResponseRecordType `json:"record_type"`
	// Meta Partner Solution ID linked to the app. Returned only when creating a hosted
	// signup.
	SolutionID string `json:"solution_id"`
	// Hosted onboarding URL to share with the end customer
	URL string `json:"url" format:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AppID       respjson.Field
		CreatedAt   respjson.Field
		ExpiresAt   respjson.Field
		Jwt         respjson.Field
		RecordType  respjson.Field
		SolutionID  respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WhatsappHostedSignupListResponse) RawJSON() string { return r.JSON.raw }
func (r *WhatsappHostedSignupListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WhatsappHostedSignupListResponseRecordType string

const (
	WhatsappHostedSignupListResponseRecordTypeWhatsappJwt WhatsappHostedSignupListResponseRecordType = "whatsapp_jwt"
)

type WhatsappHostedSignupNewParams struct {
	// Meta App ID linked to the Telnyx account or its manager account. The app must be
	// enabled.
	AppID string `json:"app_id" api:"required"`
	// When the onboarding URL expires, in RFC 3339 format. Must be in the future and
	// at most 3 days ahead. Defaults to 24 hours after creation.
	ExpiresAt param.Opt[time.Time] `json:"expires_at,omitzero" format:"date-time"`
	// Onboarding flow the hosted page launches. Use `coexistence` for a number that
	// remains active in the WhatsApp Business app.
	//
	// Any of "standard", "coexistence".
	SignupType WhatsappHostedSignupNewParamsSignupType `json:"signup_type,omitzero"`
	paramObj
}

func (r WhatsappHostedSignupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow WhatsappHostedSignupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WhatsappHostedSignupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Onboarding flow the hosted page launches. Use `coexistence` for a number that
// remains active in the WhatsApp Business app.
type WhatsappHostedSignupNewParamsSignupType string

const (
	WhatsappHostedSignupNewParamsSignupTypeStandard    WhatsappHostedSignupNewParamsSignupType = "standard"
	WhatsappHostedSignupNewParamsSignupTypeCoexistence WhatsappHostedSignupNewParamsSignupType = "coexistence"
)

type WhatsappHostedSignupListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [WhatsappHostedSignupListParams]'s query parameters as
// `url.Values`.
func (r WhatsappHostedSignupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
