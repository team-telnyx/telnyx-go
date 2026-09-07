// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Requirement Groups
//
// ExternalRequirementSubNumberOrderService contains methods and other services
// that help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewExternalRequirementSubNumberOrderService] method instead.
type ExternalRequirementSubNumberOrderService struct {
	Options []option.RequestOption
}

// NewExternalRequirementSubNumberOrderService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewExternalRequirementSubNumberOrderService(opts ...option.RequestOption) (r ExternalRequirementSubNumberOrderService) {
	r = ExternalRequirementSubNumberOrderService{}
	r.Options = opts
	return
}

// Returns the input fields an action requirement needs and the current requirement
// action for a sub number order. Action requirements are fulfilled by an external
// step rather than by uploading documents. Australia mobile ID verification is
// currently the only action requirement. Once a verification link has been
// generated, it is returned in `requirement_action.value`.
func (r *ExternalRequirementSubNumberOrderService) Get(ctx context.Context, subNumberOrderID string, query ExternalRequirementSubNumberOrderGetParams, opts ...option.RequestOption) (res *ExternalRequirementSubNumberOrderGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.RegulatoryRequirementID == "" {
		err = errors.New("missing required regulatory_requirement_id parameter")
		return nil, err
	}
	if subNumberOrderID == "" {
		err = errors.New("missing required sub_number_order_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("external_requirements/%s/sub_number_orders/%s", query.RegulatoryRequirementID, subNumberOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Submits the end user's details to the external verification provider and returns
// the requirement action. Australia mobile ID verification is currently the only
// action requirement. It generates a unique Onfido verification link, returned in
// `requirement_action.value`, which you share with the end user. The end user's
// `first_name` and `last_name` must be nested inside a `requirement` object;
// sending them at the top level is rejected.
func (r *ExternalRequirementSubNumberOrderService) Update(ctx context.Context, subNumberOrderID string, params ExternalRequirementSubNumberOrderUpdateParams, opts ...option.RequestOption) (res *ExternalRequirementSubNumberOrderUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.RegulatoryRequirementID == "" {
		err = errors.New("missing required regulatory_requirement_id parameter")
		return nil, err
	}
	if subNumberOrderID == "" {
		err = errors.New("missing required sub_number_order_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("external_requirements/%s/sub_number_orders/%s", params.RegulatoryRequirementID, subNumberOrderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type ExternalRequirementSubNumberOrderGetResponse struct {
	Data ExternalRequirementSubNumberOrderGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalRequirementSubNumberOrderGetResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalRequirementSubNumberOrderGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalRequirementSubNumberOrderGetResponseData struct {
	// The fields the end user must provide to fulfill this requirement.
	FieldsRequired          []ExternalRequirementSubNumberOrderGetResponseDataFieldsRequired  `json:"fields_required"`
	RegulatoryRequirementID string                                                            `json:"regulatory_requirement_id" format:"uuid"`
	RequirementAction       ExternalRequirementSubNumberOrderGetResponseDataRequirementAction `json:"requirement_action"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FieldsRequired          respjson.Field
		RegulatoryRequirementID respjson.Field
		RequirementAction       respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalRequirementSubNumberOrderGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *ExternalRequirementSubNumberOrderGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalRequirementSubNumberOrderGetResponseDataFieldsRequired struct {
	Description string `json:"description"`
	// The field name to send inside the `requirement` object on the POST.
	Name string `json:"name"`
	Type string `json:"type"`
	// The value already stored for this field, or null if not yet provided.
	Value string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Name        respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalRequirementSubNumberOrderGetResponseDataFieldsRequired) RawJSON() string {
	return r.JSON.raw
}
func (r *ExternalRequirementSubNumberOrderGetResponseDataFieldsRequired) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalRequirementSubNumberOrderGetResponseDataRequirementAction struct {
	// The type of action the end user must complete.
	Type string `json:"type"`
	// The action value. For ID verification this is the verification link URL, or null
	// until it has been generated.
	Value string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalRequirementSubNumberOrderGetResponseDataRequirementAction) RawJSON() string {
	return r.JSON.raw
}
func (r *ExternalRequirementSubNumberOrderGetResponseDataRequirementAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalRequirementSubNumberOrderUpdateResponse struct {
	Data ExternalRequirementSubNumberOrderUpdateResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalRequirementSubNumberOrderUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ExternalRequirementSubNumberOrderUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalRequirementSubNumberOrderUpdateResponseData struct {
	RegulatoryRequirementID string                                                               `json:"regulatory_requirement_id" format:"uuid"`
	RequirementAction       ExternalRequirementSubNumberOrderUpdateResponseDataRequirementAction `json:"requirement_action"`
	SubOrderID              string                                                               `json:"sub_order_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		RegulatoryRequirementID respjson.Field
		RequirementAction       respjson.Field
		SubOrderID              respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalRequirementSubNumberOrderUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *ExternalRequirementSubNumberOrderUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalRequirementSubNumberOrderUpdateResponseDataRequirementAction struct {
	Type string `json:"type"`
	// For Australia mobile ID verification, the unique Onfido verification link to
	// share with the end user.
	Value string `json:"value" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalRequirementSubNumberOrderUpdateResponseDataRequirementAction) RawJSON() string {
	return r.JSON.raw
}
func (r *ExternalRequirementSubNumberOrderUpdateResponseDataRequirementAction) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalRequirementSubNumberOrderGetParams struct {
	RegulatoryRequirementID string `path:"regulatory_requirement_id" api:"required" format:"uuid" json:"-"`
	paramObj
}

type ExternalRequirementSubNumberOrderUpdateParams struct {
	RegulatoryRequirementID string `path:"regulatory_requirement_id" api:"required" format:"uuid" json:"-"`
	// The end user's identity details for the action requirement. Australia mobile ID
	// verification is currently the only action requirement. It requires `first_name`
	// and `last_name`, the same fields the corresponding GET lists in
	// `fields_required`.
	Requirement ExternalRequirementSubNumberOrderUpdateParamsRequirement `json:"requirement,omitzero" api:"required"`
	paramObj
}

func (r ExternalRequirementSubNumberOrderUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ExternalRequirementSubNumberOrderUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalRequirementSubNumberOrderUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The end user's identity details for the action requirement. Australia mobile ID
// verification is currently the only action requirement. It requires `first_name`
// and `last_name`, the same fields the corresponding GET lists in
// `fields_required`.
//
// The properties FirstName, LastName are required.
type ExternalRequirementSubNumberOrderUpdateParamsRequirement struct {
	// The end user's first name.
	FirstName string `json:"first_name" api:"required"`
	// The end user's last name.
	LastName string `json:"last_name" api:"required"`
	paramObj
}

func (r ExternalRequirementSubNumberOrderUpdateParamsRequirement) MarshalJSON() (data []byte, err error) {
	type shadow ExternalRequirementSubNumberOrderUpdateParamsRequirement
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExternalRequirementSubNumberOrderUpdateParamsRequirement) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
