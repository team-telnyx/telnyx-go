# RefreshRoomClientTokenRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TokenTtlSecs** | Pointer to **int32** | The time to live in seconds of the Client Token, after that time the Client Token is invalid and can&#39;t be used to join a Room. | [optional] [default to 600]
**RefreshToken** | **string** |  | 

## Methods

### NewRefreshRoomClientTokenRequest

`func NewRefreshRoomClientTokenRequest(refreshToken string, ) *RefreshRoomClientTokenRequest`

NewRefreshRoomClientTokenRequest instantiates a new RefreshRoomClientTokenRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshRoomClientTokenRequestWithDefaults

`func NewRefreshRoomClientTokenRequestWithDefaults() *RefreshRoomClientTokenRequest`

NewRefreshRoomClientTokenRequestWithDefaults instantiates a new RefreshRoomClientTokenRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTokenTtlSecs

`func (o *RefreshRoomClientTokenRequest) GetTokenTtlSecs() int32`

GetTokenTtlSecs returns the TokenTtlSecs field if non-nil, zero value otherwise.

### GetTokenTtlSecsOk

`func (o *RefreshRoomClientTokenRequest) GetTokenTtlSecsOk() (*int32, bool)`

GetTokenTtlSecsOk returns a tuple with the TokenTtlSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenTtlSecs

`func (o *RefreshRoomClientTokenRequest) SetTokenTtlSecs(v int32)`

SetTokenTtlSecs sets TokenTtlSecs field to given value.

### HasTokenTtlSecs

`func (o *RefreshRoomClientTokenRequest) HasTokenTtlSecs() bool`

HasTokenTtlSecs returns a boolean if a field has been set.

### GetRefreshToken

`func (o *RefreshRoomClientTokenRequest) GetRefreshToken() string`

GetRefreshToken returns the RefreshToken field if non-nil, zero value otherwise.

### GetRefreshTokenOk

`func (o *RefreshRoomClientTokenRequest) GetRefreshTokenOk() (*string, bool)`

GetRefreshTokenOk returns a tuple with the RefreshToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefreshToken

`func (o *RefreshRoomClientTokenRequest) SetRefreshToken(v string)`

SetRefreshToken sets RefreshToken field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


