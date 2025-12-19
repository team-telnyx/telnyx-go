// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// TexmlAccountQueueService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTexmlAccountQueueService] method instead.
type TexmlAccountQueueService struct {
	Options []option.RequestOption
}

// NewTexmlAccountQueueService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewTexmlAccountQueueService(opts ...option.RequestOption) (r TexmlAccountQueueService) {
	r = TexmlAccountQueueService{}
	r.Options = opts
	return
}

// Creates a new queue resource.
func (r *TexmlAccountQueueService) New(ctx context.Context, accountSid string, body TexmlAccountQueueNewParams, opts ...option.RequestOption) (res *TexmlAccountQueueNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if accountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Queues", accountSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a queue resource.
func (r *TexmlAccountQueueService) Get(ctx context.Context, queueSid string, query TexmlAccountQueueGetParams, opts ...option.RequestOption) (res *TexmlAccountQueueGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if queueSid == "" {
		err = errors.New("missing required queue_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Queues/%s", query.AccountSid, queueSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a queue resource.
func (r *TexmlAccountQueueService) Update(ctx context.Context, queueSid string, params TexmlAccountQueueUpdateParams, opts ...option.RequestOption) (res *TexmlAccountQueueUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if queueSid == "" {
		err = errors.New("missing required queue_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Queues/%s", params.AccountSid, queueSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Delete a queue resource.
func (r *TexmlAccountQueueService) Delete(ctx context.Context, queueSid string, body TexmlAccountQueueDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.AccountSid == "" {
		err = errors.New("missing required account_sid parameter")
		return
	}
	if queueSid == "" {
		err = errors.New("missing required queue_sid parameter")
		return
	}
	path := fmt.Sprintf("texml/Accounts/%s/Queues/%s", body.AccountSid, queueSid)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type TexmlAccountQueueNewResponse struct {
	// The id of the account the resource belongs to.
	AccountSid string `json:"account_sid"`
	// The average wait time in seconds for members in the queue.
	AverageWaitTime int64 `json:"average_wait_time"`
	// The current number of members in the queue.
	CurrentSize int64 `json:"current_size"`
	// The timestamp of when the resource was created.
	DateCreated string `json:"date_created"`
	// The timestamp of when the resource was last updated.
	DateUpdated string `json:"date_updated"`
	// The maximum size of the queue.
	MaxSize int64 `json:"max_size"`
	// The unique identifier of the queue.
	Sid string `json:"sid"`
	// A list of related resources identified by their relative URIs.
	SubresourceUris map[string]any `json:"subresource_uris"`
	// The relative URI for this queue.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid      respjson.Field
		AverageWaitTime respjson.Field
		CurrentSize     respjson.Field
		DateCreated     respjson.Field
		DateUpdated     respjson.Field
		MaxSize         respjson.Field
		Sid             respjson.Field
		SubresourceUris respjson.Field
		Uri             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountQueueNewResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountQueueNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountQueueGetResponse struct {
	// The id of the account the resource belongs to.
	AccountSid string `json:"account_sid"`
	// The average wait time in seconds for members in the queue.
	AverageWaitTime int64 `json:"average_wait_time"`
	// The current number of members in the queue.
	CurrentSize int64 `json:"current_size"`
	// The timestamp of when the resource was created.
	DateCreated string `json:"date_created"`
	// The timestamp of when the resource was last updated.
	DateUpdated string `json:"date_updated"`
	// The maximum size of the queue.
	MaxSize int64 `json:"max_size"`
	// The unique identifier of the queue.
	Sid string `json:"sid"`
	// A list of related resources identified by their relative URIs.
	SubresourceUris map[string]any `json:"subresource_uris"`
	// The relative URI for this queue.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid      respjson.Field
		AverageWaitTime respjson.Field
		CurrentSize     respjson.Field
		DateCreated     respjson.Field
		DateUpdated     respjson.Field
		MaxSize         respjson.Field
		Sid             respjson.Field
		SubresourceUris respjson.Field
		Uri             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountQueueGetResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountQueueGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountQueueUpdateResponse struct {
	// The id of the account the resource belongs to.
	AccountSid string `json:"account_sid"`
	// The average wait time in seconds for members in the queue.
	AverageWaitTime int64 `json:"average_wait_time"`
	// The current number of members in the queue.
	CurrentSize int64 `json:"current_size"`
	// The timestamp of when the resource was created.
	DateCreated string `json:"date_created"`
	// The timestamp of when the resource was last updated.
	DateUpdated string `json:"date_updated"`
	// The maximum size of the queue.
	MaxSize int64 `json:"max_size"`
	// The unique identifier of the queue.
	Sid string `json:"sid"`
	// A list of related resources identified by their relative URIs.
	SubresourceUris map[string]any `json:"subresource_uris"`
	// The relative URI for this queue.
	Uri string `json:"uri"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountSid      respjson.Field
		AverageWaitTime respjson.Field
		CurrentSize     respjson.Field
		DateCreated     respjson.Field
		DateUpdated     respjson.Field
		MaxSize         respjson.Field
		Sid             respjson.Field
		SubresourceUris respjson.Field
		Uri             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TexmlAccountQueueUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *TexmlAccountQueueUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountQueueNewParams struct {
	// A human readable name for the queue.
	FriendlyName param.Opt[string] `json:"FriendlyName,omitzero"`
	// The maximum size of the queue.
	MaxSize param.Opt[int64] `json:"MaxSize,omitzero"`
	paramObj
}

func (r TexmlAccountQueueNewParams) MarshalJSON() (data []byte, err error) {
	type shadow TexmlAccountQueueNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlAccountQueueNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountQueueGetParams struct {
	AccountSid string `path:"account_sid,required" json:"-"`
	paramObj
}

type TexmlAccountQueueUpdateParams struct {
	AccountSid string `path:"account_sid,required" json:"-"`
	// The maximum size of the queue.
	MaxSize param.Opt[int64] `json:"MaxSize,omitzero"`
	paramObj
}

func (r TexmlAccountQueueUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow TexmlAccountQueueUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TexmlAccountQueueUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TexmlAccountQueueDeleteParams struct {
	AccountSid string `path:"account_sid,required" json:"-"`
	paramObj
}
