# GetUserBalance200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**UserBalance**](UserBalance.md) |  | [optional] 

## Methods

### NewGetUserBalance200Response

`func NewGetUserBalance200Response() *GetUserBalance200Response`

NewGetUserBalance200Response instantiates a new GetUserBalance200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserBalance200ResponseWithDefaults

`func NewGetUserBalance200ResponseWithDefaults() *GetUserBalance200Response`

NewGetUserBalance200ResponseWithDefaults instantiates a new GetUserBalance200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetUserBalance200Response) GetData() UserBalance`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetUserBalance200Response) GetDataOk() (*UserBalance, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetUserBalance200Response) SetData(v UserBalance)`

SetData sets Data field to given value.

### HasData

`func (o *GetUserBalance200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


