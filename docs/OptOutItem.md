# OptOutItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | Sending address (+E.164 formatted phone number, alphanumeric sender ID, or short code). | [optional] 
**To** | Pointer to **string** | Receiving address (+E.164 formatted phone number or short code). | [optional] 
**MessagingProfileId** | Pointer to **NullableString** | Unique identifier for a messaging profile. | [optional] 
**Keyword** | Pointer to **NullableString** | The keyword that triggered the opt-out. | [optional] 
**CreatedAt** | Pointer to **time.Time** | The timestamp when the opt-out was created | [optional] 

## Methods

### NewOptOutItem

`func NewOptOutItem() *OptOutItem`

NewOptOutItem instantiates a new OptOutItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptOutItemWithDefaults

`func NewOptOutItemWithDefaults() *OptOutItem`

NewOptOutItemWithDefaults instantiates a new OptOutItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *OptOutItem) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *OptOutItem) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *OptOutItem) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *OptOutItem) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetTo

`func (o *OptOutItem) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *OptOutItem) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *OptOutItem) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *OptOutItem) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetMessagingProfileId

`func (o *OptOutItem) GetMessagingProfileId() string`

GetMessagingProfileId returns the MessagingProfileId field if non-nil, zero value otherwise.

### GetMessagingProfileIdOk

`func (o *OptOutItem) GetMessagingProfileIdOk() (*string, bool)`

GetMessagingProfileIdOk returns a tuple with the MessagingProfileId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessagingProfileId

`func (o *OptOutItem) SetMessagingProfileId(v string)`

SetMessagingProfileId sets MessagingProfileId field to given value.

### HasMessagingProfileId

`func (o *OptOutItem) HasMessagingProfileId() bool`

HasMessagingProfileId returns a boolean if a field has been set.

### SetMessagingProfileIdNil

`func (o *OptOutItem) SetMessagingProfileIdNil(b bool)`

 SetMessagingProfileIdNil sets the value for MessagingProfileId to be an explicit nil

### UnsetMessagingProfileId
`func (o *OptOutItem) UnsetMessagingProfileId()`

UnsetMessagingProfileId ensures that no value is present for MessagingProfileId, not even an explicit nil
### GetKeyword

`func (o *OptOutItem) GetKeyword() string`

GetKeyword returns the Keyword field if non-nil, zero value otherwise.

### GetKeywordOk

`func (o *OptOutItem) GetKeywordOk() (*string, bool)`

GetKeywordOk returns a tuple with the Keyword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyword

`func (o *OptOutItem) SetKeyword(v string)`

SetKeyword sets Keyword field to given value.

### HasKeyword

`func (o *OptOutItem) HasKeyword() bool`

HasKeyword returns a boolean if a field has been set.

### SetKeywordNil

`func (o *OptOutItem) SetKeywordNil(b bool)`

 SetKeywordNil sets the value for Keyword to be an explicit nil

### UnsetKeyword
`func (o *OptOutItem) UnsetKeyword()`

UnsetKeyword ensures that no value is present for Keyword, not even an explicit nil
### GetCreatedAt

`func (o *OptOutItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *OptOutItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *OptOutItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *OptOutItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


