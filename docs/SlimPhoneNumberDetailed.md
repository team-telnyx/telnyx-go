# SlimPhoneNumberDetailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies the resource. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**PhoneNumber** | Pointer to **string** | The +E.164-formatted phone number associated with this record. | [optional] [readonly] 
**CountryIsoAlpha2** | Pointer to **string** | The ISO 3166-1 alpha-2 country code of the phone number. | [optional] [readonly] 
**Status** | Pointer to **string** | The phone number&#39;s current status. | [optional] [readonly] 
**ExternalPin** | Pointer to **string** | If someone attempts to port your phone number away from Telnyx and your phone number has an external PIN set, Telnyx will attempt to verify that you provided the correct external PIN to the winning carrier. Note that not all carriers cooperate with this security mechanism. | [optional] 
**ConnectionId** | Pointer to **string** | Identifies the connection associated with the phone number. | [optional] 
**CustomerReference** | Pointer to **string** | A customer reference string for customer look ups. | [optional] 
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

## Methods

### NewSlimPhoneNumberDetailed

`func NewSlimPhoneNumberDetailed() *SlimPhoneNumberDetailed`

NewSlimPhoneNumberDetailed instantiates a new SlimPhoneNumberDetailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSlimPhoneNumberDetailedWithDefaults

`func NewSlimPhoneNumberDetailedWithDefaults() *SlimPhoneNumberDetailed`

NewSlimPhoneNumberDetailedWithDefaults instantiates a new SlimPhoneNumberDetailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SlimPhoneNumberDetailed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SlimPhoneNumberDetailed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SlimPhoneNumberDetailed) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SlimPhoneNumberDetailed) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *SlimPhoneNumberDetailed) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *SlimPhoneNumberDetailed) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *SlimPhoneNumberDetailed) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *SlimPhoneNumberDetailed) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *SlimPhoneNumberDetailed) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *SlimPhoneNumberDetailed) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *SlimPhoneNumberDetailed) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *SlimPhoneNumberDetailed) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetCountryIsoAlpha2

`func (o *SlimPhoneNumberDetailed) GetCountryIsoAlpha2() string`

GetCountryIsoAlpha2 returns the CountryIsoAlpha2 field if non-nil, zero value otherwise.

### GetCountryIsoAlpha2Ok

`func (o *SlimPhoneNumberDetailed) GetCountryIsoAlpha2Ok() (*string, bool)`

GetCountryIsoAlpha2Ok returns a tuple with the CountryIsoAlpha2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIsoAlpha2

`func (o *SlimPhoneNumberDetailed) SetCountryIsoAlpha2(v string)`

SetCountryIsoAlpha2 sets CountryIsoAlpha2 field to given value.

### HasCountryIsoAlpha2

`func (o *SlimPhoneNumberDetailed) HasCountryIsoAlpha2() bool`

HasCountryIsoAlpha2 returns a boolean if a field has been set.

### GetStatus

`func (o *SlimPhoneNumberDetailed) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SlimPhoneNumberDetailed) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SlimPhoneNumberDetailed) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SlimPhoneNumberDetailed) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExternalPin

`func (o *SlimPhoneNumberDetailed) GetExternalPin() string`

GetExternalPin returns the ExternalPin field if non-nil, zero value otherwise.

### GetExternalPinOk

`func (o *SlimPhoneNumberDetailed) GetExternalPinOk() (*string, bool)`

GetExternalPinOk returns a tuple with the ExternalPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPin

`func (o *SlimPhoneNumberDetailed) SetExternalPin(v string)`

SetExternalPin sets ExternalPin field to given value.

### HasExternalPin

`func (o *SlimPhoneNumberDetailed) HasExternalPin() bool`

HasExternalPin returns a boolean if a field has been set.

### GetConnectionId

`func (o *SlimPhoneNumberDetailed) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *SlimPhoneNumberDetailed) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *SlimPhoneNumberDetailed) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *SlimPhoneNumberDetailed) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCustomerReference

`func (o *SlimPhoneNumberDetailed) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *SlimPhoneNumberDetailed) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *SlimPhoneNumberDetailed) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *SlimPhoneNumberDetailed) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetBillingGroupId

`func (o *SlimPhoneNumberDetailed) GetBillingGroupId() string`

GetBillingGroupId returns the BillingGroupId field if non-nil, zero value otherwise.

### GetBillingGroupIdOk

`func (o *SlimPhoneNumberDetailed) GetBillingGroupIdOk() (*string, bool)`

GetBillingGroupIdOk returns a tuple with the BillingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupId

`func (o *SlimPhoneNumberDetailed) SetBillingGroupId(v string)`

SetBillingGroupId sets BillingGroupId field to given value.

### HasBillingGroupId

`func (o *SlimPhoneNumberDetailed) HasBillingGroupId() bool`

HasBillingGroupId returns a boolean if a field has been set.

### GetEmergencyEnabled

`func (o *SlimPhoneNumberDetailed) GetEmergencyEnabled() bool`

GetEmergencyEnabled returns the EmergencyEnabled field if non-nil, zero value otherwise.

### GetEmergencyEnabledOk

`func (o *SlimPhoneNumberDetailed) GetEmergencyEnabledOk() (*bool, bool)`

GetEmergencyEnabledOk returns a tuple with the EmergencyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyEnabled

`func (o *SlimPhoneNumberDetailed) SetEmergencyEnabled(v bool)`

SetEmergencyEnabled sets EmergencyEnabled field to given value.

### HasEmergencyEnabled

`func (o *SlimPhoneNumberDetailed) HasEmergencyEnabled() bool`

HasEmergencyEnabled returns a boolean if a field has been set.

### GetEmergencyAddressId

`func (o *SlimPhoneNumberDetailed) GetEmergencyAddressId() string`

GetEmergencyAddressId returns the EmergencyAddressId field if non-nil, zero value otherwise.

### GetEmergencyAddressIdOk

`func (o *SlimPhoneNumberDetailed) GetEmergencyAddressIdOk() (*string, bool)`

GetEmergencyAddressIdOk returns a tuple with the EmergencyAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyAddressId

`func (o *SlimPhoneNumberDetailed) SetEmergencyAddressId(v string)`

SetEmergencyAddressId sets EmergencyAddressId field to given value.

### HasEmergencyAddressId

`func (o *SlimPhoneNumberDetailed) HasEmergencyAddressId() bool`

HasEmergencyAddressId returns a boolean if a field has been set.

### GetEmergencyStatus

`func (o *SlimPhoneNumberDetailed) GetEmergencyStatus() string`

GetEmergencyStatus returns the EmergencyStatus field if non-nil, zero value otherwise.

### GetEmergencyStatusOk

`func (o *SlimPhoneNumberDetailed) GetEmergencyStatusOk() (*string, bool)`

GetEmergencyStatusOk returns a tuple with the EmergencyStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyStatus

`func (o *SlimPhoneNumberDetailed) SetEmergencyStatus(v string)`

SetEmergencyStatus sets EmergencyStatus field to given value.

### HasEmergencyStatus

`func (o *SlimPhoneNumberDetailed) HasEmergencyStatus() bool`

HasEmergencyStatus returns a boolean if a field has been set.

### GetCallForwardingEnabled

`func (o *SlimPhoneNumberDetailed) GetCallForwardingEnabled() bool`

GetCallForwardingEnabled returns the CallForwardingEnabled field if non-nil, zero value otherwise.

### GetCallForwardingEnabledOk

`func (o *SlimPhoneNumberDetailed) GetCallForwardingEnabledOk() (*bool, bool)`

GetCallForwardingEnabledOk returns a tuple with the CallForwardingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallForwardingEnabled

`func (o *SlimPhoneNumberDetailed) SetCallForwardingEnabled(v bool)`

SetCallForwardingEnabled sets CallForwardingEnabled field to given value.

### HasCallForwardingEnabled

`func (o *SlimPhoneNumberDetailed) HasCallForwardingEnabled() bool`

HasCallForwardingEnabled returns a boolean if a field has been set.

### GetCnamListingEnabled

`func (o *SlimPhoneNumberDetailed) GetCnamListingEnabled() bool`

GetCnamListingEnabled returns the CnamListingEnabled field if non-nil, zero value otherwise.

### GetCnamListingEnabledOk

`func (o *SlimPhoneNumberDetailed) GetCnamListingEnabledOk() (*bool, bool)`

GetCnamListingEnabledOk returns a tuple with the CnamListingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnamListingEnabled

`func (o *SlimPhoneNumberDetailed) SetCnamListingEnabled(v bool)`

SetCnamListingEnabled sets CnamListingEnabled field to given value.

### HasCnamListingEnabled

`func (o *SlimPhoneNumberDetailed) HasCnamListingEnabled() bool`

HasCnamListingEnabled returns a boolean if a field has been set.

### GetCallerIdNameEnabled

`func (o *SlimPhoneNumberDetailed) GetCallerIdNameEnabled() bool`

GetCallerIdNameEnabled returns the CallerIdNameEnabled field if non-nil, zero value otherwise.

### GetCallerIdNameEnabledOk

`func (o *SlimPhoneNumberDetailed) GetCallerIdNameEnabledOk() (*bool, bool)`

GetCallerIdNameEnabledOk returns a tuple with the CallerIdNameEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdNameEnabled

`func (o *SlimPhoneNumberDetailed) SetCallerIdNameEnabled(v bool)`

SetCallerIdNameEnabled sets CallerIdNameEnabled field to given value.

### HasCallerIdNameEnabled

`func (o *SlimPhoneNumberDetailed) HasCallerIdNameEnabled() bool`

HasCallerIdNameEnabled returns a boolean if a field has been set.

### GetCallRecordingEnabled

`func (o *SlimPhoneNumberDetailed) GetCallRecordingEnabled() bool`

GetCallRecordingEnabled returns the CallRecordingEnabled field if non-nil, zero value otherwise.

### GetCallRecordingEnabledOk

`func (o *SlimPhoneNumberDetailed) GetCallRecordingEnabledOk() (*bool, bool)`

GetCallRecordingEnabledOk returns a tuple with the CallRecordingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordingEnabled

`func (o *SlimPhoneNumberDetailed) SetCallRecordingEnabled(v bool)`

SetCallRecordingEnabled sets CallRecordingEnabled field to given value.

### HasCallRecordingEnabled

`func (o *SlimPhoneNumberDetailed) HasCallRecordingEnabled() bool`

HasCallRecordingEnabled returns a boolean if a field has been set.

### GetT38FaxGatewayEnabled

`func (o *SlimPhoneNumberDetailed) GetT38FaxGatewayEnabled() bool`

GetT38FaxGatewayEnabled returns the T38FaxGatewayEnabled field if non-nil, zero value otherwise.

### GetT38FaxGatewayEnabledOk

`func (o *SlimPhoneNumberDetailed) GetT38FaxGatewayEnabledOk() (*bool, bool)`

GetT38FaxGatewayEnabledOk returns a tuple with the T38FaxGatewayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT38FaxGatewayEnabled

`func (o *SlimPhoneNumberDetailed) SetT38FaxGatewayEnabled(v bool)`

SetT38FaxGatewayEnabled sets T38FaxGatewayEnabled field to given value.

### HasT38FaxGatewayEnabled

`func (o *SlimPhoneNumberDetailed) HasT38FaxGatewayEnabled() bool`

HasT38FaxGatewayEnabled returns a boolean if a field has been set.

### GetPurchasedAt

`func (o *SlimPhoneNumberDetailed) GetPurchasedAt() string`

GetPurchasedAt returns the PurchasedAt field if non-nil, zero value otherwise.

### GetPurchasedAtOk

`func (o *SlimPhoneNumberDetailed) GetPurchasedAtOk() (*string, bool)`

GetPurchasedAtOk returns a tuple with the PurchasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedAt

`func (o *SlimPhoneNumberDetailed) SetPurchasedAt(v string)`

SetPurchasedAt sets PurchasedAt field to given value.

### HasPurchasedAt

`func (o *SlimPhoneNumberDetailed) HasPurchasedAt() bool`

HasPurchasedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SlimPhoneNumberDetailed) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SlimPhoneNumberDetailed) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SlimPhoneNumberDetailed) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SlimPhoneNumberDetailed) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetPhoneNumberType

`func (o *SlimPhoneNumberDetailed) GetPhoneNumberType() string`

GetPhoneNumberType returns the PhoneNumberType field if non-nil, zero value otherwise.

### GetPhoneNumberTypeOk

`func (o *SlimPhoneNumberDetailed) GetPhoneNumberTypeOk() (*string, bool)`

GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberType

`func (o *SlimPhoneNumberDetailed) SetPhoneNumberType(v string)`

SetPhoneNumberType sets PhoneNumberType field to given value.

### HasPhoneNumberType

`func (o *SlimPhoneNumberDetailed) HasPhoneNumberType() bool`

HasPhoneNumberType returns a boolean if a field has been set.

### GetInboundCallScreening

`func (o *SlimPhoneNumberDetailed) GetInboundCallScreening() string`

GetInboundCallScreening returns the InboundCallScreening field if non-nil, zero value otherwise.

### GetInboundCallScreeningOk

`func (o *SlimPhoneNumberDetailed) GetInboundCallScreeningOk() (*string, bool)`

GetInboundCallScreeningOk returns a tuple with the InboundCallScreening field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInboundCallScreening

`func (o *SlimPhoneNumberDetailed) SetInboundCallScreening(v string)`

SetInboundCallScreening sets InboundCallScreening field to given value.

### HasInboundCallScreening

`func (o *SlimPhoneNumberDetailed) HasInboundCallScreening() bool`

HasInboundCallScreening returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


