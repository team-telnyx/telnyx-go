# RCSLatLng

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Latitude** | **float32** | The latitude in degrees. It must be in the range [-90.0, +90.0]. | 
**Longitude** | **float32** | The longitude in degrees. It must be in the range [-180.0, +180.0]. | 

## Methods

### NewRCSLatLng

`func NewRCSLatLng(latitude float32, longitude float32, ) *RCSLatLng`

NewRCSLatLng instantiates a new RCSLatLng object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCSLatLngWithDefaults

`func NewRCSLatLngWithDefaults() *RCSLatLng`

NewRCSLatLngWithDefaults instantiates a new RCSLatLng object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLatitude

`func (o *RCSLatLng) GetLatitude() float32`

GetLatitude returns the Latitude field if non-nil, zero value otherwise.

### GetLatitudeOk

`func (o *RCSLatLng) GetLatitudeOk() (*float32, bool)`

GetLatitudeOk returns a tuple with the Latitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatitude

`func (o *RCSLatLng) SetLatitude(v float32)`

SetLatitude sets Latitude field to given value.


### GetLongitude

`func (o *RCSLatLng) GetLongitude() float32`

GetLongitude returns the Longitude field if non-nil, zero value otherwise.

### GetLongitudeOk

`func (o *RCSLatLng) GetLongitudeOk() (*float32, bool)`

GetLongitudeOk returns a tuple with the Longitude field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongitude

`func (o *RCSLatLng) SetLongitude(v float32)`

SetLongitude sets Longitude field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


