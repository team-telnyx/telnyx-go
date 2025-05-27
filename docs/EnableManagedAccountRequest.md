# EnableManagedAccountRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReenableAllConnections** | Pointer to **bool** | When true, all connections owned by this managed account will automatically be re-enabled. Note: Any connections that do not pass validations will not be re-enabled. | [optional] [default to false]

## Methods

### NewEnableManagedAccountRequest

`func NewEnableManagedAccountRequest() *EnableManagedAccountRequest`

NewEnableManagedAccountRequest instantiates a new EnableManagedAccountRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableManagedAccountRequestWithDefaults

`func NewEnableManagedAccountRequestWithDefaults() *EnableManagedAccountRequest`

NewEnableManagedAccountRequestWithDefaults instantiates a new EnableManagedAccountRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReenableAllConnections

`func (o *EnableManagedAccountRequest) GetReenableAllConnections() bool`

GetReenableAllConnections returns the ReenableAllConnections field if non-nil, zero value otherwise.

### GetReenableAllConnectionsOk

`func (o *EnableManagedAccountRequest) GetReenableAllConnectionsOk() (*bool, bool)`

GetReenableAllConnectionsOk returns a tuple with the ReenableAllConnections field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReenableAllConnections

`func (o *EnableManagedAccountRequest) SetReenableAllConnections(v bool)`

SetReenableAllConnections sets ReenableAllConnections field to given value.

### HasReenableAllConnections

`func (o *EnableManagedAccountRequest) HasReenableAllConnections() bool`

HasReenableAllConnections returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


