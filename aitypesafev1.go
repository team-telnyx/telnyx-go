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
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
)

// Beta API for evaluating shared context with typed questions and structured
// answers. Telnyx manages model selection.
//
// AITypesafeV1Service contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAITypesafeV1Service] method instead.
type AITypesafeV1Service struct {
	Options []option.RequestOption
}

// NewAITypesafeV1Service generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAITypesafeV1Service(opts ...option.RequestOption) (r AITypesafeV1Service) {
	r = AITypesafeV1Service{}
	r.Options = opts
	return
}

// **Beta API.** Telnyx controls model selection.
//
// Evaluate shared context using named choice, noul (yes/no), and score questions.
// Returns TypeSafe System One-compatible answer shapes, an opaque compatibility
// identifier, and token usage. See the
// [decision model guide](https://developers.telnyx.com/docs/inference/decision-models)
// for examples and compatibility limits.
//
// The supported request subset requires instructions for every question, string
// descriptions for criteria (or null for choice descriptions), 1–64 questions, and
// 2–64 options for choice and score questions. The SDK-supplied model value is
// ignored and cannot select a model. Other unknown fields are rejected. The
// endpoint is synchronous and does not stream.
//
// Use the TypeSafe Python SDK with base_url set to
// https://api.telnyx.com/v2/ai/typesafe and a Telnyx API key. The SDK appends
// /v1/systemone. Compatibility covers this operation and the documented request
// subset; it does not include TypeSafe model listing. Scores describe relative
// preference, not calibrated correctness.
func (r *AITypesafeV1Service) Systemone(ctx context.Context, body AITypesafeV1SystemoneParams, opts ...option.RequestOption) (res *AITypesafeV1SystemoneResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "ai/typesafe/v1/systemone"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// A complete synchronous evaluation. Answers are returned directly without a data
// wrapper.
type AITypesafeV1SystemoneResponse struct {
	// Answers keyed by exactly the question IDs in the request. Each answer type
	// matches its question.
	Answers map[string]AITypesafeV1SystemoneResponseAnswersUnion `json:"answers" api:"required"`
	// Opaque Telnyx-controlled identifier retained for TypeSafe SDK response
	// compatibility. It is not a selectable model name or a guarantee of a particular
	// underlying model.
	Model string `json:"model" api:"required"`
	// Token usage for the completed evaluation.
	Usage AITypesafeV1SystemoneResponseUsage `json:"usage" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Answers     respjson.Field
		Model       respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITypesafeV1SystemoneResponse) RawJSON() string { return r.JSON.raw }
func (r *AITypesafeV1SystemoneResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AITypesafeV1SystemoneResponseAnswersUnion contains all possible properties and
// values from [AITypesafeV1SystemoneResponseAnswersChoice],
// [AITypesafeV1SystemoneResponseAnswersNoul],
// [AITypesafeV1SystemoneResponseAnswersScore].
//
// Use the [AITypesafeV1SystemoneResponseAnswersUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AITypesafeV1SystemoneResponseAnswersUnion struct {
	// This field is from variant [AITypesafeV1SystemoneResponseAnswersChoice].
	Choice        string  `json:"choice"`
	Confidence    float64 `json:"confidence"`
	Probabilities float64 `json:"probabilities"`
	// Any of "choice", "noul", "score".
	Type string `json:"type"`
	// This field is from variant [AITypesafeV1SystemoneResponseAnswersNoul].
	Noul float64 `json:"noul"`
	// This field is from variant [AITypesafeV1SystemoneResponseAnswersScore].
	Legend map[string]string `json:"legend"`
	// This field is from variant [AITypesafeV1SystemoneResponseAnswersScore].
	Score float64 `json:"score"`
	JSON  struct {
		Choice        respjson.Field
		Confidence    respjson.Field
		Probabilities respjson.Field
		Type          respjson.Field
		Noul          respjson.Field
		Legend        respjson.Field
		Score         respjson.Field
		raw           string
	} `json:"-"`
}

// anyAITypesafeV1SystemoneResponseAnswer is implemented by each variant of
// [AITypesafeV1SystemoneResponseAnswersUnion] to add type safety for the return
// type of [AITypesafeV1SystemoneResponseAnswersUnion.AsAny]
type anyAITypesafeV1SystemoneResponseAnswer interface {
	implAITypesafeV1SystemoneResponseAnswersUnion()
}

func (AITypesafeV1SystemoneResponseAnswersChoice) implAITypesafeV1SystemoneResponseAnswersUnion() {}
func (AITypesafeV1SystemoneResponseAnswersNoul) implAITypesafeV1SystemoneResponseAnswersUnion()   {}
func (AITypesafeV1SystemoneResponseAnswersScore) implAITypesafeV1SystemoneResponseAnswersUnion()  {}

// Use the following switch statement to find the correct variant
//
//	switch variant := AITypesafeV1SystemoneResponseAnswersUnion.AsAny().(type) {
//	case telnyx.AITypesafeV1SystemoneResponseAnswersChoice:
//	case telnyx.AITypesafeV1SystemoneResponseAnswersNoul:
//	case telnyx.AITypesafeV1SystemoneResponseAnswersScore:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u AITypesafeV1SystemoneResponseAnswersUnion) AsAny() anyAITypesafeV1SystemoneResponseAnswer {
	switch u.Type {
	case "choice":
		return u.AsChoice()
	case "noul":
		return u.AsNoul()
	case "score":
		return u.AsScore()
	}
	return nil
}

func (u AITypesafeV1SystemoneResponseAnswersUnion) AsChoice() (v AITypesafeV1SystemoneResponseAnswersChoice) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITypesafeV1SystemoneResponseAnswersUnion) AsNoul() (v AITypesafeV1SystemoneResponseAnswersNoul) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITypesafeV1SystemoneResponseAnswersUnion) AsScore() (v AITypesafeV1SystemoneResponseAnswersScore) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AITypesafeV1SystemoneResponseAnswersUnion) RawJSON() string { return u.JSON.raw }

func (r *AITypesafeV1SystemoneResponseAnswersUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A selected option and the distribution across all supplied option keys.
type AITypesafeV1SystemoneResponseAnswersChoice struct {
	// The option key with the highest relative score. Ties favor the first option in
	// request order.
	Choice string `json:"choice" api:"required"`
	// Normalized entropy confidence: 1 - H(p) / ln(N), where H(p) = -sum(p \* ln(p))
	// and N is the number of options. Zero indicates a uniform distribution; one
	// indicates concentration on one option. This is neither the winning probability
	// nor calibrated correctness.
	Confidence float64 `json:"confidence" api:"required"`
	// Relative scores normalized across the supplied options, summing approximately
	// to 1. These are not calibrated probabilities of correctness.
	Probabilities map[string]float64 `json:"probabilities" api:"required"`
	// Answer type.
	Type constant.Choice `json:"type" default:"choice"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Choice        respjson.Field
		Confidence    respjson.Field
		Probabilities respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITypesafeV1SystemoneResponseAnswersChoice) RawJSON() string { return r.JSON.raw }
func (r *AITypesafeV1SystemoneResponseAnswersChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A yes/no score with no separate confidence or probabilities fields.
type AITypesafeV1SystemoneResponseAnswersNoul struct {
	// Score of the positive outcome. Values near 1 favor yes; values near 0 favor no.
	// This is a number, not a Boolean, and is not calibrated correctness.
	Noul float64 `json:"noul" api:"required"`
	// Answer type.
	Type constant.Noul `json:"type" default:"noul"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Noul        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITypesafeV1SystemoneResponseAnswersNoul) RawJSON() string { return r.JSON.raw }
func (r *AITypesafeV1SystemoneResponseAnswersNoul) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An expected rating over the ordered criteria.
type AITypesafeV1SystemoneResponseAnswersScore struct {
	// Normalized entropy confidence: 1 - H(p) / ln(N), where H(p) = -sum(p \* ln(p))
	// and N is the number of options. Zero indicates a uniform distribution; one
	// indicates concentration on one option. This is neither the winning probability
	// nor calibrated correctness.
	Confidence float64 `json:"confidence" api:"required"`
	// Criterion descriptions keyed by stringified zero-based indices, such as "0",
	// "1", and "2".
	Legend map[string]string `json:"legend" api:"required"`
	// Relative scores keyed by the same stringified indices as legend.
	Probabilities map[string]float64 `json:"probabilities" api:"required"`
	// Expected zero-based criterion index: sum(index \* probability). Ranges from 0 to
	// N-1 for N criteria; fractional values are valid.
	Score float64 `json:"score" api:"required"`
	// Answer type.
	Type constant.Score `json:"type" default:"score"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence    respjson.Field
		Legend        respjson.Field
		Probabilities respjson.Field
		Score         respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITypesafeV1SystemoneResponseAnswersScore) RawJSON() string { return r.JSON.raw }
func (r *AITypesafeV1SystemoneResponseAnswersScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token usage for the completed evaluation.
type AITypesafeV1SystemoneResponseUsage struct {
	// Input tokens processed, including shared-context preparation and question
	// evaluation. This can exceed the token count of the unique input text.
	InputTokens int64 `json:"input_tokens" api:"required"`
	// Output tokens used for the evaluation, including shared-context preparation.
	OutputTokens int64 `json:"output_tokens" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InputTokens  respjson.Field
		OutputTokens respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITypesafeV1SystemoneResponseUsage) RawJSON() string { return r.JSON.raw }
func (r *AITypesafeV1SystemoneResponseUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITypesafeV1SystemoneParams struct {
	// Between 1 and 64 named questions. Each key identifies the corresponding answer.
	Questions map[string]AITypesafeV1SystemoneParamsQuestionsUnion `json:"questions,omitzero" api:"required"`
	// Shared context evaluated by every question.
	State AITypesafeV1SystemoneParamsStateUnion `json:"state,omitzero" api:"required"`
	paramObj
}

func (r AITypesafeV1SystemoneParams) MarshalJSON() (data []byte, err error) {
	type shadow AITypesafeV1SystemoneParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AITypesafeV1SystemoneParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AITypesafeV1SystemoneParamsQuestionsUnion struct {
	OfChoice *AITypesafeV1SystemoneParamsQuestionsChoice `json:",omitzero,inline"`
	OfNoul   *AITypesafeV1SystemoneParamsQuestionsNoul   `json:",omitzero,inline"`
	OfScore  *AITypesafeV1SystemoneParamsQuestionsScore  `json:",omitzero,inline"`
	paramUnion
}

func (u AITypesafeV1SystemoneParamsQuestionsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfChoice, u.OfNoul, u.OfScore)
}
func (u *AITypesafeV1SystemoneParamsQuestionsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AITypesafeV1SystemoneParamsQuestionsUnion) asAny() any {
	if !param.IsOmitted(u.OfChoice) {
		return u.OfChoice
	} else if !param.IsOmitted(u.OfNoul) {
		return u.OfNoul
	} else if !param.IsOmitted(u.OfScore) {
		return u.OfScore
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u AITypesafeV1SystemoneParamsQuestionsUnion) GetType() *string {
	if vt := u.OfChoice; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfNoul; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfScore; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u AITypesafeV1SystemoneParamsQuestionsUnion) GetCriteria() (res aiTypesafeV1SystemoneParamsQuestionsUnionCriteria) {
	if vt := u.OfChoice; vt != nil {
		res.any = &vt.Criteria
	} else if vt := u.OfNoul; vt != nil {
		res.any = &vt.Criteria
	} else if vt := u.OfScore; vt != nil {
		res.any = &vt.Criteria
	}
	return
}

// Can have the runtime types [*map[string]string],
// [*AITypesafeV1SystemoneParamsQuestionsNoulCriteria], [\*[]string]
type aiTypesafeV1SystemoneParamsQuestionsUnionCriteria struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *map[string]string:
//	case *telnyx.AITypesafeV1SystemoneParamsQuestionsNoulCriteria:
//	case *[]string:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u aiTypesafeV1SystemoneParamsQuestionsUnionCriteria) AsAny() any { return u.any }

// Returns a subunion which exports methods to access subproperties
//
// Or use AsAny() to get the underlying value
func (u AITypesafeV1SystemoneParamsQuestionsUnion) GetInstructions() (res aiTypesafeV1SystemoneParamsQuestionsUnionInstructions) {
	if vt := u.OfChoice; vt != nil {
		res.any = vt.Instructions.asAny()
	} else if vt := u.OfNoul; vt != nil {
		res.any = vt.Instructions.asAny()
	} else if vt := u.OfScore; vt != nil {
		res.any = vt.Instructions.asAny()
	}
	return
}

// Can have the runtime types [*string], [*any], [\*[]any]
type aiTypesafeV1SystemoneParamsQuestionsUnionInstructions struct{ any }

// Use the following switch statement to get the type of the union:
//
//	switch u.AsAny().(type) {
//	case *string:
//	case *any:
//	case *[]any:
//	default:
//	    fmt.Errorf("not present")
//	}
func (u aiTypesafeV1SystemoneParamsQuestionsUnionInstructions) AsAny() any { return u.any }

func init() {
	apijson.RegisterUnion[AITypesafeV1SystemoneParamsQuestionsUnion](
		"type",
		apijson.Discriminator[AITypesafeV1SystemoneParamsQuestionsChoice]("choice"),
		apijson.Discriminator[AITypesafeV1SystemoneParamsQuestionsNoul]("noul"),
		apijson.Discriminator[AITypesafeV1SystemoneParamsQuestionsScore]("score"),
	)
}

// Select one of the supplied options.
//
// The properties Criteria, Instructions, Type are required.
type AITypesafeV1SystemoneParamsQuestionsChoice struct {
	// Between 2 and 64 option keys mapped to description strings or null. A null
	// description uses the option key as its text.
	Criteria map[string]string `json:"criteria,omitzero" api:"required"`
	// Required instructions describing what to decide about the shared state.
	Instructions AITypesafeV1SystemoneParamsQuestionsChoiceInstructionsUnion `json:"instructions,omitzero" api:"required"`
	// Question type.
	//
	// This field can be elided, and will marshal its zero value as "choice".
	Type constant.Choice `json:"type" default:"choice"`
	paramObj
}

func (r AITypesafeV1SystemoneParamsQuestionsChoice) MarshalJSON() (data []byte, err error) {
	type shadow AITypesafeV1SystemoneParamsQuestionsChoice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AITypesafeV1SystemoneParamsQuestionsChoice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AITypesafeV1SystemoneParamsQuestionsChoiceInstructionsUnion struct {
	OfString   param.Opt[string] `json:",omitzero,inline"`
	OfAnyMap   map[string]any    `json:",omitzero,inline"`
	OfAnyArray []any             `json:",omitzero,inline"`
	paramUnion
}

func (u AITypesafeV1SystemoneParamsQuestionsChoiceInstructionsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfAnyMap, u.OfAnyArray)
}
func (u *AITypesafeV1SystemoneParamsQuestionsChoiceInstructionsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AITypesafeV1SystemoneParamsQuestionsChoiceInstructionsUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	}
	return nil
}

// Evaluate a yes/no question. Omit criteria to use Yes and No descriptions.
//
// The properties Instructions, Type are required.
type AITypesafeV1SystemoneParamsQuestionsNoul struct {
	// Required instructions describing what to decide about the shared state.
	Instructions AITypesafeV1SystemoneParamsQuestionsNoulInstructionsUnion `json:"instructions,omitzero" api:"required"`
	// Optional descriptions for the positive and negative outcomes. Descriptions must
	// be strings.
	Criteria AITypesafeV1SystemoneParamsQuestionsNoulCriteria `json:"criteria,omitzero"`
	// Question type.
	//
	// This field can be elided, and will marshal its zero value as "noul".
	Type constant.Noul `json:"type" default:"noul"`
	paramObj
}

func (r AITypesafeV1SystemoneParamsQuestionsNoul) MarshalJSON() (data []byte, err error) {
	type shadow AITypesafeV1SystemoneParamsQuestionsNoul
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AITypesafeV1SystemoneParamsQuestionsNoul) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AITypesafeV1SystemoneParamsQuestionsNoulInstructionsUnion struct {
	OfString   param.Opt[string] `json:",omitzero,inline"`
	OfAnyMap   map[string]any    `json:",omitzero,inline"`
	OfAnyArray []any             `json:",omitzero,inline"`
	paramUnion
}

func (u AITypesafeV1SystemoneParamsQuestionsNoulInstructionsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfAnyMap, u.OfAnyArray)
}
func (u *AITypesafeV1SystemoneParamsQuestionsNoulInstructionsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AITypesafeV1SystemoneParamsQuestionsNoulInstructionsUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	}
	return nil
}

// Optional descriptions for the positive and negative outcomes. Descriptions must
// be strings.
type AITypesafeV1SystemoneParamsQuestionsNoulCriteria struct {
	// Description of the negative outcome.
	False param.Opt[string] `json:"false,omitzero"`
	// Description of the positive outcome.
	True param.Opt[string] `json:"true,omitzero"`
	paramObj
}

func (r AITypesafeV1SystemoneParamsQuestionsNoulCriteria) MarshalJSON() (data []byte, err error) {
	type shadow AITypesafeV1SystemoneParamsQuestionsNoulCriteria
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AITypesafeV1SystemoneParamsQuestionsNoulCriteria) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rate the state against an ordered rubric.
//
// The properties Criteria, Instructions, Type are required.
type AITypesafeV1SystemoneParamsQuestionsScore struct {
	// Between 2 and 64 description strings in ascending score order. Indices start at
	// zero.
	Criteria []string `json:"criteria,omitzero" api:"required"`
	// Required instructions describing what to decide about the shared state.
	Instructions AITypesafeV1SystemoneParamsQuestionsScoreInstructionsUnion `json:"instructions,omitzero" api:"required"`
	// Question type.
	//
	// This field can be elided, and will marshal its zero value as "score".
	Type constant.Score `json:"type" default:"score"`
	paramObj
}

func (r AITypesafeV1SystemoneParamsQuestionsScore) MarshalJSON() (data []byte, err error) {
	type shadow AITypesafeV1SystemoneParamsQuestionsScore
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AITypesafeV1SystemoneParamsQuestionsScore) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AITypesafeV1SystemoneParamsQuestionsScoreInstructionsUnion struct {
	OfString   param.Opt[string] `json:",omitzero,inline"`
	OfAnyMap   map[string]any    `json:",omitzero,inline"`
	OfAnyArray []any             `json:",omitzero,inline"`
	paramUnion
}

func (u AITypesafeV1SystemoneParamsQuestionsScoreInstructionsUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfAnyMap, u.OfAnyArray)
}
func (u *AITypesafeV1SystemoneParamsQuestionsScoreInstructionsUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AITypesafeV1SystemoneParamsQuestionsScoreInstructionsUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	}
	return nil
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AITypesafeV1SystemoneParamsStateUnion struct {
	OfString   param.Opt[string] `json:",omitzero,inline"`
	OfAnyMap   map[string]any    `json:",omitzero,inline"`
	OfAnyArray []any             `json:",omitzero,inline"`
	paramUnion
}

func (u AITypesafeV1SystemoneParamsStateUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfAnyMap, u.OfAnyArray)
}
func (u *AITypesafeV1SystemoneParamsStateUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AITypesafeV1SystemoneParamsStateUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	} else if !param.IsOmitted(u.OfAnyArray) {
		return &u.OfAnyArray
	}
	return nil
}
