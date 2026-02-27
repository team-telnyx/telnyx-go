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
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Background jobs performed over a phone-numbers block's phone numbers
//
// PhoneNumberBlockJobService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPhoneNumberBlockJobService] method instead.
type PhoneNumberBlockJobService struct {
	Options []option.RequestOption
}

// NewPhoneNumberBlockJobService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPhoneNumberBlockJobService(opts ...option.RequestOption) (r PhoneNumberBlockJobService) {
	r = PhoneNumberBlockJobService{}
	r.Options = opts
	return
}

// Retrieves a phone number blocks job
func (r *PhoneNumberBlockJobService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PhoneNumberBlockJobGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return
	}
	path := fmt.Sprintf("phone_number_blocks/jobs/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Lists the phone number blocks jobs
func (r *PhoneNumberBlockJobService) List(ctx context.Context, query PhoneNumberBlockJobListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[Job], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "phone_number_blocks/jobs"
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

// Lists the phone number blocks jobs
func (r *PhoneNumberBlockJobService) ListAutoPaging(ctx context.Context, query PhoneNumberBlockJobListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[Job] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Creates a new background job to delete all the phone numbers associated with the
// given block. We will only consider the phone number block as deleted after all
// phone numbers associated with it are removed, so multiple executions of this job
// may be necessary in case some of the phone numbers present errors during the
// deletion process.
func (r *PhoneNumberBlockJobService) DeletePhoneNumberBlock(ctx context.Context, body PhoneNumberBlockJobDeletePhoneNumberBlockParams, opts ...option.RequestOption) (res *PhoneNumberBlockJobDeletePhoneNumberBlockResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "phone_number_blocks/jobs/delete_phone_number_block"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type Job struct {
	// Identifies the resource.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 formatted date indicating when the resource was created.
	CreatedAt string `json:"created_at"`
	// ISO 8601 formatted date indicating when the estimated time of completion of the
	// background job.
	Etc              time.Time            `json:"etc" format:"date-time"`
	FailedOperations []JobFailedOperation `json:"failed_operations"`
	// Identifies the type of the resource.
	RecordType string `json:"record_type"`
	// Indicates the completion status of the background operation.
	//
	// Any of "pending", "in_progress", "completed", "failed".
	Status               JobStatus                `json:"status"`
	SuccessfulOperations []JobSuccessfulOperation `json:"successful_operations"`
	// Identifies the type of the background job.
	//
	// Any of "delete_phone_number_block".
	Type JobType `json:"type"`
	// ISO 8601 formatted date indicating when the resource was updated.
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                   respjson.Field
		CreatedAt            respjson.Field
		Etc                  respjson.Field
		FailedOperations     respjson.Field
		RecordType           respjson.Field
		Status               respjson.Field
		SuccessfulOperations respjson.Field
		Type                 respjson.Field
		UpdatedAt            respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Job) RawJSON() string { return r.JSON.raw }
func (r *Job) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobFailedOperation struct {
	// The phone number's ID
	ID     string     `json:"id"`
	Errors []JobError `json:"errors"`
	// The phone number in e164 format.
	PhoneNumber string `json:"phone_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Errors      respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobFailedOperation) RawJSON() string { return r.JSON.raw }
func (r *JobFailedOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicates the completion status of the background operation.
type JobStatus string

const (
	JobStatusPending    JobStatus = "pending"
	JobStatusInProgress JobStatus = "in_progress"
	JobStatusCompleted  JobStatus = "completed"
	JobStatusFailed     JobStatus = "failed"
)

// The phone numbers successfully updated.
type JobSuccessfulOperation struct {
	// The phone number's ID
	ID string `json:"id"`
	// The phone number in e164 format.
	PhoneNumber string `json:"phone_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobSuccessfulOperation) RawJSON() string { return r.JSON.raw }
func (r *JobSuccessfulOperation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the type of the background job.
type JobType string

const (
	JobTypeDeletePhoneNumberBlock JobType = "delete_phone_number_block"
)

type JobError struct {
	Code   string         `json:"code"`
	Detail string         `json:"detail"`
	Meta   JobErrorMeta   `json:"meta"`
	Source JobErrorSource `json:"source"`
	Title  string         `json:"title"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Code        respjson.Field
		Detail      respjson.Field
		Meta        respjson.Field
		Source      respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobError) RawJSON() string { return r.JSON.raw }
func (r *JobError) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobErrorMeta struct {
	// URL with additional information on the error.
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobErrorMeta) RawJSON() string { return r.JSON.raw }
func (r *JobErrorMeta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobErrorSource struct {
	// Indicates which query parameter caused the error.
	Parameter string `json:"parameter"`
	// JSON pointer (RFC6901) to the offending entity.
	Pointer string `json:"pointer"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parameter   respjson.Field
		Pointer     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobErrorSource) RawJSON() string { return r.JSON.raw }
func (r *JobErrorSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberBlockJobGetResponse struct {
	Data Job `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberBlockJobGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberBlockJobGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberBlockJobDeletePhoneNumberBlockResponse struct {
	Data Job `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PhoneNumberBlockJobDeletePhoneNumberBlockResponse) RawJSON() string { return r.JSON.raw }
func (r *PhoneNumberBlockJobDeletePhoneNumberBlockResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PhoneNumberBlockJobListParams struct {
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	PageSize   param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	// Consolidated filter parameter (deepObject style). Originally: filter[type],
	// filter[status]
	Filter PhoneNumberBlockJobListParamsFilter `query:"filter,omitzero" json:"-"`
	// Specifies the sort order for results. If not given, results are sorted by
	// created_at in descending order.
	//
	// Any of "created_at".
	Sort PhoneNumberBlockJobListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberBlockJobListParams]'s query parameters as
// `url.Values`.
func (r PhoneNumberBlockJobListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Consolidated filter parameter (deepObject style). Originally: filter[type],
// filter[status]
type PhoneNumberBlockJobListParamsFilter struct {
	// Identifies the status of the background job.
	//
	// Any of "pending", "in_progress", "completed", "failed".
	Status string `query:"status,omitzero" json:"-"`
	// Identifies the type of the background job.
	//
	// Any of "delete_phone_number_block".
	Type string `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PhoneNumberBlockJobListParamsFilter]'s query parameters as
// `url.Values`.
func (r PhoneNumberBlockJobListParamsFilter) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Specifies the sort order for results. If not given, results are sorted by
// created_at in descending order.
type PhoneNumberBlockJobListParamsSort string

const (
	PhoneNumberBlockJobListParamsSortCreatedAt PhoneNumberBlockJobListParamsSort = "created_at"
)

type PhoneNumberBlockJobDeletePhoneNumberBlockParams struct {
	PhoneNumberBlockID string `json:"phone_number_block_id" api:"required"`
	paramObj
}

func (r PhoneNumberBlockJobDeletePhoneNumberBlockParams) MarshalJSON() (data []byte, err error) {
	type shadow PhoneNumberBlockJobDeletePhoneNumberBlockParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PhoneNumberBlockJobDeletePhoneNumberBlockParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
