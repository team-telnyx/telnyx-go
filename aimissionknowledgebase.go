// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

// AIMissionKnowledgeBaseService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIMissionKnowledgeBaseService] method instead.
type AIMissionKnowledgeBaseService struct {
	Options []option.RequestOption
}

// NewAIMissionKnowledgeBaseService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIMissionKnowledgeBaseService(opts ...option.RequestOption) (r AIMissionKnowledgeBaseService) {
	r = AIMissionKnowledgeBaseService{}
	r.Options = opts
	return
}

// Create a new knowledge base for a mission
func (r *AIMissionKnowledgeBaseService) NewKnowledgeBase(ctx context.Context, missionID string, opts ...option.RequestOption) (res *AIMissionKnowledgeBaseNewKnowledgeBaseResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if missionID == "" {
		err = errors.New("missing required mission_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/missions/%s/knowledge-bases", missionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Detaches the specified knowledge base from the mission so its content is no
// longer available to agents in subsequent runs.
func (r *AIMissionKnowledgeBaseService) DeleteKnowledgeBase(ctx context.Context, knowledgeBaseID string, body AIMissionKnowledgeBaseDeleteKnowledgeBaseParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return err
	}
	if knowledgeBaseID == "" {
		err = errors.New("missing required knowledge_base_id parameter")
		return err
	}
	path := fmt.Sprintf("ai/missions/%s/knowledge-bases/%s", body.MissionID, knowledgeBaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Returns the details of a single knowledge base attached to the specified
// mission.
func (r *AIMissionKnowledgeBaseService) GetKnowledgeBase(ctx context.Context, knowledgeBaseID string, query AIMissionKnowledgeBaseGetKnowledgeBaseParams, opts ...option.RequestOption) (res *AIMissionKnowledgeBaseGetKnowledgeBaseResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return nil, err
	}
	if knowledgeBaseID == "" {
		err = errors.New("missing required knowledge_base_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/missions/%s/knowledge-bases/%s", query.MissionID, knowledgeBaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the knowledge bases attached to the specified mission. Knowledge bases
// provide reference content agents can draw on during runs.
func (r *AIMissionKnowledgeBaseService) ListKnowledgeBases(ctx context.Context, missionID string, opts ...option.RequestOption) (res *AIMissionKnowledgeBaseListKnowledgeBasesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if missionID == "" {
		err = errors.New("missing required mission_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/missions/%s/knowledge-bases", missionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Replaces the definition of the specified knowledge base on this mission.
func (r *AIMissionKnowledgeBaseService) UpdateKnowledgeBase(ctx context.Context, knowledgeBaseID string, body AIMissionKnowledgeBaseUpdateKnowledgeBaseParams, opts ...option.RequestOption) (res *AIMissionKnowledgeBaseUpdateKnowledgeBaseResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.MissionID == "" {
		err = errors.New("missing required mission_id parameter")
		return nil, err
	}
	if knowledgeBaseID == "" {
		err = errors.New("missing required knowledge_base_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/missions/%s/knowledge-bases/%s", body.MissionID, knowledgeBaseID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, nil, &res, opts...)
	return res, err
}

type AIMissionKnowledgeBaseNewKnowledgeBaseResponse = any

type AIMissionKnowledgeBaseGetKnowledgeBaseResponse = any

type AIMissionKnowledgeBaseListKnowledgeBasesResponse = any

type AIMissionKnowledgeBaseUpdateKnowledgeBaseResponse = any

type AIMissionKnowledgeBaseDeleteKnowledgeBaseParams struct {
	MissionID string `path:"mission_id" api:"required" json:"-"`
	paramObj
}

type AIMissionKnowledgeBaseGetKnowledgeBaseParams struct {
	MissionID string `path:"mission_id" api:"required" json:"-"`
	paramObj
}

type AIMissionKnowledgeBaseUpdateKnowledgeBaseParams struct {
	MissionID string `path:"mission_id" api:"required" json:"-"`
	paramObj
}
