# PortingPhoneNumberBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies this porting phone number block. | [optional] [readonly] 
**CountryCode** | Pointer to **string** | Specifies the country code for this porting phone number block. It is a two-letter ISO 3166-1 alpha-2 country code. | [optional] 
**PhoneNumberType** | Pointer to **string** | Specifies the phone number type for this porting phone number block. | [optional] 
**PhoneNumberRange** | Pointer to [**PortingPhoneNumberBlockPhoneNumberRange**](PortingPhoneNumberBlockPhoneNumberRange.md) |  | [optional] 
**ActivationRanges** | Pointer to [**[]PortingPhoneNumberBlockActivationRangesInner**](PortingPhoneNumberBlockActivationRangesInner.md) | Specifies the activation ranges for this porting phone number block. The activation range must be within the phone number range and should not overlap with other activation ranges. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was last updated. | [optional] [readonly] 

## Methods

### NewPortingPhoneNumberBlock

`func NewPortingPhoneNumberBlock() *PortingPhoneNumberBlock`

NewPortingPhoneNumberBlock instantiates a new PortingPhoneNumberBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortingPhoneNumberBlockWithDefaults

`func NewPortingPhoneNumberBlockWithDefaults() *PortingPhoneNumberBlock`

NewPortingPhoneNumberBlockWithDefaults instantiates a new PortingPhoneNumberBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortingPhoneNumberBlock) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortingPhoneNumberBlock) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortingPhoneNumberBlock) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortingPhoneNumberBlock) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCountryCode

`func (o *PortingPhoneNumberBlock) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PortingPhoneNumberBlock) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PortingPhoneNumberBlock) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PortingPhoneNumberBlock) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetPhoneNumberType

`func (o *PortingPhoneNumberBlock) GetPhoneNumberType() string`

GetPhoneNumberType returns the PhoneNumberType field if non-nil, zero value otherwise.

### GetPhoneNumberTypeOk

`func (o *PortingPhoneNumberBlock) GetPhoneNumberTypeOk() (*string, bool)`

GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberType

`func (o *PortingPhoneNumberBlock) SetPhoneNumberType(v string)`

SetPhoneNumberType sets PhoneNumberType field to given value.

### HasPhoneNumberType

`func (o *PortingPhoneNumberBlock) HasPhoneNumberType() bool`

HasPhoneNumberType returns a boolean if a field has been set.

### GetPhoneNumberRange

`func (o *PortingPhoneNumberBlock) GetPhoneNumberRange() PortingPhoneNumberBlockPhoneNumberRange`

GetPhoneNumberRange returns the PhoneNumberRange field if non-nil, zero value otherwise.

### GetPhoneNumberRangeOk

`func (o *PortingPhoneNumberBlock) GetPhoneNumberRangeOk() (*PortingPhoneNumberBlockPhoneNumberRange, bool)`

GetPhoneNumberRangeOk returns a tuple with the PhoneNumberRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberRange

`func (o *PortingPhoneNumberBlock) SetPhoneNumberRange(v PortingPhoneNumberBlockPhoneNumberRange)`

SetPhoneNumberRange sets PhoneNumberRange field to given value.

### HasPhoneNumberRange

`func (o *PortingPhoneNumberBlock) HasPhoneNumberRange() bool`

HasPhoneNumberRange returns a boolean if a field has been set.

### GetActivationRanges

`func (o *PortingPhoneNumberBlock) GetActivationRanges() []PortingPhoneNumberBlockActivationRangesInner`

GetActivationRanges returns the ActivationRanges field if non-nil, zero value otherwise.

### GetActivationRangesOk

`func (o *PortingPhoneNumberBlock) GetActivationRangesOk() (*[]PortingPhoneNumberBlockActivationRangesInner, bool)`

GetActivationRangesOk returns a tuple with the ActivationRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationRanges

`func (o *PortingPhoneNumberBlock) SetActivationRanges(v []PortingPhoneNumberBlockActivationRangesInner)`

SetActivationRanges sets ActivationRanges field to given value.

### HasActivationRanges

`func (o *PortingPhoneNumberBlock) HasActivationRanges() bool`

HasActivationRanges returns a boolean if a field has been set.

### GetRecordType

`func (o *PortingPhoneNumberBlock) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PortingPhoneNumberBlock) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PortingPhoneNumberBlock) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PortingPhoneNumberBlock) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PortingPhoneNumberBlock) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortingPhoneNumberBlock) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortingPhoneNumberBlock) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PortingPhoneNumberBlock) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PortingPhoneNumberBlock) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PortingPhoneNumberBlock) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PortingPhoneNumberBlock) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PortingPhoneNumberBlock) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


