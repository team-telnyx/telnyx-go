# SharePortingOrderRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpiresInSeconds** | Pointer to **int32** | The number of seconds the token will be valid for | [optional] 
**Permissions** | Pointer to **string** | The permissions the token will have | [optional] 

## Methods

### NewSharePortingOrderRequest

`func NewSharePortingOrderRequest() *SharePortingOrderRequest`

NewSharePortingOrderRequest instantiates a new SharePortingOrderRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharePortingOrderRequestWithDefaults

`func NewSharePortingOrderRequestWithDefaults() *SharePortingOrderRequest`

NewSharePortingOrderRequestWithDefaults instantiates a new SharePortingOrderRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpiresInSeconds

`func (o *SharePortingOrderRequest) GetExpiresInSeconds() int32`

GetExpiresInSeconds returns the ExpiresInSeconds field if non-nil, zero value otherwise.

### GetExpiresInSecondsOk

`func (o *SharePortingOrderRequest) GetExpiresInSecondsOk() (*int32, bool)`

GetExpiresInSecondsOk returns a tuple with the ExpiresInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInSeconds

`func (o *SharePortingOrderRequest) SetExpiresInSeconds(v int32)`

SetExpiresInSeconds sets ExpiresInSeconds field to given value.

### HasExpiresInSeconds

`func (o *SharePortingOrderRequest) HasExpiresInSeconds() bool`

HasExpiresInSeconds returns a boolean if a field has been set.

### GetPermissions

`func (o *SharePortingOrderRequest) GetPermissions() string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *SharePortingOrderRequest) GetPermissionsOk() (*string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *SharePortingOrderRequest) SetPermissions(v string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *SharePortingOrderRequest) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


