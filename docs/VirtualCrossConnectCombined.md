# VirtualCrossConnectCombined

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 
**NetworkId** | **string** | The id of the network associated with the interface. | 
**Name** | Pointer to **string** | A user specified name for the interface. | [optional] 
**Status** | Pointer to [**InterfaceStatus**](InterfaceStatus.md) |  | [optional] 
**CloudProvider** | **string** | The Virtual Private Cloud with which you would like to establish a cross connect. | 
**CloudProviderRegion** | **string** | The region where your Virtual Private Cloud hosts are located.&lt;br /&gt;&lt;br /&gt;The available regions can be found using the /virtual_cross_connect_regions endpoint. | 
**BgpAsn** | **float32** | The Border Gateway Protocol (BGP) Autonomous System Number (ASN). If null, value will be assigned by Telnyx. | 
**BandwidthMbps** | Pointer to **float32** | The desired throughput in Megabits per Second (Mbps) for your Virtual Cross Connect.&lt;br /&gt;&lt;br /&gt;The available bandwidths can be found using the /virtual_cross_connect_regions endpoint. | [optional] 
**PrimaryEnabled** | Pointer to **bool** | Indicates whether the primary circuit is enabled. Setting this to &#x60;false&#x60; will disable the circuit. | [optional] 
**PrimaryCloudAccountId** | **string** | The identifier for your Virtual Private Cloud. The number will be different based upon your Cloud provider. | 
**PrimaryTelnyxIp** | Pointer to **string** | The IP address assigned to the Telnyx side of the Virtual Cross Connect.&lt;br /&gt;&lt;br /&gt;If none is provided, one will be generated for you.&lt;br /&gt;&lt;br /&gt;This value should be null for GCE as Google will only inform you of your assigned IP once the connection has been accepted. | [optional] 
**PrimaryCloudIp** | Pointer to **string** | The IP address assigned for your side of the Virtual Cross Connect.&lt;br /&gt;&lt;br /&gt;If none is provided, one will be generated for you.&lt;br /&gt;&lt;br /&gt;This value can not be patched once the VXC has bene provisioned. | [optional] 
**PrimaryBgpKey** | Pointer to **string** | The authentication key for BGP peer configuration. | [optional] 
**SecondaryEnabled** | Pointer to **bool** | Indicates whether the secondary circuit is enabled. Setting this to &#x60;false&#x60; will disable the circuit. | [optional] 
**SecondaryCloudAccountId** | Pointer to **string** | The identifier for your Virtual Private Cloud. The number will be different based upon your Cloud provider.&lt;br /&gt;&lt;br /&gt;This attribute is only necessary for GCE. | [optional] 
**SecondaryTelnyxIp** | Pointer to **string** | The IP address assigned to the Telnyx side of the Virtual Cross Connect.&lt;br /&gt;&lt;br /&gt;If none is provided, one will be generated for you.&lt;br /&gt;&lt;br /&gt;This value should be null for GCE as Google will only inform you of your assigned IP once the connection has been accepted. | [optional] 
**SecondaryCloudIp** | Pointer to **string** | The IP address assigned for your side of the Virtual Cross Connect.&lt;br /&gt;&lt;br /&gt;If none is provided, one will be generated for you.&lt;br /&gt;&lt;br /&gt;This value can not be patched once the VXC has bene provisioned. | [optional] 
**SecondaryBgpKey** | Pointer to **string** | The authentication key for BGP peer configuration. | [optional] 
**RegionCode** | **string** | The region interface is deployed to. | 
**PrimaryRoutingAnnouncement** | Pointer to **bool** | Whether the primary BGP route is being announced. | [optional] 
**SecondaryRoutingAnnouncement** | Pointer to **bool** | Whether the secondary BGP route is being announced. | [optional] 
**Region** | Pointer to [**RegionOutRegion**](RegionOutRegion.md) |  | [optional] 

## Methods

### NewVirtualCrossConnectCombined

`func NewVirtualCrossConnectCombined(networkId string, cloudProvider string, cloudProviderRegion string, bgpAsn float32, primaryCloudAccountId string, regionCode string, ) *VirtualCrossConnectCombined`

NewVirtualCrossConnectCombined instantiates a new VirtualCrossConnectCombined object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCrossConnectCombinedWithDefaults

`func NewVirtualCrossConnectCombinedWithDefaults() *VirtualCrossConnectCombined`

NewVirtualCrossConnectCombinedWithDefaults instantiates a new VirtualCrossConnectCombined object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualCrossConnectCombined) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualCrossConnectCombined) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualCrossConnectCombined) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualCrossConnectCombined) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *VirtualCrossConnectCombined) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *VirtualCrossConnectCombined) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *VirtualCrossConnectCombined) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *VirtualCrossConnectCombined) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VirtualCrossConnectCombined) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VirtualCrossConnectCombined) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VirtualCrossConnectCombined) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VirtualCrossConnectCombined) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VirtualCrossConnectCombined) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VirtualCrossConnectCombined) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VirtualCrossConnectCombined) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VirtualCrossConnectCombined) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetNetworkId

`func (o *VirtualCrossConnectCombined) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *VirtualCrossConnectCombined) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *VirtualCrossConnectCombined) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetName

`func (o *VirtualCrossConnectCombined) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualCrossConnectCombined) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualCrossConnectCombined) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualCrossConnectCombined) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualCrossConnectCombined) GetStatus() InterfaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualCrossConnectCombined) GetStatusOk() (*InterfaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualCrossConnectCombined) SetStatus(v InterfaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualCrossConnectCombined) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCloudProvider

`func (o *VirtualCrossConnectCombined) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *VirtualCrossConnectCombined) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *VirtualCrossConnectCombined) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetCloudProviderRegion

`func (o *VirtualCrossConnectCombined) GetCloudProviderRegion() string`

GetCloudProviderRegion returns the CloudProviderRegion field if non-nil, zero value otherwise.

### GetCloudProviderRegionOk

`func (o *VirtualCrossConnectCombined) GetCloudProviderRegionOk() (*string, bool)`

GetCloudProviderRegionOk returns a tuple with the CloudProviderRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderRegion

`func (o *VirtualCrossConnectCombined) SetCloudProviderRegion(v string)`

SetCloudProviderRegion sets CloudProviderRegion field to given value.


### GetBgpAsn

`func (o *VirtualCrossConnectCombined) GetBgpAsn() float32`

GetBgpAsn returns the BgpAsn field if non-nil, zero value otherwise.

### GetBgpAsnOk

`func (o *VirtualCrossConnectCombined) GetBgpAsnOk() (*float32, bool)`

GetBgpAsnOk returns a tuple with the BgpAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAsn

`func (o *VirtualCrossConnectCombined) SetBgpAsn(v float32)`

SetBgpAsn sets BgpAsn field to given value.


### GetBandwidthMbps

`func (o *VirtualCrossConnectCombined) GetBandwidthMbps() float32`

GetBandwidthMbps returns the BandwidthMbps field if non-nil, zero value otherwise.

### GetBandwidthMbpsOk

`func (o *VirtualCrossConnectCombined) GetBandwidthMbpsOk() (*float32, bool)`

GetBandwidthMbpsOk returns a tuple with the BandwidthMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthMbps

`func (o *VirtualCrossConnectCombined) SetBandwidthMbps(v float32)`

SetBandwidthMbps sets BandwidthMbps field to given value.

### HasBandwidthMbps

`func (o *VirtualCrossConnectCombined) HasBandwidthMbps() bool`

HasBandwidthMbps returns a boolean if a field has been set.

### GetPrimaryEnabled

`func (o *VirtualCrossConnectCombined) GetPrimaryEnabled() bool`

GetPrimaryEnabled returns the PrimaryEnabled field if non-nil, zero value otherwise.

### GetPrimaryEnabledOk

`func (o *VirtualCrossConnectCombined) GetPrimaryEnabledOk() (*bool, bool)`

GetPrimaryEnabledOk returns a tuple with the PrimaryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEnabled

`func (o *VirtualCrossConnectCombined) SetPrimaryEnabled(v bool)`

SetPrimaryEnabled sets PrimaryEnabled field to given value.

### HasPrimaryEnabled

`func (o *VirtualCrossConnectCombined) HasPrimaryEnabled() bool`

HasPrimaryEnabled returns a boolean if a field has been set.

### GetPrimaryCloudAccountId

`func (o *VirtualCrossConnectCombined) GetPrimaryCloudAccountId() string`

GetPrimaryCloudAccountId returns the PrimaryCloudAccountId field if non-nil, zero value otherwise.

### GetPrimaryCloudAccountIdOk

`func (o *VirtualCrossConnectCombined) GetPrimaryCloudAccountIdOk() (*string, bool)`

GetPrimaryCloudAccountIdOk returns a tuple with the PrimaryCloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCloudAccountId

`func (o *VirtualCrossConnectCombined) SetPrimaryCloudAccountId(v string)`

SetPrimaryCloudAccountId sets PrimaryCloudAccountId field to given value.


### GetPrimaryTelnyxIp

`func (o *VirtualCrossConnectCombined) GetPrimaryTelnyxIp() string`

GetPrimaryTelnyxIp returns the PrimaryTelnyxIp field if non-nil, zero value otherwise.

### GetPrimaryTelnyxIpOk

`func (o *VirtualCrossConnectCombined) GetPrimaryTelnyxIpOk() (*string, bool)`

GetPrimaryTelnyxIpOk returns a tuple with the PrimaryTelnyxIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTelnyxIp

`func (o *VirtualCrossConnectCombined) SetPrimaryTelnyxIp(v string)`

SetPrimaryTelnyxIp sets PrimaryTelnyxIp field to given value.

### HasPrimaryTelnyxIp

`func (o *VirtualCrossConnectCombined) HasPrimaryTelnyxIp() bool`

HasPrimaryTelnyxIp returns a boolean if a field has been set.

### GetPrimaryCloudIp

`func (o *VirtualCrossConnectCombined) GetPrimaryCloudIp() string`

GetPrimaryCloudIp returns the PrimaryCloudIp field if non-nil, zero value otherwise.

### GetPrimaryCloudIpOk

`func (o *VirtualCrossConnectCombined) GetPrimaryCloudIpOk() (*string, bool)`

GetPrimaryCloudIpOk returns a tuple with the PrimaryCloudIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCloudIp

`func (o *VirtualCrossConnectCombined) SetPrimaryCloudIp(v string)`

SetPrimaryCloudIp sets PrimaryCloudIp field to given value.

### HasPrimaryCloudIp

`func (o *VirtualCrossConnectCombined) HasPrimaryCloudIp() bool`

HasPrimaryCloudIp returns a boolean if a field has been set.

### GetPrimaryBgpKey

`func (o *VirtualCrossConnectCombined) GetPrimaryBgpKey() string`

GetPrimaryBgpKey returns the PrimaryBgpKey field if non-nil, zero value otherwise.

### GetPrimaryBgpKeyOk

`func (o *VirtualCrossConnectCombined) GetPrimaryBgpKeyOk() (*string, bool)`

GetPrimaryBgpKeyOk returns a tuple with the PrimaryBgpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryBgpKey

`func (o *VirtualCrossConnectCombined) SetPrimaryBgpKey(v string)`

SetPrimaryBgpKey sets PrimaryBgpKey field to given value.

### HasPrimaryBgpKey

`func (o *VirtualCrossConnectCombined) HasPrimaryBgpKey() bool`

HasPrimaryBgpKey returns a boolean if a field has been set.

### GetSecondaryEnabled

`func (o *VirtualCrossConnectCombined) GetSecondaryEnabled() bool`

GetSecondaryEnabled returns the SecondaryEnabled field if non-nil, zero value otherwise.

### GetSecondaryEnabledOk

`func (o *VirtualCrossConnectCombined) GetSecondaryEnabledOk() (*bool, bool)`

GetSecondaryEnabledOk returns a tuple with the SecondaryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryEnabled

`func (o *VirtualCrossConnectCombined) SetSecondaryEnabled(v bool)`

SetSecondaryEnabled sets SecondaryEnabled field to given value.

### HasSecondaryEnabled

`func (o *VirtualCrossConnectCombined) HasSecondaryEnabled() bool`

HasSecondaryEnabled returns a boolean if a field has been set.

### GetSecondaryCloudAccountId

`func (o *VirtualCrossConnectCombined) GetSecondaryCloudAccountId() string`

GetSecondaryCloudAccountId returns the SecondaryCloudAccountId field if non-nil, zero value otherwise.

### GetSecondaryCloudAccountIdOk

`func (o *VirtualCrossConnectCombined) GetSecondaryCloudAccountIdOk() (*string, bool)`

GetSecondaryCloudAccountIdOk returns a tuple with the SecondaryCloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCloudAccountId

`func (o *VirtualCrossConnectCombined) SetSecondaryCloudAccountId(v string)`

SetSecondaryCloudAccountId sets SecondaryCloudAccountId field to given value.

### HasSecondaryCloudAccountId

`func (o *VirtualCrossConnectCombined) HasSecondaryCloudAccountId() bool`

HasSecondaryCloudAccountId returns a boolean if a field has been set.

### GetSecondaryTelnyxIp

`func (o *VirtualCrossConnectCombined) GetSecondaryTelnyxIp() string`

GetSecondaryTelnyxIp returns the SecondaryTelnyxIp field if non-nil, zero value otherwise.

### GetSecondaryTelnyxIpOk

`func (o *VirtualCrossConnectCombined) GetSecondaryTelnyxIpOk() (*string, bool)`

GetSecondaryTelnyxIpOk returns a tuple with the SecondaryTelnyxIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryTelnyxIp

`func (o *VirtualCrossConnectCombined) SetSecondaryTelnyxIp(v string)`

SetSecondaryTelnyxIp sets SecondaryTelnyxIp field to given value.

### HasSecondaryTelnyxIp

`func (o *VirtualCrossConnectCombined) HasSecondaryTelnyxIp() bool`

HasSecondaryTelnyxIp returns a boolean if a field has been set.

### GetSecondaryCloudIp

`func (o *VirtualCrossConnectCombined) GetSecondaryCloudIp() string`

GetSecondaryCloudIp returns the SecondaryCloudIp field if non-nil, zero value otherwise.

### GetSecondaryCloudIpOk

`func (o *VirtualCrossConnectCombined) GetSecondaryCloudIpOk() (*string, bool)`

GetSecondaryCloudIpOk returns a tuple with the SecondaryCloudIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCloudIp

`func (o *VirtualCrossConnectCombined) SetSecondaryCloudIp(v string)`

SetSecondaryCloudIp sets SecondaryCloudIp field to given value.

### HasSecondaryCloudIp

`func (o *VirtualCrossConnectCombined) HasSecondaryCloudIp() bool`

HasSecondaryCloudIp returns a boolean if a field has been set.

### GetSecondaryBgpKey

`func (o *VirtualCrossConnectCombined) GetSecondaryBgpKey() string`

GetSecondaryBgpKey returns the SecondaryBgpKey field if non-nil, zero value otherwise.

### GetSecondaryBgpKeyOk

`func (o *VirtualCrossConnectCombined) GetSecondaryBgpKeyOk() (*string, bool)`

GetSecondaryBgpKeyOk returns a tuple with the SecondaryBgpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryBgpKey

`func (o *VirtualCrossConnectCombined) SetSecondaryBgpKey(v string)`

SetSecondaryBgpKey sets SecondaryBgpKey field to given value.

### HasSecondaryBgpKey

`func (o *VirtualCrossConnectCombined) HasSecondaryBgpKey() bool`

HasSecondaryBgpKey returns a boolean if a field has been set.

### GetRegionCode

`func (o *VirtualCrossConnectCombined) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *VirtualCrossConnectCombined) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *VirtualCrossConnectCombined) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.


### GetPrimaryRoutingAnnouncement

`func (o *VirtualCrossConnectCombined) GetPrimaryRoutingAnnouncement() bool`

GetPrimaryRoutingAnnouncement returns the PrimaryRoutingAnnouncement field if non-nil, zero value otherwise.

### GetPrimaryRoutingAnnouncementOk

`func (o *VirtualCrossConnectCombined) GetPrimaryRoutingAnnouncementOk() (*bool, bool)`

GetPrimaryRoutingAnnouncementOk returns a tuple with the PrimaryRoutingAnnouncement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryRoutingAnnouncement

`func (o *VirtualCrossConnectCombined) SetPrimaryRoutingAnnouncement(v bool)`

SetPrimaryRoutingAnnouncement sets PrimaryRoutingAnnouncement field to given value.

### HasPrimaryRoutingAnnouncement

`func (o *VirtualCrossConnectCombined) HasPrimaryRoutingAnnouncement() bool`

HasPrimaryRoutingAnnouncement returns a boolean if a field has been set.

### GetSecondaryRoutingAnnouncement

`func (o *VirtualCrossConnectCombined) GetSecondaryRoutingAnnouncement() bool`

GetSecondaryRoutingAnnouncement returns the SecondaryRoutingAnnouncement field if non-nil, zero value otherwise.

### GetSecondaryRoutingAnnouncementOk

`func (o *VirtualCrossConnectCombined) GetSecondaryRoutingAnnouncementOk() (*bool, bool)`

GetSecondaryRoutingAnnouncementOk returns a tuple with the SecondaryRoutingAnnouncement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryRoutingAnnouncement

`func (o *VirtualCrossConnectCombined) SetSecondaryRoutingAnnouncement(v bool)`

SetSecondaryRoutingAnnouncement sets SecondaryRoutingAnnouncement field to given value.

### HasSecondaryRoutingAnnouncement

`func (o *VirtualCrossConnectCombined) HasSecondaryRoutingAnnouncement() bool`

HasSecondaryRoutingAnnouncement returns a boolean if a field has been set.

### GetRegion

`func (o *VirtualCrossConnectCombined) GetRegion() RegionOutRegion`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *VirtualCrossConnectCombined) GetRegionOk() (*RegionOutRegion, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *VirtualCrossConnectCombined) SetRegion(v RegionOutRegion)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *VirtualCrossConnectCombined) HasRegion() bool`

HasRegion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


