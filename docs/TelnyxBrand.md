# TelnyxBrand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | [**EntityType**](EntityType.md) | Entity type behind the brand. This is the form of business establishment. | 
**CspId** | Pointer to **string** | Unique identifier assigned to the csp by the registry. | [optional] 
**BrandId** | Pointer to **string** | Unique identifier assigned to the brand. | [optional] 
**TcrBrandId** | Pointer to **string** | Unique identifier assigned to the brand by the registry. | [optional] 
**DisplayName** | **string** | Display or marketing name of the brand. | 
**CompanyName** | Pointer to **string** | (Required for Non-profit/private/public) Legal company name. | [optional] 
**FirstName** | Pointer to **string** | First name of business contact. | [optional] 
**LastName** | Pointer to **string** | Last name of business contact. | [optional] 
**Ein** | Pointer to **string** | (Required for Non-profit) Government assigned corporate tax ID. EIN is 9-digits in U.S. | [optional] 
**Phone** | Pointer to **string** | Valid phone number in e.164 international format. | [optional] 
**Street** | Pointer to **string** | Street number and name. | [optional] 
**City** | Pointer to **string** | City name | [optional] 
**State** | Pointer to **string** | State. Must be 2 letters code for United States. | [optional] 
**PostalCode** | Pointer to **string** | Postal codes. Use 5 digit zipcode for United States | [optional] 
**Country** | **string** | ISO2 2 characters country code. Example: US - United States | 
**Email** | **string** | Valid email address of brand support contact. | 
**StockSymbol** | Pointer to **string** | (Required for public company) stock symbol. | [optional] 
**StockExchange** | Pointer to [**StockExchange**](StockExchange.md) | (Required for public company) stock exchange. | [optional] 
**IpAddress** | Pointer to **string** | IP address of the browser requesting to create brand identity. | [optional] 
**Website** | Pointer to **string** | Brand website URL. | [optional] 
**BrandRelationship** | [**BrandRelationship**](BrandRelationship.md) | Brand relationship to the CSP | 
**Vertical** | **string** | Vertical or industry segment of the brand. | 
**AltBusinessId** | Pointer to **string** | Alternate business identifier such as DUNS, LEI, or GIIN | [optional] 
**AltBusinessIdType** | Pointer to [**AltBusinessIdType**](AltBusinessIdType.md) |  | [optional] 
**UniversalEin** | Pointer to **string** | Universal EIN of Brand, Read Only. | [optional] 
**ReferenceId** | Pointer to **string** | Unique identifier Telnyx assigned to the brand - the brandId | [optional] 
**IdentityStatus** | Pointer to [**BrandIdentityStatus**](BrandIdentityStatus.md) |  | [optional] 
**OptionalAttributes** | Pointer to [**BrandOptionalAttributes**](BrandOptionalAttributes.md) |  | [optional] 
**Mock** | Pointer to **bool** | Mock brand for testing purposes | [optional] [default to false]
**MobilePhone** | Pointer to **string** | Valid mobile phone number in e.164 international format. | [optional] 
**IsReseller** | Pointer to **bool** | Indicates whether this brand is known to be a reseller | [optional] [default to false]
**WebhookURL** | Pointer to **string** | Webhook to which brand status updates are sent. | [optional] 
**BusinessContactEmail** | Pointer to **string** | Business contact email.  Required if &#x60;entityType&#x60; is &#x60;PUBLIC_PROFIT&#x60;. | [optional] 
**WebhookFailoverURL** | Pointer to **string** | Failover webhook to which brand status updates are sent. | [optional] 
**CreatedAt** | Pointer to **string** | Date and time that the brand was created at. | [optional] 
**UpdatedAt** | Pointer to **string** | Date and time that the brand was last updated at. | [optional] 
**Status** | Pointer to **string** | Status of the brand | [optional] 
**FailureReasons** | Pointer to **interface{}** | Failure reasons for brand | [optional] 

## Methods

### NewTelnyxBrand

`func NewTelnyxBrand(entityType EntityType, displayName string, country string, email string, brandRelationship BrandRelationship, vertical string, ) *TelnyxBrand`

NewTelnyxBrand instantiates a new TelnyxBrand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelnyxBrandWithDefaults

`func NewTelnyxBrandWithDefaults() *TelnyxBrand`

NewTelnyxBrandWithDefaults instantiates a new TelnyxBrand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *TelnyxBrand) GetEntityType() EntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *TelnyxBrand) GetEntityTypeOk() (*EntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *TelnyxBrand) SetEntityType(v EntityType)`

SetEntityType sets EntityType field to given value.


### GetCspId

`func (o *TelnyxBrand) GetCspId() string`

GetCspId returns the CspId field if non-nil, zero value otherwise.

### GetCspIdOk

`func (o *TelnyxBrand) GetCspIdOk() (*string, bool)`

GetCspIdOk returns a tuple with the CspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspId

`func (o *TelnyxBrand) SetCspId(v string)`

SetCspId sets CspId field to given value.

### HasCspId

`func (o *TelnyxBrand) HasCspId() bool`

HasCspId returns a boolean if a field has been set.

### GetBrandId

`func (o *TelnyxBrand) GetBrandId() string`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *TelnyxBrand) GetBrandIdOk() (*string, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *TelnyxBrand) SetBrandId(v string)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *TelnyxBrand) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### GetTcrBrandId

`func (o *TelnyxBrand) GetTcrBrandId() string`

GetTcrBrandId returns the TcrBrandId field if non-nil, zero value otherwise.

### GetTcrBrandIdOk

`func (o *TelnyxBrand) GetTcrBrandIdOk() (*string, bool)`

GetTcrBrandIdOk returns a tuple with the TcrBrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcrBrandId

`func (o *TelnyxBrand) SetTcrBrandId(v string)`

SetTcrBrandId sets TcrBrandId field to given value.

### HasTcrBrandId

`func (o *TelnyxBrand) HasTcrBrandId() bool`

HasTcrBrandId returns a boolean if a field has been set.

### GetDisplayName

`func (o *TelnyxBrand) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TelnyxBrand) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TelnyxBrand) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetCompanyName

`func (o *TelnyxBrand) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *TelnyxBrand) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *TelnyxBrand) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *TelnyxBrand) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetFirstName

`func (o *TelnyxBrand) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *TelnyxBrand) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *TelnyxBrand) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *TelnyxBrand) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *TelnyxBrand) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *TelnyxBrand) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *TelnyxBrand) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *TelnyxBrand) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEin

`func (o *TelnyxBrand) GetEin() string`

GetEin returns the Ein field if non-nil, zero value otherwise.

### GetEinOk

`func (o *TelnyxBrand) GetEinOk() (*string, bool)`

GetEinOk returns a tuple with the Ein field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEin

`func (o *TelnyxBrand) SetEin(v string)`

SetEin sets Ein field to given value.

### HasEin

`func (o *TelnyxBrand) HasEin() bool`

HasEin returns a boolean if a field has been set.

### GetPhone

`func (o *TelnyxBrand) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *TelnyxBrand) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *TelnyxBrand) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *TelnyxBrand) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetStreet

`func (o *TelnyxBrand) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *TelnyxBrand) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *TelnyxBrand) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *TelnyxBrand) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetCity

`func (o *TelnyxBrand) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *TelnyxBrand) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *TelnyxBrand) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *TelnyxBrand) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *TelnyxBrand) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TelnyxBrand) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TelnyxBrand) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TelnyxBrand) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *TelnyxBrand) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *TelnyxBrand) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *TelnyxBrand) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *TelnyxBrand) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *TelnyxBrand) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TelnyxBrand) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TelnyxBrand) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetEmail

`func (o *TelnyxBrand) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TelnyxBrand) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TelnyxBrand) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetStockSymbol

`func (o *TelnyxBrand) GetStockSymbol() string`

GetStockSymbol returns the StockSymbol field if non-nil, zero value otherwise.

### GetStockSymbolOk

`func (o *TelnyxBrand) GetStockSymbolOk() (*string, bool)`

GetStockSymbolOk returns a tuple with the StockSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockSymbol

`func (o *TelnyxBrand) SetStockSymbol(v string)`

SetStockSymbol sets StockSymbol field to given value.

### HasStockSymbol

`func (o *TelnyxBrand) HasStockSymbol() bool`

HasStockSymbol returns a boolean if a field has been set.

### GetStockExchange

`func (o *TelnyxBrand) GetStockExchange() StockExchange`

GetStockExchange returns the StockExchange field if non-nil, zero value otherwise.

### GetStockExchangeOk

`func (o *TelnyxBrand) GetStockExchangeOk() (*StockExchange, bool)`

GetStockExchangeOk returns a tuple with the StockExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockExchange

`func (o *TelnyxBrand) SetStockExchange(v StockExchange)`

SetStockExchange sets StockExchange field to given value.

### HasStockExchange

`func (o *TelnyxBrand) HasStockExchange() bool`

HasStockExchange returns a boolean if a field has been set.

### GetIpAddress

`func (o *TelnyxBrand) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *TelnyxBrand) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *TelnyxBrand) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *TelnyxBrand) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetWebsite

`func (o *TelnyxBrand) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *TelnyxBrand) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *TelnyxBrand) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *TelnyxBrand) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetBrandRelationship

`func (o *TelnyxBrand) GetBrandRelationship() BrandRelationship`

GetBrandRelationship returns the BrandRelationship field if non-nil, zero value otherwise.

### GetBrandRelationshipOk

`func (o *TelnyxBrand) GetBrandRelationshipOk() (*BrandRelationship, bool)`

GetBrandRelationshipOk returns a tuple with the BrandRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandRelationship

`func (o *TelnyxBrand) SetBrandRelationship(v BrandRelationship)`

SetBrandRelationship sets BrandRelationship field to given value.


### GetVertical

`func (o *TelnyxBrand) GetVertical() string`

GetVertical returns the Vertical field if non-nil, zero value otherwise.

### GetVerticalOk

`func (o *TelnyxBrand) GetVerticalOk() (*string, bool)`

GetVerticalOk returns a tuple with the Vertical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertical

`func (o *TelnyxBrand) SetVertical(v string)`

SetVertical sets Vertical field to given value.


### GetAltBusinessId

`func (o *TelnyxBrand) GetAltBusinessId() string`

GetAltBusinessId returns the AltBusinessId field if non-nil, zero value otherwise.

### GetAltBusinessIdOk

`func (o *TelnyxBrand) GetAltBusinessIdOk() (*string, bool)`

GetAltBusinessIdOk returns a tuple with the AltBusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltBusinessId

`func (o *TelnyxBrand) SetAltBusinessId(v string)`

SetAltBusinessId sets AltBusinessId field to given value.

### HasAltBusinessId

`func (o *TelnyxBrand) HasAltBusinessId() bool`

HasAltBusinessId returns a boolean if a field has been set.

### GetAltBusinessIdType

`func (o *TelnyxBrand) GetAltBusinessIdType() AltBusinessIdType`

GetAltBusinessIdType returns the AltBusinessIdType field if non-nil, zero value otherwise.

### GetAltBusinessIdTypeOk

`func (o *TelnyxBrand) GetAltBusinessIdTypeOk() (*AltBusinessIdType, bool)`

GetAltBusinessIdTypeOk returns a tuple with the AltBusinessIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltBusinessIdType

`func (o *TelnyxBrand) SetAltBusinessIdType(v AltBusinessIdType)`

SetAltBusinessIdType sets AltBusinessIdType field to given value.

### HasAltBusinessIdType

`func (o *TelnyxBrand) HasAltBusinessIdType() bool`

HasAltBusinessIdType returns a boolean if a field has been set.

### GetUniversalEin

`func (o *TelnyxBrand) GetUniversalEin() string`

GetUniversalEin returns the UniversalEin field if non-nil, zero value otherwise.

### GetUniversalEinOk

`func (o *TelnyxBrand) GetUniversalEinOk() (*string, bool)`

GetUniversalEinOk returns a tuple with the UniversalEin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalEin

`func (o *TelnyxBrand) SetUniversalEin(v string)`

SetUniversalEin sets UniversalEin field to given value.

### HasUniversalEin

`func (o *TelnyxBrand) HasUniversalEin() bool`

HasUniversalEin returns a boolean if a field has been set.

### GetReferenceId

`func (o *TelnyxBrand) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *TelnyxBrand) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *TelnyxBrand) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *TelnyxBrand) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetIdentityStatus

`func (o *TelnyxBrand) GetIdentityStatus() BrandIdentityStatus`

GetIdentityStatus returns the IdentityStatus field if non-nil, zero value otherwise.

### GetIdentityStatusOk

`func (o *TelnyxBrand) GetIdentityStatusOk() (*BrandIdentityStatus, bool)`

GetIdentityStatusOk returns a tuple with the IdentityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityStatus

`func (o *TelnyxBrand) SetIdentityStatus(v BrandIdentityStatus)`

SetIdentityStatus sets IdentityStatus field to given value.

### HasIdentityStatus

`func (o *TelnyxBrand) HasIdentityStatus() bool`

HasIdentityStatus returns a boolean if a field has been set.

### GetOptionalAttributes

`func (o *TelnyxBrand) GetOptionalAttributes() BrandOptionalAttributes`

GetOptionalAttributes returns the OptionalAttributes field if non-nil, zero value otherwise.

### GetOptionalAttributesOk

`func (o *TelnyxBrand) GetOptionalAttributesOk() (*BrandOptionalAttributes, bool)`

GetOptionalAttributesOk returns a tuple with the OptionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalAttributes

`func (o *TelnyxBrand) SetOptionalAttributes(v BrandOptionalAttributes)`

SetOptionalAttributes sets OptionalAttributes field to given value.

### HasOptionalAttributes

`func (o *TelnyxBrand) HasOptionalAttributes() bool`

HasOptionalAttributes returns a boolean if a field has been set.

### GetMock

`func (o *TelnyxBrand) GetMock() bool`

GetMock returns the Mock field if non-nil, zero value otherwise.

### GetMockOk

`func (o *TelnyxBrand) GetMockOk() (*bool, bool)`

GetMockOk returns a tuple with the Mock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMock

`func (o *TelnyxBrand) SetMock(v bool)`

SetMock sets Mock field to given value.

### HasMock

`func (o *TelnyxBrand) HasMock() bool`

HasMock returns a boolean if a field has been set.

### GetMobilePhone

`func (o *TelnyxBrand) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *TelnyxBrand) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *TelnyxBrand) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *TelnyxBrand) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetIsReseller

`func (o *TelnyxBrand) GetIsReseller() bool`

GetIsReseller returns the IsReseller field if non-nil, zero value otherwise.

### GetIsResellerOk

`func (o *TelnyxBrand) GetIsResellerOk() (*bool, bool)`

GetIsResellerOk returns a tuple with the IsReseller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReseller

`func (o *TelnyxBrand) SetIsReseller(v bool)`

SetIsReseller sets IsReseller field to given value.

### HasIsReseller

`func (o *TelnyxBrand) HasIsReseller() bool`

HasIsReseller returns a boolean if a field has been set.

### GetWebhookURL

`func (o *TelnyxBrand) GetWebhookURL() string`

GetWebhookURL returns the WebhookURL field if non-nil, zero value otherwise.

### GetWebhookURLOk

`func (o *TelnyxBrand) GetWebhookURLOk() (*string, bool)`

GetWebhookURLOk returns a tuple with the WebhookURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookURL

`func (o *TelnyxBrand) SetWebhookURL(v string)`

SetWebhookURL sets WebhookURL field to given value.

### HasWebhookURL

`func (o *TelnyxBrand) HasWebhookURL() bool`

HasWebhookURL returns a boolean if a field has been set.

### GetBusinessContactEmail

`func (o *TelnyxBrand) GetBusinessContactEmail() string`

GetBusinessContactEmail returns the BusinessContactEmail field if non-nil, zero value otherwise.

### GetBusinessContactEmailOk

`func (o *TelnyxBrand) GetBusinessContactEmailOk() (*string, bool)`

GetBusinessContactEmailOk returns a tuple with the BusinessContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessContactEmail

`func (o *TelnyxBrand) SetBusinessContactEmail(v string)`

SetBusinessContactEmail sets BusinessContactEmail field to given value.

### HasBusinessContactEmail

`func (o *TelnyxBrand) HasBusinessContactEmail() bool`

HasBusinessContactEmail returns a boolean if a field has been set.

### GetWebhookFailoverURL

`func (o *TelnyxBrand) GetWebhookFailoverURL() string`

GetWebhookFailoverURL returns the WebhookFailoverURL field if non-nil, zero value otherwise.

### GetWebhookFailoverURLOk

`func (o *TelnyxBrand) GetWebhookFailoverURLOk() (*string, bool)`

GetWebhookFailoverURLOk returns a tuple with the WebhookFailoverURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailoverURL

`func (o *TelnyxBrand) SetWebhookFailoverURL(v string)`

SetWebhookFailoverURL sets WebhookFailoverURL field to given value.

### HasWebhookFailoverURL

`func (o *TelnyxBrand) HasWebhookFailoverURL() bool`

HasWebhookFailoverURL returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TelnyxBrand) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TelnyxBrand) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TelnyxBrand) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TelnyxBrand) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TelnyxBrand) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TelnyxBrand) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TelnyxBrand) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TelnyxBrand) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *TelnyxBrand) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TelnyxBrand) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TelnyxBrand) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TelnyxBrand) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFailureReasons

`func (o *TelnyxBrand) GetFailureReasons() interface{}`

GetFailureReasons returns the FailureReasons field if non-nil, zero value otherwise.

### GetFailureReasonsOk

`func (o *TelnyxBrand) GetFailureReasonsOk() (*interface{}, bool)`

GetFailureReasonsOk returns a tuple with the FailureReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReasons

`func (o *TelnyxBrand) SetFailureReasons(v interface{})`

SetFailureReasons sets FailureReasons field to given value.

### HasFailureReasons

`func (o *TelnyxBrand) HasFailureReasons() bool`

HasFailureReasons returns a boolean if a field has been set.

### SetFailureReasonsNil

`func (o *TelnyxBrand) SetFailureReasonsNil(b bool)`

 SetFailureReasonsNil sets the value for FailureReasons to be an explicit nil

### UnsetFailureReasons
`func (o *TelnyxBrand) UnsetFailureReasons()`

UnsetFailureReasons ensures that no value is present for FailureReasons, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


