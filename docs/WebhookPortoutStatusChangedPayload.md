# WebhookPortoutStatusChangedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the port out that was moved. | [optional] 
**Status** | Pointer to **string** | The new status of the port out. | [optional] 
**PhoneNumbers** | Pointer to **[]string** | Phone numbers associated with this port-out order | [optional] 
**CarrierName** | Pointer to **string** | Carrier the number will be ported out to | [optional] 
**Spid** | Pointer to **string** | The new carrier SPID. | [optional] 
**UserId** | Pointer to **string** | Identifies the user that the port-out order belongs to. | [optional] 
**SubscriberName** | Pointer to **string** | The name of the port-out&#39;s end user. | [optional] 
**RejectionReason** | Pointer to **string** | The reason why the order is being rejected by the user. If the order is authorized, this field can be left null | [optional] 
**AttemptedPin** | Pointer to **string** | The PIN that was attempted to be used to authorize the port out. | [optional] 

## Methods

### NewWebhookPortoutStatusChangedPayload

`func NewWebhookPortoutStatusChangedPayload() *WebhookPortoutStatusChangedPayload`

NewWebhookPortoutStatusChangedPayload instantiates a new WebhookPortoutStatusChangedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPortoutStatusChangedPayloadWithDefaults

`func NewWebhookPortoutStatusChangedPayloadWithDefaults() *WebhookPortoutStatusChangedPayload`

NewWebhookPortoutStatusChangedPayloadWithDefaults instantiates a new WebhookPortoutStatusChangedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WebhookPortoutStatusChangedPayload) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WebhookPortoutStatusChangedPayload) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WebhookPortoutStatusChangedPayload) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WebhookPortoutStatusChangedPayload) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStatus

`func (o *WebhookPortoutStatusChangedPayload) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *WebhookPortoutStatusChangedPayload) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *WebhookPortoutStatusChangedPayload) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *WebhookPortoutStatusChangedPayload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *WebhookPortoutStatusChangedPayload) GetPhoneNumbers() []string`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *WebhookPortoutStatusChangedPayload) GetPhoneNumbersOk() (*[]string, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *WebhookPortoutStatusChangedPayload) SetPhoneNumbers(v []string)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *WebhookPortoutStatusChangedPayload) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetCarrierName

`func (o *WebhookPortoutStatusChangedPayload) GetCarrierName() string`

GetCarrierName returns the CarrierName field if non-nil, zero value otherwise.

### GetCarrierNameOk

`func (o *WebhookPortoutStatusChangedPayload) GetCarrierNameOk() (*string, bool)`

GetCarrierNameOk returns a tuple with the CarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierName

`func (o *WebhookPortoutStatusChangedPayload) SetCarrierName(v string)`

SetCarrierName sets CarrierName field to given value.

### HasCarrierName

`func (o *WebhookPortoutStatusChangedPayload) HasCarrierName() bool`

HasCarrierName returns a boolean if a field has been set.

### GetSpid

`func (o *WebhookPortoutStatusChangedPayload) GetSpid() string`

GetSpid returns the Spid field if non-nil, zero value otherwise.

### GetSpidOk

`func (o *WebhookPortoutStatusChangedPayload) GetSpidOk() (*string, bool)`

GetSpidOk returns a tuple with the Spid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpid

`func (o *WebhookPortoutStatusChangedPayload) SetSpid(v string)`

SetSpid sets Spid field to given value.

### HasSpid

`func (o *WebhookPortoutStatusChangedPayload) HasSpid() bool`

HasSpid returns a boolean if a field has been set.

### GetUserId

`func (o *WebhookPortoutStatusChangedPayload) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *WebhookPortoutStatusChangedPayload) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *WebhookPortoutStatusChangedPayload) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *WebhookPortoutStatusChangedPayload) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetSubscriberName

`func (o *WebhookPortoutStatusChangedPayload) GetSubscriberName() string`

GetSubscriberName returns the SubscriberName field if non-nil, zero value otherwise.

### GetSubscriberNameOk

`func (o *WebhookPortoutStatusChangedPayload) GetSubscriberNameOk() (*string, bool)`

GetSubscriberNameOk returns a tuple with the SubscriberName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubscriberName

`func (o *WebhookPortoutStatusChangedPayload) SetSubscriberName(v string)`

SetSubscriberName sets SubscriberName field to given value.

### HasSubscriberName

`func (o *WebhookPortoutStatusChangedPayload) HasSubscriberName() bool`

HasSubscriberName returns a boolean if a field has been set.

### GetRejectionReason

`func (o *WebhookPortoutStatusChangedPayload) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *WebhookPortoutStatusChangedPayload) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *WebhookPortoutStatusChangedPayload) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *WebhookPortoutStatusChangedPayload) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### GetAttemptedPin

`func (o *WebhookPortoutStatusChangedPayload) GetAttemptedPin() string`

GetAttemptedPin returns the AttemptedPin field if non-nil, zero value otherwise.

### GetAttemptedPinOk

`func (o *WebhookPortoutStatusChangedPayload) GetAttemptedPinOk() (*string, bool)`

GetAttemptedPinOk returns a tuple with the AttemptedPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttemptedPin

`func (o *WebhookPortoutStatusChangedPayload) SetAttemptedPin(v string)`

SetAttemptedPin sets AttemptedPin field to given value.

### HasAttemptedPin

`func (o *WebhookPortoutStatusChangedPayload) HasAttemptedPin() bool`

HasAttemptedPin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


