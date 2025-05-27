# GetUserTags200ResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OutboundProfileTags** | Pointer to **[]string** | A list of your tags on the given resource type. NOTE: The casing of the tags returned will not necessarily match the casing of the tags when they were added to a resource. This is because the tags will have the casing of the first time they were used within the Telnyx system regardless of source. | [optional] 
**NumberTags** | Pointer to **[]string** | A list of your tags on the given resource type. NOTE: The casing of the tags returned will not necessarily match the casing of the tags when they were added to a resource. This is because the tags will have the casing of the first time they were used within the Telnyx system regardless of source. | [optional] 

## Methods

### NewGetUserTags200ResponseData

`func NewGetUserTags200ResponseData() *GetUserTags200ResponseData`

NewGetUserTags200ResponseData instantiates a new GetUserTags200ResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetUserTags200ResponseDataWithDefaults

`func NewGetUserTags200ResponseDataWithDefaults() *GetUserTags200ResponseData`

NewGetUserTags200ResponseDataWithDefaults instantiates a new GetUserTags200ResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutboundProfileTags

`func (o *GetUserTags200ResponseData) GetOutboundProfileTags() []string`

GetOutboundProfileTags returns the OutboundProfileTags field if non-nil, zero value otherwise.

### GetOutboundProfileTagsOk

`func (o *GetUserTags200ResponseData) GetOutboundProfileTagsOk() (*[]string, bool)`

GetOutboundProfileTagsOk returns a tuple with the OutboundProfileTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundProfileTags

`func (o *GetUserTags200ResponseData) SetOutboundProfileTags(v []string)`

SetOutboundProfileTags sets OutboundProfileTags field to given value.

### HasOutboundProfileTags

`func (o *GetUserTags200ResponseData) HasOutboundProfileTags() bool`

HasOutboundProfileTags returns a boolean if a field has been set.

### GetNumberTags

`func (o *GetUserTags200ResponseData) GetNumberTags() []string`

GetNumberTags returns the NumberTags field if non-nil, zero value otherwise.

### GetNumberTagsOk

`func (o *GetUserTags200ResponseData) GetNumberTagsOk() (*[]string, bool)`

GetNumberTagsOk returns a tuple with the NumberTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberTags

`func (o *GetUserTags200ResponseData) SetNumberTags(v []string)`

SetNumberTags sets NumberTags field to given value.

### HasNumberTags

`func (o *GetUserTags200ResponseData) HasNumberTags() bool`

HasNumberTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


