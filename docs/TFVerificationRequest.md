# TFVerificationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessName** | **string** | Name of the business; there are no specific formatting requirements | 
**CorporateWebsite** | **string** | A URL, including the scheme, pointing to the corporate website | 
**BusinessAddr1** | **string** | Line 1 of the business address | 
**BusinessAddr2** | Pointer to **string** | Line 2 of the business address | [optional] 
**BusinessCity** | **string** | The city of the business address; the first letter should be capitalized | 
**BusinessState** | **string** | The full name of the state (not the 2 letter code) of the business address; the first letter should be capitalized | 
**BusinessZip** | **string** | The ZIP code of the business address | 
**BusinessContactFirstName** | **string** | First name of the business contact; there are no specific requirements on formatting | 
**BusinessContactLastName** | **string** | Last name of the business contact; there are no specific requirements on formatting | 
**BusinessContactEmail** | **string** | The email address of the business contact | 
**BusinessContactPhone** | **string** | The phone number of the business contact in E.164 format | 
**MessageVolume** | [**Volume**](Volume.md) | Estimated monthly volume of messages from the given phone numbers | 
**PhoneNumbers** | [**[]TFPhoneNumber**](TFPhoneNumber.md) | The phone numbers to request the verification of | 
**UseCase** | [**UseCaseCategories**](UseCaseCategories.md) | Machine-readable use-case for the phone numbers | 
**UseCaseSummary** | **string** | Human-readable summary of the desired use-case | 
**ProductionMessageContent** | **string** | An example of a message that will be sent from the given phone numbers | 
**OptInWorkflow** | **string** | Human-readable description of how end users will opt into receiving messages from the given phone numbers | 
**OptInWorkflowImageURLs** | [**[]Url**](Url.md) | Images showing the opt-in workflow | 
**AdditionalInformation** | **string** | Any additional information | 
**IsvReseller** | **string** | ISV name | 
**WebhookUrl** | Pointer to **string** | URL that should receive webhooks relating to this verification request | [optional] 

## Methods

### NewTFVerificationRequest

`func NewTFVerificationRequest(businessName string, corporateWebsite string, businessAddr1 string, businessCity string, businessState string, businessZip string, businessContactFirstName string, businessContactLastName string, businessContactEmail string, businessContactPhone string, messageVolume Volume, phoneNumbers []TFPhoneNumber, useCase UseCaseCategories, useCaseSummary string, productionMessageContent string, optInWorkflow string, optInWorkflowImageURLs []Url, additionalInformation string, isvReseller string, ) *TFVerificationRequest`

NewTFVerificationRequest instantiates a new TFVerificationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTFVerificationRequestWithDefaults

`func NewTFVerificationRequestWithDefaults() *TFVerificationRequest`

NewTFVerificationRequestWithDefaults instantiates a new TFVerificationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessName

`func (o *TFVerificationRequest) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *TFVerificationRequest) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *TFVerificationRequest) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.


### GetCorporateWebsite

`func (o *TFVerificationRequest) GetCorporateWebsite() string`

GetCorporateWebsite returns the CorporateWebsite field if non-nil, zero value otherwise.

### GetCorporateWebsiteOk

`func (o *TFVerificationRequest) GetCorporateWebsiteOk() (*string, bool)`

GetCorporateWebsiteOk returns a tuple with the CorporateWebsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorporateWebsite

`func (o *TFVerificationRequest) SetCorporateWebsite(v string)`

SetCorporateWebsite sets CorporateWebsite field to given value.


### GetBusinessAddr1

`func (o *TFVerificationRequest) GetBusinessAddr1() string`

GetBusinessAddr1 returns the BusinessAddr1 field if non-nil, zero value otherwise.

### GetBusinessAddr1Ok

`func (o *TFVerificationRequest) GetBusinessAddr1Ok() (*string, bool)`

GetBusinessAddr1Ok returns a tuple with the BusinessAddr1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessAddr1

`func (o *TFVerificationRequest) SetBusinessAddr1(v string)`

SetBusinessAddr1 sets BusinessAddr1 field to given value.


### GetBusinessAddr2

`func (o *TFVerificationRequest) GetBusinessAddr2() string`

GetBusinessAddr2 returns the BusinessAddr2 field if non-nil, zero value otherwise.

### GetBusinessAddr2Ok

`func (o *TFVerificationRequest) GetBusinessAddr2Ok() (*string, bool)`

GetBusinessAddr2Ok returns a tuple with the BusinessAddr2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessAddr2

`func (o *TFVerificationRequest) SetBusinessAddr2(v string)`

SetBusinessAddr2 sets BusinessAddr2 field to given value.

### HasBusinessAddr2

`func (o *TFVerificationRequest) HasBusinessAddr2() bool`

HasBusinessAddr2 returns a boolean if a field has been set.

### GetBusinessCity

`func (o *TFVerificationRequest) GetBusinessCity() string`

GetBusinessCity returns the BusinessCity field if non-nil, zero value otherwise.

### GetBusinessCityOk

`func (o *TFVerificationRequest) GetBusinessCityOk() (*string, bool)`

GetBusinessCityOk returns a tuple with the BusinessCity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessCity

`func (o *TFVerificationRequest) SetBusinessCity(v string)`

SetBusinessCity sets BusinessCity field to given value.


### GetBusinessState

`func (o *TFVerificationRequest) GetBusinessState() string`

GetBusinessState returns the BusinessState field if non-nil, zero value otherwise.

### GetBusinessStateOk

`func (o *TFVerificationRequest) GetBusinessStateOk() (*string, bool)`

GetBusinessStateOk returns a tuple with the BusinessState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessState

`func (o *TFVerificationRequest) SetBusinessState(v string)`

SetBusinessState sets BusinessState field to given value.


### GetBusinessZip

`func (o *TFVerificationRequest) GetBusinessZip() string`

GetBusinessZip returns the BusinessZip field if non-nil, zero value otherwise.

### GetBusinessZipOk

`func (o *TFVerificationRequest) GetBusinessZipOk() (*string, bool)`

GetBusinessZipOk returns a tuple with the BusinessZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessZip

`func (o *TFVerificationRequest) SetBusinessZip(v string)`

SetBusinessZip sets BusinessZip field to given value.


### GetBusinessContactFirstName

`func (o *TFVerificationRequest) GetBusinessContactFirstName() string`

GetBusinessContactFirstName returns the BusinessContactFirstName field if non-nil, zero value otherwise.

### GetBusinessContactFirstNameOk

`func (o *TFVerificationRequest) GetBusinessContactFirstNameOk() (*string, bool)`

GetBusinessContactFirstNameOk returns a tuple with the BusinessContactFirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessContactFirstName

`func (o *TFVerificationRequest) SetBusinessContactFirstName(v string)`

SetBusinessContactFirstName sets BusinessContactFirstName field to given value.


### GetBusinessContactLastName

`func (o *TFVerificationRequest) GetBusinessContactLastName() string`

GetBusinessContactLastName returns the BusinessContactLastName field if non-nil, zero value otherwise.

### GetBusinessContactLastNameOk

`func (o *TFVerificationRequest) GetBusinessContactLastNameOk() (*string, bool)`

GetBusinessContactLastNameOk returns a tuple with the BusinessContactLastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessContactLastName

`func (o *TFVerificationRequest) SetBusinessContactLastName(v string)`

SetBusinessContactLastName sets BusinessContactLastName field to given value.


### GetBusinessContactEmail

`func (o *TFVerificationRequest) GetBusinessContactEmail() string`

GetBusinessContactEmail returns the BusinessContactEmail field if non-nil, zero value otherwise.

### GetBusinessContactEmailOk

`func (o *TFVerificationRequest) GetBusinessContactEmailOk() (*string, bool)`

GetBusinessContactEmailOk returns a tuple with the BusinessContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessContactEmail

`func (o *TFVerificationRequest) SetBusinessContactEmail(v string)`

SetBusinessContactEmail sets BusinessContactEmail field to given value.


### GetBusinessContactPhone

`func (o *TFVerificationRequest) GetBusinessContactPhone() string`

GetBusinessContactPhone returns the BusinessContactPhone field if non-nil, zero value otherwise.

### GetBusinessContactPhoneOk

`func (o *TFVerificationRequest) GetBusinessContactPhoneOk() (*string, bool)`

GetBusinessContactPhoneOk returns a tuple with the BusinessContactPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessContactPhone

`func (o *TFVerificationRequest) SetBusinessContactPhone(v string)`

SetBusinessContactPhone sets BusinessContactPhone field to given value.


### GetMessageVolume

`func (o *TFVerificationRequest) GetMessageVolume() Volume`

GetMessageVolume returns the MessageVolume field if non-nil, zero value otherwise.

### GetMessageVolumeOk

`func (o *TFVerificationRequest) GetMessageVolumeOk() (*Volume, bool)`

GetMessageVolumeOk returns a tuple with the MessageVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageVolume

`func (o *TFVerificationRequest) SetMessageVolume(v Volume)`

SetMessageVolume sets MessageVolume field to given value.


### GetPhoneNumbers

`func (o *TFVerificationRequest) GetPhoneNumbers() []TFPhoneNumber`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *TFVerificationRequest) GetPhoneNumbersOk() (*[]TFPhoneNumber, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *TFVerificationRequest) SetPhoneNumbers(v []TFPhoneNumber)`

SetPhoneNumbers sets PhoneNumbers field to given value.


### GetUseCase

`func (o *TFVerificationRequest) GetUseCase() UseCaseCategories`

GetUseCase returns the UseCase field if non-nil, zero value otherwise.

### GetUseCaseOk

`func (o *TFVerificationRequest) GetUseCaseOk() (*UseCaseCategories, bool)`

GetUseCaseOk returns a tuple with the UseCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCase

`func (o *TFVerificationRequest) SetUseCase(v UseCaseCategories)`

SetUseCase sets UseCase field to given value.


### GetUseCaseSummary

`func (o *TFVerificationRequest) GetUseCaseSummary() string`

GetUseCaseSummary returns the UseCaseSummary field if non-nil, zero value otherwise.

### GetUseCaseSummaryOk

`func (o *TFVerificationRequest) GetUseCaseSummaryOk() (*string, bool)`

GetUseCaseSummaryOk returns a tuple with the UseCaseSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCaseSummary

`func (o *TFVerificationRequest) SetUseCaseSummary(v string)`

SetUseCaseSummary sets UseCaseSummary field to given value.


### GetProductionMessageContent

`func (o *TFVerificationRequest) GetProductionMessageContent() string`

GetProductionMessageContent returns the ProductionMessageContent field if non-nil, zero value otherwise.

### GetProductionMessageContentOk

`func (o *TFVerificationRequest) GetProductionMessageContentOk() (*string, bool)`

GetProductionMessageContentOk returns a tuple with the ProductionMessageContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProductionMessageContent

`func (o *TFVerificationRequest) SetProductionMessageContent(v string)`

SetProductionMessageContent sets ProductionMessageContent field to given value.


### GetOptInWorkflow

`func (o *TFVerificationRequest) GetOptInWorkflow() string`

GetOptInWorkflow returns the OptInWorkflow field if non-nil, zero value otherwise.

### GetOptInWorkflowOk

`func (o *TFVerificationRequest) GetOptInWorkflowOk() (*string, bool)`

GetOptInWorkflowOk returns a tuple with the OptInWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptInWorkflow

`func (o *TFVerificationRequest) SetOptInWorkflow(v string)`

SetOptInWorkflow sets OptInWorkflow field to given value.


### GetOptInWorkflowImageURLs

`func (o *TFVerificationRequest) GetOptInWorkflowImageURLs() []Url`

GetOptInWorkflowImageURLs returns the OptInWorkflowImageURLs field if non-nil, zero value otherwise.

### GetOptInWorkflowImageURLsOk

`func (o *TFVerificationRequest) GetOptInWorkflowImageURLsOk() (*[]Url, bool)`

GetOptInWorkflowImageURLsOk returns a tuple with the OptInWorkflowImageURLs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptInWorkflowImageURLs

`func (o *TFVerificationRequest) SetOptInWorkflowImageURLs(v []Url)`

SetOptInWorkflowImageURLs sets OptInWorkflowImageURLs field to given value.


### GetAdditionalInformation

`func (o *TFVerificationRequest) GetAdditionalInformation() string`

GetAdditionalInformation returns the AdditionalInformation field if non-nil, zero value otherwise.

### GetAdditionalInformationOk

`func (o *TFVerificationRequest) GetAdditionalInformationOk() (*string, bool)`

GetAdditionalInformationOk returns a tuple with the AdditionalInformation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalInformation

`func (o *TFVerificationRequest) SetAdditionalInformation(v string)`

SetAdditionalInformation sets AdditionalInformation field to given value.


### GetIsvReseller

`func (o *TFVerificationRequest) GetIsvReseller() string`

GetIsvReseller returns the IsvReseller field if non-nil, zero value otherwise.

### GetIsvResellerOk

`func (o *TFVerificationRequest) GetIsvResellerOk() (*string, bool)`

GetIsvResellerOk returns a tuple with the IsvReseller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsvReseller

`func (o *TFVerificationRequest) SetIsvReseller(v string)`

SetIsvReseller sets IsvReseller field to given value.


### GetWebhookUrl

`func (o *TFVerificationRequest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *TFVerificationRequest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *TFVerificationRequest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *TFVerificationRequest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


