# SIMCardOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**Quantity** | Pointer to **int32** | The amount of SIM cards requested in the SIM card order. | [optional] 
**Cost** | Pointer to [**SIMCardOrderCost**](SIMCardOrderCost.md) |  | [optional] 
**OrderAddress** | Pointer to [**SIMCardOrderOrderAddress**](SIMCardOrderOrderAddress.md) |  | [optional] 
**TrackingUrl** | Pointer to **string** | The URL used to get tracking information about the order. | [optional] 
**Status** | Pointer to **string** | The current status of the SIM Card order.&lt;ul&gt; &lt;li&gt;&lt;code&gt;pending&lt;/code&gt; - the order is waiting to be processed.&lt;/li&gt; &lt;li&gt;&lt;code&gt;processing&lt;/code&gt; - the order is currently being processed.&lt;/li&gt; &lt;li&gt;&lt;code&gt;ready_to_ship&lt;/code&gt; - the order is ready to be shipped to the specified &lt;b&gt;address&lt;/b&gt;.&lt;/li&gt; &lt;li&gt;&lt;code&gt;shipped&lt;/code&gt; - the order was shipped and is on its way to be delivered to the specified &lt;b&gt;address&lt;/b&gt;.&lt;/li&gt; &lt;li&gt;&lt;code&gt;delivered&lt;/code&gt; - the order was delivered to the specified &lt;b&gt;address&lt;/b&gt;.&lt;/li&gt; &lt;li&gt;&lt;code&gt;canceled&lt;/code&gt; - the order was canceled.&lt;/li&gt; &lt;/ul&gt; | [optional] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was last created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was last updated. | [optional] [readonly] 

## Methods

### NewSIMCardOrder

`func NewSIMCardOrder() *SIMCardOrder`

NewSIMCardOrder instantiates a new SIMCardOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSIMCardOrderWithDefaults

`func NewSIMCardOrderWithDefaults() *SIMCardOrder`

NewSIMCardOrderWithDefaults instantiates a new SIMCardOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SIMCardOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SIMCardOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SIMCardOrder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SIMCardOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *SIMCardOrder) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *SIMCardOrder) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *SIMCardOrder) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *SIMCardOrder) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetQuantity

`func (o *SIMCardOrder) GetQuantity() int32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *SIMCardOrder) GetQuantityOk() (*int32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *SIMCardOrder) SetQuantity(v int32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *SIMCardOrder) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetCost

`func (o *SIMCardOrder) GetCost() SIMCardOrderCost`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *SIMCardOrder) GetCostOk() (*SIMCardOrderCost, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *SIMCardOrder) SetCost(v SIMCardOrderCost)`

SetCost sets Cost field to given value.

### HasCost

`func (o *SIMCardOrder) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetOrderAddress

`func (o *SIMCardOrder) GetOrderAddress() SIMCardOrderOrderAddress`

GetOrderAddress returns the OrderAddress field if non-nil, zero value otherwise.

### GetOrderAddressOk

`func (o *SIMCardOrder) GetOrderAddressOk() (*SIMCardOrderOrderAddress, bool)`

GetOrderAddressOk returns a tuple with the OrderAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderAddress

`func (o *SIMCardOrder) SetOrderAddress(v SIMCardOrderOrderAddress)`

SetOrderAddress sets OrderAddress field to given value.

### HasOrderAddress

`func (o *SIMCardOrder) HasOrderAddress() bool`

HasOrderAddress returns a boolean if a field has been set.

### GetTrackingUrl

`func (o *SIMCardOrder) GetTrackingUrl() string`

GetTrackingUrl returns the TrackingUrl field if non-nil, zero value otherwise.

### GetTrackingUrlOk

`func (o *SIMCardOrder) GetTrackingUrlOk() (*string, bool)`

GetTrackingUrlOk returns a tuple with the TrackingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingUrl

`func (o *SIMCardOrder) SetTrackingUrl(v string)`

SetTrackingUrl sets TrackingUrl field to given value.

### HasTrackingUrl

`func (o *SIMCardOrder) HasTrackingUrl() bool`

HasTrackingUrl returns a boolean if a field has been set.

### GetStatus

`func (o *SIMCardOrder) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SIMCardOrder) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SIMCardOrder) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SIMCardOrder) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SIMCardOrder) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SIMCardOrder) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SIMCardOrder) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SIMCardOrder) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SIMCardOrder) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SIMCardOrder) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SIMCardOrder) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SIMCardOrder) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


