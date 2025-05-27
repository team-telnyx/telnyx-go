# ListVerificationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Verification**](Verification.md) |  | 
**Meta** | [**Meta**](Meta.md) |  | 

## Methods

### NewListVerificationsResponse

`func NewListVerificationsResponse(data []Verification, meta Meta, ) *ListVerificationsResponse`

NewListVerificationsResponse instantiates a new ListVerificationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVerificationsResponseWithDefaults

`func NewListVerificationsResponseWithDefaults() *ListVerificationsResponse`

NewListVerificationsResponseWithDefaults instantiates a new ListVerificationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *ListVerificationsResponse) GetData() []Verification`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListVerificationsResponse) GetDataOk() (*[]Verification, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListVerificationsResponse) SetData(v []Verification)`

SetData sets Data field to given value.


### GetMeta

`func (o *ListVerificationsResponse) GetMeta() Meta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListVerificationsResponse) GetMetaOk() (*Meta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListVerificationsResponse) SetMeta(v Meta)`

SetMeta sets Meta field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


