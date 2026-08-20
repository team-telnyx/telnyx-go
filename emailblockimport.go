// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apiform"
	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Async CSV import of competitor suppression lists.
//
// EmailBlockImportService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailBlockImportService] method instead.
type EmailBlockImportService struct {
	Options []option.RequestOption
}

// NewEmailBlockImportService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewEmailBlockImportService(opts ...option.RequestOption) (r EmailBlockImportService) {
	r = EmailBlockImportService{}
	r.Options = opts
	return
}

// Accepts `multipart/form-data` with a `file` field (the CSV) and an optional
// `block_ttl_days` (integer >0, default 30). Validates:
//
//   - content ≤ 25 MiB, else `413`
//   - row count ≤ 250 000, else `413`
//   - header-only / all-blank / undetectable provider → `400` Returns `202` with the
//     import record (status `pending`); an Oban worker (`EmailBlockImportWorker`,
//     max_attempts 3) transitions `pending → processing → completed | failed`.
//     `block_ttl_days` applies only to imported `manual_block` rows; other reasons
//     get `expires_at: nil`. Provider is auto-detected from the CSV header
//     (`sendgrid` / `mailgun` / `ses` / `generic`).
func (r *EmailBlockImportService) New(ctx context.Context, body EmailBlockImportNewParams, opts ...option.RequestOption) (res *EmailBlockImportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "email_blocks/import"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Account-scoped fetch (cross-account → 404; malformed UUID → 404). Nullable
// fields are omitted until terminal: `provider`/`completed_at` when nil;
// `processed_rows`/`created_count`/`existing_count`/ `skipped_count`/`error_count`
// only when `status == completed`; `errors` only when non-empty; `failure_reason`
// only on terminal failure.
func (r *EmailBlockImportService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *EmailBlockImportResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email_blocks/import/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Import job. Schema fields hidden: `account_id`, `csv_content`, `block_ttl_days`.
// Nullable fields use the omit-nullable pattern.
type EmailBlockImport struct {
	ID        string    `json:"id" api:"required" format:"uuid"`
	CreatedAt time.Time `json:"created_at" api:"required" format:"date-time"`
	// View-only.
	//
	// Any of "email_block_import".
	RecordType EmailBlockImportRecordType `json:"record_type" api:"required"`
	// Any of "pending", "processing", "completed", "failed".
	Status EmailBlockImportStatus `json:"status" api:"required"`
	// Data-row count at upload.
	Total     int64     `json:"total" api:"required"`
	UpdatedAt time.Time `json:"updated_at" api:"required" format:"date-time"`
	// Omitted until terminal success.
	CompletedAt time.Time `json:"completed_at" format:"date-time"`
	// Only when `status == completed`.
	CreatedCount int64 `json:"created_count"`
	// Only when `status == completed`.
	ErrorCount int64 `json:"error_count"`
	// `{row_number: reason}`; only rendered when non-empty.
	Errors map[string]string `json:"errors"`
	// Only when `status == completed`.
	ExistingCount int64 `json:"existing_count"`
	// Only on terminal failure.
	FailureReason string `json:"failure_reason"`
	// Only when `status == completed`.
	ProcessedRows int64 `json:"processed_rows"`
	// Omitted when nil.
	//
	// Any of "sendgrid", "mailgun", "ses", "generic".
	Provider EmailBlockImportProvider `json:"provider"`
	// Only when `status == completed`.
	SkippedCount int64 `json:"skipped_count"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		RecordType    respjson.Field
		Status        respjson.Field
		Total         respjson.Field
		UpdatedAt     respjson.Field
		CompletedAt   respjson.Field
		CreatedCount  respjson.Field
		ErrorCount    respjson.Field
		Errors        respjson.Field
		ExistingCount respjson.Field
		FailureReason respjson.Field
		ProcessedRows respjson.Field
		Provider      respjson.Field
		SkippedCount  respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailBlockImport) RawJSON() string { return r.JSON.raw }
func (r *EmailBlockImport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// View-only.
type EmailBlockImportRecordType string

const (
	EmailBlockImportRecordTypeEmailBlockImport EmailBlockImportRecordType = "email_block_import"
)

type EmailBlockImportStatus string

const (
	EmailBlockImportStatusPending    EmailBlockImportStatus = "pending"
	EmailBlockImportStatusProcessing EmailBlockImportStatus = "processing"
	EmailBlockImportStatusCompleted  EmailBlockImportStatus = "completed"
	EmailBlockImportStatusFailed     EmailBlockImportStatus = "failed"
)

// Omitted when nil.
type EmailBlockImportProvider string

const (
	EmailBlockImportProviderSendgrid EmailBlockImportProvider = "sendgrid"
	EmailBlockImportProviderMailgun  EmailBlockImportProvider = "mailgun"
	EmailBlockImportProviderSes      EmailBlockImportProvider = "ses"
	EmailBlockImportProviderGeneric  EmailBlockImportProvider = "generic"
)

type EmailBlockImportResponse struct {
	// Import job. Schema fields hidden: `account_id`, `csv_content`, `block_ttl_days`.
	// Nullable fields use the omit-nullable pattern.
	Data EmailBlockImport `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailBlockImportResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailBlockImportResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailBlockImportNewParams struct {
	// The CSV file (Plug.Upload). Missing/non-upload → 400.
	File io.Reader `json:"file,omitzero" api:"required" format:"binary"`
	// TTL for imported `manual_block` rows; other reasons get `expires_at: null`.
	// Invalid/missing → falls back to 30.
	BlockTtlDays param.Opt[int64] `json:"block_ttl_days,omitzero"`
	paramObj
}

func (r EmailBlockImportNewParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}
