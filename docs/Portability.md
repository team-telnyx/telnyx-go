# Portability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Lrn** | Pointer to **string** | Local Routing Number, if assigned to the requested phone number | [optional] 
**PortedStatus** | Pointer to **string** | Indicates whether or not the requested phone number has been ported | [optional] 
**PortedDate** | Pointer to **string** | ISO-formatted date when the requested phone number has been ported | [optional] 
**Ocn** | Pointer to **string** | Operating Company Name (OCN) as per the Local Exchange Routing Guide (LERG) database | [optional] 
**LineType** | Pointer to **string** | Type of number | [optional] 
**Spid** | Pointer to **string** | SPID (Service Provider ID) | [optional] 
**SpidCarrierName** | Pointer to **string** | Service provider name | [optional] 
**SpidCarrierType** | Pointer to **string** | Service provider type | [optional] 
**Altspid** | Pointer to **string** | Alternative SPID (Service Provider ID). Often used when a carrier is using a number from another carrier | [optional] 
**AltspidCarrierName** | Pointer to **string** | Alternative service provider name | [optional] 
**AltspidCarrierType** | Pointer to **string** | Alternative service provider type | [optional] 
**City** | Pointer to **string** | City name extracted from the locality in the Local Exchange Routing Guide (LERG) database | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewPortability

`func NewPortability() *Portability`

NewPortability instantiates a new Portability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortabilityWithDefaults

`func NewPortabilityWithDefaults() *Portability`

NewPortabilityWithDefaults instantiates a new Portability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLrn

`func (o *Portability) GetLrn() string`

GetLrn returns the Lrn field if non-nil, zero value otherwise.

### GetLrnOk

`func (o *Portability) GetLrnOk() (*string, bool)`

GetLrnOk returns a tuple with the Lrn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLrn

`func (o *Portability) SetLrn(v string)`

SetLrn sets Lrn field to given value.

### HasLrn

`func (o *Portability) HasLrn() bool`

HasLrn returns a boolean if a field has been set.

### GetPortedStatus

`func (o *Portability) GetPortedStatus() string`

GetPortedStatus returns the PortedStatus field if non-nil, zero value otherwise.

### GetPortedStatusOk

`func (o *Portability) GetPortedStatusOk() (*string, bool)`

GetPortedStatusOk returns a tuple with the PortedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortedStatus

`func (o *Portability) SetPortedStatus(v string)`

SetPortedStatus sets PortedStatus field to given value.

### HasPortedStatus

`func (o *Portability) HasPortedStatus() bool`

HasPortedStatus returns a boolean if a field has been set.

### GetPortedDate

`func (o *Portability) GetPortedDate() string`

GetPortedDate returns the PortedDate field if non-nil, zero value otherwise.

### GetPortedDateOk

`func (o *Portability) GetPortedDateOk() (*string, bool)`

GetPortedDateOk returns a tuple with the PortedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortedDate

`func (o *Portability) SetPortedDate(v string)`

SetPortedDate sets PortedDate field to given value.

### HasPortedDate

`func (o *Portability) HasPortedDate() bool`

HasPortedDate returns a boolean if a field has been set.

### GetOcn

`func (o *Portability) GetOcn() string`

GetOcn returns the Ocn field if non-nil, zero value otherwise.

### GetOcnOk

`func (o *Portability) GetOcnOk() (*string, bool)`

GetOcnOk returns a tuple with the Ocn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOcn

`func (o *Portability) SetOcn(v string)`

SetOcn sets Ocn field to given value.

### HasOcn

`func (o *Portability) HasOcn() bool`

HasOcn returns a boolean if a field has been set.

### GetLineType

`func (o *Portability) GetLineType() string`

GetLineType returns the LineType field if non-nil, zero value otherwise.

### GetLineTypeOk

`func (o *Portability) GetLineTypeOk() (*string, bool)`

GetLineTypeOk returns a tuple with the LineType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineType

`func (o *Portability) SetLineType(v string)`

SetLineType sets LineType field to given value.

### HasLineType

`func (o *Portability) HasLineType() bool`

HasLineType returns a boolean if a field has been set.

### GetSpid

`func (o *Portability) GetSpid() string`

GetSpid returns the Spid field if non-nil, zero value otherwise.

### GetSpidOk

`func (o *Portability) GetSpidOk() (*string, bool)`

GetSpidOk returns a tuple with the Spid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpid

`func (o *Portability) SetSpid(v string)`

SetSpid sets Spid field to given value.

### HasSpid

`func (o *Portability) HasSpid() bool`

HasSpid returns a boolean if a field has been set.

### GetSpidCarrierName

`func (o *Portability) GetSpidCarrierName() string`

GetSpidCarrierName returns the SpidCarrierName field if non-nil, zero value otherwise.

### GetSpidCarrierNameOk

`func (o *Portability) GetSpidCarrierNameOk() (*string, bool)`

GetSpidCarrierNameOk returns a tuple with the SpidCarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpidCarrierName

`func (o *Portability) SetSpidCarrierName(v string)`

SetSpidCarrierName sets SpidCarrierName field to given value.

### HasSpidCarrierName

`func (o *Portability) HasSpidCarrierName() bool`

HasSpidCarrierName returns a boolean if a field has been set.

### GetSpidCarrierType

`func (o *Portability) GetSpidCarrierType() string`

GetSpidCarrierType returns the SpidCarrierType field if non-nil, zero value otherwise.

### GetSpidCarrierTypeOk

`func (o *Portability) GetSpidCarrierTypeOk() (*string, bool)`

GetSpidCarrierTypeOk returns a tuple with the SpidCarrierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpidCarrierType

`func (o *Portability) SetSpidCarrierType(v string)`

SetSpidCarrierType sets SpidCarrierType field to given value.

### HasSpidCarrierType

`func (o *Portability) HasSpidCarrierType() bool`

HasSpidCarrierType returns a boolean if a field has been set.

### GetAltspid

`func (o *Portability) GetAltspid() string`

GetAltspid returns the Altspid field if non-nil, zero value otherwise.

### GetAltspidOk

`func (o *Portability) GetAltspidOk() (*string, bool)`

GetAltspidOk returns a tuple with the Altspid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltspid

`func (o *Portability) SetAltspid(v string)`

SetAltspid sets Altspid field to given value.

### HasAltspid

`func (o *Portability) HasAltspid() bool`

HasAltspid returns a boolean if a field has been set.

### GetAltspidCarrierName

`func (o *Portability) GetAltspidCarrierName() string`

GetAltspidCarrierName returns the AltspidCarrierName field if non-nil, zero value otherwise.

### GetAltspidCarrierNameOk

`func (o *Portability) GetAltspidCarrierNameOk() (*string, bool)`

GetAltspidCarrierNameOk returns a tuple with the AltspidCarrierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltspidCarrierName

`func (o *Portability) SetAltspidCarrierName(v string)`

SetAltspidCarrierName sets AltspidCarrierName field to given value.

### HasAltspidCarrierName

`func (o *Portability) HasAltspidCarrierName() bool`

HasAltspidCarrierName returns a boolean if a field has been set.

### GetAltspidCarrierType

`func (o *Portability) GetAltspidCarrierType() string`

GetAltspidCarrierType returns the AltspidCarrierType field if non-nil, zero value otherwise.

### GetAltspidCarrierTypeOk

`func (o *Portability) GetAltspidCarrierTypeOk() (*string, bool)`

GetAltspidCarrierTypeOk returns a tuple with the AltspidCarrierType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAltspidCarrierType

`func (o *Portability) SetAltspidCarrierType(v string)`

SetAltspidCarrierType sets AltspidCarrierType field to given value.

### HasAltspidCarrierType

`func (o *Portability) HasAltspidCarrierType() bool`

HasAltspidCarrierType returns a boolean if a field has been set.

### GetCity

`func (o *Portability) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Portability) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Portability) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *Portability) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetState

`func (o *Portability) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Portability) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Portability) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Portability) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


