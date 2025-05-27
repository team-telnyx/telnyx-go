# LocationResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LocationId** | Pointer to **string** |  | [optional] 
**StaticEmergencyAddressId** | Pointer to **string** |  | [optional] 
**AcceptedAddressSuggestions** | Pointer to **bool** |  | [optional] 

## Methods

### NewLocationResponseData

`func NewLocationResponseData() *LocationResponseData`

NewLocationResponseData instantiates a new LocationResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationResponseDataWithDefaults

`func NewLocationResponseDataWithDefaults() *LocationResponseData`

NewLocationResponseDataWithDefaults instantiates a new LocationResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocationId

`func (o *LocationResponseData) GetLocationId() string`

GetLocationId returns the LocationId field if non-nil, zero value otherwise.

### GetLocationIdOk

`func (o *LocationResponseData) GetLocationIdOk() (*string, bool)`

GetLocationIdOk returns a tuple with the LocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocationId

`func (o *LocationResponseData) SetLocationId(v string)`

SetLocationId sets LocationId field to given value.

### HasLocationId

`func (o *LocationResponseData) HasLocationId() bool`

HasLocationId returns a boolean if a field has been set.

### GetStaticEmergencyAddressId

`func (o *LocationResponseData) GetStaticEmergencyAddressId() string`

GetStaticEmergencyAddressId returns the StaticEmergencyAddressId field if non-nil, zero value otherwise.

### GetStaticEmergencyAddressIdOk

`func (o *LocationResponseData) GetStaticEmergencyAddressIdOk() (*string, bool)`

GetStaticEmergencyAddressIdOk returns a tuple with the StaticEmergencyAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStaticEmergencyAddressId

`func (o *LocationResponseData) SetStaticEmergencyAddressId(v string)`

SetStaticEmergencyAddressId sets StaticEmergencyAddressId field to given value.

### HasStaticEmergencyAddressId

`func (o *LocationResponseData) HasStaticEmergencyAddressId() bool`

HasStaticEmergencyAddressId returns a boolean if a field has been set.

### GetAcceptedAddressSuggestions

`func (o *LocationResponseData) GetAcceptedAddressSuggestions() bool`

GetAcceptedAddressSuggestions returns the AcceptedAddressSuggestions field if non-nil, zero value otherwise.

### GetAcceptedAddressSuggestionsOk

`func (o *LocationResponseData) GetAcceptedAddressSuggestionsOk() (*bool, bool)`

GetAcceptedAddressSuggestionsOk returns a tuple with the AcceptedAddressSuggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAddressSuggestions

`func (o *LocationResponseData) SetAcceptedAddressSuggestions(v bool)`

SetAcceptedAddressSuggestions sets AcceptedAddressSuggestions field to given value.

### HasAcceptedAddressSuggestions

`func (o *LocationResponseData) HasAcceptedAddressSuggestions() bool`

HasAcceptedAddressSuggestions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


