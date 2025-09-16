// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/param"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// TexmlAccountCallStreamService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTexmlAccountCallStreamService] method instead.
type TexmlAccountCallStreamService struct {
	Options []option.RequestOption
}

// NewTexmlAccountCallStreamService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTexmlAccountCallStreamService(opts ...option.RequestOption) (r TexmlAccountCallStreamService) {
	r = TexmlAccountCallStreamService{}
	r.Options = opts
	return
}

// Updates streaming resource for particular call.
func (r *TexmlAccountCallStreamService) StreamingSidJson(ctx context.Context, streamingSid string, params TexmlAccountCallStreamStreamingSidJsonParams, opts ...option.RequestOption) (res *TexmlAccountCallStreamStreamingSidJsonResponse, err error) {
	opts = append(r.Options[:], opts...)
	if params.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if params.CallSid == "" {
		err = errors.New("missing required call_sid parameter")
		return
	}
	if streamingSid == "" {
		err = errors.New("missing required streaming_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Calls/%s/Streams/%s.json", params.AccountSid, params.CallSid, streamingSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type TexmlAccountCallStreamStreamingSidJsonResponse struct {
	AccountSid  string    `json:"account_sid"`
	CallSid     string    `json:"call_sid"`
	DateUpdated time.Time `json:"date_updated" format:"date-time"`
	// Identifier of a resource.
	Sid string `json:"sid"`
	// The status of the streaming.
	//
	// Any of "stopped".
	Status TexmlAccountCallStreamStreamingSidJsonResponseStatus `json:"status"`
	// The relative URI for this streaming resource.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid  respjson.Field
		CallSid     respjson.Field
		DateUpdated respjson.Field
		Sid         respjson.Field
		Status      respjson.Field
		Uri         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountCallStreamStreamingSidJsonResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountCallStreamStreamingSidJsonResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the streaming.
type TexmlAccountCallStreamStreamingSidJsonResponseStatus string

const (
	TexmlAccountCallStreamStreamingSidJsonResponseStatusStopped TexmlAccountCallStreamStreamingSidJsonResponseStatus = "stopped"
)

type TexmlAccountCallStreamStreamingSidJsonParams struct {
	AccountSid string `path:"account_sid,required" json:"-"`
	CallSid    string `path:"call_sid,required" json:"-"`
	// The status of the Stream you wish to update.
	//
	// Any of "stopped".
	Status TexmlAccountCallStreamStreamingSidJsonParamsStatus `json:"Status,omitzero"`
	paramObj
}

func (r TexmlAccountCallStreamStreamingSidJsonParams) MarshalJSON() (data []byte, err error) {
	type shadow TexmlAccountCallStreamStreamingSidJsonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlAccountCallStreamStreamingSidJsonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The status of the Stream you wish to update.
type TexmlAccountCallStreamStreamingSidJsonParamsStatus string

const (
	TexmlAccountCallStreamStreamingSidJsonParamsStatusStopped TexmlAccountCallStreamStreamingSidJsonParamsStatus = "stopped"
)
