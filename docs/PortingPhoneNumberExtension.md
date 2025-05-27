# PortingPhoneNumberExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies this porting phone number extension. | [optional] [readonly] 
**PortingPhoneNumberId** | Pointer to **string** | Identifies the porting phone number associated with this porting phone number extension. | [optional] 
**ExtensionRange** | Pointer to [**PortingPhoneNumberExtensionExtensionRange**](PortingPhoneNumberExtensionExtensionRange.md) |  | [optional] 
**ActivationRanges** | Pointer to [**[]PortingPhoneNumberExtensionActivationRangesInner**](PortingPhoneNumberExtensionActivationRangesInner.md) | Specifies the activation ranges for this porting phone number extension. The activation range must be within the extension range and should not overlap with other activation ranges. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was last updated. | [optional] [readonly] 

## Methods

### NewPortingPhoneNumberExtension

`func NewPortingPhoneNumberExtension() *PortingPhoneNumberExtension`

NewPortingPhoneNumberExtension instantiates a new PortingPhoneNumberExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortingPhoneNumberExtensionWithDefaults

`func NewPortingPhoneNumberExtensionWithDefaults() *PortingPhoneNumberExtension`

NewPortingPhoneNumberExtensionWithDefaults instantiates a new PortingPhoneNumberExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortingPhoneNumberExtension) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortingPhoneNumberExtension) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortingPhoneNumberExtension) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortingPhoneNumberExtension) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPortingPhoneNumberId

`func (o *PortingPhoneNumberExtension) GetPortingPhoneNumberId() string`

GetPortingPhoneNumberId returns the PortingPhoneNumberId field if non-nil, zero value otherwise.

### GetPortingPhoneNumberIdOk

`func (o *PortingPhoneNumberExtension) GetPortingPhoneNumberIdOk() (*string, bool)`

GetPortingPhoneNumberIdOk returns a tuple with the PortingPhoneNumberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortingPhoneNumberId

`func (o *PortingPhoneNumberExtension) SetPortingPhoneNumberId(v string)`

SetPortingPhoneNumberId sets PortingPhoneNumberId field to given value.

### HasPortingPhoneNumberId

`func (o *PortingPhoneNumberExtension) HasPortingPhoneNumberId() bool`

HasPortingPhoneNumberId returns a boolean if a field has been set.

### GetExtensionRange

`func (o *PortingPhoneNumberExtension) GetExtensionRange() PortingPhoneNumberExtensionExtensionRange`

GetExtensionRange returns the ExtensionRange field if non-nil, zero value otherwise.

### GetExtensionRangeOk

`func (o *PortingPhoneNumberExtension) GetExtensionRangeOk() (*PortingPhoneNumberExtensionExtensionRange, bool)`

GetExtensionRangeOk returns a tuple with the ExtensionRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionRange

`func (o *PortingPhoneNumberExtension) SetExtensionRange(v PortingPhoneNumberExtensionExtensionRange)`

SetExtensionRange sets ExtensionRange field to given value.

### HasExtensionRange

`func (o *PortingPhoneNumberExtension) HasExtensionRange() bool`

HasExtensionRange returns a boolean if a field has been set.

### GetActivationRanges

`func (o *PortingPhoneNumberExtension) GetActivationRanges() []PortingPhoneNumberExtensionActivationRangesInner`

GetActivationRanges returns the ActivationRanges field if non-nil, zero value otherwise.

### GetActivationRangesOk

`func (o *PortingPhoneNumberExtension) GetActivationRangesOk() (*[]PortingPhoneNumberExtensionActivationRangesInner, bool)`

GetActivationRangesOk returns a tuple with the ActivationRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationRanges

`func (o *PortingPhoneNumberExtension) SetActivationRanges(v []PortingPhoneNumberExtensionActivationRangesInner)`

SetActivationRanges sets ActivationRanges field to given value.

### HasActivationRanges

`func (o *PortingPhoneNumberExtension) HasActivationRanges() bool`

HasActivationRanges returns a boolean if a field has been set.

### GetRecordType

`func (o *PortingPhoneNumberExtension) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PortingPhoneNumberExtension) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PortingPhoneNumberExtension) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PortingPhoneNumberExtension) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PortingPhoneNumberExtension) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortingPhoneNumberExtension) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortingPhoneNumberExtension) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PortingPhoneNumberExtension) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PortingPhoneNumberExtension) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PortingPhoneNumberExtension) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PortingPhoneNumberExtension) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PortingPhoneNumberExtension) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


