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

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v3/shared"
)

// MobilePhoneNumberMessagingService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMobilePhoneNumberMessagingService] method instead.
type MobilePhoneNumberMessagingService struct {
	Options []option.RequestOption
}

// NewMobilePhoneNumberMessagingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMobilePhoneNumberMessagingService(opts ...option.RequestOption) (r MobilePhoneNumberMessagingService) {
	r = MobilePhoneNumberMessagingService{}
	r.Options = opts
	return
}

// Retrieve a mobile phone number with messaging settings
func (r *MobilePhoneNumberMessagingService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MobilePhoneNumberMessagingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("mobile_phone_numbers/%s/messaging", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List mobile phone numbers with messaging settings
func (r *MobilePhoneNumberMessagingService) List(ctx context.Context, query MobilePhoneNumberMessagingListParams, opts ...option.RequestOption) (res *MobilePhoneNumberMessagingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "mobile_phone_numbers/messaging"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type MobilePhoneNumberMessagingGetResponse struct {
	Data MobilePhoneNumberMessagingGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberMessagingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberMessagingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberMessagingGetResponseData struct {
	// Identifies the type of resource.
	ID string `json:"id"`
	// ISO 3166-1 alpha-2 country code.
	CountryCode string `json:"country_code"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time                                         `json:"created_at" format:"date-time"`
	Features  MobilePhoneNumberMessagingGetResponseDataFeatures `json:"features"`
	// The messaging product that the number is registered to use
	MessagingProduct string `json:"messaging_product"`
	// Unique identifier for a messaging profile.
	MessagingProfileID string `json:"messaging_profile_id,nullable"`
	// +E.164 formatted phone number.
	PhoneNumber string `json:"phone_number" format:"e164"`
	// Identifies the type of the resource.
	//
	// Any of "messaging_phone_number", "messaging_settings".
	RecordType string `json:"record_type"`
	// The messaging traffic or use case for which the number is currently configured.
	TrafficType string `json:"traffic_type"`
	// The type of the phone number
	//
	// Any of "longcode".
	Type string `json:"type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CountryCode        respjson.Field
		CreatedAt          respjson.Field
		Features           respjson.Field
		MessagingProduct   respjson.Field
		MessagingProfileID respjson.Field
		PhoneNumber        respjson.Field
		RecordType         respjson.Field
		TrafficType        respjson.Field
		Type               respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberMessagingGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberMessagingGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberMessagingGetResponseDataFeatures struct {
	// The set of features available for a specific messaging use case (SMS or MMS).
	// Features can vary depending on the characteristics the phone number, as well as
	// its current product configuration.
	SMS shared.MessagingFeatureSet `json:"sms,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SMS         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberMessagingGetResponseDataFeatures) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberMessagingGetResponseDataFeatures) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberMessagingListResponse struct {
	Data []MobilePhoneNumberMessagingListResponseData `json:"data"`
	Meta PaginationMeta                               `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberMessagingListResponse) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberMessagingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberMessagingListResponseData struct {
	// Identifies the type of resource.
	ID string `json:"id"`
	// ISO 3166-1 alpha-2 country code.
	CountryCode string `json:"country_code"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time                                          `json:"created_at" format:"date-time"`
	Features  MobilePhoneNumberMessagingListResponseDataFeatures `json:"features"`
	// The messaging product that the number is registered to use
	MessagingProduct string `json:"messaging_product"`
	// Unique identifier for a messaging profile.
	MessagingProfileID string `json:"messaging_profile_id,nullable"`
	// +E.164 formatted phone number.
	PhoneNumber string `json:"phone_number" format:"e164"`
	// Identifies the type of the resource.
	//
	// Any of "messaging_phone_number", "messaging_settings".
	RecordType string `json:"record_type"`
	// The messaging traffic or use case for which the number is currently configured.
	TrafficType string `json:"traffic_type"`
	// The type of the phone number
	//
	// Any of "longcode".
	Type string `json:"type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CountryCode        respjson.Field
		CreatedAt          respjson.Field
		Features           respjson.Field
		MessagingProduct   respjson.Field
		MessagingProfileID respjson.Field
		PhoneNumber        respjson.Field
		RecordType         respjson.Field
		TrafficType        respjson.Field
		Type               respjson.Field
		UpdatedAt          respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberMessagingListResponseData) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberMessagingListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberMessagingListResponseDataFeatures struct {
	// The set of features available for a specific messaging use case (SMS or MMS).
	// Features can vary depending on the characteristics the phone number, as well as
	// its current product configuration.
	SMS shared.MessagingFeatureSet `json:"sms,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SMS         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MobilePhoneNumberMessagingListResponseDataFeatures) RawJSON() string { return r.JSON.raw }
func (r *MobilePhoneNumberMessagingListResponseDataFeatures) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MobilePhoneNumberMessagingListParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page MobilePhoneNumberMessagingListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MobilePhoneNumberMessagingListParams]'s query parameters as
// `url.Values`.
func (r MobilePhoneNumberMessagingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type MobilePhoneNumberMessagingListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MobilePhoneNumberMessagingListParamsPage]'s query
// parameters as `url.Values`.
func (r MobilePhoneNumberMessagingListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
