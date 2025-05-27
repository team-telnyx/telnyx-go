# CreateMessagingHostedNumberOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumbers** | Pointer to **[]string** | Phone numbers to be used for hosted messaging. | [optional] 
**MessagingProfileId** | Pointer to **string** | Automatically associate the number with this messaging profile ID when the order is complete. | [optional] 

## Methods

### NewCreateMessagingHostedNumberOrderRequest

`func NewCreateMessagingHostedNumberOrderRequest() *CreateMessagingHostedNumberOrderRequest`

NewCreateMessagingHostedNumberOrderRequest instantiates a new CreateMessagingHostedNumberOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMessagingHostedNumberOrderRequestWithDefaults

`func NewCreateMessagingHostedNumberOrderRequestWithDefaults() *CreateMessagingHostedNumberOrderRequest`

NewCreateMessagingHostedNumberOrderRequestWithDefaults instantiates a new CreateMessagingHostedNumberOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumbers

`func (o *CreateMessagingHostedNumberOrderRequest) GetPhoneNumbers() []string`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *CreateMessagingHostedNumberOrderRequest) GetPhoneNumbersOk() (*[]string, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *CreateMessagingHostedNumberOrderRequest) SetPhoneNumbers(v []string)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *CreateMessagingHostedNumberOrderRequest) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetMessagingProfileId

`func (o *CreateMessagingHostedNumberOrderRequest) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *CreateMessagingHostedNumberOrderRequest) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *CreateMessagingHostedNumberOrderRequest) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.

### HasMessagingProfileId

`func (o *CreateMessagingHostedNumberOrderRequest) HasMessagingProfileId() bool`

HasMessagingProfileId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


