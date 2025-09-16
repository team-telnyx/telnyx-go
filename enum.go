// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
)

// EnumService contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnumService] method instead.
type EnumService struct {
	Options []option.RequestOption
}

// NewEnumService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewEnumService(opts ...option.RequestOption) (r EnumService) {
	r = EnumService{}
	r.Options = opts
	return
}

// Get Enum
func (r *EnumService) Get(ctx context.Context, endpoint EnumGetParamsEndpoint, opts ...option.RequestOption) (res *[]any, err error) {
	opts = append(r.Options[:], opts...)
	path := fmt.Sprintf("enum/%v", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EnumGetParamsEndpoint string

const (
	EnumGetParamsEndpointMno                   EnumGetParamsEndpoint = "mno"
	EnumGetParamsEndpointOptionalAttributes    EnumGetParamsEndpoint = "optionalAttributes"
	EnumGetParamsEndpointUsecase               EnumGetParamsEndpoint = "usecase"
	EnumGetParamsEndpointVertical              EnumGetParamsEndpoint = "vertical"
	EnumGetParamsEndpointAltBusinessIDType     EnumGetParamsEndpoint = "altBusinessIdType"
	EnumGetParamsEndpointBrandIdentityStatus   EnumGetParamsEndpoint = "brandIdentityStatus"
	EnumGetParamsEndpointBrandRelationship     EnumGetParamsEndpoint = "brandRelationship"
	EnumGetParamsEndpointCampaignStatus        EnumGetParamsEndpoint = "campaignStatus"
	EnumGetParamsEndpointEntityType            EnumGetParamsEndpoint = "entityType"
	EnumGetParamsEndpointExtVettingProvider    EnumGetParamsEndpoint = "extVettingProvider"
	EnumGetParamsEndpointVettingStatus         EnumGetParamsEndpoint = "vettingStatus"
	EnumGetParamsEndpointBrandStatus           EnumGetParamsEndpoint = "brandStatus"
	EnumGetParamsEndpointOperationStatus       EnumGetParamsEndpoint = "operationStatus"
	EnumGetParamsEndpointApprovedPublicCompany EnumGetParamsEndpoint = "approvedPublicCompany"
	EnumGetParamsEndpointStockExchange         EnumGetParamsEndpoint = "stockExchange"
	EnumGetParamsEndpointVettingClass          EnumGetParamsEndpoint = "vettingClass"
)
