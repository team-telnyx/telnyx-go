# CnamListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CnamListingEnabled** | Pointer to **bool** | Enables CNAM listings for this number. Requires cnam_listing_details to also be set. | [optional] [default to false]
**CnamListingDetails** | Pointer to **string** | The CNAM listing details for this number. Must be alphanumeric characters or spaces with a maximum length of 15. Requires cnam_listing_enabled to also be set to true. | [optional] 

## Methods

### NewCnamListing

`func NewCnamListing() *CnamListing`

NewCnamListing instantiates a new CnamListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCnamListingWithDefaults

`func NewCnamListingWithDefaults() *CnamListing`

NewCnamListingWithDefaults instantiates a new CnamListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCnamListingEnabled

`func (o *CnamListing) GetCnamListingEnabled() bool`

GetCnamListingEnabled returns the CnamListingEnabled field if non-nil, zero value otherwise.

### GetCnamListingEnabledOk

`func (o *CnamListing) GetCnamListingEnabledOk() (*bool, bool)`

GetCnamListingEnabledOk returns a tuple with the CnamListingEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnamListingEnabled

`func (o *CnamListing) SetCnamListingEnabled(v bool)`

SetCnamListingEnabled sets CnamListingEnabled field to given value.

### HasCnamListingEnabled

`func (o *CnamListing) HasCnamListingEnabled() bool`

HasCnamListingEnabled returns a boolean if a field has been set.

### GetCnamListingDetails

`func (o *CnamListing) GetCnamListingDetails() string`

GetCnamListingDetails returns the CnamListingDetails field if non-nil, zero value otherwise.

### GetCnamListingDetailsOk

`func (o *CnamListing) GetCnamListingDetailsOk() (*string, bool)`

GetCnamListingDetailsOk returns a tuple with the CnamListingDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCnamListingDetails

`func (o *CnamListing) SetCnamListingDetails(v string)`

SetCnamListingDetails sets CnamListingDetails field to given value.

### HasCnamListingDetails

`func (o *CnamListing) HasCnamListingDetails() bool`

HasCnamListingDetails returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


