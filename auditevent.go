// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/internal/apijson"
	"github.com/team-telnyx/telnyx-go/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
	"github.com/team-telnyx/telnyx-go/packages/param"
	"github.com/team-telnyx/telnyx-go/packages/respjson"
)

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
func (r *AuditEventService) List(ctx context.Context, query AuditEventListParams, opts ...option.RequestOption) (res *AuditEventListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "audit_events"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type AuditEventListResponse struct {
	Data []AuditEventListResponseData `json:"data"`
	Meta AuditEventListResponseMeta   `json:"meta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		Meta        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AuditEventListResponse) RawJSON() string { return r.JSON.raw }
func (r *AuditEventListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuditEventListResponseData struct {
	// Unique identifier for the audit log entry.
	ID string `json:"id" format:"uuid"`
	// An alternate identifier for a resource which may be considered unique enough to
	// identify the resource but is not the primary identifier for the resource. For
	// example, this field could be used to store the phone number value for a phone
	// number when the primary database identifier is a separate distinct value.
	AlternateResourceID string `json:"alternate_resource_id,nullable"`
	// Indicates if the change was made by Telnyx on your behalf, the organization
	// owner, a member of your organization, or in the case of managed accounts, the
	// account manager.
	//
	// Any of "telnyx", "account_manager", "account_owner", "organization_member".
	ChangeMadeBy string `json:"change_made_by"`
	// The type of change that occurred.
	ChangeType string `json:"change_type"`
	// Details of the changes made to the resource.
	Changes []AuditEventListResponseDataChange `json:"changes,nullable"`
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
func (r AuditEventListResponseData) RawJSON() string { return r.JSON.raw }
func (r *AuditEventListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details of the changes made to a resource.
type AuditEventListResponseDataChange struct {
	// The name of the field that was changed. May use the dot notation to indicate
	// nested fields.
	Field string `json:"field"`
	// The previous value of the field. Can be any JSON type.
	From AuditEventListResponseDataChangeFromUnion `json:"from,nullable"`
	// The new value of the field. Can be any JSON type.
	To AuditEventListResponseDataChangeToUnion `json:"to,nullable"`
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
func (r AuditEventListResponseDataChange) RawJSON() string { return r.JSON.raw }
func (r *AuditEventListResponseDataChange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AuditEventListResponseDataChangeFromUnion contains all possible properties and
// values from [string], [float64], [bool], [map[string]any], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAuditEventListResponseDataChangeFromMapItem OfAnyArray]
type AuditEventListResponseDataChangeFromUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAuditEventListResponseDataChangeFromMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                      respjson.Field
		OfFloat                                       respjson.Field
		OfBool                                        respjson.Field
		OfAuditEventListResponseDataChangeFromMapItem respjson.Field
		OfAnyArray                                    respjson.Field
		raw                                           string
	} `json:"-"`
}

func (u AuditEventListResponseDataChangeFromUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AuditEventListResponseDataChangeFromUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AuditEventListResponseDataChangeFromUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AuditEventListResponseDataChangeFromUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AuditEventListResponseDataChangeFromUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AuditEventListResponseDataChangeFromUnion) RawJSON() string { return u.JSON.raw }

func (r *AuditEventListResponseDataChangeFromUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AuditEventListResponseDataChangeToUnion contains all possible properties and
// values from [string], [float64], [bool], [map[string]any], [[]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfFloat OfBool
// OfAuditEventListResponseDataChangeToMapItem OfAnyArray]
type AuditEventListResponseDataChangeToUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAuditEventListResponseDataChangeToMapItem any `json:",inline"`
	// This field will be present if the value is a [[]any] instead of an object.
	OfAnyArray []any `json:",inline"`
	JSON       struct {
		OfString                                    respjson.Field
		OfFloat                                     respjson.Field
		OfBool                                      respjson.Field
		OfAuditEventListResponseDataChangeToMapItem respjson.Field
		OfAnyArray                                  respjson.Field
		raw                                         string
	} `json:"-"`
}

func (u AuditEventListResponseDataChangeToUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AuditEventListResponseDataChangeToUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AuditEventListResponseDataChangeToUnion) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AuditEventListResponseDataChangeToUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AuditEventListResponseDataChangeToUnion) AsAnyArray() (v []any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AuditEventListResponseDataChangeToUnion) RawJSON() string { return u.JSON.raw }

func (r *AuditEventListResponseDataChangeToUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuditEventListResponseMeta struct {
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
func (r AuditEventListResponseMeta) RawJSON() string { return r.JSON.raw }
func (r *AuditEventListResponseMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AuditEventListParams struct {
	// Consolidated filter parameter (deepObject style). Originally:
	// filter[created_before], filter[created_after]
	Filter AuditEventListParamsFilter `query:"filter,omitzero" json:"-"`
	// Consolidated page parameter (deepObject style). Originally: page[number],
	// page[size]
	Page AuditEventListParamsPage `query:"page,omitzero" json:"-"`
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

// Consolidated page parameter (deepObject style). Originally: page[number],
// page[size]
type AuditEventListParamsPage struct {
	// Page number to load.
	Number param.Opt[int64] `query:"number,omitzero" json:"-"`
	// Number of items per page.
	Size param.Opt[int64] `query:"size,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AuditEventListParamsPage]'s query parameters as
// `url.Values`.
func (r AuditEventListParamsPage) URLQuery() (v url.Values, err error) {
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
