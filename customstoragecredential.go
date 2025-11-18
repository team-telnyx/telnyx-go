// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package telnyx

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/team-telnyx/telnyx-go/v3/internal/apijson"
	shimjson "github.com/team-telnyx/telnyx-go/v3/internal/encoding/json"
	"github.com/team-telnyx/telnyx-go/v3/internal/requestconfig"
	"github.com/team-telnyx/telnyx-go/v3/option"
	"github.com/team-telnyx/telnyx-go/v3/packages/param"
	"github.com/team-telnyx/telnyx-go/v3/packages/respjson"
)

// CustomStorageCredentialService contains methods and other services that help
// with interacting with the telnyx API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomStorageCredentialService] method instead.
type CustomStorageCredentialService struct {
	Options []option.RequestOption
}

// NewCustomStorageCredentialService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomStorageCredentialService(opts ...option.RequestOption) (r CustomStorageCredentialService) {
	r = CustomStorageCredentialService{}
	r.Options = opts
	return
}

// Creates a custom storage credentials configuration.
func (r *CustomStorageCredentialService) New(ctx context.Context, connectionID string, body CustomStorageCredentialNewParams, opts ...option.RequestOption) (res *CustomStorageCredentialNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return
	}
	path := fmt.Sprintf("custom_storage_credentials/%s", connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns the information about custom storage credentials.
func (r *CustomStorageCredentialService) Get(ctx context.Context, connectionID string, opts ...option.RequestOption) (res *CustomStorageCredentialGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return
	}
	path := fmt.Sprintf("custom_storage_credentials/%s", connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a stored custom credentials configuration.
func (r *CustomStorageCredentialService) Update(ctx context.Context, connectionID string, body CustomStorageCredentialUpdateParams, opts ...option.RequestOption) (res *CustomStorageCredentialUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return
	}
	path := fmt.Sprintf("custom_storage_credentials/%s", connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Deletes a stored custom credentials configuration.
func (r *CustomStorageCredentialService) Delete(ctx context.Context, connectionID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if connectionID == "" {
		err = errors.New("missing required connection_id parameter")
		return
	}
	path := fmt.Sprintf("custom_storage_credentials/%s", connectionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type AzureConfigurationData struct {
	// Azure Blob Storage account key.
	AccountKey string `json:"account_key"`
	// Azure Blob Storage account name.
	AccountName string `json:"account_name"`
	// Name of the bucket to be used to store recording files.
	Bucket string `json:"bucket"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountKey  respjson.Field
		AccountName respjson.Field
		Bucket      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AzureConfigurationData) RawJSON() string { return r.JSON.raw }
func (r *AzureConfigurationData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AzureConfigurationData to a AzureConfigurationDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AzureConfigurationDataParam.Overrides()
func (r AzureConfigurationData) ToParam() AzureConfigurationDataParam {
	return param.Override[AzureConfigurationDataParam](json.RawMessage(r.RawJSON()))
}

type AzureConfigurationDataParam struct {
	// Azure Blob Storage account key.
	AccountKey param.Opt[string] `json:"account_key,omitzero"`
	// Azure Blob Storage account name.
	AccountName param.Opt[string] `json:"account_name,omitzero"`
	// Name of the bucket to be used to store recording files.
	Bucket param.Opt[string] `json:"bucket,omitzero"`
	paramObj
}

func (r AzureConfigurationDataParam) MarshalJSON() (data []byte, err error) {
	type shadow AzureConfigurationDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AzureConfigurationDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomStorageConfiguration struct {
	// Any of "gcs", "s3", "azure".
	Backend       CustomStorageConfigurationBackend            `json:"backend,required"`
	Configuration CustomStorageConfigurationConfigurationUnion `json:"configuration,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Backend       respjson.Field
		Configuration respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomStorageConfiguration) RawJSON() string { return r.JSON.raw }
func (r *CustomStorageConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CustomStorageConfiguration to a
// CustomStorageConfigurationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CustomStorageConfigurationParam.Overrides()
func (r CustomStorageConfiguration) ToParam() CustomStorageConfigurationParam {
	return param.Override[CustomStorageConfigurationParam](json.RawMessage(r.RawJSON()))
}

type CustomStorageConfigurationBackend string

const (
	CustomStorageConfigurationBackendGcs   CustomStorageConfigurationBackend = "gcs"
	CustomStorageConfigurationBackendS3    CustomStorageConfigurationBackend = "s3"
	CustomStorageConfigurationBackendAzure CustomStorageConfigurationBackend = "azure"
)

// CustomStorageConfigurationConfigurationUnion contains all possible properties
// and values from [GcsConfigurationData], [S3ConfigurationData],
// [AzureConfigurationData].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type CustomStorageConfigurationConfigurationUnion struct {
	Bucket string `json:"bucket"`
	// This field is from variant [GcsConfigurationData].
	Credentials string `json:"credentials"`
	// This field is from variant [S3ConfigurationData].
	AwsAccessKeyID string `json:"aws_access_key_id"`
	// This field is from variant [S3ConfigurationData].
	AwsSecretAccessKey string `json:"aws_secret_access_key"`
	// This field is from variant [S3ConfigurationData].
	Region string `json:"region"`
	// This field is from variant [AzureConfigurationData].
	AccountKey string `json:"account_key"`
	// This field is from variant [AzureConfigurationData].
	AccountName string `json:"account_name"`
	JSON        struct {
		Bucket             respjson.Field
		Credentials        respjson.Field
		AwsAccessKeyID     respjson.Field
		AwsSecretAccessKey respjson.Field
		Region             respjson.Field
		AccountKey         respjson.Field
		AccountName        respjson.Field
		raw                string
	} `json:"-"`
}

func (u CustomStorageConfigurationConfigurationUnion) AsGoogleCloudStorageConfigurationData() (v GcsConfigurationData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomStorageConfigurationConfigurationUnion) AsAwsS3StorageConfigurationData() (v S3ConfigurationData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u CustomStorageConfigurationConfigurationUnion) AsAzureBlobStorageConfigurationData() (v AzureConfigurationData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u CustomStorageConfigurationConfigurationUnion) RawJSON() string { return u.JSON.raw }

func (r *CustomStorageConfigurationConfigurationUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties Backend, Configuration are required.
type CustomStorageConfigurationParam struct {
	// Any of "gcs", "s3", "azure".
	Backend       CustomStorageConfigurationBackend                 `json:"backend,omitzero,required"`
	Configuration CustomStorageConfigurationConfigurationUnionParam `json:"configuration,omitzero,required"`
	paramObj
}

func (r CustomStorageConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow CustomStorageConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomStorageConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type CustomStorageConfigurationConfigurationUnionParam struct {
	OfGoogleCloudStorageConfigurationData *GcsConfigurationDataParam   `json:",omitzero,inline"`
	OfAwsS3StorageConfigurationData       *S3ConfigurationDataParam    `json:",omitzero,inline"`
	OfAzureBlobStorageConfigurationData   *AzureConfigurationDataParam `json:",omitzero,inline"`
	paramUnion
}

func (u CustomStorageConfigurationConfigurationUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfGoogleCloudStorageConfigurationData, u.OfAwsS3StorageConfigurationData, u.OfAzureBlobStorageConfigurationData)
}
func (u *CustomStorageConfigurationConfigurationUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *CustomStorageConfigurationConfigurationUnionParam) asAny() any {
	if !param.IsOmitted(u.OfGoogleCloudStorageConfigurationData) {
		return u.OfGoogleCloudStorageConfigurationData
	} else if !param.IsOmitted(u.OfAwsS3StorageConfigurationData) {
		return u.OfAwsS3StorageConfigurationData
	} else if !param.IsOmitted(u.OfAzureBlobStorageConfigurationData) {
		return u.OfAzureBlobStorageConfigurationData
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CustomStorageConfigurationConfigurationUnionParam) GetCredentials() *string {
	if vt := u.OfGoogleCloudStorageConfigurationData; vt != nil && vt.Credentials.Valid() {
		return &vt.Credentials.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CustomStorageConfigurationConfigurationUnionParam) GetAwsAccessKeyID() *string {
	if vt := u.OfAwsS3StorageConfigurationData; vt != nil && vt.AwsAccessKeyID.Valid() {
		return &vt.AwsAccessKeyID.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CustomStorageConfigurationConfigurationUnionParam) GetAwsSecretAccessKey() *string {
	if vt := u.OfAwsS3StorageConfigurationData; vt != nil && vt.AwsSecretAccessKey.Valid() {
		return &vt.AwsSecretAccessKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CustomStorageConfigurationConfigurationUnionParam) GetRegion() *string {
	if vt := u.OfAwsS3StorageConfigurationData; vt != nil && vt.Region.Valid() {
		return &vt.Region.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CustomStorageConfigurationConfigurationUnionParam) GetAccountKey() *string {
	if vt := u.OfAzureBlobStorageConfigurationData; vt != nil && vt.AccountKey.Valid() {
		return &vt.AccountKey.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CustomStorageConfigurationConfigurationUnionParam) GetAccountName() *string {
	if vt := u.OfAzureBlobStorageConfigurationData; vt != nil && vt.AccountName.Valid() {
		return &vt.AccountName.Value
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u CustomStorageConfigurationConfigurationUnionParam) GetBucket() *string {
	if vt := u.OfGoogleCloudStorageConfigurationData; vt != nil && vt.Bucket.Valid() {
		return &vt.Bucket.Value
	} else if vt := u.OfAwsS3StorageConfigurationData; vt != nil && vt.Bucket.Valid() {
		return &vt.Bucket.Value
	} else if vt := u.OfAzureBlobStorageConfigurationData; vt != nil && vt.Bucket.Valid() {
		return &vt.Bucket.Value
	}
	return nil
}

type GcsConfigurationData struct {
	// Name of the bucket to be used to store recording files.
	Bucket string `json:"bucket"`
	// Opaque credential token used to authenticate and authorize with storage
	// provider.
	Credentials string `json:"credentials"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bucket      respjson.Field
		Credentials respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GcsConfigurationData) RawJSON() string { return r.JSON.raw }
func (r *GcsConfigurationData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this GcsConfigurationData to a GcsConfigurationDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// GcsConfigurationDataParam.Overrides()
func (r GcsConfigurationData) ToParam() GcsConfigurationDataParam {
	return param.Override[GcsConfigurationDataParam](json.RawMessage(r.RawJSON()))
}

type GcsConfigurationDataParam struct {
	// Name of the bucket to be used to store recording files.
	Bucket param.Opt[string] `json:"bucket,omitzero"`
	// Opaque credential token used to authenticate and authorize with storage
	// provider.
	Credentials param.Opt[string] `json:"credentials,omitzero"`
	paramObj
}

func (r GcsConfigurationDataParam) MarshalJSON() (data []byte, err error) {
	type shadow GcsConfigurationDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GcsConfigurationDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type S3ConfigurationData struct {
	// AWS credentials access key id.
	AwsAccessKeyID string `json:"aws_access_key_id"`
	// AWS secret access key.
	AwsSecretAccessKey string `json:"aws_secret_access_key"`
	// Name of the bucket to be used to store recording files.
	Bucket string `json:"bucket"`
	// Region where the bucket is located.
	Region string `json:"region"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AwsAccessKeyID     respjson.Field
		AwsSecretAccessKey respjson.Field
		Bucket             respjson.Field
		Region             respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r S3ConfigurationData) RawJSON() string { return r.JSON.raw }
func (r *S3ConfigurationData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this S3ConfigurationData to a S3ConfigurationDataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// S3ConfigurationDataParam.Overrides()
func (r S3ConfigurationData) ToParam() S3ConfigurationDataParam {
	return param.Override[S3ConfigurationDataParam](json.RawMessage(r.RawJSON()))
}

type S3ConfigurationDataParam struct {
	// AWS credentials access key id.
	AwsAccessKeyID param.Opt[string] `json:"aws_access_key_id,omitzero"`
	// AWS secret access key.
	AwsSecretAccessKey param.Opt[string] `json:"aws_secret_access_key,omitzero"`
	// Name of the bucket to be used to store recording files.
	Bucket param.Opt[string] `json:"bucket,omitzero"`
	// Region where the bucket is located.
	Region param.Opt[string] `json:"region,omitzero"`
	paramObj
}

func (r S3ConfigurationDataParam) MarshalJSON() (data []byte, err error) {
	type shadow S3ConfigurationDataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *S3ConfigurationDataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CustomStorageCredentialNewResponse struct {
	// Uniquely identifies a Telnyx application (Call Control, TeXML) or Sip connection
	// resource.
	ConnectionID string                     `json:"connection_id,required"`
	Data         CustomStorageConfiguration `json:"data,required"`
	// Identifies record type.
	//
	// Any of "custom_storage_credentials".
	RecordType CustomStorageCredentialNewResponseRecordType `json:"record_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectionID respjson.Field
		Data         respjson.Field
		RecordType   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomStorageCredentialNewResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomStorageCredentialNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies record type.
type CustomStorageCredentialNewResponseRecordType string

const (
	CustomStorageCredentialNewResponseRecordTypeCustomStorageCredentials CustomStorageCredentialNewResponseRecordType = "custom_storage_credentials"
)

type CustomStorageCredentialGetResponse struct {
	// Uniquely identifies a Telnyx application (Call Control, TeXML) or Sip connection
	// resource.
	ConnectionID string                     `json:"connection_id,required"`
	Data         CustomStorageConfiguration `json:"data,required"`
	// Identifies record type.
	//
	// Any of "custom_storage_credentials".
	RecordType CustomStorageCredentialGetResponseRecordType `json:"record_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectionID respjson.Field
		Data         respjson.Field
		RecordType   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomStorageCredentialGetResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomStorageCredentialGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies record type.
type CustomStorageCredentialGetResponseRecordType string

const (
	CustomStorageCredentialGetResponseRecordTypeCustomStorageCredentials CustomStorageCredentialGetResponseRecordType = "custom_storage_credentials"
)

type CustomStorageCredentialUpdateResponse struct {
	// Uniquely identifies a Telnyx application (Call Control, TeXML) or Sip connection
	// resource.
	ConnectionID string                     `json:"connection_id,required"`
	Data         CustomStorageConfiguration `json:"data,required"`
	// Identifies record type.
	//
	// Any of "custom_storage_credentials".
	RecordType CustomStorageCredentialUpdateResponseRecordType `json:"record_type,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectionID respjson.Field
		Data         respjson.Field
		RecordType   respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomStorageCredentialUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomStorageCredentialUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifies record type.
type CustomStorageCredentialUpdateResponseRecordType string

const (
	CustomStorageCredentialUpdateResponseRecordTypeCustomStorageCredentials CustomStorageCredentialUpdateResponseRecordType = "custom_storage_credentials"
)

type CustomStorageCredentialNewParams struct {
	CustomStorageConfiguration CustomStorageConfigurationParam
	paramObj
}

func (r CustomStorageCredentialNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CustomStorageConfiguration)
}
func (r *CustomStorageCredentialNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CustomStorageConfiguration)
}

type CustomStorageCredentialUpdateParams struct {
	CustomStorageConfiguration CustomStorageConfigurationParam
	paramObj
}

func (r CustomStorageCredentialUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CustomStorageConfiguration)
}
func (r *CustomStorageCredentialUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CustomStorageConfiguration)
}
