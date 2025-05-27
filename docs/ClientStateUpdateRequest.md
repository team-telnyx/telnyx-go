# ClientStateUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientState** | **string** | Use this field to add state to every subsequent webhook. It must be a valid Base-64 encoded string. | 

## Methods

### NewClientStateUpdateRequest

`func NewClientStateUpdateRequest(clientState string, ) *ClientStateUpdateRequest`

NewClientStateUpdateRequest instantiates a new ClientStateUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientStateUpdateRequestWithDefaults

`func NewClientStateUpdateRequestWithDefaults() *ClientStateUpdateRequest`

NewClientStateUpdateRequestWithDefaults instantiates a new ClientStateUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientState

`func (o *ClientStateUpdateRequest) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *ClientStateUpdateRequest) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *ClientStateUpdateRequest) SetClientState(v string)`

SetClientState sets ClientState field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


