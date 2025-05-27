# ListVerifyProfilesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]VerifyProfileResponse**](VerifyProfileResponse.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewListVerifyProfilesResponse

`func NewListVerifyProfilesResponse(data []VerifyProfileResponse, meta Meta, ) *ListVerifyProfilesResponse`

NewListVerifyProfilesResponse instantiates a new ListVerifyProfilesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVerifyProfilesResponseWithDefaults

`func NewListVerifyProfilesResponseWithDefaults() *ListVerifyProfilesResponse`

NewListVerifyProfilesResponseWithDefaults instantiates a new ListVerifyProfilesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListVerifyProfilesResponse) GetData() []VerifyProfileResponse`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListVerifyProfilesResponse) GetDataOk() (*[]VerifyProfileResponse, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListVerifyProfilesResponse) SetData(v []VerifyProfileResponse)`

SetData sets Data field to given value.


### GetMeta

`func (o *ListVerifyProfilesResponse) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListVerifyProfilesResponse) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListVerifyProfilesResponse) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


