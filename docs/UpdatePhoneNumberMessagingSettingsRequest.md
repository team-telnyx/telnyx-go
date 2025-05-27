# UpdatePhoneNumberMessagingSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessagingProfileId** | Pointer to **string** | Configure the messaging profile this phone number is assigned to:  * Omit this field or set its value to &#x60;null&#x60; to keep the current value. * Set this field to &#x60;\&quot;\&quot;&#x60; to unassign the number from its messaging profile * Set this field to a quoted UUID of a messaging profile to assign this number to that messaging profile | [optional] 
**MessagingProduct** | Pointer to **string** | Configure the messaging product for this number:  * Omit this field or set its value to &#x60;null&#x60; to keep the current value. * Set this field to a quoted product ID to set this phone number to that product | [optional] 

## Methods

### NewUpdatePhoneNumberMessagingSettingsRequest

`func NewUpdatePhoneNumberMessagingSettingsRequest() *UpdatePhoneNumberMessagingSettingsRequest`

NewUpdatePhoneNumberMessagingSettingsRequest instantiates a new UpdatePhoneNumberMessagingSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePhoneNumberMessagingSettingsRequestWithDefaults

`func NewUpdatePhoneNumberMessagingSettingsRequestWithDefaults() *UpdatePhoneNumberMessagingSettingsRequest`

NewUpdatePhoneNumberMessagingSettingsRequestWithDefaults instantiates a new UpdatePhoneNumberMessagingSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessagingProfileId

`func (o *UpdatePhoneNumberMessagingSettingsRequest) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *UpdatePhoneNumberMessagingSettingsRequest) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *UpdatePhoneNumberMessagingSettingsRequest) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.

### HasMessagingProfileId

`func (o *UpdatePhoneNumberMessagingSettingsRequest) HasMessagingProfileId() bool`

HasMessagingProfileId returns a boolean if a field has been set.

### GetMessagingProduct

`func (o *UpdatePhoneNumberMessagingSettingsRequest) GetMessagingProduct() string`

GetMessagingProduct returns the MessagingProduct field if non-nil, zero value otherwise.

### GetMessagingProductOk

`func (o *UpdatePhoneNumberMessagingSettingsRequest) GetMessagingProductOk() (*string, bool)`

GetMessagingProductOk returns a tuple with the MessagingProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProduct

`func (o *UpdatePhoneNumberMessagingSettingsRequest) SetMessagingProduct(v string)`

SetMessagingProduct sets MessagingProduct field to given value.

### HasMessagingProduct

`func (o *UpdatePhoneNumberMessagingSettingsRequest) HasMessagingProduct() bool`

HasMessagingProduct returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


