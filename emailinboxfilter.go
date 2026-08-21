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
// EmailInboxFilterService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailInboxFilterService] method instead.
type EmailInboxFilterService struct {
	Options []option.RequestOption
}

// NewEmailInboxFilterService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailInboxFilterService(opts ...option.RequestOption) (r EmailInboxFilterService) {
	r = EmailInboxFilterService{}
	r.Options = opts
	return
}

// Returns the inbox's sender allowlist and blocklist. Entries are normalized to
// lowercase. A blocklist match takes precedence over an allowlist match; when both
// lists are empty, all senders are accepted.
func (r *EmailInboxFilterService) List(ctx context.Context, inboxID string, opts ...option.RequestOption) (res *EmailInboxFilterListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/filters", inboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Adds entries to either the allowlist or blocklist. The operation is an
// idempotent set union: entries already present remain unchanged.
func (r *EmailInboxFilterService) Add(ctx context.Context, inboxID string, body EmailInboxFilterAddParams, opts ...option.RequestOption) (res *EmailInboxFilterAddResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/filters", inboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Removes entries from either the allowlist or blocklist. The operation is
// idempotent: removing an entry that is not present still returns the current
// filter lists.
func (r *EmailInboxFilterService) DeleteAll(ctx context.Context, inboxID string, body EmailInboxFilterDeleteAllParams, opts ...option.RequestOption) (res *EmailInboxFilterDeleteAllResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/filters", inboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return res, err
}

// Replaces both sender filter lists atomically. Omitting either list clears that
// list. Use `POST` or `DELETE` for incremental changes.
func (r *EmailInboxFilterService) Replace(ctx context.Context, inboxID string, body EmailInboxFilterReplaceParams, opts ...option.RequestOption) (res *EmailInboxFilterReplaceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if inboxID == "" {
		err = errors.New("missing required inbox_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_inboxes/%s/filters", inboxID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

type InboxFilters struct {
	Allowlist []string `json:"allowlist" api:"required"`
	Blocklist []string `json:"blocklist" api:"required"`
	// Any of "email_inbox_filters".
	RecordType InboxFiltersRecordType `json:"record_type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allowlist   respjson.Field
		Blocklist   respjson.Field
		RecordType  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InboxFilters) RawJSON() string { return r.JSON.raw }
func (r *InboxFilters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InboxFiltersRecordType string

const (
	InboxFiltersRecordTypeEmailInboxFilters InboxFiltersRecordType = "email_inbox_filters"
)

// The properties Entries, Type are required.
type MutateInboxFiltersRequestParam struct {
	Entries []string `json:"entries,omitzero" api:"required"`
	// The list to change.
	//
	// Any of "allowlist", "blocklist".
	Type MutateInboxFiltersRequestType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r MutateInboxFiltersRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow MutateInboxFiltersRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MutateInboxFiltersRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The list to change.
type MutateInboxFiltersRequestType string

const (
	MutateInboxFiltersRequestTypeAllowlist MutateInboxFiltersRequestType = "allowlist"
	MutateInboxFiltersRequestTypeBlocklist MutateInboxFiltersRequestType = "blocklist"
)

type EmailInboxFilterListResponse struct {
	Data InboxFilters `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailInboxFilterListResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailInboxFilterListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxFilterAddResponse struct {
	Data InboxFilters `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailInboxFilterAddResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailInboxFilterAddResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxFilterDeleteAllResponse struct {
	Data InboxFilters `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailInboxFilterDeleteAllResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailInboxFilterDeleteAllResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxFilterReplaceResponse struct {
	Data InboxFilters `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailInboxFilterReplaceResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailInboxFilterReplaceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxFilterAddParams struct {
	MutateInboxFiltersRequest MutateInboxFiltersRequestParam
	paramObj
}

func (r EmailInboxFilterAddParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MutateInboxFiltersRequest)
}
func (r *EmailInboxFilterAddParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxFilterDeleteAllParams struct {
	MutateInboxFiltersRequest MutateInboxFiltersRequestParam
	paramObj
}

func (r EmailInboxFilterDeleteAllParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MutateInboxFiltersRequest)
}
func (r *EmailInboxFilterDeleteAllParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailInboxFilterReplaceParams struct {
	Allowlist []string `json:"allowlist,omitzero"`
	Blocklist []string `json:"blocklist,omitzero"`
	paramObj
}

func (r EmailInboxFilterReplaceParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailInboxFilterReplaceParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailInboxFilterReplaceParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
