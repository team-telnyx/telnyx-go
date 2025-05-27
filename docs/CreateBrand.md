# CreateBrand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityType** | [**EntityType**](EntityType.md) | Entity type behind the brand. This is the form of business establishment. | 
**DisplayName** | **string** | Display name, marketing name, or DBA name of the brand. | 
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
**IsReseller** | Pointer to **bool** |  | [optional] [default to false]
**Mock** | Pointer to **bool** | Mock brand for testing purposes. Defaults to false. | [optional] [default to false]
**MobilePhone** | Pointer to **string** | Valid mobile phone number in e.164 international format. | [optional] 
**BusinessContactEmail** | Pointer to **string** | Business contact email.  Required if &#x60;entityType&#x60; is &#x60;PUBLIC_PROFIT&#x60;. | [optional] 
**WebhookURL** | Pointer to **string** | Webhook URL for brand status updates. | [optional] 
**WebhookFailoverURL** | Pointer to **string** | Webhook failover URL for brand status updates. | [optional] 

## Methods

### NewCreateBrand

`func NewCreateBrand(entityType EntityType, displayName string, country string, email string, vertical Vertical, ) *CreateBrand`

NewCreateBrand instantiates a new CreateBrand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBrandWithDefaults

`func NewCreateBrandWithDefaults() *CreateBrand`

NewCreateBrandWithDefaults instantiates a new CreateBrand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityType

`func (o *CreateBrand) GetEntityType() EntityType`

GetEntityType returns the EntityType field if non-nil, zero value otherwise.

### GetEntityTypeOk

`func (o *CreateBrand) GetEntityTypeOk() (*EntityType, bool)`

GetEntityTypeOk returns a tuple with the EntityType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityType

`func (o *CreateBrand) SetEntityType(v EntityType)`

SetEntityType sets EntityType field to given value.


### GetDisplayName

`func (o *CreateBrand) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CreateBrand) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CreateBrand) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.


### GetCompanyName

`func (o *CreateBrand) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *CreateBrand) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *CreateBrand) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *CreateBrand) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetFirstName

`func (o *CreateBrand) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CreateBrand) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CreateBrand) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CreateBrand) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *CreateBrand) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CreateBrand) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CreateBrand) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CreateBrand) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetEin

`func (o *CreateBrand) GetEin() string`

GetEin returns the Ein field if non-nil, zero value otherwise.

### GetEinOk

`func (o *CreateBrand) GetEinOk() (*string, bool)`

GetEinOk returns a tuple with the Ein field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEin

`func (o *CreateBrand) SetEin(v string)`

SetEin sets Ein field to given value.

### HasEin

`func (o *CreateBrand) HasEin() bool`

HasEin returns a boolean if a field has been set.

### GetPhone

`func (o *CreateBrand) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CreateBrand) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CreateBrand) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CreateBrand) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetStreet

`func (o *CreateBrand) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *CreateBrand) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *CreateBrand) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *CreateBrand) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetCity

`func (o *CreateBrand) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *CreateBrand) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *CreateBrand) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *CreateBrand) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *CreateBrand) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateBrand) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateBrand) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CreateBrand) HasState() bool`

HasState returns a boolean if a field has been set.

### GetPostalCode

`func (o *CreateBrand) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *CreateBrand) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *CreateBrand) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *CreateBrand) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *CreateBrand) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *CreateBrand) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *CreateBrand) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetEmail

`func (o *CreateBrand) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CreateBrand) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CreateBrand) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetStockSymbol

`func (o *CreateBrand) GetStockSymbol() string`

GetStockSymbol returns the StockSymbol field if non-nil, zero value otherwise.

### GetStockSymbolOk

`func (o *CreateBrand) GetStockSymbolOk() (*string, bool)`

GetStockSymbolOk returns a tuple with the StockSymbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockSymbol

`func (o *CreateBrand) SetStockSymbol(v string)`

SetStockSymbol sets StockSymbol field to given value.

### HasStockSymbol

`func (o *CreateBrand) HasStockSymbol() bool`

HasStockSymbol returns a boolean if a field has been set.

### GetStockExchange

`func (o *CreateBrand) GetStockExchange() StockExchange`

GetStockExchange returns the StockExchange field if non-nil, zero value otherwise.

### GetStockExchangeOk

`func (o *CreateBrand) GetStockExchangeOk() (*StockExchange, bool)`

GetStockExchangeOk returns a tuple with the StockExchange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockExchange

`func (o *CreateBrand) SetStockExchange(v StockExchange)`

SetStockExchange sets StockExchange field to given value.

### HasStockExchange

`func (o *CreateBrand) HasStockExchange() bool`

HasStockExchange returns a boolean if a field has been set.

### GetIpAddress

`func (o *CreateBrand) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *CreateBrand) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *CreateBrand) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *CreateBrand) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### GetWebsite

`func (o *CreateBrand) GetWebsite() string`

GetWebsite returns the Website field if non-nil, zero value otherwise.

### GetWebsiteOk

`func (o *CreateBrand) GetWebsiteOk() (*string, bool)`

GetWebsiteOk returns a tuple with the Website field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsite

`func (o *CreateBrand) SetWebsite(v string)`

SetWebsite sets Website field to given value.

### HasWebsite

`func (o *CreateBrand) HasWebsite() bool`

HasWebsite returns a boolean if a field has been set.

### GetVertical

`func (o *CreateBrand) GetVertical() Vertical`

GetVertical returns the Vertical field if non-nil, zero value otherwise.

### GetVerticalOk

`func (o *CreateBrand) GetVerticalOk() (*Vertical, bool)`

GetVerticalOk returns a tuple with the Vertical field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVertical

`func (o *CreateBrand) SetVertical(v Vertical)`

SetVertical sets Vertical field to given value.


### GetIsReseller

`func (o *CreateBrand) GetIsReseller() bool`

GetIsReseller returns the IsReseller field if non-nil, zero value otherwise.

### GetIsResellerOk

`func (o *CreateBrand) GetIsResellerOk() (*bool, bool)`

GetIsResellerOk returns a tuple with the IsReseller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsReseller

`func (o *CreateBrand) SetIsReseller(v bool)`

SetIsReseller sets IsReseller field to given value.

### HasIsReseller

`func (o *CreateBrand) HasIsReseller() bool`

HasIsReseller returns a boolean if a field has been set.

### GetMock

`func (o *CreateBrand) GetMock() bool`

GetMock returns the Mock field if non-nil, zero value otherwise.

### GetMockOk

`func (o *CreateBrand) GetMockOk() (*bool, bool)`

GetMockOk returns a tuple with the Mock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMock

`func (o *CreateBrand) SetMock(v bool)`

SetMock sets Mock field to given value.

### HasMock

`func (o *CreateBrand) HasMock() bool`

HasMock returns a boolean if a field has been set.

### GetMobilePhone

`func (o *CreateBrand) GetMobilePhone() string`

GetMobilePhone returns the MobilePhone field if non-nil, zero value otherwise.

### GetMobilePhoneOk

`func (o *CreateBrand) GetMobilePhoneOk() (*string, bool)`

GetMobilePhoneOk returns a tuple with the MobilePhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhone

`func (o *CreateBrand) SetMobilePhone(v string)`

SetMobilePhone sets MobilePhone field to given value.

### HasMobilePhone

`func (o *CreateBrand) HasMobilePhone() bool`

HasMobilePhone returns a boolean if a field has been set.

### GetBusinessContactEmail

`func (o *CreateBrand) GetBusinessContactEmail() string`

GetBusinessContactEmail returns the BusinessContactEmail field if non-nil, zero value otherwise.

### GetBusinessContactEmailOk

`func (o *CreateBrand) GetBusinessContactEmailOk() (*string, bool)`

GetBusinessContactEmailOk returns a tuple with the BusinessContactEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessContactEmail

`func (o *CreateBrand) SetBusinessContactEmail(v string)`

SetBusinessContactEmail sets BusinessContactEmail field to given value.

### HasBusinessContactEmail

`func (o *CreateBrand) HasBusinessContactEmail() bool`

HasBusinessContactEmail returns a boolean if a field has been set.

### GetWebhookURL

`func (o *CreateBrand) GetWebhookURL() string`

GetWebhookURL returns the WebhookURL field if non-nil, zero value otherwise.

### GetWebhookURLOk

`func (o *CreateBrand) GetWebhookURLOk() (*string, bool)`

GetWebhookURLOk returns a tuple with the WebhookURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookURL

`func (o *CreateBrand) SetWebhookURL(v string)`

SetWebhookURL sets WebhookURL field to given value.

### HasWebhookURL

`func (o *CreateBrand) HasWebhookURL() bool`

HasWebhookURL returns a boolean if a field has been set.

### GetWebhookFailoverURL

`func (o *CreateBrand) GetWebhookFailoverURL() string`

GetWebhookFailoverURL returns the WebhookFailoverURL field if non-nil, zero value otherwise.

### GetWebhookFailoverURLOk

`func (o *CreateBrand) GetWebhookFailoverURLOk() (*string, bool)`

GetWebhookFailoverURLOk returns a tuple with the WebhookFailoverURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookFailoverURL

`func (o *CreateBrand) SetWebhookFailoverURL(v string)`

SetWebhookFailoverURL sets WebhookFailoverURL field to given value.

### HasWebhookFailoverURL

`func (o *CreateBrand) HasWebhookFailoverURL() bool`

HasWebhookFailoverURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


