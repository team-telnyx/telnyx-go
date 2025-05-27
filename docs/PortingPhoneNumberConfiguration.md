# PortingPhoneNumberConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies this phone number configuration | [optional] 
**UserBundleId** | Pointer to **string** | Identifies the associated user bundle | [optional] 
**PortingPhoneNumberId** | Pointer to **string** | Identifies the associated porting phone number | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 

## Methods

### NewPortingPhoneNumberConfiguration

`func NewPortingPhoneNumberConfiguration() *PortingPhoneNumberConfiguration`

NewPortingPhoneNumberConfiguration instantiates a new PortingPhoneNumberConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortingPhoneNumberConfigurationWithDefaults

`func NewPortingPhoneNumberConfigurationWithDefaults() *PortingPhoneNumberConfiguration`

NewPortingPhoneNumberConfigurationWithDefaults instantiates a new PortingPhoneNumberConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortingPhoneNumberConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortingPhoneNumberConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortingPhoneNumberConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortingPhoneNumberConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUserBundleId

`func (o *PortingPhoneNumberConfiguration) GetUserBundleId() string`

GetUserBundleId returns the UserBundleId field if non-nil, zero value otherwise.

### GetUserBundleIdOk

`func (o *PortingPhoneNumberConfiguration) GetUserBundleIdOk() (*string, bool)`

GetUserBundleIdOk returns a tuple with the UserBundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserBundleId

`func (o *PortingPhoneNumberConfiguration) SetUserBundleId(v string)`

SetUserBundleId sets UserBundleId field to given value.

### HasUserBundleId

`func (o *PortingPhoneNumberConfiguration) HasUserBundleId() bool`

HasUserBundleId returns a boolean if a field has been set.

### GetPortingPhoneNumberId

`func (o *PortingPhoneNumberConfiguration) GetPortingPhoneNumberId() string`

GetPortingPhoneNumberId returns the PortingPhoneNumberId field if non-nil, zero value otherwise.

### GetPortingPhoneNumberIdOk

`func (o *PortingPhoneNumberConfiguration) GetPortingPhoneNumberIdOk() (*string, bool)`

GetPortingPhoneNumberIdOk returns a tuple with the PortingPhoneNumberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortingPhoneNumberId

`func (o *PortingPhoneNumberConfiguration) SetPortingPhoneNumberId(v string)`

SetPortingPhoneNumberId sets PortingPhoneNumberId field to given value.

### HasPortingPhoneNumberId

`func (o *PortingPhoneNumberConfiguration) HasPortingPhoneNumberId() bool`

HasPortingPhoneNumberId returns a boolean if a field has been set.

### GetRecordType

`func (o *PortingPhoneNumberConfiguration) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PortingPhoneNumberConfiguration) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PortingPhoneNumberConfiguration) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PortingPhoneNumberConfiguration) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PortingPhoneNumberConfiguration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortingPhoneNumberConfiguration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortingPhoneNumberConfiguration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PortingPhoneNumberConfiguration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PortingPhoneNumberConfiguration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PortingPhoneNumberConfiguration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PortingPhoneNumberConfiguration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PortingPhoneNumberConfiguration) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


