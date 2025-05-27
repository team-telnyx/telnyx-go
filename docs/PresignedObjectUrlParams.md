# PresignedObjectUrlParams

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ttl** | Pointer to **int32** | The time to live of the token in seconds | [optional] 

## Methods

### NewPresignedObjectUrlParams

`func NewPresignedObjectUrlParams() *PresignedObjectUrlParams`

NewPresignedObjectUrlParams instantiates a new PresignedObjectUrlParams object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPresignedObjectUrlParamsWithDefaults

`func NewPresignedObjectUrlParamsWithDefaults() *PresignedObjectUrlParams`

NewPresignedObjectUrlParamsWithDefaults instantiates a new PresignedObjectUrlParams object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTtl

`func (o *PresignedObjectUrlParams) GetTtl() int32`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *PresignedObjectUrlParams) GetTtlOk() (*int32, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *PresignedObjectUrlParams) SetTtl(v int32)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *PresignedObjectUrlParams) HasTtl() bool`

HasTtl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


