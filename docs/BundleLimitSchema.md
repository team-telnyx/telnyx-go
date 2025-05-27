# BundleLimitSchema

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Service** | **string** |  | 
**Metric** | **string** |  | 
**Limit** | Pointer to **int32** |  | [optional] 
**Rate** | Pointer to **string** |  | [optional] 
**CountryIso** | Pointer to **string** |  | [optional] 
**CountryCode** | Pointer to **int32** |  | [optional] 
**Country** | Pointer to **string** | Use country_iso instead | [optional] 
**Direction** | Pointer to [**BundleLimitDirection**](BundleLimitDirection.md) |  | [optional] 
**Types** | Pointer to **[]string** |  | [optional] 
**CreatedAt** | **string** |  | 
**UpdatedAt** | **string** |  | 
**BillingService** | Pointer to **string** |  | [optional] 

## Methods

### NewBundleLimitSchema

`func NewBundleLimitSchema(id string, service string, metric string, createdAt string, updatedAt string, ) *BundleLimitSchema`

NewBundleLimitSchema instantiates a new BundleLimitSchema object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBundleLimitSchemaWithDefaults

`func NewBundleLimitSchemaWithDefaults() *BundleLimitSchema`

NewBundleLimitSchemaWithDefaults instantiates a new BundleLimitSchema object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *BundleLimitSchema) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BundleLimitSchema) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BundleLimitSchema) SetId(v string)`

SetId sets Id field to given value.


### GetService

`func (o *BundleLimitSchema) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *BundleLimitSchema) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *BundleLimitSchema) SetService(v string)`

SetService sets Service field to given value.


### GetMetric

`func (o *BundleLimitSchema) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *BundleLimitSchema) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *BundleLimitSchema) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetLimit

`func (o *BundleLimitSchema) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *BundleLimitSchema) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *BundleLimitSchema) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *BundleLimitSchema) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetRate

`func (o *BundleLimitSchema) GetRate() string`

GetRate returns the Rate field if non-nil, zero value otherwise.

### GetRateOk

`func (o *BundleLimitSchema) GetRateOk() (*string, bool)`

GetRateOk returns a tuple with the Rate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRate

`func (o *BundleLimitSchema) SetRate(v string)`

SetRate sets Rate field to given value.

### HasRate

`func (o *BundleLimitSchema) HasRate() bool`

HasRate returns a boolean if a field has been set.

### GetCountryIso

`func (o *BundleLimitSchema) GetCountryIso() string`

GetCountryIso returns the CountryIso field if non-nil, zero value otherwise.

### GetCountryIsoOk

`func (o *BundleLimitSchema) GetCountryIsoOk() (*string, bool)`

GetCountryIsoOk returns a tuple with the CountryIso field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryIso

`func (o *BundleLimitSchema) SetCountryIso(v string)`

SetCountryIso sets CountryIso field to given value.

### HasCountryIso

`func (o *BundleLimitSchema) HasCountryIso() bool`

HasCountryIso returns a boolean if a field has been set.

### GetCountryCode

`func (o *BundleLimitSchema) GetCountryCode() int32`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *BundleLimitSchema) GetCountryCodeOk() (*int32, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *BundleLimitSchema) SetCountryCode(v int32)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *BundleLimitSchema) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetCountry

`func (o *BundleLimitSchema) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *BundleLimitSchema) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *BundleLimitSchema) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *BundleLimitSchema) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetDirection

`func (o *BundleLimitSchema) GetDirection() BundleLimitDirection`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *BundleLimitSchema) GetDirectionOk() (*BundleLimitDirection, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *BundleLimitSchema) SetDirection(v BundleLimitDirection)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *BundleLimitSchema) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetTypes

`func (o *BundleLimitSchema) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *BundleLimitSchema) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *BundleLimitSchema) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *BundleLimitSchema) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetCreatedAt

`func (o *BundleLimitSchema) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *BundleLimitSchema) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *BundleLimitSchema) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *BundleLimitSchema) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *BundleLimitSchema) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *BundleLimitSchema) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetBillingService

`func (o *BundleLimitSchema) GetBillingService() string`

GetBillingService returns the BillingService field if non-nil, zero value otherwise.

### GetBillingServiceOk

`func (o *BundleLimitSchema) GetBillingServiceOk() (*string, bool)`

GetBillingServiceOk returns a tuple with the BillingService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingService

`func (o *BundleLimitSchema) SetBillingService(v string)`

SetBillingService sets BillingService field to given value.

### HasBillingService

`func (o *BundleLimitSchema) HasBillingService() bool`

HasBillingService returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


