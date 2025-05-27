# RCSOpenUrlAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | **string** |  | 
**Application** | **string** | URL open application, browser or webview. | 
**WebviewViewMode** | **string** |  | 
**Description** | Pointer to **string** | Accessbility description for webview. | [optional] 

## Methods

### NewRCSOpenUrlAction

`func NewRCSOpenUrlAction(url string, application string, webviewViewMode string, ) *RCSOpenUrlAction`

NewRCSOpenUrlAction instantiates a new RCSOpenUrlAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCSOpenUrlActionWithDefaults

`func NewRCSOpenUrlActionWithDefaults() *RCSOpenUrlAction`

NewRCSOpenUrlActionWithDefaults instantiates a new RCSOpenUrlAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *RCSOpenUrlAction) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RCSOpenUrlAction) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RCSOpenUrlAction) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetApplication

`func (o *RCSOpenUrlAction) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *RCSOpenUrlAction) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *RCSOpenUrlAction) SetApplication(v string)`

SetApplication sets Application field to given value.


### GetWebviewViewMode

`func (o *RCSOpenUrlAction) GetWebviewViewMode() string`

GetWebviewViewMode returns the WebviewViewMode field if non-nil, zero value otherwise.

### GetWebviewViewModeOk

`func (o *RCSOpenUrlAction) GetWebviewViewModeOk() (*string, bool)`

GetWebviewViewModeOk returns a tuple with the WebviewViewMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebviewViewMode

`func (o *RCSOpenUrlAction) SetWebviewViewMode(v string)`

SetWebviewViewMode sets WebviewViewMode field to given value.


### GetDescription

`func (o *RCSOpenUrlAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *RCSOpenUrlAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *RCSOpenUrlAction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *RCSOpenUrlAction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


