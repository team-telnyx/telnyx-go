# CustomerServiceRecordsPostRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CustomerServiceRecordStatusChangedEvent**](CustomerServiceRecordStatusChangedEvent.md) |  | [optional] 
**Meta** | Pointer to [**CallbackWebhookMeta**](CallbackWebhookMeta.md) |  | [optional] 

## Methods

### NewCustomerServiceRecordsPostRequest

`func NewCustomerServiceRecordsPostRequest() *CustomerServiceRecordsPostRequest`

NewCustomerServiceRecordsPostRequest instantiates a new CustomerServiceRecordsPostRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerServiceRecordsPostRequestWithDefaults

`func NewCustomerServiceRecordsPostRequestWithDefaults() *CustomerServiceRecordsPostRequest`

NewCustomerServiceRecordsPostRequestWithDefaults instantiates a new CustomerServiceRecordsPostRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CustomerServiceRecordsPostRequest) GetData() CustomerServiceRecordStatusChangedEvent`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CustomerServiceRecordsPostRequest) GetDataOk() (*CustomerServiceRecordStatusChangedEvent, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CustomerServiceRecordsPostRequest) SetData(v CustomerServiceRecordStatusChangedEvent)`

SetData sets Data field to given value.

### HasData

`func (o *CustomerServiceRecordsPostRequest) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *CustomerServiceRecordsPostRequest) GetMeta() CallbackWebhookMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *CustomerServiceRecordsPostRequest) GetMetaOk() (*CallbackWebhookMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *CustomerServiceRecordsPostRequest) SetMeta(v CallbackWebhookMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *CustomerServiceRecordsPostRequest) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


