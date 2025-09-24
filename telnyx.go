// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/url"
	"time"

	"github.com/team-telnyx/telnyx-go/v3/internal/apiform"
	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	"github.com/team-telnyx/telnyx-go/v3/internal/apiquery"
	shimjson "github.com/team-telnyx/telnyx-go/v3/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

type DeleteObjectsResponse = any

type ListBucketsResponse struct {
	Buckets []ListBucketsResponseBucket `json:"Buckets"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Buckets     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListBucketsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListBucketsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListBucketsResponseBucket struct {
	CreationDate time.Time `json:"CreationDate" format:"date-time"`
	Name         string    `json:"Name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreationDate respjson.Field
		Name         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListBucketsResponseBucket) RawJSON() string { return r.JSON.raw }
func (r *ListBucketsResponseBucket) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListObjectsResponse struct {
	Contents []ListObjectsResponseContent `json:"Contents"`
	Name     string                       `json:"Name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Contents    respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListObjectsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListObjectsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListObjectsResponseContent struct {
	Key          string    `json:"Key"`
	LastModified time.Time `json:"LastModified" format:"date-time"`
	Size         float64   `json:"Size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Key          respjson.Field
		LastModified respjson.Field
		Size         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListObjectsResponseContent) RawJSON() string { return r.JSON.raw }
func (r *ListObjectsResponseContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NewBucketParams struct {
	LocationConstraint param.Opt[string] `json:"LocationConstraint,omitzero"`
	paramObj
}

func (r NewBucketParams) MarshalJSON() (data []byte, err error) {
	type shadow NewBucketParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NewBucketParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeleteObjectParams struct {
	BucketName string `path:"bucketName,required" json:"-"`
	paramObj
}

type DeleteObjectsParams struct {
	// Any of true.
	Delete bool `query:"delete,omitzero,required" json:"-"`
	Body   []DeleteObjectsParamsBody
	paramObj
}

func (r DeleteObjectsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Body)
}
func (r *DeleteObjectsParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}

// URLQuery serializes [DeleteObjectsParams]'s query parameters as `url.Values`.
func (r DeleteObjectsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type DeleteObjectsParamsBody struct {
	Key param.Opt[string] `json:"Key,omitzero"`
	paramObj
}

func (r DeleteObjectsParamsBody) MarshalJSON() (data []byte, err error) {
	type shadow DeleteObjectsParamsBody
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeleteObjectsParamsBody) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GetObjectParams struct {
	BucketName string            `path:"bucketName,required" json:"-"`
	UploadID   param.Opt[string] `query:"uploadId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [GetObjectParams]'s query parameters as `url.Values`.
func (r GetObjectParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ListObjectsParams struct {
	// Any of 2.
	ListType int64 `query:"list-type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ListObjectsParams]'s query parameters as `url.Values`.
func (r ListObjectsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type PutObjectParams struct {
	BucketName string `path:"bucketName,required" json:"-"`
	Body       io.Reader
	PartNumber param.Opt[string] `query:"partNumber,omitzero" json:"-"`
	UploadID   param.Opt[string] `query:"uploadId,omitzero" json:"-"`
	paramObj
}

func (r PutObjectParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// URLQuery serializes [PutObjectParams]'s query parameters as `url.Values`.
func (r PutObjectParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
