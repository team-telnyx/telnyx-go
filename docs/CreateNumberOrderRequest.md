# CreateNumberOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumbers** | Pointer to [**[]CreateNumberOrderRequestPhoneNumbersInner**](CreateNumberOrderRequestPhoneNumbersInner.md) |  | [optional] 
**ConnectionId** | Pointer to **string** | Identifies the connection associated with this phone number. | [optional] 
**MessagingProfileId** | Pointer to **string** | Identifies the messaging profile associated with the phone number. | [optional] 
**BillingGroupId** | Pointer to **string** | Identifies the billing group associated with the phone number. | [optional] 
**CustomerReference** | Pointer to **string** | A customer reference string for customer look ups. | [optional] 

## Methods

### NewCreateNumberOrderRequest

`func NewCreateNumberOrderRequest() *CreateNumberOrderRequest`

NewCreateNumberOrderRequest instantiates a new CreateNumberOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateNumberOrderRequestWithDefaults

`func NewCreateNumberOrderRequestWithDefaults() *CreateNumberOrderRequest`

NewCreateNumberOrderRequestWithDefaults instantiates a new CreateNumberOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumbers

`func (o *CreateNumberOrderRequest) GetPhoneNumbers() []CreateNumberOrderRequestPhoneNumbersInner`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *CreateNumberOrderRequest) GetPhoneNumbersOk() (*[]CreateNumberOrderRequestPhoneNumbersInner, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *CreateNumberOrderRequest) SetPhoneNumbers(v []CreateNumberOrderRequestPhoneNumbersInner)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *CreateNumberOrderRequest) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetConnectionId

`func (o *CreateNumberOrderRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *CreateNumberOrderRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *CreateNumberOrderRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *CreateNumberOrderRequest) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetMessagingProfileId

`func (o *CreateNumberOrderRequest) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *CreateNumberOrderRequest) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *CreateNumberOrderRequest) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.

### HasMessagingProfileId

`func (o *CreateNumberOrderRequest) HasMessagingProfileId() bool`

HasMessagingProfileId returns a boolean if a field has been set.

### GetBillingGroupId

`func (o *CreateNumberOrderRequest) GetBillingGroupId() string`

GetBillingGroupId returns the BillingGroupId field if non-nil, zero value otherwise.

### GetBillingGroupIdOk

`func (o *CreateNumberOrderRequest) GetBillingGroupIdOk() (*string, bool)`

GetBillingGroupIdOk returns a tuple with the BillingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupId

`func (o *CreateNumberOrderRequest) SetBillingGroupId(v string)`

SetBillingGroupId sets BillingGroupId field to given value.

### HasBillingGroupId

`func (o *CreateNumberOrderRequest) HasBillingGroupId() bool`

HasBillingGroupId returns a boolean if a field has been set.

### GetCustomerReference

`func (o *CreateNumberOrderRequest) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *CreateNumberOrderRequest) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *CreateNumberOrderRequest) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *CreateNumberOrderRequest) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


