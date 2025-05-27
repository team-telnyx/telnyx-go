# UpdatePortingOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Misc** | Pointer to [**PortingOrderMisc**](PortingOrderMisc.md) |  | [optional] 
**EndUser** | Pointer to [**PortingOrderEndUser**](PortingOrderEndUser.md) |  | [optional] 
**Documents** | Pointer to [**PortingOrderDocuments**](PortingOrderDocuments.md) |  | [optional] 
**ActivationSettings** | Pointer to [**UpdatePortingOrderActivationSettings**](UpdatePortingOrderActivationSettings.md) |  | [optional] 
**PhoneNumberConfiguration** | Pointer to [**PortingOrderPhoneNumberConfiguration**](PortingOrderPhoneNumberConfiguration.md) |  | [optional] 
**RequirementGroupId** | Pointer to **string** | If present, we will read the current values from the specified Requirement Group into the Documents and Requirements for this Porting Order. Note that any future changes in the Requirement Group would have no impact on this Porting Order. We will return an error if a specified Requirement Group conflicts with documents or requirements in the same request. | [optional] 
**Requirements** | Pointer to [**[]UpdatePortingOrderRequirement**](UpdatePortingOrderRequirement.md) | List of requirements for porting numbers.  | [optional] 
**UserFeedback** | Pointer to [**PortingOrderUserFeedback**](PortingOrderUserFeedback.md) |  | [optional] 
**WebhookUrl** | Pointer to **string** |  | [optional] 
**CustomerReference** | Pointer to **string** |  | [optional] 
**Messaging** | Pointer to [**UpdatePortingOrderMessaging**](UpdatePortingOrderMessaging.md) |  | [optional] 

## Methods

### NewUpdatePortingOrder

`func NewUpdatePortingOrder() *UpdatePortingOrder`

NewUpdatePortingOrder instantiates a new UpdatePortingOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePortingOrderWithDefaults

`func NewUpdatePortingOrderWithDefaults() *UpdatePortingOrder`

NewUpdatePortingOrderWithDefaults instantiates a new UpdatePortingOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMisc

`func (o *UpdatePortingOrder) GetMisc() PortingOrderMisc`

GetMisc returns the Misc field if non-nil, zero value otherwise.

### GetMiscOk

`func (o *UpdatePortingOrder) GetMiscOk() (*PortingOrderMisc, bool)`

GetMiscOk returns a tuple with the Misc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMisc

`func (o *UpdatePortingOrder) SetMisc(v PortingOrderMisc)`

SetMisc sets Misc field to given value.

### HasMisc

`func (o *UpdatePortingOrder) HasMisc() bool`

HasMisc returns a boolean if a field has been set.

### GetEndUser

`func (o *UpdatePortingOrder) GetEndUser() PortingOrderEndUser`

GetEndUser returns the EndUser field if non-nil, zero value otherwise.

### GetEndUserOk

`func (o *UpdatePortingOrder) GetEndUserOk() (*PortingOrderEndUser, bool)`

GetEndUserOk returns a tuple with the EndUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUser

`func (o *UpdatePortingOrder) SetEndUser(v PortingOrderEndUser)`

SetEndUser sets EndUser field to given value.

### HasEndUser

`func (o *UpdatePortingOrder) HasEndUser() bool`

HasEndUser returns a boolean if a field has been set.

### GetDocuments

`func (o *UpdatePortingOrder) GetDocuments() PortingOrderDocuments`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *UpdatePortingOrder) GetDocumentsOk() (*PortingOrderDocuments, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *UpdatePortingOrder) SetDocuments(v PortingOrderDocuments)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *UpdatePortingOrder) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetActivationSettings

`func (o *UpdatePortingOrder) GetActivationSettings() UpdatePortingOrderActivationSettings`

GetActivationSettings returns the ActivationSettings field if non-nil, zero value otherwise.

### GetActivationSettingsOk

`func (o *UpdatePortingOrder) GetActivationSettingsOk() (*UpdatePortingOrderActivationSettings, bool)`

GetActivationSettingsOk returns a tuple with the ActivationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationSettings

`func (o *UpdatePortingOrder) SetActivationSettings(v UpdatePortingOrderActivationSettings)`

SetActivationSettings sets ActivationSettings field to given value.

### HasActivationSettings

`func (o *UpdatePortingOrder) HasActivationSettings() bool`

HasActivationSettings returns a boolean if a field has been set.

### GetPhoneNumberConfiguration

`func (o *UpdatePortingOrder) GetPhoneNumberConfiguration() PortingOrderPhoneNumberConfiguration`

GetPhoneNumberConfiguration returns the PhoneNumberConfiguration field if non-nil, zero value otherwise.

### GetPhoneNumberConfigurationOk

`func (o *UpdatePortingOrder) GetPhoneNumberConfigurationOk() (*PortingOrderPhoneNumberConfiguration, bool)`

GetPhoneNumberConfigurationOk returns a tuple with the PhoneNumberConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberConfiguration

`func (o *UpdatePortingOrder) SetPhoneNumberConfiguration(v PortingOrderPhoneNumberConfiguration)`

SetPhoneNumberConfiguration sets PhoneNumberConfiguration field to given value.

### HasPhoneNumberConfiguration

`func (o *UpdatePortingOrder) HasPhoneNumberConfiguration() bool`

HasPhoneNumberConfiguration returns a boolean if a field has been set.

### GetRequirementGroupId

`func (o *UpdatePortingOrder) GetRequirementGroupId() string`

GetRequirementGroupId returns the RequirementGroupId field if non-nil, zero value otherwise.

### GetRequirementGroupIdOk

`func (o *UpdatePortingOrder) GetRequirementGroupIdOk() (*string, bool)`

GetRequirementGroupIdOk returns a tuple with the RequirementGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementGroupId

`func (o *UpdatePortingOrder) SetRequirementGroupId(v string)`

SetRequirementGroupId sets RequirementGroupId field to given value.

### HasRequirementGroupId

`func (o *UpdatePortingOrder) HasRequirementGroupId() bool`

HasRequirementGroupId returns a boolean if a field has been set.

### GetRequirements

`func (o *UpdatePortingOrder) GetRequirements() []UpdatePortingOrderRequirement`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *UpdatePortingOrder) GetRequirementsOk() (*[]UpdatePortingOrderRequirement, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *UpdatePortingOrder) SetRequirements(v []UpdatePortingOrderRequirement)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *UpdatePortingOrder) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### GetUserFeedback

`func (o *UpdatePortingOrder) GetUserFeedback() PortingOrderUserFeedback`

GetUserFeedback returns the UserFeedback field if non-nil, zero value otherwise.

### GetUserFeedbackOk

`func (o *UpdatePortingOrder) GetUserFeedbackOk() (*PortingOrderUserFeedback, bool)`

GetUserFeedbackOk returns a tuple with the UserFeedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFeedback

`func (o *UpdatePortingOrder) SetUserFeedback(v PortingOrderUserFeedback)`

SetUserFeedback sets UserFeedback field to given value.

### HasUserFeedback

`func (o *UpdatePortingOrder) HasUserFeedback() bool`

HasUserFeedback returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *UpdatePortingOrder) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *UpdatePortingOrder) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *UpdatePortingOrder) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *UpdatePortingOrder) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetCustomerReference

`func (o *UpdatePortingOrder) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *UpdatePortingOrder) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *UpdatePortingOrder) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *UpdatePortingOrder) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetMessaging

`func (o *UpdatePortingOrder) GetMessaging() UpdatePortingOrderMessaging`

GetMessaging returns the Messaging field if non-nil, zero value otherwise.

### GetMessagingOk

`func (o *UpdatePortingOrder) GetMessagingOk() (*UpdatePortingOrderMessaging, bool)`

GetMessagingOk returns a tuple with the Messaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessaging

`func (o *UpdatePortingOrder) SetMessaging(v UpdatePortingOrderMessaging)`

SetMessaging sets Messaging field to given value.

### HasMessaging

`func (o *UpdatePortingOrder) HasMessaging() bool`

HasMessaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


