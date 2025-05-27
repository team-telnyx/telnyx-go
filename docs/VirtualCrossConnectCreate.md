# VirtualCrossConnectCreate

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
**PrimaryEnabled** | Pointer to **bool** | Indicates whether the primary circuit is enabled. Setting this to &#x60;false&#x60; will disable the circuit. | [optional] [readonly] 
**PrimaryCloudAccountId** | **string** | The identifier for your Virtual Private Cloud. The number will be different based upon your Cloud provider. | 
**PrimaryTelnyxIp** | Pointer to **string** | The IP address assigned to the Telnyx side of the Virtual Cross Connect.&lt;br /&gt;&lt;br /&gt;If none is provided, one will be generated for you.&lt;br /&gt;&lt;br /&gt;This value should be null for GCE as Google will only inform you of your assigned IP once the connection has been accepted. | [optional] 
**PrimaryCloudIp** | Pointer to **string** | The IP address assigned for your side of the Virtual Cross Connect.&lt;br /&gt;&lt;br /&gt;If none is provided, one will be generated for you.&lt;br /&gt;&lt;br /&gt;This value should be null for GCE as Google will only inform you of your assigned IP once the connection has been accepted. | [optional] 
**PrimaryBgpKey** | Pointer to **string** | The authentication key for BGP peer configuration. | [optional] 
**SecondaryEnabled** | Pointer to **bool** | Indicates whether the secondary circuit is enabled. Setting this to &#x60;false&#x60; will disable the circuit. | [optional] [readonly] 
**SecondaryCloudAccountId** | Pointer to **string** | The identifier for your Virtual Private Cloud. The number will be different based upon your Cloud provider.&lt;br /&gt;&lt;br /&gt;This attribute is only necessary for GCE. | [optional] 
**SecondaryTelnyxIp** | Pointer to **string** | The IP address assigned to the Telnyx side of the Virtual Cross Connect.&lt;br /&gt;&lt;br /&gt;If none is provided, one will be generated for you.&lt;br /&gt;&lt;br /&gt;This value should be null for GCE as Google will only inform you of your assigned IP once the connection has been accepted. | [optional] 
**SecondaryCloudIp** | Pointer to **string** | The IP address assigned for your side of the Virtual Cross Connect.&lt;br /&gt;&lt;br /&gt;If none is provided, one will be generated for you.&lt;br /&gt;&lt;br /&gt;This value should be null for GCE as Google will only inform you of your assigned IP once the connection has been accepted. | [optional] 
**SecondaryBgpKey** | Pointer to **string** | The authentication key for BGP peer configuration. | [optional] 
**RegionCode** | **string** | The region the interface should be deployed to. | 

## Methods

### NewVirtualCrossConnectCreate

`func NewVirtualCrossConnectCreate(networkId string, cloudProvider string, cloudProviderRegion string, bgpAsn float32, primaryCloudAccountId string, regionCode string, ) *VirtualCrossConnectCreate`

NewVirtualCrossConnectCreate instantiates a new VirtualCrossConnectCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCrossConnectCreateWithDefaults

`func NewVirtualCrossConnectCreateWithDefaults() *VirtualCrossConnectCreate`

NewVirtualCrossConnectCreateWithDefaults instantiates a new VirtualCrossConnectCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualCrossConnectCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualCrossConnectCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualCrossConnectCreate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualCrossConnectCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *VirtualCrossConnectCreate) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *VirtualCrossConnectCreate) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *VirtualCrossConnectCreate) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *VirtualCrossConnectCreate) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VirtualCrossConnectCreate) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VirtualCrossConnectCreate) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VirtualCrossConnectCreate) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VirtualCrossConnectCreate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VirtualCrossConnectCreate) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VirtualCrossConnectCreate) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VirtualCrossConnectCreate) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VirtualCrossConnectCreate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetNetworkId

`func (o *VirtualCrossConnectCreate) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *VirtualCrossConnectCreate) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *VirtualCrossConnectCreate) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.


### GetName

`func (o *VirtualCrossConnectCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualCrossConnectCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualCrossConnectCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualCrossConnectCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualCrossConnectCreate) GetStatus() InterfaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualCrossConnectCreate) GetStatusOk() (*InterfaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualCrossConnectCreate) SetStatus(v InterfaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualCrossConnectCreate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCloudProvider

`func (o *VirtualCrossConnectCreate) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *VirtualCrossConnectCreate) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *VirtualCrossConnectCreate) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.


### GetCloudProviderRegion

`func (o *VirtualCrossConnectCreate) GetCloudProviderRegion() string`

GetCloudProviderRegion returns the CloudProviderRegion field if non-nil, zero value otherwise.

### GetCloudProviderRegionOk

`func (o *VirtualCrossConnectCreate) GetCloudProviderRegionOk() (*string, bool)`

GetCloudProviderRegionOk returns a tuple with the CloudProviderRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderRegion

`func (o *VirtualCrossConnectCreate) SetCloudProviderRegion(v string)`

SetCloudProviderRegion sets CloudProviderRegion field to given value.


### GetBgpAsn

`func (o *VirtualCrossConnectCreate) GetBgpAsn() float32`

GetBgpAsn returns the BgpAsn field if non-nil, zero value otherwise.

### GetBgpAsnOk

`func (o *VirtualCrossConnectCreate) GetBgpAsnOk() (*float32, bool)`

GetBgpAsnOk returns a tuple with the BgpAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAsn

`func (o *VirtualCrossConnectCreate) SetBgpAsn(v float32)`

SetBgpAsn sets BgpAsn field to given value.


### GetBandwidthMbps

`func (o *VirtualCrossConnectCreate) GetBandwidthMbps() float32`

GetBandwidthMbps returns the BandwidthMbps field if non-nil, zero value otherwise.

### GetBandwidthMbpsOk

`func (o *VirtualCrossConnectCreate) GetBandwidthMbpsOk() (*float32, bool)`

GetBandwidthMbpsOk returns a tuple with the BandwidthMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthMbps

`func (o *VirtualCrossConnectCreate) SetBandwidthMbps(v float32)`

SetBandwidthMbps sets BandwidthMbps field to given value.

### HasBandwidthMbps

`func (o *VirtualCrossConnectCreate) HasBandwidthMbps() bool`

HasBandwidthMbps returns a boolean if a field has been set.

### GetPrimaryEnabled

`func (o *VirtualCrossConnectCreate) GetPrimaryEnabled() bool`

GetPrimaryEnabled returns the PrimaryEnabled field if non-nil, zero value otherwise.

### GetPrimaryEnabledOk

`func (o *VirtualCrossConnectCreate) GetPrimaryEnabledOk() (*bool, bool)`

GetPrimaryEnabledOk returns a tuple with the PrimaryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEnabled

`func (o *VirtualCrossConnectCreate) SetPrimaryEnabled(v bool)`

SetPrimaryEnabled sets PrimaryEnabled field to given value.

### HasPrimaryEnabled

`func (o *VirtualCrossConnectCreate) HasPrimaryEnabled() bool`

HasPrimaryEnabled returns a boolean if a field has been set.

### GetPrimaryCloudAccountId

`func (o *VirtualCrossConnectCreate) GetPrimaryCloudAccountId() string`

GetPrimaryCloudAccountId returns the PrimaryCloudAccountId field if non-nil, zero value otherwise.

### GetPrimaryCloudAccountIdOk

`func (o *VirtualCrossConnectCreate) GetPrimaryCloudAccountIdOk() (*string, bool)`

GetPrimaryCloudAccountIdOk returns a tuple with the PrimaryCloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCloudAccountId

`func (o *VirtualCrossConnectCreate) SetPrimaryCloudAccountId(v string)`

SetPrimaryCloudAccountId sets PrimaryCloudAccountId field to given value.


### GetPrimaryTelnyxIp

`func (o *VirtualCrossConnectCreate) GetPrimaryTelnyxIp() string`

GetPrimaryTelnyxIp returns the PrimaryTelnyxIp field if non-nil, zero value otherwise.

### GetPrimaryTelnyxIpOk

`func (o *VirtualCrossConnectCreate) GetPrimaryTelnyxIpOk() (*string, bool)`

GetPrimaryTelnyxIpOk returns a tuple with the PrimaryTelnyxIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTelnyxIp

`func (o *VirtualCrossConnectCreate) SetPrimaryTelnyxIp(v string)`

SetPrimaryTelnyxIp sets PrimaryTelnyxIp field to given value.

### HasPrimaryTelnyxIp

`func (o *VirtualCrossConnectCreate) HasPrimaryTelnyxIp() bool`

HasPrimaryTelnyxIp returns a boolean if a field has been set.

### GetPrimaryCloudIp

`func (o *VirtualCrossConnectCreate) GetPrimaryCloudIp() string`

GetPrimaryCloudIp returns the PrimaryCloudIp field if non-nil, zero value otherwise.

### GetPrimaryCloudIpOk

`func (o *VirtualCrossConnectCreate) GetPrimaryCloudIpOk() (*string, bool)`

GetPrimaryCloudIpOk returns a tuple with the PrimaryCloudIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCloudIp

`func (o *VirtualCrossConnectCreate) SetPrimaryCloudIp(v string)`

SetPrimaryCloudIp sets PrimaryCloudIp field to given value.

### HasPrimaryCloudIp

`func (o *VirtualCrossConnectCreate) HasPrimaryCloudIp() bool`

HasPrimaryCloudIp returns a boolean if a field has been set.

### GetPrimaryBgpKey

`func (o *VirtualCrossConnectCreate) GetPrimaryBgpKey() string`

GetPrimaryBgpKey returns the PrimaryBgpKey field if non-nil, zero value otherwise.

### GetPrimaryBgpKeyOk

`func (o *VirtualCrossConnectCreate) GetPrimaryBgpKeyOk() (*string, bool)`

GetPrimaryBgpKeyOk returns a tuple with the PrimaryBgpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryBgpKey

`func (o *VirtualCrossConnectCreate) SetPrimaryBgpKey(v string)`

SetPrimaryBgpKey sets PrimaryBgpKey field to given value.

### HasPrimaryBgpKey

`func (o *VirtualCrossConnectCreate) HasPrimaryBgpKey() bool`

HasPrimaryBgpKey returns a boolean if a field has been set.

### GetSecondaryEnabled

`func (o *VirtualCrossConnectCreate) GetSecondaryEnabled() bool`

GetSecondaryEnabled returns the SecondaryEnabled field if non-nil, zero value otherwise.

### GetSecondaryEnabledOk

`func (o *VirtualCrossConnectCreate) GetSecondaryEnabledOk() (*bool, bool)`

GetSecondaryEnabledOk returns a tuple with the SecondaryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryEnabled

`func (o *VirtualCrossConnectCreate) SetSecondaryEnabled(v bool)`

SetSecondaryEnabled sets SecondaryEnabled field to given value.

### HasSecondaryEnabled

`func (o *VirtualCrossConnectCreate) HasSecondaryEnabled() bool`

HasSecondaryEnabled returns a boolean if a field has been set.

### GetSecondaryCloudAccountId

`func (o *VirtualCrossConnectCreate) GetSecondaryCloudAccountId() string`

GetSecondaryCloudAccountId returns the SecondaryCloudAccountId field if non-nil, zero value otherwise.

### GetSecondaryCloudAccountIdOk

`func (o *VirtualCrossConnectCreate) GetSecondaryCloudAccountIdOk() (*string, bool)`

GetSecondaryCloudAccountIdOk returns a tuple with the SecondaryCloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCloudAccountId

`func (o *VirtualCrossConnectCreate) SetSecondaryCloudAccountId(v string)`

SetSecondaryCloudAccountId sets SecondaryCloudAccountId field to given value.

### HasSecondaryCloudAccountId

`func (o *VirtualCrossConnectCreate) HasSecondaryCloudAccountId() bool`

HasSecondaryCloudAccountId returns a boolean if a field has been set.

### GetSecondaryTelnyxIp

`func (o *VirtualCrossConnectCreate) GetSecondaryTelnyxIp() string`

GetSecondaryTelnyxIp returns the SecondaryTelnyxIp field if non-nil, zero value otherwise.

### GetSecondaryTelnyxIpOk

`func (o *VirtualCrossConnectCreate) GetSecondaryTelnyxIpOk() (*string, bool)`

GetSecondaryTelnyxIpOk returns a tuple with the SecondaryTelnyxIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryTelnyxIp

`func (o *VirtualCrossConnectCreate) SetSecondaryTelnyxIp(v string)`

SetSecondaryTelnyxIp sets SecondaryTelnyxIp field to given value.

### HasSecondaryTelnyxIp

`func (o *VirtualCrossConnectCreate) HasSecondaryTelnyxIp() bool`

HasSecondaryTelnyxIp returns a boolean if a field has been set.

### GetSecondaryCloudIp

`func (o *VirtualCrossConnectCreate) GetSecondaryCloudIp() string`

GetSecondaryCloudIp returns the SecondaryCloudIp field if non-nil, zero value otherwise.

### GetSecondaryCloudIpOk

`func (o *VirtualCrossConnectCreate) GetSecondaryCloudIpOk() (*string, bool)`

GetSecondaryCloudIpOk returns a tuple with the SecondaryCloudIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCloudIp

`func (o *VirtualCrossConnectCreate) SetSecondaryCloudIp(v string)`

SetSecondaryCloudIp sets SecondaryCloudIp field to given value.

### HasSecondaryCloudIp

`func (o *VirtualCrossConnectCreate) HasSecondaryCloudIp() bool`

HasSecondaryCloudIp returns a boolean if a field has been set.

### GetSecondaryBgpKey

`func (o *VirtualCrossConnectCreate) GetSecondaryBgpKey() string`

GetSecondaryBgpKey returns the SecondaryBgpKey field if non-nil, zero value otherwise.

### GetSecondaryBgpKeyOk

`func (o *VirtualCrossConnectCreate) GetSecondaryBgpKeyOk() (*string, bool)`

GetSecondaryBgpKeyOk returns a tuple with the SecondaryBgpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryBgpKey

`func (o *VirtualCrossConnectCreate) SetSecondaryBgpKey(v string)`

SetSecondaryBgpKey sets SecondaryBgpKey field to given value.

### HasSecondaryBgpKey

`func (o *VirtualCrossConnectCreate) HasSecondaryBgpKey() bool`

HasSecondaryBgpKey returns a boolean if a field has been set.

### GetRegionCode

`func (o *VirtualCrossConnectCreate) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *VirtualCrossConnectCreate) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *VirtualCrossConnectCreate) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


