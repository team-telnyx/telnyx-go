// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// TexmlAccountRecordingJsonService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTexmlAccountRecordingJsonService] method instead.
type TexmlAccountRecordingJsonService struct {
	Options []option.RequestOption
}

// NewTexmlAccountRecordingJsonService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewTexmlAccountRecordingJsonService(opts ...option.RequestOption) (r TexmlAccountRecordingJsonService) {
	r = TexmlAccountRecordingJsonService{}
	r.Options = opts
	return
}

// Deletes recording resource identified by recording id.
func (r *TexmlAccountRecordingJsonService) DeleteRecordingSidJson(ctx context.Context, recordingSid string, body TexmlAccountRecordingJsonDeleteRecordingSidJsonParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if recordingSid == "" {
		err = errors.New("missing required recording_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Recordings/%s.json", body.AccountSid, recordingSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Returns recording resource identified by recording id.
func (r *TexmlAccountRecordingJsonService) GetRecordingSidJson(ctx context.Context, recordingSid string, query TexmlAccountRecordingJsonGetRecordingSidJsonParams, opts ...option.RequestOption) (res *TexmlGetCallRecordingResponseBody, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if recordingSid == "" {
		err = errors.New("missing required recording_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Recordings/%s.json", query.AccountSid, recordingSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type TexmlAccountRecordingJsonDeleteRecordingSidJsonParams struct {
	AccountSid string `path:"account_sid,required" json:"-"`
	paramObj
}

type TexmlAccountRecordingJsonGetRecordingSidJsonParams struct {
	AccountSid string `path:"account_sid,required" json:"-"`
	paramObj
}
