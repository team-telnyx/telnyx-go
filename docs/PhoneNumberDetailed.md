# PhoneNumberDetailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies the resource. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**PhoneNumber** | Pointer to **string** | The +E.164-formatted phone number associated with this record. | [optional] [readonly] 
**CountryIsoAlpha2** | Pointer to **string** | The ISO 3166-1 alpha-2 country code of the phone number. | [optional] [readonly] 
**Status** | Pointer to **string** | The phone number&#39;s current status. | [optional] [readonly] 
**Tags** | Pointer to **[]string** | A list of user-assigned tags to help manage the phone number. | [optional] 
**ExternalPin** | Pointer to **string** | If someone attempts to port your phone number away from Telnyx and your phone number has an external PIN set, Telnyx will attempt to verify that you provided the correct external PIN to the winning carrier. Note that not all carriers cooperate with this security mechanism. | [optional] 
**ConnectionName** | Pointer to **string** | The user-assigned name of the connection to be associated with this phone number. | [optional] [readonly] 
**ConnectionId** | Pointer to **string** | Identifies the connection associated with the phone number. | [optional] 
**CustomerReference** | Pointer to **string** | A customer reference string for customer look ups. | [optional] 
**MessagingProfileId** | Pointer to **string** | Identifies the messaging profile associated with the phone number. | [optional] 
**MessagingProfileName** | Pointer to **string** | The name of the messaging profile associated with the phone number. | [optional] 
**BillingGroupId** | Pointer to **string** | Identifies the billing group associated with the phone number. | [optional] 
**EmergencyEnabled** | Pointer to **bool** | Indicates whether emergency services are enabled for this number. | [optional] [readonly] 
**EmergencyAddressId** | Pointer to **string** | Identifies the emergency address associated with the phone number. | [optional] [readonly] 
**EmergencyStatus** | Pointer to **string** | Indicates the status of the provisioning of emergency services for the phone number. This field contains information about activity that may be ongoing for a number where it either is being provisioned or deprovisioned but is not yet enabled/disabled. | [optional] 
**CallForwardingEnabled** | Pointer to **bool** | Indicates if call forwarding will be enabled for this number if forwards_to and forwarding_type are filled in. Defaults to true for backwards compatibility with APIV1 use of numbers endpoints. | [optional] [readonly] [default to true]
**CnamListingEnabled** | Pointer to **bool** | Indicates whether a CNAM listing is enabled for this number. | [optional] [readonly] 
**CallerIdNameEnabled** | Pointer to **bool** | Indicates whether caller ID is enabled for this number. | [optional] [readonly] 
**CallRecordingEnabled** | Pointer to **bool** | Indicates whether call recording is enabled for this number. | [optional] [readonly] 
**T38FaxGatewayEnabled** | Pointer to **bool** | Indicates whether T38 Fax Gateway for inbound calls to this number. | [optional] [readonly] 
**PurchasedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was purchased. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was created. | [optional] [readonly] 
**PhoneNumberType** | Pointer to **string** | The phone number&#39;s type. Note: For numbers purchased prior to July 2023 or when fetching a number&#39;s details immediately after a purchase completes, the legacy values &#x60;tollfree&#x60;, &#x60;shortcode&#x60; or &#x60;longcode&#x60; may be returned instead. | [optional] [readonly] 
**InboundCallScreening** | Pointer to **string** | The inbound_call_screening setting is a phone number configuration option variable that allows users to configure their settings to block or flag fraudulent calls. It can be set to disabled, reject_calls, or flag_calls. This feature has an additional per-number monthly cost associated with it. | [optional] [default to "disabled"]
**SourceType** | Pointer to **NullableString** | Indicates if the phone number was purchased or ported in. For some numbers this information may not be available. | [optional] [readonly] 

## Methods

### NewPhoneNumberDetailed

`func NewPhoneNumberDetailed() *PhoneNumberDetailed`

NewPhoneNumberDetailed instantiates a new PhoneNumberDetailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneNumberDetailedWithDefaults

`func NewPhoneNumberDetailedWithDefaults() *PhoneNumberDetailed`

NewPhoneNumberDetailedWithDefaults instantiates a new PhoneNumberDetailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PhoneNumberDetailed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhoneNumberDetailed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhoneNumberDetailed) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PhoneNumberDetailed) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *PhoneNumberDetailed) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PhoneNumberDetailed) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PhoneNumberDetailed) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PhoneNumberDetailed) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *PhoneNumberDetailed) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PhoneNumberDetailed) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PhoneNumberDetailed) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PhoneNumberDetailed) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetCountryIsoAlpha2

`func (o *PhoneNumberDetailed) GetCountryIsoAlpha2() string`

GetCountryIsoAlpha2 returns the CountryIsoAlpha2 field if non-nil, zero value otherwise.

### GetCountryIsoAlpha2Ok

`func (o *PhoneNumberDetailed) GetCountryIsoAlpha2Ok() (*string, bool)`

GetCountryIsoAlpha2Ok returns a tuple with the CountryIsoAlpha2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsoAlpha2

`func (o *PhoneNumberDetailed) SetCountryIsoAlpha2(v string)`

SetCountryIsoAlpha2 sets CountryIsoAlpha2 field to given value.

### HasCountryIsoAlpha2

`func (o *PhoneNumberDetailed) HasCountryIsoAlpha2() bool`

HasCountryIsoAlpha2 returns a boolean if a field has been set.

### GetStatus

`func (o *PhoneNumberDetailed) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PhoneNumberDetailed) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PhoneNumberDetailed) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PhoneNumberDetailed) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *PhoneNumberDetailed) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PhoneNumberDetailed) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PhoneNumberDetailed) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PhoneNumberDetailed) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetExternalPin

`func (o *PhoneNumberDetailed) GetExternalPin() string`

GetExternalPin returns the ExternalPin field if non-nil, zero value otherwise.

### GetExternalPinOk

`func (o *PhoneNumberDetailed) GetExternalPinOk() (*string, bool)`

GetExternalPinOk returns a tuple with the ExternalPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPin

`func (o *PhoneNumberDetailed) SetExternalPin(v string)`

SetExternalPin sets ExternalPin field to given value.

### HasExternalPin

`func (o *PhoneNumberDetailed) HasExternalPin() bool`

HasExternalPin returns a boolean if a field has been set.

### GetConnectionName

`func (o *PhoneNumberDetailed) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *PhoneNumberDetailed) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *PhoneNumberDetailed) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.

### HasConnectionName

`func (o *PhoneNumberDetailed) HasConnectionName() bool`

HasConnectionName returns a boolean if a field has been set.

### GetConnectionId

`func (o *PhoneNumberDetailed) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *PhoneNumberDetailed) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *PhoneNumberDetailed) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *PhoneNumberDetailed) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCustomerReference

`func (o *PhoneNumberDetailed) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *PhoneNumberDetailed) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *PhoneNumberDetailed) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *PhoneNumberDetailed) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetMessagingProfileId

`func (o *PhoneNumberDetailed) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *PhoneNumberDetailed) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *PhoneNumberDetailed) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.

### HasMessagingProfileId

`func (o *PhoneNumberDetailed) HasMessagingProfileId() bool`

HasMessagingProfileId returns a boolean if a field has been set.

### GetMessagingProfileName

`func (o *PhoneNumberDetailed) GetMessagingProfileName() string`

GetMessagingProfileName returns the MessagingProfileName field if non-nil, zero value otherwise.

### GetMessagingProfileNameOk

`func (o *PhoneNumberDetailed) GetMessagingProfileNameOk() (*string, bool)`

GetMessagingProfileNameOk returns a tuple with the MessagingProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileName

`func (o *PhoneNumberDetailed) SetMessagingProfileName(v string)`

SetMessagingProfileName sets MessagingProfileName field to given value.

### HasMessagingProfileName

`func (o *PhoneNumberDetailed) HasMessagingProfileName() bool`

HasMessagingProfileName returns a boolean if a field has been set.

### GetBillingGroupId

`func (o *PhoneNumberDetailed) GetBillingGroupId() string`

GetBillingGroupId returns the BillingGroupId field if non-nil, zero value otherwise.

### GetBillingGroupIdOk

`func (o *PhoneNumberDetailed) GetBillingGroupIdOk() (*string, bool)`

GetBillingGroupIdOk returns a tuple with the BillingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupId

`func (o *PhoneNumberDetailed) SetBillingGroupId(v string)`

SetBillingGroupId sets BillingGroupId field to given value.

### HasBillingGroupId

`func (o *PhoneNumberDetailed) HasBillingGroupId() bool`

HasBillingGroupId returns a boolean if a field has been set.

### GetEmergencyEnabled

`func (o *PhoneNumberDetailed) GetEmergencyEnabled() bool`

GetEmergencyEnabled returns the EmergencyEnabled field if non-nil, zero value otherwise.

### GetEmergencyEnabledOk

`func (o *PhoneNumberDetailed) GetEmergencyEnabledOk() (*bool, bool)`

GetEmergencyEnabledOk returns a tuple with the EmergencyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyEnabled

`func (o *PhoneNumberDetailed) SetEmergencyEnabled(v bool)`

SetEmergencyEnabled sets EmergencyEnabled field to given value.

### HasEmergencyEnabled

`func (o *PhoneNumberDetailed) HasEmergencyEnabled() bool`

HasEmergencyEnabled returns a boolean if a field has been set.

### GetEmergencyAddressId

`func (o *PhoneNumberDetailed) GetEmergencyAddressId() string`

GetEmergencyAddressId returns the EmergencyAddressId field if non-nil, zero value otherwise.

### GetEmergencyAddressIdOk

`func (o *PhoneNumberDetailed) GetEmergencyAddressIdOk() (*string, bool)`

GetEmergencyAddressIdOk returns a tuple with the EmergencyAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyAddressId

`func (o *PhoneNumberDetailed) SetEmergencyAddressId(v string)`

SetEmergencyAddressId sets EmergencyAddressId field to given value.

### HasEmergencyAddressId

`func (o *PhoneNumberDetailed) HasEmergencyAddressId() bool`

HasEmergencyAddressId returns a boolean if a field has been set.

### GetEmergencyStatus

`func (o *PhoneNumberDetailed) GetEmergencyStatus() string`

GetEmergencyStatus returns the EmergencyStatus field if non-nil, zero value otherwise.

### GetEmergencyStatusOk

`func (o *PhoneNumberDetailed) GetEmergencyStatusOk() (*string, bool)`

GetEmergencyStatusOk returns a tuple with the EmergencyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyStatus

`func (o *PhoneNumberDetailed) SetEmergencyStatus(v string)`

SetEmergencyStatus sets EmergencyStatus field to given value.

### HasEmergencyStatus

`func (o *PhoneNumberDetailed) HasEmergencyStatus() bool`

HasEmergencyStatus returns a boolean if a field has been set.

### GetCallForwardingEnabled

`func (o *PhoneNumberDetailed) GetCallForwardingEnabled() bool`

GetCallForwardingEnabled returns the CallForwardingEnabled field if non-nil, zero value otherwise.

### GetCallForwardingEnabledOk

`func (o *PhoneNumberDetailed) GetCallForwardingEnabledOk() (*bool, bool)`

GetCallForwardingEnabledOk returns a tuple with the CallForwardingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallForwardingEnabled

`func (o *PhoneNumberDetailed) SetCallForwardingEnabled(v bool)`

SetCallForwardingEnabled sets CallForwardingEnabled field to given value.

### HasCallForwardingEnabled

`func (o *PhoneNumberDetailed) HasCallForwardingEnabled() bool`

HasCallForwardingEnabled returns a boolean if a field has been set.

### GetCnamListingEnabled

`func (o *PhoneNumberDetailed) GetCnamListingEnabled() bool`

GetCnamListingEnabled returns the CnamListingEnabled field if non-nil, zero value otherwise.

### GetCnamListingEnabledOk

`func (o *PhoneNumberDetailed) GetCnamListingEnabledOk() (*bool, bool)`

GetCnamListingEnabledOk returns a tuple with the CnamListingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnamListingEnabled

`func (o *PhoneNumberDetailed) SetCnamListingEnabled(v bool)`

SetCnamListingEnabled sets CnamListingEnabled field to given value.

### HasCnamListingEnabled

`func (o *PhoneNumberDetailed) HasCnamListingEnabled() bool`

HasCnamListingEnabled returns a boolean if a field has been set.

### GetCallerIdNameEnabled

`func (o *PhoneNumberDetailed) GetCallerIdNameEnabled() bool`

GetCallerIdNameEnabled returns the CallerIdNameEnabled field if non-nil, zero value otherwise.

### GetCallerIdNameEnabledOk

`func (o *PhoneNumberDetailed) GetCallerIdNameEnabledOk() (*bool, bool)`

GetCallerIdNameEnabledOk returns a tuple with the CallerIdNameEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdNameEnabled

`func (o *PhoneNumberDetailed) SetCallerIdNameEnabled(v bool)`

SetCallerIdNameEnabled sets CallerIdNameEnabled field to given value.

### HasCallerIdNameEnabled

`func (o *PhoneNumberDetailed) HasCallerIdNameEnabled() bool`

HasCallerIdNameEnabled returns a boolean if a field has been set.

### GetCallRecordingEnabled

`func (o *PhoneNumberDetailed) GetCallRecordingEnabled() bool`

GetCallRecordingEnabled returns the CallRecordingEnabled field if non-nil, zero value otherwise.

### GetCallRecordingEnabledOk

`func (o *PhoneNumberDetailed) GetCallRecordingEnabledOk() (*bool, bool)`

GetCallRecordingEnabledOk returns a tuple with the CallRecordingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordingEnabled

`func (o *PhoneNumberDetailed) SetCallRecordingEnabled(v bool)`

SetCallRecordingEnabled sets CallRecordingEnabled field to given value.

### HasCallRecordingEnabled

`func (o *PhoneNumberDetailed) HasCallRecordingEnabled() bool`

HasCallRecordingEnabled returns a boolean if a field has been set.

### GetT38FaxGatewayEnabled

`func (o *PhoneNumberDetailed) GetT38FaxGatewayEnabled() bool`

GetT38FaxGatewayEnabled returns the T38FaxGatewayEnabled field if non-nil, zero value otherwise.

### GetT38FaxGatewayEnabledOk

`func (o *PhoneNumberDetailed) GetT38FaxGatewayEnabledOk() (*bool, bool)`

GetT38FaxGatewayEnabledOk returns a tuple with the T38FaxGatewayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT38FaxGatewayEnabled

`func (o *PhoneNumberDetailed) SetT38FaxGatewayEnabled(v bool)`

SetT38FaxGatewayEnabled sets T38FaxGatewayEnabled field to given value.

### HasT38FaxGatewayEnabled

`func (o *PhoneNumberDetailed) HasT38FaxGatewayEnabled() bool`

HasT38FaxGatewayEnabled returns a boolean if a field has been set.

### GetPurchasedAt

`func (o *PhoneNumberDetailed) GetPurchasedAt() string`

GetPurchasedAt returns the PurchasedAt field if non-nil, zero value otherwise.

### GetPurchasedAtOk

`func (o *PhoneNumberDetailed) GetPurchasedAtOk() (*string, bool)`

GetPurchasedAtOk returns a tuple with the PurchasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedAt

`func (o *PhoneNumberDetailed) SetPurchasedAt(v string)`

SetPurchasedAt sets PurchasedAt field to given value.

### HasPurchasedAt

`func (o *PhoneNumberDetailed) HasPurchasedAt() bool`

HasPurchasedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PhoneNumberDetailed) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PhoneNumberDetailed) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PhoneNumberDetailed) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PhoneNumberDetailed) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPhoneNumberType

`func (o *PhoneNumberDetailed) GetPhoneNumberType() string`

GetPhoneNumberType returns the PhoneNumberType field if non-nil, zero value otherwise.

### GetPhoneNumberTypeOk

`func (o *PhoneNumberDetailed) GetPhoneNumberTypeOk() (*string, bool)`

GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberType

`func (o *PhoneNumberDetailed) SetPhoneNumberType(v string)`

SetPhoneNumberType sets PhoneNumberType field to given value.

### HasPhoneNumberType

`func (o *PhoneNumberDetailed) HasPhoneNumberType() bool`

HasPhoneNumberType returns a boolean if a field has been set.

### GetInboundCallScreening

`func (o *PhoneNumberDetailed) GetInboundCallScreening() string`

GetInboundCallScreening returns the InboundCallScreening field if non-nil, zero value otherwise.

### GetInboundCallScreeningOk

`func (o *PhoneNumberDetailed) GetInboundCallScreeningOk() (*string, bool)`

GetInboundCallScreeningOk returns a tuple with the InboundCallScreening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundCallScreening

`func (o *PhoneNumberDetailed) SetInboundCallScreening(v string)`

SetInboundCallScreening sets InboundCallScreening field to given value.

### HasInboundCallScreening

`func (o *PhoneNumberDetailed) HasInboundCallScreening() bool`

HasInboundCallScreening returns a boolean if a field has been set.

### GetSourceType

`func (o *PhoneNumberDetailed) GetSourceType() string`

GetSourceType returns the SourceType field if non-nil, zero value otherwise.

### GetSourceTypeOk

`func (o *PhoneNumberDetailed) GetSourceTypeOk() (*string, bool)`

GetSourceTypeOk returns a tuple with the SourceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceType

`func (o *PhoneNumberDetailed) SetSourceType(v string)`

SetSourceType sets SourceType field to given value.

### HasSourceType

`func (o *PhoneNumberDetailed) HasSourceType() bool`

HasSourceType returns a boolean if a field has been set.

### SetSourceTypeNil

`func (o *PhoneNumberDetailed) SetSourceTypeNil(b bool)`

 SetSourceTypeNil sets the value for SourceType to be an explicit nil

### UnsetSourceType
`func (o *PhoneNumberDetailed) UnsetSourceType()`

UnsetSourceType ensures that no value is present for SourceType, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


