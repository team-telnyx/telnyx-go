# PortingAdditionalDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies this additional document | [optional] 
**DocumentType** | Pointer to **string** | Identifies the type of additional document | [optional] 
**DocumentId** | Pointer to **string** | Identifies the associated document | [optional] 
**Filename** | Pointer to **string** | The filename of the related document. | [optional] 
**ContentType** | Pointer to **string** | The content type of the related document. | [optional] 
**PortingOrderId** | Pointer to **string** | Identifies the associated porting order | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 

## Methods

### NewPortingAdditionalDocument

`func NewPortingAdditionalDocument() *PortingAdditionalDocument`

NewPortingAdditionalDocument instantiates a new PortingAdditionalDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortingAdditionalDocumentWithDefaults

`func NewPortingAdditionalDocumentWithDefaults() *PortingAdditionalDocument`

NewPortingAdditionalDocumentWithDefaults instantiates a new PortingAdditionalDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortingAdditionalDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortingAdditionalDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortingAdditionalDocument) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortingAdditionalDocument) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDocumentType

`func (o *PortingAdditionalDocument) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *PortingAdditionalDocument) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *PortingAdditionalDocument) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.

### HasDocumentType

`func (o *PortingAdditionalDocument) HasDocumentType() bool`

HasDocumentType returns a boolean if a field has been set.

### GetDocumentId

`func (o *PortingAdditionalDocument) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *PortingAdditionalDocument) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *PortingAdditionalDocument) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *PortingAdditionalDocument) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### GetFilename

`func (o *PortingAdditionalDocument) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *PortingAdditionalDocument) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *PortingAdditionalDocument) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *PortingAdditionalDocument) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetContentType

`func (o *PortingAdditionalDocument) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *PortingAdditionalDocument) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *PortingAdditionalDocument) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *PortingAdditionalDocument) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetPortingOrderId

`func (o *PortingAdditionalDocument) GetPortingOrderId() string`

GetPortingOrderId returns the PortingOrderId field if non-nil, zero value otherwise.

### GetPortingOrderIdOk

`func (o *PortingAdditionalDocument) GetPortingOrderIdOk() (*string, bool)`

GetPortingOrderIdOk returns a tuple with the PortingOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortingOrderId

`func (o *PortingAdditionalDocument) SetPortingOrderId(v string)`

SetPortingOrderId sets PortingOrderId field to given value.

### HasPortingOrderId

`func (o *PortingAdditionalDocument) HasPortingOrderId() bool`

HasPortingOrderId returns a boolean if a field has been set.

### GetRecordType

`func (o *PortingAdditionalDocument) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PortingAdditionalDocument) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PortingAdditionalDocument) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PortingAdditionalDocument) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PortingAdditionalDocument) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortingAdditionalDocument) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortingAdditionalDocument) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PortingAdditionalDocument) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PortingAdditionalDocument) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PortingAdditionalDocument) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PortingAdditionalDocument) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PortingAdditionalDocument) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


