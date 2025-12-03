// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
)

// EnumerationService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnumerationService] method instead.
type EnumerationService struct {
	Options []option.RequestOption
}

// NewEnumerationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEnumerationService(opts ...option.RequestOption) (r EnumerationService) {
	r = EnumerationService{}
	r.Options = opts
	return
}

// Get Enum
func (r *EnumerationService) Get(ctx context.Context, endpoint EnumerationGetParamsEndpoint, opts ...option.RequestOption) (res *[]string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := fmt.Sprintf("enum/%v", endpoint)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type EnumerationGetParamsEndpoint string

const (
	EnumerationGetParamsEndpointMno                   EnumerationGetParamsEndpoint = "mno"
	EnumerationGetParamsEndpointOptionalAttributes    EnumerationGetParamsEndpoint = "optionalAttributes"
	EnumerationGetParamsEndpointUsecase               EnumerationGetParamsEndpoint = "usecase"
	EnumerationGetParamsEndpointVertical              EnumerationGetParamsEndpoint = "vertical"
	EnumerationGetParamsEndpointAltBusinessIDType     EnumerationGetParamsEndpoint = "altBusinessIdType"
	EnumerationGetParamsEndpointBrandIdentityStatus   EnumerationGetParamsEndpoint = "brandIdentityStatus"
	EnumerationGetParamsEndpointBrandRelationship     EnumerationGetParamsEndpoint = "brandRelationship"
	EnumerationGetParamsEndpointCampaignStatus        EnumerationGetParamsEndpoint = "campaignStatus"
	EnumerationGetParamsEndpointEntityType            EnumerationGetParamsEndpoint = "entityType"
	EnumerationGetParamsEndpointExtVettingProvider    EnumerationGetParamsEndpoint = "extVettingProvider"
	EnumerationGetParamsEndpointVettingStatus         EnumerationGetParamsEndpoint = "vettingStatus"
	EnumerationGetParamsEndpointBrandStatus           EnumerationGetParamsEndpoint = "brandStatus"
	EnumerationGetParamsEndpointOperationStatus       EnumerationGetParamsEndpoint = "operationStatus"
	EnumerationGetParamsEndpointApprovedPublicCompany EnumerationGetParamsEndpoint = "approvedPublicCompany"
	EnumerationGetParamsEndpointStockExchange         EnumerationGetParamsEndpoint = "stockExchange"
	EnumerationGetParamsEndpointVettingClass          EnumerationGetParamsEndpoint = "vettingClass"
)
