# TelnyxBrandWithCampaignsCount

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
**AssignedCampaignsCount** | Pointer to **float32** | Number of campaigns associated with the brand | [optional] 

## Methods

### NewTelnyxBrandWithCampaignsCount

`func NewTelnyxBrandWithCampaignsCount(entityType EntityType, displayName string, country string, email string, brandRelationship BrandRelationship, vertical string, ) *TelnyxBrandWithCampaignsCount`

NewTelnyxBrandWithCampaignsCount instantiates a new TelnyxBrandWithCampaignsCount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTelnyxBrandWithCampaignsCountWithDefaults

`func NewTelnyxBrandWithCampaignsCountWithDefaults() *TelnyxBrandWithCampaignsCount`

NewTelnyxBrandWithCampaignsCountWithDefaults instantiates a new TelnyxBrandWithCampaignsCount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *TelnyxBrandWithCampaignsCount) GetEntityType() EntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *TelnyxBrandWithCampaignsCount) GetEntityTypeOk() (*EntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *TelnyxBrandWithCampaignsCount) SetEntityType(v EntityType)`

SetEntityType sets EntityType field to given value.


### GetCspId

`func (o *TelnyxBrandWithCampaignsCount) GetCspId() string`

GetCspId returns the CspId field if non-nil, zero value otherwise.

### GetCspIdOk

`func (o *TelnyxBrandWithCampaignsCount) GetCspIdOk() (*string, bool)`

GetCspIdOk returns a tuple with the CspId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCspId

`func (o *TelnyxBrandWithCampaignsCount) SetCspId(v string)`

SetCspId sets CspId field to given value.

### HasCspId

`func (o *TelnyxBrandWithCampaignsCount) HasCspId() bool`

HasCspId returns a boolean if a field has been set.

### GetBrandId

`func (o *TelnyxBrandWithCampaignsCount) GetBrandId() string`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *TelnyxBrandWithCampaignsCount) GetBrandIdOk() (*string, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *TelnyxBrandWithCampaignsCount) SetBrandId(v string)`

SetBrandId sets BrandId field to given value.

### HasBrandId

`func (o *TelnyxBrandWithCampaignsCount) HasBrandId() bool`

HasBrandId returns a boolean if a field has been set.

### GetTcrBrandId

`func (o *TelnyxBrandWithCampaignsCount) GetTcrBrandId() string`

GetTcrBrandId returns the TcrBrandId field if non-nil, zero value otherwise.

### GetTcrBrandIdOk

`func (o *TelnyxBrandWithCampaignsCount) GetTcrBrandIdOk() (*string, bool)`

GetTcrBrandIdOk returns a tuple with the TcrBrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcrBrandId

`func (o *TelnyxBrandWithCampaignsCount) SetTcrBrandId(v string)`

SetTcrBrandId sets TcrBrandId field to given value.

### HasTcrBrandId

`func (o *TelnyxBrandWithCampaignsCount) HasTcrBrandId() bool`

HasTcrBrandId returns a boolean if a field has been set.

### GetDisplayName

`func (o *TelnyxBrandWithCampaignsCount) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *TelnyxBrandWithCampaignsCount) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *TelnyxBrandWithCampaignsCount) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetCompanyName

`func (o *TelnyxBrandWithCampaignsCount) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *TelnyxBrandWithCampaignsCount) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *TelnyxBrandWithCampaignsCount) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *TelnyxBrandWithCampaignsCount) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetFirstName

`func (o *TelnyxBrandWithCampaignsCount) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *TelnyxBrandWithCampaignsCount) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *TelnyxBrandWithCampaignsCount) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *TelnyxBrandWithCampaignsCount) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *TelnyxBrandWithCampaignsCount) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *TelnyxBrandWithCampaignsCount) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *TelnyxBrandWithCampaignsCount) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *TelnyxBrandWithCampaignsCount) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEin

`func (o *TelnyxBrandWithCampaignsCount) GetEin() string`

GetEin returns the Ein field if non-nil, zero value otherwise.

### GetEinOk

`func (o *TelnyxBrandWithCampaignsCount) GetEinOk() (*string, bool)`

GetEinOk returns a tuple with the Ein field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEin

`func (o *TelnyxBrandWithCampaignsCount) SetEin(v string)`

SetEin sets Ein field to given value.

### HasEin

`func (o *TelnyxBrandWithCampaignsCount) HasEin() bool`

HasEin returns a boolean if a field has been set.

### GetPhone

`func (o *TelnyxBrandWithCampaignsCount) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *TelnyxBrandWithCampaignsCount) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *TelnyxBrandWithCampaignsCount) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *TelnyxBrandWithCampaignsCount) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetStreet

`func (o *TelnyxBrandWithCampaignsCount) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *TelnyxBrandWithCampaignsCount) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *TelnyxBrandWithCampaignsCount) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *TelnyxBrandWithCampaignsCount) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetCity

`func (o *TelnyxBrandWithCampaignsCount) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *TelnyxBrandWithCampaignsCount) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *TelnyxBrandWithCampaignsCount) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *TelnyxBrandWithCampaignsCount) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *TelnyxBrandWithCampaignsCount) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *TelnyxBrandWithCampaignsCount) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *TelnyxBrandWithCampaignsCount) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *TelnyxBrandWithCampaignsCount) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *TelnyxBrandWithCampaignsCount) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *TelnyxBrandWithCampaignsCount) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *TelnyxBrandWithCampaignsCount) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *TelnyxBrandWithCampaignsCount) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *TelnyxBrandWithCampaignsCount) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TelnyxBrandWithCampaignsCount) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TelnyxBrandWithCampaignsCount) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetEmail

`func (o *TelnyxBrandWithCampaignsCount) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *TelnyxBrandWithCampaignsCount) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *TelnyxBrandWithCampaignsCount) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetStockSymbol

`func (o *TelnyxBrandWithCampaignsCount) GetStockSymbol() string`

GetStockSymbol returns the StockSymbol field if non-nil, zero value otherwise.

### GetStockSymbolOk

`func (o *TelnyxBrandWithCampaignsCount) GetStockSymbolOk() (*string, bool)`

GetStockSymbolOk returns a tuple with the StockSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockSymbol

`func (o *TelnyxBrandWithCampaignsCount) SetStockSymbol(v string)`

SetStockSymbol sets StockSymbol field to given value.

### HasStockSymbol

`func (o *TelnyxBrandWithCampaignsCount) HasStockSymbol() bool`

HasStockSymbol returns a boolean if a field has been set.

### GetStockExchange

`func (o *TelnyxBrandWithCampaignsCount) GetStockExchange() StockExchange`

GetStockExchange returns the StockExchange field if non-nil, zero value otherwise.

### GetStockExchangeOk

`func (o *TelnyxBrandWithCampaignsCount) GetStockExchangeOk() (*StockExchange, bool)`

GetStockExchangeOk returns a tuple with the StockExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockExchange

`func (o *TelnyxBrandWithCampaignsCount) SetStockExchange(v StockExchange)`

SetStockExchange sets StockExchange field to given value.

### HasStockExchange

`func (o *TelnyxBrandWithCampaignsCount) HasStockExchange() bool`

HasStockExchange returns a boolean if a field has been set.

### GetIpAddress

`func (o *TelnyxBrandWithCampaignsCount) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *TelnyxBrandWithCampaignsCount) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *TelnyxBrandWithCampaignsCount) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *TelnyxBrandWithCampaignsCount) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetWebsite

`func (o *TelnyxBrandWithCampaignsCount) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *TelnyxBrandWithCampaignsCount) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *TelnyxBrandWithCampaignsCount) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *TelnyxBrandWithCampaignsCount) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetBrandRelationship

`func (o *TelnyxBrandWithCampaignsCount) GetBrandRelationship() BrandRelationship`

GetBrandRelationship returns the BrandRelationship field if non-nil, zero value otherwise.

### GetBrandRelationshipOk

`func (o *TelnyxBrandWithCampaignsCount) GetBrandRelationshipOk() (*BrandRelationship, bool)`

GetBrandRelationshipOk returns a tuple with the BrandRelationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandRelationship

`func (o *TelnyxBrandWithCampaignsCount) SetBrandRelationship(v BrandRelationship)`

SetBrandRelationship sets BrandRelationship field to given value.


### GetVertical

`func (o *TelnyxBrandWithCampaignsCount) GetVertical() string`

GetVertical returns the Vertical field if non-nil, zero value otherwise.

### GetVerticalOk

`func (o *TelnyxBrandWithCampaignsCount) GetVerticalOk() (*string, bool)`

GetVerticalOk returns a tuple with the Vertical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertical

`func (o *TelnyxBrandWithCampaignsCount) SetVertical(v string)`

SetVertical sets Vertical field to given value.


### GetAltBusinessId

`func (o *TelnyxBrandWithCampaignsCount) GetAltBusinessId() string`

GetAltBusinessId returns the AltBusinessId field if non-nil, zero value otherwise.

### GetAltBusinessIdOk

`func (o *TelnyxBrandWithCampaignsCount) GetAltBusinessIdOk() (*string, bool)`

GetAltBusinessIdOk returns a tuple with the AltBusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltBusinessId

`func (o *TelnyxBrandWithCampaignsCount) SetAltBusinessId(v string)`

SetAltBusinessId sets AltBusinessId field to given value.

### HasAltBusinessId

`func (o *TelnyxBrandWithCampaignsCount) HasAltBusinessId() bool`

HasAltBusinessId returns a boolean if a field has been set.

### GetAltBusinessIdType

`func (o *TelnyxBrandWithCampaignsCount) GetAltBusinessIdType() AltBusinessIdType`

GetAltBusinessIdType returns the AltBusinessIdType field if non-nil, zero value otherwise.

### GetAltBusinessIdTypeOk

`func (o *TelnyxBrandWithCampaignsCount) GetAltBusinessIdTypeOk() (*AltBusinessIdType, bool)`

GetAltBusinessIdTypeOk returns a tuple with the AltBusinessIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltBusinessIdType

`func (o *TelnyxBrandWithCampaignsCount) SetAltBusinessIdType(v AltBusinessIdType)`

SetAltBusinessIdType sets AltBusinessIdType field to given value.

### HasAltBusinessIdType

`func (o *TelnyxBrandWithCampaignsCount) HasAltBusinessIdType() bool`

HasAltBusinessIdType returns a boolean if a field has been set.

### GetUniversalEin

`func (o *TelnyxBrandWithCampaignsCount) GetUniversalEin() string`

GetUniversalEin returns the UniversalEin field if non-nil, zero value otherwise.

### GetUniversalEinOk

`func (o *TelnyxBrandWithCampaignsCount) GetUniversalEinOk() (*string, bool)`

GetUniversalEinOk returns a tuple with the UniversalEin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUniversalEin

`func (o *TelnyxBrandWithCampaignsCount) SetUniversalEin(v string)`

SetUniversalEin sets UniversalEin field to given value.

### HasUniversalEin

`func (o *TelnyxBrandWithCampaignsCount) HasUniversalEin() bool`

HasUniversalEin returns a boolean if a field has been set.

### GetReferenceId

`func (o *TelnyxBrandWithCampaignsCount) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *TelnyxBrandWithCampaignsCount) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *TelnyxBrandWithCampaignsCount) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *TelnyxBrandWithCampaignsCount) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetIdentityStatus

`func (o *TelnyxBrandWithCampaignsCount) GetIdentityStatus() BrandIdentityStatus`

GetIdentityStatus returns the IdentityStatus field if non-nil, zero value otherwise.

### GetIdentityStatusOk

`func (o *TelnyxBrandWithCampaignsCount) GetIdentityStatusOk() (*BrandIdentityStatus, bool)`

GetIdentityStatusOk returns a tuple with the IdentityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityStatus

`func (o *TelnyxBrandWithCampaignsCount) SetIdentityStatus(v BrandIdentityStatus)`

SetIdentityStatus sets IdentityStatus field to given value.

### HasIdentityStatus

`func (o *TelnyxBrandWithCampaignsCount) HasIdentityStatus() bool`

HasIdentityStatus returns a boolean if a field has been set.

### GetOptionalAttributes

`func (o *TelnyxBrandWithCampaignsCount) GetOptionalAttributes() BrandOptionalAttributes`

GetOptionalAttributes returns the OptionalAttributes field if non-nil, zero value otherwise.

### GetOptionalAttributesOk

`func (o *TelnyxBrandWithCampaignsCount) GetOptionalAttributesOk() (*BrandOptionalAttributes, bool)`

GetOptionalAttributesOk returns a tuple with the OptionalAttributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionalAttributes

`func (o *TelnyxBrandWithCampaignsCount) SetOptionalAttributes(v BrandOptionalAttributes)`

SetOptionalAttributes sets OptionalAttributes field to given value.

### HasOptionalAttributes

`func (o *TelnyxBrandWithCampaignsCount) HasOptionalAttributes() bool`

HasOptionalAttributes returns a boolean if a field has been set.

### GetMock

`func (o *TelnyxBrandWithCampaignsCount) GetMock() bool`

GetMock returns the Mock field if non-nil, zero value otherwise.

### GetMockOk

`func (o *TelnyxBrandWithCampaignsCount) GetMockOk() (*bool, bool)`

GetMockOk returns a tuple with the Mock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMock

`func (o *TelnyxBrandWithCampaignsCount) SetMock(v bool)`

SetMock sets Mock field to given value.

### HasMock

`func (o *TelnyxBrandWithCampaignsCount) HasMock() bool`

HasMock returns a boolean if a field has been set.

### GetMobilePhone

`func (o *TelnyxBrandWithCampaignsCount) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *TelnyxBrandWithCampaignsCount) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *TelnyxBrandWithCampaignsCount) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *TelnyxBrandWithCampaignsCount) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetIsReseller

`func (o *TelnyxBrandWithCampaignsCount) GetIsReseller() bool`

GetIsReseller returns the IsReseller field if non-nil, zero value otherwise.

### GetIsResellerOk

`func (o *TelnyxBrandWithCampaignsCount) GetIsResellerOk() (*bool, bool)`

GetIsResellerOk returns a tuple with the IsReseller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReseller

`func (o *TelnyxBrandWithCampaignsCount) SetIsReseller(v bool)`

SetIsReseller sets IsReseller field to given value.

### HasIsReseller

`func (o *TelnyxBrandWithCampaignsCount) HasIsReseller() bool`

HasIsReseller returns a boolean if a field has been set.

### GetWebhookURL

`func (o *TelnyxBrandWithCampaignsCount) GetWebhookURL() string`

GetWebhookURL returns the WebhookURL field if non-nil, zero value otherwise.

### GetWebhookURLOk

`func (o *TelnyxBrandWithCampaignsCount) GetWebhookURLOk() (*string, bool)`

GetWebhookURLOk returns a tuple with the WebhookURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookURL

`func (o *TelnyxBrandWithCampaignsCount) SetWebhookURL(v string)`

SetWebhookURL sets WebhookURL field to given value.

### HasWebhookURL

`func (o *TelnyxBrandWithCampaignsCount) HasWebhookURL() bool`

HasWebhookURL returns a boolean if a field has been set.

### GetBusinessContactEmail

`func (o *TelnyxBrandWithCampaignsCount) GetBusinessContactEmail() string`

GetBusinessContactEmail returns the BusinessContactEmail field if non-nil, zero value otherwise.

### GetBusinessContactEmailOk

`func (o *TelnyxBrandWithCampaignsCount) GetBusinessContactEmailOk() (*string, bool)`

GetBusinessContactEmailOk returns a tuple with the BusinessContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessContactEmail

`func (o *TelnyxBrandWithCampaignsCount) SetBusinessContactEmail(v string)`

SetBusinessContactEmail sets BusinessContactEmail field to given value.

### HasBusinessContactEmail

`func (o *TelnyxBrandWithCampaignsCount) HasBusinessContactEmail() bool`

HasBusinessContactEmail returns a boolean if a field has been set.

### GetWebhookFailoverURL

`func (o *TelnyxBrandWithCampaignsCount) GetWebhookFailoverURL() string`

GetWebhookFailoverURL returns the WebhookFailoverURL field if non-nil, zero value otherwise.

### GetWebhookFailoverURLOk

`func (o *TelnyxBrandWithCampaignsCount) GetWebhookFailoverURLOk() (*string, bool)`

GetWebhookFailoverURLOk returns a tuple with the WebhookFailoverURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailoverURL

`func (o *TelnyxBrandWithCampaignsCount) SetWebhookFailoverURL(v string)`

SetWebhookFailoverURL sets WebhookFailoverURL field to given value.

### HasWebhookFailoverURL

`func (o *TelnyxBrandWithCampaignsCount) HasWebhookFailoverURL() bool`

HasWebhookFailoverURL returns a boolean if a field has been set.

### GetCreatedAt

`func (o *TelnyxBrandWithCampaignsCount) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *TelnyxBrandWithCampaignsCount) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *TelnyxBrandWithCampaignsCount) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *TelnyxBrandWithCampaignsCount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *TelnyxBrandWithCampaignsCount) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *TelnyxBrandWithCampaignsCount) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *TelnyxBrandWithCampaignsCount) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *TelnyxBrandWithCampaignsCount) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *TelnyxBrandWithCampaignsCount) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TelnyxBrandWithCampaignsCount) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TelnyxBrandWithCampaignsCount) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *TelnyxBrandWithCampaignsCount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetFailureReasons

`func (o *TelnyxBrandWithCampaignsCount) GetFailureReasons() interface{}`

GetFailureReasons returns the FailureReasons field if non-nil, zero value otherwise.

### GetFailureReasonsOk

`func (o *TelnyxBrandWithCampaignsCount) GetFailureReasonsOk() (*interface{}, bool)`

GetFailureReasonsOk returns a tuple with the FailureReasons field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReasons

`func (o *TelnyxBrandWithCampaignsCount) SetFailureReasons(v interface{})`

SetFailureReasons sets FailureReasons field to given value.

### HasFailureReasons

`func (o *TelnyxBrandWithCampaignsCount) HasFailureReasons() bool`

HasFailureReasons returns a boolean if a field has been set.

### SetFailureReasonsNil

`func (o *TelnyxBrandWithCampaignsCount) SetFailureReasonsNil(b bool)`

 SetFailureReasonsNil sets the value for FailureReasons to be an explicit nil

### UnsetFailureReasons
`func (o *TelnyxBrandWithCampaignsCount) UnsetFailureReasons()`

UnsetFailureReasons ensures that no value is present for FailureReasons, not even an explicit nil
### GetAssignedCampaignsCount

`func (o *TelnyxBrandWithCampaignsCount) GetAssignedCampaignsCount() float32`

GetAssignedCampaignsCount returns the AssignedCampaignsCount field if non-nil, zero value otherwise.

### GetAssignedCampaignsCountOk

`func (o *TelnyxBrandWithCampaignsCount) GetAssignedCampaignsCountOk() (*float32, bool)`

GetAssignedCampaignsCountOk returns a tuple with the AssignedCampaignsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssignedCampaignsCount

`func (o *TelnyxBrandWithCampaignsCount) SetAssignedCampaignsCount(v float32)`

SetAssignedCampaignsCount sets AssignedCampaignsCount field to given value.

### HasAssignedCampaignsCount

`func (o *TelnyxBrandWithCampaignsCount) HasAssignedCampaignsCount() bool`

HasAssignedCampaignsCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


