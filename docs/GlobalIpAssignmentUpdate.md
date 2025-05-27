# GlobalIpAssignmentUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 
**GlobalIpId** | Pointer to **interface{}** |  | [optional] [readonly] 
**WireguardPeerId** | Pointer to **interface{}** |  | [optional] [readonly] 
**Status** | Pointer to [**InterfaceStatus**](InterfaceStatus.md) |  | [optional] 
**IsConnected** | Pointer to **bool** | Wireguard peer is connected. | [optional] [readonly] 
**IsInMaintenance** | Pointer to **bool** | Enable/disable BGP announcement. | [optional] 
**IsAnnounced** | Pointer to **bool** | Status of BGP announcement. | [optional] [readonly] 

## Methods

### NewGlobalIpAssignmentUpdate

`func NewGlobalIpAssignmentUpdate() *GlobalIpAssignmentUpdate`

NewGlobalIpAssignmentUpdate instantiates a new GlobalIpAssignmentUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGlobalIpAssignmentUpdateWithDefaults

`func NewGlobalIpAssignmentUpdateWithDefaults() *GlobalIpAssignmentUpdate`

NewGlobalIpAssignmentUpdateWithDefaults instantiates a new GlobalIpAssignmentUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GlobalIpAssignmentUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GlobalIpAssignmentUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GlobalIpAssignmentUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GlobalIpAssignmentUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *GlobalIpAssignmentUpdate) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *GlobalIpAssignmentUpdate) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *GlobalIpAssignmentUpdate) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *GlobalIpAssignmentUpdate) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *GlobalIpAssignmentUpdate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GlobalIpAssignmentUpdate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GlobalIpAssignmentUpdate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *GlobalIpAssignmentUpdate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *GlobalIpAssignmentUpdate) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GlobalIpAssignmentUpdate) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GlobalIpAssignmentUpdate) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *GlobalIpAssignmentUpdate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetGlobalIpId

`func (o *GlobalIpAssignmentUpdate) GetGlobalIpId() interface{}`

GetGlobalIpId returns the GlobalIpId field if non-nil, zero value otherwise.

### GetGlobalIpIdOk

`func (o *GlobalIpAssignmentUpdate) GetGlobalIpIdOk() (*interface{}, bool)`

GetGlobalIpIdOk returns a tuple with the GlobalIpId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalIpId

`func (o *GlobalIpAssignmentUpdate) SetGlobalIpId(v interface{})`

SetGlobalIpId sets GlobalIpId field to given value.

### HasGlobalIpId

`func (o *GlobalIpAssignmentUpdate) HasGlobalIpId() bool`

HasGlobalIpId returns a boolean if a field has been set.

### SetGlobalIpIdNil

`func (o *GlobalIpAssignmentUpdate) SetGlobalIpIdNil(b bool)`

 SetGlobalIpIdNil sets the value for GlobalIpId to be an explicit nil

### UnsetGlobalIpId
`func (o *GlobalIpAssignmentUpdate) UnsetGlobalIpId()`

UnsetGlobalIpId ensures that no value is present for GlobalIpId, not even an explicit nil
### GetWireguardPeerId

`func (o *GlobalIpAssignmentUpdate) GetWireguardPeerId() interface{}`

GetWireguardPeerId returns the WireguardPeerId field if non-nil, zero value otherwise.

### GetWireguardPeerIdOk

`func (o *GlobalIpAssignmentUpdate) GetWireguardPeerIdOk() (*interface{}, bool)`

GetWireguardPeerIdOk returns a tuple with the WireguardPeerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWireguardPeerId

`func (o *GlobalIpAssignmentUpdate) SetWireguardPeerId(v interface{})`

SetWireguardPeerId sets WireguardPeerId field to given value.

### HasWireguardPeerId

`func (o *GlobalIpAssignmentUpdate) HasWireguardPeerId() bool`

HasWireguardPeerId returns a boolean if a field has been set.

### SetWireguardPeerIdNil

`func (o *GlobalIpAssignmentUpdate) SetWireguardPeerIdNil(b bool)`

 SetWireguardPeerIdNil sets the value for WireguardPeerId to be an explicit nil

### UnsetWireguardPeerId
`func (o *GlobalIpAssignmentUpdate) UnsetWireguardPeerId()`

UnsetWireguardPeerId ensures that no value is present for WireguardPeerId, not even an explicit nil
### GetStatus

`func (o *GlobalIpAssignmentUpdate) GetStatus() InterfaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GlobalIpAssignmentUpdate) GetStatusOk() (*InterfaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GlobalIpAssignmentUpdate) SetStatus(v InterfaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GlobalIpAssignmentUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetIsConnected

`func (o *GlobalIpAssignmentUpdate) GetIsConnected() bool`

GetIsConnected returns the IsConnected field if non-nil, zero value otherwise.

### GetIsConnectedOk

`func (o *GlobalIpAssignmentUpdate) GetIsConnectedOk() (*bool, bool)`

GetIsConnectedOk returns a tuple with the IsConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConnected

`func (o *GlobalIpAssignmentUpdate) SetIsConnected(v bool)`

SetIsConnected sets IsConnected field to given value.

### HasIsConnected

`func (o *GlobalIpAssignmentUpdate) HasIsConnected() bool`

HasIsConnected returns a boolean if a field has been set.

### GetIsInMaintenance

`func (o *GlobalIpAssignmentUpdate) GetIsInMaintenance() bool`

GetIsInMaintenance returns the IsInMaintenance field if non-nil, zero value otherwise.

### GetIsInMaintenanceOk

`func (o *GlobalIpAssignmentUpdate) GetIsInMaintenanceOk() (*bool, bool)`

GetIsInMaintenanceOk returns a tuple with the IsInMaintenance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsInMaintenance

`func (o *GlobalIpAssignmentUpdate) SetIsInMaintenance(v bool)`

SetIsInMaintenance sets IsInMaintenance field to given value.

### HasIsInMaintenance

`func (o *GlobalIpAssignmentUpdate) HasIsInMaintenance() bool`

HasIsInMaintenance returns a boolean if a field has been set.

### GetIsAnnounced

`func (o *GlobalIpAssignmentUpdate) GetIsAnnounced() bool`

GetIsAnnounced returns the IsAnnounced field if non-nil, zero value otherwise.

### GetIsAnnouncedOk

`func (o *GlobalIpAssignmentUpdate) GetIsAnnouncedOk() (*bool, bool)`

GetIsAnnouncedOk returns a tuple with the IsAnnounced field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAnnounced

`func (o *GlobalIpAssignmentUpdate) SetIsAnnounced(v bool)`

SetIsAnnounced sets IsAnnounced field to given value.

### HasIsAnnounced

`func (o *GlobalIpAssignmentUpdate) HasIsAnnounced() bool`

HasIsAnnounced returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


