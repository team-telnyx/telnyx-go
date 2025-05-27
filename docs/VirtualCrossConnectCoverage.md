# VirtualCrossConnectCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**CloudProvider** | Pointer to **string** | The Virtual Private Cloud with which you would like to establish a cross connect. | [optional] 
**CloudProviderRegion** | Pointer to **string** | The region where your Virtual Private Cloud hosts are located. Should be identical to how the cloud provider names region, i.e. us-east-1 for AWS but Frankfurt for Azure | [optional] 
**AvailableBandwidth** | Pointer to **[]float32** | The available throughput in Megabits per Second (Mbps) for your Virtual Cross Connect. | [optional] 

## Methods

### NewVirtualCrossConnectCoverage

`func NewVirtualCrossConnectCoverage() *VirtualCrossConnectCoverage`

NewVirtualCrossConnectCoverage instantiates a new VirtualCrossConnectCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVirtualCrossConnectCoverageWithDefaults

`func NewVirtualCrossConnectCoverageWithDefaults() *VirtualCrossConnectCoverage`

NewVirtualCrossConnectCoverageWithDefaults instantiates a new VirtualCrossConnectCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *VirtualCrossConnectCoverage) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *VirtualCrossConnectCoverage) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *VirtualCrossConnectCoverage) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *VirtualCrossConnectCoverage) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetLocation

`func (o *VirtualCrossConnectCoverage) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *VirtualCrossConnectCoverage) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *VirtualCrossConnectCoverage) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *VirtualCrossConnectCoverage) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetCloudProvider

`func (o *VirtualCrossConnectCoverage) GetCloudProvider() string`

GetCloudProvider returns the CloudProvider field if non-nil, zero value otherwise.

### GetCloudProviderOk

`func (o *VirtualCrossConnectCoverage) GetCloudProviderOk() (*string, bool)`

GetCloudProviderOk returns a tuple with the CloudProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProvider

`func (o *VirtualCrossConnectCoverage) SetCloudProvider(v string)`

SetCloudProvider sets CloudProvider field to given value.

### HasCloudProvider

`func (o *VirtualCrossConnectCoverage) HasCloudProvider() bool`

HasCloudProvider returns a boolean if a field has been set.

### GetCloudProviderRegion

`func (o *VirtualCrossConnectCoverage) GetCloudProviderRegion() string`

GetCloudProviderRegion returns the CloudProviderRegion field if non-nil, zero value otherwise.

### GetCloudProviderRegionOk

`func (o *VirtualCrossConnectCoverage) GetCloudProviderRegionOk() (*string, bool)`

GetCloudProviderRegionOk returns a tuple with the CloudProviderRegion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloudProviderRegion

`func (o *VirtualCrossConnectCoverage) SetCloudProviderRegion(v string)`

SetCloudProviderRegion sets CloudProviderRegion field to given value.

### HasCloudProviderRegion

`func (o *VirtualCrossConnectCoverage) HasCloudProviderRegion() bool`

HasCloudProviderRegion returns a boolean if a field has been set.

### GetAvailableBandwidth

`func (o *VirtualCrossConnectCoverage) GetAvailableBandwidth() []float32`

GetAvailableBandwidth returns the AvailableBandwidth field if non-nil, zero value otherwise.

### GetAvailableBandwidthOk

`func (o *VirtualCrossConnectCoverage) GetAvailableBandwidthOk() (*[]float32, bool)`

GetAvailableBandwidthOk returns a tuple with the AvailableBandwidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBandwidth

`func (o *VirtualCrossConnectCoverage) SetAvailableBandwidth(v []float32)`

SetAvailableBandwidth sets AvailableBandwidth field to given value.

### HasAvailableBandwidth

`func (o *VirtualCrossConnectCoverage) HasAvailableBandwidth() bool`

HasAvailableBandwidth returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


