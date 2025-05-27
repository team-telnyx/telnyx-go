# PortOutSupportingDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**RecordType** | **string** | Identifies the type of the resource. | [readonly] 
**Type** | **string** | Identifies the type of the document | [readonly] 
**PortoutId** | **string** | Identifies the associated port request | [readonly] 
**DocumentId** | **string** | Identifies the associated document | 
**CreatedAt** | **string** | Supporting document creation timestamp in ISO 8601 format | 
**UpdatedAt** | **string** | Supporting document last changed timestamp in ISO 8601 format | 

## Methods

### NewPortOutSupportingDocument

`func NewPortOutSupportingDocument(id string, recordType string, type_ string, portoutId string, documentId string, createdAt string, updatedAt string, ) *PortOutSupportingDocument`

NewPortOutSupportingDocument instantiates a new PortOutSupportingDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortOutSupportingDocumentWithDefaults

`func NewPortOutSupportingDocumentWithDefaults() *PortOutSupportingDocument`

NewPortOutSupportingDocumentWithDefaults instantiates a new PortOutSupportingDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortOutSupportingDocument) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortOutSupportingDocument) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortOutSupportingDocument) SetId(v string)`

SetId sets Id field to given value.


### GetRecordType

`func (o *PortOutSupportingDocument) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PortOutSupportingDocument) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PortOutSupportingDocument) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetType

`func (o *PortOutSupportingDocument) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PortOutSupportingDocument) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PortOutSupportingDocument) SetType(v string)`

SetType sets Type field to given value.


### GetPortoutId

`func (o *PortOutSupportingDocument) GetPortoutId() string`

GetPortoutId returns the PortoutId field if non-nil, zero value otherwise.

### GetPortoutIdOk

`func (o *PortOutSupportingDocument) GetPortoutIdOk() (*string, bool)`

GetPortoutIdOk returns a tuple with the PortoutId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortoutId

`func (o *PortOutSupportingDocument) SetPortoutId(v string)`

SetPortoutId sets PortoutId field to given value.


### GetDocumentId

`func (o *PortOutSupportingDocument) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *PortOutSupportingDocument) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *PortOutSupportingDocument) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.


### GetCreatedAt

`func (o *PortOutSupportingDocument) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortOutSupportingDocument) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortOutSupportingDocument) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PortOutSupportingDocument) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PortOutSupportingDocument) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PortOutSupportingDocument) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


