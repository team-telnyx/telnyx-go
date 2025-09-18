// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

// BillingGroupService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewBillingGroupService] method instead.
type BillingGroupService struct {
	Options []option.RequestOption
}

// NewBillingGroupService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewBillingGroupService(opts ...option.RequestOption) (r BillingGroupService) {
	r = BillingGroupService{}
	r.Options = opts
	return
}

// Create a billing group
func (r *BillingGroupService) New(ctx context.Context, body BillingGroupNewParams, opts ...option.RequestOption) (res *BillingGroupNewResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "billing_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a billing group
func (r *BillingGroupService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *BillingGroupGetResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("billing_groups/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a billing group
func (r *BillingGroupService) Update(ctx context.Context, id string, body BillingGroupUpdateParams, opts ...option.RequestOption) (res *BillingGroupUpdateResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("billing_groups/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// List all billing groups
func (r *BillingGroupService) List(ctx context.Context, query BillingGroupListParams, opts ...option.RequestOption) (res *BillingGroupListResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "billing_groups"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a billing group
func (r *BillingGroupService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (res *BillingGroupDeleteResponse, err error) {
	opts = append(r.Options[:], opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("billing_groups/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type BillingGroup struct {
	// Identifies the type of resource.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// ISO 8601 formatted date indicating when the resource was removed.
	DeletedAt time.Time `json:"deleted_at,nullable" format:"date-time"`
	// A user-specified name for the billing group
	Name string `json:"name"`
	// Identifies the organization that owns the resource.
	OrganizationID string `json:"organization_id" format:"uuid"`
	// Identifies the type of the resource.
	//
	// Any of "billing_group".
	RecordType BillingGroupRecordType `json:"record_type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CreatedAt      respjson.Field
		DeletedAt      respjson.Field
		Name           respjson.Field
		OrganizationID respjson.Field
		RecordType     respjson.Field
		UpdatedAt      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingGroup) RawJSON() string { return r.JSON.raw }
func (r *BillingGroup) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the resource.
type BillingGroupRecordType string

const (
	BillingGroupRecordTypeBillingGroup BillingGroupRecordType = "billing_group"
)

type BillingGroupNewResponse struct {
	Data BillingGroup `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingGroupNewResponse) RawJSON() string { return r.JSON.raw }
func (r *BillingGroupNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillingGroupGetResponse struct {
	Data BillingGroup `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingGroupGetResponse) RawJSON() string { return r.JSON.raw }
func (r *BillingGroupGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillingGroupUpdateResponse struct {
	Data BillingGroup `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingGroupUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *BillingGroupUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillingGroupListResponse struct {
	Data []BillingGroup `json:"data"`
	Meta PaginationMeta `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingGroupListResponse) RawJSON() string { return r.JSON.raw }
func (r *BillingGroupListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillingGroupDeleteResponse struct {
	Data BillingGroup `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r BillingGroupDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *BillingGroupDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillingGroupNewParams struct {
	// A name for the billing group
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r BillingGroupNewParams) MarshalJSON() (data []byte, err error) {
	type shadow BillingGroupNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BillingGroupNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillingGroupUpdateParams struct {
	// A name for the billing group
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r BillingGroupUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow BillingGroupUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *BillingGroupUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type BillingGroupListParams struct {
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page BillingGroupListParamsPage `query:"page,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BillingGroupListParams]'s query parameters as `url.Values`.
func (r BillingGroupListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type BillingGroupListParamsPage struct {
	// The page number to load
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// The size of the page
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [BillingGroupListParamsPage]'s query parameters as
// `url.Values`.
func (r BillingGroupListParamsPage) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
