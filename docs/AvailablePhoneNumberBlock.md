# AvailablePhoneNumberBlock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** |  | [optional] 
**StartingNumber** | Pointer to **string** |  | [optional] 
**Range** | Pointer to **int32** |  | [optional] 
**RegionInformation** | Pointer to [**[]RegionInformation**](RegionInformation.md) |  | [optional] 
**CostInformation** | Pointer to [**CostInformation**](CostInformation.md) |  | [optional] 
**Features** | Pointer to [**[]Feature**](Feature.md) |  | [optional] 

## Methods

### NewAvailablePhoneNumberBlock

`func NewAvailablePhoneNumberBlock() *AvailablePhoneNumberBlock`

NewAvailablePhoneNumberBlock instantiates a new AvailablePhoneNumberBlock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailablePhoneNumberBlockWithDefaults

`func NewAvailablePhoneNumberBlockWithDefaults() *AvailablePhoneNumberBlock`

NewAvailablePhoneNumberBlockWithDefaults instantiates a new AvailablePhoneNumberBlock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *AvailablePhoneNumberBlock) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *AvailablePhoneNumberBlock) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *AvailablePhoneNumberBlock) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *AvailablePhoneNumberBlock) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetStartingNumber

`func (o *AvailablePhoneNumberBlock) GetStartingNumber() string`

GetStartingNumber returns the StartingNumber field if non-nil, zero value otherwise.

### GetStartingNumberOk

`func (o *AvailablePhoneNumberBlock) GetStartingNumberOk() (*string, bool)`

GetStartingNumberOk returns a tuple with the StartingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingNumber

`func (o *AvailablePhoneNumberBlock) SetStartingNumber(v string)`

SetStartingNumber sets StartingNumber field to given value.

### HasStartingNumber

`func (o *AvailablePhoneNumberBlock) HasStartingNumber() bool`

HasStartingNumber returns a boolean if a field has been set.

### GetRange

`func (o *AvailablePhoneNumberBlock) GetRange() int32`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *AvailablePhoneNumberBlock) GetRangeOk() (*int32, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *AvailablePhoneNumberBlock) SetRange(v int32)`

SetRange sets Range field to given value.

### HasRange

`func (o *AvailablePhoneNumberBlock) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRegionInformation

`func (o *AvailablePhoneNumberBlock) GetRegionInformation() []RegionInformation`

GetRegionInformation returns the RegionInformation field if non-nil, zero value otherwise.

### GetRegionInformationOk

`func (o *AvailablePhoneNumberBlock) GetRegionInformationOk() (*[]RegionInformation, bool)`

GetRegionInformationOk returns a tuple with the RegionInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionInformation

`func (o *AvailablePhoneNumberBlock) SetRegionInformation(v []RegionInformation)`

SetRegionInformation sets RegionInformation field to given value.

### HasRegionInformation

`func (o *AvailablePhoneNumberBlock) HasRegionInformation() bool`

HasRegionInformation returns a boolean if a field has been set.

### GetCostInformation

`func (o *AvailablePhoneNumberBlock) GetCostInformation() CostInformation`

GetCostInformation returns the CostInformation field if non-nil, zero value otherwise.

### GetCostInformationOk

`func (o *AvailablePhoneNumberBlock) GetCostInformationOk() (*CostInformation, bool)`

GetCostInformationOk returns a tuple with the CostInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostInformation

`func (o *AvailablePhoneNumberBlock) SetCostInformation(v CostInformation)`

SetCostInformation sets CostInformation field to given value.

### HasCostInformation

`func (o *AvailablePhoneNumberBlock) HasCostInformation() bool`

HasCostInformation returns a boolean if a field has been set.

### GetFeatures

`func (o *AvailablePhoneNumberBlock) GetFeatures() []Feature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *AvailablePhoneNumberBlock) GetFeaturesOk() (*[]Feature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *AvailablePhoneNumberBlock) SetFeatures(v []Feature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *AvailablePhoneNumberBlock) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


