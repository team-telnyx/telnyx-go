# CreateDocumentRequestOneOf1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**File** | **string** | The Base64 encoded contents of the file you are uploading. | 
**Filename** | Pointer to **string** | The filename of the document. | [optional] 
**CustomerReference** | Pointer to **string** | A customer reference string for customer look ups. | [optional] 

## Methods

### NewCreateDocumentRequestOneOf1

`func NewCreateDocumentRequestOneOf1(file string, ) *CreateDocumentRequestOneOf1`

NewCreateDocumentRequestOneOf1 instantiates a new CreateDocumentRequestOneOf1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDocumentRequestOneOf1WithDefaults

`func NewCreateDocumentRequestOneOf1WithDefaults() *CreateDocumentRequestOneOf1`

NewCreateDocumentRequestOneOf1WithDefaults instantiates a new CreateDocumentRequestOneOf1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFile

`func (o *CreateDocumentRequestOneOf1) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *CreateDocumentRequestOneOf1) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *CreateDocumentRequestOneOf1) SetFile(v string)`

SetFile sets File field to given value.


### GetFilename

`func (o *CreateDocumentRequestOneOf1) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CreateDocumentRequestOneOf1) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CreateDocumentRequestOneOf1) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *CreateDocumentRequestOneOf1) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetCustomerReference

`func (o *CreateDocumentRequestOneOf1) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *CreateDocumentRequestOneOf1) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *CreateDocumentRequestOneOf1) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *CreateDocumentRequestOneOf1) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


