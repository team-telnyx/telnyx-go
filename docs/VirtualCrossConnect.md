# VirtualCrossConnect

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Identifies the resource. | [optional] [readonly] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **string** | ISO 8601 formatted date-time indicating when the resource was updated. | [optional] [readonly] 
**NetworkId** | Pointer to **string** | The id of the network associated with the interface. | [optional] 
**Name** | Pointer to **string** | A user specified name for the interface. | [optional] 
**Status** | Pointer to [**InterfaceStatus**](InterfaceStatus.md) |  | [optional] 
**CloudProvider** | Pointer to **string** | The Virtual Private Cloud with which you would like to establish a cross connect. | [optional] 
**CloudProviderRegion** | Pointer to **string** | The region where your Virtual Private Cloud hosts are located.&lt;br /&gt;&lt;br /&gt;The available regions can be found using the /virtual_cross_connect_regions endpoint. | [optional] 
**BgpAsn** | Pointer to **float32** | The Border Gateway Protocol (BGP) Autonomous System Number (ASN). If null, value will be assigned by Telnyx. | [optional] 
**BandwidthMbps** | Pointer to **float32** | The desired throughput in Megabits per Second (Mbps) for your Virtual Cross Connect.&lt;br /&gt;&lt;br /&gt;The available bandwidths can be found using the /virtual_cross_connect_regions endpoint. | [optional] 
**PrimaryEnabled** | Pointer to **bool** | Indicates whether the primary circuit is enabled. Setting this to &#x60;false&#x60; will disable the circuit. | [optional] [readonly] 
**PrimaryCloudAccountId** | Pointer to **string** | The identifier for your Virtual Private Cloud. The number will be different based upon your Cloud provider. | [optional] 
**PrimaryTelnyxIp** | Pointer to **string** | The IP address assigned to the Telnyx side of the Virtual Cross Connect.&lt;br /&gt;&lt;br /&gt;If none is provided, one will be generated for you.&lt;br /&gt;&lt;br /&gt;This value should be null for GCE as Google will only inform you of your assigned IP once the connection has been accepted. | [optional] 
**PrimaryCloudIp** | Pointer to **string** | The IP address assigned for your side of the Virtual Cross Connect.&lt;br /&gt;&lt;br /&gt;If none is provided, one will be generated for you.&lt;br /&gt;&lt;br /&gt;This value should be null for GCE as Google will only inform you of your assigned IP once the connection has been accepted. | [optional] 
**PrimaryBgpKey** | Pointer to **string** | The authentication key for BGP peer configuration. | [optional] 
**SecondaryEnabled** | Pointer to **bool** | Indicates whether the secondary circuit is enabled. Setting this to &#x60;false&#x60; will disable the circuit. | [optional] [readonly] 
**SecondaryCloudAccountId** | Pointer to **string** | The identifier for your Virtual Private Cloud. The number will be different based upon your Cloud provider.&lt;br /&gt;&lt;br /&gt;This attribute is only necessary for GCE. | [optional] 
**SecondaryTelnyxIp** | Pointer to **string** | The IP address assigned to the Telnyx side of the Virtual Cross Connect.&lt;br /&gt;&lt;br /&gt;If none is provided, one will be generated for you.&lt;br /&gt;&lt;br /&gt;This value should be null for GCE as Google will only inform you of your assigned IP once the connection has been accepted. | [optional] 
**SecondaryCloudIp** | Pointer to **string** | The IP address assigned for your side of the Virtual Cross Connect.&lt;br /&gt;&lt;br /&gt;If none is provided, one will be generated for you.&lt;br /&gt;&lt;br /&gt;This value should be null for GCE as Google will only inform you of your assigned IP once the connection has been accepted. | [optional] 
**SecondaryBgpKey** | Pointer to **string** | The authentication key for BGP peer configuration. | [optional] 

## Methods

### NewVirtualCrossConnect

`func NewVirtualCrossConnect() *VirtualCrossConnect`

NewVirtualCrossConnect instantiates a new VirtualCrossConnect object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCrossConnectWithDefaults

`func NewVirtualCrossConnectWithDefaults() *VirtualCrossConnect`

NewVirtualCrossConnectWithDefaults instantiates a new VirtualCrossConnect object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *VirtualCrossConnect) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VirtualCrossConnect) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VirtualCrossConnect) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VirtualCrossConnect) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *VirtualCrossConnect) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *VirtualCrossConnect) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *VirtualCrossConnect) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *VirtualCrossConnect) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VirtualCrossConnect) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VirtualCrossConnect) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VirtualCrossConnect) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VirtualCrossConnect) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VirtualCrossConnect) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VirtualCrossConnect) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VirtualCrossConnect) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VirtualCrossConnect) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetNetworkId

`func (o *VirtualCrossConnect) GetNetworkId() string`

GetNetworkId returns the NetworkId field if non-nil, zero value otherwise.

### GetNetworkIdOk

`func (o *VirtualCrossConnect) GetNetworkIdOk() (*string, bool)`

GetNetworkIdOk returns a tuple with the NetworkId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkId

`func (o *VirtualCrossConnect) SetNetworkId(v string)`

SetNetworkId sets NetworkId field to given value.

### HasNetworkId

`func (o *VirtualCrossConnect) HasNetworkId() bool`

HasNetworkId returns a boolean if a field has been set.

### GetName

`func (o *VirtualCrossConnect) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *VirtualCrossConnect) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *VirtualCrossConnect) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *VirtualCrossConnect) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *VirtualCrossConnect) GetStatus() InterfaceStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VirtualCrossConnect) GetStatusOk() (*InterfaceStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VirtualCrossConnect) SetStatus(v InterfaceStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VirtualCrossConnect) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCloudProvider

`func (o *VirtualCrossConnect) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *VirtualCrossConnect) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *VirtualCrossConnect) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *VirtualCrossConnect) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCloudProviderRegion

`func (o *VirtualCrossConnect) GetCloudProviderRegion() string`

GetCloudProviderRegion returns the CloudProviderRegion field if non-nil, zero value otherwise.

### GetCloudProviderRegionOk

`func (o *VirtualCrossConnect) GetCloudProviderRegionOk() (*string, bool)`

GetCloudProviderRegionOk returns a tuple with the CloudProviderRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderRegion

`func (o *VirtualCrossConnect) SetCloudProviderRegion(v string)`

SetCloudProviderRegion sets CloudProviderRegion field to given value.

### HasCloudProviderRegion

`func (o *VirtualCrossConnect) HasCloudProviderRegion() bool`

HasCloudProviderRegion returns a boolean if a field has been set.

### GetBgpAsn

`func (o *VirtualCrossConnect) GetBgpAsn() float32`

GetBgpAsn returns the BgpAsn field if non-nil, zero value otherwise.

### GetBgpAsnOk

`func (o *VirtualCrossConnect) GetBgpAsnOk() (*float32, bool)`

GetBgpAsnOk returns a tuple with the BgpAsn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpAsn

`func (o *VirtualCrossConnect) SetBgpAsn(v float32)`

SetBgpAsn sets BgpAsn field to given value.

### HasBgpAsn

`func (o *VirtualCrossConnect) HasBgpAsn() bool`

HasBgpAsn returns a boolean if a field has been set.

### GetBandwidthMbps

`func (o *VirtualCrossConnect) GetBandwidthMbps() float32`

GetBandwidthMbps returns the BandwidthMbps field if non-nil, zero value otherwise.

### GetBandwidthMbpsOk

`func (o *VirtualCrossConnect) GetBandwidthMbpsOk() (*float32, bool)`

GetBandwidthMbpsOk returns a tuple with the BandwidthMbps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBandwidthMbps

`func (o *VirtualCrossConnect) SetBandwidthMbps(v float32)`

SetBandwidthMbps sets BandwidthMbps field to given value.

### HasBandwidthMbps

`func (o *VirtualCrossConnect) HasBandwidthMbps() bool`

HasBandwidthMbps returns a boolean if a field has been set.

### GetPrimaryEnabled

`func (o *VirtualCrossConnect) GetPrimaryEnabled() bool`

GetPrimaryEnabled returns the PrimaryEnabled field if non-nil, zero value otherwise.

### GetPrimaryEnabledOk

`func (o *VirtualCrossConnect) GetPrimaryEnabledOk() (*bool, bool)`

GetPrimaryEnabledOk returns a tuple with the PrimaryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryEnabled

`func (o *VirtualCrossConnect) SetPrimaryEnabled(v bool)`

SetPrimaryEnabled sets PrimaryEnabled field to given value.

### HasPrimaryEnabled

`func (o *VirtualCrossConnect) HasPrimaryEnabled() bool`

HasPrimaryEnabled returns a boolean if a field has been set.

### GetPrimaryCloudAccountId

`func (o *VirtualCrossConnect) GetPrimaryCloudAccountId() string`

GetPrimaryCloudAccountId returns the PrimaryCloudAccountId field if non-nil, zero value otherwise.

### GetPrimaryCloudAccountIdOk

`func (o *VirtualCrossConnect) GetPrimaryCloudAccountIdOk() (*string, bool)`

GetPrimaryCloudAccountIdOk returns a tuple with the PrimaryCloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCloudAccountId

`func (o *VirtualCrossConnect) SetPrimaryCloudAccountId(v string)`

SetPrimaryCloudAccountId sets PrimaryCloudAccountId field to given value.

### HasPrimaryCloudAccountId

`func (o *VirtualCrossConnect) HasPrimaryCloudAccountId() bool`

HasPrimaryCloudAccountId returns a boolean if a field has been set.

### GetPrimaryTelnyxIp

`func (o *VirtualCrossConnect) GetPrimaryTelnyxIp() string`

GetPrimaryTelnyxIp returns the PrimaryTelnyxIp field if non-nil, zero value otherwise.

### GetPrimaryTelnyxIpOk

`func (o *VirtualCrossConnect) GetPrimaryTelnyxIpOk() (*string, bool)`

GetPrimaryTelnyxIpOk returns a tuple with the PrimaryTelnyxIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryTelnyxIp

`func (o *VirtualCrossConnect) SetPrimaryTelnyxIp(v string)`

SetPrimaryTelnyxIp sets PrimaryTelnyxIp field to given value.

### HasPrimaryTelnyxIp

`func (o *VirtualCrossConnect) HasPrimaryTelnyxIp() bool`

HasPrimaryTelnyxIp returns a boolean if a field has been set.

### GetPrimaryCloudIp

`func (o *VirtualCrossConnect) GetPrimaryCloudIp() string`

GetPrimaryCloudIp returns the PrimaryCloudIp field if non-nil, zero value otherwise.

### GetPrimaryCloudIpOk

`func (o *VirtualCrossConnect) GetPrimaryCloudIpOk() (*string, bool)`

GetPrimaryCloudIpOk returns a tuple with the PrimaryCloudIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryCloudIp

`func (o *VirtualCrossConnect) SetPrimaryCloudIp(v string)`

SetPrimaryCloudIp sets PrimaryCloudIp field to given value.

### HasPrimaryCloudIp

`func (o *VirtualCrossConnect) HasPrimaryCloudIp() bool`

HasPrimaryCloudIp returns a boolean if a field has been set.

### GetPrimaryBgpKey

`func (o *VirtualCrossConnect) GetPrimaryBgpKey() string`

GetPrimaryBgpKey returns the PrimaryBgpKey field if non-nil, zero value otherwise.

### GetPrimaryBgpKeyOk

`func (o *VirtualCrossConnect) GetPrimaryBgpKeyOk() (*string, bool)`

GetPrimaryBgpKeyOk returns a tuple with the PrimaryBgpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimaryBgpKey

`func (o *VirtualCrossConnect) SetPrimaryBgpKey(v string)`

SetPrimaryBgpKey sets PrimaryBgpKey field to given value.

### HasPrimaryBgpKey

`func (o *VirtualCrossConnect) HasPrimaryBgpKey() bool`

HasPrimaryBgpKey returns a boolean if a field has been set.

### GetSecondaryEnabled

`func (o *VirtualCrossConnect) GetSecondaryEnabled() bool`

GetSecondaryEnabled returns the SecondaryEnabled field if non-nil, zero value otherwise.

### GetSecondaryEnabledOk

`func (o *VirtualCrossConnect) GetSecondaryEnabledOk() (*bool, bool)`

GetSecondaryEnabledOk returns a tuple with the SecondaryEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryEnabled

`func (o *VirtualCrossConnect) SetSecondaryEnabled(v bool)`

SetSecondaryEnabled sets SecondaryEnabled field to given value.

### HasSecondaryEnabled

`func (o *VirtualCrossConnect) HasSecondaryEnabled() bool`

HasSecondaryEnabled returns a boolean if a field has been set.

### GetSecondaryCloudAccountId

`func (o *VirtualCrossConnect) GetSecondaryCloudAccountId() string`

GetSecondaryCloudAccountId returns the SecondaryCloudAccountId field if non-nil, zero value otherwise.

### GetSecondaryCloudAccountIdOk

`func (o *VirtualCrossConnect) GetSecondaryCloudAccountIdOk() (*string, bool)`

GetSecondaryCloudAccountIdOk returns a tuple with the SecondaryCloudAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCloudAccountId

`func (o *VirtualCrossConnect) SetSecondaryCloudAccountId(v string)`

SetSecondaryCloudAccountId sets SecondaryCloudAccountId field to given value.

### HasSecondaryCloudAccountId

`func (o *VirtualCrossConnect) HasSecondaryCloudAccountId() bool`

HasSecondaryCloudAccountId returns a boolean if a field has been set.

### GetSecondaryTelnyxIp

`func (o *VirtualCrossConnect) GetSecondaryTelnyxIp() string`

GetSecondaryTelnyxIp returns the SecondaryTelnyxIp field if non-nil, zero value otherwise.

### GetSecondaryTelnyxIpOk

`func (o *VirtualCrossConnect) GetSecondaryTelnyxIpOk() (*string, bool)`

GetSecondaryTelnyxIpOk returns a tuple with the SecondaryTelnyxIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryTelnyxIp

`func (o *VirtualCrossConnect) SetSecondaryTelnyxIp(v string)`

SetSecondaryTelnyxIp sets SecondaryTelnyxIp field to given value.

### HasSecondaryTelnyxIp

`func (o *VirtualCrossConnect) HasSecondaryTelnyxIp() bool`

HasSecondaryTelnyxIp returns a boolean if a field has been set.

### GetSecondaryCloudIp

`func (o *VirtualCrossConnect) GetSecondaryCloudIp() string`

GetSecondaryCloudIp returns the SecondaryCloudIp field if non-nil, zero value otherwise.

### GetSecondaryCloudIpOk

`func (o *VirtualCrossConnect) GetSecondaryCloudIpOk() (*string, bool)`

GetSecondaryCloudIpOk returns a tuple with the SecondaryCloudIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryCloudIp

`func (o *VirtualCrossConnect) SetSecondaryCloudIp(v string)`

SetSecondaryCloudIp sets SecondaryCloudIp field to given value.

### HasSecondaryCloudIp

`func (o *VirtualCrossConnect) HasSecondaryCloudIp() bool`

HasSecondaryCloudIp returns a boolean if a field has been set.

### GetSecondaryBgpKey

`func (o *VirtualCrossConnect) GetSecondaryBgpKey() string`

GetSecondaryBgpKey returns the SecondaryBgpKey field if non-nil, zero value otherwise.

### GetSecondaryBgpKeyOk

`func (o *VirtualCrossConnect) GetSecondaryBgpKeyOk() (*string, bool)`

GetSecondaryBgpKeyOk returns a tuple with the SecondaryBgpKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryBgpKey

`func (o *VirtualCrossConnect) SetSecondaryBgpKey(v string)`

SetSecondaryBgpKey sets SecondaryBgpKey field to given value.

### HasSecondaryBgpKey

`func (o *VirtualCrossConnect) HasSecondaryBgpKey() bool`

HasSecondaryBgpKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


