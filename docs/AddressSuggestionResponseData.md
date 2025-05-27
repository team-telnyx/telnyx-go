# AddressSuggestionResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accepted** | Pointer to **bool** | Indicates if the address suggestions are accepted. | [optional] 
**Id** | Pointer to **string** | The UUID of the location. | [optional] 
**RecordType** | Pointer to **string** |  | [optional] 

## Methods

### NewAddressSuggestionResponseData

`func NewAddressSuggestionResponseData() *AddressSuggestionResponseData`

NewAddressSuggestionResponseData instantiates a new AddressSuggestionResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressSuggestionResponseDataWithDefaults

`func NewAddressSuggestionResponseDataWithDefaults() *AddressSuggestionResponseData`

NewAddressSuggestionResponseDataWithDefaults instantiates a new AddressSuggestionResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccepted

`func (o *AddressSuggestionResponseData) GetAccepted() bool`

GetAccepted returns the Accepted field if non-nil, zero value otherwise.

### GetAcceptedOk

`func (o *AddressSuggestionResponseData) GetAcceptedOk() (*bool, bool)`

GetAcceptedOk returns a tuple with the Accepted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccepted

`func (o *AddressSuggestionResponseData) SetAccepted(v bool)`

SetAccepted sets Accepted field to given value.

### HasAccepted

`func (o *AddressSuggestionResponseData) HasAccepted() bool`

HasAccepted returns a boolean if a field has been set.

### GetId

`func (o *AddressSuggestionResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddressSuggestionResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddressSuggestionResponseData) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AddressSuggestionResponseData) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRecordType

`func (o *AddressSuggestionResponseData) GetRecordType() string`

GetRecordType returns the RecordType field if non-nil, zero value otherwise.

### GetRecordTypeOk

`func (o *AddressSuggestionResponseData) GetRecordTypeOk() (*string, bool)`

GetRecordTypeOk returns a tuple with the RecordType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordType

`func (o *AddressSuggestionResponseData) SetRecordType(v string)`

SetRecordType sets RecordType field to given value.

### HasRecordType

`func (o *AddressSuggestionResponseData) HasRecordType() bool`

HasRecordType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


