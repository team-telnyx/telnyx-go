// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Group dedicated IPs into pools for sending.
//
// EmailIPPoolIPService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailIPPoolIPService] method instead.
type EmailIPPoolIPService struct {
	Options []option.RequestOption
}

// NewEmailIPPoolIPService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailIPPoolIPService(opts ...option.RequestOption) (r EmailIPPoolIPService) {
	r = EmailIPPoolIPService{}
	r.Options = opts
	return
}

// Add a dedicated IP to the pool with an optional routing weight. An IP can belong
// to one pool; to move an IP that is already in a pool, use
// `PATCH /email/dedicated_ips/{id}` with `pool_id`.
func (r *EmailIPPoolIPService) New(ctx context.Context, id string, body EmailIPPoolIPNewParams, opts ...option.RequestOption) (res *EmailIPPoolIPNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email/ip_pools/%s/ips", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Remove the dedicated IP from the pool. The IP is kept but carries no mail until
// added to a pool again. Removing the last IP requires `force=true` and
// `drain=true`.
func (r *EmailIPPoolIPService) Delete(ctx context.Context, ipID string, params EmailIPPoolIPDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if params.ID == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	if ipID == "" {
		err = errors.New("missing required ip_id parameter")
		return err
	}
	path := fmt.Sprintf("email/ip_pools/%s/ips/%s", params.ID, ipID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, params, nil, opts...)
	return err
}

type EmailIPPoolIPNewResponse struct {
	Data EmailIPPoolIPNewResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailIPPoolIPNewResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailIPPoolIPNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailIPPoolIPNewResponseData struct {
	CreatedAt     time.Time `json:"created_at" format:"date-time"`
	DedicatedIPID string    `json:"dedicated_ip_id" format:"uuid"`
	PoolID        string    `json:"pool_id" format:"uuid"`
	Weight        int64     `json:"weight"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt     respjson.Field
		DedicatedIPID respjson.Field
		PoolID        respjson.Field
		Weight        respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailIPPoolIPNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *EmailIPPoolIPNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailIPPoolIPNewParams struct {
	// The dedicated IP to add. It must belong to the account, must not be released,
	// and must not already be in a pool.
	DedicatedIPID string `json:"dedicated_ip_id" api:"required" format:"uuid"`
	// Relative routing weight for this IP in the pool. Positive integer; defaults
	// to 100.
	Weight param.Opt[int64] `json:"weight,omitzero"`
	paramObj
}

func (r EmailIPPoolIPNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailIPPoolIPNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailIPPoolIPNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailIPPoolIPDeleteParams struct {
	ID string `path:"id" api:"required" format:"uuid" json:"-"`
	// Set together with `force=true` to remove the last IP in the pool. Acknowledges
	// that mail routed to the emptied pool is held until the pool has a sendable IP
	// again.
	Drain param.Opt[bool] `query:"drain,omitzero" json:"-"`
	// Set together with `drain=true` to remove the last IP in the pool. Either flag
	// alone has no effect.
	Force param.Opt[bool] `query:"force,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailIPPoolIPDeleteParams]'s query parameters as
// `url.Values`.
func (r EmailIPPoolIPDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
