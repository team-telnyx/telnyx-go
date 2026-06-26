// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Trademark or impersonation claims filed against your DIR. Customers may contest
// a claim with supporting evidence.
//
// InfringementClaimService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewInfringementClaimService] method instead.
type InfringementClaimService struct {
	Options []option.RequestOption
}

// NewInfringementClaimService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewInfringementClaimService(opts ...option.RequestOption) (r InfringementClaimService) {
	r = InfringementClaimService{}
	r.Options = opts
	return
}

// Retrieve a single claim by id. Returns `404` if the claim does not exist or is
// not against a DIR you own.
func (r *InfringementClaimService) Get(ctx context.Context, claimID string, opts ...option.RequestOption) (res *InfringementClaimWrapped, err error) {
	opts = slices.Concat(r.Options, opts)
	if claimID == "" {
		err = errors.New("missing required claim_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("infringement_claims/%s", claimID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Submit a written response and supporting documents disputing the claim. The
// first call moves the claim from `pending` to `contested`; subsequent calls
// append supplementary evidence without changing status. The `documents[]` you
// attach are aggregated across rounds in the claim's `contest_documents` field.
//
// Only `pending` and `contested` claims accept new evidence. A `resolved` claim
// returns `400`.
//
// Failure modes:
//
//   - `400` - the claim is `resolved` (terminal); cannot be contested further.
//   - `404` - the claim does not exist or is not against a DIR you own.
//   - `422` - `contest_notes` is too short (< 10 chars), too long (> 2000 chars),
//     `documents` is > 20 entries, or a `document_id` is duplicated within the same
//     submission.
func (r *InfringementClaimService) Contest(ctx context.Context, claimID string, body InfringementClaimContestParams, opts ...option.RequestOption) (res *InfringementClaimWrapped, err error) {
	opts = slices.Concat(r.Options, opts)
	if claimID == "" {
		err = errors.New("missing required claim_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("infringement_claims/%s/contest", claimID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type InfringementClaim struct {
	ID string `json:"id" format:"uuid"`
	// When the claim was filed (set by the claimant's representative at file time).
	ClaimDate        time.Time `json:"claim_date" format:"date-time"`
	ClaimDescription string    `json:"claim_description"`
	// Category of infringement being claimed.
	//
	// Any of "trademark", "copyright".
	ClaimType       InfringementClaimClaimType `json:"claim_type"`
	ClaimantContact string                     `json:"claimant_contact"`
	ClaimantName    string                     `json:"claimant_name"`
	// Aggregated across all customer contest submissions on this claim.
	ContestDocuments []Document `json:"contest_documents"`
	// Per-round submission audit trail. Each entry records one
	// `POST /infringement_claims/{id}/contest` call (notes, timestamp, document
	// count). Aggregated documents live on `contest_documents`.
	ContestHistory []InfringementClaimContestHistory `json:"contest_history"`
	CreatedAt      time.Time                         `json:"created_at" format:"date-time"`
	// Snapshot of the DIR the claim is filed against, embedded for convenience.
	Dir          InfringementClaimDir `json:"dir"`
	DirID        string               `json:"dir_id" format:"uuid"`
	EnterpriseID string               `json:"enterprise_id" format:"uuid"`
	// Set only when `status` is `resolved`.
	//
	// Any of "upheld", "rejected", "modified".
	Resolution      InfringementClaimResolution `json:"resolution" api:"nullable"`
	ResolutionDate  time.Time                   `json:"resolution_date" api:"nullable" format:"date-time"`
	ResolutionNotes string                      `json:"resolution_notes" api:"nullable"`
	// Lifecycle status. `pending` - newly filed; the DIR is auto-suspended.
	// `contested` - you have submitted contest evidence; awaiting Telnyx review.
	// `resolved` - final.
	//
	// Any of "pending", "contested", "resolved".
	Status    InfringementClaimStatus `json:"status"`
	UpdatedAt time.Time               `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		ClaimDate        respjson.Field
		ClaimDescription respjson.Field
		ClaimType        respjson.Field
		ClaimantContact  respjson.Field
		ClaimantName     respjson.Field
		ContestDocuments respjson.Field
		ContestHistory   respjson.Field
		CreatedAt        respjson.Field
		Dir              respjson.Field
		DirID            respjson.Field
		EnterpriseID     respjson.Field
		Resolution       respjson.Field
		ResolutionDate   respjson.Field
		ResolutionNotes  respjson.Field
		Status           respjson.Field
		UpdatedAt        respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InfringementClaim) RawJSON() string { return r.JSON.raw }
func (r *InfringementClaim) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Category of infringement being claimed.
type InfringementClaimClaimType string

const (
	InfringementClaimClaimTypeTrademark InfringementClaimClaimType = "trademark"
	InfringementClaimClaimTypeCopyright InfringementClaimClaimType = "copyright"
)

// One round of customer contest evidence on an infringement claim. The aggregated
// documents across rounds live on the parent claim's `contest_documents`; this
// submission record carries only the per-round notes and document count.
type InfringementClaimContestHistory struct {
	DocumentCount int64     `json:"document_count"`
	Notes         string    `json:"notes"`
	SubmittedAt   time.Time `json:"submitted_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocumentCount respjson.Field
		Notes         respjson.Field
		SubmittedAt   respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InfringementClaimContestHistory) RawJSON() string { return r.JSON.raw }
func (r *InfringementClaimContestHistory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Snapshot of the DIR the claim is filed against, embedded for convenience.
type InfringementClaimDir struct {
	ID           string `json:"id" format:"uuid"`
	DisplayName  string `json:"display_name"`
	EnterpriseID string `json:"enterprise_id" format:"uuid"`
	// DIR lifecycle status.
	//
	//   - `draft` - newly created; editable; not yet submitted.
	//   - `submitted` / `in_review` - Telnyx is reviewing.
	//   - `verified` - approved; phone numbers may be attached.
	//   - `rejected` - Telnyx rejected this submission; `rejection_reasons` is
	//     populated; customer can edit and resubmit.
	//   - `unsuccessful` - system-side error during processing; customer can edit and
	//     resubmit.
	//   - `suspended` - temporarily disabled (e.g. by an active infringement claim).
	//   - `expired` - verification expired; customer must resubmit.
	//   - `infringement_claimed` - a trademark/impersonation claim is open against this
	//     DIR.
	//   - `permanently_rejected` - terminal; cannot be resubmitted.
	//
	// Any of "draft", "submitted", "in_review", "verified", "rejected",
	// "unsuccessful", "suspended", "expired", "infringement_claimed",
	// "permanently_rejected".
	Status DirStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		DisplayName  respjson.Field
		EnterpriseID respjson.Field
		Status       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InfringementClaimDir) RawJSON() string { return r.JSON.raw }
func (r *InfringementClaimDir) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Set only when `status` is `resolved`.
type InfringementClaimResolution string

const (
	InfringementClaimResolutionUpheld   InfringementClaimResolution = "upheld"
	InfringementClaimResolutionRejected InfringementClaimResolution = "rejected"
	InfringementClaimResolutionModified InfringementClaimResolution = "modified"
)

// Lifecycle status. `pending` - newly filed; the DIR is auto-suspended.
// `contested` - you have submitted contest evidence; awaiting Telnyx review.
// `resolved` - final.
type InfringementClaimStatus string

const (
	InfringementClaimStatusPending   InfringementClaimStatus = "pending"
	InfringementClaimStatusContested InfringementClaimStatus = "contested"
	InfringementClaimStatusResolved  InfringementClaimStatus = "resolved"
)

type InfringementClaimWrapped struct {
	Data InfringementClaim `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InfringementClaimWrapped) RawJSON() string { return r.JSON.raw }
func (r *InfringementClaimWrapped) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InfringementClaimContestParams struct {
	// Customer's response to the claim. 10–2000 characters.
	ContestNotes string `json:"contest_notes" api:"required"`
	// Up to 20 supporting documents per submission. `document_id` must be unique
	// within this submission. Documents are aggregated into the claim's
	// `contest_documents` across all submissions.
	Documents []DocumentParam `json:"documents,omitzero"`
	paramObj
}

func (r InfringementClaimContestParams) MarshalJSON() (data []byte, err error) {
	type shadow InfringementClaimContestParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InfringementClaimContestParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
