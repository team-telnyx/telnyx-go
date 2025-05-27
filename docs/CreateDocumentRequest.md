# CreateDocumentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | If the file is already hosted publicly, you can provide a URL and have the documents service fetch it for you. | 
**Filename** | Pointer to **string** | The filename of the document. | [optional] 
**CustomerReference** | Pointer to **string** | A customer reference string for customer look ups. | [optional] 
**File** | **string** | The Base64 encoded contents of the file you are uploading. | 

## Methods

### NewCreateDocumentRequest

`func NewCreateDocumentRequest(url string, file string, ) *CreateDocumentRequest`

NewCreateDocumentRequest instantiates a new CreateDocumentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDocumentRequestWithDefaults

`func NewCreateDocumentRequestWithDefaults() *CreateDocumentRequest`

NewCreateDocumentRequestWithDefaults instantiates a new CreateDocumentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *CreateDocumentRequest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateDocumentRequest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateDocumentRequest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetFilename

`func (o *CreateDocumentRequest) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CreateDocumentRequest) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CreateDocumentRequest) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *CreateDocumentRequest) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetCustomerReference

`func (o *CreateDocumentRequest) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *CreateDocumentRequest) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *CreateDocumentRequest) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *CreateDocumentRequest) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetFile

`func (o *CreateDocumentRequest) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *CreateDocumentRequest) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *CreateDocumentRequest) SetFile(v string)`

SetFile sets File field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


