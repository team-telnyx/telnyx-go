// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Create and manage agent inboxes, retrieve inbound messages and threads, and
// reply to or forward messages.
//
// EmailInboxThreadLabelService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailInboxThreadLabelService] method instead.
type EmailInboxThreadLabelService struct {
	Options []option.RequestOption
}

// NewEmailInboxThreadLabelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailInboxThreadLabelService(opts ...option.RequestOption) (r EmailInboxThreadLabelService) {
	r = EmailInboxThreadLabelService{}
	r.Options = opts
	return
}

// Adds one or more mutable labels to a thread, letting an agent mark a whole
// conversation (for example `needs_review`) without labelling each message
// individually.
//
// Thread labels are independent of message labels: labelling a thread does not
// label its messages, and labelling a message does not label its thread.
// Idempotent and case-sensitive.
func (r *EmailInboxThreadLabelService) New(ctx context.Context, threadID string, params EmailInboxThreadLabelNewParams, opts ...option.RequestOption) (res *EmailInboxThreadLabelNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/threads/%s/labels", params.InboxID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Removes one or more labels from a thread. Idempotent — removing a label the
// thread does not carry is a no-op and still returns 200.
func (r *EmailInboxThreadLabelService) DeleteAll(ctx context.Context, threadID string, params EmailInboxThreadLabelDeleteAllParams, opts ...option.RequestOption) (res *EmailInboxThreadLabelDeleteAllResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	if threadID == "" {
		err = errors.New("missing required thread_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/threads/%s/labels", params.InboxID, threadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

type EmailInboxThreadLabelNewResponse struct {
	Data EmailInboxThreadLabelNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailInboxThreadLabelNewResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailInboxThreadLabelNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxThreadLabelNewResponseData struct {
	ID     string   `json:"id" api:"required" format:"uuid"`
	Labels []string `json:"labels" api:"required"`
	// Any of "email_thread".
	RecordType string `json:"record_type" api:"required"`
	InboxID    string `json:"inbox_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Labels      respjson.Field
		RecordType  respjson.Field
		InboxID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailInboxThreadLabelNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailInboxThreadLabelNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxThreadLabelDeleteAllResponse struct {
	Data EmailInboxThreadLabelDeleteAllResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailInboxThreadLabelDeleteAllResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailInboxThreadLabelDeleteAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxThreadLabelDeleteAllResponseData struct {
	ID     string   `json:"id" api:"required" format:"uuid"`
	Labels []string `json:"labels" api:"required"`
	// Any of "email_thread".
	RecordType string `json:"record_type" api:"required"`
	InboxID    string `json:"inbox_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Labels      respjson.Field
		RecordType  respjson.Field
		InboxID     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailInboxThreadLabelDeleteAllResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailInboxThreadLabelDeleteAllResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxThreadLabelNewParams struct {
	InboxID string `path:"inbox_id" api:"required" format:"uuid" json:"-"`
	// Labels to add or remove. Both operations are idempotent set operations, so a
	// retried request converges instead of failing.
	LabelMutationRequest LabelMutationRequestParam
	paramObj
}

func (r EmailInboxThreadLabelNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.LabelMutationRequest)
}
func (r *EmailInboxThreadLabelNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxThreadLabelDeleteAllParams struct {
	InboxID string `path:"inbox_id" api:"required" format:"uuid" json:"-"`
	// Labels to add or remove. Both operations are idempotent set operations, so a
	// retried request converges instead of failing.
	LabelMutationRequest LabelMutationRequestParam
	paramObj
}

func (r EmailInboxThreadLabelDeleteAllParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.LabelMutationRequest)
}
func (r *EmailInboxThreadLabelDeleteAllParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
