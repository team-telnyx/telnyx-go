# PortingPhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PortingOrderStatus** | Pointer to **string** | The current status of the porting order | [optional] 
**PhoneNumberType** | Pointer to **string** | The type of the phone number | [optional] 
**PhoneNumber** | Pointer to **string** | E164 formatted phone number | [optional] 
**PortingOrderId** | Pointer to **string** | Identifies the associated port request | [optional] 
**SupportKey** | Pointer to **string** | A key to reference this porting order when contacting Telnyx customer support | [optional] 
**ActivationStatus** | Pointer to [**PortingOrderActivationStatus**](PortingOrderActivationStatus.md) |  | [optional] 
**PortabilityStatus** | Pointer to [**PortabilityStatus**](PortabilityStatus.md) |  | [optional] 
**RequirementsStatus** | Pointer to **string** | The current status of the requirements in a INTL porting order | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 

## Methods

### NewPortingPhoneNumber

`func NewPortingPhoneNumber() *PortingPhoneNumber`

NewPortingPhoneNumber instantiates a new PortingPhoneNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortingPhoneNumberWithDefaults

`func NewPortingPhoneNumberWithDefaults() *PortingPhoneNumber`

NewPortingPhoneNumberWithDefaults instantiates a new PortingPhoneNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPortingOrderStatus

`func (o *PortingPhoneNumber) GetPortingOrderStatus() string`

GetPortingOrderStatus returns the PortingOrderStatus field if non-nil, zero value otherwise.

### GetPortingOrderStatusOk

`func (o *PortingPhoneNumber) GetPortingOrderStatusOk() (*string, bool)`

GetPortingOrderStatusOk returns a tuple with the PortingOrderStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortingOrderStatus

`func (o *PortingPhoneNumber) SetPortingOrderStatus(v string)`

SetPortingOrderStatus sets PortingOrderStatus field to given value.

### HasPortingOrderStatus

`func (o *PortingPhoneNumber) HasPortingOrderStatus() bool`

HasPortingOrderStatus returns a boolean if a field has been set.

### GetPhoneNumberType

`func (o *PortingPhoneNumber) GetPhoneNumberType() string`

GetPhoneNumberType returns the PhoneNumberType field if non-nil, zero value otherwise.

### GetPhoneNumberTypeOk

`func (o *PortingPhoneNumber) GetPhoneNumberTypeOk() (*string, bool)`

GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberType

`func (o *PortingPhoneNumber) SetPhoneNumberType(v string)`

SetPhoneNumberType sets PhoneNumberType field to given value.

### HasPhoneNumberType

`func (o *PortingPhoneNumber) HasPhoneNumberType() bool`

HasPhoneNumberType returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *PortingPhoneNumber) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PortingPhoneNumber) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PortingPhoneNumber) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PortingPhoneNumber) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetPortingOrderId

`func (o *PortingPhoneNumber) GetPortingOrderId() string`

GetPortingOrderId returns the PortingOrderId field if non-nil, zero value otherwise.

### GetPortingOrderIdOk

`func (o *PortingPhoneNumber) GetPortingOrderIdOk() (*string, bool)`

GetPortingOrderIdOk returns a tuple with the PortingOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortingOrderId

`func (o *PortingPhoneNumber) SetPortingOrderId(v string)`

SetPortingOrderId sets PortingOrderId field to given value.

### HasPortingOrderId

`func (o *PortingPhoneNumber) HasPortingOrderId() bool`

HasPortingOrderId returns a boolean if a field has been set.

### GetSupportKey

`func (o *PortingPhoneNumber) GetSupportKey() string`

GetSupportKey returns the SupportKey field if non-nil, zero value otherwise.

### GetSupportKeyOk

`func (o *PortingPhoneNumber) GetSupportKeyOk() (*string, bool)`

GetSupportKeyOk returns a tuple with the SupportKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportKey

`func (o *PortingPhoneNumber) SetSupportKey(v string)`

SetSupportKey sets SupportKey field to given value.

### HasSupportKey

`func (o *PortingPhoneNumber) HasSupportKey() bool`

HasSupportKey returns a boolean if a field has been set.

### GetActivationStatus

`func (o *PortingPhoneNumber) GetActivationStatus() PortingOrderActivationStatus`

GetActivationStatus returns the ActivationStatus field if non-nil, zero value otherwise.

### GetActivationStatusOk

`func (o *PortingPhoneNumber) GetActivationStatusOk() (*PortingOrderActivationStatus, bool)`

GetActivationStatusOk returns a tuple with the ActivationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationStatus

`func (o *PortingPhoneNumber) SetActivationStatus(v PortingOrderActivationStatus)`

SetActivationStatus sets ActivationStatus field to given value.

### HasActivationStatus

`func (o *PortingPhoneNumber) HasActivationStatus() bool`

HasActivationStatus returns a boolean if a field has been set.

### GetPortabilityStatus

`func (o *PortingPhoneNumber) GetPortabilityStatus() PortabilityStatus`

GetPortabilityStatus returns the PortabilityStatus field if non-nil, zero value otherwise.

### GetPortabilityStatusOk

`func (o *PortingPhoneNumber) GetPortabilityStatusOk() (*PortabilityStatus, bool)`

GetPortabilityStatusOk returns a tuple with the PortabilityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortabilityStatus

`func (o *PortingPhoneNumber) SetPortabilityStatus(v PortabilityStatus)`

SetPortabilityStatus sets PortabilityStatus field to given value.

### HasPortabilityStatus

`func (o *PortingPhoneNumber) HasPortabilityStatus() bool`

HasPortabilityStatus returns a boolean if a field has been set.

### GetRequirementsStatus

`func (o *PortingPhoneNumber) GetRequirementsStatus() string`

GetRequirementsStatus returns the RequirementsStatus field if non-nil, zero value otherwise.

### GetRequirementsStatusOk

`func (o *PortingPhoneNumber) GetRequirementsStatusOk() (*string, bool)`

GetRequirementsStatusOk returns a tuple with the RequirementsStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementsStatus

`func (o *PortingPhoneNumber) SetRequirementsStatus(v string)`

SetRequirementsStatus sets RequirementsStatus field to given value.

### HasRequirementsStatus

`func (o *PortingPhoneNumber) HasRequirementsStatus() bool`

HasRequirementsStatus returns a boolean if a field has been set.

### GetRecordType

`func (o *PortingPhoneNumber) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PortingPhoneNumber) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PortingPhoneNumber) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PortingPhoneNumber) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


