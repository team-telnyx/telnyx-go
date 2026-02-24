// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

// MessagingHostedNumberService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingHostedNumberService] method instead.
type MessagingHostedNumberService struct {
	Options []option.RequestOption
}

// NewMessagingHostedNumberService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMessagingHostedNumberService(opts ...option.RequestOption) (r MessagingHostedNumberService) {
	r = MessagingHostedNumberService{}
	r.Options = opts
	return
}

// Retrieve a specific messaging hosted number by its ID or phone number.
func (r *MessagingHostedNumberService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *MessagingHostedNumberGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging_hosted_numbers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update the messaging settings for a hosted number.
func (r *MessagingHostedNumberService) Update(ctx context.Context, id string, body MessagingHostedNumberUpdateParams, opts ...option.RequestOption) (res *MessagingHostedNumberUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging_hosted_numbers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all hosted numbers associated with the authenticated user.
func (r *MessagingHostedNumberService) List(ctx context.Context, query MessagingHostedNumberListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[shared.PhoneNumberWithMessagingSettings], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "messaging_hosted_numbers"
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

// List all hosted numbers associated with the authenticated user.
func (r *MessagingHostedNumberService) ListAutoPaging(ctx context.Context, query MessagingHostedNumberListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[shared.PhoneNumberWithMessagingSettings] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete a messaging hosted number
func (r *MessagingHostedNumberService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *MessagingHostedNumberDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("messaging_hosted_numbers/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type MessagingHostedNumberGetResponse struct {
	Data shared.PhoneNumberWithMessagingSettings `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingHostedNumberGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingHostedNumberGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingHostedNumberUpdateResponse struct {
	Data shared.PhoneNumberWithMessagingSettings `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingHostedNumberUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingHostedNumberUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingHostedNumberDeleteResponse struct {
	Data shared.MessagingHostedNumberOrder `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingHostedNumberDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingHostedNumberDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingHostedNumberUpdateParams struct {
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
	// Tags to set on this phone number.
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r MessagingHostedNumberUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow MessagingHostedNumberUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessagingHostedNumberUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingHostedNumberListParams struct {
	// Filter by messaging profile ID.
	FilterMessagingProfileID param.Opt[string] `query:"filter[messaging_profile_id],omitzero" format:"uuid" json:"-"`
	// Filter by exact phone number.
	FilterPhoneNumber param.Opt[string] `query:"filter[phone_number],omitzero" json:"-"`
	// Filter by phone number substring.
	FilterPhoneNumberContains param.Opt[string] `query:"filter[phone_number][contains],omitzero" json:"-"`
	PageNumber                param.Opt[int64]  `query:"page[number],omitzero" json:"-"`
	PageSize                  param.Opt[int64]  `query:"page[size],omitzero" json:"-"`
	// Sort by phone number.
	//
	// Any of "asc", "desc".
	SortPhoneNumber MessagingHostedNumberListParamsSortPhoneNumber `query:"sort[phone_number],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [MessagingHostedNumberListParams]'s query parameters as
// `url.Values`.
func (r MessagingHostedNumberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort by phone number.
type MessagingHostedNumberListParamsSortPhoneNumber string

const (
	MessagingHostedNumberListParamsSortPhoneNumberAsc  MessagingHostedNumberListParamsSortPhoneNumber = "asc"
	MessagingHostedNumberListParamsSortPhoneNumberDesc MessagingHostedNumberListParamsSortPhoneNumber = "desc"
)
