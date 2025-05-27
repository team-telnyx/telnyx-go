# UpdateBrand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | [**EntityType**](EntityType.md) | Entity type behind the brand. This is the form of business establishment. | 
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
**Vertical** | [**Vertical**](Vertical.md) | Vertical or industry segment of the brand. | 
**AltBusinessId** | Pointer to **string** | Alternate business identifier such as DUNS, LEI, or GIIN | [optional] 
**AltBusinessIdType** | Pointer to [**AltBusinessIdType**](AltBusinessIdType.md) |  | [optional] 
**IsReseller** | Pointer to **bool** |  | [optional] 
**IdentityStatus** | Pointer to [**BrandIdentityStatus**](BrandIdentityStatus.md) |  | [optional] 
**BusinessContactEmail** | Pointer to **string** | Business contact email.  Required if &#x60;entityType&#x60; will be changed to &#x60;PUBLIC_PROFIT&#x60;. | [optional] 
**WebhookURL** | Pointer to **string** | Webhook URL for brand status updates. | [optional] 
**WebhookFailoverURL** | Pointer to **string** | Webhook failover URL for brand status updates. | [optional] 

## Methods

### NewUpdateBrand

`func NewUpdateBrand(entityType EntityType, displayName string, country string, email string, vertical Vertical, ) *UpdateBrand`

NewUpdateBrand instantiates a new UpdateBrand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBrandWithDefaults

`func NewUpdateBrandWithDefaults() *UpdateBrand`

NewUpdateBrandWithDefaults instantiates a new UpdateBrand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *UpdateBrand) GetEntityType() EntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *UpdateBrand) GetEntityTypeOk() (*EntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *UpdateBrand) SetEntityType(v EntityType)`

SetEntityType sets EntityType field to given value.


### GetDisplayName

`func (o *UpdateBrand) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UpdateBrand) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UpdateBrand) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetCompanyName

`func (o *UpdateBrand) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *UpdateBrand) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *UpdateBrand) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *UpdateBrand) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetFirstName

`func (o *UpdateBrand) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdateBrand) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdateBrand) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UpdateBrand) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UpdateBrand) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdateBrand) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdateBrand) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UpdateBrand) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEin

`func (o *UpdateBrand) GetEin() string`

GetEin returns the Ein field if non-nil, zero value otherwise.

### GetEinOk

`func (o *UpdateBrand) GetEinOk() (*string, bool)`

GetEinOk returns a tuple with the Ein field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEin

`func (o *UpdateBrand) SetEin(v string)`

SetEin sets Ein field to given value.

### HasEin

`func (o *UpdateBrand) HasEin() bool`

HasEin returns a boolean if a field has been set.

### GetPhone

`func (o *UpdateBrand) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UpdateBrand) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UpdateBrand) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UpdateBrand) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetStreet

`func (o *UpdateBrand) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *UpdateBrand) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *UpdateBrand) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *UpdateBrand) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetCity

`func (o *UpdateBrand) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *UpdateBrand) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *UpdateBrand) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *UpdateBrand) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *UpdateBrand) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateBrand) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateBrand) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UpdateBrand) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *UpdateBrand) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *UpdateBrand) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *UpdateBrand) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *UpdateBrand) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *UpdateBrand) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *UpdateBrand) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *UpdateBrand) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetEmail

`func (o *UpdateBrand) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateBrand) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateBrand) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetStockSymbol

`func (o *UpdateBrand) GetStockSymbol() string`

GetStockSymbol returns the StockSymbol field if non-nil, zero value otherwise.

### GetStockSymbolOk

`func (o *UpdateBrand) GetStockSymbolOk() (*string, bool)`

GetStockSymbolOk returns a tuple with the StockSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockSymbol

`func (o *UpdateBrand) SetStockSymbol(v string)`

SetStockSymbol sets StockSymbol field to given value.

### HasStockSymbol

`func (o *UpdateBrand) HasStockSymbol() bool`

HasStockSymbol returns a boolean if a field has been set.

### GetStockExchange

`func (o *UpdateBrand) GetStockExchange() StockExchange`

GetStockExchange returns the StockExchange field if non-nil, zero value otherwise.

### GetStockExchangeOk

`func (o *UpdateBrand) GetStockExchangeOk() (*StockExchange, bool)`

GetStockExchangeOk returns a tuple with the StockExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockExchange

`func (o *UpdateBrand) SetStockExchange(v StockExchange)`

SetStockExchange sets StockExchange field to given value.

### HasStockExchange

`func (o *UpdateBrand) HasStockExchange() bool`

HasStockExchange returns a boolean if a field has been set.

### GetIpAddress

`func (o *UpdateBrand) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *UpdateBrand) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *UpdateBrand) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *UpdateBrand) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetWebsite

`func (o *UpdateBrand) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *UpdateBrand) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *UpdateBrand) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *UpdateBrand) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetVertical

`func (o *UpdateBrand) GetVertical() Vertical`

GetVertical returns the Vertical field if non-nil, zero value otherwise.

### GetVerticalOk

`func (o *UpdateBrand) GetVerticalOk() (*Vertical, bool)`

GetVerticalOk returns a tuple with the Vertical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertical

`func (o *UpdateBrand) SetVertical(v Vertical)`

SetVertical sets Vertical field to given value.


### GetAltBusinessId

`func (o *UpdateBrand) GetAltBusinessId() string`

GetAltBusinessId returns the AltBusinessId field if non-nil, zero value otherwise.

### GetAltBusinessIdOk

`func (o *UpdateBrand) GetAltBusinessIdOk() (*string, bool)`

GetAltBusinessIdOk returns a tuple with the AltBusinessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltBusinessId

`func (o *UpdateBrand) SetAltBusinessId(v string)`

SetAltBusinessId sets AltBusinessId field to given value.

### HasAltBusinessId

`func (o *UpdateBrand) HasAltBusinessId() bool`

HasAltBusinessId returns a boolean if a field has been set.

### GetAltBusinessIdType

`func (o *UpdateBrand) GetAltBusinessIdType() AltBusinessIdType`

GetAltBusinessIdType returns the AltBusinessIdType field if non-nil, zero value otherwise.

### GetAltBusinessIdTypeOk

`func (o *UpdateBrand) GetAltBusinessIdTypeOk() (*AltBusinessIdType, bool)`

GetAltBusinessIdTypeOk returns a tuple with the AltBusinessIdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltBusinessIdType

`func (o *UpdateBrand) SetAltBusinessIdType(v AltBusinessIdType)`

SetAltBusinessIdType sets AltBusinessIdType field to given value.

### HasAltBusinessIdType

`func (o *UpdateBrand) HasAltBusinessIdType() bool`

HasAltBusinessIdType returns a boolean if a field has been set.

### GetIsReseller

`func (o *UpdateBrand) GetIsReseller() bool`

GetIsReseller returns the IsReseller field if non-nil, zero value otherwise.

### GetIsResellerOk

`func (o *UpdateBrand) GetIsResellerOk() (*bool, bool)`

GetIsResellerOk returns a tuple with the IsReseller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReseller

`func (o *UpdateBrand) SetIsReseller(v bool)`

SetIsReseller sets IsReseller field to given value.

### HasIsReseller

`func (o *UpdateBrand) HasIsReseller() bool`

HasIsReseller returns a boolean if a field has been set.

### GetIdentityStatus

`func (o *UpdateBrand) GetIdentityStatus() BrandIdentityStatus`

GetIdentityStatus returns the IdentityStatus field if non-nil, zero value otherwise.

### GetIdentityStatusOk

`func (o *UpdateBrand) GetIdentityStatusOk() (*BrandIdentityStatus, bool)`

GetIdentityStatusOk returns a tuple with the IdentityStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityStatus

`func (o *UpdateBrand) SetIdentityStatus(v BrandIdentityStatus)`

SetIdentityStatus sets IdentityStatus field to given value.

### HasIdentityStatus

`func (o *UpdateBrand) HasIdentityStatus() bool`

HasIdentityStatus returns a boolean if a field has been set.

### GetBusinessContactEmail

`func (o *UpdateBrand) GetBusinessContactEmail() string`

GetBusinessContactEmail returns the BusinessContactEmail field if non-nil, zero value otherwise.

### GetBusinessContactEmailOk

`func (o *UpdateBrand) GetBusinessContactEmailOk() (*string, bool)`

GetBusinessContactEmailOk returns a tuple with the BusinessContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessContactEmail

`func (o *UpdateBrand) SetBusinessContactEmail(v string)`

SetBusinessContactEmail sets BusinessContactEmail field to given value.

### HasBusinessContactEmail

`func (o *UpdateBrand) HasBusinessContactEmail() bool`

HasBusinessContactEmail returns a boolean if a field has been set.

### GetWebhookURL

`func (o *UpdateBrand) GetWebhookURL() string`

GetWebhookURL returns the WebhookURL field if non-nil, zero value otherwise.

### GetWebhookURLOk

`func (o *UpdateBrand) GetWebhookURLOk() (*string, bool)`

GetWebhookURLOk returns a tuple with the WebhookURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookURL

`func (o *UpdateBrand) SetWebhookURL(v string)`

SetWebhookURL sets WebhookURL field to given value.

### HasWebhookURL

`func (o *UpdateBrand) HasWebhookURL() bool`

HasWebhookURL returns a boolean if a field has been set.

### GetWebhookFailoverURL

`func (o *UpdateBrand) GetWebhookFailoverURL() string`

GetWebhookFailoverURL returns the WebhookFailoverURL field if non-nil, zero value otherwise.

### GetWebhookFailoverURLOk

`func (o *UpdateBrand) GetWebhookFailoverURLOk() (*string, bool)`

GetWebhookFailoverURLOk returns a tuple with the WebhookFailoverURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailoverURL

`func (o *UpdateBrand) SetWebhookFailoverURL(v string)`

SetWebhookFailoverURL sets WebhookFailoverURL field to given value.

### HasWebhookFailoverURL

`func (o *UpdateBrand) HasWebhookFailoverURL() bool`

HasWebhookFailoverURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


