// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// Number10dlcBrandExternalVettingService contains methods and other services that
// help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumber10dlcBrandExternalVettingService] method instead.
type Number10dlcBrandExternalVettingService struct {
	Options []option.RequestOption
}

// NewNumber10dlcBrandExternalVettingService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewNumber10dlcBrandExternalVettingService(opts ...option.RequestOption) (r Number10dlcBrandExternalVettingService) {
	r = Number10dlcBrandExternalVettingService{}
	r.Options = opts
	return
}

// Get list of valid external vetting record for a given brand
func (r *Number10dlcBrandExternalVettingService) List(ctx context.Context, brandID string, opts ...option.RequestOption) (res *[]Number10dlcBrandExternalVettingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if brandID == "" {
		err = errors.New("missing required brandId parameter")
		return
	}
	path := fmt.Sprintf("10dlc/brand/%s/externalVetting", brandID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// This operation can be used to import an external vetting record from a
// TCR-approved vetting provider. If the vetting provider confirms validity of the
// record, it will be saved with the brand and will be considered for future
// campaign qualification.
func (r *Number10dlcBrandExternalVettingService) Imports(ctx context.Context, brandID string, body Number10dlcBrandExternalVettingImportsParams, opts ...option.RequestOption) (res *Number10dlcBrandExternalVettingImportsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if brandID == "" {
		err = errors.New("missing required brandId parameter")
		return
	}
	path := fmt.Sprintf("10dlc/brand/%s/externalVetting", brandID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Order new external vetting for a brand
func (r *Number10dlcBrandExternalVettingService) Order(ctx context.Context, brandID string, body Number10dlcBrandExternalVettingOrderParams, opts ...option.RequestOption) (res *Number10dlcBrandExternalVettingOrderResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if brandID == "" {
		err = errors.New("missing required brandId parameter")
		return
	}
	path := fmt.Sprintf("10dlc/brand/%s/externalVetting", brandID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Number10dlcBrandExternalVettingListResponse struct {
	// Vetting submission date. This is the date when the vetting request is generated
	// in ISO 8601 format.
	CreateDate string `json:"createDate"`
	// External vetting provider ID for the brand.
	EvpID string `json:"evpId"`
	// Vetting effective date. This is the date when vetting was completed, or the
	// starting effective date in ISO 8601 format. If this date is missing, then the
	// vetting was not complete or not valid.
	VettedDate string `json:"vettedDate"`
	// Identifies the vetting classification.
	VettingClass string `json:"vettingClass"`
	// Unique ID that identifies a vetting transaction performed by a vetting provider.
	// This ID is provided by the vetting provider at time of vetting.
	VettingID string `json:"vettingId"`
	// Vetting score ranging from 0-100.
	VettingScore int64 `json:"vettingScore"`
	// Required by some providers for vetting record confirmation.
	VettingToken string `json:"vettingToken"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateDate   respjson.Field
		EvpID        respjson.Field
		VettedDate   respjson.Field
		VettingClass respjson.Field
		VettingID    respjson.Field
		VettingScore respjson.Field
		VettingToken respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcBrandExternalVettingListResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcBrandExternalVettingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcBrandExternalVettingImportsResponse struct {
	// Vetting submission date. This is the date when the vetting request is generated
	// in ISO 8601 format.
	CreateDate string `json:"createDate"`
	// External vetting provider ID for the brand.
	EvpID string `json:"evpId"`
	// Vetting effective date. This is the date when vetting was completed, or the
	// starting effective date in ISO 8601 format. If this date is missing, then the
	// vetting was not complete or not valid.
	VettedDate string `json:"vettedDate"`
	// Identifies the vetting classification.
	VettingClass string `json:"vettingClass"`
	// Unique ID that identifies a vetting transaction performed by a vetting provider.
	// This ID is provided by the vetting provider at time of vetting.
	VettingID string `json:"vettingId"`
	// Vetting score ranging from 0-100.
	VettingScore int64 `json:"vettingScore"`
	// Required by some providers for vetting record confirmation.
	VettingToken string `json:"vettingToken"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateDate   respjson.Field
		EvpID        respjson.Field
		VettedDate   respjson.Field
		VettingClass respjson.Field
		VettingID    respjson.Field
		VettingScore respjson.Field
		VettingToken respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcBrandExternalVettingImportsResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcBrandExternalVettingImportsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcBrandExternalVettingOrderResponse struct {
	// Vetting submission date. This is the date when the vetting request is generated
	// in ISO 8601 format.
	CreateDate string `json:"createDate"`
	// External vetting provider ID for the brand.
	EvpID string `json:"evpId"`
	// Vetting effective date. This is the date when vetting was completed, or the
	// starting effective date in ISO 8601 format. If this date is missing, then the
	// vetting was not complete or not valid.
	VettedDate string `json:"vettedDate"`
	// Identifies the vetting classification.
	VettingClass string `json:"vettingClass"`
	// Unique ID that identifies a vetting transaction performed by a vetting provider.
	// This ID is provided by the vetting provider at time of vetting.
	VettingID string `json:"vettingId"`
	// Vetting score ranging from 0-100.
	VettingScore int64 `json:"vettingScore"`
	// Required by some providers for vetting record confirmation.
	VettingToken string `json:"vettingToken"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreateDate   respjson.Field
		EvpID        respjson.Field
		VettedDate   respjson.Field
		VettingClass respjson.Field
		VettingID    respjson.Field
		VettingScore respjson.Field
		VettingToken respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Number10dlcBrandExternalVettingOrderResponse) RawJSON() string { return r.JSON.raw }
func (r *Number10dlcBrandExternalVettingOrderResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcBrandExternalVettingImportsParams struct {
	// External vetting provider ID for the brand.
	EvpID string `json:"evpId,required"`
	// Unique ID that identifies a vetting transaction performed by a vetting provider.
	// This ID is provided by the vetting provider at time of vetting.
	VettingID string `json:"vettingId,required"`
	// Required by some providers for vetting record confirmation.
	VettingToken param.Opt[string] `json:"vettingToken,omitzero"`
	paramObj
}

func (r Number10dlcBrandExternalVettingImportsParams) MarshalJSON() (data []byte, err error) {
	type shadow Number10dlcBrandExternalVettingImportsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Number10dlcBrandExternalVettingImportsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Number10dlcBrandExternalVettingOrderParams struct {
	// External vetting provider ID for the brand.
	EvpID string `json:"evpId,required"`
	// Identifies the vetting classification.
	VettingClass string `json:"vettingClass,required"`
	paramObj
}

func (r Number10dlcBrandExternalVettingOrderParams) MarshalJSON() (data []byte, err error) {
	type shadow Number10dlcBrandExternalVettingOrderParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *Number10dlcBrandExternalVettingOrderParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
