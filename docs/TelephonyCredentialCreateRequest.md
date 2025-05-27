# TelephonyCredentialCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** |  | [optional] 
**Tag** | Pointer to **string** | Tags a credential. A single tag can hold at maximum 1000 credentials. | [optional] 
**ConnectionId** | **string** | Identifies the Credential Connection this credential is associated with. | 
**ExpiresAt** | Pointer to **string** | ISO-8601 formatted date indicating when the credential will expire. | [optional] 

## Methods

### NewTelephonyCredentialCreateRequest

`func NewTelephonyCredentialCreateRequest(connectionId string, ) *TelephonyCredentialCreateRequest`

NewTelephonyCredentialCreateRequest instantiates a new TelephonyCredentialCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelephonyCredentialCreateRequestWithDefaults

`func NewTelephonyCredentialCreateRequestWithDefaults() *TelephonyCredentialCreateRequest`

NewTelephonyCredentialCreateRequestWithDefaults instantiates a new TelephonyCredentialCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *TelephonyCredentialCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TelephonyCredentialCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TelephonyCredentialCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TelephonyCredentialCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTag

`func (o *TelephonyCredentialCreateRequest) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *TelephonyCredentialCreateRequest) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *TelephonyCredentialCreateRequest) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *TelephonyCredentialCreateRequest) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetConnectionId

`func (o *TelephonyCredentialCreateRequest) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *TelephonyCredentialCreateRequest) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *TelephonyCredentialCreateRequest) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.


### GetExpiresAt

`func (o *TelephonyCredentialCreateRequest) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *TelephonyCredentialCreateRequest) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *TelephonyCredentialCreateRequest) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *TelephonyCredentialCreateRequest) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


