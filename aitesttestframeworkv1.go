// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Test agents, evaluator scenarios, metrics, and run results for the AI test
// framework.
//
// AITestTestFrameworkV1Service contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAITestTestFrameworkV1Service] method instead.
type AITestTestFrameworkV1Service struct {
	Options []option.RequestOption
	// Test agents, evaluator scenarios, metrics, and run results for the AI test
	// framework.
	Aiagents AITestTestFrameworkV1AiagentService
	// Test agents, evaluator scenarios, metrics, and run results for the AI test
	// framework.
	Metrics AITestTestFrameworkV1MetricService
	// Test agents, evaluator scenarios, metrics, and run results for the AI test
	// framework.
	Results AITestTestFrameworkV1ResultService
	// Test agents, evaluator scenarios, metrics, and run results for the AI test
	// framework.
	Scenarios AITestTestFrameworkV1ScenarioService
}

// NewAITestTestFrameworkV1Service generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewAITestTestFrameworkV1Service(opts ...option.RequestOption) (r AITestTestFrameworkV1Service) {
	r = AITestTestFrameworkV1Service{}
	r.Options = opts
	r.Aiagents = NewAITestTestFrameworkV1AiagentService(opts...)
	r.Metrics = NewAITestTestFrameworkV1MetricService(opts...)
	r.Results = NewAITestTestFrameworkV1ResultService(opts...)
	r.Scenarios = NewAITestTestFrameworkV1ScenarioService(opts...)
	return
}

// Returns the catalogue of test personalities available to your organization. A
// personality describes how a simulated caller speaks: prompt, language, accent,
// voice, interruption level, and call-behaviour settings.
func (r *AITestTestFrameworkV1Service) GetPersonalities(ctx context.Context, opts ...option.RequestOption) (res *[]AITestTestFrameworkV1GetPersonalitiesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/tests/test_framework/v1/personalities"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type AITestTestFrameworkV1GetPersonalitiesResponse struct {
	ID             int64 `json:"id" api:"required"`
	EndCallEnabled bool  `json:"end_call_enabled" api:"required"`
	// ISO language code the simulated caller speaks.
	Language string `json:"language" api:"required"`
	// Display name of the personality.
	Name string `json:"name" api:"required"`
	// Instructions the simulated caller follows.
	Prompt                string  `json:"prompt" api:"required"`
	Accent                string  `json:"accent"`
	BackgroundNoise       string  `json:"background_noise" api:"nullable"`
	BackgroundSoundVolume float64 `json:"background_sound_volume" api:"nullable"`
	EffectiveVoiceVolume  float64 `json:"effective_voice_volume" api:"nullable"`
	// ID of the personality this one was forked from.
	ForkedFrom       int64          `json:"forked_from" api:"nullable"`
	GenerationConfig map[string]any `json:"generation_config" api:"nullable"`
	// How aggressively the personality interrupts the assistant.
	InterruptionLevel string `json:"interruption_level" api:"nullable"`
	// Idle-timeout plan for messages during a call.
	MessagePlan       map[string]any                                                      `json:"message_plan" api:"nullable"`
	NetworkSimulation map[string]any                                                      `json:"network_simulation" api:"nullable"`
	Organization      int64                                                               `json:"organization" api:"nullable"`
	Project           int64                                                               `json:"project" api:"nullable"`
	Speed             AITestTestFrameworkV1GetPersonalitiesResponseSpeedUnion             `json:"speed" api:"nullable"`
	StartSpeakingPlan AITestTestFrameworkV1GetPersonalitiesResponseStartSpeakingPlanUnion `json:"start_speaking_plan" api:"nullable"`
	StopSpeakingPlan  map[string]any                                                      `json:"stop_speaking_plan" api:"nullable"`
	VoiceVolume       float64                                                             `json:"voice_volume" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                    respjson.Field
		EndCallEnabled        respjson.Field
		Language              respjson.Field
		Name                  respjson.Field
		Prompt                respjson.Field
		Accent                respjson.Field
		BackgroundNoise       respjson.Field
		BackgroundSoundVolume respjson.Field
		EffectiveVoiceVolume  respjson.Field
		ForkedFrom            respjson.Field
		GenerationConfig      respjson.Field
		InterruptionLevel     respjson.Field
		MessagePlan           respjson.Field
		NetworkSimulation     respjson.Field
		Organization          respjson.Field
		Project               respjson.Field
		Speed                 respjson.Field
		StartSpeakingPlan     respjson.Field
		StopSpeakingPlan      respjson.Field
		VoiceVolume           respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITestTestFrameworkV1GetPersonalitiesResponse) RawJSON() string { return r.JSON.raw }
func (r *AITestTestFrameworkV1GetPersonalitiesResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AITestTestFrameworkV1GetPersonalitiesResponseSpeedUnion contains all possible
// properties and values from [float64], [string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat OfString]
type AITestTestFrameworkV1GetPersonalitiesResponseSpeedUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	JSON     struct {
		OfFloat  respjson.Field
		OfString respjson.Field
		raw      string
	} `json:"-"`
}

func (u AITestTestFrameworkV1GetPersonalitiesResponseSpeedUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITestTestFrameworkV1GetPersonalitiesResponseSpeedUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AITestTestFrameworkV1GetPersonalitiesResponseSpeedUnion) RawJSON() string { return u.JSON.raw }

func (r *AITestTestFrameworkV1GetPersonalitiesResponseSpeedUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AITestTestFrameworkV1GetPersonalitiesResponseStartSpeakingPlanUnion contains all
// possible properties and values from [float64], [map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFloat
// OfAITestTestFrameworkV1GetPersonalitiesResponseStartSpeakingPlanUnionMember1]
type AITestTestFrameworkV1GetPersonalitiesResponseStartSpeakingPlanUnion struct {
	// This field will be present if the value is a [float64] instead of an object.
	OfFloat float64 `json:",inline"`
	// This field will be present if the value is a [any] instead of an object.
	OfAITestTestFrameworkV1GetPersonalitiesResponseStartSpeakingPlanUnionMember1 any `json:",inline"`
	JSON                                                                         struct {
		OfFloat                                                                      respjson.Field
		OfAITestTestFrameworkV1GetPersonalitiesResponseStartSpeakingPlanUnionMember1 respjson.Field
		raw                                                                          string
	} `json:"-"`
}

func (u AITestTestFrameworkV1GetPersonalitiesResponseStartSpeakingPlanUnion) AsFloat() (v float64) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITestTestFrameworkV1GetPersonalitiesResponseStartSpeakingPlanUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AITestTestFrameworkV1GetPersonalitiesResponseStartSpeakingPlanUnion) RawJSON() string {
	return u.JSON.raw
}

func (r *AITestTestFrameworkV1GetPersonalitiesResponseStartSpeakingPlanUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
