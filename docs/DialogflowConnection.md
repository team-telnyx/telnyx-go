# DialogflowConnection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** |  | [optional] 
**ConnectionId** | Pointer to **string** | Uniquely identifies a Telnyx application (Call Control). | [optional] 
**ConversationProfileId** | Pointer to **string** | The id of a configured conversation profile on your Dialogflow account. (If you use Dialogflow CX, this param is required) | [optional] 
**Environment** | Pointer to **string** | Which Dialogflow environment will be used. | [optional] 
**ServiceAccount** | Pointer to **string** | The JSON map to connect your Dialoglow account. | [optional] 

## Methods

### NewDialogflowConnection

`func NewDialogflowConnection() *DialogflowConnection`

NewDialogflowConnection instantiates a new DialogflowConnection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDialogflowConnectionWithDefaults

`func NewDialogflowConnectionWithDefaults() *DialogflowConnection`

NewDialogflowConnectionWithDefaults instantiates a new DialogflowConnection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *DialogflowConnection) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *DialogflowConnection) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *DialogflowConnection) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *DialogflowConnection) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetConnectionId

`func (o *DialogflowConnection) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *DialogflowConnection) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *DialogflowConnection) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *DialogflowConnection) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetConversationProfileId

`func (o *DialogflowConnection) GetConversationProfileId() string`

GetConversationProfileId returns the ConversationProfileId field if non-nil, zero value otherwise.

### GetConversationProfileIdOk

`func (o *DialogflowConnection) GetConversationProfileIdOk() (*string, bool)`

GetConversationProfileIdOk returns a tuple with the ConversationProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationProfileId

`func (o *DialogflowConnection) SetConversationProfileId(v string)`

SetConversationProfileId sets ConversationProfileId field to given value.

### HasConversationProfileId

`func (o *DialogflowConnection) HasConversationProfileId() bool`

HasConversationProfileId returns a boolean if a field has been set.

### GetEnvironment

`func (o *DialogflowConnection) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *DialogflowConnection) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *DialogflowConnection) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *DialogflowConnection) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetServiceAccount

`func (o *DialogflowConnection) GetServiceAccount() string`

GetServiceAccount returns the ServiceAccount field if non-nil, zero value otherwise.

### GetServiceAccountOk

`func (o *DialogflowConnection) GetServiceAccountOk() (*string, bool)`

GetServiceAccountOk returns a tuple with the ServiceAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceAccount

`func (o *DialogflowConnection) SetServiceAccount(v string)`

SetServiceAccount sets ServiceAccount field to given value.

### HasServiceAccount

`func (o *DialogflowConnection) HasServiceAccount() bool`

HasServiceAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


