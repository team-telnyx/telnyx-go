# GetAllTelephonyCredentialResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]TelephonyCredential**](TelephonyCredential.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewGetAllTelephonyCredentialResponse

`func NewGetAllTelephonyCredentialResponse() *GetAllTelephonyCredentialResponse`

NewGetAllTelephonyCredentialResponse instantiates a new GetAllTelephonyCredentialResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllTelephonyCredentialResponseWithDefaults

`func NewGetAllTelephonyCredentialResponseWithDefaults() *GetAllTelephonyCredentialResponse`

NewGetAllTelephonyCredentialResponseWithDefaults instantiates a new GetAllTelephonyCredentialResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetAllTelephonyCredentialResponse) GetData() []TelephonyCredential`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetAllTelephonyCredentialResponse) GetDataOk() (*[]TelephonyCredential, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetAllTelephonyCredentialResponse) SetData(v []TelephonyCredential)`

SetData sets Data field to given value.

### HasData

`func (o *GetAllTelephonyCredentialResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetAllTelephonyCredentialResponse) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetAllTelephonyCredentialResponse) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetAllTelephonyCredentialResponse) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetAllTelephonyCredentialResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


