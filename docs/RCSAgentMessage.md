# RCSAgentMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentMessage** | Pointer to [**RCSContentMessage**](RCSContentMessage.md) |  | [optional] 
**Event** | Pointer to [**RCSEvent**](RCSEvent.md) |  | [optional] 
**ExpireTime** | Pointer to **time.Time** | Timestamp in UTC of when this message is considered expired | [optional] 
**Ttl** | Pointer to **string** | Duration in seconds ending with &#39;s&#39; | [optional] 

## Methods

### NewRCSAgentMessage

`func NewRCSAgentMessage() *RCSAgentMessage`

NewRCSAgentMessage instantiates a new RCSAgentMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCSAgentMessageWithDefaults

`func NewRCSAgentMessageWithDefaults() *RCSAgentMessage`

NewRCSAgentMessageWithDefaults instantiates a new RCSAgentMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentMessage

`func (o *RCSAgentMessage) GetContentMessage() RCSContentMessage`

GetContentMessage returns the ContentMessage field if non-nil, zero value otherwise.

### GetContentMessageOk

`func (o *RCSAgentMessage) GetContentMessageOk() (*RCSContentMessage, bool)`

GetContentMessageOk returns a tuple with the ContentMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentMessage

`func (o *RCSAgentMessage) SetContentMessage(v RCSContentMessage)`

SetContentMessage sets ContentMessage field to given value.

### HasContentMessage

`func (o *RCSAgentMessage) HasContentMessage() bool`

HasContentMessage returns a boolean if a field has been set.

### GetEvent

`func (o *RCSAgentMessage) GetEvent() RCSEvent`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *RCSAgentMessage) GetEventOk() (*RCSEvent, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *RCSAgentMessage) SetEvent(v RCSEvent)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *RCSAgentMessage) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetExpireTime

`func (o *RCSAgentMessage) GetExpireTime() time.Time`

GetExpireTime returns the ExpireTime field if non-nil, zero value otherwise.

### GetExpireTimeOk

`func (o *RCSAgentMessage) GetExpireTimeOk() (*time.Time, bool)`

GetExpireTimeOk returns a tuple with the ExpireTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireTime

`func (o *RCSAgentMessage) SetExpireTime(v time.Time)`

SetExpireTime sets ExpireTime field to given value.

### HasExpireTime

`func (o *RCSAgentMessage) HasExpireTime() bool`

HasExpireTime returns a boolean if a field has been set.

### GetTtl

`func (o *RCSAgentMessage) GetTtl() string`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *RCSAgentMessage) GetTtlOk() (*string, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *RCSAgentMessage) SetTtl(v string)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *RCSAgentMessage) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


