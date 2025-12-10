// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// Number10dlcService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumber10dlcService] method instead.
type Number10dlcService struct {
	Options              []option.RequestOption
	Brand                Number10dlcBrandService
	Campaign             Number10dlcCampaignService
	CampaignBuilder      Number10dlcCampaignBuilderService
	PhoneNumberCampaigns Number10dlcPhoneNumberCampaignService
}

// NewNumber10dlcService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNumber10dlcService(opts ...option.RequestOption) (r Number10dlcService) {
	r = Number10dlcService{}
	r.Options = opts
	r.Brand = NewNumber10dlcBrandService(opts...)
	r.Campaign = NewNumber10dlcCampaignService(opts...)
	r.CampaignBuilder = NewNumber10dlcCampaignBuilderService(opts...)
	r.PhoneNumberCampaigns = NewNumber10dlcPhoneNumberCampaignService(opts...)
	return
}

// Get Enum
func (r *Number10dlcService) GetEnum(ctx context.Context, endpoint Number10dlcGetEnumParamsEndpoint, opts ...option.RequestOption) (res *Number10dlcGetEnumResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("10dlc/enum/%v", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Number10dlcGetEnumResponseUnion contains all possible properties and values from
// [[]string], [[]map[string]any], [map[string]any], [map[string]any],
// [Number10dlcGetEnumResponseEnumPaginatedResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfEnumStringListResponse OfEnumObjectListResponse
// OfNumber10dlcGetEnumResponseEnumObjecToObjecttResponseItem]
type Number10dlcGetEnumResponseUnion struct {
	// This field will be present if the value is a [[]string] instead of an object.
	OfEnumStringListResponse []string `json:",inline"`
	// This field will be present if the value is a [[]map[string]any] instead of an
	// object.
	OfEnumObjectListResponse []map[string]any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfNumber10dlcGetEnumResponseEnumObjecToObjecttResponseItem any `json:",inline"`
	// This field is from variant [Number10dlcGetEnumResponseEnumPaginatedResponse].
	Page int64 `json:"page"`
	// This field is from variant [Number10dlcGetEnumResponseEnumPaginatedResponse].
	Records []map[string]any `json:"records"`
	// This field is from variant [Number10dlcGetEnumResponseEnumPaginatedResponse].
	TotalRecords int64 `json:"totalRecords"`
	JSON         struct {
		OfEnumStringListResponse                                   respjson.Field
		OfEnumObjectListResponse                                   respjson.Field
		OfNumber10dlcGetEnumResponseEnumObjecToObjecttResponseItem respjson.Field
		Page                                                       respjson.Field
		Records                                                    respjson.Field
		TotalRecords                                               respjson.Field
		raw                                                        string
	} `json:"-"`
}

func (u Number10dlcGetEnumResponseUnion) AsEnumStringListResponse() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u Number10dlcGetEnumResponseUnion) AsEnumObjectListResponse() (v []map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u Number10dlcGetEnumResponseUnion) AsEnumObjectToStringResponse() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u Number10dlcGetEnumResponseUnion) AsEnumObjecToObjecttResponse() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u Number10dlcGetEnumResponseUnion) AsEnumPaginatedResponse() (v Number10dlcGetEnumResponseEnumPaginatedResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u Number10dlcGetEnumResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *Number10dlcGetEnumResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcGetEnumResponseEnumPaginatedResponse struct {
	Page         int64            `json:"page,required"`
	Records      []map[string]any `json:"records,required"`
	TotalRecords int64            `json:"totalRecords,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Page         respjson.Field
		Records      respjson.Field
		TotalRecords respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcGetEnumResponseEnumPaginatedResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcGetEnumResponseEnumPaginatedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcGetEnumParamsEndpoint string

const (
	Number10dlcGetEnumParamsEndpointMno                   Number10dlcGetEnumParamsEndpoint = "mno"
	Number10dlcGetEnumParamsEndpointOptionalAttributes    Number10dlcGetEnumParamsEndpoint = "optionalAttributes"
	Number10dlcGetEnumParamsEndpointUsecase               Number10dlcGetEnumParamsEndpoint = "usecase"
	Number10dlcGetEnumParamsEndpointVertical              Number10dlcGetEnumParamsEndpoint = "vertical"
	Number10dlcGetEnumParamsEndpointAltBusinessIDType     Number10dlcGetEnumParamsEndpoint = "altBusinessIdType"
	Number10dlcGetEnumParamsEndpointBrandIdentityStatus   Number10dlcGetEnumParamsEndpoint = "brandIdentityStatus"
	Number10dlcGetEnumParamsEndpointBrandRelationship     Number10dlcGetEnumParamsEndpoint = "brandRelationship"
	Number10dlcGetEnumParamsEndpointCampaignStatus        Number10dlcGetEnumParamsEndpoint = "campaignStatus"
	Number10dlcGetEnumParamsEndpointEntityType            Number10dlcGetEnumParamsEndpoint = "entityType"
	Number10dlcGetEnumParamsEndpointExtVettingProvider    Number10dlcGetEnumParamsEndpoint = "extVettingProvider"
	Number10dlcGetEnumParamsEndpointVettingStatus         Number10dlcGetEnumParamsEndpoint = "vettingStatus"
	Number10dlcGetEnumParamsEndpointBrandStatus           Number10dlcGetEnumParamsEndpoint = "brandStatus"
	Number10dlcGetEnumParamsEndpointOperationStatus       Number10dlcGetEnumParamsEndpoint = "operationStatus"
	Number10dlcGetEnumParamsEndpointApprovedPublicCompany Number10dlcGetEnumParamsEndpoint = "approvedPublicCompany"
	Number10dlcGetEnumParamsEndpointStockExchange         Number10dlcGetEnumParamsEndpoint = "stockExchange"
	Number10dlcGetEnumParamsEndpointVettingClass          Number10dlcGetEnumParamsEndpoint = "vettingClass"
)
