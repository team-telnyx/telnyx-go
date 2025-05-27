# CreateConversationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** | Metadata associated with the conversation. | [optional] 

## Methods

### NewCreateConversationRequest

`func NewCreateConversationRequest() *CreateConversationRequest`

NewCreateConversationRequest instantiates a new CreateConversationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateConversationRequestWithDefaults

`func NewCreateConversationRequestWithDefaults() *CreateConversationRequest`

NewCreateConversationRequestWithDefaults instantiates a new CreateConversationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateConversationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateConversationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateConversationRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateConversationRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetMetadata

`func (o *CreateConversationRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CreateConversationRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CreateConversationRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CreateConversationRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


