# TexmlCreateCallRecordingResponseBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** |  | [optional] 
**CallSid** | Pointer to **string** |  | [optional] 
**ConferenceSid** | Pointer to **NullableString** |  | [optional] 
**Channels** | Pointer to [**TwimlRecordingChannels**](TwimlRecordingChannels.md) |  | [optional] [default to _2]
**DateCreated** | Pointer to **string** |  | [optional] 
**DateUpdated** | Pointer to **string** |  | [optional] 
**StartTime** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **NullableString** | The price of this recording, the currency is specified in the price_unit field.  | [optional] 
**PriceUnit** | Pointer to **NullableString** | The unit in which the price is given. | [optional] 
**Duration** | Pointer to **NullableString** | The duration of this recording, given in seconds. | [optional] 
**Sid** | Pointer to **string** | Identifier of a resource. | [optional] 
**Source** | Pointer to [**RecordingSource**](RecordingSource.md) |  | [optional] 
**ErrorCode** | Pointer to **NullableString** |  | [optional] 
**Track** | Pointer to [**RecordingTrack**](RecordingTrack.md) |  | [optional] 
**Uri** | Pointer to **string** | The relative URI for this recording resource. | [optional] 

## Methods

### NewTexmlCreateCallRecordingResponseBody

`func NewTexmlCreateCallRecordingResponseBody() *TexmlCreateCallRecordingResponseBody`

NewTexmlCreateCallRecordingResponseBody instantiates a new TexmlCreateCallRecordingResponseBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTexmlCreateCallRecordingResponseBodyWithDefaults

`func NewTexmlCreateCallRecordingResponseBodyWithDefaults() *TexmlCreateCallRecordingResponseBody`

NewTexmlCreateCallRecordingResponseBodyWithDefaults instantiates a new TexmlCreateCallRecordingResponseBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *TexmlCreateCallRecordingResponseBody) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *TexmlCreateCallRecordingResponseBody) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *TexmlCreateCallRecordingResponseBody) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *TexmlCreateCallRecordingResponseBody) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### GetCallSid

`func (o *TexmlCreateCallRecordingResponseBody) GetCallSid() string`

GetCallSid returns the CallSid field if non-nil, zero value otherwise.

### GetCallSidOk

`func (o *TexmlCreateCallRecordingResponseBody) GetCallSidOk() (*string, bool)`

GetCallSidOk returns a tuple with the CallSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallSid

`func (o *TexmlCreateCallRecordingResponseBody) SetCallSid(v string)`

SetCallSid sets CallSid field to given value.

### HasCallSid

`func (o *TexmlCreateCallRecordingResponseBody) HasCallSid() bool`

HasCallSid returns a boolean if a field has been set.

### GetConferenceSid

`func (o *TexmlCreateCallRecordingResponseBody) GetConferenceSid() string`

GetConferenceSid returns the ConferenceSid field if non-nil, zero value otherwise.

### GetConferenceSidOk

`func (o *TexmlCreateCallRecordingResponseBody) GetConferenceSidOk() (*string, bool)`

GetConferenceSidOk returns a tuple with the ConferenceSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceSid

`func (o *TexmlCreateCallRecordingResponseBody) SetConferenceSid(v string)`

SetConferenceSid sets ConferenceSid field to given value.

### HasConferenceSid

`func (o *TexmlCreateCallRecordingResponseBody) HasConferenceSid() bool`

HasConferenceSid returns a boolean if a field has been set.

### SetConferenceSidNil

`func (o *TexmlCreateCallRecordingResponseBody) SetConferenceSidNil(b bool)`

 SetConferenceSidNil sets the value for ConferenceSid to be an explicit nil

### UnsetConferenceSid
`func (o *TexmlCreateCallRecordingResponseBody) UnsetConferenceSid()`

UnsetConferenceSid ensures that no value is present for ConferenceSid, not even an explicit nil
### GetChannels

`func (o *TexmlCreateCallRecordingResponseBody) GetChannels() TwimlRecordingChannels`

GetChannels returns the Channels field if non-nil, zero value otherwise.

### GetChannelsOk

`func (o *TexmlCreateCallRecordingResponseBody) GetChannelsOk() (*TwimlRecordingChannels, bool)`

GetChannelsOk returns a tuple with the Channels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannels

`func (o *TexmlCreateCallRecordingResponseBody) SetChannels(v TwimlRecordingChannels)`

SetChannels sets Channels field to given value.

### HasChannels

`func (o *TexmlCreateCallRecordingResponseBody) HasChannels() bool`

HasChannels returns a boolean if a field has been set.

### GetDateCreated

`func (o *TexmlCreateCallRecordingResponseBody) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *TexmlCreateCallRecordingResponseBody) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *TexmlCreateCallRecordingResponseBody) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *TexmlCreateCallRecordingResponseBody) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateUpdated

`func (o *TexmlCreateCallRecordingResponseBody) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *TexmlCreateCallRecordingResponseBody) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *TexmlCreateCallRecordingResponseBody) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *TexmlCreateCallRecordingResponseBody) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetStartTime

`func (o *TexmlCreateCallRecordingResponseBody) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *TexmlCreateCallRecordingResponseBody) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *TexmlCreateCallRecordingResponseBody) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *TexmlCreateCallRecordingResponseBody) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetPrice

`func (o *TexmlCreateCallRecordingResponseBody) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *TexmlCreateCallRecordingResponseBody) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *TexmlCreateCallRecordingResponseBody) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *TexmlCreateCallRecordingResponseBody) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *TexmlCreateCallRecordingResponseBody) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *TexmlCreateCallRecordingResponseBody) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetPriceUnit

`func (o *TexmlCreateCallRecordingResponseBody) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *TexmlCreateCallRecordingResponseBody) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *TexmlCreateCallRecordingResponseBody) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *TexmlCreateCallRecordingResponseBody) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### SetPriceUnitNil

`func (o *TexmlCreateCallRecordingResponseBody) SetPriceUnitNil(b bool)`

 SetPriceUnitNil sets the value for PriceUnit to be an explicit nil

### UnsetPriceUnit
`func (o *TexmlCreateCallRecordingResponseBody) UnsetPriceUnit()`

UnsetPriceUnit ensures that no value is present for PriceUnit, not even an explicit nil
### GetDuration

`func (o *TexmlCreateCallRecordingResponseBody) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *TexmlCreateCallRecordingResponseBody) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *TexmlCreateCallRecordingResponseBody) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *TexmlCreateCallRecordingResponseBody) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### SetDurationNil

`func (o *TexmlCreateCallRecordingResponseBody) SetDurationNil(b bool)`

 SetDurationNil sets the value for Duration to be an explicit nil

### UnsetDuration
`func (o *TexmlCreateCallRecordingResponseBody) UnsetDuration()`

UnsetDuration ensures that no value is present for Duration, not even an explicit nil
### GetSid

`func (o *TexmlCreateCallRecordingResponseBody) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *TexmlCreateCallRecordingResponseBody) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *TexmlCreateCallRecordingResponseBody) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *TexmlCreateCallRecordingResponseBody) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetSource

`func (o *TexmlCreateCallRecordingResponseBody) GetSource() RecordingSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TexmlCreateCallRecordingResponseBody) GetSourceOk() (*RecordingSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TexmlCreateCallRecordingResponseBody) SetSource(v RecordingSource)`

SetSource sets Source field to given value.

### HasSource

`func (o *TexmlCreateCallRecordingResponseBody) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetErrorCode

`func (o *TexmlCreateCallRecordingResponseBody) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *TexmlCreateCallRecordingResponseBody) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *TexmlCreateCallRecordingResponseBody) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *TexmlCreateCallRecordingResponseBody) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *TexmlCreateCallRecordingResponseBody) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *TexmlCreateCallRecordingResponseBody) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetTrack

`func (o *TexmlCreateCallRecordingResponseBody) GetTrack() RecordingTrack`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *TexmlCreateCallRecordingResponseBody) GetTrackOk() (*RecordingTrack, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *TexmlCreateCallRecordingResponseBody) SetTrack(v RecordingTrack)`

SetTrack sets Track field to given value.

### HasTrack

`func (o *TexmlCreateCallRecordingResponseBody) HasTrack() bool`

HasTrack returns a boolean if a field has been set.

### GetUri

`func (o *TexmlCreateCallRecordingResponseBody) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *TexmlCreateCallRecordingResponseBody) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *TexmlCreateCallRecordingResponseBody) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *TexmlCreateCallRecordingResponseBody) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


