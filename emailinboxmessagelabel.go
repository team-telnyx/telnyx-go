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
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Create and manage agent inboxes, retrieve inbound messages and threads, and
// reply to or forward messages.
//
// EmailInboxMessageLabelService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailInboxMessageLabelService] method instead.
type EmailInboxMessageLabelService struct {
	Options []option.RequestOption
}

// NewEmailInboxMessageLabelService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailInboxMessageLabelService(opts ...option.RequestOption) (r EmailInboxMessageLabelService) {
	r = EmailInboxMessageLabelService{}
	r.Options = opts
	return
}

// Adds one or more mutable labels to a message. Labels carry agent workflow state
// such as `spam`, `needs_review`, or `processed`.
//
// Labels are **not** the same as the send-time `tags` on outbound messages: `tags`
// are immutable and propagate to Email Detail Records and Mission Control for
// billing attribution, while labels are mailbox state that never reaches the
// reporting contract.
//
// The operation is an idempotent set union — adding a label the message already
// carries is a no-op and still returns 200. Labels are case-sensitive, and message
// labels are independent of thread labels.
func (r *EmailInboxMessageLabelService) New(ctx context.Context, messageID string, params EmailInboxMessageLabelNewParams, opts ...option.RequestOption) (res *EmailInboxMessageLabelNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/messages/%s/labels", params.InboxID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Removes one or more labels from a message. Idempotent — removing a label the
// message does not carry is a no-op and still returns 200. Removal is
// case-sensitive.
func (r *EmailInboxMessageLabelService) DeleteAll(ctx context.Context, messageID string, params EmailInboxMessageLabelDeleteAllParams, opts ...option.RequestOption) (res *EmailInboxMessageLabelDeleteAllResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/messages/%s/labels", params.InboxID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, &res, opts...)
	return res, err
}

// Labels to add or remove. Both operations are idempotent set operations, so a
// retried request converges instead of failing.
//
// The property Labels is required.
type LabelMutationRequestParam struct {
	// One or more labels. Each label is a freeform, case-sensitive string of at most
	// 255 characters; a message or thread may carry at most 50 labels. The `telnyx:`
	// prefix is a reserved system namespace and is rejected on customer writes.
	Labels []string `json:"labels,omitzero" api:"required"`
	paramObj
}

func (r LabelMutationRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow LabelMutationRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *LabelMutationRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxMessageLabelNewResponse struct {
	Data InboundMessage `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailInboxMessageLabelNewResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailInboxMessageLabelNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxMessageLabelDeleteAllResponse struct {
	Data InboundMessage `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailInboxMessageLabelDeleteAllResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailInboxMessageLabelDeleteAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxMessageLabelNewParams struct {
	InboxID string `path:"inbox_id" api:"required" format:"uuid" json:"-"`
	// Labels to add or remove. Both operations are idempotent set operations, so a
	// retried request converges instead of failing.
	LabelMutationRequest LabelMutationRequestParam
	paramObj
}

func (r EmailInboxMessageLabelNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.LabelMutationRequest)
}
func (r *EmailInboxMessageLabelNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxMessageLabelDeleteAllParams struct {
	InboxID string `path:"inbox_id" api:"required" format:"uuid" json:"-"`
	// Labels to add or remove. Both operations are idempotent set operations, so a
	// retried request converges instead of failing.
	LabelMutationRequest LabelMutationRequestParam
	paramObj
}

func (r EmailInboxMessageLabelDeleteAllParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.LabelMutationRequest)
}
func (r *EmailInboxMessageLabelDeleteAllParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
