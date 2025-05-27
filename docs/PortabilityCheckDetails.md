# PortabilityCheckDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**FastPortable** | Pointer to **bool** | Indicates whether this phone number is FastPort eligible | [optional] [readonly] 
**NotPortableReason** | Pointer to **string** | If this phone number is not portable, explains why. Empty string if the number is portable. | [optional] [readonly] 
**PhoneNumber** | Pointer to **string** | The +E.164 formatted phone number this result is about | [optional] [readonly] 
**Portable** | Pointer to **bool** | Indicates whether this phone number is portable | [optional] [readonly] 

## Methods

### NewPortabilityCheckDetails

`func NewPortabilityCheckDetails() *PortabilityCheckDetails`

NewPortabilityCheckDetails instantiates a new PortabilityCheckDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortabilityCheckDetailsWithDefaults

`func NewPortabilityCheckDetailsWithDefaults() *PortabilityCheckDetails`

NewPortabilityCheckDetailsWithDefaults instantiates a new PortabilityCheckDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *PortabilityCheckDetails) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PortabilityCheckDetails) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PortabilityCheckDetails) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PortabilityCheckDetails) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetFastPortable

`func (o *PortabilityCheckDetails) GetFastPortable() bool`

GetFastPortable returns the FastPortable field if non-nil, zero value otherwise.

### GetFastPortableOk

`func (o *PortabilityCheckDetails) GetFastPortableOk() (*bool, bool)`

GetFastPortableOk returns a tuple with the FastPortable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFastPortable

`func (o *PortabilityCheckDetails) SetFastPortable(v bool)`

SetFastPortable sets FastPortable field to given value.

### HasFastPortable

`func (o *PortabilityCheckDetails) HasFastPortable() bool`

HasFastPortable returns a boolean if a field has been set.

### GetNotPortableReason

`func (o *PortabilityCheckDetails) GetNotPortableReason() string`

GetNotPortableReason returns the NotPortableReason field if non-nil, zero value otherwise.

### GetNotPortableReasonOk

`func (o *PortabilityCheckDetails) GetNotPortableReasonOk() (*string, bool)`

GetNotPortableReasonOk returns a tuple with the NotPortableReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotPortableReason

`func (o *PortabilityCheckDetails) SetNotPortableReason(v string)`

SetNotPortableReason sets NotPortableReason field to given value.

### HasNotPortableReason

`func (o *PortabilityCheckDetails) HasNotPortableReason() bool`

HasNotPortableReason returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *PortabilityCheckDetails) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PortabilityCheckDetails) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PortabilityCheckDetails) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PortabilityCheckDetails) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPortable

`func (o *PortabilityCheckDetails) GetPortable() bool`

GetPortable returns the Portable field if non-nil, zero value otherwise.

### GetPortableOk

`func (o *PortabilityCheckDetails) GetPortableOk() (*bool, bool)`

GetPortableOk returns a tuple with the Portable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortable

`func (o *PortabilityCheckDetails) SetPortable(v bool)`

SetPortable sets Portable field to given value.

### HasPortable

`func (o *PortabilityCheckDetails) HasPortable() bool`

HasPortable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


