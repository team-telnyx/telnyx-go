# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Region** | Pointer to **string** | Identifies the geographical region of location. | [optional] 
**Site** | Pointer to **string** | Site of location. | [optional] 
**Pop** | Pointer to **string** | Point of presence of location. | [optional] 
**Code** | Pointer to **string** | Location code. | [optional] 
**Name** | Pointer to **string** | Human readable name of location. | [optional] 

## Methods

### NewLocation

`func NewLocation() *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRegion

`func (o *Location) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Location) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Location) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Location) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSite

`func (o *Location) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *Location) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *Location) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *Location) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetPop

`func (o *Location) GetPop() string`

GetPop returns the Pop field if non-nil, zero value otherwise.

### GetPopOk

`func (o *Location) GetPopOk() (*string, bool)`

GetPopOk returns a tuple with the Pop field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPop

`func (o *Location) SetPop(v string)`

SetPop sets Pop field to given value.

### HasPop

`func (o *Location) HasPop() bool`

HasPop returns a boolean if a field has been set.

### GetCode

`func (o *Location) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Location) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Location) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Location) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *Location) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Location) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Location) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Location) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


