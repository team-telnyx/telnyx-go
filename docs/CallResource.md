# CallResource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountSid** | Pointer to **string** | The id of the account the resource belongs to. | [optional] 
**AnsweredBy** | Pointer to **string** | The value of the answering machine detection result, if this feature was enabled for the call. | [optional] 
**CallerName** | Pointer to **string** | Caller ID, if present. | [optional] 
**DateCreated** | Pointer to **string** | The timestamp of when the resource was created. | [optional] 
**DateUpdated** | Pointer to **string** | The timestamp of when the resource was last updated. | [optional] 
**Direction** | Pointer to **string** | The direction of this call. | [optional] 
**Duration** | Pointer to **string** | The duration of this call, given in seconds. | [optional] 
**EndTime** | Pointer to **string** | The end time of this call. | [optional] 
**From** | Pointer to **string** | The phone number or SIP address that made this call. | [optional] 
**FromFormatted** | Pointer to **string** | The from number formatted for display. | [optional] 
**Price** | Pointer to **string** | The price of this call, the currency is specified in the price_unit field. Only populated when the call cost feature is enabled for the account. | [optional] 
**PriceUnit** | Pointer to **string** | The unit in which the price is given. | [optional] 
**Sid** | Pointer to **string** | The identifier of this call. | [optional] 
**StartTime** | Pointer to **string** | The start time of this call. | [optional] 
**Status** | Pointer to **string** | The status of this call. | [optional] 
**To** | Pointer to **string** | The phone number or SIP address that received this call. | [optional] 
**ToFormatted** | Pointer to **string** | The to number formatted for display. | [optional] 
**Uri** | Pointer to **string** | The relative URI for this call. | [optional] 

## Methods

### NewCallResource

`func NewCallResource() *CallResource`

NewCallResource instantiates a new CallResource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallResourceWithDefaults

`func NewCallResourceWithDefaults() *CallResource`

NewCallResourceWithDefaults instantiates a new CallResource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountSid

`func (o *CallResource) GetAccountSid() string`

GetAccountSid returns the AccountSid field if non-nil, zero value otherwise.

### GetAccountSidOk

`func (o *CallResource) GetAccountSidOk() (*string, bool)`

GetAccountSidOk returns a tuple with the AccountSid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSid

`func (o *CallResource) SetAccountSid(v string)`

SetAccountSid sets AccountSid field to given value.

### HasAccountSid

`func (o *CallResource) HasAccountSid() bool`

HasAccountSid returns a boolean if a field has been set.

### GetAnsweredBy

`func (o *CallResource) GetAnsweredBy() string`

GetAnsweredBy returns the AnsweredBy field if non-nil, zero value otherwise.

### GetAnsweredByOk

`func (o *CallResource) GetAnsweredByOk() (*string, bool)`

GetAnsweredByOk returns a tuple with the AnsweredBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweredBy

`func (o *CallResource) SetAnsweredBy(v string)`

SetAnsweredBy sets AnsweredBy field to given value.

### HasAnsweredBy

`func (o *CallResource) HasAnsweredBy() bool`

HasAnsweredBy returns a boolean if a field has been set.

### GetCallerName

`func (o *CallResource) GetCallerName() string`

GetCallerName returns the CallerName field if non-nil, zero value otherwise.

### GetCallerNameOk

`func (o *CallResource) GetCallerNameOk() (*string, bool)`

GetCallerNameOk returns a tuple with the CallerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerName

`func (o *CallResource) SetCallerName(v string)`

SetCallerName sets CallerName field to given value.

### HasCallerName

`func (o *CallResource) HasCallerName() bool`

HasCallerName returns a boolean if a field has been set.

### GetDateCreated

`func (o *CallResource) GetDateCreated() string`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *CallResource) GetDateCreatedOk() (*string, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *CallResource) SetDateCreated(v string)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *CallResource) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetDateUpdated

`func (o *CallResource) GetDateUpdated() string`

GetDateUpdated returns the DateUpdated field if non-nil, zero value otherwise.

### GetDateUpdatedOk

`func (o *CallResource) GetDateUpdatedOk() (*string, bool)`

GetDateUpdatedOk returns a tuple with the DateUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateUpdated

`func (o *CallResource) SetDateUpdated(v string)`

SetDateUpdated sets DateUpdated field to given value.

### HasDateUpdated

`func (o *CallResource) HasDateUpdated() bool`

HasDateUpdated returns a boolean if a field has been set.

### GetDirection

`func (o *CallResource) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *CallResource) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *CallResource) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *CallResource) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetDuration

`func (o *CallResource) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *CallResource) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *CallResource) SetDuration(v string)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *CallResource) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetEndTime

`func (o *CallResource) GetEndTime() string`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *CallResource) GetEndTimeOk() (*string, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *CallResource) SetEndTime(v string)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *CallResource) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFrom

`func (o *CallResource) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *CallResource) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *CallResource) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *CallResource) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetFromFormatted

`func (o *CallResource) GetFromFormatted() string`

GetFromFormatted returns the FromFormatted field if non-nil, zero value otherwise.

### GetFromFormattedOk

`func (o *CallResource) GetFromFormattedOk() (*string, bool)`

GetFromFormattedOk returns a tuple with the FromFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromFormatted

`func (o *CallResource) SetFromFormatted(v string)`

SetFromFormatted sets FromFormatted field to given value.

### HasFromFormatted

`func (o *CallResource) HasFromFormatted() bool`

HasFromFormatted returns a boolean if a field has been set.

### GetPrice

`func (o *CallResource) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CallResource) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CallResource) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CallResource) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetPriceUnit

`func (o *CallResource) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *CallResource) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *CallResource) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *CallResource) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetSid

`func (o *CallResource) GetSid() string`

GetSid returns the Sid field if non-nil, zero value otherwise.

### GetSidOk

`func (o *CallResource) GetSidOk() (*string, bool)`

GetSidOk returns a tuple with the Sid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSid

`func (o *CallResource) SetSid(v string)`

SetSid sets Sid field to given value.

### HasSid

`func (o *CallResource) HasSid() bool`

HasSid returns a boolean if a field has been set.

### GetStartTime

`func (o *CallResource) GetStartTime() string`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *CallResource) GetStartTimeOk() (*string, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *CallResource) SetStartTime(v string)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *CallResource) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatus

`func (o *CallResource) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CallResource) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CallResource) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CallResource) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTo

`func (o *CallResource) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *CallResource) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *CallResource) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *CallResource) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetToFormatted

`func (o *CallResource) GetToFormatted() string`

GetToFormatted returns the ToFormatted field if non-nil, zero value otherwise.

### GetToFormattedOk

`func (o *CallResource) GetToFormattedOk() (*string, bool)`

GetToFormattedOk returns a tuple with the ToFormatted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToFormatted

`func (o *CallResource) SetToFormatted(v string)`

SetToFormatted sets ToFormatted field to given value.

### HasToFormatted

`func (o *CallResource) HasToFormatted() bool`

HasToFormatted returns a boolean if a field has been set.

### GetUri

`func (o *CallResource) GetUri() string`

GetUri returns the Uri field if non-nil, zero value otherwise.

### GetUriOk

`func (o *CallResource) GetUriOk() (*string, bool)`

GetUriOk returns a tuple with the Uri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUri

`func (o *CallResource) SetUri(v string)`

SetUri sets Uri field to given value.

### HasUri

`func (o *CallResource) HasUri() bool`

HasUri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


