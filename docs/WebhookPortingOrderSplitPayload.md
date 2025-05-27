# WebhookPortingOrderSplitPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to [**WebhookPortingOrderSplitPayloadFrom**](WebhookPortingOrderSplitPayloadFrom.md) |  | [optional] 
**To** | Pointer to [**WebhookPortingOrderSplitPayloadTo**](WebhookPortingOrderSplitPayloadTo.md) |  | [optional] 
**PortingPhoneNumbers** | Pointer to [**[]WebhookPortingOrderSplitPayloadPortingPhoneNumbersInner**](WebhookPortingOrderSplitPayloadPortingPhoneNumbersInner.md) | The list of porting phone numbers that were moved to the new porting order. | [optional] 

## Methods

### NewWebhookPortingOrderSplitPayload

`func NewWebhookPortingOrderSplitPayload() *WebhookPortingOrderSplitPayload`

NewWebhookPortingOrderSplitPayload instantiates a new WebhookPortingOrderSplitPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWebhookPortingOrderSplitPayloadWithDefaults

`func NewWebhookPortingOrderSplitPayloadWithDefaults() *WebhookPortingOrderSplitPayload`

NewWebhookPortingOrderSplitPayloadWithDefaults instantiates a new WebhookPortingOrderSplitPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *WebhookPortingOrderSplitPayload) GetFrom() WebhookPortingOrderSplitPayloadFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *WebhookPortingOrderSplitPayload) GetFromOk() (*WebhookPortingOrderSplitPayloadFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *WebhookPortingOrderSplitPayload) SetFrom(v WebhookPortingOrderSplitPayloadFrom)`

SetFrom sets From field to given value.

### HasFrom

`func (o *WebhookPortingOrderSplitPayload) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *WebhookPortingOrderSplitPayload) GetTo() WebhookPortingOrderSplitPayloadTo`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *WebhookPortingOrderSplitPayload) GetToOk() (*WebhookPortingOrderSplitPayloadTo, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *WebhookPortingOrderSplitPayload) SetTo(v WebhookPortingOrderSplitPayloadTo)`

SetTo sets To field to given value.

### HasTo

`func (o *WebhookPortingOrderSplitPayload) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetPortingPhoneNumbers

`func (o *WebhookPortingOrderSplitPayload) GetPortingPhoneNumbers() []WebhookPortingOrderSplitPayloadPortingPhoneNumbersInner`

GetPortingPhoneNumbers returns the PortingPhoneNumbers field if non-nil, zero value otherwise.

### GetPortingPhoneNumbersOk

`func (o *WebhookPortingOrderSplitPayload) GetPortingPhoneNumbersOk() (*[]WebhookPortingOrderSplitPayloadPortingPhoneNumbersInner, bool)`

GetPortingPhoneNumbersOk returns a tuple with the PortingPhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortingPhoneNumbers

`func (o *WebhookPortingOrderSplitPayload) SetPortingPhoneNumbers(v []WebhookPortingOrderSplitPayloadPortingPhoneNumbersInner)`

SetPortingPhoneNumbers sets PortingPhoneNumbers field to given value.

### HasPortingPhoneNumbers

`func (o *WebhookPortingOrderSplitPayload) HasPortingPhoneNumbers() bool`

HasPortingPhoneNumbers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


