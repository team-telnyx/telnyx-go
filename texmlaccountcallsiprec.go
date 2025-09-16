// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// TexmlAccountCallSiprecService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTexmlAccountCallSiprecService] method instead.
type TexmlAccountCallSiprecService struct {
	Options []option.RequestOption
}

// NewTexmlAccountCallSiprecService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTexmlAccountCallSiprecService(opts ...option.RequestOption) (r TexmlAccountCallSiprecService) {
	r = TexmlAccountCallSiprecService{}
	r.Options = opts
	return
}

// Updates siprec session identified by siprec_sid.
func (r *TexmlAccountCallSiprecService) SiprecSidJson(ctx context.Context, siprecSid string, params TexmlAccountCallSiprecSiprecSidJsonParams, opts ...option.RequestOption) (res *TexmlAccountCallSiprecSiprecSidJsonResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if params.CallSid == "" {
		err = errors.New("missing required call_sid parameter")
		return
	}
	if siprecSid == "" {
		err = errors.New("missing required siprec_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Calls/%s/Siprec/%s.json", params.AccountSid, params.CallSid, siprecSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type TexmlAccountCallSiprecSiprecSidJsonResponse struct {
	// The id of the account the resource belongs to.
	AccountSid string `json:"account_sid"`
	// The id of the call the resource belongs to.
	CallSid string `json:"call_sid"`
	// The date and time the siprec session was last updated.
	DateUpdated string `json:"date_updated"`
	// The error code of the siprec session.
	ErrorCode string `json:"error_code"`
	// The SID of the siprec session.
	Sid string `json:"sid"`
	// The status of the siprec session.
	//
	// Any of "in-progress", "stopped".
	Status TexmlAccountCallSiprecSiprecSidJsonResponseStatus `json:"status"`
	// The URI of the siprec session.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid  respjson.Field
		CallSid     respjson.Field
		DateUpdated respjson.Field
		ErrorCode   respjson.Field
		Sid         respjson.Field
		Status      respjson.Field
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountCallSiprecSiprecSidJsonResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountCallSiprecSiprecSidJsonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the siprec session.
type TexmlAccountCallSiprecSiprecSidJsonResponseStatus string

const (
	TexmlAccountCallSiprecSiprecSidJsonResponseStatusInProgress TexmlAccountCallSiprecSiprecSidJsonResponseStatus = "in-progress"
	TexmlAccountCallSiprecSiprecSidJsonResponseStatusStopped    TexmlAccountCallSiprecSiprecSidJsonResponseStatus = "stopped"
)

type TexmlAccountCallSiprecSiprecSidJsonParams struct {
	AccountSid string `path:"account_sid,required" json:"-"`
	CallSid    string `path:"call_sid,required" json:"-"`
	// The new status of the resource. Specifying `stopped` will end the siprec
	// session.
	//
	// Any of "stopped".
	Status TexmlAccountCallSiprecSiprecSidJsonParamsStatus `json:"Status,omitzero"`
	paramObj
}

func (r TexmlAccountCallSiprecSiprecSidJsonParams) MarshalJSON() (data []byte, err error) {
	type shadow TexmlAccountCallSiprecSiprecSidJsonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlAccountCallSiprecSiprecSidJsonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The new status of the resource. Specifying `stopped` will end the siprec
// session.
type TexmlAccountCallSiprecSiprecSidJsonParamsStatus string

const (
	TexmlAccountCallSiprecSiprecSidJsonParamsStatusStopped TexmlAccountCallSiprecSiprecSidJsonParamsStatus = "stopped"
)
