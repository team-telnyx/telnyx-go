# CustomerServiceRecordStatusChangedEventPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies the customer service record. | [optional] 
**PhoneNumber** | Pointer to **string** | The phone number of the customer service record. | [optional] 
**Status** | Pointer to **string** | The status of the customer service record | [optional] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating the last time where the resource was updated. | [optional] 

## Methods

### NewCustomerServiceRecordStatusChangedEventPayload

`func NewCustomerServiceRecordStatusChangedEventPayload() *CustomerServiceRecordStatusChangedEventPayload`

NewCustomerServiceRecordStatusChangedEventPayload instantiates a new CustomerServiceRecordStatusChangedEventPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerServiceRecordStatusChangedEventPayloadWithDefaults

`func NewCustomerServiceRecordStatusChangedEventPayloadWithDefaults() *CustomerServiceRecordStatusChangedEventPayload`

NewCustomerServiceRecordStatusChangedEventPayloadWithDefaults instantiates a new CustomerServiceRecordStatusChangedEventPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerServiceRecordStatusChangedEventPayload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerServiceRecordStatusChangedEventPayload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerServiceRecordStatusChangedEventPayload) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerServiceRecordStatusChangedEventPayload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *CustomerServiceRecordStatusChangedEventPayload) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CustomerServiceRecordStatusChangedEventPayload) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CustomerServiceRecordStatusChangedEventPayload) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *CustomerServiceRecordStatusChangedEventPayload) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetStatus

`func (o *CustomerServiceRecordStatusChangedEventPayload) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CustomerServiceRecordStatusChangedEventPayload) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CustomerServiceRecordStatusChangedEventPayload) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CustomerServiceRecordStatusChangedEventPayload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CustomerServiceRecordStatusChangedEventPayload) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CustomerServiceRecordStatusChangedEventPayload) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CustomerServiceRecordStatusChangedEventPayload) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CustomerServiceRecordStatusChangedEventPayload) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


