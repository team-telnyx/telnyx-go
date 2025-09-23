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
)

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
func (r *RecordingActionService) Delete(ctx context.Context, body RecordingActionDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "")}, opts...)
	path := "recordings/actions/delete"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, nil, opts...)
	return
}

type RecordingActionDeleteParams struct {
	// List of call recording IDs to delete.
	IDs []string `json:"ids,omitzero,required"`
	paramObj
}

func (r RecordingActionDeleteParams) MarshalJSON() (data []byte, err error) {
	type shadow RecordingActionDeleteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RecordingActionDeleteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
