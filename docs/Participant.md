# Participant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | **string** |  | 
**Id** | **string** | Uniquely identifies the participant | 
**CallLegId** | **string** | Uniquely identifies the call leg associated with the participant | 
**CallControlId** | **string** | Call Control ID associated with the partiipant of the conference | 
**Conference** | [**ParticipantConference**](ParticipantConference.md) |  | 
**WhisperCallControlIds** | **[]string** | Array of unique call_control_ids the participant can whisper to.. | 
**CreatedAt** | **string** | ISO 8601 formatted date of when the participant was created | 
**UpdatedAt** | **string** | ISO 8601 formatted date of when the participant was last updated | 
**EndConferenceOnExit** | **bool** | Whether the conference will end and all remaining participants be hung up after the participant leaves the conference. | 
**SoftEndConferenceOnExit** | **bool** | Whether the conference will end after the participant leaves the conference. | 
**Status** | **string** | The status of the participant with respect to the lifecycle within the conference | 
**Muted** | **bool** | Whether the participant is muted. | 
**OnHold** | **bool** | Whether the participant is put on_hold. | 

## Methods

### NewParticipant

`func NewParticipant(recordType string, id string, callLegId string, callControlId string, conference ParticipantConference, whisperCallControlIds []string, createdAt string, updatedAt string, endConferenceOnExit bool, softEndConferenceOnExit bool, status string, muted bool, onHold bool, ) *Participant`

NewParticipant instantiates a new Participant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantWithDefaults

`func NewParticipantWithDefaults() *Participant`

NewParticipantWithDefaults instantiates a new Participant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *Participant) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *Participant) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *Participant) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetId

`func (o *Participant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Participant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Participant) SetId(v string)`

SetId sets Id field to given value.


### GetCallLegId

`func (o *Participant) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *Participant) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *Participant) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.


### GetCallControlId

`func (o *Participant) GetCallControlId() string`

GetCallControlId returns the CallControlId field if non-nil, zero value otherwise.

### GetCallControlIdOk

`func (o *Participant) GetCallControlIdOk() (*string, bool)`

GetCallControlIdOk returns a tuple with the CallControlId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallControlId

`func (o *Participant) SetCallControlId(v string)`

SetCallControlId sets CallControlId field to given value.


### GetConference

`func (o *Participant) GetConference() ParticipantConference`

GetConference returns the Conference field if non-nil, zero value otherwise.

### GetConferenceOk

`func (o *Participant) GetConferenceOk() (*ParticipantConference, bool)`

GetConferenceOk returns a tuple with the Conference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConference

`func (o *Participant) SetConference(v ParticipantConference)`

SetConference sets Conference field to given value.


### GetWhisperCallControlIds

`func (o *Participant) GetWhisperCallControlIds() []string`

GetWhisperCallControlIds returns the WhisperCallControlIds field if non-nil, zero value otherwise.

### GetWhisperCallControlIdsOk

`func (o *Participant) GetWhisperCallControlIdsOk() (*[]string, bool)`

GetWhisperCallControlIdsOk returns a tuple with the WhisperCallControlIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhisperCallControlIds

`func (o *Participant) SetWhisperCallControlIds(v []string)`

SetWhisperCallControlIds sets WhisperCallControlIds field to given value.


### GetCreatedAt

`func (o *Participant) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Participant) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Participant) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *Participant) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Participant) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Participant) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetEndConferenceOnExit

`func (o *Participant) GetEndConferenceOnExit() bool`

GetEndConferenceOnExit returns the EndConferenceOnExit field if non-nil, zero value otherwise.

### GetEndConferenceOnExitOk

`func (o *Participant) GetEndConferenceOnExitOk() (*bool, bool)`

GetEndConferenceOnExitOk returns a tuple with the EndConferenceOnExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndConferenceOnExit

`func (o *Participant) SetEndConferenceOnExit(v bool)`

SetEndConferenceOnExit sets EndConferenceOnExit field to given value.


### GetSoftEndConferenceOnExit

`func (o *Participant) GetSoftEndConferenceOnExit() bool`

GetSoftEndConferenceOnExit returns the SoftEndConferenceOnExit field if non-nil, zero value otherwise.

### GetSoftEndConferenceOnExitOk

`func (o *Participant) GetSoftEndConferenceOnExitOk() (*bool, bool)`

GetSoftEndConferenceOnExitOk returns a tuple with the SoftEndConferenceOnExit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftEndConferenceOnExit

`func (o *Participant) SetSoftEndConferenceOnExit(v bool)`

SetSoftEndConferenceOnExit sets SoftEndConferenceOnExit field to given value.


### GetStatus

`func (o *Participant) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Participant) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Participant) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMuted

`func (o *Participant) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *Participant) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *Participant) SetMuted(v bool)`

SetMuted sets Muted field to given value.


### GetOnHold

`func (o *Participant) GetOnHold() bool`

GetOnHold returns the OnHold field if non-nil, zero value otherwise.

### GetOnHoldOk

`func (o *Participant) GetOnHoldOk() (*bool, bool)`

GetOnHoldOk returns a tuple with the OnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHold

`func (o *Participant) SetOnHold(v bool)`

SetOnHold sets OnHold field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


