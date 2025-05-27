# RCSResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** |  | [optional] 
**Direction** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | message ID | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**OrganizationId** | Pointer to **string** |  | [optional] 
**MessagingProfileId** | Pointer to **string** |  | [optional] 
**From** | Pointer to [**RCSFrom**](RCSFrom.md) |  | [optional] 
**To** | Pointer to [**[]RCSToItem**](RCSToItem.md) |  | [optional] 
**Body** | Pointer to [**RCSAgentMessage**](RCSAgentMessage.md) |  | [optional] 
**Encoding** | Pointer to **string** |  | [optional] 
**ReceivedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewRCSResponseData

`func NewRCSResponseData() *RCSResponseData`

NewRCSResponseData instantiates a new RCSResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRCSResponseDataWithDefaults

`func NewRCSResponseDataWithDefaults() *RCSResponseData`

NewRCSResponseDataWithDefaults instantiates a new RCSResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *RCSResponseData) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *RCSResponseData) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *RCSResponseData) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *RCSResponseData) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetDirection

`func (o *RCSResponseData) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *RCSResponseData) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *RCSResponseData) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *RCSResponseData) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetId

`func (o *RCSResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RCSResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RCSResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *RCSResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *RCSResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *RCSResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *RCSResponseData) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *RCSResponseData) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrganizationId

`func (o *RCSResponseData) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *RCSResponseData) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *RCSResponseData) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *RCSResponseData) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetMessagingProfileId

`func (o *RCSResponseData) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *RCSResponseData) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *RCSResponseData) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.

### HasMessagingProfileId

`func (o *RCSResponseData) HasMessagingProfileId() bool`

HasMessagingProfileId returns a boolean if a field has been set.

### GetFrom

`func (o *RCSResponseData) GetFrom() RCSFrom`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *RCSResponseData) GetFromOk() (*RCSFrom, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *RCSResponseData) SetFrom(v RCSFrom)`

SetFrom sets From field to given value.

### HasFrom

`func (o *RCSResponseData) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *RCSResponseData) GetTo() []RCSToItem`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *RCSResponseData) GetToOk() (*[]RCSToItem, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *RCSResponseData) SetTo(v []RCSToItem)`

SetTo sets To field to given value.

### HasTo

`func (o *RCSResponseData) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetBody

`func (o *RCSResponseData) GetBody() RCSAgentMessage`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *RCSResponseData) GetBodyOk() (*RCSAgentMessage, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *RCSResponseData) SetBody(v RCSAgentMessage)`

SetBody sets Body field to given value.

### HasBody

`func (o *RCSResponseData) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetEncoding

`func (o *RCSResponseData) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *RCSResponseData) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *RCSResponseData) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *RCSResponseData) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetReceivedAt

`func (o *RCSResponseData) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *RCSResponseData) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *RCSResponseData) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.

### HasReceivedAt

`func (o *RCSResponseData) HasReceivedAt() bool`

HasReceivedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


