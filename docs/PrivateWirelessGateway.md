# PrivateWirelessGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**NetworkId** | Pointer to **string** | The identification of the related network resource. | [optional] 
**RecordType** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 
**Name** | Pointer to **string** | The private wireless gateway name. | [optional] 
**RegionCode** | Pointer to **string** | The name of the region where the Private Wireless Gateway is deployed. | [optional] 
**Status** | Pointer to [**PrivateWirelessGatewayStatus**](PrivateWirelessGatewayStatus.md) |  | [optional] 
**IpRange** | Pointer to **string** | IP block used to assign IPs to the SIM cards in the Private Wireless Gateway. | [optional] [readonly] 
**AssignedResources** | Pointer to [**[]PWGAssignedResourcesSummary**](PWGAssignedResourcesSummary.md) | A list of the resources that have been assigned to the Private Wireless Gateway. | [optional] 

## Methods

### NewPrivateWirelessGateway

`func NewPrivateWirelessGateway() *PrivateWirelessGateway`

NewPrivateWirelessGateway instantiates a new PrivateWirelessGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrivateWirelessGatewayWithDefaults

`func NewPrivateWirelessGatewayWithDefaults() *PrivateWirelessGateway`

NewPrivateWirelessGatewayWithDefaults instantiates a new PrivateWirelessGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PrivateWirelessGateway) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrivateWirelessGateway) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrivateWirelessGateway) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrivateWirelessGateway) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNetworkId

`func (o *PrivateWirelessGateway) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *PrivateWirelessGateway) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *PrivateWirelessGateway) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *PrivateWirelessGateway) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetRecordType

`func (o *PrivateWirelessGateway) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PrivateWirelessGateway) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PrivateWirelessGateway) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PrivateWirelessGateway) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PrivateWirelessGateway) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PrivateWirelessGateway) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PrivateWirelessGateway) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PrivateWirelessGateway) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PrivateWirelessGateway) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PrivateWirelessGateway) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PrivateWirelessGateway) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PrivateWirelessGateway) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetName

`func (o *PrivateWirelessGateway) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PrivateWirelessGateway) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PrivateWirelessGateway) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PrivateWirelessGateway) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegionCode

`func (o *PrivateWirelessGateway) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *PrivateWirelessGateway) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *PrivateWirelessGateway) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *PrivateWirelessGateway) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetStatus

`func (o *PrivateWirelessGateway) GetStatus() PrivateWirelessGatewayStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrivateWirelessGateway) GetStatusOk() (*PrivateWirelessGatewayStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrivateWirelessGateway) SetStatus(v PrivateWirelessGatewayStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrivateWirelessGateway) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIpRange

`func (o *PrivateWirelessGateway) GetIpRange() string`

GetIpRange returns the IpRange field if non-nil, zero value otherwise.

### GetIpRangeOk

`func (o *PrivateWirelessGateway) GetIpRangeOk() (*string, bool)`

GetIpRangeOk returns a tuple with the IpRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRange

`func (o *PrivateWirelessGateway) SetIpRange(v string)`

SetIpRange sets IpRange field to given value.

### HasIpRange

`func (o *PrivateWirelessGateway) HasIpRange() bool`

HasIpRange returns a boolean if a field has been set.

### GetAssignedResources

`func (o *PrivateWirelessGateway) GetAssignedResources() []PWGAssignedResourcesSummary`

GetAssignedResources returns the AssignedResources field if non-nil, zero value otherwise.

### GetAssignedResourcesOk

`func (o *PrivateWirelessGateway) GetAssignedResourcesOk() (*[]PWGAssignedResourcesSummary, bool)`

GetAssignedResourcesOk returns a tuple with the AssignedResources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedResources

`func (o *PrivateWirelessGateway) SetAssignedResources(v []PWGAssignedResourcesSummary)`

SetAssignedResources sets AssignedResources field to given value.

### HasAssignedResources

`func (o *PrivateWirelessGateway) HasAssignedResources() bool`

HasAssignedResources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


