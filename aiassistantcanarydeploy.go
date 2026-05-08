// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	shimjson "github.com/team-telnyx/telnyx-go/v4/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Configure AI assistant specifications
//
// AIAssistantCanaryDeployService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIAssistantCanaryDeployService] method instead.
type AIAssistantCanaryDeployService struct {
	Options []option.RequestOption
}

// NewAIAssistantCanaryDeployService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAIAssistantCanaryDeployService(opts ...option.RequestOption) (r AIAssistantCanaryDeployService) {
	r = AIAssistantCanaryDeployService{}
	r.Options = opts
	return
}

// Endpoint to create a canary deploy configuration for an assistant.
//
// Creates a new canary deploy configuration with multiple version IDs and their
// traffic percentages for A/B testing or gradual rollouts of assistant versions.
func (r *AIAssistantCanaryDeployService) New(ctx context.Context, assistantID string, body AIAssistantCanaryDeployNewParams, opts ...option.RequestOption) (res *CanaryDeployResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/assistants/%s/canary-deploys", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Endpoint to get a canary deploy configuration for an assistant.
//
// Retrieves the current canary deploy configuration with all version IDs and their
// traffic percentages for the specified assistant.
func (r *AIAssistantCanaryDeployService) Get(ctx context.Context, assistantID string, opts ...option.RequestOption) (res *CanaryDeployResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/assistants/%s/canary-deploys", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Endpoint to update a canary deploy configuration for an assistant.
//
// Updates the existing canary deploy configuration with new version IDs and
// percentages. All old versions and percentages are replaces by new ones from this
// request.
func (r *AIAssistantCanaryDeployService) Update(ctx context.Context, assistantID string, body AIAssistantCanaryDeployUpdateParams, opts ...option.RequestOption) (res *CanaryDeployResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/assistants/%s/canary-deploys", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Endpoint to delete a canary deploy configuration for an assistant.
//
// Removes all canary deploy configurations for the specified assistant.
func (r *AIAssistantCanaryDeployService) Delete(ctx context.Context, assistantID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if assistantID == "" {
		err = errors.New("missing required assistant_id parameter")
		return err
	}
	path := fmt.Sprintf("ai/assistants/%s/canary-deploys", assistantID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Create/update request body. Accepts:
//
// - `rules` — canonical ordered list of routing rules
type CanaryDeployParam struct {
	Rules []RuleInputParam `json:"rules,omitzero"`
	paramObj
}

func (r CanaryDeployParam) MarshalJSON() (data []byte, err error) {
	type shadow CanaryDeployParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CanaryDeployParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response shape.
//
// Always carries `rules` (canonical).
type CanaryDeployResponse struct {
	AssistantID string       `json:"assistant_id" api:"required"`
	CreatedAt   time.Time    `json:"created_at" api:"required" format:"date-time"`
	Rules       []RuleOutput `json:"rules" api:"required"`
	UpdatedAt   time.Time    `json:"updated_at" api:"required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AssistantID respjson.Field
		CreatedAt   respjson.Field
		Rules       respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CanaryDeployResponse) RawJSON() string { return r.JSON.raw }
func (r *CanaryDeployResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single attribute/operator/values check.
//
// A clause matches when the routing context's value for `attribute` satisfies
// `operator` against any of `values`.
type Clause struct {
	// Attribute name from the routing context
	Attribute string `json:"attribute" api:"required"`
	// Match operator
	//
	// Any of "in", "not_in", "starts_with".
	Operator ClauseOperator `json:"operator" api:"required"`
	Values   []string       `json:"values" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Attribute   respjson.Field
		Operator    respjson.Field
		Values      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Clause) RawJSON() string { return r.JSON.raw }
func (r *Clause) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Clause to a ClauseParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ClauseParam.Overrides()
func (r Clause) ToParam() ClauseParam {
	return param.Override[ClauseParam](json.RawMessage(r.RawJSON()))
}

// Match operator
type ClauseOperator string

const (
	ClauseOperatorIn         ClauseOperator = "in"
	ClauseOperatorNotIn      ClauseOperator = "not_in"
	ClauseOperatorStartsWith ClauseOperator = "starts_with"
)

// A single attribute/operator/values check.
//
// A clause matches when the routing context's value for `attribute` satisfies
// `operator` against any of `values`.
//
// The properties Attribute, Operator, Values are required.
type ClauseParam struct {
	// Attribute name from the routing context
	Attribute string `json:"attribute" api:"required"`
	// Match operator
	//
	// Any of "in", "not_in", "starts_with".
	Operator ClauseOperator `json:"operator,omitzero" api:"required"`
	Values   []string       `json:"values,omitzero" api:"required"`
	paramObj
}

func (r ClauseParam) MarshalJSON() (data []byte, err error) {
	type shadow ClauseParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClauseParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One slot in a percentage rollout.
type RolloutSlot struct {
	VersionID string  `json:"version_id" api:"required"`
	Weight    float64 `json:"weight" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		VersionID   respjson.Field
		Weight      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RolloutSlot) RawJSON() string { return r.JSON.raw }
func (r *RolloutSlot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RolloutSlot to a RolloutSlotParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RolloutSlotParam.Overrides()
func (r RolloutSlot) ToParam() RolloutSlotParam {
	return param.Override[RolloutSlotParam](json.RawMessage(r.RawJSON()))
}

// One slot in a percentage rollout.
//
// The properties VersionID, Weight are required.
type RolloutSlotParam struct {
	VersionID string  `json:"version_id" api:"required"`
	Weight    float64 `json:"weight" api:"required"`
	paramObj
}

func (r RolloutSlotParam) MarshalJSON() (data []byte, err error) {
	type shadow RolloutSlotParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RolloutSlotParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A targeting rule: `match` clauses (AND) gate `serve`.
//
// An empty `match` is a catch-all (always fires).
//
// The property Serve is required.
type RuleInputParam struct {
	// What a rule serves when matched.
	//
	// Exactly one of:
	//
	//   - `version_id` — serve a specific version
	//   - `rollout` — weighted random across versions; weights must sum to less than
	//     100, with the leftover routing to the main version
	Serve ServeParam    `json:"serve,omitzero" api:"required"`
	Match []ClauseParam `json:"match,omitzero"`
	paramObj
}

func (r RuleInputParam) MarshalJSON() (data []byte, err error) {
	type shadow RuleInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RuleInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A targeting rule: `match` clauses (AND) gate `serve`.
//
// An empty `match` is a catch-all (always fires).
type RuleOutput struct {
	// What a rule serves when matched.
	//
	// Exactly one of:
	//
	//   - `version_id` — serve a specific version
	//   - `rollout` — weighted random across versions; weights must sum to less than
	//     100, with the leftover routing to the main version
	Serve Serve    `json:"serve" api:"required"`
	Match []Clause `json:"match"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Serve       respjson.Field
		Match       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RuleOutput) RawJSON() string { return r.JSON.raw }
func (r *RuleOutput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// What a rule serves when matched.
//
// Exactly one of:
//
//   - `version_id` — serve a specific version
//   - `rollout` — weighted random across versions; weights must sum to less than
//     100, with the leftover routing to the main version
type Serve struct {
	Rollout   []RolloutSlot `json:"rollout"`
	VersionID string        `json:"version_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Rollout     respjson.Field
		VersionID   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Serve) RawJSON() string { return r.JSON.raw }
func (r *Serve) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Serve to a ServeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ServeParam.Overrides()
func (r Serve) ToParam() ServeParam {
	return param.Override[ServeParam](json.RawMessage(r.RawJSON()))
}

// What a rule serves when matched.
//
// Exactly one of:
//
//   - `version_id` — serve a specific version
//   - `rollout` — weighted random across versions; weights must sum to less than
//     100, with the leftover routing to the main version
type ServeParam struct {
	VersionID param.Opt[string]  `json:"version_id,omitzero"`
	Rollout   []RolloutSlotParam `json:"rollout,omitzero"`
	paramObj
}

func (r ServeParam) MarshalJSON() (data []byte, err error) {
	type shadow ServeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ServeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantCanaryDeployNewParams struct {
	// Create/update request body. Accepts:
	//
	// - `rules` — canonical ordered list of routing rules
	CanaryDeploy CanaryDeployParam
	paramObj
}

func (r AIAssistantCanaryDeployNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CanaryDeploy)
}
func (r *AIAssistantCanaryDeployNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIAssistantCanaryDeployUpdateParams struct {
	// Create/update request body. Accepts:
	//
	// - `rules` — canonical ordered list of routing rules
	CanaryDeploy CanaryDeployParam
	paramObj
}

func (r AIAssistantCanaryDeployUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CanaryDeploy)
}
func (r *AIAssistantCanaryDeployUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
