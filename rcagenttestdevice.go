// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Manage RCS agent registration, testing, verification, and launch.
//
// RcAgentTestDeviceService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRcAgentTestDeviceService] method instead.
type RcAgentTestDeviceService struct {
	Options []option.RequestOption
}

// NewRcAgentTestDeviceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRcAgentTestDeviceService(opts ...option.RequestOption) (r RcAgentTestDeviceService) {
	r = RcAgentTestDeviceService{}
	r.Options = opts
	return
}

// Adds an RCS-capable test number after provider agent creation. Repeating the
// request for a number already attached to the agent returns the existing test
// device.
func (r *RcAgentTestDeviceService) New(ctx context.Context, id string, body RcAgentTestDeviceNewParams, opts ...option.RequestOption) (res *TestDeviceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rcs/agents/%s/test_devices", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists test devices attached to an RCS agent.
func (r *RcAgentTestDeviceService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *[]TestDeviceResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("rcs/agents/%s/test_devices", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Removes a test device from an RCS agent and its provider registration.
func (r *RcAgentTestDeviceService) Delete(ctx context.Context, testDeviceID string, body RcAgentTestDeviceDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ID == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if testDeviceID == "" {
		err = errors.New("missing required test_device_id parameter")
		return err
	}
	path := fmt.Sprintf("rcs/agents/%s/test_devices/%s", body.ID, testDeviceID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type TestDeviceResponse struct {
	// Any of "PENDING", "ACCEPTED", "DECLINED".
	InviteStatus TestDeviceResponseInviteStatus `json:"invite_status" api:"required"`
	PhoneNumber  string                         `json:"phone_number" api:"required"`
	TestDeviceID string                         `json:"test_device_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InviteStatus respjson.Field
		PhoneNumber  respjson.Field
		TestDeviceID respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TestDeviceResponse) RawJSON() string { return r.JSON.raw }
func (r *TestDeviceResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TestDeviceResponseInviteStatus string

const (
	TestDeviceResponseInviteStatusPending  TestDeviceResponseInviteStatus = "PENDING"
	TestDeviceResponseInviteStatusAccepted TestDeviceResponseInviteStatus = "ACCEPTED"
	TestDeviceResponseInviteStatusDeclined TestDeviceResponseInviteStatus = "DECLINED"
)

type RcAgentTestDeviceNewParams struct {
	PhoneNumber string `json:"phone_number" api:"required"`
	paramObj
}

func (r RcAgentTestDeviceNewParams) MarshalJSON() (data []byte, err error) {
	type shadow RcAgentTestDeviceNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RcAgentTestDeviceNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RcAgentTestDeviceDeleteParams struct {
	ID string `path:"id" api:"required" format:"uuid" json:"-"`
	paramObj
}
