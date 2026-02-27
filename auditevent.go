// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Audit log operations.
//
// AuditEventService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAuditEventService] method instead.
type AuditEventService struct {
	Options []option.RequestOption
}

// NewAuditEventService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAuditEventService(opts ...option.RequestOption) (r AuditEventService) {
	r = AuditEventService{}
	r.Options = opts
	return
}

// Retrieve a list of audit log entries. Audit logs are a best-effort, eventually
// consistent record of significant account-related changes.
func (r *AuditEventService) List(ctx context.Context, query AuditEventListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[AuditEventListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "audit_events"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve a list of audit log entries. Audit logs are a best-effort, eventually
// consistent record of significant account-related changes.
func (r *AuditEventService) ListAutoPaging(ctx context.Context, query AuditEventListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[AuditEventListResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

type AuditEventListResponse struct {
	// Unique identifier for the audit log entry.
	ID string `json:"id" format:"uuid"`
	// An alternate identifier for a resource which may be considered unique enough to
	// identify the resource but is not the primary identifier for the resource. For
	// example, this field could be used to store the phone number value for a phone
	// number when the primary database identifier is a separate distinct value.
	AlternateResourceID string `json:"alternate_resource_id" api:"nullable"`
	// Indicates if the change was made by Telnyx on your behalf, the organization
	// owner, a member of your organization, or in the case of managed accounts, the
	// account manager.
	//
	// Any of "telnyx", "account_manager", "account_owner", "organization_member".
	ChangeMadeBy AuditEventListResponseChangeMadeBy `json:"change_made_by"`
	// The type of change that occurred.
	ChangeType string `json:"change_type"`
	// Details of the changes made to the resource.
	Changes []AuditEventListResponseChange `json:"changes" api:"nullable"`
	// ISO 8601 formatted date indicating when the change occurred.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// Unique identifier for the organization that owns the resource.
	OrganizationID string `json:"organization_id" format:"uuid"`
	// The type of the resource being audited.
	RecordType string `json:"record_type"`
	// Unique identifier for the resource that was changed.
	ResourceID string `json:"resource_id"`
	// Unique identifier for the user who made the change.
	UserID string `json:"user_id" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AlternateResourceID respjson.Field
		ChangeMadeBy        respjson.Field
		ChangeType          respjson.Field
		Changes             respjson.Field
		CreatedAt           respjson.Field
		OrganizationID      respjson.Field
		RecordType          respjson.Field
		ResourceID          respjson.Field
		UserID              respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditEventListResponse) RawJSON() string { return r.JSON.raw }
func (r *AuditEventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates if the change was made by Telnyx on your behalf, the organization
// owner, a member of your organization, or in the case of managed accounts, the
// account manager.
type AuditEventListResponseChangeMadeBy string

const (
	AuditEventListResponseChangeMadeByTelnyx             AuditEventListResponseChangeMadeBy = "telnyx"
	AuditEventListResponseChangeMadeByAccountManager     AuditEventListResponseChangeMadeBy = "account_manager"
	AuditEventListResponseChangeMadeByAccountOwner       AuditEventListResponseChangeMadeBy = "account_owner"
	AuditEventListResponseChangeMadeByOrganizationMember AuditEventListResponseChangeMadeBy = "organization_member"
)

// Details of the changes made to a resource.
type AuditEventListResponseChange struct {
	// The name of the field that was changed. May use the dot notation to indicate
	// nested fields.
	Field string `json:"field"`
	// The previous value of the field. Can be any JSON type.
	From AuditEventListResponseChangeFromUnion `json:"from" api:"nullable"`
	// The new value of the field. Can be any JSON type.
	To AuditEventListResponseChangeToUnion `json:"to" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Field       respjson.Field
		From        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditEventListResponseChange) RawJSON() string { return r.JSON.raw }
func (r *AuditEventListResponseChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AuditEventListResponseChangeFromUnion contains all possible properties and
// values from [string], [float64], [bool], [map[string]any], [[]map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAuditEventListResponseChangeFromChangesObjectItem OfMapOfAnyMap]
type AuditEventListResponseChangeFromUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAuditEventListResponseChangeFromChangesObjectItem any `json:",inline"`
	// This field will be present if the value is a [[]map[string]any] instead of an
	// object.
	OfMapOfAnyMap []map[string]any `json:",inline"`
	JSON          struct {
		OfString                                            respjson.Field
		OfFloat                                             respjson.Field
		OfBool                                              respjson.Field
		OfAuditEventListResponseChangeFromChangesObjectItem respjson.Field
		OfMapOfAnyMap                                       respjson.Field
		raw                                                 string
	} `json:"-"`
}

func (u AuditEventListResponseChangeFromUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AuditEventListResponseChangeFromUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AuditEventListResponseChangeFromUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AuditEventListResponseChangeFromUnion) AsChangesObject() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AuditEventListResponseChangeFromUnion) AsMapOfAnyMap() (v []map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AuditEventListResponseChangeFromUnion) RawJSON() string { return u.JSON.raw }

func (r *AuditEventListResponseChangeFromUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AuditEventListResponseChangeToUnion contains all possible properties and values
// from [string], [float64], [bool], [map[string]any], [[]map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAuditEventListResponseChangeToChangesObjectItem OfMapOfAnyMap]
type AuditEventListResponseChangeToUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAuditEventListResponseChangeToChangesObjectItem any `json:",inline"`
	// This field will be present if the value is a [[]map[string]any] instead of an
	// object.
	OfMapOfAnyMap []map[string]any `json:",inline"`
	JSON          struct {
		OfString                                          respjson.Field
		OfFloat                                           respjson.Field
		OfBool                                            respjson.Field
		OfAuditEventListResponseChangeToChangesObjectItem respjson.Field
		OfMapOfAnyMap                                     respjson.Field
		raw                                               string
	} `json:"-"`
}

func (u AuditEventListResponseChangeToUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AuditEventListResponseChangeToUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AuditEventListResponseChangeToUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AuditEventListResponseChangeToUnion) AsChangesObject() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AuditEventListResponseChangeToUnion) AsMapOfAnyMap() (v []map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AuditEventListResponseChangeToUnion) RawJSON() string { return u.JSON.raw }

func (r *AuditEventListResponseChangeToUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuditEventListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[created_before], filter[created_after]
	Filter AuditEventListParamsFilter `query:"filter,omitzero" json:"-"`
	// Set the order of the results by the creation date.
	//
	// Any of "asc", "desc".
	Sort AuditEventListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AuditEventListParams]'s query parameters as `url.Values`.
func (r AuditEventListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally:
// filter[created_before], filter[created_after]
type AuditEventListParamsFilter struct {
	// Filter for audit events created after a specific date.
	CreatedAfter param.Opt[time.Time] `query:"created_after,omitzero" format:"date-time" json:"-"`
	// Filter for audit events created before a specific date.
	CreatedBefore param.Opt[time.Time] `query:"created_before,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [AuditEventListParamsFilter]'s query parameters as
// `url.Values`.
func (r AuditEventListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Set the order of the results by the creation date.
type AuditEventListParamsSort string

const (
	AuditEventListParamsSortAsc  AuditEventListParamsSort = "asc"
	AuditEventListParamsSortDesc AuditEventListParamsSort = "desc"
)
