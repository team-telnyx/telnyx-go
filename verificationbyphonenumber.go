// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
)

// VerificationByPhoneNumberService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVerificationByPhoneNumberService] method instead.
type VerificationByPhoneNumberService struct {
	Options []option.RequestOption
	Actions VerificationByPhoneNumberActionService
}

// NewVerificationByPhoneNumberService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewVerificationByPhoneNumberService(opts ...option.RequestOption) (r VerificationByPhoneNumberService) {
	r = VerificationByPhoneNumberService{}
	r.Options = opts
	r.Actions = NewVerificationByPhoneNumberActionService(opts...)
	return
}

// List verifications by phone number
func (r *VerificationByPhoneNumberService) List(ctx context.Context, phoneNumber string, opts ...option.RequestOption) (res *VerificationByPhoneNumberListResponse, err error) {
	opts = append(r.Options[:], opts...)
	if phoneNumber == "" {
		err = errors.New("missing required phone_number parameter")
		return
	}
	path := fmt.Sprintf("verifications/by_phone_number/%s", phoneNumber)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type VerificationByPhoneNumberListResponse struct {
	Data []Verification                            `json:"data,required"`
	Meta VerificationByPhoneNumberListResponseMeta `json:"meta,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerificationByPhoneNumberListResponse) RawJSON() string { return r.JSON.raw }
func (r *VerificationByPhoneNumberListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VerificationByPhoneNumberListResponseMeta struct {
	PageNumber   int64 `json:"page_number"`
	PageSize     int64 `json:"page_size"`
	TotalPages   int64 `json:"total_pages"`
	TotalResults int64 `json:"total_results"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PageNumber   respjson.Field
		PageSize     respjson.Field
		TotalPages   respjson.Field
		TotalResults respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VerificationByPhoneNumberListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *VerificationByPhoneNumberListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
