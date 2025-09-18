// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// DialogflowConnectionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDialogflowConnectionService] method instead.
type DialogflowConnectionService struct {
	Options []option.RequestOption
}

// NewDialogflowConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDialogflowConnectionService(opts ...option.RequestOption) (r DialogflowConnectionService) {
	r = DialogflowConnectionService{}
	r.Options = opts
	return
}

// Save Dialogflow Credentiails to Telnyx, so it can be used with other Telnyx
// services.
func (r *DialogflowConnectionService) New(ctx context.Context, connectionID string, body DialogflowConnectionNewParams, opts ...option.RequestOption) (res *DialogflowConnectionNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return
	}
	path := fmt.Sprintf("dialogflow_connections/%s", connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Return details of the Dialogflow connection associated with the given
// CallControl connection.
func (r *DialogflowConnectionService) Get(ctx context.Context, connectionID string, opts ...option.RequestOption) (res *DialogflowConnectionGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return
	}
	path := fmt.Sprintf("dialogflow_connections/%s", connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a stored Dialogflow Connection.
func (r *DialogflowConnectionService) Update(ctx context.Context, connectionID string, body DialogflowConnectionUpdateParams, opts ...option.RequestOption) (res *DialogflowConnectionUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return
	}
	path := fmt.Sprintf("dialogflow_connections/%s", connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a stored Dialogflow Connection.
func (r *DialogflowConnectionService) Delete(ctx context.Context, connectionID string, opts ...option.RequestOption) (err error) {
	opts = append(r.Options[:], opts...)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return
	}
	path := fmt.Sprintf("dialogflow_connections/%s", connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type DialogflowConnectionNewResponse struct {
	Data DialogflowConnectionNewResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DialogflowConnectionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *DialogflowConnectionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DialogflowConnectionNewResponseData struct {
	// Uniquely identifies a Telnyx application (Call Control).
	ConnectionID string `json:"connection_id"`
	// The id of a configured conversation profile on your Dialogflow account. (If you
	// use Dialogflow CX, this param is required)
	ConversationProfileID string `json:"conversation_profile_id"`
	// Which Dialogflow environment will be used.
	Environment string `json:"environment"`
	RecordType  string `json:"record_type"`
	// The JSON map to connect your Dialoglow account.
	ServiceAccount string `json:"service_account"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectionID          respjson.Field
		ConversationProfileID respjson.Field
		Environment           respjson.Field
		RecordType            respjson.Field
		ServiceAccount        respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DialogflowConnectionNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *DialogflowConnectionNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DialogflowConnectionGetResponse struct {
	Data DialogflowConnectionGetResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DialogflowConnectionGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DialogflowConnectionGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DialogflowConnectionGetResponseData struct {
	// Uniquely identifies a Telnyx application (Call Control).
	ConnectionID string `json:"connection_id"`
	// The id of a configured conversation profile on your Dialogflow account. (If you
	// use Dialogflow CX, this param is required)
	ConversationProfileID string `json:"conversation_profile_id"`
	// Which Dialogflow environment will be used.
	Environment string `json:"environment"`
	RecordType  string `json:"record_type"`
	// The JSON map to connect your Dialoglow account.
	ServiceAccount string `json:"service_account"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectionID          respjson.Field
		ConversationProfileID respjson.Field
		Environment           respjson.Field
		RecordType            respjson.Field
		ServiceAccount        respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DialogflowConnectionGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *DialogflowConnectionGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DialogflowConnectionUpdateResponse struct {
	Data DialogflowConnectionUpdateResponseData `json:"data,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DialogflowConnectionUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *DialogflowConnectionUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DialogflowConnectionUpdateResponseData struct {
	// Uniquely identifies a Telnyx application (Call Control).
	ConnectionID string `json:"connection_id"`
	// The id of a configured conversation profile on your Dialogflow account. (If you
	// use Dialogflow CX, this param is required)
	ConversationProfileID string `json:"conversation_profile_id"`
	// Which Dialogflow environment will be used.
	Environment string `json:"environment"`
	RecordType  string `json:"record_type"`
	// The JSON map to connect your Dialoglow account.
	ServiceAccount string `json:"service_account"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectionID          respjson.Field
		ConversationProfileID respjson.Field
		Environment           respjson.Field
		RecordType            respjson.Field
		ServiceAccount        respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DialogflowConnectionUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *DialogflowConnectionUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DialogflowConnectionNewParams struct {
	// The JSON map to connect your Dialoglow account.
	ServiceAccount map[string]any `json:"service_account,omitzero,required"`
	// The id of a configured conversation profile on your Dialogflow account. (If you
	// use Dialogflow CX, this param is required)
	ConversationProfileID param.Opt[string] `json:"conversation_profile_id,omitzero"`
	// Which Dialogflow environment will be used.
	Environment param.Opt[string] `json:"environment,omitzero"`
	// The region of your agent is. (If you use Dialogflow CX, this param is required)
	Location param.Opt[string] `json:"location,omitzero"`
	// Determine which Dialogflow will be used.
	//
	// Any of "cx", "es".
	DialogflowAPI DialogflowConnectionNewParamsDialogflowAPI `json:"dialogflow_api,omitzero"`
	paramObj
}

func (r DialogflowConnectionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DialogflowConnectionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DialogflowConnectionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Determine which Dialogflow will be used.
type DialogflowConnectionNewParamsDialogflowAPI string

const (
	DialogflowConnectionNewParamsDialogflowAPICx DialogflowConnectionNewParamsDialogflowAPI = "cx"
	DialogflowConnectionNewParamsDialogflowAPIEs DialogflowConnectionNewParamsDialogflowAPI = "es"
)

type DialogflowConnectionUpdateParams struct {
	// The JSON map to connect your Dialoglow account.
	ServiceAccount map[string]any `json:"service_account,omitzero,required"`
	// The id of a configured conversation profile on your Dialogflow account. (If you
	// use Dialogflow CX, this param is required)
	ConversationProfileID param.Opt[string] `json:"conversation_profile_id,omitzero"`
	// Which Dialogflow environment will be used.
	Environment param.Opt[string] `json:"environment,omitzero"`
	// The region of your agent is. (If you use Dialogflow CX, this param is required)
	Location param.Opt[string] `json:"location,omitzero"`
	// Determine which Dialogflow will be used.
	//
	// Any of "cx", "es".
	DialogflowAPI DialogflowConnectionUpdateParamsDialogflowAPI `json:"dialogflow_api,omitzero"`
	paramObj
}

func (r DialogflowConnectionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DialogflowConnectionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DialogflowConnectionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Determine which Dialogflow will be used.
type DialogflowConnectionUpdateParamsDialogflowAPI string

const (
	DialogflowConnectionUpdateParamsDialogflowAPICx DialogflowConnectionUpdateParamsDialogflowAPI = "cx"
	DialogflowConnectionUpdateParamsDialogflowAPIEs DialogflowConnectionUpdateParamsDialogflowAPI = "es"
)
