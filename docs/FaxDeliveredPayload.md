# FaxDeliveredPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallDurationSecs** | Pointer to **int32** | The duration of the call in seconds. | [optional] 
**ConnectionId** | Pointer to **string** | The ID of the connection used to send the fax. | [optional] 
**Direction** | Pointer to [**Direction**](Direction.md) |  | [optional] 
**FaxId** | Pointer to **string** | Identifies the fax. | [optional] 
**OriginalMediaUrl** | Pointer to **string** | The original URL to the PDF used for the fax&#39;s media. If media_name was supplied, this is omitted | [optional] 
**MediaName** | Pointer to **string** | The media_name used for the fax&#39;s media. Must point to a file previously uploaded to api.telnyx.com/v2/media by the same user/organization. media_name and media_url/contents can&#39;t be submitted together. | [optional] 
**To** | Pointer to **string** | The phone number, in E.164 format, the fax will be sent to or SIP URI | [optional] 
**From** | Pointer to **string** | The phone number, in E.164 format, the fax will be sent from. | [optional] 
**UserId** | Pointer to **string** | Identifier of the user to whom the fax belongs | [optional] 
**PageCount** | Pointer to **int32** | Number of transferred pages | [optional] 
**Status** | Pointer to **string** | The status of the fax. | [optional] 
**ClientState** | Pointer to **string** | State received from a command. | [optional] 

## Methods

### NewFaxDeliveredPayload

`func NewFaxDeliveredPayload() *FaxDeliveredPayload`

NewFaxDeliveredPayload instantiates a new FaxDeliveredPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFaxDeliveredPayloadWithDefaults

`func NewFaxDeliveredPayloadWithDefaults() *FaxDeliveredPayload`

NewFaxDeliveredPayloadWithDefaults instantiates a new FaxDeliveredPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallDurationSecs

`func (o *FaxDeliveredPayload) GetCallDurationSecs() int32`

GetCallDurationSecs returns the CallDurationSecs field if non-nil, zero value otherwise.

### GetCallDurationSecsOk

`func (o *FaxDeliveredPayload) GetCallDurationSecsOk() (*int32, bool)`

GetCallDurationSecsOk returns a tuple with the CallDurationSecs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallDurationSecs

`func (o *FaxDeliveredPayload) SetCallDurationSecs(v int32)`

SetCallDurationSecs sets CallDurationSecs field to given value.

### HasCallDurationSecs

`func (o *FaxDeliveredPayload) HasCallDurationSecs() bool`

HasCallDurationSecs returns a boolean if a field has been set.

### GetConnectionId

`func (o *FaxDeliveredPayload) GetConnectionId() string`

GetConnectionId returns the ConnectionId field if non-nil, zero value otherwise.

### GetConnectionIdOk

`func (o *FaxDeliveredPayload) GetConnectionIdOk() (*string, bool)`

GetConnectionIdOk returns a tuple with the ConnectionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectionId

`func (o *FaxDeliveredPayload) SetConnectionId(v string)`

SetConnectionId sets ConnectionId field to given value.

### HasConnectionId

`func (o *FaxDeliveredPayload) HasConnectionId() bool`

HasConnectionId returns a boolean if a field has been set.

### GetDirection

`func (o *FaxDeliveredPayload) GetDirection() Direction`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *FaxDeliveredPayload) GetDirectionOk() (*Direction, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *FaxDeliveredPayload) SetDirection(v Direction)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *FaxDeliveredPayload) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetFaxId

`func (o *FaxDeliveredPayload) GetFaxId() string`

GetFaxId returns the FaxId field if non-nil, zero value otherwise.

### GetFaxIdOk

`func (o *FaxDeliveredPayload) GetFaxIdOk() (*string, bool)`

GetFaxIdOk returns a tuple with the FaxId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaxId

`func (o *FaxDeliveredPayload) SetFaxId(v string)`

SetFaxId sets FaxId field to given value.

### HasFaxId

`func (o *FaxDeliveredPayload) HasFaxId() bool`

HasFaxId returns a boolean if a field has been set.

### GetOriginalMediaUrl

`func (o *FaxDeliveredPayload) GetOriginalMediaUrl() string`

GetOriginalMediaUrl returns the OriginalMediaUrl field if non-nil, zero value otherwise.

### GetOriginalMediaUrlOk

`func (o *FaxDeliveredPayload) GetOriginalMediaUrlOk() (*string, bool)`

GetOriginalMediaUrlOk returns a tuple with the OriginalMediaUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalMediaUrl

`func (o *FaxDeliveredPayload) SetOriginalMediaUrl(v string)`

SetOriginalMediaUrl sets OriginalMediaUrl field to given value.

### HasOriginalMediaUrl

`func (o *FaxDeliveredPayload) HasOriginalMediaUrl() bool`

HasOriginalMediaUrl returns a boolean if a field has been set.

### GetMediaName

`func (o *FaxDeliveredPayload) GetMediaName() string`

GetMediaName returns the MediaName field if non-nil, zero value otherwise.

### GetMediaNameOk

`func (o *FaxDeliveredPayload) GetMediaNameOk() (*string, bool)`

GetMediaNameOk returns a tuple with the MediaName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaName

`func (o *FaxDeliveredPayload) SetMediaName(v string)`

SetMediaName sets MediaName field to given value.

### HasMediaName

`func (o *FaxDeliveredPayload) HasMediaName() bool`

HasMediaName returns a boolean if a field has been set.

### GetTo

`func (o *FaxDeliveredPayload) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *FaxDeliveredPayload) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *FaxDeliveredPayload) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *FaxDeliveredPayload) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetFrom

`func (o *FaxDeliveredPayload) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *FaxDeliveredPayload) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *FaxDeliveredPayload) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *FaxDeliveredPayload) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetUserId

`func (o *FaxDeliveredPayload) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *FaxDeliveredPayload) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *FaxDeliveredPayload) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *FaxDeliveredPayload) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetPageCount

`func (o *FaxDeliveredPayload) GetPageCount() int32`

GetPageCount returns the PageCount field if non-nil, zero value otherwise.

### GetPageCountOk

`func (o *FaxDeliveredPayload) GetPageCountOk() (*int32, bool)`

GetPageCountOk returns a tuple with the PageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPageCount

`func (o *FaxDeliveredPayload) SetPageCount(v int32)`

SetPageCount sets PageCount field to given value.

### HasPageCount

`func (o *FaxDeliveredPayload) HasPageCount() bool`

HasPageCount returns a boolean if a field has been set.

### GetStatus

`func (o *FaxDeliveredPayload) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FaxDeliveredPayload) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FaxDeliveredPayload) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FaxDeliveredPayload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetClientState

`func (o *FaxDeliveredPayload) GetClientState() string`

GetClientState returns the ClientState field if non-nil, zero value otherwise.

### GetClientStateOk

`func (o *FaxDeliveredPayload) GetClientStateOk() (*string, bool)`

GetClientStateOk returns a tuple with the ClientState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientState

`func (o *FaxDeliveredPayload) SetClientState(v string)`

SetClientState sets ClientState field to given value.

### HasClientState

`func (o *FaxDeliveredPayload) HasClientState() bool`

HasClientState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


