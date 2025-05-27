# MessagingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultMessagingProfileId** | Pointer to **string** | Default Messaging Profile used for messaging exchanges with your assistant. This will be created automatically on assistant creation. | [optional] 
**DeliveryStatusWebhookUrl** | Pointer to **string** | The URL where webhooks related to delivery statused for assistant messages will be sent. | [optional] 

## Methods

### NewMessagingSettings

`func NewMessagingSettings() *MessagingSettings`

NewMessagingSettings instantiates a new MessagingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessagingSettingsWithDefaults

`func NewMessagingSettingsWithDefaults() *MessagingSettings`

NewMessagingSettingsWithDefaults instantiates a new MessagingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultMessagingProfileId

`func (o *MessagingSettings) GetDefaultMessagingProfileId() string`

GetDefaultMessagingProfileId returns the DefaultMessagingProfileId field if non-nil, zero value otherwise.

### GetDefaultMessagingProfileIdOk

`func (o *MessagingSettings) GetDefaultMessagingProfileIdOk() (*string, bool)`

GetDefaultMessagingProfileIdOk returns a tuple with the DefaultMessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMessagingProfileId

`func (o *MessagingSettings) SetDefaultMessagingProfileId(v string)`

SetDefaultMessagingProfileId sets DefaultMessagingProfileId field to given value.

### HasDefaultMessagingProfileId

`func (o *MessagingSettings) HasDefaultMessagingProfileId() bool`

HasDefaultMessagingProfileId returns a boolean if a field has been set.

### GetDeliveryStatusWebhookUrl

`func (o *MessagingSettings) GetDeliveryStatusWebhookUrl() string`

GetDeliveryStatusWebhookUrl returns the DeliveryStatusWebhookUrl field if non-nil, zero value otherwise.

### GetDeliveryStatusWebhookUrlOk

`func (o *MessagingSettings) GetDeliveryStatusWebhookUrlOk() (*string, bool)`

GetDeliveryStatusWebhookUrlOk returns a tuple with the DeliveryStatusWebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatusWebhookUrl

`func (o *MessagingSettings) SetDeliveryStatusWebhookUrl(v string)`

SetDeliveryStatusWebhookUrl sets DeliveryStatusWebhookUrl field to given value.

### HasDeliveryStatusWebhookUrl

`func (o *MessagingSettings) HasDeliveryStatusWebhookUrl() bool`

HasDeliveryStatusWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


