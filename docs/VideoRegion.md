# VideoRegion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**XPos** | Pointer to **NullableInt32** | X axis value (in pixels) of the region&#39;s upper left corner relative to the upper left corner of the whole room composition viewport. | [optional] [default to 0]
**YPos** | Pointer to **NullableInt32** | Y axis value (in pixels) of the region&#39;s upper left corner relative to the upper left corner of the whole room composition viewport. | [optional] [default to 0]
**ZPos** | Pointer to **NullableInt32** | Regions with higher z_pos values are stacked on top of regions with lower z_pos values | [optional] [default to 0]
**Height** | Pointer to **NullableInt32** | Height of the video region | [optional] 
**Width** | Pointer to **NullableInt32** | Width of the video region | [optional] 
**MaxColumns** | Pointer to **NullableInt32** | Maximum number of columns of the region&#39;s placement grid. By default, the region has as many columns as needed to layout all the specified video sources. | [optional] 
**MaxRows** | Pointer to **NullableInt32** | Maximum number of rows of the region&#39;s placement grid. By default, the region has as many rows as needed to layout all the specified video sources. | [optional] 
**VideoSources** | Pointer to **[]string** | Array of video recording ids to be composed in the region. Can be \&quot;*\&quot; to specify all video recordings in the session | [optional] 

## Methods

### NewVideoRegion

`func NewVideoRegion() *VideoRegion`

NewVideoRegion instantiates a new VideoRegion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVideoRegionWithDefaults

`func NewVideoRegionWithDefaults() *VideoRegion`

NewVideoRegionWithDefaults instantiates a new VideoRegion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetXPos

`func (o *VideoRegion) GetXPos() int32`

GetXPos returns the XPos field if non-nil, zero value otherwise.

### GetXPosOk

`func (o *VideoRegion) GetXPosOk() (*int32, bool)`

GetXPosOk returns a tuple with the XPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetXPos

`func (o *VideoRegion) SetXPos(v int32)`

SetXPos sets XPos field to given value.

### HasXPos

`func (o *VideoRegion) HasXPos() bool`

HasXPos returns a boolean if a field has been set.

### SetXPosNil

`func (o *VideoRegion) SetXPosNil(b bool)`

 SetXPosNil sets the value for XPos to be an explicit nil

### UnsetXPos
`func (o *VideoRegion) UnsetXPos()`

UnsetXPos ensures that no value is present for XPos, not even an explicit nil
### GetYPos

`func (o *VideoRegion) GetYPos() int32`

GetYPos returns the YPos field if non-nil, zero value otherwise.

### GetYPosOk

`func (o *VideoRegion) GetYPosOk() (*int32, bool)`

GetYPosOk returns a tuple with the YPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYPos

`func (o *VideoRegion) SetYPos(v int32)`

SetYPos sets YPos field to given value.

### HasYPos

`func (o *VideoRegion) HasYPos() bool`

HasYPos returns a boolean if a field has been set.

### SetYPosNil

`func (o *VideoRegion) SetYPosNil(b bool)`

 SetYPosNil sets the value for YPos to be an explicit nil

### UnsetYPos
`func (o *VideoRegion) UnsetYPos()`

UnsetYPos ensures that no value is present for YPos, not even an explicit nil
### GetZPos

`func (o *VideoRegion) GetZPos() int32`

GetZPos returns the ZPos field if non-nil, zero value otherwise.

### GetZPosOk

`func (o *VideoRegion) GetZPosOk() (*int32, bool)`

GetZPosOk returns a tuple with the ZPos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZPos

`func (o *VideoRegion) SetZPos(v int32)`

SetZPos sets ZPos field to given value.

### HasZPos

`func (o *VideoRegion) HasZPos() bool`

HasZPos returns a boolean if a field has been set.

### SetZPosNil

`func (o *VideoRegion) SetZPosNil(b bool)`

 SetZPosNil sets the value for ZPos to be an explicit nil

### UnsetZPos
`func (o *VideoRegion) UnsetZPos()`

UnsetZPos ensures that no value is present for ZPos, not even an explicit nil
### GetHeight

`func (o *VideoRegion) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *VideoRegion) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *VideoRegion) SetHeight(v int32)`

SetHeight sets Height field to given value.

### HasHeight

`func (o *VideoRegion) HasHeight() bool`

HasHeight returns a boolean if a field has been set.

### SetHeightNil

`func (o *VideoRegion) SetHeightNil(b bool)`

 SetHeightNil sets the value for Height to be an explicit nil

### UnsetHeight
`func (o *VideoRegion) UnsetHeight()`

UnsetHeight ensures that no value is present for Height, not even an explicit nil
### GetWidth

`func (o *VideoRegion) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *VideoRegion) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *VideoRegion) SetWidth(v int32)`

SetWidth sets Width field to given value.

### HasWidth

`func (o *VideoRegion) HasWidth() bool`

HasWidth returns a boolean if a field has been set.

### SetWidthNil

`func (o *VideoRegion) SetWidthNil(b bool)`

 SetWidthNil sets the value for Width to be an explicit nil

### UnsetWidth
`func (o *VideoRegion) UnsetWidth()`

UnsetWidth ensures that no value is present for Width, not even an explicit nil
### GetMaxColumns

`func (o *VideoRegion) GetMaxColumns() int32`

GetMaxColumns returns the MaxColumns field if non-nil, zero value otherwise.

### GetMaxColumnsOk

`func (o *VideoRegion) GetMaxColumnsOk() (*int32, bool)`

GetMaxColumnsOk returns a tuple with the MaxColumns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxColumns

`func (o *VideoRegion) SetMaxColumns(v int32)`

SetMaxColumns sets MaxColumns field to given value.

### HasMaxColumns

`func (o *VideoRegion) HasMaxColumns() bool`

HasMaxColumns returns a boolean if a field has been set.

### SetMaxColumnsNil

`func (o *VideoRegion) SetMaxColumnsNil(b bool)`

 SetMaxColumnsNil sets the value for MaxColumns to be an explicit nil

### UnsetMaxColumns
`func (o *VideoRegion) UnsetMaxColumns()`

UnsetMaxColumns ensures that no value is present for MaxColumns, not even an explicit nil
### GetMaxRows

`func (o *VideoRegion) GetMaxRows() int32`

GetMaxRows returns the MaxRows field if non-nil, zero value otherwise.

### GetMaxRowsOk

`func (o *VideoRegion) GetMaxRowsOk() (*int32, bool)`

GetMaxRowsOk returns a tuple with the MaxRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRows

`func (o *VideoRegion) SetMaxRows(v int32)`

SetMaxRows sets MaxRows field to given value.

### HasMaxRows

`func (o *VideoRegion) HasMaxRows() bool`

HasMaxRows returns a boolean if a field has been set.

### SetMaxRowsNil

`func (o *VideoRegion) SetMaxRowsNil(b bool)`

 SetMaxRowsNil sets the value for MaxRows to be an explicit nil

### UnsetMaxRows
`func (o *VideoRegion) UnsetMaxRows()`

UnsetMaxRows ensures that no value is present for MaxRows, not even an explicit nil
### GetVideoSources

`func (o *VideoRegion) GetVideoSources() []string`

GetVideoSources returns the VideoSources field if non-nil, zero value otherwise.

### GetVideoSourcesOk

`func (o *VideoRegion) GetVideoSourcesOk() (*[]string, bool)`

GetVideoSourcesOk returns a tuple with the VideoSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoSources

`func (o *VideoRegion) SetVideoSources(v []string)`

SetVideoSources sets VideoSources field to given value.

### HasVideoSources

`func (o *VideoRegion) HasVideoSources() bool`

HasVideoSources returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


