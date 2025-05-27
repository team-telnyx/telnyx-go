# GcbPhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | **string** |  | 
**ChannelZoneId** | **string** |  | 
**Id** | Pointer to **string** |  | [optional] 
**PhoneNumber** | **string** |  | 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date of when the phone number was created | [optional] 

## Methods

### NewGcbPhoneNumber

`func NewGcbPhoneNumber(recordType string, channelZoneId string, phoneNumber string, ) *GcbPhoneNumber`

NewGcbPhoneNumber instantiates a new GcbPhoneNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGcbPhoneNumberWithDefaults

`func NewGcbPhoneNumberWithDefaults() *GcbPhoneNumber`

NewGcbPhoneNumberWithDefaults instantiates a new GcbPhoneNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *GcbPhoneNumber) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *GcbPhoneNumber) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *GcbPhoneNumber) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.


### GetChannelZoneId

`func (o *GcbPhoneNumber) GetChannelZoneId() string`

GetChannelZoneId returns the ChannelZoneId field if non-nil, zero value otherwise.

### GetChannelZoneIdOk

`func (o *GcbPhoneNumber) GetChannelZoneIdOk() (*string, bool)`

GetChannelZoneIdOk returns a tuple with the ChannelZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelZoneId

`func (o *GcbPhoneNumber) SetChannelZoneId(v string)`

SetChannelZoneId sets ChannelZoneId field to given value.


### GetId

`func (o *GcbPhoneNumber) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GcbPhoneNumber) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GcbPhoneNumber) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GcbPhoneNumber) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *GcbPhoneNumber) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *GcbPhoneNumber) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *GcbPhoneNumber) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetCreatedAt

`func (o *GcbPhoneNumber) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GcbPhoneNumber) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GcbPhoneNumber) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GcbPhoneNumber) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


