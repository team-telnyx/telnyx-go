# GlobalIpAssignment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 
**GlobalIpId** | Pointer to **string** | Global IP ID. | [optional] 
**WireguardPeerId** | Pointer to **string** | Wireguard peer ID. | [optional] 
**Status** | Pointer to [**InterfaceStatus**](InterfaceStatus.md) |  | [optional] 
**IsConnected** | Pointer to **bool** | Wireguard peer is connected. | [optional] [readonly] 
**IsInMaintenance** | Pointer to **bool** | Enable/disable BGP announcement. | [optional] 
**IsAnnounced** | Pointer to **bool** | Status of BGP announcement. | [optional] [readonly] 

## Methods

### NewGlobalIpAssignment

`func NewGlobalIpAssignment() *GlobalIpAssignment`

NewGlobalIpAssignment instantiates a new GlobalIpAssignment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalIpAssignmentWithDefaults

`func NewGlobalIpAssignmentWithDefaults() *GlobalIpAssignment`

NewGlobalIpAssignmentWithDefaults instantiates a new GlobalIpAssignment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GlobalIpAssignment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalIpAssignment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalIpAssignment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GlobalIpAssignment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *GlobalIpAssignment) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *GlobalIpAssignment) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *GlobalIpAssignment) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *GlobalIpAssignment) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GlobalIpAssignment) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GlobalIpAssignment) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GlobalIpAssignment) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GlobalIpAssignment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GlobalIpAssignment) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GlobalIpAssignment) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GlobalIpAssignment) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GlobalIpAssignment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetGlobalIpId

`func (o *GlobalIpAssignment) GetGlobalIpId() string`

GetGlobalIpId returns the GlobalIpId field if non-nil, zero value otherwise.

### GetGlobalIpIdOk

`func (o *GlobalIpAssignment) GetGlobalIpIdOk() (*string, bool)`

GetGlobalIpIdOk returns a tuple with the GlobalIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIpId

`func (o *GlobalIpAssignment) SetGlobalIpId(v string)`

SetGlobalIpId sets GlobalIpId field to given value.

### HasGlobalIpId

`func (o *GlobalIpAssignment) HasGlobalIpId() bool`

HasGlobalIpId returns a boolean if a field has been set.

### GetWireguardPeerId

`func (o *GlobalIpAssignment) GetWireguardPeerId() string`

GetWireguardPeerId returns the WireguardPeerId field if non-nil, zero value otherwise.

### GetWireguardPeerIdOk

`func (o *GlobalIpAssignment) GetWireguardPeerIdOk() (*string, bool)`

GetWireguardPeerIdOk returns a tuple with the WireguardPeerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguardPeerId

`func (o *GlobalIpAssignment) SetWireguardPeerId(v string)`

SetWireguardPeerId sets WireguardPeerId field to given value.

### HasWireguardPeerId

`func (o *GlobalIpAssignment) HasWireguardPeerId() bool`

HasWireguardPeerId returns a boolean if a field has been set.

### GetStatus

`func (o *GlobalIpAssignment) GetStatus() InterfaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GlobalIpAssignment) GetStatusOk() (*InterfaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GlobalIpAssignment) SetStatus(v InterfaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GlobalIpAssignment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsConnected

`func (o *GlobalIpAssignment) GetIsConnected() bool`

GetIsConnected returns the IsConnected field if non-nil, zero value otherwise.

### GetIsConnectedOk

`func (o *GlobalIpAssignment) GetIsConnectedOk() (*bool, bool)`

GetIsConnectedOk returns a tuple with the IsConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConnected

`func (o *GlobalIpAssignment) SetIsConnected(v bool)`

SetIsConnected sets IsConnected field to given value.

### HasIsConnected

`func (o *GlobalIpAssignment) HasIsConnected() bool`

HasIsConnected returns a boolean if a field has been set.

### GetIsInMaintenance

`func (o *GlobalIpAssignment) GetIsInMaintenance() bool`

GetIsInMaintenance returns the IsInMaintenance field if non-nil, zero value otherwise.

### GetIsInMaintenanceOk

`func (o *GlobalIpAssignment) GetIsInMaintenanceOk() (*bool, bool)`

GetIsInMaintenanceOk returns a tuple with the IsInMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInMaintenance

`func (o *GlobalIpAssignment) SetIsInMaintenance(v bool)`

SetIsInMaintenance sets IsInMaintenance field to given value.

### HasIsInMaintenance

`func (o *GlobalIpAssignment) HasIsInMaintenance() bool`

HasIsInMaintenance returns a boolean if a field has been set.

### GetIsAnnounced

`func (o *GlobalIpAssignment) GetIsAnnounced() bool`

GetIsAnnounced returns the IsAnnounced field if non-nil, zero value otherwise.

### GetIsAnnouncedOk

`func (o *GlobalIpAssignment) GetIsAnnouncedOk() (*bool, bool)`

GetIsAnnouncedOk returns a tuple with the IsAnnounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnnounced

`func (o *GlobalIpAssignment) SetIsAnnounced(v bool)`

SetIsAnnounced sets IsAnnounced field to given value.

### HasIsAnnounced

`func (o *GlobalIpAssignment) HasIsAnnounced() bool`

HasIsAnnounced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


