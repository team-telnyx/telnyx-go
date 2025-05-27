# PortingOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies this porting order | [optional] [readonly] 
**CustomerReference** | Pointer to **string** | A customer-specified reference number for customer bookkeeping purposes | [optional] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] [readonly] 
**Status** | Pointer to [**PortingOrderStatus**](PortingOrderStatus.md) |  | [optional] 
**SupportKey** | Pointer to **string** | A key to reference this porting order when contacting Telnyx customer support. This information is not available in draft porting orders. | [optional] [readonly] 
**ParentSupportKey** | Pointer to **string** | A key to reference for the porting order group when contacting Telnyx customer support. This information is not available for porting orders in &#x60;draft&#x60; state | [optional] [readonly] 
**PortingPhoneNumbersCount** | Pointer to **int32** | Count of phone numbers associated with this porting order | [optional] [readonly] 
**OldServiceProviderOcn** | Pointer to **string** | Identifies the old service provider | [optional] [readonly] 
**Documents** | Pointer to [**PortingOrderDocuments**](PortingOrderDocuments.md) |  | [optional] 
**Misc** | Pointer to [**PortingOrderMisc**](PortingOrderMisc.md) |  | [optional] 
**EndUser** | Pointer to [**PortingOrderEndUser**](PortingOrderEndUser.md) |  | [optional] 
**ActivationSettings** | Pointer to [**PortingOrderActivationSettings**](PortingOrderActivationSettings.md) |  | [optional] 
**PhoneNumberConfiguration** | Pointer to [**PortingOrderPhoneNumberConfiguration**](PortingOrderPhoneNumberConfiguration.md) |  | [optional] 
**PhoneNumberType** | Pointer to **string** | The type of the phone number | [optional] 
**Description** | Pointer to **string** | A description of the porting order | [optional] [readonly] 
**Requirements** | Pointer to [**[]PortingOrderRequirement**](PortingOrderRequirement.md) | List of documentation requirements for porting numbers. Can be set directly or via the &#x60;requirement_group_id&#x60; parameter. | [optional] 
**RequirementsMet** | Pointer to **bool** | Is true when the required documentation is met | [optional] 
**UserFeedback** | Pointer to [**PortingOrderUserFeedback**](PortingOrderUserFeedback.md) |  | [optional] 
**UserId** | Pointer to **string** | Identifies the user (or organization) who requested the porting order | [optional] 
**WebhookUrl** | Pointer to **string** |  | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**Messaging** | Pointer to [**PortingOrderMessaging**](PortingOrderMessaging.md) |  | [optional] 

## Methods

### NewPortingOrder

`func NewPortingOrder() *PortingOrder`

NewPortingOrder instantiates a new PortingOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortingOrderWithDefaults

`func NewPortingOrderWithDefaults() *PortingOrder`

NewPortingOrderWithDefaults instantiates a new PortingOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortingOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortingOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortingOrder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortingOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCustomerReference

`func (o *PortingOrder) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *PortingOrder) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *PortingOrder) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *PortingOrder) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PortingOrder) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortingOrder) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortingOrder) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PortingOrder) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PortingOrder) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PortingOrder) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PortingOrder) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PortingOrder) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *PortingOrder) GetStatus() PortingOrderStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PortingOrder) GetStatusOk() (*PortingOrderStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PortingOrder) SetStatus(v PortingOrderStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PortingOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSupportKey

`func (o *PortingOrder) GetSupportKey() string`

GetSupportKey returns the SupportKey field if non-nil, zero value otherwise.

### GetSupportKeyOk

`func (o *PortingOrder) GetSupportKeyOk() (*string, bool)`

GetSupportKeyOk returns a tuple with the SupportKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportKey

`func (o *PortingOrder) SetSupportKey(v string)`

SetSupportKey sets SupportKey field to given value.

### HasSupportKey

`func (o *PortingOrder) HasSupportKey() bool`

HasSupportKey returns a boolean if a field has been set.

### GetParentSupportKey

`func (o *PortingOrder) GetParentSupportKey() string`

GetParentSupportKey returns the ParentSupportKey field if non-nil, zero value otherwise.

### GetParentSupportKeyOk

`func (o *PortingOrder) GetParentSupportKeyOk() (*string, bool)`

GetParentSupportKeyOk returns a tuple with the ParentSupportKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentSupportKey

`func (o *PortingOrder) SetParentSupportKey(v string)`

SetParentSupportKey sets ParentSupportKey field to given value.

### HasParentSupportKey

`func (o *PortingOrder) HasParentSupportKey() bool`

HasParentSupportKey returns a boolean if a field has been set.

### GetPortingPhoneNumbersCount

`func (o *PortingOrder) GetPortingPhoneNumbersCount() int32`

GetPortingPhoneNumbersCount returns the PortingPhoneNumbersCount field if non-nil, zero value otherwise.

### GetPortingPhoneNumbersCountOk

`func (o *PortingOrder) GetPortingPhoneNumbersCountOk() (*int32, bool)`

GetPortingPhoneNumbersCountOk returns a tuple with the PortingPhoneNumbersCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortingPhoneNumbersCount

`func (o *PortingOrder) SetPortingPhoneNumbersCount(v int32)`

SetPortingPhoneNumbersCount sets PortingPhoneNumbersCount field to given value.

### HasPortingPhoneNumbersCount

`func (o *PortingOrder) HasPortingPhoneNumbersCount() bool`

HasPortingPhoneNumbersCount returns a boolean if a field has been set.

### GetOldServiceProviderOcn

`func (o *PortingOrder) GetOldServiceProviderOcn() string`

GetOldServiceProviderOcn returns the OldServiceProviderOcn field if non-nil, zero value otherwise.

### GetOldServiceProviderOcnOk

`func (o *PortingOrder) GetOldServiceProviderOcnOk() (*string, bool)`

GetOldServiceProviderOcnOk returns a tuple with the OldServiceProviderOcn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldServiceProviderOcn

`func (o *PortingOrder) SetOldServiceProviderOcn(v string)`

SetOldServiceProviderOcn sets OldServiceProviderOcn field to given value.

### HasOldServiceProviderOcn

`func (o *PortingOrder) HasOldServiceProviderOcn() bool`

HasOldServiceProviderOcn returns a boolean if a field has been set.

### GetDocuments

`func (o *PortingOrder) GetDocuments() PortingOrderDocuments`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *PortingOrder) GetDocumentsOk() (*PortingOrderDocuments, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *PortingOrder) SetDocuments(v PortingOrderDocuments)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *PortingOrder) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### GetMisc

`func (o *PortingOrder) GetMisc() PortingOrderMisc`

GetMisc returns the Misc field if non-nil, zero value otherwise.

### GetMiscOk

`func (o *PortingOrder) GetMiscOk() (*PortingOrderMisc, bool)`

GetMiscOk returns a tuple with the Misc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMisc

`func (o *PortingOrder) SetMisc(v PortingOrderMisc)`

SetMisc sets Misc field to given value.

### HasMisc

`func (o *PortingOrder) HasMisc() bool`

HasMisc returns a boolean if a field has been set.

### GetEndUser

`func (o *PortingOrder) GetEndUser() PortingOrderEndUser`

GetEndUser returns the EndUser field if non-nil, zero value otherwise.

### GetEndUserOk

`func (o *PortingOrder) GetEndUserOk() (*PortingOrderEndUser, bool)`

GetEndUserOk returns a tuple with the EndUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndUser

`func (o *PortingOrder) SetEndUser(v PortingOrderEndUser)`

SetEndUser sets EndUser field to given value.

### HasEndUser

`func (o *PortingOrder) HasEndUser() bool`

HasEndUser returns a boolean if a field has been set.

### GetActivationSettings

`func (o *PortingOrder) GetActivationSettings() PortingOrderActivationSettings`

GetActivationSettings returns the ActivationSettings field if non-nil, zero value otherwise.

### GetActivationSettingsOk

`func (o *PortingOrder) GetActivationSettingsOk() (*PortingOrderActivationSettings, bool)`

GetActivationSettingsOk returns a tuple with the ActivationSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationSettings

`func (o *PortingOrder) SetActivationSettings(v PortingOrderActivationSettings)`

SetActivationSettings sets ActivationSettings field to given value.

### HasActivationSettings

`func (o *PortingOrder) HasActivationSettings() bool`

HasActivationSettings returns a boolean if a field has been set.

### GetPhoneNumberConfiguration

`func (o *PortingOrder) GetPhoneNumberConfiguration() PortingOrderPhoneNumberConfiguration`

GetPhoneNumberConfiguration returns the PhoneNumberConfiguration field if non-nil, zero value otherwise.

### GetPhoneNumberConfigurationOk

`func (o *PortingOrder) GetPhoneNumberConfigurationOk() (*PortingOrderPhoneNumberConfiguration, bool)`

GetPhoneNumberConfigurationOk returns a tuple with the PhoneNumberConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberConfiguration

`func (o *PortingOrder) SetPhoneNumberConfiguration(v PortingOrderPhoneNumberConfiguration)`

SetPhoneNumberConfiguration sets PhoneNumberConfiguration field to given value.

### HasPhoneNumberConfiguration

`func (o *PortingOrder) HasPhoneNumberConfiguration() bool`

HasPhoneNumberConfiguration returns a boolean if a field has been set.

### GetPhoneNumberType

`func (o *PortingOrder) GetPhoneNumberType() string`

GetPhoneNumberType returns the PhoneNumberType field if non-nil, zero value otherwise.

### GetPhoneNumberTypeOk

`func (o *PortingOrder) GetPhoneNumberTypeOk() (*string, bool)`

GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberType

`func (o *PortingOrder) SetPhoneNumberType(v string)`

SetPhoneNumberType sets PhoneNumberType field to given value.

### HasPhoneNumberType

`func (o *PortingOrder) HasPhoneNumberType() bool`

HasPhoneNumberType returns a boolean if a field has been set.

### GetDescription

`func (o *PortingOrder) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PortingOrder) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PortingOrder) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PortingOrder) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetRequirements

`func (o *PortingOrder) GetRequirements() []PortingOrderRequirement`

GetRequirements returns the Requirements field if non-nil, zero value otherwise.

### GetRequirementsOk

`func (o *PortingOrder) GetRequirementsOk() (*[]PortingOrderRequirement, bool)`

GetRequirementsOk returns a tuple with the Requirements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirements

`func (o *PortingOrder) SetRequirements(v []PortingOrderRequirement)`

SetRequirements sets Requirements field to given value.

### HasRequirements

`func (o *PortingOrder) HasRequirements() bool`

HasRequirements returns a boolean if a field has been set.

### GetRequirementsMet

`func (o *PortingOrder) GetRequirementsMet() bool`

GetRequirementsMet returns the RequirementsMet field if non-nil, zero value otherwise.

### GetRequirementsMetOk

`func (o *PortingOrder) GetRequirementsMetOk() (*bool, bool)`

GetRequirementsMetOk returns a tuple with the RequirementsMet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequirementsMet

`func (o *PortingOrder) SetRequirementsMet(v bool)`

SetRequirementsMet sets RequirementsMet field to given value.

### HasRequirementsMet

`func (o *PortingOrder) HasRequirementsMet() bool`

HasRequirementsMet returns a boolean if a field has been set.

### GetUserFeedback

`func (o *PortingOrder) GetUserFeedback() PortingOrderUserFeedback`

GetUserFeedback returns the UserFeedback field if non-nil, zero value otherwise.

### GetUserFeedbackOk

`func (o *PortingOrder) GetUserFeedbackOk() (*PortingOrderUserFeedback, bool)`

GetUserFeedbackOk returns a tuple with the UserFeedback field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserFeedback

`func (o *PortingOrder) SetUserFeedback(v PortingOrderUserFeedback)`

SetUserFeedback sets UserFeedback field to given value.

### HasUserFeedback

`func (o *PortingOrder) HasUserFeedback() bool`

HasUserFeedback returns a boolean if a field has been set.

### GetUserId

`func (o *PortingOrder) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *PortingOrder) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *PortingOrder) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *PortingOrder) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *PortingOrder) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *PortingOrder) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *PortingOrder) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *PortingOrder) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetRecordType

`func (o *PortingOrder) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PortingOrder) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PortingOrder) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PortingOrder) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetMessaging

`func (o *PortingOrder) GetMessaging() PortingOrderMessaging`

GetMessaging returns the Messaging field if non-nil, zero value otherwise.

### GetMessagingOk

`func (o *PortingOrder) GetMessagingOk() (*PortingOrderMessaging, bool)`

GetMessagingOk returns a tuple with the Messaging field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessaging

`func (o *PortingOrder) SetMessaging(v PortingOrderMessaging)`

SetMessaging sets Messaging field to given value.

### HasMessaging

`func (o *PortingOrder) HasMessaging() bool`

HasMessaging returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


