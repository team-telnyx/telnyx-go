# RCSAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | Pointer to **string** | Text that is shown in the suggested action. Maximum 25 characters. | [optional] 
**PostbackData** | Pointer to **string** | Payload (base64 encoded) that will be sent to the agent in the user event that results when the user taps the suggested action. Maximum 2048 characters. | [optional] 
**FallbackUrl** | Pointer to **string** | Fallback URL to use if a client doesn&#39;t support a suggested action. Fallback URLs open in new browser windows. Maximum 2048 characters. | [optional] 
**DialAction** | Pointer to [**RCSDialAction**](RCSDialAction.md) |  | [optional] 
**ViewLocationAction** | Pointer to [**RCSViewLocationAction**](RCSViewLocationAction.md) |  | [optional] 
**CreateCalendarEventAction** | Pointer to [**RCSCreateCalendarEventAction**](RCSCreateCalendarEventAction.md) |  | [optional] 
**OpenUrlAction** | Pointer to [**RCSOpenUrlAction**](RCSOpenUrlAction.md) |  | [optional] 
**ShareLocationAction** | Pointer to **map[string]interface{}** | Opens the RCS app&#39;s location chooser so the user can pick a location to send back to the agent. | [optional] 
**ComposeAction** | Pointer to [**RCSComposeAction**](RCSComposeAction.md) |  | [optional] 

## Methods

### NewRCSAction

`func NewRCSAction() *RCSAction`

NewRCSAction instantiates a new RCSAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCSActionWithDefaults

`func NewRCSActionWithDefaults() *RCSAction`

NewRCSActionWithDefaults instantiates a new RCSAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *RCSAction) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *RCSAction) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *RCSAction) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *RCSAction) HasText() bool`

HasText returns a boolean if a field has been set.

### GetPostbackData

`func (o *RCSAction) GetPostbackData() string`

GetPostbackData returns the PostbackData field if non-nil, zero value otherwise.

### GetPostbackDataOk

`func (o *RCSAction) GetPostbackDataOk() (*string, bool)`

GetPostbackDataOk returns a tuple with the PostbackData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostbackData

`func (o *RCSAction) SetPostbackData(v string)`

SetPostbackData sets PostbackData field to given value.

### HasPostbackData

`func (o *RCSAction) HasPostbackData() bool`

HasPostbackData returns a boolean if a field has been set.

### GetFallbackUrl

`func (o *RCSAction) GetFallbackUrl() string`

GetFallbackUrl returns the FallbackUrl field if non-nil, zero value otherwise.

### GetFallbackUrlOk

`func (o *RCSAction) GetFallbackUrlOk() (*string, bool)`

GetFallbackUrlOk returns a tuple with the FallbackUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbackUrl

`func (o *RCSAction) SetFallbackUrl(v string)`

SetFallbackUrl sets FallbackUrl field to given value.

### HasFallbackUrl

`func (o *RCSAction) HasFallbackUrl() bool`

HasFallbackUrl returns a boolean if a field has been set.

### GetDialAction

`func (o *RCSAction) GetDialAction() RCSDialAction`

GetDialAction returns the DialAction field if non-nil, zero value otherwise.

### GetDialActionOk

`func (o *RCSAction) GetDialActionOk() (*RCSDialAction, bool)`

GetDialActionOk returns a tuple with the DialAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialAction

`func (o *RCSAction) SetDialAction(v RCSDialAction)`

SetDialAction sets DialAction field to given value.

### HasDialAction

`func (o *RCSAction) HasDialAction() bool`

HasDialAction returns a boolean if a field has been set.

### GetViewLocationAction

`func (o *RCSAction) GetViewLocationAction() RCSViewLocationAction`

GetViewLocationAction returns the ViewLocationAction field if non-nil, zero value otherwise.

### GetViewLocationActionOk

`func (o *RCSAction) GetViewLocationActionOk() (*RCSViewLocationAction, bool)`

GetViewLocationActionOk returns a tuple with the ViewLocationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewLocationAction

`func (o *RCSAction) SetViewLocationAction(v RCSViewLocationAction)`

SetViewLocationAction sets ViewLocationAction field to given value.

### HasViewLocationAction

`func (o *RCSAction) HasViewLocationAction() bool`

HasViewLocationAction returns a boolean if a field has been set.

### GetCreateCalendarEventAction

`func (o *RCSAction) GetCreateCalendarEventAction() RCSCreateCalendarEventAction`

GetCreateCalendarEventAction returns the CreateCalendarEventAction field if non-nil, zero value otherwise.

### GetCreateCalendarEventActionOk

`func (o *RCSAction) GetCreateCalendarEventActionOk() (*RCSCreateCalendarEventAction, bool)`

GetCreateCalendarEventActionOk returns a tuple with the CreateCalendarEventAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateCalendarEventAction

`func (o *RCSAction) SetCreateCalendarEventAction(v RCSCreateCalendarEventAction)`

SetCreateCalendarEventAction sets CreateCalendarEventAction field to given value.

### HasCreateCalendarEventAction

`func (o *RCSAction) HasCreateCalendarEventAction() bool`

HasCreateCalendarEventAction returns a boolean if a field has been set.

### GetOpenUrlAction

`func (o *RCSAction) GetOpenUrlAction() RCSOpenUrlAction`

GetOpenUrlAction returns the OpenUrlAction field if non-nil, zero value otherwise.

### GetOpenUrlActionOk

`func (o *RCSAction) GetOpenUrlActionOk() (*RCSOpenUrlAction, bool)`

GetOpenUrlActionOk returns a tuple with the OpenUrlAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenUrlAction

`func (o *RCSAction) SetOpenUrlAction(v RCSOpenUrlAction)`

SetOpenUrlAction sets OpenUrlAction field to given value.

### HasOpenUrlAction

`func (o *RCSAction) HasOpenUrlAction() bool`

HasOpenUrlAction returns a boolean if a field has been set.

### GetShareLocationAction

`func (o *RCSAction) GetShareLocationAction() map[string]interface{}`

GetShareLocationAction returns the ShareLocationAction field if non-nil, zero value otherwise.

### GetShareLocationActionOk

`func (o *RCSAction) GetShareLocationActionOk() (*map[string]interface{}, bool)`

GetShareLocationActionOk returns a tuple with the ShareLocationAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareLocationAction

`func (o *RCSAction) SetShareLocationAction(v map[string]interface{})`

SetShareLocationAction sets ShareLocationAction field to given value.

### HasShareLocationAction

`func (o *RCSAction) HasShareLocationAction() bool`

HasShareLocationAction returns a boolean if a field has been set.

### GetComposeAction

`func (o *RCSAction) GetComposeAction() RCSComposeAction`

GetComposeAction returns the ComposeAction field if non-nil, zero value otherwise.

### GetComposeActionOk

`func (o *RCSAction) GetComposeActionOk() (*RCSComposeAction, bool)`

GetComposeActionOk returns a tuple with the ComposeAction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComposeAction

`func (o *RCSAction) SetComposeAction(v RCSComposeAction)`

SetComposeAction sets ComposeAction field to given value.

### HasComposeAction

`func (o *RCSAction) HasComposeAction() bool`

HasComposeAction returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


