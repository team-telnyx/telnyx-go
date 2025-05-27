# SSLCertificateIssuedBy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommonName** | Pointer to **string** | The common name of the entity the certificate was issued by | [optional] 
**Organization** | Pointer to **string** | The organization the certificate was issued by | [optional] 
**OrganizationUnit** | Pointer to **string** | The organizational unit the certificate was issued by | [optional] 

## Methods

### NewSSLCertificateIssuedBy

`func NewSSLCertificateIssuedBy() *SSLCertificateIssuedBy`

NewSSLCertificateIssuedBy instantiates a new SSLCertificateIssuedBy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSLCertificateIssuedByWithDefaults

`func NewSSLCertificateIssuedByWithDefaults() *SSLCertificateIssuedBy`

NewSSLCertificateIssuedByWithDefaults instantiates a new SSLCertificateIssuedBy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommonName

`func (o *SSLCertificateIssuedBy) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *SSLCertificateIssuedBy) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *SSLCertificateIssuedBy) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *SSLCertificateIssuedBy) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### GetOrganization

`func (o *SSLCertificateIssuedBy) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *SSLCertificateIssuedBy) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *SSLCertificateIssuedBy) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *SSLCertificateIssuedBy) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOrganizationUnit

`func (o *SSLCertificateIssuedBy) GetOrganizationUnit() string`

GetOrganizationUnit returns the OrganizationUnit field if non-nil, zero value otherwise.

### GetOrganizationUnitOk

`func (o *SSLCertificateIssuedBy) GetOrganizationUnitOk() (*string, bool)`

GetOrganizationUnitOk returns a tuple with the OrganizationUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationUnit

`func (o *SSLCertificateIssuedBy) SetOrganizationUnit(v string)`

SetOrganizationUnit sets OrganizationUnit field to given value.

### HasOrganizationUnit

`func (o *SSLCertificateIssuedBy) HasOrganizationUnit() bool`

HasOrganizationUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


