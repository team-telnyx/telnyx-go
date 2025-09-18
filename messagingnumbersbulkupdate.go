// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// MessagingNumbersBulkUpdateService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMessagingNumbersBulkUpdateService] method instead.
type MessagingNumbersBulkUpdateService struct {
	Options []option.RequestOption
}

// NewMessagingNumbersBulkUpdateService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewMessagingNumbersBulkUpdateService(opts ...option.RequestOption) (r MessagingNumbersBulkUpdateService) {
	r = MessagingNumbersBulkUpdateService{}
	r.Options = opts
	return
}

// Update the messaging profile of multiple phone numbers
func (r *MessagingNumbersBulkUpdateService) New(ctx context.Context, body MessagingNumbersBulkUpdateNewParams, opts ...option.RequestOption) (res *MessagingNumbersBulkUpdateNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "messaging_numbers_bulk_updates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve bulk update status
func (r *MessagingNumbersBulkUpdateService) Get(ctx context.Context, orderID string, opts ...option.RequestOption) (res *MessagingNumbersBulkUpdateGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if orderID == "" {
		err = errors.New("missing required order_id parameter")
		return
	}
	path := fmt.Sprintf("messaging_numbers_bulk_updates/%s", orderID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type MessagingNumbersBulkUpdateNewResponse struct {
	Data MessagingNumbersBulkUpdateNewResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingNumbersBulkUpdateNewResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingNumbersBulkUpdateNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingNumbersBulkUpdateNewResponseData struct {
	// Phone numbers that failed to update.
	Failed []string `json:"failed" format:"+E.164"`
	// Order ID to verify bulk update status.
	OrderID string `json:"order_id" format:"uuid"`
	// Phone numbers pending to be updated.
	Pending []string `json:"pending" format:"+E.164"`
	// Identifies the type of the resource.
	//
	// Any of "messaging_numbers_bulk_update".
	RecordType string `json:"record_type"`
	// Phoned numbers updated successfully.
	Success []string `json:"success" format:"+E.164"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Failed      respjson.Field
		OrderID     respjson.Field
		Pending     respjson.Field
		RecordType  respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingNumbersBulkUpdateNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *MessagingNumbersBulkUpdateNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingNumbersBulkUpdateGetResponse struct {
	Data MessagingNumbersBulkUpdateGetResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingNumbersBulkUpdateGetResponse) RawJSON() string { return r.JSON.raw }
func (r *MessagingNumbersBulkUpdateGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingNumbersBulkUpdateGetResponseData struct {
	// Phone numbers that failed to update.
	Failed []string `json:"failed" format:"+E.164"`
	// Order ID to verify bulk update status.
	OrderID string `json:"order_id" format:"uuid"`
	// Phone numbers pending to be updated.
	Pending []string `json:"pending" format:"+E.164"`
	// Identifies the type of the resource.
	//
	// Any of "messaging_numbers_bulk_update".
	RecordType string `json:"record_type"`
	// Phoned numbers updated successfully.
	Success []string `json:"success" format:"+E.164"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Failed      respjson.Field
		OrderID     respjson.Field
		Pending     respjson.Field
		RecordType  respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MessagingNumbersBulkUpdateGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *MessagingNumbersBulkUpdateGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MessagingNumbersBulkUpdateNewParams struct {
	// Configure the messaging profile these phone numbers are assigned to:
	//
	//   - Set this field to `""` to unassign each number from their respective messaging
	//     profile
	//   - Set this field to a quoted UUID of a messaging profile to assign these numbers
	//     to that messaging profile
	MessagingProfileID string `json:"messaging_profile_id,required"`
	// The list of phone numbers to update.
	Numbers []string `json:"numbers,omitzero,required" format:"+E.164"`
	paramObj
}

func (r MessagingNumbersBulkUpdateNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MessagingNumbersBulkUpdateNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MessagingNumbersBulkUpdateNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
