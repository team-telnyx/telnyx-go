// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Configure AI assistant specifications
//
// AIToolService contains methods and other services that help with interacting
// with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAIToolService] method instead.
type AIToolService struct {
	Options []option.RequestOption
}

// NewAIToolService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAIToolService(opts ...option.RequestOption) (r AIToolService) {
	r = AIToolService{}
	r.Options = opts
	return
}

// Create a new custom AI tool that can be attached to AI assistants.
func (r *AIToolService) New(ctx context.Context, params AIToolNewParams, opts ...option.RequestOption) (res *SharedToolResponse, err error) {
	if !param.IsOmitted(params.IdempotencyKey) {
		opts = append(opts, option.WithHeader("Idempotency-Key", fmt.Sprintf("%v", params.IdempotencyKey.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "ai/tools"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve the details of a specific AI tool.
func (r *AIToolService) Get(ctx context.Context, toolID string, opts ...option.RequestOption) (res *SharedToolResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if toolID == "" {
		err = errors.New("missing required tool_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/tools/%s", url.PathEscape(toolID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update the configuration of an existing AI tool.
func (r *AIToolService) Update(ctx context.Context, toolID string, body AIToolUpdateParams, opts ...option.RequestOption) (res *SharedToolResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if toolID == "" {
		err = errors.New("missing required tool_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("ai/tools/%s", url.PathEscape(toolID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Retrieve a list of the custom AI tools configured on your account.
func (r *AIToolService) List(ctx context.Context, query AIToolListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[SharedToolResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "ai/tools"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Retrieve a list of the custom AI tools configured on your account.
func (r *AIToolService) ListAutoPaging(ctx context.Context, query AIToolListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[SharedToolResponse] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes the specified custom AI tool from your account.
func (r *AIToolService) Delete(ctx context.Context, toolID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if toolID == "" {
		err = errors.New("missing required tool_id parameter")
		return err
	}
	path := fmt.Sprintf("ai/tools/%s", url.PathEscape(toolID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

type PayToolParamsResp struct {
	// The name of the pay connector configured in the Telnyx API. Must reference an
	// existing pay connector for this organization.
	ConnectorName string `json:"connector_name" api:"required"`
	// Custom key/value parameters forwarded to the pay connector as `parameters` in
	// the processor request. String values may use `{{dynamic_variable}}` templating,
	// resolved per call.
	ConnectorParameters map[string]any `json:"connector_parameters"`
	// Default currency for payments processed by this tool.
	Currency string `json:"currency"`
	// Description of the pay tool shown to the assistant so it knows when to call the
	// tool. Not sent to the payment connector — the per-call `description` argument
	// the assistant passes is what reaches the processor.
	Description string `json:"description" api:"nullable"`
	// Time in milliseconds to wait between consecutive DTMF digits. Omitted means the
	// platform default (5000) applies.
	InterDigitTimeoutMillis int64 `json:"inter_digit_timeout_millis"`
	// Maximum number of attempts for each payment collection step. Omitted means the
	// platform default (3) applies.
	MaxAttempts int64 `json:"max_attempts"`
	// Minimum number of digits the caller must enter when collecting a credit-card
	// billing postal code. Applies only when `postal_code` prompts the caller. Omitted
	// means the platform default (1) applies. The maximum number of digits collected
	// is the greater of 10 and this value. This does not validate the length of a
	// postal code supplied directly in `postal_code`.
	MinPostalCodeLength int64 `json:"min_postal_code_length"`
	// Default payment method for payments processed by this tool.
	PaymentMethod string `json:"payment_method"`
	// Controls billing postal code collection for credit-card payments. Set to `false`
	// to skip the postal code prompt without supplying a postal code to the payment
	// connector (non-US cards, e.g. Irish Eircodes). Set to `true` to prompt the
	// caller. Supply a non-empty string, such as `"90210"`, to use that postal code
	// without prompting. This option does not affect ACH debit payments or payments
	// using an existing token.
	PostalCode PayToolParamsPostalCodeUnionResp `json:"postal_code"`
	// Custom text-to-speech prompts keyed by payment collection step. The AI assistant
	// merges its own card-number prompt into this map, matching on the same step keys.
	Prompts PayToolParamsPromptsResp `json:"prompts"`
	// Time in milliseconds to wait for DTMF input for each collection step. Omitted
	// means the platform default (5000) applies. Increase this for slower entry, e.g.
	// a 16-digit card number.
	TimeoutMillis int64 `json:"timeout_millis"`
	// Restricts accepted card numbers to the listed card types. When the caller enters
	// a card number that does not match one of the listed types, Pay treats the input
	// as invalid and re-prompts for the card number.
	//
	// Any of "visa", "mastercard", "amex", "maestro", "discover", "optima", "jcb",
	// "diners-club", "enroute".
	ValidCardTypes []string `json:"valid_card_types"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectorName           respjson.Field
		ConnectorParameters     respjson.Field
		Currency                respjson.Field
		Description             respjson.Field
		InterDigitTimeoutMillis respjson.Field
		MaxAttempts             respjson.Field
		MinPostalCodeLength     respjson.Field
		PaymentMethod           respjson.Field
		PostalCode              respjson.Field
		Prompts                 respjson.Field
		TimeoutMillis           respjson.Field
		ValidCardTypes          respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayToolParamsResp) RawJSON() string { return r.JSON.raw }
func (r *PayToolParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PayToolParamsResp to a PayToolParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PayToolParams.Overrides()
func (r PayToolParamsResp) ToParam() PayToolParams {
	return param.Override[PayToolParams](json.RawMessage(r.RawJSON()))
}

// PayToolParamsPostalCodeUnionResp contains all possible properties and values
// from [string], [bool].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfBool]
type PayToolParamsPostalCodeUnionResp struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [bool] instead of an object.
	OfBool bool `json:",inline"`
	JSON   struct {
		OfString respjson.Field
		OfBool   respjson.Field
		raw      string
	} `json:"-"`
}

func (u PayToolParamsPostalCodeUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PayToolParamsPostalCodeUnionResp) AsBool() (v bool) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PayToolParamsPostalCodeUnionResp) RawJSON() string { return u.JSON.raw }

func (r *PayToolParamsPostalCodeUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom text-to-speech prompts keyed by payment collection step. The AI assistant
// merges its own card-number prompt into this map, matching on the same step keys.
type PayToolParamsPromptsResp struct {
	// A default prompt string or an ordered list of qualified prompts.
	BankAccountNumber PayToolParamsPromptsBankAccountNumberUnionResp `json:"bank-account-number"`
	// A default prompt string or an ordered list of qualified prompts.
	BankRoutingNumber PayToolParamsPromptsBankRoutingNumberUnionResp `json:"bank-routing-number"`
	// A default prompt string or an ordered list of qualified prompts.
	ExpirationDate PayToolParamsPromptsExpirationDateUnionResp `json:"expiration-date"`
	// A default prompt string or an ordered list of qualified prompts.
	PaymentCardNumber PayToolParamsPromptsPaymentCardNumberUnionResp `json:"payment-card-number"`
	// A default prompt string or an ordered list of qualified prompts.
	PostalCode PayToolParamsPromptsPostalCodeUnionResp `json:"postal-code"`
	// A default prompt string or an ordered list of qualified prompts.
	SecurityCode PayToolParamsPromptsSecurityCodeUnionResp `json:"security-code"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		BankAccountNumber respjson.Field
		BankRoutingNumber respjson.Field
		ExpirationDate    respjson.Field
		PaymentCardNumber respjson.Field
		PostalCode        respjson.Field
		SecurityCode      respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayToolParamsPromptsResp) RawJSON() string { return r.JSON.raw }
func (r *PayToolParamsPromptsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PayToolParamsPromptsBankAccountNumberUnionResp contains all possible properties
// and values from [string],
// [[]PayToolParamsPromptsBankAccountNumberPayToolPromptListItemResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfPayToolPromptList]
type PayToolParamsPromptsBankAccountNumberUnionResp struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]PayToolParamsPromptsBankAccountNumberPayToolPromptListItemResp] instead of an
	// object.
	OfPayToolPromptList []PayToolParamsPromptsBankAccountNumberPayToolPromptListItemResp `json:",inline"`
	JSON                struct {
		OfString            respjson.Field
		OfPayToolPromptList respjson.Field
		raw                 string
	} `json:"-"`
}

func (u PayToolParamsPromptsBankAccountNumberUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PayToolParamsPromptsBankAccountNumberUnionResp) AsPayToolPromptList() (v []PayToolParamsPromptsBankAccountNumberPayToolPromptListItemResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PayToolParamsPromptsBankAccountNumberUnionResp) RawJSON() string { return u.JSON.raw }

func (r *PayToolParamsPromptsBankAccountNumberUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text-to-speech prompt with optional matching qualifiers, mirroring the Voice
// API PayPrompt schema.
type PayToolParamsPromptsBankAccountNumberPayToolPromptListItemResp struct {
	// Text spoken for the payment collection step.
	Text string `json:"text" api:"required"`
	// Space-separated 1-based attempt numbers (1-9 followed by digits only, e.g.
	// `"2 3"`) for which this prompt applies. A prompt qualified with an attempt
	// number above `max_attempts` (default 3) is rejected at write time — it could
	// never play.
	Attempt string `json:"attempt"`
	// Lowercase, case-sensitive detected card type for which this prompt applies.
	//
	// Any of "visa", "mastercard", "amex", "optima", "discover", "diners-club", "jcb",
	// "maestro", "enroute".
	CardType string `json:"card_type"`
	// Step error for which this prompt applies.
	//
	// Any of "timeout", "invalid-card-number", "invalid-card-type", "invalid-date",
	// "invalid-security-code", "invalid-postal-code", "invalid-bank-routing-number",
	// "invalid-bank-account-number", "input-matching-failed".
	ErrorType string `json:"error_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Attempt     respjson.Field
		CardType    respjson.Field
		ErrorType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayToolParamsPromptsBankAccountNumberPayToolPromptListItemResp) RawJSON() string {
	return r.JSON.raw
}
func (r *PayToolParamsPromptsBankAccountNumberPayToolPromptListItemResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PayToolParamsPromptsBankRoutingNumberUnionResp contains all possible properties
// and values from [string],
// [[]PayToolParamsPromptsBankRoutingNumberPayToolPromptListItemResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfPayToolPromptList]
type PayToolParamsPromptsBankRoutingNumberUnionResp struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]PayToolParamsPromptsBankRoutingNumberPayToolPromptListItemResp] instead of an
	// object.
	OfPayToolPromptList []PayToolParamsPromptsBankRoutingNumberPayToolPromptListItemResp `json:",inline"`
	JSON                struct {
		OfString            respjson.Field
		OfPayToolPromptList respjson.Field
		raw                 string
	} `json:"-"`
}

func (u PayToolParamsPromptsBankRoutingNumberUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PayToolParamsPromptsBankRoutingNumberUnionResp) AsPayToolPromptList() (v []PayToolParamsPromptsBankRoutingNumberPayToolPromptListItemResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PayToolParamsPromptsBankRoutingNumberUnionResp) RawJSON() string { return u.JSON.raw }

func (r *PayToolParamsPromptsBankRoutingNumberUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text-to-speech prompt with optional matching qualifiers, mirroring the Voice
// API PayPrompt schema.
type PayToolParamsPromptsBankRoutingNumberPayToolPromptListItemResp struct {
	// Text spoken for the payment collection step.
	Text string `json:"text" api:"required"`
	// Space-separated 1-based attempt numbers (1-9 followed by digits only, e.g.
	// `"2 3"`) for which this prompt applies. A prompt qualified with an attempt
	// number above `max_attempts` (default 3) is rejected at write time — it could
	// never play.
	Attempt string `json:"attempt"`
	// Lowercase, case-sensitive detected card type for which this prompt applies.
	//
	// Any of "visa", "mastercard", "amex", "optima", "discover", "diners-club", "jcb",
	// "maestro", "enroute".
	CardType string `json:"card_type"`
	// Step error for which this prompt applies.
	//
	// Any of "timeout", "invalid-card-number", "invalid-card-type", "invalid-date",
	// "invalid-security-code", "invalid-postal-code", "invalid-bank-routing-number",
	// "invalid-bank-account-number", "input-matching-failed".
	ErrorType string `json:"error_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Attempt     respjson.Field
		CardType    respjson.Field
		ErrorType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayToolParamsPromptsBankRoutingNumberPayToolPromptListItemResp) RawJSON() string {
	return r.JSON.raw
}
func (r *PayToolParamsPromptsBankRoutingNumberPayToolPromptListItemResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PayToolParamsPromptsExpirationDateUnionResp contains all possible properties and
// values from [string],
// [[]PayToolParamsPromptsExpirationDatePayToolPromptListItemResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfPayToolPromptList]
type PayToolParamsPromptsExpirationDateUnionResp struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]PayToolParamsPromptsExpirationDatePayToolPromptListItemResp] instead of an
	// object.
	OfPayToolPromptList []PayToolParamsPromptsExpirationDatePayToolPromptListItemResp `json:",inline"`
	JSON                struct {
		OfString            respjson.Field
		OfPayToolPromptList respjson.Field
		raw                 string
	} `json:"-"`
}

func (u PayToolParamsPromptsExpirationDateUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PayToolParamsPromptsExpirationDateUnionResp) AsPayToolPromptList() (v []PayToolParamsPromptsExpirationDatePayToolPromptListItemResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PayToolParamsPromptsExpirationDateUnionResp) RawJSON() string { return u.JSON.raw }

func (r *PayToolParamsPromptsExpirationDateUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text-to-speech prompt with optional matching qualifiers, mirroring the Voice
// API PayPrompt schema.
type PayToolParamsPromptsExpirationDatePayToolPromptListItemResp struct {
	// Text spoken for the payment collection step.
	Text string `json:"text" api:"required"`
	// Space-separated 1-based attempt numbers (1-9 followed by digits only, e.g.
	// `"2 3"`) for which this prompt applies. A prompt qualified with an attempt
	// number above `max_attempts` (default 3) is rejected at write time — it could
	// never play.
	Attempt string `json:"attempt"`
	// Lowercase, case-sensitive detected card type for which this prompt applies.
	//
	// Any of "visa", "mastercard", "amex", "optima", "discover", "diners-club", "jcb",
	// "maestro", "enroute".
	CardType string `json:"card_type"`
	// Step error for which this prompt applies.
	//
	// Any of "timeout", "invalid-card-number", "invalid-card-type", "invalid-date",
	// "invalid-security-code", "invalid-postal-code", "invalid-bank-routing-number",
	// "invalid-bank-account-number", "input-matching-failed".
	ErrorType string `json:"error_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Attempt     respjson.Field
		CardType    respjson.Field
		ErrorType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayToolParamsPromptsExpirationDatePayToolPromptListItemResp) RawJSON() string {
	return r.JSON.raw
}
func (r *PayToolParamsPromptsExpirationDatePayToolPromptListItemResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PayToolParamsPromptsPaymentCardNumberUnionResp contains all possible properties
// and values from [string],
// [[]PayToolParamsPromptsPaymentCardNumberPayToolPromptListItemResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfPayToolPromptList]
type PayToolParamsPromptsPaymentCardNumberUnionResp struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]PayToolParamsPromptsPaymentCardNumberPayToolPromptListItemResp] instead of an
	// object.
	OfPayToolPromptList []PayToolParamsPromptsPaymentCardNumberPayToolPromptListItemResp `json:",inline"`
	JSON                struct {
		OfString            respjson.Field
		OfPayToolPromptList respjson.Field
		raw                 string
	} `json:"-"`
}

func (u PayToolParamsPromptsPaymentCardNumberUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PayToolParamsPromptsPaymentCardNumberUnionResp) AsPayToolPromptList() (v []PayToolParamsPromptsPaymentCardNumberPayToolPromptListItemResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PayToolParamsPromptsPaymentCardNumberUnionResp) RawJSON() string { return u.JSON.raw }

func (r *PayToolParamsPromptsPaymentCardNumberUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text-to-speech prompt with optional matching qualifiers, mirroring the Voice
// API PayPrompt schema.
type PayToolParamsPromptsPaymentCardNumberPayToolPromptListItemResp struct {
	// Text spoken for the payment collection step.
	Text string `json:"text" api:"required"`
	// Space-separated 1-based attempt numbers (1-9 followed by digits only, e.g.
	// `"2 3"`) for which this prompt applies. A prompt qualified with an attempt
	// number above `max_attempts` (default 3) is rejected at write time — it could
	// never play.
	Attempt string `json:"attempt"`
	// Lowercase, case-sensitive detected card type for which this prompt applies.
	//
	// Any of "visa", "mastercard", "amex", "optima", "discover", "diners-club", "jcb",
	// "maestro", "enroute".
	CardType string `json:"card_type"`
	// Step error for which this prompt applies.
	//
	// Any of "timeout", "invalid-card-number", "invalid-card-type", "invalid-date",
	// "invalid-security-code", "invalid-postal-code", "invalid-bank-routing-number",
	// "invalid-bank-account-number", "input-matching-failed".
	ErrorType string `json:"error_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Attempt     respjson.Field
		CardType    respjson.Field
		ErrorType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayToolParamsPromptsPaymentCardNumberPayToolPromptListItemResp) RawJSON() string {
	return r.JSON.raw
}
func (r *PayToolParamsPromptsPaymentCardNumberPayToolPromptListItemResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PayToolParamsPromptsPostalCodeUnionResp contains all possible properties and
// values from [string],
// [[]PayToolParamsPromptsPostalCodePayToolPromptListItemResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfPayToolPromptList]
type PayToolParamsPromptsPostalCodeUnionResp struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]PayToolParamsPromptsPostalCodePayToolPromptListItemResp] instead of an
	// object.
	OfPayToolPromptList []PayToolParamsPromptsPostalCodePayToolPromptListItemResp `json:",inline"`
	JSON                struct {
		OfString            respjson.Field
		OfPayToolPromptList respjson.Field
		raw                 string
	} `json:"-"`
}

func (u PayToolParamsPromptsPostalCodeUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PayToolParamsPromptsPostalCodeUnionResp) AsPayToolPromptList() (v []PayToolParamsPromptsPostalCodePayToolPromptListItemResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PayToolParamsPromptsPostalCodeUnionResp) RawJSON() string { return u.JSON.raw }

func (r *PayToolParamsPromptsPostalCodeUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text-to-speech prompt with optional matching qualifiers, mirroring the Voice
// API PayPrompt schema.
type PayToolParamsPromptsPostalCodePayToolPromptListItemResp struct {
	// Text spoken for the payment collection step.
	Text string `json:"text" api:"required"`
	// Space-separated 1-based attempt numbers (1-9 followed by digits only, e.g.
	// `"2 3"`) for which this prompt applies. A prompt qualified with an attempt
	// number above `max_attempts` (default 3) is rejected at write time — it could
	// never play.
	Attempt string `json:"attempt"`
	// Lowercase, case-sensitive detected card type for which this prompt applies.
	//
	// Any of "visa", "mastercard", "amex", "optima", "discover", "diners-club", "jcb",
	// "maestro", "enroute".
	CardType string `json:"card_type"`
	// Step error for which this prompt applies.
	//
	// Any of "timeout", "invalid-card-number", "invalid-card-type", "invalid-date",
	// "invalid-security-code", "invalid-postal-code", "invalid-bank-routing-number",
	// "invalid-bank-account-number", "input-matching-failed".
	ErrorType string `json:"error_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Attempt     respjson.Field
		CardType    respjson.Field
		ErrorType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayToolParamsPromptsPostalCodePayToolPromptListItemResp) RawJSON() string { return r.JSON.raw }
func (r *PayToolParamsPromptsPostalCodePayToolPromptListItemResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PayToolParamsPromptsSecurityCodeUnionResp contains all possible properties and
// values from [string],
// [[]PayToolParamsPromptsSecurityCodePayToolPromptListItemResp].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfPayToolPromptList]
type PayToolParamsPromptsSecurityCodeUnionResp struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a
	// [[]PayToolParamsPromptsSecurityCodePayToolPromptListItemResp] instead of an
	// object.
	OfPayToolPromptList []PayToolParamsPromptsSecurityCodePayToolPromptListItemResp `json:",inline"`
	JSON                struct {
		OfString            respjson.Field
		OfPayToolPromptList respjson.Field
		raw                 string
	} `json:"-"`
}

func (u PayToolParamsPromptsSecurityCodeUnionResp) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PayToolParamsPromptsSecurityCodeUnionResp) AsPayToolPromptList() (v []PayToolParamsPromptsSecurityCodePayToolPromptListItemResp) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PayToolParamsPromptsSecurityCodeUnionResp) RawJSON() string { return u.JSON.raw }

func (r *PayToolParamsPromptsSecurityCodeUnionResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A text-to-speech prompt with optional matching qualifiers, mirroring the Voice
// API PayPrompt schema.
type PayToolParamsPromptsSecurityCodePayToolPromptListItemResp struct {
	// Text spoken for the payment collection step.
	Text string `json:"text" api:"required"`
	// Space-separated 1-based attempt numbers (1-9 followed by digits only, e.g.
	// `"2 3"`) for which this prompt applies. A prompt qualified with an attempt
	// number above `max_attempts` (default 3) is rejected at write time — it could
	// never play.
	Attempt string `json:"attempt"`
	// Lowercase, case-sensitive detected card type for which this prompt applies.
	//
	// Any of "visa", "mastercard", "amex", "optima", "discover", "diners-club", "jcb",
	// "maestro", "enroute".
	CardType string `json:"card_type"`
	// Step error for which this prompt applies.
	//
	// Any of "timeout", "invalid-card-number", "invalid-card-type", "invalid-date",
	// "invalid-security-code", "invalid-postal-code", "invalid-bank-routing-number",
	// "invalid-bank-account-number", "input-matching-failed".
	ErrorType string `json:"error_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Text        respjson.Field
		Attempt     respjson.Field
		CardType    respjson.Field
		ErrorType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PayToolParamsPromptsSecurityCodePayToolPromptListItemResp) RawJSON() string {
	return r.JSON.raw
}
func (r *PayToolParamsPromptsSecurityCodePayToolPromptListItemResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property ConnectorName is required.
type PayToolParams struct {
	// The name of the pay connector configured in the Telnyx API. Must reference an
	// existing pay connector for this organization.
	ConnectorName string `json:"connector_name" api:"required"`
	// Description of the pay tool shown to the assistant so it knows when to call the
	// tool. Not sent to the payment connector — the per-call `description` argument
	// the assistant passes is what reaches the processor.
	Description param.Opt[string] `json:"description,omitzero"`
	// Default currency for payments processed by this tool.
	Currency param.Opt[string] `json:"currency,omitzero"`
	// Time in milliseconds to wait between consecutive DTMF digits. Omitted means the
	// platform default (5000) applies.
	InterDigitTimeoutMillis param.Opt[int64] `json:"inter_digit_timeout_millis,omitzero"`
	// Maximum number of attempts for each payment collection step. Omitted means the
	// platform default (3) applies.
	MaxAttempts param.Opt[int64] `json:"max_attempts,omitzero"`
	// Minimum number of digits the caller must enter when collecting a credit-card
	// billing postal code. Applies only when `postal_code` prompts the caller. Omitted
	// means the platform default (1) applies. The maximum number of digits collected
	// is the greater of 10 and this value. This does not validate the length of a
	// postal code supplied directly in `postal_code`.
	MinPostalCodeLength param.Opt[int64] `json:"min_postal_code_length,omitzero"`
	// Default payment method for payments processed by this tool.
	PaymentMethod param.Opt[string] `json:"payment_method,omitzero"`
	// Time in milliseconds to wait for DTMF input for each collection step. Omitted
	// means the platform default (5000) applies. Increase this for slower entry, e.g.
	// a 16-digit card number.
	TimeoutMillis param.Opt[int64] `json:"timeout_millis,omitzero"`
	// Custom key/value parameters forwarded to the pay connector as `parameters` in
	// the processor request. String values may use `{{dynamic_variable}}` templating,
	// resolved per call.
	ConnectorParameters map[string]any `json:"connector_parameters,omitzero"`
	// Controls billing postal code collection for credit-card payments. Set to `false`
	// to skip the postal code prompt without supplying a postal code to the payment
	// connector (non-US cards, e.g. Irish Eircodes). Set to `true` to prompt the
	// caller. Supply a non-empty string, such as `"90210"`, to use that postal code
	// without prompting. This option does not affect ACH debit payments or payments
	// using an existing token.
	PostalCode PayToolParamsPostalCodeUnion `json:"postal_code,omitzero"`
	// Custom text-to-speech prompts keyed by payment collection step. The AI assistant
	// merges its own card-number prompt into this map, matching on the same step keys.
	Prompts PayToolParamsPrompts `json:"prompts,omitzero"`
	// Restricts accepted card numbers to the listed card types. When the caller enters
	// a card number that does not match one of the listed types, Pay treats the input
	// as invalid and re-prompts for the card number.
	//
	// Any of "visa", "mastercard", "amex", "maestro", "discover", "optima", "jcb",
	// "diners-club", "enroute".
	ValidCardTypes []string `json:"valid_card_types,omitzero"`
	paramObj
}

func (r PayToolParams) MarshalJSON() (data []byte, err error) {
	type shadow PayToolParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PayToolParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PayToolParamsPostalCodeUnion struct {
	OfString param.Opt[string] `json:",omitzero,inline"`
	OfBool   param.Opt[bool]   `json:",omitzero,inline"`
	paramUnion
}

func (u PayToolParamsPostalCodeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfBool)
}
func (u *PayToolParamsPostalCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PayToolParamsPostalCodeUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfBool) {
		return &u.OfBool.Value
	}
	return nil
}

// Custom text-to-speech prompts keyed by payment collection step. The AI assistant
// merges its own card-number prompt into this map, matching on the same step keys.
type PayToolParamsPrompts struct {
	// A default prompt string or an ordered list of qualified prompts.
	BankAccountNumber PayToolParamsPromptsBankAccountNumberUnion `json:"bank-account-number,omitzero"`
	// A default prompt string or an ordered list of qualified prompts.
	BankRoutingNumber PayToolParamsPromptsBankRoutingNumberUnion `json:"bank-routing-number,omitzero"`
	// A default prompt string or an ordered list of qualified prompts.
	ExpirationDate PayToolParamsPromptsExpirationDateUnion `json:"expiration-date,omitzero"`
	// A default prompt string or an ordered list of qualified prompts.
	PaymentCardNumber PayToolParamsPromptsPaymentCardNumberUnion `json:"payment-card-number,omitzero"`
	// A default prompt string or an ordered list of qualified prompts.
	PostalCode PayToolParamsPromptsPostalCodeUnion `json:"postal-code,omitzero"`
	// A default prompt string or an ordered list of qualified prompts.
	SecurityCode PayToolParamsPromptsSecurityCodeUnion `json:"security-code,omitzero"`
	paramObj
}

func (r PayToolParamsPrompts) MarshalJSON() (data []byte, err error) {
	type shadow PayToolParamsPrompts
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PayToolParamsPrompts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PayToolParamsPromptsBankAccountNumberUnion struct {
	OfString            param.Opt[string]                                            `json:",omitzero,inline"`
	OfPayToolPromptList []PayToolParamsPromptsBankAccountNumberPayToolPromptListItem `json:",omitzero,inline"`
	paramUnion
}

func (u PayToolParamsPromptsBankAccountNumberUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfPayToolPromptList)
}
func (u *PayToolParamsPromptsBankAccountNumberUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PayToolParamsPromptsBankAccountNumberUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfPayToolPromptList) {
		return &u.OfPayToolPromptList
	}
	return nil
}

// A text-to-speech prompt with optional matching qualifiers, mirroring the Voice
// API PayPrompt schema.
//
// The property Text is required.
type PayToolParamsPromptsBankAccountNumberPayToolPromptListItem struct {
	// Text spoken for the payment collection step.
	Text string `json:"text" api:"required"`
	// Space-separated 1-based attempt numbers (1-9 followed by digits only, e.g.
	// `"2 3"`) for which this prompt applies. A prompt qualified with an attempt
	// number above `max_attempts` (default 3) is rejected at write time — it could
	// never play.
	Attempt param.Opt[string] `json:"attempt,omitzero"`
	// Lowercase, case-sensitive detected card type for which this prompt applies.
	//
	// Any of "visa", "mastercard", "amex", "optima", "discover", "diners-club", "jcb",
	// "maestro", "enroute".
	CardType string `json:"card_type,omitzero"`
	// Step error for which this prompt applies.
	//
	// Any of "timeout", "invalid-card-number", "invalid-card-type", "invalid-date",
	// "invalid-security-code", "invalid-postal-code", "invalid-bank-routing-number",
	// "invalid-bank-account-number", "input-matching-failed".
	ErrorType string `json:"error_type,omitzero"`
	paramObj
}

func (r PayToolParamsPromptsBankAccountNumberPayToolPromptListItem) MarshalJSON() (data []byte, err error) {
	type shadow PayToolParamsPromptsBankAccountNumberPayToolPromptListItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PayToolParamsPromptsBankAccountNumberPayToolPromptListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PayToolParamsPromptsBankAccountNumberPayToolPromptListItem](
		"card_type", "visa", "mastercard", "amex", "optima", "discover", "diners-club", "jcb", "maestro", "enroute",
	)
	apijson.RegisterFieldValidator[PayToolParamsPromptsBankAccountNumberPayToolPromptListItem](
		"error_type", "timeout", "invalid-card-number", "invalid-card-type", "invalid-date", "invalid-security-code", "invalid-postal-code", "invalid-bank-routing-number", "invalid-bank-account-number", "input-matching-failed",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PayToolParamsPromptsBankRoutingNumberUnion struct {
	OfString            param.Opt[string]                                            `json:",omitzero,inline"`
	OfPayToolPromptList []PayToolParamsPromptsBankRoutingNumberPayToolPromptListItem `json:",omitzero,inline"`
	paramUnion
}

func (u PayToolParamsPromptsBankRoutingNumberUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfPayToolPromptList)
}
func (u *PayToolParamsPromptsBankRoutingNumberUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PayToolParamsPromptsBankRoutingNumberUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfPayToolPromptList) {
		return &u.OfPayToolPromptList
	}
	return nil
}

// A text-to-speech prompt with optional matching qualifiers, mirroring the Voice
// API PayPrompt schema.
//
// The property Text is required.
type PayToolParamsPromptsBankRoutingNumberPayToolPromptListItem struct {
	// Text spoken for the payment collection step.
	Text string `json:"text" api:"required"`
	// Space-separated 1-based attempt numbers (1-9 followed by digits only, e.g.
	// `"2 3"`) for which this prompt applies. A prompt qualified with an attempt
	// number above `max_attempts` (default 3) is rejected at write time — it could
	// never play.
	Attempt param.Opt[string] `json:"attempt,omitzero"`
	// Lowercase, case-sensitive detected card type for which this prompt applies.
	//
	// Any of "visa", "mastercard", "amex", "optima", "discover", "diners-club", "jcb",
	// "maestro", "enroute".
	CardType string `json:"card_type,omitzero"`
	// Step error for which this prompt applies.
	//
	// Any of "timeout", "invalid-card-number", "invalid-card-type", "invalid-date",
	// "invalid-security-code", "invalid-postal-code", "invalid-bank-routing-number",
	// "invalid-bank-account-number", "input-matching-failed".
	ErrorType string `json:"error_type,omitzero"`
	paramObj
}

func (r PayToolParamsPromptsBankRoutingNumberPayToolPromptListItem) MarshalJSON() (data []byte, err error) {
	type shadow PayToolParamsPromptsBankRoutingNumberPayToolPromptListItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PayToolParamsPromptsBankRoutingNumberPayToolPromptListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PayToolParamsPromptsBankRoutingNumberPayToolPromptListItem](
		"card_type", "visa", "mastercard", "amex", "optima", "discover", "diners-club", "jcb", "maestro", "enroute",
	)
	apijson.RegisterFieldValidator[PayToolParamsPromptsBankRoutingNumberPayToolPromptListItem](
		"error_type", "timeout", "invalid-card-number", "invalid-card-type", "invalid-date", "invalid-security-code", "invalid-postal-code", "invalid-bank-routing-number", "invalid-bank-account-number", "input-matching-failed",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PayToolParamsPromptsExpirationDateUnion struct {
	OfString            param.Opt[string]                                         `json:",omitzero,inline"`
	OfPayToolPromptList []PayToolParamsPromptsExpirationDatePayToolPromptListItem `json:",omitzero,inline"`
	paramUnion
}

func (u PayToolParamsPromptsExpirationDateUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfPayToolPromptList)
}
func (u *PayToolParamsPromptsExpirationDateUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PayToolParamsPromptsExpirationDateUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfPayToolPromptList) {
		return &u.OfPayToolPromptList
	}
	return nil
}

// A text-to-speech prompt with optional matching qualifiers, mirroring the Voice
// API PayPrompt schema.
//
// The property Text is required.
type PayToolParamsPromptsExpirationDatePayToolPromptListItem struct {
	// Text spoken for the payment collection step.
	Text string `json:"text" api:"required"`
	// Space-separated 1-based attempt numbers (1-9 followed by digits only, e.g.
	// `"2 3"`) for which this prompt applies. A prompt qualified with an attempt
	// number above `max_attempts` (default 3) is rejected at write time — it could
	// never play.
	Attempt param.Opt[string] `json:"attempt,omitzero"`
	// Lowercase, case-sensitive detected card type for which this prompt applies.
	//
	// Any of "visa", "mastercard", "amex", "optima", "discover", "diners-club", "jcb",
	// "maestro", "enroute".
	CardType string `json:"card_type,omitzero"`
	// Step error for which this prompt applies.
	//
	// Any of "timeout", "invalid-card-number", "invalid-card-type", "invalid-date",
	// "invalid-security-code", "invalid-postal-code", "invalid-bank-routing-number",
	// "invalid-bank-account-number", "input-matching-failed".
	ErrorType string `json:"error_type,omitzero"`
	paramObj
}

func (r PayToolParamsPromptsExpirationDatePayToolPromptListItem) MarshalJSON() (data []byte, err error) {
	type shadow PayToolParamsPromptsExpirationDatePayToolPromptListItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PayToolParamsPromptsExpirationDatePayToolPromptListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PayToolParamsPromptsExpirationDatePayToolPromptListItem](
		"card_type", "visa", "mastercard", "amex", "optima", "discover", "diners-club", "jcb", "maestro", "enroute",
	)
	apijson.RegisterFieldValidator[PayToolParamsPromptsExpirationDatePayToolPromptListItem](
		"error_type", "timeout", "invalid-card-number", "invalid-card-type", "invalid-date", "invalid-security-code", "invalid-postal-code", "invalid-bank-routing-number", "invalid-bank-account-number", "input-matching-failed",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PayToolParamsPromptsPaymentCardNumberUnion struct {
	OfString            param.Opt[string]                                            `json:",omitzero,inline"`
	OfPayToolPromptList []PayToolParamsPromptsPaymentCardNumberPayToolPromptListItem `json:",omitzero,inline"`
	paramUnion
}

func (u PayToolParamsPromptsPaymentCardNumberUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfPayToolPromptList)
}
func (u *PayToolParamsPromptsPaymentCardNumberUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PayToolParamsPromptsPaymentCardNumberUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfPayToolPromptList) {
		return &u.OfPayToolPromptList
	}
	return nil
}

// A text-to-speech prompt with optional matching qualifiers, mirroring the Voice
// API PayPrompt schema.
//
// The property Text is required.
type PayToolParamsPromptsPaymentCardNumberPayToolPromptListItem struct {
	// Text spoken for the payment collection step.
	Text string `json:"text" api:"required"`
	// Space-separated 1-based attempt numbers (1-9 followed by digits only, e.g.
	// `"2 3"`) for which this prompt applies. A prompt qualified with an attempt
	// number above `max_attempts` (default 3) is rejected at write time — it could
	// never play.
	Attempt param.Opt[string] `json:"attempt,omitzero"`
	// Lowercase, case-sensitive detected card type for which this prompt applies.
	//
	// Any of "visa", "mastercard", "amex", "optima", "discover", "diners-club", "jcb",
	// "maestro", "enroute".
	CardType string `json:"card_type,omitzero"`
	// Step error for which this prompt applies.
	//
	// Any of "timeout", "invalid-card-number", "invalid-card-type", "invalid-date",
	// "invalid-security-code", "invalid-postal-code", "invalid-bank-routing-number",
	// "invalid-bank-account-number", "input-matching-failed".
	ErrorType string `json:"error_type,omitzero"`
	paramObj
}

func (r PayToolParamsPromptsPaymentCardNumberPayToolPromptListItem) MarshalJSON() (data []byte, err error) {
	type shadow PayToolParamsPromptsPaymentCardNumberPayToolPromptListItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PayToolParamsPromptsPaymentCardNumberPayToolPromptListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PayToolParamsPromptsPaymentCardNumberPayToolPromptListItem](
		"card_type", "visa", "mastercard", "amex", "optima", "discover", "diners-club", "jcb", "maestro", "enroute",
	)
	apijson.RegisterFieldValidator[PayToolParamsPromptsPaymentCardNumberPayToolPromptListItem](
		"error_type", "timeout", "invalid-card-number", "invalid-card-type", "invalid-date", "invalid-security-code", "invalid-postal-code", "invalid-bank-routing-number", "invalid-bank-account-number", "input-matching-failed",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PayToolParamsPromptsPostalCodeUnion struct {
	OfString            param.Opt[string]                                     `json:",omitzero,inline"`
	OfPayToolPromptList []PayToolParamsPromptsPostalCodePayToolPromptListItem `json:",omitzero,inline"`
	paramUnion
}

func (u PayToolParamsPromptsPostalCodeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfPayToolPromptList)
}
func (u *PayToolParamsPromptsPostalCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PayToolParamsPromptsPostalCodeUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfPayToolPromptList) {
		return &u.OfPayToolPromptList
	}
	return nil
}

// A text-to-speech prompt with optional matching qualifiers, mirroring the Voice
// API PayPrompt schema.
//
// The property Text is required.
type PayToolParamsPromptsPostalCodePayToolPromptListItem struct {
	// Text spoken for the payment collection step.
	Text string `json:"text" api:"required"`
	// Space-separated 1-based attempt numbers (1-9 followed by digits only, e.g.
	// `"2 3"`) for which this prompt applies. A prompt qualified with an attempt
	// number above `max_attempts` (default 3) is rejected at write time — it could
	// never play.
	Attempt param.Opt[string] `json:"attempt,omitzero"`
	// Lowercase, case-sensitive detected card type for which this prompt applies.
	//
	// Any of "visa", "mastercard", "amex", "optima", "discover", "diners-club", "jcb",
	// "maestro", "enroute".
	CardType string `json:"card_type,omitzero"`
	// Step error for which this prompt applies.
	//
	// Any of "timeout", "invalid-card-number", "invalid-card-type", "invalid-date",
	// "invalid-security-code", "invalid-postal-code", "invalid-bank-routing-number",
	// "invalid-bank-account-number", "input-matching-failed".
	ErrorType string `json:"error_type,omitzero"`
	paramObj
}

func (r PayToolParamsPromptsPostalCodePayToolPromptListItem) MarshalJSON() (data []byte, err error) {
	type shadow PayToolParamsPromptsPostalCodePayToolPromptListItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PayToolParamsPromptsPostalCodePayToolPromptListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PayToolParamsPromptsPostalCodePayToolPromptListItem](
		"card_type", "visa", "mastercard", "amex", "optima", "discover", "diners-club", "jcb", "maestro", "enroute",
	)
	apijson.RegisterFieldValidator[PayToolParamsPromptsPostalCodePayToolPromptListItem](
		"error_type", "timeout", "invalid-card-number", "invalid-card-type", "invalid-date", "invalid-security-code", "invalid-postal-code", "invalid-bank-routing-number", "invalid-bank-account-number", "input-matching-failed",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PayToolParamsPromptsSecurityCodeUnion struct {
	OfString            param.Opt[string]                                       `json:",omitzero,inline"`
	OfPayToolPromptList []PayToolParamsPromptsSecurityCodePayToolPromptListItem `json:",omitzero,inline"`
	paramUnion
}

func (u PayToolParamsPromptsSecurityCodeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfString, u.OfPayToolPromptList)
}
func (u *PayToolParamsPromptsSecurityCodeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PayToolParamsPromptsSecurityCodeUnion) asAny() any {
	if !param.IsOmitted(u.OfString) {
		return &u.OfString.Value
	} else if !param.IsOmitted(u.OfPayToolPromptList) {
		return &u.OfPayToolPromptList
	}
	return nil
}

// A text-to-speech prompt with optional matching qualifiers, mirroring the Voice
// API PayPrompt schema.
//
// The property Text is required.
type PayToolParamsPromptsSecurityCodePayToolPromptListItem struct {
	// Text spoken for the payment collection step.
	Text string `json:"text" api:"required"`
	// Space-separated 1-based attempt numbers (1-9 followed by digits only, e.g.
	// `"2 3"`) for which this prompt applies. A prompt qualified with an attempt
	// number above `max_attempts` (default 3) is rejected at write time — it could
	// never play.
	Attempt param.Opt[string] `json:"attempt,omitzero"`
	// Lowercase, case-sensitive detected card type for which this prompt applies.
	//
	// Any of "visa", "mastercard", "amex", "optima", "discover", "diners-club", "jcb",
	// "maestro", "enroute".
	CardType string `json:"card_type,omitzero"`
	// Step error for which this prompt applies.
	//
	// Any of "timeout", "invalid-card-number", "invalid-card-type", "invalid-date",
	// "invalid-security-code", "invalid-postal-code", "invalid-bank-routing-number",
	// "invalid-bank-account-number", "input-matching-failed".
	ErrorType string `json:"error_type,omitzero"`
	paramObj
}

func (r PayToolParamsPromptsSecurityCodePayToolPromptListItem) MarshalJSON() (data []byte, err error) {
	type shadow PayToolParamsPromptsSecurityCodePayToolPromptListItem
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PayToolParamsPromptsSecurityCodePayToolPromptListItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PayToolParamsPromptsSecurityCodePayToolPromptListItem](
		"card_type", "visa", "mastercard", "amex", "optima", "discover", "diners-club", "jcb", "maestro", "enroute",
	)
	apijson.RegisterFieldValidator[PayToolParamsPromptsSecurityCodePayToolPromptListItem](
		"error_type", "timeout", "invalid-card-number", "invalid-card-type", "invalid-date", "invalid-security-code", "invalid-postal-code", "invalid-bank-routing-number", "invalid-bank-account-number", "input-matching-failed",
	)
}

type SharedToolResponse struct {
	ID             string         `json:"id" api:"required"`
	ToolDefinition map[string]any `json:"tool_definition" api:"required"`
	Type           string         `json:"type" api:"required"`
	CreatedAt      string         `json:"created_at"`
	DisplayName    string         `json:"display_name"`
	TimeoutMs      int64          `json:"timeout_ms"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		ToolDefinition respjson.Field
		Type           respjson.Field
		CreatedAt      respjson.Field
		DisplayName    respjson.Field
		TimeoutMs      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SharedToolResponse) RawJSON() string { return r.JSON.raw }
func (r *SharedToolResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for an update_dynamic_variables tool.
type UpdateDynamicVariablesToolParamsResp struct {
	// Description of the tool passed to the assistant, guiding when to call it and
	// which variables to update.
	Description string `json:"description" api:"required"`
	// The function name surfaced to the LLM. Must match the OpenAI function-name
	// pattern `^[a-zA-Z0-9_-]+$` and be unique across the assistant's function,
	// webhook, and client_side tools.
	Name string `json:"name" api:"required"`
	// The dynamic variables the assistant is allowed to write. At least one is
	// required.
	UpdatableVariables []UpdateDynamicVariablesToolParamsUpdatableVariableResp `json:"updatable_variables" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description        respjson.Field
		Name               respjson.Field
		UpdatableVariables respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UpdateDynamicVariablesToolParamsResp) RawJSON() string { return r.JSON.raw }
func (r *UpdateDynamicVariablesToolParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this UpdateDynamicVariablesToolParamsResp to a
// UpdateDynamicVariablesToolParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// UpdateDynamicVariablesToolParams.Overrides()
func (r UpdateDynamicVariablesToolParamsResp) ToParam() UpdateDynamicVariablesToolParams {
	return param.Override[UpdateDynamicVariablesToolParams](json.RawMessage(r.RawJSON()))
}

type UpdateDynamicVariablesToolParamsUpdatableVariableResp struct {
	// The dynamic-variable key to update. Must match `^[a-zA-Z0-9._-]+$` and may not
	// start with the reserved `telnyx_` prefix (reserved for system variables). The
	// `pattern` encodes both rules via a negative lookahead.
	Name string `json:"name" api:"required"`
	// Optional description of the variable, guiding the assistant on what value to
	// capture.
	Description string `json:"description"`
	// Optional hint for the variable's value type (e.g. `string`).
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Description respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UpdateDynamicVariablesToolParamsUpdatableVariableResp) RawJSON() string { return r.JSON.raw }
func (r *UpdateDynamicVariablesToolParamsUpdatableVariableResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for an update_dynamic_variables tool.
//
// The properties Description, Name, UpdatableVariables are required.
type UpdateDynamicVariablesToolParams struct {
	// Description of the tool passed to the assistant, guiding when to call it and
	// which variables to update.
	Description string `json:"description" api:"required"`
	// The function name surfaced to the LLM. Must match the OpenAI function-name
	// pattern `^[a-zA-Z0-9_-]+$` and be unique across the assistant's function,
	// webhook, and client_side tools.
	Name string `json:"name" api:"required"`
	// The dynamic variables the assistant is allowed to write. At least one is
	// required.
	UpdatableVariables []UpdateDynamicVariablesToolParamsUpdatableVariable `json:"updatable_variables,omitzero" api:"required"`
	paramObj
}

func (r UpdateDynamicVariablesToolParams) MarshalJSON() (data []byte, err error) {
	type shadow UpdateDynamicVariablesToolParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateDynamicVariablesToolParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type UpdateDynamicVariablesToolParamsUpdatableVariable struct {
	// The dynamic-variable key to update. Must match `^[a-zA-Z0-9._-]+$` and may not
	// start with the reserved `telnyx_` prefix (reserved for system variables). The
	// `pattern` encodes both rules via a negative lookahead.
	Name string `json:"name" api:"required"`
	// Optional description of the variable, guiding the assistant on what value to
	// capture.
	Description param.Opt[string] `json:"description,omitzero"`
	// Optional hint for the variable's value type (e.g. `string`).
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r UpdateDynamicVariablesToolParamsUpdatableVariable) MarshalJSON() (data []byte, err error) {
	type shadow UpdateDynamicVariablesToolParamsUpdatableVariable
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpdateDynamicVariablesToolParamsUpdatableVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VisionToolParamsResp struct {
	// Frames received longer ago than this many seconds are not shown to the
	// assistant, so it never describes a frozen or disconnected camera as live.
	MaxFrameAgeSecs int64 `json:"max_frame_age_secs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxFrameAgeSecs respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VisionToolParamsResp) RawJSON() string { return r.JSON.raw }
func (r *VisionToolParamsResp) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this VisionToolParamsResp to a VisionToolParams.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VisionToolParams.Overrides()
func (r VisionToolParamsResp) ToParam() VisionToolParams {
	return param.Override[VisionToolParams](json.RawMessage(r.RawJSON()))
}

type VisionToolParams struct {
	// Frames received longer ago than this many seconds are not shown to the
	// assistant, so it never describes a frozen or disconnected camera as live.
	MaxFrameAgeSecs param.Opt[int64] `json:"max_frame_age_secs,omitzero"`
	paramObj
}

func (r VisionToolParams) MarshalJSON() (data []byte, err error) {
	type shadow VisionToolParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VisionToolParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIToolNewParams struct {
	DisplayName    string            `json:"display_name" api:"required"`
	Type           string            `json:"type" api:"required"`
	TimeoutMs      param.Opt[int64]  `json:"timeout_ms,omitzero"`
	IdempotencyKey param.Opt[string] `header:"Idempotency-Key,omitzero" json:"-"`
	ClientSideTool map[string]any    `json:"client_side_tool,omitzero"`
	Function       map[string]any    `json:"function,omitzero"`
	Handoff        map[string]any    `json:"handoff,omitzero"`
	Invite         map[string]any    `json:"invite,omitzero"`
	Pay            PayToolParams     `json:"pay,omitzero"`
	Retrieval      map[string]any    `json:"retrieval,omitzero"`
	// Configuration for an update_dynamic_variables tool.
	UpdateDynamicVariables UpdateDynamicVariablesToolParams `json:"update_dynamic_variables,omitzero"`
	Vision                 VisionToolParams                 `json:"vision,omitzero"`
	Webhook                map[string]any                   `json:"webhook,omitzero"`
	paramObj
}

func (r AIToolNewParams) MarshalJSON() (data []byte, err error) {
	type shadow AIToolNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIToolNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIToolUpdateParams struct {
	DisplayName    param.Opt[string] `json:"display_name,omitzero"`
	TimeoutMs      param.Opt[int64]  `json:"timeout_ms,omitzero"`
	Type           param.Opt[string] `json:"type,omitzero"`
	ClientSideTool map[string]any    `json:"client_side_tool,omitzero"`
	Function       map[string]any    `json:"function,omitzero"`
	Handoff        map[string]any    `json:"handoff,omitzero"`
	Invite         map[string]any    `json:"invite,omitzero"`
	Pay            PayToolParams     `json:"pay,omitzero"`
	Retrieval      map[string]any    `json:"retrieval,omitzero"`
	// Configuration for an update_dynamic_variables tool.
	UpdateDynamicVariables UpdateDynamicVariablesToolParams `json:"update_dynamic_variables,omitzero"`
	Vision                 VisionToolParams                 `json:"vision,omitzero"`
	Webhook                map[string]any                   `json:"webhook,omitzero"`
	paramObj
}

func (r AIToolUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow AIToolUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIToolUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIToolListParams struct {
	// Filter results by filter name.
	FilterName param.Opt[string] `query:"filter[name],omitzero" json:"-"`
	// Filter results by filter type.
	FilterType param.Opt[string] `query:"filter[type],omitzero" json:"-"`
	// Page number to retrieve (1-based).
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of items to return per page.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AIToolListParams]'s query parameters as `url.Values`.
func (r AIToolListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
