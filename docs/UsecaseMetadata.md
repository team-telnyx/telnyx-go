# UsecaseMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnnualFee** | Pointer to **float32** | Campaign annual subscription fee | [optional] 
**MaxSubUsecases** | Pointer to **int32** | Maximum number of sub-usecases declaration required. | [optional] 
**MinSubUsecases** | Pointer to **int32** | Minimum number of sub-usecases declaration required. | [optional] 
**MnoMetadata** | Pointer to **map[string]interface{}** | Map of usecase metadata for each MNO. Key is the network ID of the MNO (e.g. 10017), Value is the mno metadata for the usecase. | [optional] 
**MonthlyFee** | Pointer to **float32** | Campaign monthly subscription fee | [optional] 
**QuarterlyFee** | Pointer to **float32** | Campaign quarterly subscription fee | [optional] 
**Usecase** | Pointer to **string** | Campaign usecase | [optional] 

## Methods

### NewUsecaseMetadata

`func NewUsecaseMetadata() *UsecaseMetadata`

NewUsecaseMetadata instantiates a new UsecaseMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUsecaseMetadataWithDefaults

`func NewUsecaseMetadataWithDefaults() *UsecaseMetadata`

NewUsecaseMetadataWithDefaults instantiates a new UsecaseMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnnualFee

`func (o *UsecaseMetadata) GetAnnualFee() float32`

GetAnnualFee returns the AnnualFee field if non-nil, zero value otherwise.

### GetAnnualFeeOk

`func (o *UsecaseMetadata) GetAnnualFeeOk() (*float32, bool)`

GetAnnualFeeOk returns a tuple with the AnnualFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualFee

`func (o *UsecaseMetadata) SetAnnualFee(v float32)`

SetAnnualFee sets AnnualFee field to given value.

### HasAnnualFee

`func (o *UsecaseMetadata) HasAnnualFee() bool`

HasAnnualFee returns a boolean if a field has been set.

### GetMaxSubUsecases

`func (o *UsecaseMetadata) GetMaxSubUsecases() int32`

GetMaxSubUsecases returns the MaxSubUsecases field if non-nil, zero value otherwise.

### GetMaxSubUsecasesOk

`func (o *UsecaseMetadata) GetMaxSubUsecasesOk() (*int32, bool)`

GetMaxSubUsecasesOk returns a tuple with the MaxSubUsecases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSubUsecases

`func (o *UsecaseMetadata) SetMaxSubUsecases(v int32)`

SetMaxSubUsecases sets MaxSubUsecases field to given value.

### HasMaxSubUsecases

`func (o *UsecaseMetadata) HasMaxSubUsecases() bool`

HasMaxSubUsecases returns a boolean if a field has been set.

### GetMinSubUsecases

`func (o *UsecaseMetadata) GetMinSubUsecases() int32`

GetMinSubUsecases returns the MinSubUsecases field if non-nil, zero value otherwise.

### GetMinSubUsecasesOk

`func (o *UsecaseMetadata) GetMinSubUsecasesOk() (*int32, bool)`

GetMinSubUsecasesOk returns a tuple with the MinSubUsecases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSubUsecases

`func (o *UsecaseMetadata) SetMinSubUsecases(v int32)`

SetMinSubUsecases sets MinSubUsecases field to given value.

### HasMinSubUsecases

`func (o *UsecaseMetadata) HasMinSubUsecases() bool`

HasMinSubUsecases returns a boolean if a field has been set.

### GetMnoMetadata

`func (o *UsecaseMetadata) GetMnoMetadata() map[string]interface{}`

GetMnoMetadata returns the MnoMetadata field if non-nil, zero value otherwise.

### GetMnoMetadataOk

`func (o *UsecaseMetadata) GetMnoMetadataOk() (*map[string]interface{}, bool)`

GetMnoMetadataOk returns a tuple with the MnoMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMnoMetadata

`func (o *UsecaseMetadata) SetMnoMetadata(v map[string]interface{})`

SetMnoMetadata sets MnoMetadata field to given value.

### HasMnoMetadata

`func (o *UsecaseMetadata) HasMnoMetadata() bool`

HasMnoMetadata returns a boolean if a field has been set.

### GetMonthlyFee

`func (o *UsecaseMetadata) GetMonthlyFee() float32`

GetMonthlyFee returns the MonthlyFee field if non-nil, zero value otherwise.

### GetMonthlyFeeOk

`func (o *UsecaseMetadata) GetMonthlyFeeOk() (*float32, bool)`

GetMonthlyFeeOk returns a tuple with the MonthlyFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyFee

`func (o *UsecaseMetadata) SetMonthlyFee(v float32)`

SetMonthlyFee sets MonthlyFee field to given value.

### HasMonthlyFee

`func (o *UsecaseMetadata) HasMonthlyFee() bool`

HasMonthlyFee returns a boolean if a field has been set.

### GetQuarterlyFee

`func (o *UsecaseMetadata) GetQuarterlyFee() float32`

GetQuarterlyFee returns the QuarterlyFee field if non-nil, zero value otherwise.

### GetQuarterlyFeeOk

`func (o *UsecaseMetadata) GetQuarterlyFeeOk() (*float32, bool)`

GetQuarterlyFeeOk returns a tuple with the QuarterlyFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuarterlyFee

`func (o *UsecaseMetadata) SetQuarterlyFee(v float32)`

SetQuarterlyFee sets QuarterlyFee field to given value.

### HasQuarterlyFee

`func (o *UsecaseMetadata) HasQuarterlyFee() bool`

HasQuarterlyFee returns a boolean if a field has been set.

### GetUsecase

`func (o *UsecaseMetadata) GetUsecase() string`

GetUsecase returns the Usecase field if non-nil, zero value otherwise.

### GetUsecaseOk

`func (o *UsecaseMetadata) GetUsecaseOk() (*string, bool)`

GetUsecaseOk returns a tuple with the Usecase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsecase

`func (o *UsecaseMetadata) SetUsecase(v string)`

SetUsecase sets Usecase field to given value.

### HasUsecase

`func (o *UsecaseMetadata) HasUsecase() bool`

HasUsecase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


