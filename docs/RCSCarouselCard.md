# RCSCarouselCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CardWidth** | **string** | The width of the cards in the carousel. | 
**CardContents** | [**[]RCSCardContent**](RCSCardContent.md) | The list of contents for each card in the carousel. A carousel can have a minimum of 2 cards and a maximum 10 cards. | 

## Methods

### NewRCSCarouselCard

`func NewRCSCarouselCard(cardWidth string, cardContents []RCSCardContent, ) *RCSCarouselCard`

NewRCSCarouselCard instantiates a new RCSCarouselCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCSCarouselCardWithDefaults

`func NewRCSCarouselCardWithDefaults() *RCSCarouselCard`

NewRCSCarouselCardWithDefaults instantiates a new RCSCarouselCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardWidth

`func (o *RCSCarouselCard) GetCardWidth() string`

GetCardWidth returns the CardWidth field if non-nil, zero value otherwise.

### GetCardWidthOk

`func (o *RCSCarouselCard) GetCardWidthOk() (*string, bool)`

GetCardWidthOk returns a tuple with the CardWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardWidth

`func (o *RCSCarouselCard) SetCardWidth(v string)`

SetCardWidth sets CardWidth field to given value.


### GetCardContents

`func (o *RCSCarouselCard) GetCardContents() []RCSCardContent`

GetCardContents returns the CardContents field if non-nil, zero value otherwise.

### GetCardContentsOk

`func (o *RCSCarouselCard) GetCardContentsOk() (*[]RCSCardContent, bool)`

GetCardContentsOk returns a tuple with the CardContents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardContents

`func (o *RCSCarouselCard) SetCardContents(v []RCSCardContent)`

SetCardContents sets CardContents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


