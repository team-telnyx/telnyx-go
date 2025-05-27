# CreateCustomerServiceRecordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PhoneNumber** | **string** | A valid US phone number in E164 format. | 
**WebhookUrl** | Pointer to **string** | Callback URL to receive webhook notifications. | [optional] 
**AdditionalData** | Pointer to [**CustomerServiceRecordAdditionalData**](CustomerServiceRecordAdditionalData.md) |  | [optional] 

## Methods

### NewCreateCustomerServiceRecordRequest

`func NewCreateCustomerServiceRecordRequest(phoneNumber string, ) *CreateCustomerServiceRecordRequest`

NewCreateCustomerServiceRecordRequest instantiates a new CreateCustomerServiceRecordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCustomerServiceRecordRequestWithDefaults

`func NewCreateCustomerServiceRecordRequestWithDefaults() *CreateCustomerServiceRecordRequest`

NewCreateCustomerServiceRecordRequestWithDefaults instantiates a new CreateCustomerServiceRecordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPhoneNumber

`func (o *CreateCustomerServiceRecordRequest) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *CreateCustomerServiceRecordRequest) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *CreateCustomerServiceRecordRequest) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.


### GetWebhookUrl

`func (o *CreateCustomerServiceRecordRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *CreateCustomerServiceRecordRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *CreateCustomerServiceRecordRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *CreateCustomerServiceRecordRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetAdditionalData

`func (o *CreateCustomerServiceRecordRequest) GetAdditionalData() CustomerServiceRecordAdditionalData`

GetAdditionalData returns the AdditionalData field if non-nil, zero value otherwise.

### GetAdditionalDataOk

`func (o *CreateCustomerServiceRecordRequest) GetAdditionalDataOk() (*CustomerServiceRecordAdditionalData, bool)`

GetAdditionalDataOk returns a tuple with the AdditionalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalData

`func (o *CreateCustomerServiceRecordRequest) SetAdditionalData(v CustomerServiceRecordAdditionalData)`

SetAdditionalData sets AdditionalData field to given value.

### HasAdditionalData

`func (o *CreateCustomerServiceRecordRequest) HasAdditionalData() bool`

HasAdditionalData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


