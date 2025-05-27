# NetworkCoverage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**AvailableServices** | Pointer to [**[]NetworkCoverageAvailableServicesInner**](NetworkCoverageAvailableServicesInner.md) | List of interface types supported in this region. | [optional] 

## Methods

### NewNetworkCoverage

`func NewNetworkCoverage() *NetworkCoverage`

NewNetworkCoverage instantiates a new NetworkCoverage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkCoverageWithDefaults

`func NewNetworkCoverageWithDefaults() *NetworkCoverage`

NewNetworkCoverageWithDefaults instantiates a new NetworkCoverage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *NetworkCoverage) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *NetworkCoverage) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *NetworkCoverage) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *NetworkCoverage) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetLocation

`func (o *NetworkCoverage) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *NetworkCoverage) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *NetworkCoverage) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *NetworkCoverage) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetAvailableServices

`func (o *NetworkCoverage) GetAvailableServices() []NetworkCoverageAvailableServicesInner`

GetAvailableServices returns the AvailableServices field if non-nil, zero value otherwise.

### GetAvailableServicesOk

`func (o *NetworkCoverage) GetAvailableServicesOk() (*[]NetworkCoverageAvailableServicesInner, bool)`

GetAvailableServicesOk returns a tuple with the AvailableServices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableServices

`func (o *NetworkCoverage) SetAvailableServices(v []NetworkCoverageAvailableServicesInner)`

SetAvailableServices sets AvailableServices field to given value.

### HasAvailableServices

`func (o *NetworkCoverage) HasAvailableServices() bool`

HasAvailableServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


