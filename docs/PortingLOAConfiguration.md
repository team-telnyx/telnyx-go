# PortingLOAConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Uniquely identifies the LOA configuration. | [optional] 
**CompanyName** | Pointer to **string** | The name of the company | [optional] 
**OrganizationId** | Pointer to **string** | The organization that owns the LOA configuration | [optional] 
**Name** | Pointer to **string** | The name of the LOA configuration | [optional] 
**Logo** | Pointer to [**PortingLOAConfigurationLogo**](PortingLOAConfigurationLogo.md) |  | [optional] 
**Address** | Pointer to [**PortingLOAConfigurationAddress**](PortingLOAConfigurationAddress.md) |  | [optional] 
**Contact** | Pointer to [**PortingLOAConfigurationContact**](PortingLOAConfigurationContact.md) |  | [optional] 
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was created. | [optional] 
**UpdatedAt** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the resource was updated. | [optional] 

## Methods

### NewPortingLOAConfiguration

`func NewPortingLOAConfiguration() *PortingLOAConfiguration`

NewPortingLOAConfiguration instantiates a new PortingLOAConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortingLOAConfigurationWithDefaults

`func NewPortingLOAConfigurationWithDefaults() *PortingLOAConfiguration`

NewPortingLOAConfigurationWithDefaults instantiates a new PortingLOAConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PortingLOAConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PortingLOAConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PortingLOAConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PortingLOAConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCompanyName

`func (o *PortingLOAConfiguration) GetCompanyName() string`

GetCompanyName returns the CompanyName field if non-nil, zero value otherwise.

### GetCompanyNameOk

`func (o *PortingLOAConfiguration) GetCompanyNameOk() (*string, bool)`

GetCompanyNameOk returns a tuple with the CompanyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompanyName

`func (o *PortingLOAConfiguration) SetCompanyName(v string)`

SetCompanyName sets CompanyName field to given value.

### HasCompanyName

`func (o *PortingLOAConfiguration) HasCompanyName() bool`

HasCompanyName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *PortingLOAConfiguration) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *PortingLOAConfiguration) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *PortingLOAConfiguration) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *PortingLOAConfiguration) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetName

`func (o *PortingLOAConfiguration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PortingLOAConfiguration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PortingLOAConfiguration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PortingLOAConfiguration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLogo

`func (o *PortingLOAConfiguration) GetLogo() PortingLOAConfigurationLogo`

GetLogo returns the Logo field if non-nil, zero value otherwise.

### GetLogoOk

`func (o *PortingLOAConfiguration) GetLogoOk() (*PortingLOAConfigurationLogo, bool)`

GetLogoOk returns a tuple with the Logo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogo

`func (o *PortingLOAConfiguration) SetLogo(v PortingLOAConfigurationLogo)`

SetLogo sets Logo field to given value.

### HasLogo

`func (o *PortingLOAConfiguration) HasLogo() bool`

HasLogo returns a boolean if a field has been set.

### GetAddress

`func (o *PortingLOAConfiguration) GetAddress() PortingLOAConfigurationAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PortingLOAConfiguration) GetAddressOk() (*PortingLOAConfigurationAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PortingLOAConfiguration) SetAddress(v PortingLOAConfigurationAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PortingLOAConfiguration) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetContact

`func (o *PortingLOAConfiguration) GetContact() PortingLOAConfigurationContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *PortingLOAConfiguration) GetContactOk() (*PortingLOAConfigurationContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *PortingLOAConfiguration) SetContact(v PortingLOAConfigurationContact)`

SetContact sets Contact field to given value.

### HasContact

`func (o *PortingLOAConfiguration) HasContact() bool`

HasContact returns a boolean if a field has been set.

### GetRecordType

`func (o *PortingLOAConfiguration) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *PortingLOAConfiguration) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *PortingLOAConfiguration) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *PortingLOAConfiguration) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetCreatedAt

`func (o *PortingLOAConfiguration) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PortingLOAConfiguration) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PortingLOAConfiguration) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *PortingLOAConfiguration) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *PortingLOAConfiguration) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PortingLOAConfiguration) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PortingLOAConfiguration) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PortingLOAConfiguration) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


