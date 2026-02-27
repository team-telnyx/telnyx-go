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
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Customize LLMs for your unique needs
//
// AIFineTuningJobService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIFineTuningJobService] method instead.
type AIFineTuningJobService struct {
	Options []option.RequestOption
}

// NewAIFineTuningJobService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAIFineTuningJobService(opts ...option.RequestOption) (r AIFineTuningJobService) {
	r = AIFineTuningJobService{}
	r.Options = opts
	return
}

// Create a new fine tuning job.
func (r *AIFineTuningJobService) New(ctx context.Context, body AIFineTuningJobNewParams, opts ...option.RequestOption) (res *FineTuningJob, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/fine_tuning/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve a fine tuning job by `job_id`.
func (r *AIFineTuningJobService) Get(ctx context.Context, jobID string, opts ...option.RequestOption) (res *FineTuningJob, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("ai/fine_tuning/jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a list of all fine tuning jobs created by the user.
func (r *AIFineTuningJobService) List(ctx context.Context, opts ...option.RequestOption) (res *AIFineTuningJobListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/fine_tuning/jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Cancel a fine tuning job.
func (r *AIFineTuningJobService) Cancel(ctx context.Context, jobID string, opts ...option.RequestOption) (res *FineTuningJob, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required job_id parameter")
		return
	}
	path := fmt.Sprintf("ai/fine_tuning/jobs/%s/cancel", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// The `fine_tuning.job` object represents a fine-tuning job that has been created
// through the API.
type FineTuningJob struct {
	// The name of the fine-tuned model that is being created.
	ID string `json:"id" api:"required"`
	// The Unix timestamp (in seconds) for when the fine-tuning job was created.
	CreatedAt int64 `json:"created_at" api:"required"`
	// The Unix timestamp (in seconds) for when the fine-tuning job was finished. The
	// value will be null if the fine-tuning job is still running.
	FinishedAt int64 `json:"finished_at" api:"required"`
	// The hyperparameters used for the fine-tuning job.
	Hyperparameters FineTuningJobHyperparameters `json:"hyperparameters" api:"required"`
	// The base model that is being fine-tuned.
	Model string `json:"model" api:"required"`
	// The organization that owns the fine-tuning job.
	OrganizationID string `json:"organization_id" api:"required"`
	// The current status of the fine-tuning job.
	//
	// Any of "queued", "running", "succeeded", "failed", "cancelled".
	Status FineTuningJobStatus `json:"status" api:"required"`
	// The total number of billable tokens processed by this fine-tuning job. The value
	// will be null if the fine-tuning job is still running.
	TrainedTokens int64 `json:"trained_tokens" api:"required"`
	// The storage bucket or object used for training.
	TrainingFile string `json:"training_file" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		CreatedAt       respjson.Field
		FinishedAt      respjson.Field
		Hyperparameters respjson.Field
		Model           respjson.Field
		OrganizationID  respjson.Field
		Status          respjson.Field
		TrainedTokens   respjson.Field
		TrainingFile    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningJob) RawJSON() string { return r.JSON.raw }
func (r *FineTuningJob) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The hyperparameters used for the fine-tuning job.
type FineTuningJobHyperparameters struct {
	// The number of epochs to train the model for. An epoch refers to one full cycle
	// through the training dataset.
	NEpochs int64 `json:"n_epochs" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		NEpochs     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FineTuningJobHyperparameters) RawJSON() string { return r.JSON.raw }
func (r *FineTuningJobHyperparameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The current status of the fine-tuning job.
type FineTuningJobStatus string

const (
	FineTuningJobStatusQueued    FineTuningJobStatus = "queued"
	FineTuningJobStatusRunning   FineTuningJobStatus = "running"
	FineTuningJobStatusSucceeded FineTuningJobStatus = "succeeded"
	FineTuningJobStatusFailed    FineTuningJobStatus = "failed"
	FineTuningJobStatusCancelled FineTuningJobStatus = "cancelled"
)

type AIFineTuningJobListResponse struct {
	Data []FineTuningJob `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIFineTuningJobListResponse) RawJSON() string { return r.JSON.raw }
func (r *AIFineTuningJobListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIFineTuningJobNewParams struct {
	// The base model that is being fine-tuned.
	Model string `json:"model" api:"required"`
	// The storage bucket or object used for training.
	TrainingFile string `json:"training_file" api:"required"`
	// Optional suffix to append to the fine tuned model's name.
	Suffix param.Opt[string] `json:"suffix,omitzero"`
	// The hyperparameters used for the fine-tuning job.
	Hyperparameters AIFineTuningJobNewParamsHyperparameters `json:"hyperparameters,omitzero"`
	paramObj
}

func (r AIFineTuningJobNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AIFineTuningJobNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIFineTuningJobNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The hyperparameters used for the fine-tuning job.
type AIFineTuningJobNewParamsHyperparameters struct {
	// The number of epochs to train the model for. An epoch refers to one full cycle
	// through the training dataset. 'auto' decides the optimal number of epochs based
	// on the size of the dataset. If setting the number manually, we support any
	// number between 1 and 50 epochs.
	NEpochs param.Opt[int64] `json:"n_epochs,omitzero"`
	paramObj
}

func (r AIFineTuningJobNewParamsHyperparameters) MarshalJSON() (data []byte, err error) {
	type shadow AIFineTuningJobNewParamsHyperparameters
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIFineTuningJobNewParamsHyperparameters) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
