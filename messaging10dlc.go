// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Messaging10dlcService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessaging10dlcService] method instead.
type Messaging10dlcService struct {
	Options                        []option.RequestOption
	Brand                          Messaging10dlcBrandService
	Campaign                       Messaging10dlcCampaignService
	CampaignBuilder                Messaging10dlcCampaignBuilderService
	PartnerCampaigns               Messaging10dlcPartnerCampaignService
	PhoneNumberCampaigns           Messaging10dlcPhoneNumberCampaignService
	PhoneNumberAssignmentByProfile Messaging10dlcPhoneNumberAssignmentByProfileService
}

// NewMessaging10dlcService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewMessaging10dlcService(opts ...option.RequestOption) (r Messaging10dlcService) {
	r = Messaging10dlcService{}
	r.Options = opts
	r.Brand = NewMessaging10dlcBrandService(opts...)
	r.Campaign = NewMessaging10dlcCampaignService(opts...)
	r.CampaignBuilder = NewMessaging10dlcCampaignBuilderService(opts...)
	r.PartnerCampaigns = NewMessaging10dlcPartnerCampaignService(opts...)
	r.PhoneNumberCampaigns = NewMessaging10dlcPhoneNumberCampaignService(opts...)
	r.PhoneNumberAssignmentByProfile = NewMessaging10dlcPhoneNumberAssignmentByProfileService(opts...)
	return
}

// Get Enum
func (r *Messaging10dlcService) GetEnum(ctx context.Context, endpoint Messaging10dlcGetEnumParamsEndpoint, opts ...option.RequestOption) (res *Messaging10dlcGetEnumResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("10dlc/enum/%v", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Messaging10dlcGetEnumResponseUnion contains all possible properties and values
// from [[]string], [[]map[string]any], [map[string]any], [map[string]any],
// [Messaging10dlcGetEnumResponseEnumPaginatedResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfEnumStringListResponse OfEnumObjectListResponse
// OfMessaging10dlcGetEnumResponseEnumObjecToObjecttResponseItem]
type Messaging10dlcGetEnumResponseUnion struct {
	// This field will be present if the value is a [[]string] instead of an object.
	OfEnumStringListResponse []string `json:",inline"`
	// This field will be present if the value is a [[]map[string]any] instead of an
	// object.
	OfEnumObjectListResponse []map[string]any `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfMessaging10dlcGetEnumResponseEnumObjecToObjecttResponseItem any `json:",inline"`
	// This field is from variant [Messaging10dlcGetEnumResponseEnumPaginatedResponse].
	Page int64 `json:"page"`
	// This field is from variant [Messaging10dlcGetEnumResponseEnumPaginatedResponse].
	Records []map[string]any `json:"records"`
	// This field is from variant [Messaging10dlcGetEnumResponseEnumPaginatedResponse].
	TotalRecords int64 `json:"totalRecords"`
	JSON         struct {
		OfEnumStringListResponse                                      respjson.Field
		OfEnumObjectListResponse                                      respjson.Field
		OfMessaging10dlcGetEnumResponseEnumObjecToObjecttResponseItem respjson.Field
		Page                                                          respjson.Field
		Records                                                       respjson.Field
		TotalRecords                                                  respjson.Field
		raw                                                           string
	} `json:"-"`
}

func (u Messaging10dlcGetEnumResponseUnion) AsEnumStringListResponse() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u Messaging10dlcGetEnumResponseUnion) AsEnumObjectListResponse() (v []map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u Messaging10dlcGetEnumResponseUnion) AsEnumObjectToStringResponse() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u Messaging10dlcGetEnumResponseUnion) AsEnumObjecToObjecttResponse() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u Messaging10dlcGetEnumResponseUnion) AsEnumPaginatedResponse() (v Messaging10dlcGetEnumResponseEnumPaginatedResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u Messaging10dlcGetEnumResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *Messaging10dlcGetEnumResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Messaging10dlcGetEnumResponseEnumPaginatedResponse struct {
	Page         int64            `json:"page" api:"required"`
	Records      []map[string]any `json:"records" api:"required"`
	TotalRecords int64            `json:"totalRecords" api:"required"`
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
func (r Messaging10dlcGetEnumResponseEnumPaginatedResponse) RawJSON() string { return r.JSON.raw }
func (r *Messaging10dlcGetEnumResponseEnumPaginatedResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Messaging10dlcGetEnumParamsEndpoint string

const (
	Messaging10dlcGetEnumParamsEndpointMno                   Messaging10dlcGetEnumParamsEndpoint = "mno"
	Messaging10dlcGetEnumParamsEndpointOptionalAttributes    Messaging10dlcGetEnumParamsEndpoint = "optionalAttributes"
	Messaging10dlcGetEnumParamsEndpointUsecase               Messaging10dlcGetEnumParamsEndpoint = "usecase"
	Messaging10dlcGetEnumParamsEndpointVertical              Messaging10dlcGetEnumParamsEndpoint = "vertical"
	Messaging10dlcGetEnumParamsEndpointAltBusinessIDType     Messaging10dlcGetEnumParamsEndpoint = "altBusinessIdType"
	Messaging10dlcGetEnumParamsEndpointBrandIdentityStatus   Messaging10dlcGetEnumParamsEndpoint = "brandIdentityStatus"
	Messaging10dlcGetEnumParamsEndpointBrandRelationship     Messaging10dlcGetEnumParamsEndpoint = "brandRelationship"
	Messaging10dlcGetEnumParamsEndpointCampaignStatus        Messaging10dlcGetEnumParamsEndpoint = "campaignStatus"
	Messaging10dlcGetEnumParamsEndpointEntityType            Messaging10dlcGetEnumParamsEndpoint = "entityType"
	Messaging10dlcGetEnumParamsEndpointExtVettingProvider    Messaging10dlcGetEnumParamsEndpoint = "extVettingProvider"
	Messaging10dlcGetEnumParamsEndpointVettingStatus         Messaging10dlcGetEnumParamsEndpoint = "vettingStatus"
	Messaging10dlcGetEnumParamsEndpointBrandStatus           Messaging10dlcGetEnumParamsEndpoint = "brandStatus"
	Messaging10dlcGetEnumParamsEndpointOperationStatus       Messaging10dlcGetEnumParamsEndpoint = "operationStatus"
	Messaging10dlcGetEnumParamsEndpointApprovedPublicCompany Messaging10dlcGetEnumParamsEndpoint = "approvedPublicCompany"
	Messaging10dlcGetEnumParamsEndpointStockExchange         Messaging10dlcGetEnumParamsEndpoint = "stockExchange"
	Messaging10dlcGetEnumParamsEndpointVettingClass          Messaging10dlcGetEnumParamsEndpoint = "vettingClass"
)
