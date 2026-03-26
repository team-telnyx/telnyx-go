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
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

// Associate phone numbers with an enterprise for reputation monitoring and
// retrieve reputation scores
//
// EnterpriseReputationNumberService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnterpriseReputationNumberService] method instead.
type EnterpriseReputationNumberService struct {
	Options []option.RequestOption
}

// NewEnterpriseReputationNumberService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewEnterpriseReputationNumberService(opts ...option.RequestOption) (r EnterpriseReputationNumberService) {
	r = EnterpriseReputationNumberService{}
	r.Options = opts
	return
}

// Associate one or more phone numbers with an enterprise for Number Reputation
// monitoring.
//
// **Validations:**
//
//   - Phone numbers must be in E.164 format (e.g., `+16035551234`)
//   - Phone numbers must be in-service and belong to your account (verified via
//     Warehouse)
//   - Phone numbers must be US local numbers
//   - Phone numbers cannot already be associated with any enterprise
//
// **Note:** This operation is atomic — if any number fails validation, the entire
// request fails.
//
// **Maximum:** 100 phone numbers per request.
func (r *EnterpriseReputationNumberService) New(ctx context.Context, enterpriseID string, body EnterpriseReputationNumberNewParams, opts ...option.RequestOption) (res *EnterpriseReputationNumberNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation/numbers", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get detailed reputation data for a specific phone number associated with an
// enterprise.
//
// **Query Parameters:**
//
//   - `fresh` (default: `false`): When `true`, fetches fresh reputation data (incurs
//     API cost). When `false`, returns cached data. If no cached data exists, fresh
//     data is automatically fetched.
//
// **Returns:**
//
// - `spam_risk`: Overall spam risk level (`low`, `medium`, `high`)
// - `spam_category`: Spam category classification
// - `maturity_score`: Maturity metric (0–100)
// - `connection_score`: Connection quality metric (0–100)
// - `engagement_score`: Engagement metric (0–100)
// - `sentiment_score`: Sentiment metric (0–100)
// - `last_refreshed_at`: Timestamp of last data refresh
func (r *EnterpriseReputationNumberService) Get(ctx context.Context, phoneNumber string, params EnterpriseReputationNumberGetParams, opts ...option.RequestOption) (res *EnterpriseReputationNumberGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.EnterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation/numbers/%s", params.EnterpriseID, phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// List all phone numbers associated with an enterprise for Number Reputation
// monitoring.
//
// Returns phone numbers with their cached reputation data (if available). Supports
// pagination and filtering by phone number.
func (r *EnterpriseReputationNumberService) List(ctx context.Context, enterpriseID string, query EnterpriseReputationNumberListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[shared.ReputationPhoneNumberWithReputationData], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation/numbers", enterpriseID)
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

// List all phone numbers associated with an enterprise for Number Reputation
// monitoring.
//
// Returns phone numbers with their cached reputation data (if available). Supports
// pagination and filtering by phone number.
func (r *EnterpriseReputationNumberService) ListAutoPaging(ctx context.Context, enterpriseID string, query EnterpriseReputationNumberListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[shared.ReputationPhoneNumberWithReputationData] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, enterpriseID, query, opts...))
}

// Remove a phone number from Number Reputation monitoring for an enterprise.
//
// The number will no longer be tracked and reputation data will no longer be
// refreshed.
func (r *EnterpriseReputationNumberService) Delete(ctx context.Context, phoneNumber string, body EnterpriseReputationNumberDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.EnterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return err
	}
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return err
	}
	path := fmt.Sprintf("enterprises/%s/reputation/numbers/%s", body.EnterpriseID, phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type EnterpriseReputationNumberNewResponse struct {
	Data []EnterpriseReputationNumberNewResponseData `json:"data"`
	Meta shared.MetaInfo                             `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationNumberNewResponse) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationNumberNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationNumberNewResponseData struct {
	// Unique identifier
	ID string `json:"id" format:"uuid"`
	// When the number was associated
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// ID of the associated enterprise
	EnterpriseID string `json:"enterprise_id" format:"uuid"`
	// Phone number in E.164 format
	PhoneNumber string `json:"phone_number"`
	// When the record was last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		EnterpriseID respjson.Field
		PhoneNumber  respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationNumberNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationNumberNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationNumberGetResponse struct {
	Data shared.ReputationPhoneNumberWithReputationData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EnterpriseReputationNumberGetResponse) RawJSON() string { return r.JSON.raw }
func (r *EnterpriseReputationNumberGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationNumberNewParams struct {
	// List of phone numbers to associate for reputation monitoring (max 100)
	PhoneNumbers []string `json:"phone_numbers,omitzero" api:"required"`
	paramObj
}

func (r EnterpriseReputationNumberNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseReputationNumberNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseReputationNumberNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationNumberGetParams struct {
	EnterpriseID string `path:"enterprise_id" api:"required" format:"uuid" json:"-"`
	// When true, fetches fresh reputation data (incurs API cost). When false, returns
	// cached data.
	Fresh param.Opt[bool] `query:"fresh,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EnterpriseReputationNumberGetParams]'s query parameters as
// `url.Values`.
func (r EnterpriseReputationNumberGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EnterpriseReputationNumberListParams struct {
	// Page number (1-indexed)
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of items per page
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Filter by specific phone number (E.164 format)
	PhoneNumber param.Opt[string] `query:"phone_number,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EnterpriseReputationNumberListParams]'s query parameters as
// `url.Values`.
func (r EnterpriseReputationNumberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type EnterpriseReputationNumberDeleteParams struct {
	EnterpriseID string `path:"enterprise_id" api:"required" format:"uuid" json:"-"`
	paramObj
}
