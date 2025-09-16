// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/telnyx-go/internal/apijson"
	"github.com/stainless-sdks/telnyx-go/internal/requestconfig"
	"github.com/stainless-sdks/telnyx-go/option"
	"github.com/stainless-sdks/telnyx-go/packages/respjson"
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
	opts = append(r.Options[:], opts...)
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
	ID string `json:"id,required"`
	// The average time that the calls currently in the queue have spent waiting, given
	// in seconds.
	AverageWaitTimeSecs int64 `json:"average_wait_time_secs,required"`
	// ISO 8601 formatted date of when the queue was created
	CreatedAt string `json:"created_at,required"`
	// The number of calls currently in the queue
	CurrentSize int64 `json:"current_size,required"`
	// The maximum number of calls allowed in the queue
	MaxSize int64 `json:"max_size,required"`
	// Name of the queue
	Name string `json:"name,required"`
	// Any of "queue".
	RecordType string `json:"record_type,required"`
	// ISO 8601 formatted date of when the queue was last updated
	UpdatedAt string `json:"updated_at,required"`
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
