// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared"
)

// ActionPurchaseService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActionPurchaseService] method instead.
type ActionPurchaseService struct {
	Options []option.RequestOption
}

// NewActionPurchaseService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActionPurchaseService(opts ...option.RequestOption) (r ActionPurchaseService) {
	r = ActionPurchaseService{}
	r.Options = opts
	return
}

// Purchases and registers the specified amount of eSIMs to the current user's
// account.<br/><br/> If <code>sim_card_group_id</code> is provided, the eSIMs will
// be associated with that group. Otherwise, the default group for the current user
// will be used.<br/><br/>
func (r *ActionPurchaseService) New(ctx context.Context, body ActionPurchaseNewParams, opts ...option.RequestOption) (res *ActionPurchaseNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "actions/purchase/esims"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ActionPurchaseNewResponse struct {
	// Successfully registered SIM cards.
	Data   []shared.SimpleSimCard           `json:"data"`
	Errors []ActionPurchaseNewResponseError `json:"errors"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Errors      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionPurchaseNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ActionPurchaseNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionPurchaseNewResponseError struct {
	Code   string                               `json:"code" api:"required"`
	Title  string                               `json:"title" api:"required"`
	Detail string                               `json:"detail"`
	Meta   map[string]any                       `json:"meta"`
	Source ActionPurchaseNewResponseErrorSource `json:"source"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Title       respjson.Field
		Detail      respjson.Field
		Meta        respjson.Field
		Source      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionPurchaseNewResponseError) RawJSON() string { return r.JSON.raw }
func (r *ActionPurchaseNewResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionPurchaseNewResponseErrorSource struct {
	// Indicates which query parameter caused the error.
	Parameter string `json:"parameter"`
	// JSON pointer (RFC6901) to the offending entity.
	Pointer string `json:"pointer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parameter   respjson.Field
		Pointer     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionPurchaseNewResponseErrorSource) RawJSON() string { return r.JSON.raw }
func (r *ActionPurchaseNewResponseErrorSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionPurchaseNewParams struct {
	// The amount of eSIMs to be purchased.
	Amount int64 `json:"amount" api:"required"`
	// Type of product to be purchased, specify "whitelabel" to use a custom SPN
	Product param.Opt[string] `json:"product,omitzero"`
	// The group SIMCardGroup identification. This attribute can be <code>null</code>
	// when it's present in an associated resource.
	SimCardGroupID param.Opt[string] `json:"sim_card_group_id,omitzero" format:"uuid"`
	// Service Provider Name (SPN) for the Whitelabel eSIM product. It will be
	// displayed as the mobile service name by operating systems of smartphones. This
	// parameter must only contain letters, numbers and whitespaces.
	WhitelabelName param.Opt[string] `json:"whitelabel_name,omitzero"`
	// Status on which the SIM cards will be set after being successfully registered.
	//
	// Any of "enabled", "disabled", "standby".
	Status ActionPurchaseNewParamsStatus `json:"status,omitzero"`
	// Searchable tags associated with the SIM cards
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ActionPurchaseNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ActionPurchaseNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionPurchaseNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status on which the SIM cards will be set after being successfully registered.
type ActionPurchaseNewParamsStatus string

const (
	ActionPurchaseNewParamsStatusEnabled  ActionPurchaseNewParamsStatus = "enabled"
	ActionPurchaseNewParamsStatusDisabled ActionPurchaseNewParamsStatus = "disabled"
	ActionPurchaseNewParamsStatusStandby  ActionPurchaseNewParamsStatus = "standby"
)
