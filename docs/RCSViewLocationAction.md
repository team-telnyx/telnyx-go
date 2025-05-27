# RCSViewLocationAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LatLong** | Pointer to [**RCSLatLng**](RCSLatLng.md) |  | [optional] 
**Label** | Pointer to **string** | The label of the pin dropped | [optional] 
**Query** | Pointer to **string** | query string (Android only) | [optional] 

## Methods

### NewRCSViewLocationAction

`func NewRCSViewLocationAction() *RCSViewLocationAction`

NewRCSViewLocationAction instantiates a new RCSViewLocationAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCSViewLocationActionWithDefaults

`func NewRCSViewLocationActionWithDefaults() *RCSViewLocationAction`

NewRCSViewLocationActionWithDefaults instantiates a new RCSViewLocationAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatLong

`func (o *RCSViewLocationAction) GetLatLong() RCSLatLng`

GetLatLong returns the LatLong field if non-nil, zero value otherwise.

### GetLatLongOk

`func (o *RCSViewLocationAction) GetLatLongOk() (*RCSLatLng, bool)`

GetLatLongOk returns a tuple with the LatLong field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatLong

`func (o *RCSViewLocationAction) SetLatLong(v RCSLatLng)`

SetLatLong sets LatLong field to given value.

### HasLatLong

`func (o *RCSViewLocationAction) HasLatLong() bool`

HasLatLong returns a boolean if a field has been set.

### GetLabel

`func (o *RCSViewLocationAction) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *RCSViewLocationAction) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *RCSViewLocationAction) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *RCSViewLocationAction) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetQuery

`func (o *RCSViewLocationAction) GetQuery() string`

GetQuery returns the Query field if non-nil, zero value otherwise.

### GetQueryOk

`func (o *RCSViewLocationAction) GetQueryOk() (*string, bool)`

GetQueryOk returns a tuple with the Query field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuery

`func (o *RCSViewLocationAction) SetQuery(v string)`

SetQuery sets Query field to given value.

### HasQuery

`func (o *RCSViewLocationAction) HasQuery() bool`

HasQuery returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


