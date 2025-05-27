# TelephonyCredentialUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** | Tags a credential. A single tag can hold at maximum 1000 credentials. | [optional] 
**ConnectionId** | Pointer to **string** | Identifies the Credential Connection this credential is associated with. | [optional] 
**ExpiresAt** | Pointer to **string** | ISO-8601 formatted date indicating when the credential will expire. | [optional] 

## Methods

### NewTelephonyCredentialUpdateRequest

`func NewTelephonyCredentialUpdateRequest() *TelephonyCredentialUpdateRequest`

NewTelephonyCredentialUpdateRequest instantiates a new TelephonyCredentialUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelephonyCredentialUpdateRequestWithDefaults

`func NewTelephonyCredentialUpdateRequestWithDefaults() *TelephonyCredentialUpdateRequest`

NewTelephonyCredentialUpdateRequestWithDefaults instantiates a new TelephonyCredentialUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TelephonyCredentialUpdateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelephonyCredentialUpdateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelephonyCredentialUpdateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TelephonyCredentialUpdateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTag

`func (o *TelephonyCredentialUpdateRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *TelephonyCredentialUpdateRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *TelephonyCredentialUpdateRequest) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *TelephonyCredentialUpdateRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetConnectionId

`func (o *TelephonyCredentialUpdateRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *TelephonyCredentialUpdateRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *TelephonyCredentialUpdateRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *TelephonyCredentialUpdateRequest) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetExpiresAt

`func (o *TelephonyCredentialUpdateRequest) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TelephonyCredentialUpdateRequest) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TelephonyCredentialUpdateRequest) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *TelephonyCredentialUpdateRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


