// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// ActionService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActionService] method instead.
type ActionService struct {
	Options []option.RequestOption
	// SIM Cards operations
	Purchase ActionPurchaseService
	// SIM Cards operations
	Register ActionRegisterService
}

// NewActionService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewActionService(opts ...option.RequestOption) (r ActionService) {
	r = ActionService{}
	r.Options = opts
	r.Purchase = NewActionPurchaseService(opts...)
	r.Register = NewActionRegisterService(opts...)
	return
}

type WirelessError struct {
	Code   string              `json:"code" api:"required"`
	Title  string              `json:"title" api:"required"`
	Detail string              `json:"detail"`
	Meta   map[string]any      `json:"meta"`
	Source WirelessErrorSource `json:"source"`
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
func (r WirelessError) RawJSON() string { return r.JSON.raw }
func (r *WirelessError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WirelessErrorSource struct {
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
func (r WirelessErrorSource) RawJSON() string { return r.JSON.raw }
func (r *WirelessErrorSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
