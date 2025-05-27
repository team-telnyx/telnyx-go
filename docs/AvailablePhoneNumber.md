# AvailablePhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** |  | [optional] 
**PhoneNumber** | Pointer to **string** |  | [optional] 
**VanityFormat** | Pointer to **string** |  | [optional] 
**BestEffort** | Pointer to **bool** | Specifies whether the phone number is an exact match based on the search criteria or not. | [optional] 
**Quickship** | Pointer to **bool** | Specifies whether the phone number can receive calls immediately after purchase or not. | [optional] 
**Reservable** | Pointer to **bool** | Specifies whether the phone number can be reserved before purchase or not. | [optional] 
**RegionInformation** | Pointer to [**[]RegionInformation**](RegionInformation.md) |  | [optional] 
**CostInformation** | Pointer to [**CostInformation**](CostInformation.md) |  | [optional] 
**Features** | Pointer to [**[]Feature**](Feature.md) |  | [optional] 

## Methods

### NewAvailablePhoneNumber

`func NewAvailablePhoneNumber() *AvailablePhoneNumber`

NewAvailablePhoneNumber instantiates a new AvailablePhoneNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailablePhoneNumberWithDefaults

`func NewAvailablePhoneNumberWithDefaults() *AvailablePhoneNumber`

NewAvailablePhoneNumberWithDefaults instantiates a new AvailablePhoneNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *AvailablePhoneNumber) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *AvailablePhoneNumber) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *AvailablePhoneNumber) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *AvailablePhoneNumber) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *AvailablePhoneNumber) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *AvailablePhoneNumber) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *AvailablePhoneNumber) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *AvailablePhoneNumber) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetVanityFormat

`func (o *AvailablePhoneNumber) GetVanityFormat() string`

GetVanityFormat returns the VanityFormat field if non-nil, zero value otherwise.

### GetVanityFormatOk

`func (o *AvailablePhoneNumber) GetVanityFormatOk() (*string, bool)`

GetVanityFormatOk returns a tuple with the VanityFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVanityFormat

`func (o *AvailablePhoneNumber) SetVanityFormat(v string)`

SetVanityFormat sets VanityFormat field to given value.

### HasVanityFormat

`func (o *AvailablePhoneNumber) HasVanityFormat() bool`

HasVanityFormat returns a boolean if a field has been set.

### GetBestEffort

`func (o *AvailablePhoneNumber) GetBestEffort() bool`

GetBestEffort returns the BestEffort field if non-nil, zero value otherwise.

### GetBestEffortOk

`func (o *AvailablePhoneNumber) GetBestEffortOk() (*bool, bool)`

GetBestEffortOk returns a tuple with the BestEffort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestEffort

`func (o *AvailablePhoneNumber) SetBestEffort(v bool)`

SetBestEffort sets BestEffort field to given value.

### HasBestEffort

`func (o *AvailablePhoneNumber) HasBestEffort() bool`

HasBestEffort returns a boolean if a field has been set.

### GetQuickship

`func (o *AvailablePhoneNumber) GetQuickship() bool`

GetQuickship returns the Quickship field if non-nil, zero value otherwise.

### GetQuickshipOk

`func (o *AvailablePhoneNumber) GetQuickshipOk() (*bool, bool)`

GetQuickshipOk returns a tuple with the Quickship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuickship

`func (o *AvailablePhoneNumber) SetQuickship(v bool)`

SetQuickship sets Quickship field to given value.

### HasQuickship

`func (o *AvailablePhoneNumber) HasQuickship() bool`

HasQuickship returns a boolean if a field has been set.

### GetReservable

`func (o *AvailablePhoneNumber) GetReservable() bool`

GetReservable returns the Reservable field if non-nil, zero value otherwise.

### GetReservableOk

`func (o *AvailablePhoneNumber) GetReservableOk() (*bool, bool)`

GetReservableOk returns a tuple with the Reservable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReservable

`func (o *AvailablePhoneNumber) SetReservable(v bool)`

SetReservable sets Reservable field to given value.

### HasReservable

`func (o *AvailablePhoneNumber) HasReservable() bool`

HasReservable returns a boolean if a field has been set.

### GetRegionInformation

`func (o *AvailablePhoneNumber) GetRegionInformation() []RegionInformation`

GetRegionInformation returns the RegionInformation field if non-nil, zero value otherwise.

### GetRegionInformationOk

`func (o *AvailablePhoneNumber) GetRegionInformationOk() (*[]RegionInformation, bool)`

GetRegionInformationOk returns a tuple with the RegionInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionInformation

`func (o *AvailablePhoneNumber) SetRegionInformation(v []RegionInformation)`

SetRegionInformation sets RegionInformation field to given value.

### HasRegionInformation

`func (o *AvailablePhoneNumber) HasRegionInformation() bool`

HasRegionInformation returns a boolean if a field has been set.

### GetCostInformation

`func (o *AvailablePhoneNumber) GetCostInformation() CostInformation`

GetCostInformation returns the CostInformation field if non-nil, zero value otherwise.

### GetCostInformationOk

`func (o *AvailablePhoneNumber) GetCostInformationOk() (*CostInformation, bool)`

GetCostInformationOk returns a tuple with the CostInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostInformation

`func (o *AvailablePhoneNumber) SetCostInformation(v CostInformation)`

SetCostInformation sets CostInformation field to given value.

### HasCostInformation

`func (o *AvailablePhoneNumber) HasCostInformation() bool`

HasCostInformation returns a boolean if a field has been set.

### GetFeatures

`func (o *AvailablePhoneNumber) GetFeatures() []Feature`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *AvailablePhoneNumber) GetFeaturesOk() (*[]Feature, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *AvailablePhoneNumber) SetFeatures(v []Feature)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *AvailablePhoneNumber) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


