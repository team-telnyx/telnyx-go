// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v3/shared"
)

// PhoneNumberMessagingService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhoneNumberMessagingService] method instead.
type PhoneNumberMessagingService struct {
	Options []option.RequestOption
}

// NewPhoneNumberMessagingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPhoneNumberMessagingService(opts ...option.RequestOption) (r PhoneNumberMessagingService) {
	r = PhoneNumberMessagingService{}
	r.Options = opts
	return
}

// Retrieve a phone number with messaging settings
func (r *PhoneNumberMessagingService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PhoneNumberMessagingGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("phone_numbers/%s/messaging", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the messaging profile and/or messaging product of a phone number
func (r *PhoneNumberMessagingService) Update(ctx context.Context, id string, body PhoneNumberMessagingUpdateParams, opts ...option.RequestOption) (res *PhoneNumberMessagingUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("phone_numbers/%s/messaging", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List phone numbers with messaging settings
func (r *PhoneNumberMessagingService) List(ctx context.Context, query PhoneNumberMessagingListParams, opts ...option.RequestOption) (res *PhoneNumberMessagingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "phone_numbers/messaging"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PhoneNumberMessagingGetResponse struct {
	Data shared.PhoneNumberWithMessagingSettings `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberMessagingGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberMessagingGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberMessagingUpdateResponse struct {
	Data shared.PhoneNumberWithMessagingSettings `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberMessagingUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberMessagingUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberMessagingListResponse struct {
	Data []shared.PhoneNumberWithMessagingSettings `json:"data"`
	Meta PaginationMeta                            `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberMessagingListResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberMessagingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberMessagingUpdateParams struct {
	// Configure the messaging product for this number:
	//
	// - Omit this field or set its value to `null` to keep the current value.
	// - Set this field to a quoted product ID to set this phone number to that product
	MessagingProduct param.Opt[string] `json:"messaging_product,omitzero"`
	// Configure the messaging profile this phone number is assigned to:
	//
	//   - Omit this field or set its value to `null` to keep the current value.
	//   - Set this field to `""` to unassign the number from its messaging profile
	//   - Set this field to a quoted UUID of a messaging profile to assign this number
	//     to that messaging profile
	MessagingProfileID param.Opt[string] `json:"messaging_profile_id,omitzero"`
	paramObj
}

func (r PhoneNumberMessagingUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberMessagingUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberMessagingUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberMessagingListParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page PhoneNumberMessagingListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberMessagingListParams]'s query parameters as
// `url.Values`.
func (r PhoneNumberMessagingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type PhoneNumberMessagingListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberMessagingListParamsPage]'s query parameters as
// `url.Values`.
func (r PhoneNumberMessagingListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
