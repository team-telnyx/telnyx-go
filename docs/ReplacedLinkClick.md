# ReplacedLinkClick

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordType** | Pointer to **string** | Identifies the type of the resource. | [optional] 
**Url** | Pointer to **string** | The original link that was sent in the message. | [optional] 
**To** | Pointer to **string** | Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short code). | [optional] 
**MessageId** | Pointer to **string** | The message ID associated with the clicked link. | [optional] 
**TimeClicked** | Pointer to **time.Time** | ISO 8601 formatted date indicating when the message request was received. | [optional] 

## Methods

### NewReplacedLinkClick

`func NewReplacedLinkClick() *ReplacedLinkClick`

NewReplacedLinkClick instantiates a new ReplacedLinkClick object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReplacedLinkClickWithDefaults

`func NewReplacedLinkClickWithDefaults() *ReplacedLinkClick`

NewReplacedLinkClickWithDefaults instantiates a new ReplacedLinkClick object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordType

`func (o *ReplacedLinkClick) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *ReplacedLinkClick) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *ReplacedLinkClick) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *ReplacedLinkClick) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.

### GetUrl

`func (o *ReplacedLinkClick) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ReplacedLinkClick) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ReplacedLinkClick) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ReplacedLinkClick) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetTo

`func (o *ReplacedLinkClick) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ReplacedLinkClick) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ReplacedLinkClick) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *ReplacedLinkClick) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetMessageId

`func (o *ReplacedLinkClick) GetMessageId() string`

GetMessageId returns the MessageId field if non-nil, zero value otherwise.

### GetMessageIdOk

`func (o *ReplacedLinkClick) GetMessageIdOk() (*string, bool)`

GetMessageIdOk returns a tuple with the MessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageId

`func (o *ReplacedLinkClick) SetMessageId(v string)`

SetMessageId sets MessageId field to given value.

### HasMessageId

`func (o *ReplacedLinkClick) HasMessageId() bool`

HasMessageId returns a boolean if a field has been set.

### GetTimeClicked

`func (o *ReplacedLinkClick) GetTimeClicked() time.Time`

GetTimeClicked returns the TimeClicked field if non-nil, zero value otherwise.

### GetTimeClickedOk

`func (o *ReplacedLinkClick) GetTimeClickedOk() (*time.Time, bool)`

GetTimeClickedOk returns a tuple with the TimeClicked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeClicked

`func (o *ReplacedLinkClick) SetTimeClicked(v time.Time)`

SetTimeClicked sets TimeClicked field to given value.

### HasTimeClicked

`func (o *ReplacedLinkClick) HasTimeClicked() bool`

HasTimeClicked returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


