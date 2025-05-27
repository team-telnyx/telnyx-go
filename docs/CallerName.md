# CallerName

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallerName** | Pointer to **string** | The name of the requested phone number&#39;s owner as per the CNAM database | [optional] 
**ErrorCode** | Pointer to **string** | A caller-name lookup specific error code, expressed as a stringified 5-digit integer | [optional] 

## Methods

### NewCallerName

`func NewCallerName() *CallerName`

NewCallerName instantiates a new CallerName object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallerNameWithDefaults

`func NewCallerNameWithDefaults() *CallerName`

NewCallerNameWithDefaults instantiates a new CallerName object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallerName

`func (o *CallerName) GetCallerName() string`

GetCallerName returns the CallerName field if non-nil, zero value otherwise.

### GetCallerNameOk

`func (o *CallerName) GetCallerNameOk() (*string, bool)`

GetCallerNameOk returns a tuple with the CallerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerName

`func (o *CallerName) SetCallerName(v string)`

SetCallerName sets CallerName field to given value.

### HasCallerName

`func (o *CallerName) HasCallerName() bool`

HasCallerName returns a boolean if a field has been set.

### GetErrorCode

`func (o *CallerName) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *CallerName) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *CallerName) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *CallerName) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


