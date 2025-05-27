# SIMCardPublicIP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** |  | [optional] [readonly] 
**RegionCode** | Pointer to **string** |  | [optional] [readonly] 
**SimCardId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] [readonly] [default to "ipv4"]
**Ip** | Pointer to **string** | The provisioned IP address. This attribute will only be available when underlying resource status is in a \&quot;provisioned\&quot; status. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 

## Methods

### NewSIMCardPublicIP

`func NewSIMCardPublicIP() *SIMCardPublicIP`

NewSIMCardPublicIP instantiates a new SIMCardPublicIP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSIMCardPublicIPWithDefaults

`func NewSIMCardPublicIPWithDefaults() *SIMCardPublicIP`

NewSIMCardPublicIPWithDefaults instantiates a new SIMCardPublicIP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *SIMCardPublicIP) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *SIMCardPublicIP) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *SIMCardPublicIP) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *SIMCardPublicIP) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetRegionCode

`func (o *SIMCardPublicIP) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *SIMCardPublicIP) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *SIMCardPublicIP) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *SIMCardPublicIP) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetSimCardId

`func (o *SIMCardPublicIP) GetSimCardId() string`

GetSimCardId returns the SimCardId field if non-nil, zero value otherwise.

### GetSimCardIdOk

`func (o *SIMCardPublicIP) GetSimCardIdOk() (*string, bool)`

GetSimCardIdOk returns a tuple with the SimCardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimCardId

`func (o *SIMCardPublicIP) SetSimCardId(v string)`

SetSimCardId sets SimCardId field to given value.

### HasSimCardId

`func (o *SIMCardPublicIP) HasSimCardId() bool`

HasSimCardId returns a boolean if a field has been set.

### GetType

`func (o *SIMCardPublicIP) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SIMCardPublicIP) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SIMCardPublicIP) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SIMCardPublicIP) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIp

`func (o *SIMCardPublicIP) GetIp() string`

GetIp returns the Ip field if non-nil, zero value otherwise.

### GetIpOk

`func (o *SIMCardPublicIP) GetIpOk() (*string, bool)`

GetIpOk returns a tuple with the Ip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIp

`func (o *SIMCardPublicIP) SetIp(v string)`

SetIp sets Ip field to given value.

### HasIp

`func (o *SIMCardPublicIP) HasIp() bool`

HasIp returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SIMCardPublicIP) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SIMCardPublicIP) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SIMCardPublicIP) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SIMCardPublicIP) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SIMCardPublicIP) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SIMCardPublicIP) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SIMCardPublicIP) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SIMCardPublicIP) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


