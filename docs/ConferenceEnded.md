# ConferenceEnded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**EventType** | Pointer to **string** | The type of event being delivered. | [optional] 
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] 
**Payload** | Pointer to [**ConferenceEndedPayload**](ConferenceEndedPayload.md) |  | [optional] 

## Methods

### NewConferenceEnded

`func NewConferenceEnded() *ConferenceEnded`

NewConferenceEnded instantiates a new ConferenceEnded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceEndedWithDefaults

`func NewConferenceEndedWithDefaults() *ConferenceEnded`

NewConferenceEndedWithDefaults instantiates a new ConferenceEnded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *ConferenceEnded) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ConferenceEnded) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ConferenceEnded) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *ConferenceEnded) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetEventType

`func (o *ConferenceEnded) GetEventType() string`

GetEventType returns the EventType field if non-nil, zero value otherwise.

### GetEventTypeOk

`func (o *ConferenceEnded) GetEventTypeOk() (*string, bool)`

GetEventTypeOk returns a tuple with the EventType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventType

`func (o *ConferenceEnded) SetEventType(v string)`

SetEventType sets EventType field to given value.

### HasEventType

`func (o *ConferenceEnded) HasEventType() bool`

HasEventType returns a boolean if a field has been set.

### GetId

`func (o *ConferenceEnded) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConferenceEnded) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConferenceEnded) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConferenceEnded) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPayload

`func (o *ConferenceEnded) GetPayload() ConferenceEndedPayload`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *ConferenceEnded) GetPayloadOk() (*ConferenceEndedPayload, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *ConferenceEnded) SetPayload(v ConferenceEndedPayload)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *ConferenceEnded) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


