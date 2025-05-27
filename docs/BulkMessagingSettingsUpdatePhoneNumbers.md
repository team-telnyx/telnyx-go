# BulkMessagingSettingsUpdatePhoneNumbers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**OrderId** | Pointer to **string** | Order ID to verify bulk update status. | [optional] [readonly] 
**Success** | Pointer to **[]string** | Phoned numbers updated successfully. | [optional] 
**Pending** | Pointer to **[]string** | Phone numbers pending to be updated. | [optional] 
**Failed** | Pointer to **[]string** | Phone numbers that failed to update. | [optional] 

## Methods

### NewBulkMessagingSettingsUpdatePhoneNumbers

`func NewBulkMessagingSettingsUpdatePhoneNumbers() *BulkMessagingSettingsUpdatePhoneNumbers`

NewBulkMessagingSettingsUpdatePhoneNumbers instantiates a new BulkMessagingSettingsUpdatePhoneNumbers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkMessagingSettingsUpdatePhoneNumbersWithDefaults

`func NewBulkMessagingSettingsUpdatePhoneNumbersWithDefaults() *BulkMessagingSettingsUpdatePhoneNumbers`

NewBulkMessagingSettingsUpdatePhoneNumbersWithDefaults instantiates a new BulkMessagingSettingsUpdatePhoneNumbers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *BulkMessagingSettingsUpdatePhoneNumbers) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *BulkMessagingSettingsUpdatePhoneNumbers) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *BulkMessagingSettingsUpdatePhoneNumbers) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *BulkMessagingSettingsUpdatePhoneNumbers) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetOrderId

`func (o *BulkMessagingSettingsUpdatePhoneNumbers) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *BulkMessagingSettingsUpdatePhoneNumbers) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *BulkMessagingSettingsUpdatePhoneNumbers) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *BulkMessagingSettingsUpdatePhoneNumbers) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### GetSuccess

`func (o *BulkMessagingSettingsUpdatePhoneNumbers) GetSuccess() []string`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *BulkMessagingSettingsUpdatePhoneNumbers) GetSuccessOk() (*[]string, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *BulkMessagingSettingsUpdatePhoneNumbers) SetSuccess(v []string)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *BulkMessagingSettingsUpdatePhoneNumbers) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetPending

`func (o *BulkMessagingSettingsUpdatePhoneNumbers) GetPending() []string`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *BulkMessagingSettingsUpdatePhoneNumbers) GetPendingOk() (*[]string, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *BulkMessagingSettingsUpdatePhoneNumbers) SetPending(v []string)`

SetPending sets Pending field to given value.

### HasPending

`func (o *BulkMessagingSettingsUpdatePhoneNumbers) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetFailed

`func (o *BulkMessagingSettingsUpdatePhoneNumbers) GetFailed() []string`

GetFailed returns the Failed field if non-nil, zero value otherwise.

### GetFailedOk

`func (o *BulkMessagingSettingsUpdatePhoneNumbers) GetFailedOk() (*[]string, bool)`

GetFailedOk returns a tuple with the Failed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailed

`func (o *BulkMessagingSettingsUpdatePhoneNumbers) SetFailed(v []string)`

SetFailed sets Failed field to given value.

### HasFailed

`func (o *BulkMessagingSettingsUpdatePhoneNumbers) HasFailed() bool`

HasFailed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


