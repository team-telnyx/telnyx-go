// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Create and retrieve asynchronous summaries and action-item artifacts.
//
// MeetingSessionArtifactService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMeetingSessionArtifactService] method instead.
type MeetingSessionArtifactService struct {
	Options []option.RequestOption
}

// NewMeetingSessionArtifactService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewMeetingSessionArtifactService(opts ...option.RequestOption) (r MeetingSessionArtifactService) {
	r = MeetingSessionArtifactService{}
	r.Options = opts
	return
}

// Requests asynchronous generation of one `summary` or `action_items` artifact.
// Each type requires its own request. Generation requires transcript content and
// configured inference and currently reads at most the first 10,000 segments, so
// exceptionally long transcripts may produce incomplete artifacts or fail model
// limits.
func (r *MeetingSessionArtifactService) New(ctx context.Context, id string, body MeetingSessionArtifactNewParams, opts ...option.RequestOption) (res *MeetingSessionArtifactResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("meeting_sessions/%s/artifacts", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves a single meeting session artifact by ID.
func (r *MeetingSessionArtifactService) Get(ctx context.Context, artifactID string, query MeetingSessionArtifactGetParams, opts ...option.RequestOption) (res *MeetingSessionArtifactResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ID == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	if artifactID == "" {
		err = errors.New("missing required artifact_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("meeting_sessions/%s/artifacts/%s", query.ID, artifactID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns a list of artifacts for a meeting session.
func (r *MeetingSessionArtifactService) List(ctx context.Context, id string, opts ...option.RequestOption) (res *MeetingSessionArtifactListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("meeting_sessions/%s/artifacts", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type MeetingSessionArtifact struct {
	ID              string                                `json:"id" api:"required"`
	Content         MeetingSessionArtifactContent         `json:"content" api:"required"`
	CreatedAt       time.Time                             `json:"created_at" api:"required" format:"date-time"`
	FailureReason   string                                `json:"failure_reason" api:"required"`
	ModelProvenance MeetingSessionArtifactModelProvenance `json:"model_provenance" api:"required"`
	SessionID       string                                `json:"session_id" api:"required"`
	// Any of "pending", "completed", "failed".
	Status MeetingSessionArtifactStatus `json:"status" api:"required"`
	// Any of "summary", "action_items".
	Type      MeetingSessionArtifactType `json:"type" api:"required"`
	UpdatedAt time.Time                  `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID              respjson.Field
		Content         respjson.Field
		CreatedAt       respjson.Field
		FailureReason   respjson.Field
		ModelProvenance respjson.Field
		SessionID       respjson.Field
		Status          respjson.Field
		Type            respjson.Field
		UpdatedAt       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionArtifact) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionArtifact) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionArtifactContent struct {
	Text string `json:"text" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionArtifactContent) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionArtifactContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionArtifactModelProvenance struct {
	Model    string `json:"model" api:"required"`
	Provider string `json:"provider" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Model       respjson.Field
		Provider    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionArtifactModelProvenance) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionArtifactModelProvenance) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionArtifactStatus string

const (
	MeetingSessionArtifactStatusPending   MeetingSessionArtifactStatus = "pending"
	MeetingSessionArtifactStatusCompleted MeetingSessionArtifactStatus = "completed"
	MeetingSessionArtifactStatusFailed    MeetingSessionArtifactStatus = "failed"
)

type MeetingSessionArtifactType string

const (
	MeetingSessionArtifactTypeSummary     MeetingSessionArtifactType = "summary"
	MeetingSessionArtifactTypeActionItems MeetingSessionArtifactType = "action_items"
)

type MeetingSessionArtifactResponse struct {
	Data MeetingSessionArtifact `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionArtifactResponse) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionArtifactResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionArtifactListResponse struct {
	Data []MeetingSessionArtifact `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MeetingSessionArtifactListResponse) RawJSON() string { return r.JSON.raw }
func (r *MeetingSessionArtifactListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type MeetingSessionArtifactNewParams struct {
	// Type of artifact to generate from the session.
	//
	// Any of "summary", "action_items".
	Type MeetingSessionArtifactNewParamsType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r MeetingSessionArtifactNewParams) MarshalJSON() (data []byte, err error) {
	type shadow MeetingSessionArtifactNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MeetingSessionArtifactNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of artifact to generate from the session.
type MeetingSessionArtifactNewParamsType string

const (
	MeetingSessionArtifactNewParamsTypeSummary     MeetingSessionArtifactNewParamsType = "summary"
	MeetingSessionArtifactNewParamsTypeActionItems MeetingSessionArtifactNewParamsType = "action_items"
)

type MeetingSessionArtifactGetParams struct {
	ID string `path:"id" api:"required" json:"-"`
	paramObj
}
