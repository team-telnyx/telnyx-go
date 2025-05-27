# ConferenceDetailRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Conference id | [optional] 
**Name** | Pointer to **string** | Conference name | [optional] 
**UserId** | Pointer to **string** | User id | [optional] 
**StartedAt** | Pointer to **time.Time** | Conference start time | [optional] 
**EndedAt** | Pointer to **time.Time** | Conference end time | [optional] 
**ExpiresAt** | Pointer to **time.Time** | Conference expiry time | [optional] 
**Region** | Pointer to **string** | Region where the conference is hosted | [optional] 
**CallLegId** | Pointer to **string** | Telnyx UUID that identifies the conference call leg | [optional] 
**CallSessionId** | Pointer to **string** | Telnyx UUID that identifies with conference call session | [optional] 
**ConnectionId** | Pointer to **string** | Connection id | [optional] 
**CallSec** | Pointer to **int32** | Duration of the conference call in seconds | [optional] 
**ParticipantCount** | Pointer to **int32** | Number of participants that joined the conference call | [optional] 
**ParticipantCallSec** | Pointer to **int32** | Sum of the conference call duration for all participants in seconds | [optional] 
**IsTelnyxBillable** | Pointer to **bool** | Indicates whether Telnyx billing charges might be applicable | [optional] 
**RecordType** | **string** |  | [default to "conference_detail_record"]

## Methods

### NewConferenceDetailRecord

`func NewConferenceDetailRecord(recordType string, ) *ConferenceDetailRecord`

NewConferenceDetailRecord instantiates a new ConferenceDetailRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceDetailRecordWithDefaults

`func NewConferenceDetailRecordWithDefaults() *ConferenceDetailRecord`

NewConferenceDetailRecordWithDefaults instantiates a new ConferenceDetailRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConferenceDetailRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConferenceDetailRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConferenceDetailRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConferenceDetailRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ConferenceDetailRecord) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ConferenceDetailRecord) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ConferenceDetailRecord) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ConferenceDetailRecord) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUserId

`func (o *ConferenceDetailRecord) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ConferenceDetailRecord) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ConferenceDetailRecord) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ConferenceDetailRecord) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetStartedAt

`func (o *ConferenceDetailRecord) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ConferenceDetailRecord) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ConferenceDetailRecord) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ConferenceDetailRecord) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### GetEndedAt

`func (o *ConferenceDetailRecord) GetEndedAt() time.Time`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *ConferenceDetailRecord) GetEndedAtOk() (*time.Time, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *ConferenceDetailRecord) SetEndedAt(v time.Time)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *ConferenceDetailRecord) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ConferenceDetailRecord) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ConferenceDetailRecord) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ConferenceDetailRecord) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ConferenceDetailRecord) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetRegion

`func (o *ConferenceDetailRecord) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *ConferenceDetailRecord) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *ConferenceDetailRecord) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *ConferenceDetailRecord) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetCallLegId

`func (o *ConferenceDetailRecord) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *ConferenceDetailRecord) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *ConferenceDetailRecord) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *ConferenceDetailRecord) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *ConferenceDetailRecord) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *ConferenceDetailRecord) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *ConferenceDetailRecord) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *ConferenceDetailRecord) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetConnectionId

`func (o *ConferenceDetailRecord) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *ConferenceDetailRecord) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *ConferenceDetailRecord) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *ConferenceDetailRecord) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCallSec

`func (o *ConferenceDetailRecord) GetCallSec() int32`

GetCallSec returns the CallSec field if non-nil, zero value otherwise.

### GetCallSecOk

`func (o *ConferenceDetailRecord) GetCallSecOk() (*int32, bool)`

GetCallSecOk returns a tuple with the CallSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSec

`func (o *ConferenceDetailRecord) SetCallSec(v int32)`

SetCallSec sets CallSec field to given value.

### HasCallSec

`func (o *ConferenceDetailRecord) HasCallSec() bool`

HasCallSec returns a boolean if a field has been set.

### GetParticipantCount

`func (o *ConferenceDetailRecord) GetParticipantCount() int32`

GetParticipantCount returns the ParticipantCount field if non-nil, zero value otherwise.

### GetParticipantCountOk

`func (o *ConferenceDetailRecord) GetParticipantCountOk() (*int32, bool)`

GetParticipantCountOk returns a tuple with the ParticipantCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantCount

`func (o *ConferenceDetailRecord) SetParticipantCount(v int32)`

SetParticipantCount sets ParticipantCount field to given value.

### HasParticipantCount

`func (o *ConferenceDetailRecord) HasParticipantCount() bool`

HasParticipantCount returns a boolean if a field has been set.

### GetParticipantCallSec

`func (o *ConferenceDetailRecord) GetParticipantCallSec() int32`

GetParticipantCallSec returns the ParticipantCallSec field if non-nil, zero value otherwise.

### GetParticipantCallSecOk

`func (o *ConferenceDetailRecord) GetParticipantCallSecOk() (*int32, bool)`

GetParticipantCallSecOk returns a tuple with the ParticipantCallSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParticipantCallSec

`func (o *ConferenceDetailRecord) SetParticipantCallSec(v int32)`

SetParticipantCallSec sets ParticipantCallSec field to given value.

### HasParticipantCallSec

`func (o *ConferenceDetailRecord) HasParticipantCallSec() bool`

HasParticipantCallSec returns a boolean if a field has been set.

### GetIsTelnyxBillable

`func (o *ConferenceDetailRecord) GetIsTelnyxBillable() bool`

GetIsTelnyxBillable returns the IsTelnyxBillable field if non-nil, zero value otherwise.

### GetIsTelnyxBillableOk

`func (o *ConferenceDetailRecord) GetIsTelnyxBillableOk() (*bool, bool)`

GetIsTelnyxBillableOk returns a tuple with the IsTelnyxBillable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTelnyxBillable

`func (o *ConferenceDetailRecord) SetIsTelnyxBillable(v bool)`

SetIsTelnyxBillable sets IsTelnyxBillable field to given value.

### HasIsTelnyxBillable

`func (o *ConferenceDetailRecord) HasIsTelnyxBillable() bool`

HasIsTelnyxBillable returns a boolean if a field has been set.

### GetRecordType

`func (o *ConferenceDetailRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ConferenceDetailRecord) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ConferenceDetailRecord) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


