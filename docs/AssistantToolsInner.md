# AssistantToolsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Webhook** | [**WebhookToolParams**](WebhookToolParams.md) |  | 
**Retrieval** | [**BucketIds**](BucketIds.md) |  | 
**Hangup** | [**HangupToolParams**](HangupToolParams.md) |  | 
**Transfer** | [**TransferToolParams**](TransferToolParams.md) |  | 
**SendDtmf** | **map[string]interface{}** |  | 
**BookAppointment** | [**BookAppointmentToolParams**](BookAppointmentToolParams.md) |  | 
**CheckAvailability** | [**CheckAvailabilityToolParams**](CheckAvailabilityToolParams.md) |  | 

## Methods

### NewAssistantToolsInner

`func NewAssistantToolsInner(type_ string, webhook WebhookToolParams, retrieval BucketIds, hangup HangupToolParams, transfer TransferToolParams, sendDtmf map[string]interface{}, bookAppointment BookAppointmentToolParams, checkAvailability CheckAvailabilityToolParams, ) *AssistantToolsInner`

NewAssistantToolsInner instantiates a new AssistantToolsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssistantToolsInnerWithDefaults

`func NewAssistantToolsInnerWithDefaults() *AssistantToolsInner`

NewAssistantToolsInnerWithDefaults instantiates a new AssistantToolsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AssistantToolsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AssistantToolsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AssistantToolsInner) SetType(v string)`

SetType sets Type field to given value.


### GetWebhook

`func (o *AssistantToolsInner) GetWebhook() WebhookToolParams`

GetWebhook returns the Webhook field if non-nil, zero value otherwise.

### GetWebhookOk

`func (o *AssistantToolsInner) GetWebhookOk() (*WebhookToolParams, bool)`

GetWebhookOk returns a tuple with the Webhook field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhook

`func (o *AssistantToolsInner) SetWebhook(v WebhookToolParams)`

SetWebhook sets Webhook field to given value.


### GetRetrieval

`func (o *AssistantToolsInner) GetRetrieval() BucketIds`

GetRetrieval returns the Retrieval field if non-nil, zero value otherwise.

### GetRetrievalOk

`func (o *AssistantToolsInner) GetRetrievalOk() (*BucketIds, bool)`

GetRetrievalOk returns a tuple with the Retrieval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetrieval

`func (o *AssistantToolsInner) SetRetrieval(v BucketIds)`

SetRetrieval sets Retrieval field to given value.


### GetHangup

`func (o *AssistantToolsInner) GetHangup() HangupToolParams`

GetHangup returns the Hangup field if non-nil, zero value otherwise.

### GetHangupOk

`func (o *AssistantToolsInner) GetHangupOk() (*HangupToolParams, bool)`

GetHangupOk returns a tuple with the Hangup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHangup

`func (o *AssistantToolsInner) SetHangup(v HangupToolParams)`

SetHangup sets Hangup field to given value.


### GetTransfer

`func (o *AssistantToolsInner) GetTransfer() TransferToolParams`

GetTransfer returns the Transfer field if non-nil, zero value otherwise.

### GetTransferOk

`func (o *AssistantToolsInner) GetTransferOk() (*TransferToolParams, bool)`

GetTransferOk returns a tuple with the Transfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransfer

`func (o *AssistantToolsInner) SetTransfer(v TransferToolParams)`

SetTransfer sets Transfer field to given value.


### GetSendDtmf

`func (o *AssistantToolsInner) GetSendDtmf() map[string]interface{}`

GetSendDtmf returns the SendDtmf field if non-nil, zero value otherwise.

### GetSendDtmfOk

`func (o *AssistantToolsInner) GetSendDtmfOk() (*map[string]interface{}, bool)`

GetSendDtmfOk returns a tuple with the SendDtmf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendDtmf

`func (o *AssistantToolsInner) SetSendDtmf(v map[string]interface{})`

SetSendDtmf sets SendDtmf field to given value.


### GetBookAppointment

`func (o *AssistantToolsInner) GetBookAppointment() BookAppointmentToolParams`

GetBookAppointment returns the BookAppointment field if non-nil, zero value otherwise.

### GetBookAppointmentOk

`func (o *AssistantToolsInner) GetBookAppointmentOk() (*BookAppointmentToolParams, bool)`

GetBookAppointmentOk returns a tuple with the BookAppointment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBookAppointment

`func (o *AssistantToolsInner) SetBookAppointment(v BookAppointmentToolParams)`

SetBookAppointment sets BookAppointment field to given value.


### GetCheckAvailability

`func (o *AssistantToolsInner) GetCheckAvailability() CheckAvailabilityToolParams`

GetCheckAvailability returns the CheckAvailability field if non-nil, zero value otherwise.

### GetCheckAvailabilityOk

`func (o *AssistantToolsInner) GetCheckAvailabilityOk() (*CheckAvailabilityToolParams, bool)`

GetCheckAvailabilityOk returns a tuple with the CheckAvailability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckAvailability

`func (o *AssistantToolsInner) SetCheckAvailability(v CheckAvailabilityToolParams)`

SetCheckAvailability sets CheckAvailability field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


