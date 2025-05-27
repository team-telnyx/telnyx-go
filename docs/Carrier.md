# Carrier

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MobileCountryCode** | Pointer to **string** | Region code that matches the specific country calling code if the requested phone number type is mobile | [optional] 
**MobileNetworkCode** | Pointer to **string** | National destination code (NDC), with a 0 prefix, if an NDC is found and the requested phone number type is mobile | [optional] 
**Name** | Pointer to **string** | SPID (Service Provider ID) name, if the requested phone number has been ported; otherwise, the name of carrier who owns the phone number block | [optional] 
**Type** | Pointer to **string** | A phone number type that identifies the type of service associated with the requested phone number | [optional] 
**ErrorCode** | Pointer to **string** | Unused | [optional] 
**NormalizedCarrier** | Pointer to **string** | If known to Telnyx and applicable, the primary network carrier. | [optional] 

## Methods

### NewCarrier

`func NewCarrier() *Carrier`

NewCarrier instantiates a new Carrier object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCarrierWithDefaults

`func NewCarrierWithDefaults() *Carrier`

NewCarrierWithDefaults instantiates a new Carrier object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMobileCountryCode

`func (o *Carrier) GetMobileCountryCode() string`

GetMobileCountryCode returns the MobileCountryCode field if non-nil, zero value otherwise.

### GetMobileCountryCodeOk

`func (o *Carrier) GetMobileCountryCodeOk() (*string, bool)`

GetMobileCountryCodeOk returns a tuple with the MobileCountryCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileCountryCode

`func (o *Carrier) SetMobileCountryCode(v string)`

SetMobileCountryCode sets MobileCountryCode field to given value.

### HasMobileCountryCode

`func (o *Carrier) HasMobileCountryCode() bool`

HasMobileCountryCode returns a boolean if a field has been set.

### GetMobileNetworkCode

`func (o *Carrier) GetMobileNetworkCode() string`

GetMobileNetworkCode returns the MobileNetworkCode field if non-nil, zero value otherwise.

### GetMobileNetworkCodeOk

`func (o *Carrier) GetMobileNetworkCodeOk() (*string, bool)`

GetMobileNetworkCodeOk returns a tuple with the MobileNetworkCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobileNetworkCode

`func (o *Carrier) SetMobileNetworkCode(v string)`

SetMobileNetworkCode sets MobileNetworkCode field to given value.

### HasMobileNetworkCode

`func (o *Carrier) HasMobileNetworkCode() bool`

HasMobileNetworkCode returns a boolean if a field has been set.

### GetName

`func (o *Carrier) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Carrier) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Carrier) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Carrier) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *Carrier) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Carrier) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Carrier) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Carrier) HasType() bool`

HasType returns a boolean if a field has been set.

### GetErrorCode

`func (o *Carrier) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Carrier) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Carrier) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *Carrier) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetNormalizedCarrier

`func (o *Carrier) GetNormalizedCarrier() string`

GetNormalizedCarrier returns the NormalizedCarrier field if non-nil, zero value otherwise.

### GetNormalizedCarrierOk

`func (o *Carrier) GetNormalizedCarrierOk() (*string, bool)`

GetNormalizedCarrierOk returns a tuple with the NormalizedCarrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedCarrier

`func (o *Carrier) SetNormalizedCarrier(v string)`

SetNormalizedCarrier sets NormalizedCarrier field to given value.

### HasNormalizedCarrier

`func (o *Carrier) HasNormalizedCarrier() bool`

HasNormalizedCarrier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


