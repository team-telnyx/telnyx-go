# GetWirelessConnectivityLogs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]WirelessConnectivityLog**](WirelessConnectivityLog.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewGetWirelessConnectivityLogs200Response

`func NewGetWirelessConnectivityLogs200Response() *GetWirelessConnectivityLogs200Response`

NewGetWirelessConnectivityLogs200Response instantiates a new GetWirelessConnectivityLogs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWirelessConnectivityLogs200ResponseWithDefaults

`func NewGetWirelessConnectivityLogs200ResponseWithDefaults() *GetWirelessConnectivityLogs200Response`

NewGetWirelessConnectivityLogs200ResponseWithDefaults instantiates a new GetWirelessConnectivityLogs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetWirelessConnectivityLogs200Response) GetData() []WirelessConnectivityLog`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetWirelessConnectivityLogs200Response) GetDataOk() (*[]WirelessConnectivityLog, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetWirelessConnectivityLogs200Response) SetData(v []WirelessConnectivityLog)`

SetData sets Data field to given value.

### HasData

`func (o *GetWirelessConnectivityLogs200Response) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *GetWirelessConnectivityLogs200Response) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetWirelessConnectivityLogs200Response) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetWirelessConnectivityLogs200Response) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetWirelessConnectivityLogs200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


