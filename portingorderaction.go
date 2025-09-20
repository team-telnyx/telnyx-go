// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// PortingOrderActionService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPortingOrderActionService] method instead.
type PortingOrderActionService struct {
	Options []option.RequestOption
}

// NewPortingOrderActionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPortingOrderActionService(opts ...option.RequestOption) (r PortingOrderActionService) {
	r = PortingOrderActionService{}
	r.Options = opts
	return
}

// Activate each number in a porting order asynchronously. This operation is
// limited to US FastPort orders only.
func (r *PortingOrderActionService) Activate(ctx context.Context, id string, opts ...option.RequestOption) (res *PortingOrderActionActivateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/actions/activate", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Cancel a porting order
func (r *PortingOrderActionService) Cancel(ctx context.Context, id string, opts ...option.RequestOption) (res *PortingOrderActionCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/actions/cancel", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Confirm and submit your porting order.
func (r *PortingOrderActionService) Confirm(ctx context.Context, id string, opts ...option.RequestOption) (res *PortingOrderActionConfirmResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/actions/confirm", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Creates a sharing token for a porting order. The token can be used to share the
// porting order with non-Telnyx users.
func (r *PortingOrderActionService) Share(ctx context.Context, id string, body PortingOrderActionShareParams, opts ...option.RequestOption) (res *PortingOrderActionShareResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("porting_orders/%s/actions/share", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type PortingOrderActionActivateResponse struct {
	Data PortingOrdersActivationJob `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderActionActivateResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderActionActivateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderActionCancelResponse struct {
	Data PortingOrder                         `json:"data"`
	Meta PortingOrderActionCancelResponseMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderActionCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderActionCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderActionCancelResponseMeta struct {
	// Link to list all phone numbers
	PhoneNumbersURL string `json:"phone_numbers_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumbersURL respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderActionCancelResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderActionCancelResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderActionConfirmResponse struct {
	Data PortingOrder                          `json:"data"`
	Meta PortingOrderActionConfirmResponseMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderActionConfirmResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderActionConfirmResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderActionConfirmResponseMeta struct {
	// Link to list all phone numbers
	PhoneNumbersURL string `json:"phone_numbers_url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumbersURL respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderActionConfirmResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderActionConfirmResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderActionShareResponse struct {
	Data PortingOrderActionShareResponseData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderActionShareResponse) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderActionShareResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderActionShareResponseData struct {
	// Uniquely identifies this sharing token
	ID string `json:"id" format:"uuid"`
	// A signed JWT token that can be used to access the shared resource
	Token string `json:"token"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// ISO 8601 formatted date indicating when the sharing token expires.
	ExpiresAt time.Time `json:"expires_at" format:"date-time"`
	// The number of seconds until the sharing token expires
	ExpiresInSeconds int64 `json:"expires_in_seconds"`
	// The permissions granted to the sharing token
	//
	// Any of "porting_order.document.read", "porting_order.document.update".
	Permissions []string `json:"permissions"`
	// Identifies the porting order resource being shared
	PortingOrderID string `json:"porting_order_id" format:"uuid"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Token            respjson.Field
		CreatedAt        respjson.Field
		ExpiresAt        respjson.Field
		ExpiresInSeconds respjson.Field
		Permissions      respjson.Field
		PortingOrderID   respjson.Field
		RecordType       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PortingOrderActionShareResponseData) RawJSON() string { return r.JSON.raw }
func (r *PortingOrderActionShareResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PortingOrderActionShareParams struct {
	// The number of seconds the token will be valid for
	ExpiresInSeconds param.Opt[int64] `json:"expires_in_seconds,omitzero"`
	// The permissions the token will have
	//
	// Any of "porting_order.document.read", "porting_order.document.update".
	Permissions PortingOrderActionShareParamsPermissions `json:"permissions,omitzero"`
	paramObj
}

func (r PortingOrderActionShareParams) MarshalJSON() (data []byte, err error) {
	type shadow PortingOrderActionShareParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortingOrderActionShareParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The permissions the token will have
type PortingOrderActionShareParamsPermissions string

const (
	PortingOrderActionShareParamsPermissionsPortingOrderDocumentRead   PortingOrderActionShareParamsPermissions = "porting_order.document.read"
	PortingOrderActionShareParamsPermissionsPortingOrderDocumentUpdate PortingOrderActionShareParamsPermissions = "porting_order.document.update"
)
