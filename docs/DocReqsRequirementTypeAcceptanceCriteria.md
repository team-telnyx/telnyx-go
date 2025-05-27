# DocReqsRequirementTypeAcceptanceCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TimeLimit** | Pointer to **string** | Specifies time-based acceptance criteria | [optional] 
**LocalityLimit** | Pointer to **string** | Specifies geography-based acceptance criteria | [optional] 
**AcceptableValues** | Pointer to **[]string** | Specifies the list of strictly possible values for the requirement. Ignored when empty | [optional] 
**MaxLength** | Pointer to **int32** | Maximum length allowed for the value | [optional] 
**MinLength** | Pointer to **int32** | Minimum length allowed for the value | [optional] 
**AcceptableCharacters** | Pointer to **string** | Specifies the allowed characters as a string | [optional] 
**CaseSensitive** | Pointer to **bool** | Specifies whether string matching should be case sensitive | [optional] 
**Regex** | Pointer to **string** | A regular expression pattern that the value must match | [optional] 

## Methods

### NewDocReqsRequirementTypeAcceptanceCriteria

`func NewDocReqsRequirementTypeAcceptanceCriteria() *DocReqsRequirementTypeAcceptanceCriteria`

NewDocReqsRequirementTypeAcceptanceCriteria instantiates a new DocReqsRequirementTypeAcceptanceCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocReqsRequirementTypeAcceptanceCriteriaWithDefaults

`func NewDocReqsRequirementTypeAcceptanceCriteriaWithDefaults() *DocReqsRequirementTypeAcceptanceCriteria`

NewDocReqsRequirementTypeAcceptanceCriteriaWithDefaults instantiates a new DocReqsRequirementTypeAcceptanceCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeLimit

`func (o *DocReqsRequirementTypeAcceptanceCriteria) GetTimeLimit() string`

GetTimeLimit returns the TimeLimit field if non-nil, zero value otherwise.

### GetTimeLimitOk

`func (o *DocReqsRequirementTypeAcceptanceCriteria) GetTimeLimitOk() (*string, bool)`

GetTimeLimitOk returns a tuple with the TimeLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeLimit

`func (o *DocReqsRequirementTypeAcceptanceCriteria) SetTimeLimit(v string)`

SetTimeLimit sets TimeLimit field to given value.

### HasTimeLimit

`func (o *DocReqsRequirementTypeAcceptanceCriteria) HasTimeLimit() bool`

HasTimeLimit returns a boolean if a field has been set.

### GetLocalityLimit

`func (o *DocReqsRequirementTypeAcceptanceCriteria) GetLocalityLimit() string`

GetLocalityLimit returns the LocalityLimit field if non-nil, zero value otherwise.

### GetLocalityLimitOk

`func (o *DocReqsRequirementTypeAcceptanceCriteria) GetLocalityLimitOk() (*string, bool)`

GetLocalityLimitOk returns a tuple with the LocalityLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalityLimit

`func (o *DocReqsRequirementTypeAcceptanceCriteria) SetLocalityLimit(v string)`

SetLocalityLimit sets LocalityLimit field to given value.

### HasLocalityLimit

`func (o *DocReqsRequirementTypeAcceptanceCriteria) HasLocalityLimit() bool`

HasLocalityLimit returns a boolean if a field has been set.

### GetAcceptableValues

`func (o *DocReqsRequirementTypeAcceptanceCriteria) GetAcceptableValues() []string`

GetAcceptableValues returns the AcceptableValues field if non-nil, zero value otherwise.

### GetAcceptableValuesOk

`func (o *DocReqsRequirementTypeAcceptanceCriteria) GetAcceptableValuesOk() (*[]string, bool)`

GetAcceptableValuesOk returns a tuple with the AcceptableValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptableValues

`func (o *DocReqsRequirementTypeAcceptanceCriteria) SetAcceptableValues(v []string)`

SetAcceptableValues sets AcceptableValues field to given value.

### HasAcceptableValues

`func (o *DocReqsRequirementTypeAcceptanceCriteria) HasAcceptableValues() bool`

HasAcceptableValues returns a boolean if a field has been set.

### GetMaxLength

`func (o *DocReqsRequirementTypeAcceptanceCriteria) GetMaxLength() int32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *DocReqsRequirementTypeAcceptanceCriteria) GetMaxLengthOk() (*int32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *DocReqsRequirementTypeAcceptanceCriteria) SetMaxLength(v int32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *DocReqsRequirementTypeAcceptanceCriteria) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### GetMinLength

`func (o *DocReqsRequirementTypeAcceptanceCriteria) GetMinLength() int32`

GetMinLength returns the MinLength field if non-nil, zero value otherwise.

### GetMinLengthOk

`func (o *DocReqsRequirementTypeAcceptanceCriteria) GetMinLengthOk() (*int32, bool)`

GetMinLengthOk returns a tuple with the MinLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLength

`func (o *DocReqsRequirementTypeAcceptanceCriteria) SetMinLength(v int32)`

SetMinLength sets MinLength field to given value.

### HasMinLength

`func (o *DocReqsRequirementTypeAcceptanceCriteria) HasMinLength() bool`

HasMinLength returns a boolean if a field has been set.

### GetAcceptableCharacters

`func (o *DocReqsRequirementTypeAcceptanceCriteria) GetAcceptableCharacters() string`

GetAcceptableCharacters returns the AcceptableCharacters field if non-nil, zero value otherwise.

### GetAcceptableCharactersOk

`func (o *DocReqsRequirementTypeAcceptanceCriteria) GetAcceptableCharactersOk() (*string, bool)`

GetAcceptableCharactersOk returns a tuple with the AcceptableCharacters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptableCharacters

`func (o *DocReqsRequirementTypeAcceptanceCriteria) SetAcceptableCharacters(v string)`

SetAcceptableCharacters sets AcceptableCharacters field to given value.

### HasAcceptableCharacters

`func (o *DocReqsRequirementTypeAcceptanceCriteria) HasAcceptableCharacters() bool`

HasAcceptableCharacters returns a boolean if a field has been set.

### GetCaseSensitive

`func (o *DocReqsRequirementTypeAcceptanceCriteria) GetCaseSensitive() bool`

GetCaseSensitive returns the CaseSensitive field if non-nil, zero value otherwise.

### GetCaseSensitiveOk

`func (o *DocReqsRequirementTypeAcceptanceCriteria) GetCaseSensitiveOk() (*bool, bool)`

GetCaseSensitiveOk returns a tuple with the CaseSensitive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaseSensitive

`func (o *DocReqsRequirementTypeAcceptanceCriteria) SetCaseSensitive(v bool)`

SetCaseSensitive sets CaseSensitive field to given value.

### HasCaseSensitive

`func (o *DocReqsRequirementTypeAcceptanceCriteria) HasCaseSensitive() bool`

HasCaseSensitive returns a boolean if a field has been set.

### GetRegex

`func (o *DocReqsRequirementTypeAcceptanceCriteria) GetRegex() string`

GetRegex returns the Regex field if non-nil, zero value otherwise.

### GetRegexOk

`func (o *DocReqsRequirementTypeAcceptanceCriteria) GetRegexOk() (*string, bool)`

GetRegexOk returns a tuple with the Regex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegex

`func (o *DocReqsRequirementTypeAcceptanceCriteria) SetRegex(v string)`

SetRegex sets Regex field to given value.

### HasRegex

`func (o *DocReqsRequirementTypeAcceptanceCriteria) HasRegex() bool`

HasRegex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


