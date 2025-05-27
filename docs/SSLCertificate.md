# SSLCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | Unique identifier for the SSL certificate | [optional] 
**IssuedTo** | Pointer to [**SSLCertificateIssuedTo**](SSLCertificateIssuedTo.md) |  | [optional] 
**IssuedBy** | Pointer to [**SSLCertificateIssuedBy**](SSLCertificateIssuedBy.md) |  | [optional] 
**ValidFrom** | Pointer to **time.Time** | The time the certificate is valid from | [optional] 
**ValidTo** | Pointer to **time.Time** | The time the certificate is valid to | [optional] 
**CreatedAt** | Pointer to **time.Time** | Time when SSL certificate was uploaded | [optional] 

## Methods

### NewSSLCertificate

`func NewSSLCertificate() *SSLCertificate`

NewSSLCertificate instantiates a new SSLCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSSLCertificateWithDefaults

`func NewSSLCertificateWithDefaults() *SSLCertificate`

NewSSLCertificateWithDefaults instantiates a new SSLCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SSLCertificate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SSLCertificate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SSLCertificate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SSLCertificate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIssuedTo

`func (o *SSLCertificate) GetIssuedTo() SSLCertificateIssuedTo`

GetIssuedTo returns the IssuedTo field if non-nil, zero value otherwise.

### GetIssuedToOk

`func (o *SSLCertificate) GetIssuedToOk() (*SSLCertificateIssuedTo, bool)`

GetIssuedToOk returns a tuple with the IssuedTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedTo

`func (o *SSLCertificate) SetIssuedTo(v SSLCertificateIssuedTo)`

SetIssuedTo sets IssuedTo field to given value.

### HasIssuedTo

`func (o *SSLCertificate) HasIssuedTo() bool`

HasIssuedTo returns a boolean if a field has been set.

### GetIssuedBy

`func (o *SSLCertificate) GetIssuedBy() SSLCertificateIssuedBy`

GetIssuedBy returns the IssuedBy field if non-nil, zero value otherwise.

### GetIssuedByOk

`func (o *SSLCertificate) GetIssuedByOk() (*SSLCertificateIssuedBy, bool)`

GetIssuedByOk returns a tuple with the IssuedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssuedBy

`func (o *SSLCertificate) SetIssuedBy(v SSLCertificateIssuedBy)`

SetIssuedBy sets IssuedBy field to given value.

### HasIssuedBy

`func (o *SSLCertificate) HasIssuedBy() bool`

HasIssuedBy returns a boolean if a field has been set.

### GetValidFrom

`func (o *SSLCertificate) GetValidFrom() time.Time`

GetValidFrom returns the ValidFrom field if non-nil, zero value otherwise.

### GetValidFromOk

`func (o *SSLCertificate) GetValidFromOk() (*time.Time, bool)`

GetValidFromOk returns a tuple with the ValidFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidFrom

`func (o *SSLCertificate) SetValidFrom(v time.Time)`

SetValidFrom sets ValidFrom field to given value.

### HasValidFrom

`func (o *SSLCertificate) HasValidFrom() bool`

HasValidFrom returns a boolean if a field has been set.

### GetValidTo

`func (o *SSLCertificate) GetValidTo() time.Time`

GetValidTo returns the ValidTo field if non-nil, zero value otherwise.

### GetValidToOk

`func (o *SSLCertificate) GetValidToOk() (*time.Time, bool)`

GetValidToOk returns a tuple with the ValidTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidTo

`func (o *SSLCertificate) SetValidTo(v time.Time)`

SetValidTo sets ValidTo field to given value.

### HasValidTo

`func (o *SSLCertificate) HasValidTo() bool`

HasValidTo returns a boolean if a field has been set.

### GetCreatedAt

`func (o *SSLCertificate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SSLCertificate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SSLCertificate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SSLCertificate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


