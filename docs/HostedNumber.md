# HostedNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] [readonly] 
**PhoneNumber** | Pointer to **string** | The messaging hosted phone number (+E.164 format) | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewHostedNumber

`func NewHostedNumber() *HostedNumber`

NewHostedNumber instantiates a new HostedNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHostedNumberWithDefaults

`func NewHostedNumberWithDefaults() *HostedNumber`

NewHostedNumberWithDefaults instantiates a new HostedNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *HostedNumber) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *HostedNumber) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *HostedNumber) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *HostedNumber) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetId

`func (o *HostedNumber) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *HostedNumber) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *HostedNumber) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *HostedNumber) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *HostedNumber) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *HostedNumber) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *HostedNumber) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *HostedNumber) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetStatus

`func (o *HostedNumber) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HostedNumber) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HostedNumber) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HostedNumber) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


