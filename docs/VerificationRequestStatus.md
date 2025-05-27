# VerificationRequestStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessName** | **string** |  | 
**CorporateWebsite** | **string** |  | 
**BusinessAddr1** | **string** |  | 
**BusinessAddr2** | Pointer to **string** |  | [optional] 
**BusinessCity** | **string** |  | 
**BusinessState** | **string** |  | 
**BusinessZip** | **string** |  | 
**BusinessContactFirstName** | **string** |  | 
**BusinessContactLastName** | **string** |  | 
**BusinessContactEmail** | **string** |  | 
**BusinessContactPhone** | **string** |  | 
**MessageVolume** | [**Volume**](Volume.md) | One of the following exact values: 10; 100; 1,000; 10,000; 100,000; 250,000; 500,000; 750,000; 1,000,000; 5,000,000; 10,000,000+ | 
**PhoneNumbers** | [**[]TFPhoneNumber**](TFPhoneNumber.md) |  | 
**UseCase** | [**UseCaseCategories**](UseCaseCategories.md) |  | 
**UseCaseSummary** | **string** |  | 
**ProductionMessageContent** | **string** |  | 
**OptInWorkflow** | **string** |  | 
**OptInWorkflowImageURLs** | [**[]Url**](Url.md) |  | 
**AdditionalInformation** | **string** |  | 
**IsvReseller** | **string** |  | 
**WebhookUrl** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 
**VerificationStatus** | [**TFVerificationStatus**](TFVerificationStatus.md) |  | 
**Reason** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewVerificationRequestStatus

`func NewVerificationRequestStatus(businessName string, corporateWebsite string, businessAddr1 string, businessCity string, businessState string, businessZip string, businessContactFirstName string, businessContactLastName string, businessContactEmail string, businessContactPhone string, messageVolume Volume, phoneNumbers []TFPhoneNumber, useCase UseCaseCategories, useCaseSummary string, productionMessageContent string, optInWorkflow string, optInWorkflowImageURLs []Url, additionalInformation string, isvReseller string, id string, verificationStatus TFVerificationStatus, ) *VerificationRequestStatus`

NewVerificationRequestStatus instantiates a new VerificationRequestStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationRequestStatusWithDefaults

`func NewVerificationRequestStatusWithDefaults() *VerificationRequestStatus`

NewVerificationRequestStatusWithDefaults instantiates a new VerificationRequestStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessName

`func (o *VerificationRequestStatus) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *VerificationRequestStatus) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *VerificationRequestStatus) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.


### GetCorporateWebsite

`func (o *VerificationRequestStatus) GetCorporateWebsite() string`

GetCorporateWebsite returns the CorporateWebsite field if non-nil, zero value otherwise.

### GetCorporateWebsiteOk

`func (o *VerificationRequestStatus) GetCorporateWebsiteOk() (*string, bool)`

GetCorporateWebsiteOk returns a tuple with the CorporateWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateWebsite

`func (o *VerificationRequestStatus) SetCorporateWebsite(v string)`

SetCorporateWebsite sets CorporateWebsite field to given value.


### GetBusinessAddr1

`func (o *VerificationRequestStatus) GetBusinessAddr1() string`

GetBusinessAddr1 returns the BusinessAddr1 field if non-nil, zero value otherwise.

### GetBusinessAddr1Ok

`func (o *VerificationRequestStatus) GetBusinessAddr1Ok() (*string, bool)`

GetBusinessAddr1Ok returns a tuple with the BusinessAddr1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessAddr1

`func (o *VerificationRequestStatus) SetBusinessAddr1(v string)`

SetBusinessAddr1 sets BusinessAddr1 field to given value.


### GetBusinessAddr2

`func (o *VerificationRequestStatus) GetBusinessAddr2() string`

GetBusinessAddr2 returns the BusinessAddr2 field if non-nil, zero value otherwise.

### GetBusinessAddr2Ok

`func (o *VerificationRequestStatus) GetBusinessAddr2Ok() (*string, bool)`

GetBusinessAddr2Ok returns a tuple with the BusinessAddr2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessAddr2

`func (o *VerificationRequestStatus) SetBusinessAddr2(v string)`

SetBusinessAddr2 sets BusinessAddr2 field to given value.

### HasBusinessAddr2

`func (o *VerificationRequestStatus) HasBusinessAddr2() bool`

HasBusinessAddr2 returns a boolean if a field has been set.

### GetBusinessCity

`func (o *VerificationRequestStatus) GetBusinessCity() string`

GetBusinessCity returns the BusinessCity field if non-nil, zero value otherwise.

### GetBusinessCityOk

`func (o *VerificationRequestStatus) GetBusinessCityOk() (*string, bool)`

GetBusinessCityOk returns a tuple with the BusinessCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessCity

`func (o *VerificationRequestStatus) SetBusinessCity(v string)`

SetBusinessCity sets BusinessCity field to given value.


### GetBusinessState

`func (o *VerificationRequestStatus) GetBusinessState() string`

GetBusinessState returns the BusinessState field if non-nil, zero value otherwise.

### GetBusinessStateOk

`func (o *VerificationRequestStatus) GetBusinessStateOk() (*string, bool)`

GetBusinessStateOk returns a tuple with the BusinessState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessState

`func (o *VerificationRequestStatus) SetBusinessState(v string)`

SetBusinessState sets BusinessState field to given value.


### GetBusinessZip

`func (o *VerificationRequestStatus) GetBusinessZip() string`

GetBusinessZip returns the BusinessZip field if non-nil, zero value otherwise.

### GetBusinessZipOk

`func (o *VerificationRequestStatus) GetBusinessZipOk() (*string, bool)`

GetBusinessZipOk returns a tuple with the BusinessZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessZip

`func (o *VerificationRequestStatus) SetBusinessZip(v string)`

SetBusinessZip sets BusinessZip field to given value.


### GetBusinessContactFirstName

`func (o *VerificationRequestStatus) GetBusinessContactFirstName() string`

GetBusinessContactFirstName returns the BusinessContactFirstName field if non-nil, zero value otherwise.

### GetBusinessContactFirstNameOk

`func (o *VerificationRequestStatus) GetBusinessContactFirstNameOk() (*string, bool)`

GetBusinessContactFirstNameOk returns a tuple with the BusinessContactFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessContactFirstName

`func (o *VerificationRequestStatus) SetBusinessContactFirstName(v string)`

SetBusinessContactFirstName sets BusinessContactFirstName field to given value.


### GetBusinessContactLastName

`func (o *VerificationRequestStatus) GetBusinessContactLastName() string`

GetBusinessContactLastName returns the BusinessContactLastName field if non-nil, zero value otherwise.

### GetBusinessContactLastNameOk

`func (o *VerificationRequestStatus) GetBusinessContactLastNameOk() (*string, bool)`

GetBusinessContactLastNameOk returns a tuple with the BusinessContactLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessContactLastName

`func (o *VerificationRequestStatus) SetBusinessContactLastName(v string)`

SetBusinessContactLastName sets BusinessContactLastName field to given value.


### GetBusinessContactEmail

`func (o *VerificationRequestStatus) GetBusinessContactEmail() string`

GetBusinessContactEmail returns the BusinessContactEmail field if non-nil, zero value otherwise.

### GetBusinessContactEmailOk

`func (o *VerificationRequestStatus) GetBusinessContactEmailOk() (*string, bool)`

GetBusinessContactEmailOk returns a tuple with the BusinessContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessContactEmail

`func (o *VerificationRequestStatus) SetBusinessContactEmail(v string)`

SetBusinessContactEmail sets BusinessContactEmail field to given value.


### GetBusinessContactPhone

`func (o *VerificationRequestStatus) GetBusinessContactPhone() string`

GetBusinessContactPhone returns the BusinessContactPhone field if non-nil, zero value otherwise.

### GetBusinessContactPhoneOk

`func (o *VerificationRequestStatus) GetBusinessContactPhoneOk() (*string, bool)`

GetBusinessContactPhoneOk returns a tuple with the BusinessContactPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessContactPhone

`func (o *VerificationRequestStatus) SetBusinessContactPhone(v string)`

SetBusinessContactPhone sets BusinessContactPhone field to given value.


### GetMessageVolume

`func (o *VerificationRequestStatus) GetMessageVolume() Volume`

GetMessageVolume returns the MessageVolume field if non-nil, zero value otherwise.

### GetMessageVolumeOk

`func (o *VerificationRequestStatus) GetMessageVolumeOk() (*Volume, bool)`

GetMessageVolumeOk returns a tuple with the MessageVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageVolume

`func (o *VerificationRequestStatus) SetMessageVolume(v Volume)`

SetMessageVolume sets MessageVolume field to given value.


### GetPhoneNumbers

`func (o *VerificationRequestStatus) GetPhoneNumbers() []TFPhoneNumber`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *VerificationRequestStatus) GetPhoneNumbersOk() (*[]TFPhoneNumber, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *VerificationRequestStatus) SetPhoneNumbers(v []TFPhoneNumber)`

SetPhoneNumbers sets PhoneNumbers field to given value.


### GetUseCase

`func (o *VerificationRequestStatus) GetUseCase() UseCaseCategories`

GetUseCase returns the UseCase field if non-nil, zero value otherwise.

### GetUseCaseOk

`func (o *VerificationRequestStatus) GetUseCaseOk() (*UseCaseCategories, bool)`

GetUseCaseOk returns a tuple with the UseCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCase

`func (o *VerificationRequestStatus) SetUseCase(v UseCaseCategories)`

SetUseCase sets UseCase field to given value.


### GetUseCaseSummary

`func (o *VerificationRequestStatus) GetUseCaseSummary() string`

GetUseCaseSummary returns the UseCaseSummary field if non-nil, zero value otherwise.

### GetUseCaseSummaryOk

`func (o *VerificationRequestStatus) GetUseCaseSummaryOk() (*string, bool)`

GetUseCaseSummaryOk returns a tuple with the UseCaseSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCaseSummary

`func (o *VerificationRequestStatus) SetUseCaseSummary(v string)`

SetUseCaseSummary sets UseCaseSummary field to given value.


### GetProductionMessageContent

`func (o *VerificationRequestStatus) GetProductionMessageContent() string`

GetProductionMessageContent returns the ProductionMessageContent field if non-nil, zero value otherwise.

### GetProductionMessageContentOk

`func (o *VerificationRequestStatus) GetProductionMessageContentOk() (*string, bool)`

GetProductionMessageContentOk returns a tuple with the ProductionMessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionMessageContent

`func (o *VerificationRequestStatus) SetProductionMessageContent(v string)`

SetProductionMessageContent sets ProductionMessageContent field to given value.


### GetOptInWorkflow

`func (o *VerificationRequestStatus) GetOptInWorkflow() string`

GetOptInWorkflow returns the OptInWorkflow field if non-nil, zero value otherwise.

### GetOptInWorkflowOk

`func (o *VerificationRequestStatus) GetOptInWorkflowOk() (*string, bool)`

GetOptInWorkflowOk returns a tuple with the OptInWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptInWorkflow

`func (o *VerificationRequestStatus) SetOptInWorkflow(v string)`

SetOptInWorkflow sets OptInWorkflow field to given value.


### GetOptInWorkflowImageURLs

`func (o *VerificationRequestStatus) GetOptInWorkflowImageURLs() []Url`

GetOptInWorkflowImageURLs returns the OptInWorkflowImageURLs field if non-nil, zero value otherwise.

### GetOptInWorkflowImageURLsOk

`func (o *VerificationRequestStatus) GetOptInWorkflowImageURLsOk() (*[]Url, bool)`

GetOptInWorkflowImageURLsOk returns a tuple with the OptInWorkflowImageURLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptInWorkflowImageURLs

`func (o *VerificationRequestStatus) SetOptInWorkflowImageURLs(v []Url)`

SetOptInWorkflowImageURLs sets OptInWorkflowImageURLs field to given value.


### GetAdditionalInformation

`func (o *VerificationRequestStatus) GetAdditionalInformation() string`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *VerificationRequestStatus) GetAdditionalInformationOk() (*string, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *VerificationRequestStatus) SetAdditionalInformation(v string)`

SetAdditionalInformation sets AdditionalInformation field to given value.


### GetIsvReseller

`func (o *VerificationRequestStatus) GetIsvReseller() string`

GetIsvReseller returns the IsvReseller field if non-nil, zero value otherwise.

### GetIsvResellerOk

`func (o *VerificationRequestStatus) GetIsvResellerOk() (*string, bool)`

GetIsvResellerOk returns a tuple with the IsvReseller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsvReseller

`func (o *VerificationRequestStatus) SetIsvReseller(v string)`

SetIsvReseller sets IsvReseller field to given value.


### GetWebhookUrl

`func (o *VerificationRequestStatus) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *VerificationRequestStatus) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *VerificationRequestStatus) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *VerificationRequestStatus) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetId

`func (o *VerificationRequestStatus) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerificationRequestStatus) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerificationRequestStatus) SetId(v string)`

SetId sets Id field to given value.


### GetVerificationStatus

`func (o *VerificationRequestStatus) GetVerificationStatus() TFVerificationStatus`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *VerificationRequestStatus) GetVerificationStatusOk() (*TFVerificationStatus, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *VerificationRequestStatus) SetVerificationStatus(v TFVerificationStatus)`

SetVerificationStatus sets VerificationStatus field to given value.


### GetReason

`func (o *VerificationRequestStatus) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *VerificationRequestStatus) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *VerificationRequestStatus) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *VerificationRequestStatus) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetCreatedAt

`func (o *VerificationRequestStatus) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *VerificationRequestStatus) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *VerificationRequestStatus) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *VerificationRequestStatus) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *VerificationRequestStatus) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *VerificationRequestStatus) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *VerificationRequestStatus) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *VerificationRequestStatus) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


