# DocServiceDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 
**ContentType** | Pointer to **string** | The document&#39;s content_type. | [optional] [readonly] 
**Size** | Pointer to [**DocServiceDocumentAllOfSize**](DocServiceDocumentAllOfSize.md) |  | [optional] 
**Status** | Pointer to **string** | Indicates the current document reviewing status | [optional] [readonly] 
**Sha256** | Pointer to **string** | The document&#39;s SHA256 hash provided for optional verification purposes. | [optional] [readonly] 
**Filename** | Pointer to **string** | The filename of the document. | [optional] 
**CustomerReference** | Pointer to **string** | Optional reference string for customer tracking. | [optional] 

## Methods

### NewDocServiceDocument

`func NewDocServiceDocument() *DocServiceDocument`

NewDocServiceDocument instantiates a new DocServiceDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocServiceDocumentWithDefaults

`func NewDocServiceDocumentWithDefaults() *DocServiceDocument`

NewDocServiceDocumentWithDefaults instantiates a new DocServiceDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DocServiceDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DocServiceDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DocServiceDocument) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DocServiceDocument) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *DocServiceDocument) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *DocServiceDocument) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *DocServiceDocument) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *DocServiceDocument) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *DocServiceDocument) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *DocServiceDocument) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *DocServiceDocument) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *DocServiceDocument) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *DocServiceDocument) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *DocServiceDocument) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *DocServiceDocument) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *DocServiceDocument) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetContentType

`func (o *DocServiceDocument) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *DocServiceDocument) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *DocServiceDocument) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *DocServiceDocument) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetSize

`func (o *DocServiceDocument) GetSize() DocServiceDocumentAllOfSize`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *DocServiceDocument) GetSizeOk() (*DocServiceDocumentAllOfSize, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *DocServiceDocument) SetSize(v DocServiceDocumentAllOfSize)`

SetSize sets Size field to given value.

### HasSize

`func (o *DocServiceDocument) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStatus

`func (o *DocServiceDocument) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocServiceDocument) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocServiceDocument) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DocServiceDocument) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSha256

`func (o *DocServiceDocument) GetSha256() string`

GetSha256 returns the Sha256 field if non-nil, zero value otherwise.

### GetSha256Ok

`func (o *DocServiceDocument) GetSha256Ok() (*string, bool)`

GetSha256Ok returns a tuple with the Sha256 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSha256

`func (o *DocServiceDocument) SetSha256(v string)`

SetSha256 sets Sha256 field to given value.

### HasSha256

`func (o *DocServiceDocument) HasSha256() bool`

HasSha256 returns a boolean if a field has been set.

### GetFilename

`func (o *DocServiceDocument) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *DocServiceDocument) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *DocServiceDocument) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *DocServiceDocument) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetCustomerReference

`func (o *DocServiceDocument) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *DocServiceDocument) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *DocServiceDocument) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *DocServiceDocument) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


