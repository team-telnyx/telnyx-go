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
)

// Create and manage agent inboxes, retrieve inbound messages and threads, and
// reply to or forward messages.
//
// EmailInboxMessageActionService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailInboxMessageActionService] method instead.
type EmailInboxMessageActionService struct {
	Options []option.RequestOption
}

// NewEmailInboxMessageActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailInboxMessageActionService(opts ...option.RequestOption) (r EmailInboxMessageActionService) {
	r = EmailInboxMessageActionService{}
	r.Options = opts
	return
}

// Sends from the inbox address through the standard email send pipeline to
// caller-supplied To, Cc, and Bcc recipients. `to` must contain at least one
// recipient. Optional `text` and `html` are prepended to a forwarded-message block
// containing the original metadata and available body content. The subject is
// prefixed with `Fwd:` unless it already has that prefix.
//
// Threading headers are derived from the original message: `In-Reply-To` is set to
// its RFC Message-ID, and `References` contains the original References values
// plus that Message-ID, de-duplicated and limited to the most recent 20 values.
func (r *EmailInboxMessageActionService) Forward(ctx context.Context, messageID string, params EmailInboxMessageActionForwardParams, opts ...option.RequestOption) (res *EmailMessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/messages/%s/actions/forward", params.InboxID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Sends from the inbox address through the standard email send pipeline. The
// recipient is the original `Reply-To`, falling back to `From`; original Cc
// recipients are not included. The subject is prefixed with `Re:` unless it
// already has that prefix.
//
// Threading headers are derived from the original message: `In-Reply-To` is set to
// its RFC Message-ID, and `References` contains the original References values
// plus that Message-ID, de-duplicated and limited to the most recent 20 values.
func (r *EmailInboxMessageActionService) Reply(ctx context.Context, messageID string, params EmailInboxMessageActionReplyParams, opts ...option.RequestOption) (res *EmailMessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/messages/%s/actions/reply", params.InboxID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Sends from the inbox address through the standard email send pipeline. The To
// list starts with the original `Reply-To` (or `From`) and includes original To
// recipients; the Cc list includes original Cc recipients. The inbox address is
// excluded, and recipients are de-duplicated case-insensitively across To and Cc.
// Bcc is always empty. The subject is prefixed with `Re:` unless it already has
// that prefix.
//
// Threading headers are derived from the original message: `In-Reply-To` is set to
// its RFC Message-ID, and `References` contains the original References values
// plus that Message-ID, de-duplicated and limited to the most recent 20 values.
func (r *EmailInboxMessageActionService) ReplyAll(ctx context.Context, messageID string, params EmailInboxMessageActionReplyAllParams, opts ...option.RequestOption) (res *EmailMessageResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.InboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	if messageID == "" {
		err = errors.New("missing required message_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/messages/%s/actions/reply_all", params.InboxID, messageID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

func InboxActionEmailAddressInputParamOfInboxActionEmailAddressInputUnionMember1(email string) InboxActionEmailAddressInputUnionParam {
	var variant InboxActionEmailAddressInputUnionMember1Param
	variant.Email = email
	return InboxActionEmailAddressInputUnionParam{OfInboxActionEmailAddressInputUnionMember1: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InboxActionEmailAddressInputUnionParam struct {
	OfString                                   param.Opt[string]                              `json:",omitzero,inline"`
	OfInboxActionEmailAddressInputUnionMember1 *InboxActionEmailAddressInputUnionMember1Param `json:",omitzero,inline"`
	paramUnion
}

func (u InboxActionEmailAddressInputUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInboxActionEmailAddressInputUnionMember1)
}
func (u *InboxActionEmailAddressInputUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InboxActionEmailAddressInputUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInboxActionEmailAddressInputUnionMember1) {
		return u.OfInboxActionEmailAddressInputUnionMember1
	}
	return nil
}

// The property Email is required.
type InboxActionEmailAddressInputUnionMember1Param struct {
	Email string            `json:"email" api:"required"`
	Name  param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r InboxActionEmailAddressInputUnionMember1Param) MarshalJSON() (data []byte, err error) {
	type shadow InboxActionEmailAddressInputUnionMember1Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboxActionEmailAddressInputUnionMember1Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func InboxActionRecipientInputParamOfInboxActionRecipientInputUnionMember1(email string) InboxActionRecipientInputUnionParam {
	var variant InboxActionRecipientInputUnionMember1Param
	variant.Email = email
	return InboxActionRecipientInputUnionParam{OfInboxActionRecipientInputUnionMember1: &variant}
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type InboxActionRecipientInputUnionParam struct {
	OfString                                param.Opt[string]                           `json:",omitzero,inline"`
	OfInboxActionRecipientInputUnionMember1 *InboxActionRecipientInputUnionMember1Param `json:",omitzero,inline"`
	OfInboxActionEmailAddressInputArray     []InboxActionEmailAddressInputUnionParam    `json:",omitzero,inline"`
	paramUnion
}

func (u InboxActionRecipientInputUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfInboxActionRecipientInputUnionMember1, u.OfInboxActionEmailAddressInputArray)
}
func (u *InboxActionRecipientInputUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *InboxActionRecipientInputUnionParam) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfInboxActionRecipientInputUnionMember1) {
		return u.OfInboxActionRecipientInputUnionMember1
	} else if !param.IsOmitted(u.OfInboxActionEmailAddressInputArray) {
		return &u.OfInboxActionEmailAddressInputArray
	}
	return nil
}

// The property Email is required.
type InboxActionRecipientInputUnionMember1Param struct {
	Email string            `json:"email" api:"required"`
	Name  param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r InboxActionRecipientInputUnionMember1Param) MarshalJSON() (data []byte, err error) {
	type shadow InboxActionRecipientInputUnionMember1Param
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InboxActionRecipientInputUnionMember1Param) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// At least one of `text` or `html` must contain a non-whitespace body. Recipients
// are derived from the source message; caller-supplied `to`, `cc`, or `bcc` values
// are ignored.
type ReplyEmailInboxMessageRequestParam struct {
	// HTML reply body.
	HTML param.Opt[string] `json:"html,omitzero"`
	// Plain-text reply body.
	Text param.Opt[string] `json:"text,omitzero"`
	paramObj
}

func (r ReplyEmailInboxMessageRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ReplyEmailInboxMessageRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReplyEmailInboxMessageRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxMessageActionForwardParams struct {
	InboxID string `path:"inbox_id" api:"required" format:"uuid" json:"-"`
	// One recipient or a non-empty recipient array. Each recipient may be an email
	// string or an object with `email` and optional `name`.
	To EmailInboxMessageActionForwardParamsToUnion `json:"to,omitzero" api:"required"`
	// Optional HTML note prepended to the generated forwarded-message block. Blank
	// values are treated as omitted.
	HTML param.Opt[string] `json:"html,omitzero"`
	// Optional plain-text note prepended to the generated forwarded-message block.
	// Blank values are treated as omitted.
	Text param.Opt[string] `json:"text,omitzero"`
	// One recipient or a recipient array. Each recipient may be an email string or an
	// object with `email` and optional `name`.
	Bcc InboxActionRecipientInputUnionParam `json:"bcc,omitzero"`
	// One recipient or a recipient array. Each recipient may be an email string or an
	// object with `email` and optional `name`.
	Cc InboxActionRecipientInputUnionParam `json:"cc,omitzero"`
	paramObj
}

func (r EmailInboxMessageActionForwardParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailInboxMessageActionForwardParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailInboxMessageActionForwardParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type EmailInboxMessageActionForwardParamsToUnion struct {
	OfString                                        param.Opt[string]                                   `json:",omitzero,inline"`
	OfEmailInboxMessageActionForwardsToUnionMember1 *EmailInboxMessageActionForwardParamsToUnionMember1 `json:",omitzero,inline"`
	OfInboxActionEmailAddressInputArray             []InboxActionEmailAddressInputUnionParam            `json:",omitzero,inline"`
	paramUnion
}

func (u EmailInboxMessageActionForwardParamsToUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfEmailInboxMessageActionForwardsToUnionMember1, u.OfInboxActionEmailAddressInputArray)
}
func (u *EmailInboxMessageActionForwardParamsToUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *EmailInboxMessageActionForwardParamsToUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfEmailInboxMessageActionForwardsToUnionMember1) {
		return u.OfEmailInboxMessageActionForwardsToUnionMember1
	} else if !param.IsOmitted(u.OfInboxActionEmailAddressInputArray) {
		return &u.OfInboxActionEmailAddressInputArray
	}
	return nil
}

// The property Email is required.
type EmailInboxMessageActionForwardParamsToUnionMember1 struct {
	Email string            `json:"email" api:"required"`
	Name  param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r EmailInboxMessageActionForwardParamsToUnionMember1) MarshalJSON() (data []byte, err error) {
	type shadow EmailInboxMessageActionForwardParamsToUnionMember1
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailInboxMessageActionForwardParamsToUnionMember1) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxMessageActionReplyParams struct {
	InboxID string `path:"inbox_id" api:"required" format:"uuid" json:"-"`
	// At least one of `text` or `html` must contain a non-whitespace body. Recipients
	// are derived from the source message; caller-supplied `to`, `cc`, or `bcc` values
	// are ignored.
	ReplyEmailInboxMessageRequest ReplyEmailInboxMessageRequestParam
	paramObj
}

func (r EmailInboxMessageActionReplyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ReplyEmailInboxMessageRequest)
}
func (r *EmailInboxMessageActionReplyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxMessageActionReplyAllParams struct {
	InboxID string `path:"inbox_id" api:"required" format:"uuid" json:"-"`
	// At least one of `text` or `html` must contain a non-whitespace body. Recipients
	// are derived from the source message; caller-supplied `to`, `cc`, or `bcc` values
	// are ignored.
	ReplyEmailInboxMessageRequest ReplyEmailInboxMessageRequestParam
	paramObj
}

func (r EmailInboxMessageActionReplyAllParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ReplyEmailInboxMessageRequest)
}
func (r *EmailInboxMessageActionReplyAllParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
