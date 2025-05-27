# PhoneNumberWithMessagingSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**Id** | Pointer to **string** | Identifies the type of resource. | [optional] [readonly] 
**PhoneNumber** | Pointer to **string** | +E.164 formatted phone number. | [optional] [readonly] 
**MessagingProfileId** | Pointer to **NullableString** | Unique identifier for a messaging profile. | [optional] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was updated. | [optional] [readonly] 
**CountryCode** | Pointer to **string** | ISO 3166-1 alpha-2 country code. | [optional] [readonly] 
**Type** | Pointer to **string** | The type of the phone number | [optional] [readonly] 
**Health** | Pointer to [**NumberHealthMetrics**](NumberHealthMetrics.md) |  | [optional] 
**EligibleMessagingProducts** | Pointer to **[]string** | The messaging products that this number can be registered to use | [optional] [readonly] 
**TrafficType** | Pointer to **string** | The messaging traffic or use case for which the number is currently configured. | [optional] [readonly] 
**MessagingProduct** | Pointer to **string** | The messaging product that the number is registered to use | [optional] 
**Features** | Pointer to [**PhoneNumberWithMessagingSettingsFeatures**](PhoneNumberWithMessagingSettingsFeatures.md) |  | [optional] 

## Methods

### NewPhoneNumberWithMessagingSettings

`func NewPhoneNumberWithMessagingSettings() *PhoneNumberWithMessagingSettings`

NewPhoneNumberWithMessagingSettings instantiates a new PhoneNumberWithMessagingSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneNumberWithMessagingSettingsWithDefaults

`func NewPhoneNumberWithMessagingSettingsWithDefaults() *PhoneNumberWithMessagingSettings`

NewPhoneNumberWithMessagingSettingsWithDefaults instantiates a new PhoneNumberWithMessagingSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *PhoneNumberWithMessagingSettings) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PhoneNumberWithMessagingSettings) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PhoneNumberWithMessagingSettings) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PhoneNumberWithMessagingSettings) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetId

`func (o *PhoneNumberWithMessagingSettings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PhoneNumberWithMessagingSettings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PhoneNumberWithMessagingSettings) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PhoneNumberWithMessagingSettings) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPhoneNumber

`func (o *PhoneNumberWithMessagingSettings) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *PhoneNumberWithMessagingSettings) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *PhoneNumberWithMessagingSettings) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *PhoneNumberWithMessagingSettings) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetMessagingProfileId

`func (o *PhoneNumberWithMessagingSettings) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *PhoneNumberWithMessagingSettings) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *PhoneNumberWithMessagingSettings) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.

### HasMessagingProfileId

`func (o *PhoneNumberWithMessagingSettings) HasMessagingProfileId() bool`

HasMessagingProfileId returns a boolean if a field has been set.

### SetMessagingProfileIdNil

`func (o *PhoneNumberWithMessagingSettings) SetMessagingProfileIdNil(b bool)`

 SetMessagingProfileIdNil sets the value for MessagingProfileId to be an explicit nil

### UnsetMessagingProfileId
`func (o *PhoneNumberWithMessagingSettings) UnsetMessagingProfileId()`

UnsetMessagingProfileId ensures that no value is present for MessagingProfileId, not even an explicit nil
### GetCreatedAt

`func (o *PhoneNumberWithMessagingSettings) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PhoneNumberWithMessagingSettings) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PhoneNumberWithMessagingSettings) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PhoneNumberWithMessagingSettings) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PhoneNumberWithMessagingSettings) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PhoneNumberWithMessagingSettings) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PhoneNumberWithMessagingSettings) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PhoneNumberWithMessagingSettings) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetCountryCode

`func (o *PhoneNumberWithMessagingSettings) GetCountryCode() string`

GetCountryCode returns the CountryCode field if non-nil, zero value otherwise.

### GetCountryCodeOk

`func (o *PhoneNumberWithMessagingSettings) GetCountryCodeOk() (*string, bool)`

GetCountryCodeOk returns a tuple with the CountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryCode

`func (o *PhoneNumberWithMessagingSettings) SetCountryCode(v string)`

SetCountryCode sets CountryCode field to given value.

### HasCountryCode

`func (o *PhoneNumberWithMessagingSettings) HasCountryCode() bool`

HasCountryCode returns a boolean if a field has been set.

### GetType

`func (o *PhoneNumberWithMessagingSettings) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PhoneNumberWithMessagingSettings) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PhoneNumberWithMessagingSettings) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PhoneNumberWithMessagingSettings) HasType() bool`

HasType returns a boolean if a field has been set.

### GetHealth

`func (o *PhoneNumberWithMessagingSettings) GetHealth() NumberHealthMetrics`

GetHealth returns the Health field if non-nil, zero value otherwise.

### GetHealthOk

`func (o *PhoneNumberWithMessagingSettings) GetHealthOk() (*NumberHealthMetrics, bool)`

GetHealthOk returns a tuple with the Health field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealth

`func (o *PhoneNumberWithMessagingSettings) SetHealth(v NumberHealthMetrics)`

SetHealth sets Health field to given value.

### HasHealth

`func (o *PhoneNumberWithMessagingSettings) HasHealth() bool`

HasHealth returns a boolean if a field has been set.

### GetEligibleMessagingProducts

`func (o *PhoneNumberWithMessagingSettings) GetEligibleMessagingProducts() []string`

GetEligibleMessagingProducts returns the EligibleMessagingProducts field if non-nil, zero value otherwise.

### GetEligibleMessagingProductsOk

`func (o *PhoneNumberWithMessagingSettings) GetEligibleMessagingProductsOk() (*[]string, bool)`

GetEligibleMessagingProductsOk returns a tuple with the EligibleMessagingProducts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleMessagingProducts

`func (o *PhoneNumberWithMessagingSettings) SetEligibleMessagingProducts(v []string)`

SetEligibleMessagingProducts sets EligibleMessagingProducts field to given value.

### HasEligibleMessagingProducts

`func (o *PhoneNumberWithMessagingSettings) HasEligibleMessagingProducts() bool`

HasEligibleMessagingProducts returns a boolean if a field has been set.

### GetTrafficType

`func (o *PhoneNumberWithMessagingSettings) GetTrafficType() string`

GetTrafficType returns the TrafficType field if non-nil, zero value otherwise.

### GetTrafficTypeOk

`func (o *PhoneNumberWithMessagingSettings) GetTrafficTypeOk() (*string, bool)`

GetTrafficTypeOk returns a tuple with the TrafficType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrafficType

`func (o *PhoneNumberWithMessagingSettings) SetTrafficType(v string)`

SetTrafficType sets TrafficType field to given value.

### HasTrafficType

`func (o *PhoneNumberWithMessagingSettings) HasTrafficType() bool`

HasTrafficType returns a boolean if a field has been set.

### GetMessagingProduct

`func (o *PhoneNumberWithMessagingSettings) GetMessagingProduct() string`

GetMessagingProduct returns the MessagingProduct field if non-nil, zero value otherwise.

### GetMessagingProductOk

`func (o *PhoneNumberWithMessagingSettings) GetMessagingProductOk() (*string, bool)`

GetMessagingProductOk returns a tuple with the MessagingProduct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProduct

`func (o *PhoneNumberWithMessagingSettings) SetMessagingProduct(v string)`

SetMessagingProduct sets MessagingProduct field to given value.

### HasMessagingProduct

`func (o *PhoneNumberWithMessagingSettings) HasMessagingProduct() bool`

HasMessagingProduct returns a boolean if a field has been set.

### GetFeatures

`func (o *PhoneNumberWithMessagingSettings) GetFeatures() PhoneNumberWithMessagingSettingsFeatures`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *PhoneNumberWithMessagingSettings) GetFeaturesOk() (*PhoneNumberWithMessagingSettingsFeatures, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *PhoneNumberWithMessagingSettings) SetFeatures(v PhoneNumberWithMessagingSettingsFeatures)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *PhoneNumberWithMessagingSettings) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


