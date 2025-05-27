# VerificationRequestEgress

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
**VerificationRequestId** | **string** |  | 
**VerificationStatus** | Pointer to [**TFVerificationStatus**](TFVerificationStatus.md) |  | [optional] [default to IN_PROGRESS]

## Methods

### NewVerificationRequestEgress

`func NewVerificationRequestEgress(businessName string, corporateWebsite string, businessAddr1 string, businessCity string, businessState string, businessZip string, businessContactFirstName string, businessContactLastName string, businessContactEmail string, businessContactPhone string, messageVolume Volume, phoneNumbers []TFPhoneNumber, useCase UseCaseCategories, useCaseSummary string, productionMessageContent string, optInWorkflow string, optInWorkflowImageURLs []Url, additionalInformation string, isvReseller string, id string, verificationRequestId string, ) *VerificationRequestEgress`

NewVerificationRequestEgress instantiates a new VerificationRequestEgress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVerificationRequestEgressWithDefaults

`func NewVerificationRequestEgressWithDefaults() *VerificationRequestEgress`

NewVerificationRequestEgressWithDefaults instantiates a new VerificationRequestEgress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessName

`func (o *VerificationRequestEgress) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *VerificationRequestEgress) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *VerificationRequestEgress) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.


### GetCorporateWebsite

`func (o *VerificationRequestEgress) GetCorporateWebsite() string`

GetCorporateWebsite returns the CorporateWebsite field if non-nil, zero value otherwise.

### GetCorporateWebsiteOk

`func (o *VerificationRequestEgress) GetCorporateWebsiteOk() (*string, bool)`

GetCorporateWebsiteOk returns a tuple with the CorporateWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateWebsite

`func (o *VerificationRequestEgress) SetCorporateWebsite(v string)`

SetCorporateWebsite sets CorporateWebsite field to given value.


### GetBusinessAddr1

`func (o *VerificationRequestEgress) GetBusinessAddr1() string`

GetBusinessAddr1 returns the BusinessAddr1 field if non-nil, zero value otherwise.

### GetBusinessAddr1Ok

`func (o *VerificationRequestEgress) GetBusinessAddr1Ok() (*string, bool)`

GetBusinessAddr1Ok returns a tuple with the BusinessAddr1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessAddr1

`func (o *VerificationRequestEgress) SetBusinessAddr1(v string)`

SetBusinessAddr1 sets BusinessAddr1 field to given value.


### GetBusinessAddr2

`func (o *VerificationRequestEgress) GetBusinessAddr2() string`

GetBusinessAddr2 returns the BusinessAddr2 field if non-nil, zero value otherwise.

### GetBusinessAddr2Ok

`func (o *VerificationRequestEgress) GetBusinessAddr2Ok() (*string, bool)`

GetBusinessAddr2Ok returns a tuple with the BusinessAddr2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessAddr2

`func (o *VerificationRequestEgress) SetBusinessAddr2(v string)`

SetBusinessAddr2 sets BusinessAddr2 field to given value.

### HasBusinessAddr2

`func (o *VerificationRequestEgress) HasBusinessAddr2() bool`

HasBusinessAddr2 returns a boolean if a field has been set.

### GetBusinessCity

`func (o *VerificationRequestEgress) GetBusinessCity() string`

GetBusinessCity returns the BusinessCity field if non-nil, zero value otherwise.

### GetBusinessCityOk

`func (o *VerificationRequestEgress) GetBusinessCityOk() (*string, bool)`

GetBusinessCityOk returns a tuple with the BusinessCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessCity

`func (o *VerificationRequestEgress) SetBusinessCity(v string)`

SetBusinessCity sets BusinessCity field to given value.


### GetBusinessState

`func (o *VerificationRequestEgress) GetBusinessState() string`

GetBusinessState returns the BusinessState field if non-nil, zero value otherwise.

### GetBusinessStateOk

`func (o *VerificationRequestEgress) GetBusinessStateOk() (*string, bool)`

GetBusinessStateOk returns a tuple with the BusinessState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessState

`func (o *VerificationRequestEgress) SetBusinessState(v string)`

SetBusinessState sets BusinessState field to given value.


### GetBusinessZip

`func (o *VerificationRequestEgress) GetBusinessZip() string`

GetBusinessZip returns the BusinessZip field if non-nil, zero value otherwise.

### GetBusinessZipOk

`func (o *VerificationRequestEgress) GetBusinessZipOk() (*string, bool)`

GetBusinessZipOk returns a tuple with the BusinessZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessZip

`func (o *VerificationRequestEgress) SetBusinessZip(v string)`

SetBusinessZip sets BusinessZip field to given value.


### GetBusinessContactFirstName

`func (o *VerificationRequestEgress) GetBusinessContactFirstName() string`

GetBusinessContactFirstName returns the BusinessContactFirstName field if non-nil, zero value otherwise.

### GetBusinessContactFirstNameOk

`func (o *VerificationRequestEgress) GetBusinessContactFirstNameOk() (*string, bool)`

GetBusinessContactFirstNameOk returns a tuple with the BusinessContactFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessContactFirstName

`func (o *VerificationRequestEgress) SetBusinessContactFirstName(v string)`

SetBusinessContactFirstName sets BusinessContactFirstName field to given value.


### GetBusinessContactLastName

`func (o *VerificationRequestEgress) GetBusinessContactLastName() string`

GetBusinessContactLastName returns the BusinessContactLastName field if non-nil, zero value otherwise.

### GetBusinessContactLastNameOk

`func (o *VerificationRequestEgress) GetBusinessContactLastNameOk() (*string, bool)`

GetBusinessContactLastNameOk returns a tuple with the BusinessContactLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessContactLastName

`func (o *VerificationRequestEgress) SetBusinessContactLastName(v string)`

SetBusinessContactLastName sets BusinessContactLastName field to given value.


### GetBusinessContactEmail

`func (o *VerificationRequestEgress) GetBusinessContactEmail() string`

GetBusinessContactEmail returns the BusinessContactEmail field if non-nil, zero value otherwise.

### GetBusinessContactEmailOk

`func (o *VerificationRequestEgress) GetBusinessContactEmailOk() (*string, bool)`

GetBusinessContactEmailOk returns a tuple with the BusinessContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessContactEmail

`func (o *VerificationRequestEgress) SetBusinessContactEmail(v string)`

SetBusinessContactEmail sets BusinessContactEmail field to given value.


### GetBusinessContactPhone

`func (o *VerificationRequestEgress) GetBusinessContactPhone() string`

GetBusinessContactPhone returns the BusinessContactPhone field if non-nil, zero value otherwise.

### GetBusinessContactPhoneOk

`func (o *VerificationRequestEgress) GetBusinessContactPhoneOk() (*string, bool)`

GetBusinessContactPhoneOk returns a tuple with the BusinessContactPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessContactPhone

`func (o *VerificationRequestEgress) SetBusinessContactPhone(v string)`

SetBusinessContactPhone sets BusinessContactPhone field to given value.


### GetMessageVolume

`func (o *VerificationRequestEgress) GetMessageVolume() Volume`

GetMessageVolume returns the MessageVolume field if non-nil, zero value otherwise.

### GetMessageVolumeOk

`func (o *VerificationRequestEgress) GetMessageVolumeOk() (*Volume, bool)`

GetMessageVolumeOk returns a tuple with the MessageVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageVolume

`func (o *VerificationRequestEgress) SetMessageVolume(v Volume)`

SetMessageVolume sets MessageVolume field to given value.


### GetPhoneNumbers

`func (o *VerificationRequestEgress) GetPhoneNumbers() []TFPhoneNumber`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *VerificationRequestEgress) GetPhoneNumbersOk() (*[]TFPhoneNumber, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *VerificationRequestEgress) SetPhoneNumbers(v []TFPhoneNumber)`

SetPhoneNumbers sets PhoneNumbers field to given value.


### GetUseCase

`func (o *VerificationRequestEgress) GetUseCase() UseCaseCategories`

GetUseCase returns the UseCase field if non-nil, zero value otherwise.

### GetUseCaseOk

`func (o *VerificationRequestEgress) GetUseCaseOk() (*UseCaseCategories, bool)`

GetUseCaseOk returns a tuple with the UseCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCase

`func (o *VerificationRequestEgress) SetUseCase(v UseCaseCategories)`

SetUseCase sets UseCase field to given value.


### GetUseCaseSummary

`func (o *VerificationRequestEgress) GetUseCaseSummary() string`

GetUseCaseSummary returns the UseCaseSummary field if non-nil, zero value otherwise.

### GetUseCaseSummaryOk

`func (o *VerificationRequestEgress) GetUseCaseSummaryOk() (*string, bool)`

GetUseCaseSummaryOk returns a tuple with the UseCaseSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCaseSummary

`func (o *VerificationRequestEgress) SetUseCaseSummary(v string)`

SetUseCaseSummary sets UseCaseSummary field to given value.


### GetProductionMessageContent

`func (o *VerificationRequestEgress) GetProductionMessageContent() string`

GetProductionMessageContent returns the ProductionMessageContent field if non-nil, zero value otherwise.

### GetProductionMessageContentOk

`func (o *VerificationRequestEgress) GetProductionMessageContentOk() (*string, bool)`

GetProductionMessageContentOk returns a tuple with the ProductionMessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionMessageContent

`func (o *VerificationRequestEgress) SetProductionMessageContent(v string)`

SetProductionMessageContent sets ProductionMessageContent field to given value.


### GetOptInWorkflow

`func (o *VerificationRequestEgress) GetOptInWorkflow() string`

GetOptInWorkflow returns the OptInWorkflow field if non-nil, zero value otherwise.

### GetOptInWorkflowOk

`func (o *VerificationRequestEgress) GetOptInWorkflowOk() (*string, bool)`

GetOptInWorkflowOk returns a tuple with the OptInWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptInWorkflow

`func (o *VerificationRequestEgress) SetOptInWorkflow(v string)`

SetOptInWorkflow sets OptInWorkflow field to given value.


### GetOptInWorkflowImageURLs

`func (o *VerificationRequestEgress) GetOptInWorkflowImageURLs() []Url`

GetOptInWorkflowImageURLs returns the OptInWorkflowImageURLs field if non-nil, zero value otherwise.

### GetOptInWorkflowImageURLsOk

`func (o *VerificationRequestEgress) GetOptInWorkflowImageURLsOk() (*[]Url, bool)`

GetOptInWorkflowImageURLsOk returns a tuple with the OptInWorkflowImageURLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptInWorkflowImageURLs

`func (o *VerificationRequestEgress) SetOptInWorkflowImageURLs(v []Url)`

SetOptInWorkflowImageURLs sets OptInWorkflowImageURLs field to given value.


### GetAdditionalInformation

`func (o *VerificationRequestEgress) GetAdditionalInformation() string`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *VerificationRequestEgress) GetAdditionalInformationOk() (*string, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *VerificationRequestEgress) SetAdditionalInformation(v string)`

SetAdditionalInformation sets AdditionalInformation field to given value.


### GetIsvReseller

`func (o *VerificationRequestEgress) GetIsvReseller() string`

GetIsvReseller returns the IsvReseller field if non-nil, zero value otherwise.

### GetIsvResellerOk

`func (o *VerificationRequestEgress) GetIsvResellerOk() (*string, bool)`

GetIsvResellerOk returns a tuple with the IsvReseller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsvReseller

`func (o *VerificationRequestEgress) SetIsvReseller(v string)`

SetIsvReseller sets IsvReseller field to given value.


### GetWebhookUrl

`func (o *VerificationRequestEgress) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *VerificationRequestEgress) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *VerificationRequestEgress) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *VerificationRequestEgress) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetId

`func (o *VerificationRequestEgress) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VerificationRequestEgress) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VerificationRequestEgress) SetId(v string)`

SetId sets Id field to given value.


### GetVerificationRequestId

`func (o *VerificationRequestEgress) GetVerificationRequestId() string`

GetVerificationRequestId returns the VerificationRequestId field if non-nil, zero value otherwise.

### GetVerificationRequestIdOk

`func (o *VerificationRequestEgress) GetVerificationRequestIdOk() (*string, bool)`

GetVerificationRequestIdOk returns a tuple with the VerificationRequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationRequestId

`func (o *VerificationRequestEgress) SetVerificationRequestId(v string)`

SetVerificationRequestId sets VerificationRequestId field to given value.


### GetVerificationStatus

`func (o *VerificationRequestEgress) GetVerificationStatus() TFVerificationStatus`

GetVerificationStatus returns the VerificationStatus field if non-nil, zero value otherwise.

### GetVerificationStatusOk

`func (o *VerificationRequestEgress) GetVerificationStatusOk() (*TFVerificationStatus, bool)`

GetVerificationStatusOk returns a tuple with the VerificationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationStatus

`func (o *VerificationRequestEgress) SetVerificationStatus(v TFVerificationStatus)`

SetVerificationStatus sets VerificationStatus field to given value.

### HasVerificationStatus

`func (o *VerificationRequestEgress) HasVerificationStatus() bool`

HasVerificationStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


