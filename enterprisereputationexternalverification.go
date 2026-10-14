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

// Phone-number reputation monitoring (spam-score lookup and tracking).
//
// EnterpriseReputationExternalVerificationService contains methods and other
// services that help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEnterpriseReputationExternalVerificationService] method instead.
type EnterpriseReputationExternalVerificationService struct {
	Options []option.RequestOption
}

// NewEnterpriseReputationExternalVerificationService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewEnterpriseReputationExternalVerificationService(opts ...option.RequestOption) (r EnterpriseReputationExternalVerificationService) {
	r = EnterpriseReputationExternalVerificationService{}
	r.Options = opts
	return
}

// Request permission to add phone numbers that the enterprise owns but that are
// hosted with another provider. Submit one or more uploaded documents proving
// ownership, such as a carrier invoice or account statement. The request is
// reviewed once per enterprise. While it is `pending` or `rejected` you can submit
// again to replace the documents; once `approved` it cannot be changed.
func (r *EnterpriseReputationExternalVerificationService) New(ctx context.Context, enterpriseID string, body EnterpriseReputationExternalVerificationNewParams, opts ...option.RequestOption) (res *ExternalVerificationWrapped, err error) {
	opts = slices.Concat(r.Options, opts)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation/external_verification", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get the review status of the enterprise's ownership proof for phone numbers
// hosted with another provider. Returns 404 when no proof has been submitted.
func (r *EnterpriseReputationExternalVerificationService) List(ctx context.Context, enterpriseID string, opts ...option.RequestOption) (res *ExternalVerificationWrapped, err error) {
	opts = slices.Concat(r.Options, opts)
	if enterpriseID == "" {
		err = errors.New("missing required enterprise_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("enterprises/%s/reputation/external_verification", enterpriseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Review state of the ownership proof that allows adding phone numbers hosted with
// another provider.
type ExternalVerificationStatus string

const (
	ExternalVerificationStatusPending  ExternalVerificationStatus = "pending"
	ExternalVerificationStatusApproved ExternalVerificationStatus = "approved"
	ExternalVerificationStatusRejected ExternalVerificationStatus = "rejected"
)

type ExternalVerificationWrapped struct {
	Data ExternalVerificationWrappedData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalVerificationWrapped) RawJSON() string { return r.JSON.raw }
func (r *ExternalVerificationWrapped) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ExternalVerificationWrappedData struct {
	// The ownership-proof documents on file.
	DocumentIDs  []string `json:"document_ids"`
	EnterpriseID string   `json:"enterprise_id" format:"uuid"`
	// The phone numbers you listed with the documents.
	PhoneNumbers []string `json:"phone_numbers"`
	// Why the proof was rejected. Set only when `status` is `rejected`.
	RejectionReason string    `json:"rejection_reason" api:"nullable"`
	ReviewedAt      time.Time `json:"reviewed_at" api:"nullable" format:"date-time"`
	// Review state of the ownership proof that allows adding phone numbers hosted with
	// another provider.
	//
	// Any of "pending", "approved", "rejected".
	Status      ExternalVerificationStatus `json:"status"`
	SubmittedAt time.Time                  `json:"submitted_at" api:"nullable" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DocumentIDs     respjson.Field
		EnterpriseID    respjson.Field
		PhoneNumbers    respjson.Field
		RejectionReason respjson.Field
		ReviewedAt      respjson.Field
		Status          respjson.Field
		SubmittedAt     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExternalVerificationWrappedData) RawJSON() string { return r.JSON.raw }
func (r *ExternalVerificationWrappedData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EnterpriseReputationExternalVerificationNewParams struct {
	// Ids of uploaded documents (Documents API) proving the enterprise owns its phone
	// numbers hosted with another provider, for example a carrier invoice or account
	// statement.
	DocumentIDs []string `json:"document_ids,omitzero" api:"required"`
	// Optional. The phone numbers hosted with another provider that the documents
	// cover, in E.164 format. Shown to the reviewer to match the documents to your
	// numbers; you can still add numbers that are not listed here, and each is checked
	// on its own.
	PhoneNumbers []string `json:"phone_numbers,omitzero"`
	paramObj
}

func (r EnterpriseReputationExternalVerificationNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EnterpriseReputationExternalVerificationNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnterpriseReputationExternalVerificationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
