# CreateRoomClientTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenTtlSecs** | Pointer to **int32** | The time to live in seconds of the Client Token, after that time the Client Token is invalid and can&#39;t be used to join a Room. | [optional] [default to 600]
**RefreshTokenTtlSecs** | Pointer to **int32** | The time to live in seconds of the Refresh Token, after that time the Refresh Token is invalid and can&#39;t be used to refresh Client Token. | [optional] [default to 3600]

## Methods

### NewCreateRoomClientTokenRequest

`func NewCreateRoomClientTokenRequest() *CreateRoomClientTokenRequest`

NewCreateRoomClientTokenRequest instantiates a new CreateRoomClientTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateRoomClientTokenRequestWithDefaults

`func NewCreateRoomClientTokenRequestWithDefaults() *CreateRoomClientTokenRequest`

NewCreateRoomClientTokenRequestWithDefaults instantiates a new CreateRoomClientTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenTtlSecs

`func (o *CreateRoomClientTokenRequest) GetTokenTtlSecs() int32`

GetTokenTtlSecs returns the TokenTtlSecs field if non-nil, zero value otherwise.

### GetTokenTtlSecsOk

`func (o *CreateRoomClientTokenRequest) GetTokenTtlSecsOk() (*int32, bool)`

GetTokenTtlSecsOk returns a tuple with the TokenTtlSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtlSecs

`func (o *CreateRoomClientTokenRequest) SetTokenTtlSecs(v int32)`

SetTokenTtlSecs sets TokenTtlSecs field to given value.

### HasTokenTtlSecs

`func (o *CreateRoomClientTokenRequest) HasTokenTtlSecs() bool`

HasTokenTtlSecs returns a boolean if a field has been set.

### GetRefreshTokenTtlSecs

`func (o *CreateRoomClientTokenRequest) GetRefreshTokenTtlSecs() int32`

GetRefreshTokenTtlSecs returns the RefreshTokenTtlSecs field if non-nil, zero value otherwise.

### GetRefreshTokenTtlSecsOk

`func (o *CreateRoomClientTokenRequest) GetRefreshTokenTtlSecsOk() (*int32, bool)`

GetRefreshTokenTtlSecsOk returns a tuple with the RefreshTokenTtlSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshTokenTtlSecs

`func (o *CreateRoomClientTokenRequest) SetRefreshTokenTtlSecs(v int32)`

SetRefreshTokenTtlSecs sets RefreshTokenTtlSecs field to given value.

### HasRefreshTokenTtlSecs

`func (o *CreateRoomClientTokenRequest) HasRefreshTokenTtlSecs() bool`

HasRefreshTokenTtlSecs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


