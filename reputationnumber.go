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

// Associate phone numbers with an enterprise for reputation monitoring and
// retrieve reputation scores
//
// ReputationNumberService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReputationNumberService] method instead.
type ReputationNumberService struct {
	Options []option.RequestOption
}

// NewReputationNumberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewReputationNumberService(opts ...option.RequestOption) (r ReputationNumberService) {
	r = ReputationNumberService{}
	r.Options = opts
	return
}

// Get reputation data for a specific phone number without requiring an
// `enterprise_id`.
//
// Same response as the enterprise-scoped endpoint. Uses cached data by default.
func (r *ReputationNumberService) Get(ctx context.Context, phoneNumber string, query ReputationNumberGetParams, opts ...option.RequestOption) (res *ReputationNumberGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return nil, err
	}
	path := fmt.Sprintf("reputation/numbers/%s", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// List all phone numbers enrolled in Number Reputation monitoring for your
// account. This is a simplified endpoint that does not require an `enterprise_id`
// — it returns numbers across all your enterprises.
//
// Supports pagination and filtering by phone number.
func (r *ReputationNumberService) List(ctx context.Context, query ReputationNumberListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[ReputationNumberListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "reputation/numbers"
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

// List all phone numbers enrolled in Number Reputation monitoring for your
// account. This is a simplified endpoint that does not require an `enterprise_id`
// — it returns numbers across all your enterprises.
//
// Supports pagination and filtering by phone number.
func (r *ReputationNumberService) ListAutoPaging(ctx context.Context, query ReputationNumberListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[ReputationNumberListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Remove a phone number from Number Reputation monitoring without requiring an
// `enterprise_id`.
func (r *ReputationNumberService) Delete(ctx context.Context, phoneNumber string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return err
	}
	path := fmt.Sprintf("reputation/numbers/%s", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type ReputationNumberGetResponse struct {
	Data ReputationNumberGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReputationNumberGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ReputationNumberGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReputationNumberGetResponseData struct {
	// Unique identifier
	ID string `json:"id" format:"uuid"`
	// When the number was associated
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// ID of the associated enterprise
	EnterpriseID string `json:"enterprise_id" format:"uuid"`
	// Phone number in E.164 format
	PhoneNumber string `json:"phone_number"`
	// Reputation metrics (null if not yet fetched)
	ReputationData ReputationNumberGetResponseDataReputationData `json:"reputation_data"`
	// When the record was last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		EnterpriseID   respjson.Field
		PhoneNumber    respjson.Field
		ReputationData respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReputationNumberGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *ReputationNumberGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reputation metrics
type ReputationNumberGetResponseDataReputationData struct {
	// Connection quality metric (0–100)
	ConnectionScore int64 `json:"connection_score" api:"nullable"`
	// Engagement metric (0–100). Higher = more positive engagement
	EngagementScore int64 `json:"engagement_score" api:"nullable"`
	// Timestamp of the last reputation data refresh
	LastRefreshedAt time.Time `json:"last_refreshed_at" api:"nullable" format:"date-time"`
	// Maturity metric (0–100). Higher = more established number
	MaturityScore int64 `json:"maturity_score" api:"nullable"`
	// Sentiment metric (0–100). Higher = more positive sentiment
	SentimentScore int64 `json:"sentiment_score" api:"nullable"`
	// Spam category classification (e.g., Telemarketing, Debt Collector)
	SpamCategory string `json:"spam_category" api:"nullable"`
	// Overall spam risk level
	//
	// Any of "low", "medium", "high".
	SpamRisk string `json:"spam_risk" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectionScore respjson.Field
		EngagementScore respjson.Field
		LastRefreshedAt respjson.Field
		MaturityScore   respjson.Field
		SentimentScore  respjson.Field
		SpamCategory    respjson.Field
		SpamRisk        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReputationNumberGetResponseDataReputationData) RawJSON() string { return r.JSON.raw }
func (r *ReputationNumberGetResponseDataReputationData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReputationNumberListResponse struct {
	// Unique identifier
	ID string `json:"id" format:"uuid"`
	// When the number was associated
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// ID of the associated enterprise
	EnterpriseID string `json:"enterprise_id" format:"uuid"`
	// Phone number in E.164 format
	PhoneNumber string `json:"phone_number"`
	// Reputation metrics (null if not yet fetched)
	ReputationData ReputationNumberListResponseReputationData `json:"reputation_data"`
	// When the record was last updated
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		EnterpriseID   respjson.Field
		PhoneNumber    respjson.Field
		ReputationData respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReputationNumberListResponse) RawJSON() string { return r.JSON.raw }
func (r *ReputationNumberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Reputation metrics
type ReputationNumberListResponseReputationData struct {
	// Connection quality metric (0–100)
	ConnectionScore int64 `json:"connection_score" api:"nullable"`
	// Engagement metric (0–100). Higher = more positive engagement
	EngagementScore int64 `json:"engagement_score" api:"nullable"`
	// Timestamp of the last reputation data refresh
	LastRefreshedAt time.Time `json:"last_refreshed_at" api:"nullable" format:"date-time"`
	// Maturity metric (0–100). Higher = more established number
	MaturityScore int64 `json:"maturity_score" api:"nullable"`
	// Sentiment metric (0–100). Higher = more positive sentiment
	SentimentScore int64 `json:"sentiment_score" api:"nullable"`
	// Spam category classification (e.g., Telemarketing, Debt Collector)
	SpamCategory string `json:"spam_category" api:"nullable"`
	// Overall spam risk level
	//
	// Any of "low", "medium", "high".
	SpamRisk string `json:"spam_risk" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectionScore respjson.Field
		EngagementScore respjson.Field
		LastRefreshedAt respjson.Field
		MaturityScore   respjson.Field
		SentimentScore  respjson.Field
		SpamCategory    respjson.Field
		SpamRisk        respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ReputationNumberListResponseReputationData) RawJSON() string { return r.JSON.raw }
func (r *ReputationNumberListResponseReputationData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ReputationNumberGetParams struct {
	// When true, fetches fresh reputation data (incurs API cost). When false, returns
	// cached data.
	Fresh param.Opt[bool] `query:"fresh,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReputationNumberGetParams]'s query parameters as
// `url.Values`.
func (r ReputationNumberGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ReputationNumberListParams struct {
	// Page number (1-indexed)
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of items per page
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Filter by specific phone number (E.164 format)
	PhoneNumber param.Opt[string] `query:"phone_number,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ReputationNumberListParams]'s query parameters as
// `url.Values`.
func (r ReputationNumberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
