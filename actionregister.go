// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
	"github.com/team-telnyx/telnyx-go/shared"
)

// ActionRegisterService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewActionRegisterService] method instead.
type ActionRegisterService struct {
	Options []option.RequestOption
}

// NewActionRegisterService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewActionRegisterService(opts ...option.RequestOption) (r ActionRegisterService) {
	r = ActionRegisterService{}
	r.Options = opts
	return
}

// Register the SIM cards associated with the provided registration codes to the
// current user's account.<br/><br/> If <code>sim_card_group_id</code> is provided,
// the SIM cards will be associated with that group. Otherwise, the default group
// for the current user will be used.<br/><br/>
func (r *ActionRegisterService) New(ctx context.Context, body ActionRegisterNewParams, opts ...option.RequestOption) (res *ActionRegisterNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "actions/register/sim_cards"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ActionRegisterNewResponse struct {
	// Successfully registered SIM cards.
	Data   []shared.SimpleSimCard           `json:"data"`
	Errors []ActionRegisterNewResponseError `json:"errors"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Errors      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ActionRegisterNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ActionRegisterNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionRegisterNewResponseError struct {
	Code   string                               `json:"code,required"`
	Title  string                               `json:"title,required"`
	Detail string                               `json:"detail"`
	Meta   map[string]any                       `json:"meta"`
	Source ActionRegisterNewResponseErrorSource `json:"source"`
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
func (r ActionRegisterNewResponseError) RawJSON() string { return r.JSON.raw }
func (r *ActionRegisterNewResponseError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionRegisterNewResponseErrorSource struct {
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
func (r ActionRegisterNewResponseErrorSource) RawJSON() string { return r.JSON.raw }
func (r *ActionRegisterNewResponseErrorSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ActionRegisterNewParams struct {
	RegistrationCodes []string `json:"registration_codes,omitzero,required"`
	// The group SIMCardGroup identification. This attribute can be <code>null</code>
	// when it's present in an associated resource.
	SimCardGroupID param.Opt[string] `json:"sim_card_group_id,omitzero" format:"uuid"`
	// Status on which the SIM card will be set after being successful registered.
	//
	// Any of "enabled", "disabled", "standby".
	Status ActionRegisterNewParamsStatus `json:"status,omitzero"`
	// Searchable tags associated with the SIM card
	Tags []string `json:"tags,omitzero"`
	paramObj
}

func (r ActionRegisterNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ActionRegisterNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ActionRegisterNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status on which the SIM card will be set after being successful registered.
type ActionRegisterNewParamsStatus string

const (
	ActionRegisterNewParamsStatusEnabled  ActionRegisterNewParamsStatus = "enabled"
	ActionRegisterNewParamsStatusDisabled ActionRegisterNewParamsStatus = "disabled"
	ActionRegisterNewParamsStatusStandby  ActionRegisterNewParamsStatus = "standby"
)
