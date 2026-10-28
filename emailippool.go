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
// EmailIPPoolService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEmailIPPoolService] method instead.
type EmailIPPoolService struct {
	Options []option.RequestOption
	// Group dedicated IPs into pools for sending.
	IPs EmailIPPoolIPService
}

// NewEmailIPPoolService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEmailIPPoolService(opts ...option.RequestOption) (r EmailIPPoolService) {
	r = EmailIPPoolService{}
	r.Options = opts
	r.IPs = NewEmailIPPoolIPService(opts...)
	return
}

// Create an account-scoped pool of dedicated IPs. Set `is_default` or `domain_id`
// for the pool to carry mail; a pool with neither carries none.
func (r *EmailIPPoolService) New(ctx context.Context, body EmailIPPoolNewParams, opts ...option.RequestOption) (res *IPPoolShowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "email/ip_pools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a single IP pool by its identifier, including its dedicated-IP members
// with billing state.
func (r *EmailIPPoolService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *IPPoolShowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email/ip_pools/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Rename the pool, set or clear the default flag, or change its domain scope.
// Routing follows the new settings for subsequent mail.
func (r *EmailIPPoolService) Update(ctx context.Context, id string, body EmailIPPoolUpdateParams, opts ...option.RequestOption) (res *IPPoolShowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("email/ip_pools/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List the account's dedicated-IP pools. Member IPs are not included (`ips` is
// empty); retrieve a pool to see its members.
func (r *EmailIPPoolService) List(ctx context.Context, opts ...option.RequestOption) (res *EmailIPPoolListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "email/ip_pools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete the pool. Member IPs are kept but leave the pool, so they carry no mail
// until added to another pool. Deleting a pool changes routing for subsequently
// resolved messages. It does not reroute messages already assigned to the deleted
// pool. Deleting a pool that still has IPs requires `force=true` and `drain=true`.
func (r *EmailIPPoolService) Delete(ctx context.Context, id string, body EmailIPPoolDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("email/ip_pools/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, nil, opts...)
	return err
}

// An account-scoped pool of dedicated IPs. Mail uses a pool only when the pool is
// scoped to the sending domain (`domain_id`) or is the account default
// (`is_default`).
type IPPool struct {
	ID        string    `json:"id" format:"uuid"`
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Email domain this pool is scoped to. The pool carries that domain's mail. Null
	// when the pool is not domain-scoped.
	DomainID string `json:"domain_id" api:"nullable" format:"uuid"`
	// Member IPs with their routing weights. Populated by Retrieve an IP pool; list,
	// create, and update responses return an empty array.
	IPs []IPPoolIP `json:"ips"`
	// Whether this is the account's default pool. The default pool carries mail for
	// sending domains that have no domain-scoped pool. At most one per account.
	IsDefault bool `json:"is_default"`
	// Pool name, unique within the account.
	Name string `json:"name"`
	// Any of "ip_pool".
	RecordType IPPoolRecordType `json:"record_type"`
	UpdatedAt  time.Time        `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		DomainID    respjson.Field
		IPs         respjson.Field
		IsDefault   respjson.Field
		Name        respjson.Field
		RecordType  respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IPPool) RawJSON() string { return r.JSON.raw }
func (r *IPPool) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A dedicated IP as a member of a pool, with its routing weight in this pool.
// Weight is a positive integer (default 100).
type IPPoolIP struct {
	// Relative routing weight of this IP within the pool.
	Weight int64 `json:"weight"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Weight      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	DedicatedIP
}

// Returns the unmodified JSON received from the API
func (r IPPoolIP) RawJSON() string { return r.JSON.raw }
func (r *IPPoolIP) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IPPoolRecordType string

const (
	IPPoolRecordTypeIPPool IPPoolRecordType = "ip_pool"
)

type IPPoolShowResponse struct {
	// An account-scoped pool of dedicated IPs. Mail uses a pool only when the pool is
	// scoped to the sending domain (`domain_id`) or is the account default
	// (`is_default`).
	Data IPPool `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IPPoolShowResponse) RawJSON() string { return r.JSON.raw }
func (r *IPPoolShowResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailIPPoolListResponse struct {
	Data []IPPool `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EmailIPPoolListResponse) RawJSON() string { return r.JSON.raw }
func (r *EmailIPPoolListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailIPPoolNewParams struct {
	// Pool name: 1-64 characters of lowercase letters, digits, `.`, `_`, and `-`,
	// starting and ending with a letter or digit. Must be unique within the account
	// and must not start with `dip_`.
	Name string `json:"name" api:"required"`
	// Scope the pool to one of the account's email domains; the pool then carries that
	// domain's mail. Scope at most one pool to each domain.
	DomainID param.Opt[string] `json:"domain_id,omitzero" format:"uuid"`
	// Make this the account's default pool, which carries mail for sending domains
	// without a domain-scoped pool. An account can have one default pool. Defaults to
	// `false`.
	IsDefault param.Opt[bool] `json:"is_default,omitzero"`
	paramObj
}

func (r EmailIPPoolNewParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailIPPoolNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailIPPoolNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailIPPoolUpdateParams struct {
	// Scope the pool to one of the account's email domains, or `null` to remove the
	// scope.
	DomainID param.Opt[string] `json:"domain_id,omitzero" format:"uuid"`
	// Set or clear the account default flag. An account can have one default pool.
	IsDefault param.Opt[bool] `json:"is_default,omitzero"`
	// Pool name: 1-64 characters of lowercase letters, digits, `.`, `_`, and `-`,
	// starting and ending with a letter or digit. Must be unique within the account
	// and must not start with `dip_`.
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r EmailIPPoolUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow EmailIPPoolUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EmailIPPoolUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EmailIPPoolDeleteParams struct {
	// Set together with `force=true` to delete a pool that still has IPs. Acknowledges
	// deletion while the pool still has members; this flag does not reroute already
	// queued messages.
	Drain param.Opt[bool] `query:"drain,omitzero" json:"-"`
	// Set together with `drain=true` to delete a pool that still has IPs. Either flag
	// alone has no effect.
	Force param.Opt[bool] `query:"force,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [EmailIPPoolDeleteParams]'s query parameters as
// `url.Values`.
func (r EmailIPPoolDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
