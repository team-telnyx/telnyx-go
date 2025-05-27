# PortingOrderUserFeedback

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserRating** | Pointer to **int32** | Once an order is ported, cancellation is requested or the request is cancelled, the user may rate their experience | [optional] 
**UserComment** | Pointer to **string** | A comment related to the customer rating. | [optional] 

## Methods

### NewPortingOrderUserFeedback

`func NewPortingOrderUserFeedback() *PortingOrderUserFeedback`

NewPortingOrderUserFeedback instantiates a new PortingOrderUserFeedback object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPortingOrderUserFeedbackWithDefaults

`func NewPortingOrderUserFeedbackWithDefaults() *PortingOrderUserFeedback`

NewPortingOrderUserFeedbackWithDefaults instantiates a new PortingOrderUserFeedback object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserRating

`func (o *PortingOrderUserFeedback) GetUserRating() int32`

GetUserRating returns the UserRating field if non-nil, zero value otherwise.

### GetUserRatingOk

`func (o *PortingOrderUserFeedback) GetUserRatingOk() (*int32, bool)`

GetUserRatingOk returns a tuple with the UserRating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRating

`func (o *PortingOrderUserFeedback) SetUserRating(v int32)`

SetUserRating sets UserRating field to given value.

### HasUserRating

`func (o *PortingOrderUserFeedback) HasUserRating() bool`

HasUserRating returns a boolean if a field has been set.

### GetUserComment

`func (o *PortingOrderUserFeedback) GetUserComment() string`

GetUserComment returns the UserComment field if non-nil, zero value otherwise.

### GetUserCommentOk

`func (o *PortingOrderUserFeedback) GetUserCommentOk() (*string, bool)`

GetUserCommentOk returns a tuple with the UserComment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserComment

`func (o *PortingOrderUserFeedback) SetUserComment(v string)`

SetUserComment sets UserComment field to given value.

### HasUserComment

`func (o *PortingOrderUserFeedback) HasUserComment() bool`

HasUserComment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


