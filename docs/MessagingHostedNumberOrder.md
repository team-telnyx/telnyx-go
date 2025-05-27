# MessagingHostedNumberOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**Id** | Pointer to **string** | Resource unique identifier. | [optional] [readonly] 
**MessagingProfileId** | Pointer to **NullableString** | Automatically associate the number with this messaging profile ID when the order is complete. | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**PhoneNumbers** | Pointer to [**[]HostedNumber**](HostedNumber.md) |  | [optional] 

## Methods

### NewMessagingHostedNumberOrder

`func NewMessagingHostedNumberOrder() *MessagingHostedNumberOrder`

NewMessagingHostedNumberOrder instantiates a new MessagingHostedNumberOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagingHostedNumberOrderWithDefaults

`func NewMessagingHostedNumberOrderWithDefaults() *MessagingHostedNumberOrder`

NewMessagingHostedNumberOrderWithDefaults instantiates a new MessagingHostedNumberOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *MessagingHostedNumberOrder) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *MessagingHostedNumberOrder) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *MessagingHostedNumberOrder) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *MessagingHostedNumberOrder) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetId

`func (o *MessagingHostedNumberOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MessagingHostedNumberOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MessagingHostedNumberOrder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MessagingHostedNumberOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMessagingProfileId

`func (o *MessagingHostedNumberOrder) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *MessagingHostedNumberOrder) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *MessagingHostedNumberOrder) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.

### HasMessagingProfileId

`func (o *MessagingHostedNumberOrder) HasMessagingProfileId() bool`

HasMessagingProfileId returns a boolean if a field has been set.

### SetMessagingProfileIdNil

`func (o *MessagingHostedNumberOrder) SetMessagingProfileIdNil(b bool)`

 SetMessagingProfileIdNil sets the value for MessagingProfileId to be an explicit nil

### UnsetMessagingProfileId
`func (o *MessagingHostedNumberOrder) UnsetMessagingProfileId()`

UnsetMessagingProfileId ensures that no value is present for MessagingProfileId, not even an explicit nil
### GetStatus

`func (o *MessagingHostedNumberOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MessagingHostedNumberOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MessagingHostedNumberOrder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MessagingHostedNumberOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *MessagingHostedNumberOrder) GetPhoneNumbers() []HostedNumber`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *MessagingHostedNumberOrder) GetPhoneNumbersOk() (*[]HostedNumber, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *MessagingHostedNumberOrder) SetPhoneNumbers(v []HostedNumber)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *MessagingHostedNumberOrder) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


