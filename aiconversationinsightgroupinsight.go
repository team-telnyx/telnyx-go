// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/option"
)

// AIConversationInsightGroupInsightService contains methods and other services
// that help with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIConversationInsightGroupInsightService] method instead.
type AIConversationInsightGroupInsightService struct {
	Options []option.RequestOption
}

// NewAIConversationInsightGroupInsightService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewAIConversationInsightGroupInsightService(opts ...option.RequestOption) (r AIConversationInsightGroupInsightService) {
	r = AIConversationInsightGroupInsightService{}
	r.Options = opts
	return
}

// Assign an insight to a group
func (r *AIConversationInsightGroupInsightService) Assign(ctx context.Context, insightID string, body AIConversationInsightGroupInsightAssignParams, opts ...option.RequestOption) (res *AIConversationInsightGroupInsightAssignResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.GroupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	if insightID == "" {
		err = errors.New("missing required insight_id parameter")
		return
	}
	path := fmt.Sprintf("ai/conversations/insight-groups/%s/insights/%s/assign", body.GroupID, insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Remove an insight from a group
func (r *AIConversationInsightGroupInsightService) DeleteUnassign(ctx context.Context, insightID string, body AIConversationInsightGroupInsightDeleteUnassignParams, opts ...option.RequestOption) (res *AIConversationInsightGroupInsightDeleteUnassignResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.GroupID == "" {
		err = errors.New("missing required group_id parameter")
		return
	}
	if insightID == "" {
		err = errors.New("missing required insight_id parameter")
		return
	}
	path := fmt.Sprintf("ai/conversations/insight-groups/%s/insights/%s/unassign", body.GroupID, insightID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type AIConversationInsightGroupInsightAssignResponse = any

type AIConversationInsightGroupInsightDeleteUnassignResponse = any

type AIConversationInsightGroupInsightAssignParams struct {
	// The ID of the insight group
	GroupID string `path:"group_id,required" format:"uuid" json:"-"`
	paramObj
}

type AIConversationInsightGroupInsightDeleteUnassignParams struct {
	// The ID of the insight group
	GroupID string `path:"group_id,required" format:"uuid" json:"-"`
	paramObj
}
