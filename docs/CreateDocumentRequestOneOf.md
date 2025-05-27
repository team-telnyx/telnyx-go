# CreateDocumentRequestOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** | If the file is already hosted publicly, you can provide a URL and have the documents service fetch it for you. | 
**Filename** | Pointer to **string** | The filename of the document. | [optional] 
**CustomerReference** | Pointer to **string** | Optional reference string for customer tracking. | [optional] 

## Methods

### NewCreateDocumentRequestOneOf

`func NewCreateDocumentRequestOneOf(url string, ) *CreateDocumentRequestOneOf`

NewCreateDocumentRequestOneOf instantiates a new CreateDocumentRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateDocumentRequestOneOfWithDefaults

`func NewCreateDocumentRequestOneOfWithDefaults() *CreateDocumentRequestOneOf`

NewCreateDocumentRequestOneOfWithDefaults instantiates a new CreateDocumentRequestOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *CreateDocumentRequestOneOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateDocumentRequestOneOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateDocumentRequestOneOf) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetFilename

`func (o *CreateDocumentRequestOneOf) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *CreateDocumentRequestOneOf) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *CreateDocumentRequestOneOf) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *CreateDocumentRequestOneOf) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetCustomerReference

`func (o *CreateDocumentRequestOneOf) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *CreateDocumentRequestOneOf) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *CreateDocumentRequestOneOf) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *CreateDocumentRequestOneOf) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


