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
	"time"

	"github.com/team-telnyx/telnyx-go/v4/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v4/internal/apiquery"
	"github.com/team-telnyx/telnyx-go/v4/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v4/option"
	"github.com/team-telnyx/telnyx-go/v4/packages/pagination"
	"github.com/team-telnyx/telnyx-go/v4/packages/param"
	"github.com/team-telnyx/telnyx-go/v4/packages/respjson"
)

// Manage pronunciation dictionaries for text-to-speech synthesis. Dictionaries
// contain alias items (text replacement) and phoneme items (IPA pronunciation
// notation) that control how specific words are spoken.
//
// PronunciationDictService contains methods and other services that help with
// interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPronunciationDictService] method instead.
type PronunciationDictService struct {
	Options []option.RequestOption
}

// NewPronunciationDictService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPronunciationDictService(opts ...option.RequestOption) (r PronunciationDictService) {
	r = PronunciationDictService{}
	r.Options = opts
	return
}

// Create a new pronunciation dictionary for the authenticated organization. Each
// dictionary contains a list of items that control how specific words are spoken.
// Items can be alias type (text replacement) or phoneme type (IPA pronunciation
// notation).
//
// As an alternative to providing items directly as JSON, you can upload a
// dictionary file (PLS/XML or plain text format, max 1MB) using
// multipart/form-data. PLS files use the standard W3C Pronunciation Lexicon
// Specification XML format. Text files use a line-based format: `word=alias` for
// aliases, `word:/phoneme/` for IPA phonemes.
//
// Limits:
//
// - Maximum 50 dictionaries per organization
// - Maximum 100 items per dictionary
// - Text: max 200 characters
// - Alias/phoneme value: max 500 characters
// - File upload: max 1MB (1,048,576 bytes)
func (r *PronunciationDictService) New(ctx context.Context, body PronunciationDictNewParams, opts ...option.RequestOption) (res *PronunciationDictNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pronunciation_dicts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieve a single pronunciation dictionary by ID.
func (r *PronunciationDictService) Get(ctx context.Context, id string, opts ...option.RequestOption) (res *PronunciationDictGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("pronunciation_dicts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Update the name and/or items of an existing pronunciation dictionary. Uses
// optimistic locking — if the dictionary was modified concurrently, the request
// returns 409 Conflict.
func (r *PronunciationDictService) Update(ctx context.Context, id string, body PronunciationDictUpdateParams, opts ...option.RequestOption) (res *PronunciationDictUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if id == "" {
		err = errors.New("missing required id parameter")
		return nil, err
	}
	path := fmt.Sprintf("pronunciation_dicts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// List all pronunciation dictionaries for the authenticated organization. Results
// are paginated using offset-based pagination.
func (r *PronunciationDictService) List(ctx context.Context, query PronunciationDictListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[PronunciationDictData], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "pronunciation_dicts"
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

// List all pronunciation dictionaries for the authenticated organization. Results
// are paginated using offset-based pagination.
func (r *PronunciationDictService) ListAutoPaging(ctx context.Context, query PronunciationDictListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[PronunciationDictData] {
	return pagination.NewDefaultFlatPaginationAutoPager(r.List(ctx, query, opts...))
}

// Permanently delete a pronunciation dictionary.
func (r *PronunciationDictService) Delete(ctx context.Context, id string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if id == "" {
		err = errors.New("missing required id parameter")
		return err
	}
	path := fmt.Sprintf("pronunciation_dicts/%s", id)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// An alias pronunciation item. When the `text` value is found in input, it is
// replaced with the `alias` before speech synthesis.
type PronunciationDictAliasItem struct {
	// The replacement text that will be spoken instead.
	Alias string `json:"alias" api:"required"`
	// The text to match in the input. Case-insensitive matching is used during
	// synthesis.
	Text string `json:"text" api:"required"`
	// The item type.
	//
	// Any of "alias".
	Type PronunciationDictAliasItemType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alias       respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PronunciationDictAliasItem) RawJSON() string { return r.JSON.raw }
func (r *PronunciationDictAliasItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PronunciationDictAliasItem to a
// PronunciationDictAliasItemParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PronunciationDictAliasItemParam.Overrides()
func (r PronunciationDictAliasItem) ToParam() PronunciationDictAliasItemParam {
	return param.Override[PronunciationDictAliasItemParam](json.RawMessage(r.RawJSON()))
}

// The item type.
type PronunciationDictAliasItemType string

const (
	PronunciationDictAliasItemTypeAlias PronunciationDictAliasItemType = "alias"
)

// An alias pronunciation item. When the `text` value is found in input, it is
// replaced with the `alias` before speech synthesis.
//
// The properties Alias, Text, Type are required.
type PronunciationDictAliasItemParam struct {
	// The replacement text that will be spoken instead.
	Alias string `json:"alias" api:"required"`
	// The text to match in the input. Case-insensitive matching is used during
	// synthesis.
	Text string `json:"text" api:"required"`
	// The item type.
	//
	// Any of "alias".
	Type PronunciationDictAliasItemType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r PronunciationDictAliasItemParam) MarshalJSON() (data []byte, err error) {
	type shadow PronunciationDictAliasItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PronunciationDictAliasItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A pronunciation dictionary record.
type PronunciationDictData struct {
	// Unique identifier for the pronunciation dictionary.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 timestamp with millisecond precision.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// List of pronunciation items (alias or phoneme type).
	Items []PronunciationDictDataItemUnion `json:"items"`
	// Human-readable name for the dictionary. Must be unique within the organization.
	Name string `json:"name"`
	// Identifies the resource type.
	//
	// Any of "pronunciation_dict".
	RecordType PronunciationDictDataRecordType `json:"record_type"`
	// ISO 8601 timestamp with millisecond precision.
	UpdatedAt time.Time `json:"updated_at" format:"date-time"`
	// Auto-incrementing version number. Increases by 1 on each update. Used for
	// optimistic concurrency control and cache invalidation.
	Version int64 `json:"version"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Items       respjson.Field
		Name        respjson.Field
		RecordType  respjson.Field
		UpdatedAt   respjson.Field
		Version     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PronunciationDictData) RawJSON() string { return r.JSON.raw }
func (r *PronunciationDictData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PronunciationDictDataItemUnion contains all possible properties and values from
// [PronunciationDictAliasItem], [PronunciationDictPhonemeItem].
//
// Use the [PronunciationDictDataItemUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PronunciationDictDataItemUnion struct {
	// This field is from variant [PronunciationDictAliasItem].
	Alias string `json:"alias"`
	Text  string `json:"text"`
	// Any of "alias", "phoneme".
	Type string `json:"type"`
	// This field is from variant [PronunciationDictPhonemeItem].
	Alphabet PronunciationDictPhonemeItemAlphabet `json:"alphabet"`
	// This field is from variant [PronunciationDictPhonemeItem].
	Phoneme string `json:"phoneme"`
	JSON    struct {
		Alias    respjson.Field
		Text     respjson.Field
		Type     respjson.Field
		Alphabet respjson.Field
		Phoneme  respjson.Field
		raw      string
	} `json:"-"`
}

// anyPronunciationDictDataItem is implemented by each variant of
// [PronunciationDictDataItemUnion] to add type safety for the return type of
// [PronunciationDictDataItemUnion.AsAny]
type anyPronunciationDictDataItem interface {
	implPronunciationDictDataItemUnion()
}

func (PronunciationDictAliasItem) implPronunciationDictDataItemUnion()   {}
func (PronunciationDictPhonemeItem) implPronunciationDictDataItemUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PronunciationDictDataItemUnion.AsAny().(type) {
//	case telnyx.PronunciationDictAliasItem:
//	case telnyx.PronunciationDictPhonemeItem:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PronunciationDictDataItemUnion) AsAny() anyPronunciationDictDataItem {
	switch u.Type {
	case "alias":
		return u.AsAlias()
	case "phoneme":
		return u.AsPhoneme()
	}
	return nil
}

func (u PronunciationDictDataItemUnion) AsAlias() (v PronunciationDictAliasItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PronunciationDictDataItemUnion) AsPhoneme() (v PronunciationDictPhonemeItem) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PronunciationDictDataItemUnion) RawJSON() string { return u.JSON.raw }

func (r *PronunciationDictDataItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the resource type.
type PronunciationDictDataRecordType string

const (
	PronunciationDictDataRecordTypePronunciationDict PronunciationDictDataRecordType = "pronunciation_dict"
)

// A phoneme pronunciation item. When the `text` value is found in input, it is
// pronounced using the specified IPA phoneme notation.
type PronunciationDictPhonemeItem struct {
	// The phonetic alphabet used for the phoneme notation.
	//
	// Any of "ipa".
	Alphabet PronunciationDictPhonemeItemAlphabet `json:"alphabet" api:"required"`
	// The phoneme notation representing the desired pronunciation.
	Phoneme string `json:"phoneme" api:"required"`
	// The text to match in the input. Case-insensitive matching is used during
	// synthesis.
	Text string `json:"text" api:"required"`
	// The item type.
	//
	// Any of "phoneme".
	Type PronunciationDictPhonemeItemType `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Alphabet    respjson.Field
		Phoneme     respjson.Field
		Text        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PronunciationDictPhonemeItem) RawJSON() string { return r.JSON.raw }
func (r *PronunciationDictPhonemeItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PronunciationDictPhonemeItem to a
// PronunciationDictPhonemeItemParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PronunciationDictPhonemeItemParam.Overrides()
func (r PronunciationDictPhonemeItem) ToParam() PronunciationDictPhonemeItemParam {
	return param.Override[PronunciationDictPhonemeItemParam](json.RawMessage(r.RawJSON()))
}

// The phonetic alphabet used for the phoneme notation.
type PronunciationDictPhonemeItemAlphabet string

const (
	PronunciationDictPhonemeItemAlphabetIpa PronunciationDictPhonemeItemAlphabet = "ipa"
)

// The item type.
type PronunciationDictPhonemeItemType string

const (
	PronunciationDictPhonemeItemTypePhoneme PronunciationDictPhonemeItemType = "phoneme"
)

// A phoneme pronunciation item. When the `text` value is found in input, it is
// pronounced using the specified IPA phoneme notation.
//
// The properties Alphabet, Phoneme, Text, Type are required.
type PronunciationDictPhonemeItemParam struct {
	// The phonetic alphabet used for the phoneme notation.
	//
	// Any of "ipa".
	Alphabet PronunciationDictPhonemeItemAlphabet `json:"alphabet,omitzero" api:"required"`
	// The phoneme notation representing the desired pronunciation.
	Phoneme string `json:"phoneme" api:"required"`
	// The text to match in the input. Case-insensitive matching is used during
	// synthesis.
	Text string `json:"text" api:"required"`
	// The item type.
	//
	// Any of "phoneme".
	Type PronunciationDictPhonemeItemType `json:"type,omitzero" api:"required"`
	paramObj
}

func (r PronunciationDictPhonemeItemParam) MarshalJSON() (data []byte, err error) {
	type shadow PronunciationDictPhonemeItemParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PronunciationDictPhonemeItemParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing a single pronunciation dictionary.
type PronunciationDictNewResponse struct {
	// A pronunciation dictionary record.
	Data PronunciationDictData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PronunciationDictNewResponse) RawJSON() string { return r.JSON.raw }
func (r *PronunciationDictNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing a single pronunciation dictionary.
type PronunciationDictGetResponse struct {
	// A pronunciation dictionary record.
	Data PronunciationDictData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PronunciationDictGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PronunciationDictGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing a single pronunciation dictionary.
type PronunciationDictUpdateResponse struct {
	// A pronunciation dictionary record.
	Data PronunciationDictData `json:"data"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PronunciationDictUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *PronunciationDictUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PronunciationDictNewParams struct {
	// List of pronunciation items (alias or phoneme type). At least one item is
	// required.
	Items []PronunciationDictNewParamsItemUnion `json:"items,omitzero" api:"required"`
	// Human-readable name. Must be unique within the organization.
	Name string `json:"name" api:"required"`
	paramObj
}

func (r PronunciationDictNewParams) MarshalJSON() (data []byte, err error) {
	type shadow PronunciationDictNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PronunciationDictNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PronunciationDictNewParamsItemUnion struct {
	OfAlias   *PronunciationDictAliasItemParam   `json:",omitzero,inline"`
	OfPhoneme *PronunciationDictPhonemeItemParam `json:",omitzero,inline"`
	paramUnion
}

func (u PronunciationDictNewParamsItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAlias, u.OfPhoneme)
}
func (u *PronunciationDictNewParamsItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PronunciationDictNewParamsItemUnion) asAny() any {
	if !param.IsOmitted(u.OfAlias) {
		return u.OfAlias
	} else if !param.IsOmitted(u.OfPhoneme) {
		return u.OfPhoneme
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PronunciationDictNewParamsItemUnion) GetAlias() *string {
	if vt := u.OfAlias; vt != nil {
		return &vt.Alias
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PronunciationDictNewParamsItemUnion) GetAlphabet() *string {
	if vt := u.OfPhoneme; vt != nil {
		return (*string)(&vt.Alphabet)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PronunciationDictNewParamsItemUnion) GetPhoneme() *string {
	if vt := u.OfPhoneme; vt != nil {
		return &vt.Phoneme
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PronunciationDictNewParamsItemUnion) GetText() *string {
	if vt := u.OfAlias; vt != nil {
		return (*string)(&vt.Text)
	} else if vt := u.OfPhoneme; vt != nil {
		return (*string)(&vt.Text)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PronunciationDictNewParamsItemUnion) GetType() *string {
	if vt := u.OfAlias; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPhoneme; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[PronunciationDictNewParamsItemUnion](
		"type",
		apijson.Discriminator[PronunciationDictAliasItemParam]("alias"),
		apijson.Discriminator[PronunciationDictPhonemeItemParam]("phoneme"),
	)
}

type PronunciationDictUpdateParams struct {
	// Updated dictionary name.
	Name param.Opt[string] `json:"name,omitzero"`
	// Updated list of pronunciation items (alias or phoneme type).
	Items []PronunciationDictUpdateParamsItemUnion `json:"items,omitzero"`
	paramObj
}

func (r PronunciationDictUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow PronunciationDictUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PronunciationDictUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type PronunciationDictUpdateParamsItemUnion struct {
	OfAlias   *PronunciationDictAliasItemParam   `json:",omitzero,inline"`
	OfPhoneme *PronunciationDictPhonemeItemParam `json:",omitzero,inline"`
	paramUnion
}

func (u PronunciationDictUpdateParamsItemUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAlias, u.OfPhoneme)
}
func (u *PronunciationDictUpdateParamsItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *PronunciationDictUpdateParamsItemUnion) asAny() any {
	if !param.IsOmitted(u.OfAlias) {
		return u.OfAlias
	} else if !param.IsOmitted(u.OfPhoneme) {
		return u.OfPhoneme
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PronunciationDictUpdateParamsItemUnion) GetAlias() *string {
	if vt := u.OfAlias; vt != nil {
		return &vt.Alias
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PronunciationDictUpdateParamsItemUnion) GetAlphabet() *string {
	if vt := u.OfPhoneme; vt != nil {
		return (*string)(&vt.Alphabet)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PronunciationDictUpdateParamsItemUnion) GetPhoneme() *string {
	if vt := u.OfPhoneme; vt != nil {
		return &vt.Phoneme
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PronunciationDictUpdateParamsItemUnion) GetText() *string {
	if vt := u.OfAlias; vt != nil {
		return (*string)(&vt.Text)
	} else if vt := u.OfPhoneme; vt != nil {
		return (*string)(&vt.Text)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u PronunciationDictUpdateParamsItemUnion) GetType() *string {
	if vt := u.OfAlias; vt != nil {
		return (*string)(&vt.Type)
	} else if vt := u.OfPhoneme; vt != nil {
		return (*string)(&vt.Type)
	}
	return nil
}

func init() {
	apijson.RegisterUnion[PronunciationDictUpdateParamsItemUnion](
		"type",
		apijson.Discriminator[PronunciationDictAliasItemParam]("alias"),
		apijson.Discriminator[PronunciationDictPhonemeItemParam]("phoneme"),
	)
}

type PronunciationDictListParams struct {
	// Page number (1-based). Defaults to 1.
	PageNumber param.Opt[int64] `query:"page[number],omitzero" json:"-"`
	// Number of results per page. Defaults to 20, maximum 250.
	PageSize param.Opt[int64] `query:"page[size],omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PronunciationDictListParams]'s query parameters as
// `url.Values`.
func (r PronunciationDictListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
