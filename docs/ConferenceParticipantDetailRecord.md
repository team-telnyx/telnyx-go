# ConferenceParticipantDetailRecord

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Participant id | [optional] 
**UserId** | Pointer to **string** | User id | [optional] 
**ConferenceId** | Pointer to **string** | Conference id | [optional] 
**JoinedAt** | Pointer to **time.Time** | Participant join time | [optional] 
**LeftAt** | Pointer to **time.Time** | Participant leave time | [optional] 
**DestinationNumber** | Pointer to **string** | Number called by the participant to join the conference | [optional] 
**OriginatingNumber** | Pointer to **string** | Participant origin number used in the conference call | [optional] 
**CallLegId** | Pointer to **string** | Telnyx UUID that identifies the conference call leg | [optional] 
**CallSessionId** | Pointer to **string** | Telnyx UUID that identifies with conference call session | [optional] 
**CallSec** | Pointer to **int32** | Duration of the conference call in seconds | [optional] 
**BilledSec** | Pointer to **int32** | Duration of the conference call for billing purposes | [optional] 
**IsTelnyxBillable** | Pointer to **bool** | Indicates whether Telnyx billing charges might be applicable | [optional] 
**Rate** | Pointer to **string** | Currency amount per billing unit used to calculate the Telnyx billing cost | [optional] 
**RateMeasuredIn** | Pointer to **string** | Billing unit used to calculate the Telnyx billing cost | [optional] 
**Cost** | Pointer to **string** | Currency amount for Telnyx billing cost | [optional] 
**Currency** | Pointer to **string** | Telnyx account currency used to describe monetary values, including billing cost | [optional] 
**RecordType** | **string** |  | [default to "conference_participant_detail_record"]

## Methods

### NewConferenceParticipantDetailRecord

`func NewConferenceParticipantDetailRecord(recordType string, ) *ConferenceParticipantDetailRecord`

NewConferenceParticipantDetailRecord instantiates a new ConferenceParticipantDetailRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceParticipantDetailRecordWithDefaults

`func NewConferenceParticipantDetailRecordWithDefaults() *ConferenceParticipantDetailRecord`

NewConferenceParticipantDetailRecordWithDefaults instantiates a new ConferenceParticipantDetailRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConferenceParticipantDetailRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConferenceParticipantDetailRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConferenceParticipantDetailRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ConferenceParticipantDetailRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserId

`func (o *ConferenceParticipantDetailRecord) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ConferenceParticipantDetailRecord) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ConferenceParticipantDetailRecord) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ConferenceParticipantDetailRecord) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetConferenceId

`func (o *ConferenceParticipantDetailRecord) GetConferenceId() string`

GetConferenceId returns the ConferenceId field if non-nil, zero value otherwise.

### GetConferenceIdOk

`func (o *ConferenceParticipantDetailRecord) GetConferenceIdOk() (*string, bool)`

GetConferenceIdOk returns a tuple with the ConferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceId

`func (o *ConferenceParticipantDetailRecord) SetConferenceId(v string)`

SetConferenceId sets ConferenceId field to given value.

### HasConferenceId

`func (o *ConferenceParticipantDetailRecord) HasConferenceId() bool`

HasConferenceId returns a boolean if a field has been set.

### GetJoinedAt

`func (o *ConferenceParticipantDetailRecord) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *ConferenceParticipantDetailRecord) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *ConferenceParticipantDetailRecord) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.

### HasJoinedAt

`func (o *ConferenceParticipantDetailRecord) HasJoinedAt() bool`

HasJoinedAt returns a boolean if a field has been set.

### GetLeftAt

`func (o *ConferenceParticipantDetailRecord) GetLeftAt() time.Time`

GetLeftAt returns the LeftAt field if non-nil, zero value otherwise.

### GetLeftAtOk

`func (o *ConferenceParticipantDetailRecord) GetLeftAtOk() (*time.Time, bool)`

GetLeftAtOk returns a tuple with the LeftAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeftAt

`func (o *ConferenceParticipantDetailRecord) SetLeftAt(v time.Time)`

SetLeftAt sets LeftAt field to given value.

### HasLeftAt

`func (o *ConferenceParticipantDetailRecord) HasLeftAt() bool`

HasLeftAt returns a boolean if a field has been set.

### GetDestinationNumber

`func (o *ConferenceParticipantDetailRecord) GetDestinationNumber() string`

GetDestinationNumber returns the DestinationNumber field if non-nil, zero value otherwise.

### GetDestinationNumberOk

`func (o *ConferenceParticipantDetailRecord) GetDestinationNumberOk() (*string, bool)`

GetDestinationNumberOk returns a tuple with the DestinationNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationNumber

`func (o *ConferenceParticipantDetailRecord) SetDestinationNumber(v string)`

SetDestinationNumber sets DestinationNumber field to given value.

### HasDestinationNumber

`func (o *ConferenceParticipantDetailRecord) HasDestinationNumber() bool`

HasDestinationNumber returns a boolean if a field has been set.

### GetOriginatingNumber

`func (o *ConferenceParticipantDetailRecord) GetOriginatingNumber() string`

GetOriginatingNumber returns the OriginatingNumber field if non-nil, zero value otherwise.

### GetOriginatingNumberOk

`func (o *ConferenceParticipantDetailRecord) GetOriginatingNumberOk() (*string, bool)`

GetOriginatingNumberOk returns a tuple with the OriginatingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginatingNumber

`func (o *ConferenceParticipantDetailRecord) SetOriginatingNumber(v string)`

SetOriginatingNumber sets OriginatingNumber field to given value.

### HasOriginatingNumber

`func (o *ConferenceParticipantDetailRecord) HasOriginatingNumber() bool`

HasOriginatingNumber returns a boolean if a field has been set.

### GetCallLegId

`func (o *ConferenceParticipantDetailRecord) GetCallLegId() string`

GetCallLegId returns the CallLegId field if non-nil, zero value otherwise.

### GetCallLegIdOk

`func (o *ConferenceParticipantDetailRecord) GetCallLegIdOk() (*string, bool)`

GetCallLegIdOk returns a tuple with the CallLegId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallLegId

`func (o *ConferenceParticipantDetailRecord) SetCallLegId(v string)`

SetCallLegId sets CallLegId field to given value.

### HasCallLegId

`func (o *ConferenceParticipantDetailRecord) HasCallLegId() bool`

HasCallLegId returns a boolean if a field has been set.

### GetCallSessionId

`func (o *ConferenceParticipantDetailRecord) GetCallSessionId() string`

GetCallSessionId returns the CallSessionId field if non-nil, zero value otherwise.

### GetCallSessionIdOk

`func (o *ConferenceParticipantDetailRecord) GetCallSessionIdOk() (*string, bool)`

GetCallSessionIdOk returns a tuple with the CallSessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSessionId

`func (o *ConferenceParticipantDetailRecord) SetCallSessionId(v string)`

SetCallSessionId sets CallSessionId field to given value.

### HasCallSessionId

`func (o *ConferenceParticipantDetailRecord) HasCallSessionId() bool`

HasCallSessionId returns a boolean if a field has been set.

### GetCallSec

`func (o *ConferenceParticipantDetailRecord) GetCallSec() int32`

GetCallSec returns the CallSec field if non-nil, zero value otherwise.

### GetCallSecOk

`func (o *ConferenceParticipantDetailRecord) GetCallSecOk() (*int32, bool)`

GetCallSecOk returns a tuple with the CallSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSec

`func (o *ConferenceParticipantDetailRecord) SetCallSec(v int32)`

SetCallSec sets CallSec field to given value.

### HasCallSec

`func (o *ConferenceParticipantDetailRecord) HasCallSec() bool`

HasCallSec returns a boolean if a field has been set.

### GetBilledSec

`func (o *ConferenceParticipantDetailRecord) GetBilledSec() int32`

GetBilledSec returns the BilledSec field if non-nil, zero value otherwise.

### GetBilledSecOk

`func (o *ConferenceParticipantDetailRecord) GetBilledSecOk() (*int32, bool)`

GetBilledSecOk returns a tuple with the BilledSec field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBilledSec

`func (o *ConferenceParticipantDetailRecord) SetBilledSec(v int32)`

SetBilledSec sets BilledSec field to given value.

### HasBilledSec

`func (o *ConferenceParticipantDetailRecord) HasBilledSec() bool`

HasBilledSec returns a boolean if a field has been set.

### GetIsTelnyxBillable

`func (o *ConferenceParticipantDetailRecord) GetIsTelnyxBillable() bool`

GetIsTelnyxBillable returns the IsTelnyxBillable field if non-nil, zero value otherwise.

### GetIsTelnyxBillableOk

`func (o *ConferenceParticipantDetailRecord) GetIsTelnyxBillableOk() (*bool, bool)`

GetIsTelnyxBillableOk returns a tuple with the IsTelnyxBillable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsTelnyxBillable

`func (o *ConferenceParticipantDetailRecord) SetIsTelnyxBillable(v bool)`

SetIsTelnyxBillable sets IsTelnyxBillable field to given value.

### HasIsTelnyxBillable

`func (o *ConferenceParticipantDetailRecord) HasIsTelnyxBillable() bool`

HasIsTelnyxBillable returns a boolean if a field has been set.

### GetRate

`func (o *ConferenceParticipantDetailRecord) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *ConferenceParticipantDetailRecord) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *ConferenceParticipantDetailRecord) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *ConferenceParticipantDetailRecord) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetRateMeasuredIn

`func (o *ConferenceParticipantDetailRecord) GetRateMeasuredIn() string`

GetRateMeasuredIn returns the RateMeasuredIn field if non-nil, zero value otherwise.

### GetRateMeasuredInOk

`func (o *ConferenceParticipantDetailRecord) GetRateMeasuredInOk() (*string, bool)`

GetRateMeasuredInOk returns a tuple with the RateMeasuredIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateMeasuredIn

`func (o *ConferenceParticipantDetailRecord) SetRateMeasuredIn(v string)`

SetRateMeasuredIn sets RateMeasuredIn field to given value.

### HasRateMeasuredIn

`func (o *ConferenceParticipantDetailRecord) HasRateMeasuredIn() bool`

HasRateMeasuredIn returns a boolean if a field has been set.

### GetCost

`func (o *ConferenceParticipantDetailRecord) GetCost() string`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ConferenceParticipantDetailRecord) GetCostOk() (*string, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ConferenceParticipantDetailRecord) SetCost(v string)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ConferenceParticipantDetailRecord) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCurrency

`func (o *ConferenceParticipantDetailRecord) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ConferenceParticipantDetailRecord) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ConferenceParticipantDetailRecord) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ConferenceParticipantDetailRecord) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetRecordType

`func (o *ConferenceParticipantDetailRecord) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ConferenceParticipantDetailRecord) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ConferenceParticipantDetailRecord) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


