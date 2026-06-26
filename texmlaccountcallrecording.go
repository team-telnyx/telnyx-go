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
)

// TeXML REST Commands
//
// TexmlAccountCallRecordingService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTexmlAccountCallRecordingService] method instead.
type TexmlAccountCallRecordingService struct {
	Options []option.RequestOption
}

// NewTexmlAccountCallRecordingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTexmlAccountCallRecordingService(opts ...option.RequestOption) (r TexmlAccountCallRecordingService) {
	r = TexmlAccountCallRecordingService{}
	r.Options = opts
	return
}

// Updates recording resource for particular call.
func (r *TexmlAccountCallRecordingService) RecordingSidJson(ctx context.Context, recordingSid string, params TexmlAccountCallRecordingRecordingSidJsonParams, opts ...option.RequestOption) (res *TexmlCreateCallRecordingResponseBody, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return nil, err
	}
	if params.CallSid == "" {
		err = errors.New("missing required call_sid parameter")
		return nil, err
	}
	if recordingSid == "" {
		err = errors.New("missing required recording_sid parameter")
		return nil, err
	}
	path := fmt.Sprintf("texml/Accounts/%s/Calls/%s/Recordings/%s.json", params.AccountSid, params.CallSid, recordingSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type TexmlAccountCallRecordingRecordingSidJsonParams struct {
	AccountSid string `path:"account_sid" api:"required" json:"-"`
	CallSid    string `path:"call_sid" api:"required" json:"-"`
	// Any of "in-progress", "paused", "stopped".
	Status TexmlAccountCallRecordingRecordingSidJsonParamsStatus `json:"Status,omitzero"`
	paramObj
}

func (r TexmlAccountCallRecordingRecordingSidJsonParams) MarshalJSON() (data []byte, err error) {
	type shadow TexmlAccountCallRecordingRecordingSidJsonParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlAccountCallRecordingRecordingSidJsonParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountCallRecordingRecordingSidJsonParamsStatus string

const (
	TexmlAccountCallRecordingRecordingSidJsonParamsStatusInProgress TexmlAccountCallRecordingRecordingSidJsonParamsStatus = "in-progress"
	TexmlAccountCallRecordingRecordingSidJsonParamsStatusPaused     TexmlAccountCallRecordingRecordingSidJsonParamsStatus = "paused"
	TexmlAccountCallRecordingRecordingSidJsonParamsStatusStopped    TexmlAccountCallRecordingRecordingSidJsonParamsStatus = "stopped"
)
