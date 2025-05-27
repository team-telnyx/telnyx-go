# PortingOrderSharingToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies this sharing token | [optional] 
**PortingOrderId** | Pointer to **string** | Identifies the porting order resource being shared | [optional] 
**ExpiresInSeconds** | Pointer to **int32** | The number of seconds until the sharing token expires | [optional] 
**Permissions** | Pointer to **[]string** | The permissions granted to the sharing token | [optional] 
**Token** | Pointer to **string** | A signed JWT token that can be used to access the shared resource | [optional] 
**ExpiresAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the sharing token expires. | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] 

## Methods

### NewPortingOrderSharingToken

`func NewPortingOrderSharingToken() *PortingOrderSharingToken`

NewPortingOrderSharingToken instantiates a new PortingOrderSharingToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortingOrderSharingTokenWithDefaults

`func NewPortingOrderSharingTokenWithDefaults() *PortingOrderSharingToken`

NewPortingOrderSharingTokenWithDefaults instantiates a new PortingOrderSharingToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortingOrderSharingToken) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortingOrderSharingToken) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortingOrderSharingToken) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortingOrderSharingToken) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPortingOrderId

`func (o *PortingOrderSharingToken) GetPortingOrderId() string`

GetPortingOrderId returns the PortingOrderId field if non-nil, zero value otherwise.

### GetPortingOrderIdOk

`func (o *PortingOrderSharingToken) GetPortingOrderIdOk() (*string, bool)`

GetPortingOrderIdOk returns a tuple with the PortingOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortingOrderId

`func (o *PortingOrderSharingToken) SetPortingOrderId(v string)`

SetPortingOrderId sets PortingOrderId field to given value.

### HasPortingOrderId

`func (o *PortingOrderSharingToken) HasPortingOrderId() bool`

HasPortingOrderId returns a boolean if a field has been set.

### GetExpiresInSeconds

`func (o *PortingOrderSharingToken) GetExpiresInSeconds() int32`

GetExpiresInSeconds returns the ExpiresInSeconds field if non-nil, zero value otherwise.

### GetExpiresInSecondsOk

`func (o *PortingOrderSharingToken) GetExpiresInSecondsOk() (*int32, bool)`

GetExpiresInSecondsOk returns a tuple with the ExpiresInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInSeconds

`func (o *PortingOrderSharingToken) SetExpiresInSeconds(v int32)`

SetExpiresInSeconds sets ExpiresInSeconds field to given value.

### HasExpiresInSeconds

`func (o *PortingOrderSharingToken) HasExpiresInSeconds() bool`

HasExpiresInSeconds returns a boolean if a field has been set.

### GetPermissions

`func (o *PortingOrderSharingToken) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PortingOrderSharingToken) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PortingOrderSharingToken) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PortingOrderSharingToken) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetToken

`func (o *PortingOrderSharingToken) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *PortingOrderSharingToken) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *PortingOrderSharingToken) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *PortingOrderSharingToken) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetExpiresAt

`func (o *PortingOrderSharingToken) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *PortingOrderSharingToken) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *PortingOrderSharingToken) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *PortingOrderSharingToken) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetRecordType

`func (o *PortingOrderSharingToken) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PortingOrderSharingToken) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PortingOrderSharingToken) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PortingOrderSharingToken) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PortingOrderSharingToken) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortingOrderSharingToken) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortingOrderSharingToken) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PortingOrderSharingToken) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


