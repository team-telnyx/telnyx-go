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
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// QueueService contains methods and other services that help with interacting with
// the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQueueService] method instead.
type QueueService struct {
	Options []option.RequestOption
	Calls   QueueCallService
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

type QueueGetResponse struct {
	Data QueueGetResponseData `json:"data"`
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

type QueueGetResponseData struct {
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
	RecordType string `json:"record_type" api:"required"`
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
func (r QueueGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *QueueGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
