# FaxSendingStartedPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConnectionId** | Pointer to **string** | The ID of the connection used to send the fax. | [optional] 
**Direction** | Pointer to [**Direction**](Direction.md) |  | [optional] 
**FaxId** | Pointer to **string** | Identifies the fax. | [optional] 
**OriginalMediaUrl** | Pointer to **string** | The original URL to the PDF used for the fax&#39;s media. If media_name was supplied, this is omitted | [optional] 
**MediaName** | Pointer to **string** | The media_name used for the fax&#39;s media. Must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. media_name and media_url/contents can&#39;t be submitted together. | [optional] 
**To** | Pointer to **string** | The phone number, in E.164 format, the fax will be sent to or SIP URI | [optional] 
**From** | Pointer to **string** | The phone number, in E.164 format, the fax will be sent from. | [optional] 
**UserId** | Pointer to **string** | Identifier of the user to whom the fax belongs | [optional] 
**Status** | Pointer to **string** | The status of the fax. | [optional] 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 

## Methods

### NewFaxSendingStartedPayload

`func NewFaxSendingStartedPayload() *FaxSendingStartedPayload`

NewFaxSendingStartedPayload instantiates a new FaxSendingStartedPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaxSendingStartedPayloadWithDefaults

`func NewFaxSendingStartedPayloadWithDefaults() *FaxSendingStartedPayload`

NewFaxSendingStartedPayloadWithDefaults instantiates a new FaxSendingStartedPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectionId

`func (o *FaxSendingStartedPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *FaxSendingStartedPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *FaxSendingStartedPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *FaxSendingStartedPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetDirection

`func (o *FaxSendingStartedPayload) GetDirection() Direction`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FaxSendingStartedPayload) GetDirectionOk() (*Direction, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FaxSendingStartedPayload) SetDirection(v Direction)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *FaxSendingStartedPayload) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetFaxId

`func (o *FaxSendingStartedPayload) GetFaxId() string`

GetFaxId returns the FaxId field if non-nil, zero value otherwise.

### GetFaxIdOk

`func (o *FaxSendingStartedPayload) GetFaxIdOk() (*string, bool)`

GetFaxIdOk returns a tuple with the FaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxId

`func (o *FaxSendingStartedPayload) SetFaxId(v string)`

SetFaxId sets FaxId field to given value.

### HasFaxId

`func (o *FaxSendingStartedPayload) HasFaxId() bool`

HasFaxId returns a boolean if a field has been set.

### GetOriginalMediaUrl

`func (o *FaxSendingStartedPayload) GetOriginalMediaUrl() string`

GetOriginalMediaUrl returns the OriginalMediaUrl field if non-nil, zero value otherwise.

### GetOriginalMediaUrlOk

`func (o *FaxSendingStartedPayload) GetOriginalMediaUrlOk() (*string, bool)`

GetOriginalMediaUrlOk returns a tuple with the OriginalMediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalMediaUrl

`func (o *FaxSendingStartedPayload) SetOriginalMediaUrl(v string)`

SetOriginalMediaUrl sets OriginalMediaUrl field to given value.

### HasOriginalMediaUrl

`func (o *FaxSendingStartedPayload) HasOriginalMediaUrl() bool`

HasOriginalMediaUrl returns a boolean if a field has been set.

### GetMediaName

`func (o *FaxSendingStartedPayload) GetMediaName() string`

GetMediaName returns the MediaName field if non-nil, zero value otherwise.

### GetMediaNameOk

`func (o *FaxSendingStartedPayload) GetMediaNameOk() (*string, bool)`

GetMediaNameOk returns a tuple with the MediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaName

`func (o *FaxSendingStartedPayload) SetMediaName(v string)`

SetMediaName sets MediaName field to given value.

### HasMediaName

`func (o *FaxSendingStartedPayload) HasMediaName() bool`

HasMediaName returns a boolean if a field has been set.

### GetTo

`func (o *FaxSendingStartedPayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *FaxSendingStartedPayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *FaxSendingStartedPayload) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *FaxSendingStartedPayload) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetFrom

`func (o *FaxSendingStartedPayload) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *FaxSendingStartedPayload) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *FaxSendingStartedPayload) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *FaxSendingStartedPayload) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetUserId

`func (o *FaxSendingStartedPayload) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FaxSendingStartedPayload) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FaxSendingStartedPayload) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FaxSendingStartedPayload) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetStatus

`func (o *FaxSendingStartedPayload) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FaxSendingStartedPayload) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FaxSendingStartedPayload) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FaxSendingStartedPayload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetClientState

`func (o *FaxSendingStartedPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *FaxSendingStartedPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *FaxSendingStartedPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *FaxSendingStartedPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


