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
)

// Call Recordings operations.
//
// RecordingActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRecordingActionService] method instead.
type RecordingActionService struct {
	Options []option.RequestOption
}

// NewRecordingActionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewRecordingActionService(opts ...option.RequestOption) (r RecordingActionService) {
	r = RecordingActionService{}
	r.Options = opts
	return
}

// Permanently deletes a list of call recordings.
func (r *RecordingActionService) Delete(ctx context.Context, body RecordingActionDeleteParams, opts ...option.RequestOption) (res *RecordingActionDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "recordings/actions/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type RecordingActionDeleteResponse struct {
	// Any of "ok".
	Status RecordingActionDeleteResponseStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RecordingActionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *RecordingActionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RecordingActionDeleteResponseStatus string

const (
	RecordingActionDeleteResponseStatusOk RecordingActionDeleteResponseStatus = "ok"
)

type RecordingActionDeleteParams struct {
	// List of call recording IDs to delete.
	IDs []string `json:"ids,omitzero" api:"required"`
	paramObj
}

func (r RecordingActionDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow RecordingActionDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecordingActionDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
