# AuditLogList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]AuditLog**](AuditLog.md) |  | [optional] 
**Meta** | Pointer to [**PaginationMeta**](PaginationMeta.md) |  | [optional] 

## Methods

### NewAuditLogList

`func NewAuditLogList() *AuditLogList`

NewAuditLogList instantiates a new AuditLogList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogListWithDefaults

`func NewAuditLogListWithDefaults() *AuditLogList`

NewAuditLogListWithDefaults instantiates a new AuditLogList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AuditLogList) GetData() []AuditLog`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuditLogList) GetDataOk() (*[]AuditLog, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuditLogList) SetData(v []AuditLog)`

SetData sets Data field to given value.

### HasData

`func (o *AuditLogList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMeta

`func (o *AuditLogList) GetMeta() PaginationMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AuditLogList) GetMetaOk() (*PaginationMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AuditLogList) SetMeta(v PaginationMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AuditLogList) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


