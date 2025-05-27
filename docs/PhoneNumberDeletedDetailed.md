# PhoneNumberDeletedDetailed

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies the resource. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**PhoneNumber** | Pointer to **string** | The +E.164-formatted phone number associated with this record. | [optional] [readonly] 
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
**CallForwardingEnabled** | Pointer to **bool** | Indicates if call forwarding will be enabled for this number if forwards_to and forwarding_type are filled in. Defaults to true for backwards compatibility with APIV1 use of numbers endpoints. | [optional] [readonly] [default to true]
**CnamListingEnabled** | Pointer to **bool** | Indicates whether a CNAM listing is enabled for this number. | [optional] [readonly] 
**CallerIdNameEnabled** | Pointer to **bool** | Indicates whether caller ID is enabled for this number. | [optional] [readonly] 
**CallRecordingEnabled** | Pointer to **bool** | Indicates whether call recording is enabled for this number. | [optional] [readonly] 
**T38FaxGatewayEnabled** | Pointer to **bool** | Indicates whether T38 Fax Gateway for inbound calls to this number. | [optional] [readonly] 
**PurchasedAt** | Pointer to **string** | ISO 8601 formatted date indicating the time the request was made to purchase the number. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the time it took to activate after the purchase. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date indicating when the resource was updated. | [optional] [readonly] 
**PhoneNumberType** | Pointer to **string** | The phone number&#39;s type. | [optional] [readonly] 

## Methods

### NewPhoneNumberDeletedDetailed

`func NewPhoneNumberDeletedDetailed() *PhoneNumberDeletedDetailed`

NewPhoneNumberDeletedDetailed instantiates a new PhoneNumberDeletedDetailed object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneNumberDeletedDetailedWithDefaults

`func NewPhoneNumberDeletedDetailedWithDefaults() *PhoneNumberDeletedDetailed`

NewPhoneNumberDeletedDetailedWithDefaults instantiates a new PhoneNumberDeletedDetailed object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PhoneNumberDeletedDetailed) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhoneNumberDeletedDetailed) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhoneNumberDeletedDetailed) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PhoneNumberDeletedDetailed) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *PhoneNumberDeletedDetailed) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PhoneNumberDeletedDetailed) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PhoneNumberDeletedDetailed) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PhoneNumberDeletedDetailed) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *PhoneNumberDeletedDetailed) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PhoneNumberDeletedDetailed) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PhoneNumberDeletedDetailed) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PhoneNumberDeletedDetailed) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetStatus

`func (o *PhoneNumberDeletedDetailed) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PhoneNumberDeletedDetailed) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PhoneNumberDeletedDetailed) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PhoneNumberDeletedDetailed) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTags

`func (o *PhoneNumberDeletedDetailed) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PhoneNumberDeletedDetailed) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PhoneNumberDeletedDetailed) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PhoneNumberDeletedDetailed) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetExternalPin

`func (o *PhoneNumberDeletedDetailed) GetExternalPin() string`

GetExternalPin returns the ExternalPin field if non-nil, zero value otherwise.

### GetExternalPinOk

`func (o *PhoneNumberDeletedDetailed) GetExternalPinOk() (*string, bool)`

GetExternalPinOk returns a tuple with the ExternalPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalPin

`func (o *PhoneNumberDeletedDetailed) SetExternalPin(v string)`

SetExternalPin sets ExternalPin field to given value.

### HasExternalPin

`func (o *PhoneNumberDeletedDetailed) HasExternalPin() bool`

HasExternalPin returns a boolean if a field has been set.

### GetConnectionName

`func (o *PhoneNumberDeletedDetailed) GetConnectionName() string`

GetConnectionName returns the ConnectionName field if non-nil, zero value otherwise.

### GetConnectionNameOk

`func (o *PhoneNumberDeletedDetailed) GetConnectionNameOk() (*string, bool)`

GetConnectionNameOk returns a tuple with the ConnectionName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionName

`func (o *PhoneNumberDeletedDetailed) SetConnectionName(v string)`

SetConnectionName sets ConnectionName field to given value.

### HasConnectionName

`func (o *PhoneNumberDeletedDetailed) HasConnectionName() bool`

HasConnectionName returns a boolean if a field has been set.

### GetConnectionId

`func (o *PhoneNumberDeletedDetailed) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *PhoneNumberDeletedDetailed) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *PhoneNumberDeletedDetailed) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *PhoneNumberDeletedDetailed) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetCustomerReference

`func (o *PhoneNumberDeletedDetailed) GetCustomerReference() string`

GetCustomerReference returns the CustomerReference field if non-nil, zero value otherwise.

### GetCustomerReferenceOk

`func (o *PhoneNumberDeletedDetailed) GetCustomerReferenceOk() (*string, bool)`

GetCustomerReferenceOk returns a tuple with the CustomerReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerReference

`func (o *PhoneNumberDeletedDetailed) SetCustomerReference(v string)`

SetCustomerReference sets CustomerReference field to given value.

### HasCustomerReference

`func (o *PhoneNumberDeletedDetailed) HasCustomerReference() bool`

HasCustomerReference returns a boolean if a field has been set.

### GetMessagingProfileId

`func (o *PhoneNumberDeletedDetailed) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *PhoneNumberDeletedDetailed) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *PhoneNumberDeletedDetailed) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.

### HasMessagingProfileId

`func (o *PhoneNumberDeletedDetailed) HasMessagingProfileId() bool`

HasMessagingProfileId returns a boolean if a field has been set.

### GetMessagingProfileName

`func (o *PhoneNumberDeletedDetailed) GetMessagingProfileName() string`

GetMessagingProfileName returns the MessagingProfileName field if non-nil, zero value otherwise.

### GetMessagingProfileNameOk

`func (o *PhoneNumberDeletedDetailed) GetMessagingProfileNameOk() (*string, bool)`

GetMessagingProfileNameOk returns a tuple with the MessagingProfileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileName

`func (o *PhoneNumberDeletedDetailed) SetMessagingProfileName(v string)`

SetMessagingProfileName sets MessagingProfileName field to given value.

### HasMessagingProfileName

`func (o *PhoneNumberDeletedDetailed) HasMessagingProfileName() bool`

HasMessagingProfileName returns a boolean if a field has been set.

### GetBillingGroupId

`func (o *PhoneNumberDeletedDetailed) GetBillingGroupId() string`

GetBillingGroupId returns the BillingGroupId field if non-nil, zero value otherwise.

### GetBillingGroupIdOk

`func (o *PhoneNumberDeletedDetailed) GetBillingGroupIdOk() (*string, bool)`

GetBillingGroupIdOk returns a tuple with the BillingGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingGroupId

`func (o *PhoneNumberDeletedDetailed) SetBillingGroupId(v string)`

SetBillingGroupId sets BillingGroupId field to given value.

### HasBillingGroupId

`func (o *PhoneNumberDeletedDetailed) HasBillingGroupId() bool`

HasBillingGroupId returns a boolean if a field has been set.

### GetEmergencyEnabled

`func (o *PhoneNumberDeletedDetailed) GetEmergencyEnabled() bool`

GetEmergencyEnabled returns the EmergencyEnabled field if non-nil, zero value otherwise.

### GetEmergencyEnabledOk

`func (o *PhoneNumberDeletedDetailed) GetEmergencyEnabledOk() (*bool, bool)`

GetEmergencyEnabledOk returns a tuple with the EmergencyEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyEnabled

`func (o *PhoneNumberDeletedDetailed) SetEmergencyEnabled(v bool)`

SetEmergencyEnabled sets EmergencyEnabled field to given value.

### HasEmergencyEnabled

`func (o *PhoneNumberDeletedDetailed) HasEmergencyEnabled() bool`

HasEmergencyEnabled returns a boolean if a field has been set.

### GetEmergencyAddressId

`func (o *PhoneNumberDeletedDetailed) GetEmergencyAddressId() string`

GetEmergencyAddressId returns the EmergencyAddressId field if non-nil, zero value otherwise.

### GetEmergencyAddressIdOk

`func (o *PhoneNumberDeletedDetailed) GetEmergencyAddressIdOk() (*string, bool)`

GetEmergencyAddressIdOk returns a tuple with the EmergencyAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmergencyAddressId

`func (o *PhoneNumberDeletedDetailed) SetEmergencyAddressId(v string)`

SetEmergencyAddressId sets EmergencyAddressId field to given value.

### HasEmergencyAddressId

`func (o *PhoneNumberDeletedDetailed) HasEmergencyAddressId() bool`

HasEmergencyAddressId returns a boolean if a field has been set.

### GetCallForwardingEnabled

`func (o *PhoneNumberDeletedDetailed) GetCallForwardingEnabled() bool`

GetCallForwardingEnabled returns the CallForwardingEnabled field if non-nil, zero value otherwise.

### GetCallForwardingEnabledOk

`func (o *PhoneNumberDeletedDetailed) GetCallForwardingEnabledOk() (*bool, bool)`

GetCallForwardingEnabledOk returns a tuple with the CallForwardingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallForwardingEnabled

`func (o *PhoneNumberDeletedDetailed) SetCallForwardingEnabled(v bool)`

SetCallForwardingEnabled sets CallForwardingEnabled field to given value.

### HasCallForwardingEnabled

`func (o *PhoneNumberDeletedDetailed) HasCallForwardingEnabled() bool`

HasCallForwardingEnabled returns a boolean if a field has been set.

### GetCnamListingEnabled

`func (o *PhoneNumberDeletedDetailed) GetCnamListingEnabled() bool`

GetCnamListingEnabled returns the CnamListingEnabled field if non-nil, zero value otherwise.

### GetCnamListingEnabledOk

`func (o *PhoneNumberDeletedDetailed) GetCnamListingEnabledOk() (*bool, bool)`

GetCnamListingEnabledOk returns a tuple with the CnamListingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnamListingEnabled

`func (o *PhoneNumberDeletedDetailed) SetCnamListingEnabled(v bool)`

SetCnamListingEnabled sets CnamListingEnabled field to given value.

### HasCnamListingEnabled

`func (o *PhoneNumberDeletedDetailed) HasCnamListingEnabled() bool`

HasCnamListingEnabled returns a boolean if a field has been set.

### GetCallerIdNameEnabled

`func (o *PhoneNumberDeletedDetailed) GetCallerIdNameEnabled() bool`

GetCallerIdNameEnabled returns the CallerIdNameEnabled field if non-nil, zero value otherwise.

### GetCallerIdNameEnabledOk

`func (o *PhoneNumberDeletedDetailed) GetCallerIdNameEnabledOk() (*bool, bool)`

GetCallerIdNameEnabledOk returns a tuple with the CallerIdNameEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdNameEnabled

`func (o *PhoneNumberDeletedDetailed) SetCallerIdNameEnabled(v bool)`

SetCallerIdNameEnabled sets CallerIdNameEnabled field to given value.

### HasCallerIdNameEnabled

`func (o *PhoneNumberDeletedDetailed) HasCallerIdNameEnabled() bool`

HasCallerIdNameEnabled returns a boolean if a field has been set.

### GetCallRecordingEnabled

`func (o *PhoneNumberDeletedDetailed) GetCallRecordingEnabled() bool`

GetCallRecordingEnabled returns the CallRecordingEnabled field if non-nil, zero value otherwise.

### GetCallRecordingEnabledOk

`func (o *PhoneNumberDeletedDetailed) GetCallRecordingEnabledOk() (*bool, bool)`

GetCallRecordingEnabledOk returns a tuple with the CallRecordingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordingEnabled

`func (o *PhoneNumberDeletedDetailed) SetCallRecordingEnabled(v bool)`

SetCallRecordingEnabled sets CallRecordingEnabled field to given value.

### HasCallRecordingEnabled

`func (o *PhoneNumberDeletedDetailed) HasCallRecordingEnabled() bool`

HasCallRecordingEnabled returns a boolean if a field has been set.

### GetT38FaxGatewayEnabled

`func (o *PhoneNumberDeletedDetailed) GetT38FaxGatewayEnabled() bool`

GetT38FaxGatewayEnabled returns the T38FaxGatewayEnabled field if non-nil, zero value otherwise.

### GetT38FaxGatewayEnabledOk

`func (o *PhoneNumberDeletedDetailed) GetT38FaxGatewayEnabledOk() (*bool, bool)`

GetT38FaxGatewayEnabledOk returns a tuple with the T38FaxGatewayEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetT38FaxGatewayEnabled

`func (o *PhoneNumberDeletedDetailed) SetT38FaxGatewayEnabled(v bool)`

SetT38FaxGatewayEnabled sets T38FaxGatewayEnabled field to given value.

### HasT38FaxGatewayEnabled

`func (o *PhoneNumberDeletedDetailed) HasT38FaxGatewayEnabled() bool`

HasT38FaxGatewayEnabled returns a boolean if a field has been set.

### GetPurchasedAt

`func (o *PhoneNumberDeletedDetailed) GetPurchasedAt() string`

GetPurchasedAt returns the PurchasedAt field if non-nil, zero value otherwise.

### GetPurchasedAtOk

`func (o *PhoneNumberDeletedDetailed) GetPurchasedAtOk() (*string, bool)`

GetPurchasedAtOk returns a tuple with the PurchasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchasedAt

`func (o *PhoneNumberDeletedDetailed) SetPurchasedAt(v string)`

SetPurchasedAt sets PurchasedAt field to given value.

### HasPurchasedAt

`func (o *PhoneNumberDeletedDetailed) HasPurchasedAt() bool`

HasPurchasedAt returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PhoneNumberDeletedDetailed) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PhoneNumberDeletedDetailed) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PhoneNumberDeletedDetailed) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PhoneNumberDeletedDetailed) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PhoneNumberDeletedDetailed) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PhoneNumberDeletedDetailed) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PhoneNumberDeletedDetailed) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PhoneNumberDeletedDetailed) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPhoneNumberType

`func (o *PhoneNumberDeletedDetailed) GetPhoneNumberType() string`

GetPhoneNumberType returns the PhoneNumberType field if non-nil, zero value otherwise.

### GetPhoneNumberTypeOk

`func (o *PhoneNumberDeletedDetailed) GetPhoneNumberTypeOk() (*string, bool)`

GetPhoneNumberTypeOk returns a tuple with the PhoneNumberType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumberType

`func (o *PhoneNumberDeletedDetailed) SetPhoneNumberType(v string)`

SetPhoneNumberType sets PhoneNumberType field to given value.

### HasPhoneNumberType

`func (o *PhoneNumberDeletedDetailed) HasPhoneNumberType() bool`

HasPhoneNumberType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


