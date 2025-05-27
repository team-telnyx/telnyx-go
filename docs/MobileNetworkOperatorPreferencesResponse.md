# MobileNetworkOperatorPreferencesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MobileNetworkOperatorId** | Pointer to **string** | The mobile network operator resource identification UUID. | [optional] 
**MobileNetworkOperatorName** | Pointer to **string** | The mobile network operator resource name. | [optional] 
**Priority** | Pointer to **int32** | It determines what is the priority of a specific network operator that should be assumed by a SIM card when connecting to a network. The highest priority is 0, the second highest is 1 and so on. | [optional] 

## Methods

### NewMobileNetworkOperatorPreferencesResponse

`func NewMobileNetworkOperatorPreferencesResponse() *MobileNetworkOperatorPreferencesResponse`

NewMobileNetworkOperatorPreferencesResponse instantiates a new MobileNetworkOperatorPreferencesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMobileNetworkOperatorPreferencesResponseWithDefaults

`func NewMobileNetworkOperatorPreferencesResponseWithDefaults() *MobileNetworkOperatorPreferencesResponse`

NewMobileNetworkOperatorPreferencesResponseWithDefaults instantiates a new MobileNetworkOperatorPreferencesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMobileNetworkOperatorId

`func (o *MobileNetworkOperatorPreferencesResponse) GetMobileNetworkOperatorId() string`

GetMobileNetworkOperatorId returns the MobileNetworkOperatorId field if non-nil, zero value otherwise.

### GetMobileNetworkOperatorIdOk

`func (o *MobileNetworkOperatorPreferencesResponse) GetMobileNetworkOperatorIdOk() (*string, bool)`

GetMobileNetworkOperatorIdOk returns a tuple with the MobileNetworkOperatorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNetworkOperatorId

`func (o *MobileNetworkOperatorPreferencesResponse) SetMobileNetworkOperatorId(v string)`

SetMobileNetworkOperatorId sets MobileNetworkOperatorId field to given value.

### HasMobileNetworkOperatorId

`func (o *MobileNetworkOperatorPreferencesResponse) HasMobileNetworkOperatorId() bool`

HasMobileNetworkOperatorId returns a boolean if a field has been set.

### GetMobileNetworkOperatorName

`func (o *MobileNetworkOperatorPreferencesResponse) GetMobileNetworkOperatorName() string`

GetMobileNetworkOperatorName returns the MobileNetworkOperatorName field if non-nil, zero value otherwise.

### GetMobileNetworkOperatorNameOk

`func (o *MobileNetworkOperatorPreferencesResponse) GetMobileNetworkOperatorNameOk() (*string, bool)`

GetMobileNetworkOperatorNameOk returns a tuple with the MobileNetworkOperatorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNetworkOperatorName

`func (o *MobileNetworkOperatorPreferencesResponse) SetMobileNetworkOperatorName(v string)`

SetMobileNetworkOperatorName sets MobileNetworkOperatorName field to given value.

### HasMobileNetworkOperatorName

`func (o *MobileNetworkOperatorPreferencesResponse) HasMobileNetworkOperatorName() bool`

HasMobileNetworkOperatorName returns a boolean if a field has been set.

### GetPriority

`func (o *MobileNetworkOperatorPreferencesResponse) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *MobileNetworkOperatorPreferencesResponse) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *MobileNetworkOperatorPreferencesResponse) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *MobileNetworkOperatorPreferencesResponse) HasPriority() bool`

HasPriority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


