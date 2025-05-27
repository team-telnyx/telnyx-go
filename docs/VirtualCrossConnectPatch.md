# VirtualCrossConnectPatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PrimaryEnabled** | Pointer to **bool** | Indicates whether the primary circuit is enabled. Setting this to &#x60;false&#x60; will disable the circuit. | [optional] 
**PrimaryRoutingAnnouncement** | Pointer to **bool** | Whether the primary BGP route is being announced. | [optional] 
**PrimaryCloudIp** | Pointer to **string** | The IP address assigned for your side of the Virtual Cross Connect.&lt;br /&gt;&lt;br /&gt;If none is provided, one will be generated for you.&lt;br /&gt;&lt;br /&gt;This value can not be patched once the VXC has bene provisioned. | [optional] 
**SecondaryEnabled** | Pointer to **bool** | Indicates whether the secondary circuit is enabled. Setting this to &#x60;false&#x60; will disable the circuit. | [optional] 
**SecondaryRoutingAnnouncement** | Pointer to **bool** | Whether the secondary BGP route is being announced. | [optional] 
**SecondaryCloudIp** | Pointer to **string** | The IP address assigned for your side of the Virtual Cross Connect.&lt;br /&gt;&lt;br /&gt;If none is provided, one will be generated for you.&lt;br /&gt;&lt;br /&gt;This value can not be patched once the VXC has bene provisioned. | [optional] 

## Methods

### NewVirtualCrossConnectPatch

`func NewVirtualCrossConnectPatch() *VirtualCrossConnectPatch`

NewVirtualCrossConnectPatch instantiates a new VirtualCrossConnectPatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCrossConnectPatchWithDefaults

`func NewVirtualCrossConnectPatchWithDefaults() *VirtualCrossConnectPatch`

NewVirtualCrossConnectPatchWithDefaults instantiates a new VirtualCrossConnectPatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPrimaryEnabled

`func (o *VirtualCrossConnectPatch) GetPrimaryEnabled() bool`

GetPrimaryEnabled returns the PrimaryEnabled field if non-nil, zero value otherwise.

### GetPrimaryEnabledOk

`func (o *VirtualCrossConnectPatch) GetPrimaryEnabledOk() (*bool, bool)`

GetPrimaryEnabledOk returns a tuple with the PrimaryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEnabled

`func (o *VirtualCrossConnectPatch) SetPrimaryEnabled(v bool)`

SetPrimaryEnabled sets PrimaryEnabled field to given value.

### HasPrimaryEnabled

`func (o *VirtualCrossConnectPatch) HasPrimaryEnabled() bool`

HasPrimaryEnabled returns a boolean if a field has been set.

### GetPrimaryRoutingAnnouncement

`func (o *VirtualCrossConnectPatch) GetPrimaryRoutingAnnouncement() bool`

GetPrimaryRoutingAnnouncement returns the PrimaryRoutingAnnouncement field if non-nil, zero value otherwise.

### GetPrimaryRoutingAnnouncementOk

`func (o *VirtualCrossConnectPatch) GetPrimaryRoutingAnnouncementOk() (*bool, bool)`

GetPrimaryRoutingAnnouncementOk returns a tuple with the PrimaryRoutingAnnouncement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryRoutingAnnouncement

`func (o *VirtualCrossConnectPatch) SetPrimaryRoutingAnnouncement(v bool)`

SetPrimaryRoutingAnnouncement sets PrimaryRoutingAnnouncement field to given value.

### HasPrimaryRoutingAnnouncement

`func (o *VirtualCrossConnectPatch) HasPrimaryRoutingAnnouncement() bool`

HasPrimaryRoutingAnnouncement returns a boolean if a field has been set.

### GetPrimaryCloudIp

`func (o *VirtualCrossConnectPatch) GetPrimaryCloudIp() string`

GetPrimaryCloudIp returns the PrimaryCloudIp field if non-nil, zero value otherwise.

### GetPrimaryCloudIpOk

`func (o *VirtualCrossConnectPatch) GetPrimaryCloudIpOk() (*string, bool)`

GetPrimaryCloudIpOk returns a tuple with the PrimaryCloudIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCloudIp

`func (o *VirtualCrossConnectPatch) SetPrimaryCloudIp(v string)`

SetPrimaryCloudIp sets PrimaryCloudIp field to given value.

### HasPrimaryCloudIp

`func (o *VirtualCrossConnectPatch) HasPrimaryCloudIp() bool`

HasPrimaryCloudIp returns a boolean if a field has been set.

### GetSecondaryEnabled

`func (o *VirtualCrossConnectPatch) GetSecondaryEnabled() bool`

GetSecondaryEnabled returns the SecondaryEnabled field if non-nil, zero value otherwise.

### GetSecondaryEnabledOk

`func (o *VirtualCrossConnectPatch) GetSecondaryEnabledOk() (*bool, bool)`

GetSecondaryEnabledOk returns a tuple with the SecondaryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryEnabled

`func (o *VirtualCrossConnectPatch) SetSecondaryEnabled(v bool)`

SetSecondaryEnabled sets SecondaryEnabled field to given value.

### HasSecondaryEnabled

`func (o *VirtualCrossConnectPatch) HasSecondaryEnabled() bool`

HasSecondaryEnabled returns a boolean if a field has been set.

### GetSecondaryRoutingAnnouncement

`func (o *VirtualCrossConnectPatch) GetSecondaryRoutingAnnouncement() bool`

GetSecondaryRoutingAnnouncement returns the SecondaryRoutingAnnouncement field if non-nil, zero value otherwise.

### GetSecondaryRoutingAnnouncementOk

`func (o *VirtualCrossConnectPatch) GetSecondaryRoutingAnnouncementOk() (*bool, bool)`

GetSecondaryRoutingAnnouncementOk returns a tuple with the SecondaryRoutingAnnouncement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryRoutingAnnouncement

`func (o *VirtualCrossConnectPatch) SetSecondaryRoutingAnnouncement(v bool)`

SetSecondaryRoutingAnnouncement sets SecondaryRoutingAnnouncement field to given value.

### HasSecondaryRoutingAnnouncement

`func (o *VirtualCrossConnectPatch) HasSecondaryRoutingAnnouncement() bool`

HasSecondaryRoutingAnnouncement returns a boolean if a field has been set.

### GetSecondaryCloudIp

`func (o *VirtualCrossConnectPatch) GetSecondaryCloudIp() string`

GetSecondaryCloudIp returns the SecondaryCloudIp field if non-nil, zero value otherwise.

### GetSecondaryCloudIpOk

`func (o *VirtualCrossConnectPatch) GetSecondaryCloudIpOk() (*string, bool)`

GetSecondaryCloudIpOk returns a tuple with the SecondaryCloudIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCloudIp

`func (o *VirtualCrossConnectPatch) SetSecondaryCloudIp(v string)`

SetSecondaryCloudIp sets SecondaryCloudIp field to given value.

### HasSecondaryCloudIp

`func (o *VirtualCrossConnectPatch) HasSecondaryCloudIp() bool`

HasSecondaryCloudIp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


