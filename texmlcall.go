// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// TeXML REST Commands
//
// TexmlCallService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTexmlCallService] method instead.
type TexmlCallService struct {
	Options []option.RequestOption
}

// NewTexmlCallService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTexmlCallService(opts ...option.RequestOption) (r TexmlCallService) {
	r = TexmlCallService{}
	r.Options = opts
	return
}

// Initiate an outbound TeXML call using a TeXML application connection ID, not an
// account SID. Request parameter names are case-sensitive. From and To are
// required; Texml supplies inline instructions and Url overrides the application
// XML request URL. When neither is supplied, the application configuration
// supplies the instructions. The response is a flat call object without a data
// wrapper.
func (r *TexmlCallService) New(ctx context.Context, connectionID string, body TexmlCallNewParams, opts ...option.RequestOption) (res *TexmlCallNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("texml/calls/%s", url.PathEscape(connectionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type TexmlCallNewResponse struct {
	// The call control ID of the created call.
	CallSid string `json:"call_sid" api:"required"`
	// The caller address.
	From string `json:"from" api:"required"`
	// The initial status of the outbound call.
	//
	// Any of "queued".
	Status TexmlCallNewResponseStatus `json:"status" api:"required"`
	// The called address.
	To string `json:"to" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CallSid     respjson.Field
		From        respjson.Field
		Status      respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlCallNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlCallNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The initial status of the outbound call.
type TexmlCallNewResponseStatus string

const (
	TexmlCallNewResponseStatusQueued TexmlCallNewResponseStatus = "queued"
)

type TexmlCallNewParams struct {
	// The E.164-formatted phone number or SIP URI to present as the caller.
	From string `json:"From" api:"required"`
	// The E.164-formatted phone number or SIP URI to call.
	To string `json:"To" api:"required"`
	// Inline TeXML instructions to execute when the call is answered.
	Texml param.Opt[string] `json:"Texml,omitzero"`
	// The URL from which to retrieve TeXML instructions. Overrides the TeXML
	// application XML request URL.
	URL param.Opt[string] `json:"Url,omitzero"`
	// HTTP method used to retrieve TeXML instructions from Url.
	//
	// Any of "GET", "POST".
	Method TexmlCallNewParamsMethod `json:"Method,omitzero"`
	paramObj
}

func (r TexmlCallNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TexmlCallNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlCallNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// HTTP method used to retrieve TeXML instructions from Url.
type TexmlCallNewParamsMethod string

const (
	TexmlCallNewParamsMethodGet  TexmlCallNewParamsMethod = "GET"
	TexmlCallNewParamsMethodPost TexmlCallNewParamsMethod = "POST"
)
