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
	"github.com/team-telnyx/telnyx-go/v4/shared/constant"
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
func (r *PronunciationDictService) List(ctx context.Context, query PronunciationDictListParams, opts ...option.RequestOption) (res *pagination.DefaultFlatPagination[PronunciationDictListResponse], err error) {
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
func (r *PronunciationDictService) ListAutoPaging(ctx context.Context, query PronunciationDictListParams, opts ...option.RequestOption) *pagination.DefaultFlatPaginationAutoPager[PronunciationDictListResponse] {
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

// Response containing a single pronunciation dictionary.
type PronunciationDictNewResponse struct {
	// A pronunciation dictionary record.
	Data PronunciationDictNewResponseData `json:"data"`
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

// A pronunciation dictionary record.
type PronunciationDictNewResponseData struct {
	// Unique identifier for the pronunciation dictionary.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 timestamp with millisecond precision.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// List of pronunciation items (alias or phoneme type).
	Items []PronunciationDictNewResponseDataItemUnion `json:"items"`
	// Human-readable name for the dictionary. Must be unique within the organization.
	Name string `json:"name"`
	// Identifies the resource type.
	//
	// Any of "pronunciation_dict".
	RecordType string `json:"record_type"`
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
func (r PronunciationDictNewResponseData) RawJSON() string { return r.JSON.raw }
func (r *PronunciationDictNewResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PronunciationDictNewResponseDataItemUnion contains all possible properties and
// values from [PronunciationDictNewResponseDataItemAlias],
// [PronunciationDictNewResponseDataItemPhoneme].
//
// Use the [PronunciationDictNewResponseDataItemUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PronunciationDictNewResponseDataItemUnion struct {
	// This field is from variant [PronunciationDictNewResponseDataItemAlias].
	Alias string `json:"alias"`
	Text  string `json:"text"`
	// Any of "alias", "phoneme".
	Type string `json:"type"`
	// This field is from variant [PronunciationDictNewResponseDataItemPhoneme].
	Alphabet string `json:"alphabet"`
	// This field is from variant [PronunciationDictNewResponseDataItemPhoneme].
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

// anyPronunciationDictNewResponseDataItem is implemented by each variant of
// [PronunciationDictNewResponseDataItemUnion] to add type safety for the return
// type of [PronunciationDictNewResponseDataItemUnion.AsAny]
type anyPronunciationDictNewResponseDataItem interface {
	implPronunciationDictNewResponseDataItemUnion()
}

func (PronunciationDictNewResponseDataItemAlias) implPronunciationDictNewResponseDataItemUnion()   {}
func (PronunciationDictNewResponseDataItemPhoneme) implPronunciationDictNewResponseDataItemUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PronunciationDictNewResponseDataItemUnion.AsAny().(type) {
//	case telnyx.PronunciationDictNewResponseDataItemAlias:
//	case telnyx.PronunciationDictNewResponseDataItemPhoneme:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PronunciationDictNewResponseDataItemUnion) AsAny() anyPronunciationDictNewResponseDataItem {
	switch u.Type {
	case "alias":
		return u.AsAlias()
	case "phoneme":
		return u.AsPhoneme()
	}
	return nil
}

func (u PronunciationDictNewResponseDataItemUnion) AsAlias() (v PronunciationDictNewResponseDataItemAlias) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PronunciationDictNewResponseDataItemUnion) AsPhoneme() (v PronunciationDictNewResponseDataItemPhoneme) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PronunciationDictNewResponseDataItemUnion) RawJSON() string { return u.JSON.raw }

func (r *PronunciationDictNewResponseDataItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An alias pronunciation item. When the `text` value is found in input, it is
// replaced with the `alias` before speech synthesis.
type PronunciationDictNewResponseDataItemAlias struct {
	// The replacement text that will be spoken instead.
	Alias string `json:"alias" api:"required"`
	// The text to match in the input. Case-insensitive matching is used during
	// synthesis.
	Text string `json:"text" api:"required"`
	// The item type.
	Type constant.Alias `json:"type" default:"alias"`
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
func (r PronunciationDictNewResponseDataItemAlias) RawJSON() string { return r.JSON.raw }
func (r *PronunciationDictNewResponseDataItemAlias) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A phoneme pronunciation item. When the `text` value is found in input, it is
// pronounced using the specified IPA phoneme notation.
type PronunciationDictNewResponseDataItemPhoneme struct {
	// The phonetic alphabet used for the phoneme notation.
	//
	// Any of "ipa".
	Alphabet string `json:"alphabet" api:"required"`
	// The phoneme notation representing the desired pronunciation.
	Phoneme string `json:"phoneme" api:"required"`
	// The text to match in the input. Case-insensitive matching is used during
	// synthesis.
	Text string `json:"text" api:"required"`
	// The item type.
	Type constant.Phoneme `json:"type" default:"phoneme"`
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
func (r PronunciationDictNewResponseDataItemPhoneme) RawJSON() string { return r.JSON.raw }
func (r *PronunciationDictNewResponseDataItemPhoneme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing a single pronunciation dictionary.
type PronunciationDictGetResponse struct {
	// A pronunciation dictionary record.
	Data PronunciationDictGetResponseData `json:"data"`
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

// A pronunciation dictionary record.
type PronunciationDictGetResponseData struct {
	// Unique identifier for the pronunciation dictionary.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 timestamp with millisecond precision.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// List of pronunciation items (alias or phoneme type).
	Items []PronunciationDictGetResponseDataItemUnion `json:"items"`
	// Human-readable name for the dictionary. Must be unique within the organization.
	Name string `json:"name"`
	// Identifies the resource type.
	//
	// Any of "pronunciation_dict".
	RecordType string `json:"record_type"`
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
func (r PronunciationDictGetResponseData) RawJSON() string { return r.JSON.raw }
func (r *PronunciationDictGetResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PronunciationDictGetResponseDataItemUnion contains all possible properties and
// values from [PronunciationDictGetResponseDataItemAlias],
// [PronunciationDictGetResponseDataItemPhoneme].
//
// Use the [PronunciationDictGetResponseDataItemUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PronunciationDictGetResponseDataItemUnion struct {
	// This field is from variant [PronunciationDictGetResponseDataItemAlias].
	Alias string `json:"alias"`
	Text  string `json:"text"`
	// Any of "alias", "phoneme".
	Type string `json:"type"`
	// This field is from variant [PronunciationDictGetResponseDataItemPhoneme].
	Alphabet string `json:"alphabet"`
	// This field is from variant [PronunciationDictGetResponseDataItemPhoneme].
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

// anyPronunciationDictGetResponseDataItem is implemented by each variant of
// [PronunciationDictGetResponseDataItemUnion] to add type safety for the return
// type of [PronunciationDictGetResponseDataItemUnion.AsAny]
type anyPronunciationDictGetResponseDataItem interface {
	implPronunciationDictGetResponseDataItemUnion()
}

func (PronunciationDictGetResponseDataItemAlias) implPronunciationDictGetResponseDataItemUnion()   {}
func (PronunciationDictGetResponseDataItemPhoneme) implPronunciationDictGetResponseDataItemUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PronunciationDictGetResponseDataItemUnion.AsAny().(type) {
//	case telnyx.PronunciationDictGetResponseDataItemAlias:
//	case telnyx.PronunciationDictGetResponseDataItemPhoneme:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PronunciationDictGetResponseDataItemUnion) AsAny() anyPronunciationDictGetResponseDataItem {
	switch u.Type {
	case "alias":
		return u.AsAlias()
	case "phoneme":
		return u.AsPhoneme()
	}
	return nil
}

func (u PronunciationDictGetResponseDataItemUnion) AsAlias() (v PronunciationDictGetResponseDataItemAlias) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PronunciationDictGetResponseDataItemUnion) AsPhoneme() (v PronunciationDictGetResponseDataItemPhoneme) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PronunciationDictGetResponseDataItemUnion) RawJSON() string { return u.JSON.raw }

func (r *PronunciationDictGetResponseDataItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An alias pronunciation item. When the `text` value is found in input, it is
// replaced with the `alias` before speech synthesis.
type PronunciationDictGetResponseDataItemAlias struct {
	// The replacement text that will be spoken instead.
	Alias string `json:"alias" api:"required"`
	// The text to match in the input. Case-insensitive matching is used during
	// synthesis.
	Text string `json:"text" api:"required"`
	// The item type.
	Type constant.Alias `json:"type" default:"alias"`
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
func (r PronunciationDictGetResponseDataItemAlias) RawJSON() string { return r.JSON.raw }
func (r *PronunciationDictGetResponseDataItemAlias) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A phoneme pronunciation item. When the `text` value is found in input, it is
// pronounced using the specified IPA phoneme notation.
type PronunciationDictGetResponseDataItemPhoneme struct {
	// The phonetic alphabet used for the phoneme notation.
	//
	// Any of "ipa".
	Alphabet string `json:"alphabet" api:"required"`
	// The phoneme notation representing the desired pronunciation.
	Phoneme string `json:"phoneme" api:"required"`
	// The text to match in the input. Case-insensitive matching is used during
	// synthesis.
	Text string `json:"text" api:"required"`
	// The item type.
	Type constant.Phoneme `json:"type" default:"phoneme"`
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
func (r PronunciationDictGetResponseDataItemPhoneme) RawJSON() string { return r.JSON.raw }
func (r *PronunciationDictGetResponseDataItemPhoneme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response containing a single pronunciation dictionary.
type PronunciationDictUpdateResponse struct {
	// A pronunciation dictionary record.
	Data PronunciationDictUpdateResponseData `json:"data"`
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

// A pronunciation dictionary record.
type PronunciationDictUpdateResponseData struct {
	// Unique identifier for the pronunciation dictionary.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 timestamp with millisecond precision.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// List of pronunciation items (alias or phoneme type).
	Items []PronunciationDictUpdateResponseDataItemUnion `json:"items"`
	// Human-readable name for the dictionary. Must be unique within the organization.
	Name string `json:"name"`
	// Identifies the resource type.
	//
	// Any of "pronunciation_dict".
	RecordType string `json:"record_type"`
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
func (r PronunciationDictUpdateResponseData) RawJSON() string { return r.JSON.raw }
func (r *PronunciationDictUpdateResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PronunciationDictUpdateResponseDataItemUnion contains all possible properties
// and values from [PronunciationDictUpdateResponseDataItemAlias],
// [PronunciationDictUpdateResponseDataItemPhoneme].
//
// Use the [PronunciationDictUpdateResponseDataItemUnion.AsAny] method to switch on
// the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PronunciationDictUpdateResponseDataItemUnion struct {
	// This field is from variant [PronunciationDictUpdateResponseDataItemAlias].
	Alias string `json:"alias"`
	Text  string `json:"text"`
	// Any of "alias", "phoneme".
	Type string `json:"type"`
	// This field is from variant [PronunciationDictUpdateResponseDataItemPhoneme].
	Alphabet string `json:"alphabet"`
	// This field is from variant [PronunciationDictUpdateResponseDataItemPhoneme].
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

// anyPronunciationDictUpdateResponseDataItem is implemented by each variant of
// [PronunciationDictUpdateResponseDataItemUnion] to add type safety for the return
// type of [PronunciationDictUpdateResponseDataItemUnion.AsAny]
type anyPronunciationDictUpdateResponseDataItem interface {
	implPronunciationDictUpdateResponseDataItemUnion()
}

func (PronunciationDictUpdateResponseDataItemAlias) implPronunciationDictUpdateResponseDataItemUnion() {
}
func (PronunciationDictUpdateResponseDataItemPhoneme) implPronunciationDictUpdateResponseDataItemUnion() {
}

// Use the following switch statement to find the correct variant
//
//	switch variant := PronunciationDictUpdateResponseDataItemUnion.AsAny().(type) {
//	case telnyx.PronunciationDictUpdateResponseDataItemAlias:
//	case telnyx.PronunciationDictUpdateResponseDataItemPhoneme:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PronunciationDictUpdateResponseDataItemUnion) AsAny() anyPronunciationDictUpdateResponseDataItem {
	switch u.Type {
	case "alias":
		return u.AsAlias()
	case "phoneme":
		return u.AsPhoneme()
	}
	return nil
}

func (u PronunciationDictUpdateResponseDataItemUnion) AsAlias() (v PronunciationDictUpdateResponseDataItemAlias) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PronunciationDictUpdateResponseDataItemUnion) AsPhoneme() (v PronunciationDictUpdateResponseDataItemPhoneme) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PronunciationDictUpdateResponseDataItemUnion) RawJSON() string { return u.JSON.raw }

func (r *PronunciationDictUpdateResponseDataItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An alias pronunciation item. When the `text` value is found in input, it is
// replaced with the `alias` before speech synthesis.
type PronunciationDictUpdateResponseDataItemAlias struct {
	// The replacement text that will be spoken instead.
	Alias string `json:"alias" api:"required"`
	// The text to match in the input. Case-insensitive matching is used during
	// synthesis.
	Text string `json:"text" api:"required"`
	// The item type.
	Type constant.Alias `json:"type" default:"alias"`
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
func (r PronunciationDictUpdateResponseDataItemAlias) RawJSON() string { return r.JSON.raw }
func (r *PronunciationDictUpdateResponseDataItemAlias) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A phoneme pronunciation item. When the `text` value is found in input, it is
// pronounced using the specified IPA phoneme notation.
type PronunciationDictUpdateResponseDataItemPhoneme struct {
	// The phonetic alphabet used for the phoneme notation.
	//
	// Any of "ipa".
	Alphabet string `json:"alphabet" api:"required"`
	// The phoneme notation representing the desired pronunciation.
	Phoneme string `json:"phoneme" api:"required"`
	// The text to match in the input. Case-insensitive matching is used during
	// synthesis.
	Text string `json:"text" api:"required"`
	// The item type.
	Type constant.Phoneme `json:"type" default:"phoneme"`
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
func (r PronunciationDictUpdateResponseDataItemPhoneme) RawJSON() string { return r.JSON.raw }
func (r *PronunciationDictUpdateResponseDataItemPhoneme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A pronunciation dictionary record.
type PronunciationDictListResponse struct {
	// Unique identifier for the pronunciation dictionary.
	ID string `json:"id" format:"uuid"`
	// ISO 8601 timestamp with millisecond precision.
	CreatedAt time.Time `json:"created_at" format:"date-time"`
	// List of pronunciation items (alias or phoneme type).
	Items []PronunciationDictListResponseItemUnion `json:"items"`
	// Human-readable name for the dictionary. Must be unique within the organization.
	Name string `json:"name"`
	// Identifies the resource type.
	//
	// Any of "pronunciation_dict".
	RecordType PronunciationDictListResponseRecordType `json:"record_type"`
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
func (r PronunciationDictListResponse) RawJSON() string { return r.JSON.raw }
func (r *PronunciationDictListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PronunciationDictListResponseItemUnion contains all possible properties and
// values from [PronunciationDictListResponseItemAlias],
// [PronunciationDictListResponseItemPhoneme].
//
// Use the [PronunciationDictListResponseItemUnion.AsAny] method to switch on the
// variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type PronunciationDictListResponseItemUnion struct {
	// This field is from variant [PronunciationDictListResponseItemAlias].
	Alias string `json:"alias"`
	Text  string `json:"text"`
	// Any of "alias", "phoneme".
	Type string `json:"type"`
	// This field is from variant [PronunciationDictListResponseItemPhoneme].
	Alphabet string `json:"alphabet"`
	// This field is from variant [PronunciationDictListResponseItemPhoneme].
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

// anyPronunciationDictListResponseItem is implemented by each variant of
// [PronunciationDictListResponseItemUnion] to add type safety for the return type
// of [PronunciationDictListResponseItemUnion.AsAny]
type anyPronunciationDictListResponseItem interface {
	implPronunciationDictListResponseItemUnion()
}

func (PronunciationDictListResponseItemAlias) implPronunciationDictListResponseItemUnion()   {}
func (PronunciationDictListResponseItemPhoneme) implPronunciationDictListResponseItemUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := PronunciationDictListResponseItemUnion.AsAny().(type) {
//	case telnyx.PronunciationDictListResponseItemAlias:
//	case telnyx.PronunciationDictListResponseItemPhoneme:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u PronunciationDictListResponseItemUnion) AsAny() anyPronunciationDictListResponseItem {
	switch u.Type {
	case "alias":
		return u.AsAlias()
	case "phoneme":
		return u.AsPhoneme()
	}
	return nil
}

func (u PronunciationDictListResponseItemUnion) AsAlias() (v PronunciationDictListResponseItemAlias) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u PronunciationDictListResponseItemUnion) AsPhoneme() (v PronunciationDictListResponseItemPhoneme) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u PronunciationDictListResponseItemUnion) RawJSON() string { return u.JSON.raw }

func (r *PronunciationDictListResponseItemUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An alias pronunciation item. When the `text` value is found in input, it is
// replaced with the `alias` before speech synthesis.
type PronunciationDictListResponseItemAlias struct {
	// The replacement text that will be spoken instead.
	Alias string `json:"alias" api:"required"`
	// The text to match in the input. Case-insensitive matching is used during
	// synthesis.
	Text string `json:"text" api:"required"`
	// The item type.
	Type constant.Alias `json:"type" default:"alias"`
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
func (r PronunciationDictListResponseItemAlias) RawJSON() string { return r.JSON.raw }
func (r *PronunciationDictListResponseItemAlias) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A phoneme pronunciation item. When the `text` value is found in input, it is
// pronounced using the specified IPA phoneme notation.
type PronunciationDictListResponseItemPhoneme struct {
	// The phonetic alphabet used for the phoneme notation.
	//
	// Any of "ipa".
	Alphabet string `json:"alphabet" api:"required"`
	// The phoneme notation representing the desired pronunciation.
	Phoneme string `json:"phoneme" api:"required"`
	// The text to match in the input. Case-insensitive matching is used during
	// synthesis.
	Text string `json:"text" api:"required"`
	// The item type.
	Type constant.Phoneme `json:"type" default:"phoneme"`
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
func (r PronunciationDictListResponseItemPhoneme) RawJSON() string { return r.JSON.raw }
func (r *PronunciationDictListResponseItemPhoneme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies the resource type.
type PronunciationDictListResponseRecordType string

const (
	PronunciationDictListResponseRecordTypePronunciationDict PronunciationDictListResponseRecordType = "pronunciation_dict"
)

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
	OfAlias   *PronunciationDictNewParamsItemAlias   `json:",omitzero,inline"`
	OfPhoneme *PronunciationDictNewParamsItemPhoneme `json:",omitzero,inline"`
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
		return &vt.Alphabet
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
		apijson.Discriminator[PronunciationDictNewParamsItemAlias]("alias"),
		apijson.Discriminator[PronunciationDictNewParamsItemPhoneme]("phoneme"),
	)
}

// An alias pronunciation item. When the `text` value is found in input, it is
// replaced with the `alias` before speech synthesis.
//
// The properties Alias, Text, Type are required.
type PronunciationDictNewParamsItemAlias struct {
	// The replacement text that will be spoken instead.
	Alias string `json:"alias" api:"required"`
	// The text to match in the input. Case-insensitive matching is used during
	// synthesis.
	Text string `json:"text" api:"required"`
	// The item type.
	//
	// This field can be elided, and will marshal its zero value as "alias".
	Type constant.Alias `json:"type" default:"alias"`
	paramObj
}

func (r PronunciationDictNewParamsItemAlias) MarshalJSON() (data []byte, err error) {
	type shadow PronunciationDictNewParamsItemAlias
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PronunciationDictNewParamsItemAlias) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A phoneme pronunciation item. When the `text` value is found in input, it is
// pronounced using the specified IPA phoneme notation.
//
// The properties Alphabet, Phoneme, Text, Type are required.
type PronunciationDictNewParamsItemPhoneme struct {
	// The phonetic alphabet used for the phoneme notation.
	//
	// Any of "ipa".
	Alphabet string `json:"alphabet,omitzero" api:"required"`
	// The phoneme notation representing the desired pronunciation.
	Phoneme string `json:"phoneme" api:"required"`
	// The text to match in the input. Case-insensitive matching is used during
	// synthesis.
	Text string `json:"text" api:"required"`
	// The item type.
	//
	// This field can be elided, and will marshal its zero value as "phoneme".
	Type constant.Phoneme `json:"type" default:"phoneme"`
	paramObj
}

func (r PronunciationDictNewParamsItemPhoneme) MarshalJSON() (data []byte, err error) {
	type shadow PronunciationDictNewParamsItemPhoneme
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PronunciationDictNewParamsItemPhoneme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PronunciationDictNewParamsItemPhoneme](
		"alphabet", "ipa",
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
	OfAlias   *PronunciationDictUpdateParamsItemAlias   `json:",omitzero,inline"`
	OfPhoneme *PronunciationDictUpdateParamsItemPhoneme `json:",omitzero,inline"`
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
		return &vt.Alphabet
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
		apijson.Discriminator[PronunciationDictUpdateParamsItemAlias]("alias"),
		apijson.Discriminator[PronunciationDictUpdateParamsItemPhoneme]("phoneme"),
	)
}

// An alias pronunciation item. When the `text` value is found in input, it is
// replaced with the `alias` before speech synthesis.
//
// The properties Alias, Text, Type are required.
type PronunciationDictUpdateParamsItemAlias struct {
	// The replacement text that will be spoken instead.
	Alias string `json:"alias" api:"required"`
	// The text to match in the input. Case-insensitive matching is used during
	// synthesis.
	Text string `json:"text" api:"required"`
	// The item type.
	//
	// This field can be elided, and will marshal its zero value as "alias".
	Type constant.Alias `json:"type" default:"alias"`
	paramObj
}

func (r PronunciationDictUpdateParamsItemAlias) MarshalJSON() (data []byte, err error) {
	type shadow PronunciationDictUpdateParamsItemAlias
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PronunciationDictUpdateParamsItemAlias) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A phoneme pronunciation item. When the `text` value is found in input, it is
// pronounced using the specified IPA phoneme notation.
//
// The properties Alphabet, Phoneme, Text, Type are required.
type PronunciationDictUpdateParamsItemPhoneme struct {
	// The phonetic alphabet used for the phoneme notation.
	//
	// Any of "ipa".
	Alphabet string `json:"alphabet,omitzero" api:"required"`
	// The phoneme notation representing the desired pronunciation.
	Phoneme string `json:"phoneme" api:"required"`
	// The text to match in the input. Case-insensitive matching is used during
	// synthesis.
	Text string `json:"text" api:"required"`
	// The item type.
	//
	// This field can be elided, and will marshal its zero value as "phoneme".
	Type constant.Phoneme `json:"type" default:"phoneme"`
	paramObj
}

func (r PronunciationDictUpdateParamsItemPhoneme) MarshalJSON() (data []byte, err error) {
	type shadow PronunciationDictUpdateParamsItemPhoneme
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PronunciationDictUpdateParamsItemPhoneme) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PronunciationDictUpdateParamsItemPhoneme](
		"alphabet", "ipa",
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
