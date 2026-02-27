// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Queue commands operations
//
// QueueService contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueueService] method instead.
type QueueService struct {
	Options []option.RequestOption
	// Queue commands operations
	Calls QueueCallService
}

// NewQueueService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewQueueService(opts ...option.RequestOption) (r QueueService) {
	r = QueueService{}
	r.Options = opts
	r.Calls = NewQueueCallService(opts...)
	return
}

// Create a new call queue.
func (r *QueueService) New(ctx context.Context, body QueueNewParams, opts ...option.RequestOption) (res *QueueNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "queues"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve an existing call queue
func (r *QueueService) Get(ctx context.Context, queueName string, opts ...option.RequestOption) (res *QueueGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if queueName == "" {
		err = errors.New("missing required queue_name parameter")
		return
	}
	path := fmt.Sprintf("queues/%s", queueName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update properties of an existing call queue.
func (r *QueueService) Update(ctx context.Context, queueName string, body QueueUpdateParams, opts ...option.RequestOption) (res *QueueUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if queueName == "" {
		err = errors.New("missing required queue_name parameter")
		return
	}
	path := fmt.Sprintf("queues/%s", queueName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all queues for the authenticated user.
func (r *QueueService) List(ctx context.Context, query QueueListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[Queue], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "queues"
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

// List all queues for the authenticated user.
func (r *QueueService) ListAutoPaging(ctx context.Context, query QueueListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[Queue] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Delete an existing call queue.
func (r *QueueService) Delete(ctx context.Context, queueName string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if queueName == "" {
		err = errors.New("missing required queue_name parameter")
		return
	}
	path := fmt.Sprintf("queues/%s", queueName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type Queue struct {
	// Uniquely identifies the queue
	ID string `json:"id" api:"required"`
	// The average time that the calls currently in the queue have spent waiting, given
	// in seconds.
	AverageWaitTimeSecs int64 `json:"average_wait_time_secs" api:"required"`
	// ISO 8601 formatted date of when the queue was created
	CreatedAt string `json:"created_at" api:"required"`
	// The number of calls currently in the queue
	CurrentSize int64 `json:"current_size" api:"required"`
	// The maximum number of calls allowed in the queue
	MaxSize int64 `json:"max_size" api:"required"`
	// Name of the queue
	Name string `json:"name" api:"required"`
	// Any of "queue".
	RecordType QueueRecordType `json:"record_type" api:"required"`
	// ISO 8601 formatted date of when the queue was last updated
	UpdatedAt string `json:"updated_at" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		AverageWaitTimeSecs respjson.Field
		CreatedAt           respjson.Field
		CurrentSize         respjson.Field
		MaxSize             respjson.Field
		Name                respjson.Field
		RecordType          respjson.Field
		UpdatedAt           respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Queue) RawJSON() string { return r.JSON.raw }
func (r *Queue) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueueRecordType string

const (
	QueueRecordTypeQueue QueueRecordType = "queue"
)

type QueueNewResponse struct {
	Data Queue `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueueNewResponse) RawJSON() string { return r.JSON.raw }
func (r *QueueNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueueGetResponse struct {
	Data Queue `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueueGetResponse) RawJSON() string { return r.JSON.raw }
func (r *QueueGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueueUpdateResponse struct {
	Data Queue `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QueueUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *QueueUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueueNewParams struct {
	// The name of the queue. Must be between 1 and 255 characters.
	QueueName string `json:"queue_name" api:"required"`
	// The maximum number of calls allowed in the queue.
	MaxSize param.Opt[int64] `json:"max_size,omitzero"`
	paramObj
}

func (r QueueNewParams) MarshalJSON() (data []byte, err error) {
	type shadow QueueNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueueNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueueUpdateParams struct {
	// The maximum number of calls allowed in the queue.
	MaxSize int64 `json:"max_size" api:"required"`
	paramObj
}

func (r QueueUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow QueueUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QueueUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QueueListParams struct {
	// The page number to load
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// The size of the page
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [QueueListParams]'s query parameters as `url.Values`.
func (r QueueListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
