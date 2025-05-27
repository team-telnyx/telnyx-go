# ModelMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Object** | Pointer to **string** |  | [optional] [default to "model"]
**Created** | **int32** |  | 
**OwnedBy** | **string** |  | 

## Methods

### NewModelMetadata

`func NewModelMetadata(id string, created int32, ownedBy string, ) *ModelMetadata`

NewModelMetadata instantiates a new ModelMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModelMetadataWithDefaults

`func NewModelMetadataWithDefaults() *ModelMetadata`

NewModelMetadataWithDefaults instantiates a new ModelMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ModelMetadata) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ModelMetadata) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ModelMetadata) SetId(v string)`

SetId sets Id field to given value.


### GetObject

`func (o *ModelMetadata) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *ModelMetadata) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *ModelMetadata) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *ModelMetadata) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetCreated

`func (o *ModelMetadata) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ModelMetadata) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ModelMetadata) SetCreated(v int32)`

SetCreated sets Created field to given value.


### GetOwnedBy

`func (o *ModelMetadata) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *ModelMetadata) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *ModelMetadata) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


