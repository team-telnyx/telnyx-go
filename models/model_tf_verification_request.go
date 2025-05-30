/*
Telnyx API

SIP trunking, SMS, MMS, Call Control and Telephony Data Services.

API version: 2.0.0
Contact: support@telnyx.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package telnyx

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the TFVerificationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TFVerificationRequest{}

// TFVerificationRequest The body of a tollfree verification request
type TFVerificationRequest struct {
	// Name of the business; there are no specific formatting requirements
	BusinessName string `json:"businessName"`
	// A URL, including the scheme, pointing to the corporate website
	CorporateWebsite string `json:"corporateWebsite"`
	// Line 1 of the business address
	BusinessAddr1 string `json:"businessAddr1"`
	// Line 2 of the business address
	BusinessAddr2 *string `json:"businessAddr2,omitempty"`
	// The city of the business address; the first letter should be capitalized
	BusinessCity string `json:"businessCity"`
	// The full name of the state (not the 2 letter code) of the business address; the first letter should be capitalized
	BusinessState string `json:"businessState"`
	// The ZIP code of the business address
	BusinessZip string `json:"businessZip"`
	// First name of the business contact; there are no specific requirements on formatting
	BusinessContactFirstName string `json:"businessContactFirstName"`
	// Last name of the business contact; there are no specific requirements on formatting
	BusinessContactLastName string `json:"businessContactLastName"`
	// The email address of the business contact
	BusinessContactEmail string `json:"businessContactEmail"`
	// The phone number of the business contact in E.164 format
	BusinessContactPhone string `json:"businessContactPhone"`
	// Estimated monthly volume of messages from the given phone numbers
	MessageVolume Volume `json:"messageVolume"`
	// The phone numbers to request the verification of
	PhoneNumbers []TFPhoneNumber `json:"phoneNumbers"`
	// Machine-readable use-case for the phone numbers
	UseCase UseCaseCategories `json:"useCase"`
	// Human-readable summary of the desired use-case
	UseCaseSummary string `json:"useCaseSummary"`
	// An example of a message that will be sent from the given phone numbers
	ProductionMessageContent string `json:"productionMessageContent"`
	// Human-readable description of how end users will opt into receiving messages from the given phone numbers
	OptInWorkflow string `json:"optInWorkflow"`
	// Images showing the opt-in workflow
	OptInWorkflowImageURLs []Url `json:"optInWorkflowImageURLs"`
	// Any additional information
	AdditionalInformation string `json:"additionalInformation"`
	// ISV name
	IsvReseller string `json:"isvReseller"`
	// URL that should receive webhooks relating to this verification request
	WebhookUrl *string `json:"webhookUrl,omitempty"`
}

type _TFVerificationRequest TFVerificationRequest

// NewTFVerificationRequest instantiates a new TFVerificationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTFVerificationRequest(businessName string, corporateWebsite string, businessAddr1 string, businessCity string, businessState string, businessZip string, businessContactFirstName string, businessContactLastName string, businessContactEmail string, businessContactPhone string, messageVolume Volume, phoneNumbers []TFPhoneNumber, useCase UseCaseCategories, useCaseSummary string, productionMessageContent string, optInWorkflow string, optInWorkflowImageURLs []Url, additionalInformation string, isvReseller string) *TFVerificationRequest {
	this := TFVerificationRequest{}
	this.BusinessName = businessName
	this.CorporateWebsite = corporateWebsite
	this.BusinessAddr1 = businessAddr1
	this.BusinessCity = businessCity
	this.BusinessState = businessState
	this.BusinessZip = businessZip
	this.BusinessContactFirstName = businessContactFirstName
	this.BusinessContactLastName = businessContactLastName
	this.BusinessContactEmail = businessContactEmail
	this.BusinessContactPhone = businessContactPhone
	this.MessageVolume = messageVolume
	this.PhoneNumbers = phoneNumbers
	this.UseCase = useCase
	this.UseCaseSummary = useCaseSummary
	this.ProductionMessageContent = productionMessageContent
	this.OptInWorkflow = optInWorkflow
	this.OptInWorkflowImageURLs = optInWorkflowImageURLs
	this.AdditionalInformation = additionalInformation
	this.IsvReseller = isvReseller
	return &this
}

// NewTFVerificationRequestWithDefaults instantiates a new TFVerificationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTFVerificationRequestWithDefaults() *TFVerificationRequest {
	this := TFVerificationRequest{}
	return &this
}

// GetBusinessName returns the BusinessName field value
func (o *TFVerificationRequest) GetBusinessName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessName
}

// GetBusinessNameOk returns a tuple with the BusinessName field value
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetBusinessNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessName, true
}

// SetBusinessName sets field value
func (o *TFVerificationRequest) SetBusinessName(v string) {
	o.BusinessName = v
}

// GetCorporateWebsite returns the CorporateWebsite field value
func (o *TFVerificationRequest) GetCorporateWebsite() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.CorporateWebsite
}

// GetCorporateWebsiteOk returns a tuple with the CorporateWebsite field value
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetCorporateWebsiteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CorporateWebsite, true
}

// SetCorporateWebsite sets field value
func (o *TFVerificationRequest) SetCorporateWebsite(v string) {
	o.CorporateWebsite = v
}

// GetBusinessAddr1 returns the BusinessAddr1 field value
func (o *TFVerificationRequest) GetBusinessAddr1() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessAddr1
}

// GetBusinessAddr1Ok returns a tuple with the BusinessAddr1 field value
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetBusinessAddr1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessAddr1, true
}

// SetBusinessAddr1 sets field value
func (o *TFVerificationRequest) SetBusinessAddr1(v string) {
	o.BusinessAddr1 = v
}

// GetBusinessAddr2 returns the BusinessAddr2 field value if set, zero value otherwise.
func (o *TFVerificationRequest) GetBusinessAddr2() string {
	if o == nil || IsNil(o.BusinessAddr2) {
		var ret string
		return ret
	}
	return *o.BusinessAddr2
}

// GetBusinessAddr2Ok returns a tuple with the BusinessAddr2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetBusinessAddr2Ok() (*string, bool) {
	if o == nil || IsNil(o.BusinessAddr2) {
		return nil, false
	}
	return o.BusinessAddr2, true
}

// HasBusinessAddr2 returns a boolean if a field has been set.
func (o *TFVerificationRequest) HasBusinessAddr2() bool {
	if o != nil && !IsNil(o.BusinessAddr2) {
		return true
	}

	return false
}

// SetBusinessAddr2 gets a reference to the given string and assigns it to the BusinessAddr2 field.
func (o *TFVerificationRequest) SetBusinessAddr2(v string) {
	o.BusinessAddr2 = &v
}

// GetBusinessCity returns the BusinessCity field value
func (o *TFVerificationRequest) GetBusinessCity() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessCity
}

// GetBusinessCityOk returns a tuple with the BusinessCity field value
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetBusinessCityOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessCity, true
}

// SetBusinessCity sets field value
func (o *TFVerificationRequest) SetBusinessCity(v string) {
	o.BusinessCity = v
}

// GetBusinessState returns the BusinessState field value
func (o *TFVerificationRequest) GetBusinessState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessState
}

// GetBusinessStateOk returns a tuple with the BusinessState field value
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetBusinessStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessState, true
}

// SetBusinessState sets field value
func (o *TFVerificationRequest) SetBusinessState(v string) {
	o.BusinessState = v
}

// GetBusinessZip returns the BusinessZip field value
func (o *TFVerificationRequest) GetBusinessZip() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessZip
}

// GetBusinessZipOk returns a tuple with the BusinessZip field value
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetBusinessZipOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessZip, true
}

// SetBusinessZip sets field value
func (o *TFVerificationRequest) SetBusinessZip(v string) {
	o.BusinessZip = v
}

// GetBusinessContactFirstName returns the BusinessContactFirstName field value
func (o *TFVerificationRequest) GetBusinessContactFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessContactFirstName
}

// GetBusinessContactFirstNameOk returns a tuple with the BusinessContactFirstName field value
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetBusinessContactFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessContactFirstName, true
}

// SetBusinessContactFirstName sets field value
func (o *TFVerificationRequest) SetBusinessContactFirstName(v string) {
	o.BusinessContactFirstName = v
}

// GetBusinessContactLastName returns the BusinessContactLastName field value
func (o *TFVerificationRequest) GetBusinessContactLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessContactLastName
}

// GetBusinessContactLastNameOk returns a tuple with the BusinessContactLastName field value
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetBusinessContactLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessContactLastName, true
}

// SetBusinessContactLastName sets field value
func (o *TFVerificationRequest) SetBusinessContactLastName(v string) {
	o.BusinessContactLastName = v
}

// GetBusinessContactEmail returns the BusinessContactEmail field value
func (o *TFVerificationRequest) GetBusinessContactEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessContactEmail
}

// GetBusinessContactEmailOk returns a tuple with the BusinessContactEmail field value
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetBusinessContactEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessContactEmail, true
}

// SetBusinessContactEmail sets field value
func (o *TFVerificationRequest) SetBusinessContactEmail(v string) {
	o.BusinessContactEmail = v
}

// GetBusinessContactPhone returns the BusinessContactPhone field value
func (o *TFVerificationRequest) GetBusinessContactPhone() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BusinessContactPhone
}

// GetBusinessContactPhoneOk returns a tuple with the BusinessContactPhone field value
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetBusinessContactPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BusinessContactPhone, true
}

// SetBusinessContactPhone sets field value
func (o *TFVerificationRequest) SetBusinessContactPhone(v string) {
	o.BusinessContactPhone = v
}

// GetMessageVolume returns the MessageVolume field value
func (o *TFVerificationRequest) GetMessageVolume() Volume {
	if o == nil {
		var ret Volume
		return ret
	}

	return o.MessageVolume
}

// GetMessageVolumeOk returns a tuple with the MessageVolume field value
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetMessageVolumeOk() (*Volume, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MessageVolume, true
}

// SetMessageVolume sets field value
func (o *TFVerificationRequest) SetMessageVolume(v Volume) {
	o.MessageVolume = v
}

// GetPhoneNumbers returns the PhoneNumbers field value
func (o *TFVerificationRequest) GetPhoneNumbers() []TFPhoneNumber {
	if o == nil {
		var ret []TFPhoneNumber
		return ret
	}

	return o.PhoneNumbers
}

// GetPhoneNumbersOk returns a tuple with the PhoneNumbers field value
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetPhoneNumbersOk() ([]TFPhoneNumber, bool) {
	if o == nil {
		return nil, false
	}
	return o.PhoneNumbers, true
}

// SetPhoneNumbers sets field value
func (o *TFVerificationRequest) SetPhoneNumbers(v []TFPhoneNumber) {
	o.PhoneNumbers = v
}

// GetUseCase returns the UseCase field value
func (o *TFVerificationRequest) GetUseCase() UseCaseCategories {
	if o == nil {
		var ret UseCaseCategories
		return ret
	}

	return o.UseCase
}

// GetUseCaseOk returns a tuple with the UseCase field value
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetUseCaseOk() (*UseCaseCategories, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseCase, true
}

// SetUseCase sets field value
func (o *TFVerificationRequest) SetUseCase(v UseCaseCategories) {
	o.UseCase = v
}

// GetUseCaseSummary returns the UseCaseSummary field value
func (o *TFVerificationRequest) GetUseCaseSummary() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UseCaseSummary
}

// GetUseCaseSummaryOk returns a tuple with the UseCaseSummary field value
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetUseCaseSummaryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UseCaseSummary, true
}

// SetUseCaseSummary sets field value
func (o *TFVerificationRequest) SetUseCaseSummary(v string) {
	o.UseCaseSummary = v
}

// GetProductionMessageContent returns the ProductionMessageContent field value
func (o *TFVerificationRequest) GetProductionMessageContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProductionMessageContent
}

// GetProductionMessageContentOk returns a tuple with the ProductionMessageContent field value
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetProductionMessageContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProductionMessageContent, true
}

// SetProductionMessageContent sets field value
func (o *TFVerificationRequest) SetProductionMessageContent(v string) {
	o.ProductionMessageContent = v
}

// GetOptInWorkflow returns the OptInWorkflow field value
func (o *TFVerificationRequest) GetOptInWorkflow() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OptInWorkflow
}

// GetOptInWorkflowOk returns a tuple with the OptInWorkflow field value
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetOptInWorkflowOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OptInWorkflow, true
}

// SetOptInWorkflow sets field value
func (o *TFVerificationRequest) SetOptInWorkflow(v string) {
	o.OptInWorkflow = v
}

// GetOptInWorkflowImageURLs returns the OptInWorkflowImageURLs field value
func (o *TFVerificationRequest) GetOptInWorkflowImageURLs() []Url {
	if o == nil {
		var ret []Url
		return ret
	}

	return o.OptInWorkflowImageURLs
}

// GetOptInWorkflowImageURLsOk returns a tuple with the OptInWorkflowImageURLs field value
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetOptInWorkflowImageURLsOk() ([]Url, bool) {
	if o == nil {
		return nil, false
	}
	return o.OptInWorkflowImageURLs, true
}

// SetOptInWorkflowImageURLs sets field value
func (o *TFVerificationRequest) SetOptInWorkflowImageURLs(v []Url) {
	o.OptInWorkflowImageURLs = v
}

// GetAdditionalInformation returns the AdditionalInformation field value
func (o *TFVerificationRequest) GetAdditionalInformation() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AdditionalInformation
}

// GetAdditionalInformationOk returns a tuple with the AdditionalInformation field value
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetAdditionalInformationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AdditionalInformation, true
}

// SetAdditionalInformation sets field value
func (o *TFVerificationRequest) SetAdditionalInformation(v string) {
	o.AdditionalInformation = v
}

// GetIsvReseller returns the IsvReseller field value
func (o *TFVerificationRequest) GetIsvReseller() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.IsvReseller
}

// GetIsvResellerOk returns a tuple with the IsvReseller field value
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetIsvResellerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsvReseller, true
}

// SetIsvReseller sets field value
func (o *TFVerificationRequest) SetIsvReseller(v string) {
	o.IsvReseller = v
}

// GetWebhookUrl returns the WebhookUrl field value if set, zero value otherwise.
func (o *TFVerificationRequest) GetWebhookUrl() string {
	if o == nil || IsNil(o.WebhookUrl) {
		var ret string
		return ret
	}
	return *o.WebhookUrl
}

// GetWebhookUrlOk returns a tuple with the WebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TFVerificationRequest) GetWebhookUrlOk() (*string, bool) {
	if o == nil || IsNil(o.WebhookUrl) {
		return nil, false
	}
	return o.WebhookUrl, true
}

// HasWebhookUrl returns a boolean if a field has been set.
func (o *TFVerificationRequest) HasWebhookUrl() bool {
	if o != nil && !IsNil(o.WebhookUrl) {
		return true
	}

	return false
}

// SetWebhookUrl gets a reference to the given string and assigns it to the WebhookUrl field.
func (o *TFVerificationRequest) SetWebhookUrl(v string) {
	o.WebhookUrl = &v
}

func (o TFVerificationRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TFVerificationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["businessName"] = o.BusinessName
	toSerialize["corporateWebsite"] = o.CorporateWebsite
	toSerialize["businessAddr1"] = o.BusinessAddr1
	if !IsNil(o.BusinessAddr2) {
		toSerialize["businessAddr2"] = o.BusinessAddr2
	}
	toSerialize["businessCity"] = o.BusinessCity
	toSerialize["businessState"] = o.BusinessState
	toSerialize["businessZip"] = o.BusinessZip
	toSerialize["businessContactFirstName"] = o.BusinessContactFirstName
	toSerialize["businessContactLastName"] = o.BusinessContactLastName
	toSerialize["businessContactEmail"] = o.BusinessContactEmail
	toSerialize["businessContactPhone"] = o.BusinessContactPhone
	toSerialize["messageVolume"] = o.MessageVolume
	toSerialize["phoneNumbers"] = o.PhoneNumbers
	toSerialize["useCase"] = o.UseCase
	toSerialize["useCaseSummary"] = o.UseCaseSummary
	toSerialize["productionMessageContent"] = o.ProductionMessageContent
	toSerialize["optInWorkflow"] = o.OptInWorkflow
	toSerialize["optInWorkflowImageURLs"] = o.OptInWorkflowImageURLs
	toSerialize["additionalInformation"] = o.AdditionalInformation
	toSerialize["isvReseller"] = o.IsvReseller
	if !IsNil(o.WebhookUrl) {
		toSerialize["webhookUrl"] = o.WebhookUrl
	}
	return toSerialize, nil
}

func (o *TFVerificationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"businessName",
		"corporateWebsite",
		"businessAddr1",
		"businessCity",
		"businessState",
		"businessZip",
		"businessContactFirstName",
		"businessContactLastName",
		"businessContactEmail",
		"businessContactPhone",
		"messageVolume",
		"phoneNumbers",
		"useCase",
		"useCaseSummary",
		"productionMessageContent",
		"optInWorkflow",
		"optInWorkflowImageURLs",
		"additionalInformation",
		"isvReseller",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varTFVerificationRequest := _TFVerificationRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varTFVerificationRequest)

	if err != nil {
		return err
	}

	*o = TFVerificationRequest(varTFVerificationRequest)

	return err
}

type NullableTFVerificationRequest struct {
	value *TFVerificationRequest
	isSet bool
}

func (v NullableTFVerificationRequest) Get() *TFVerificationRequest {
	return v.value
}

func (v *NullableTFVerificationRequest) Set(val *TFVerificationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableTFVerificationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableTFVerificationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTFVerificationRequest(val *TFVerificationRequest) *NullableTFVerificationRequest {
	return &NullableTFVerificationRequest{value: val, isSet: true}
}

func (v NullableTFVerificationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTFVerificationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


