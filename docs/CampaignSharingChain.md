# CampaignSharingChain

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SharedByMe** | Pointer to [**CampaignSharingStatus**](CampaignSharingStatus.md) |  | [optional] 
**SharedWithMe** | Pointer to [**CampaignSharingStatus**](CampaignSharingStatus.md) |  | [optional] 

## Methods

### NewCampaignSharingChain

`func NewCampaignSharingChain() *CampaignSharingChain`

NewCampaignSharingChain instantiates a new CampaignSharingChain object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignSharingChainWithDefaults

`func NewCampaignSharingChainWithDefaults() *CampaignSharingChain`

NewCampaignSharingChainWithDefaults instantiates a new CampaignSharingChain object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSharedByMe

`func (o *CampaignSharingChain) GetSharedByMe() CampaignSharingStatus`

GetSharedByMe returns the SharedByMe field if non-nil, zero value otherwise.

### GetSharedByMeOk

`func (o *CampaignSharingChain) GetSharedByMeOk() (*CampaignSharingStatus, bool)`

GetSharedByMeOk returns a tuple with the SharedByMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedByMe

`func (o *CampaignSharingChain) SetSharedByMe(v CampaignSharingStatus)`

SetSharedByMe sets SharedByMe field to given value.

### HasSharedByMe

`func (o *CampaignSharingChain) HasSharedByMe() bool`

HasSharedByMe returns a boolean if a field has been set.

### GetSharedWithMe

`func (o *CampaignSharingChain) GetSharedWithMe() CampaignSharingStatus`

GetSharedWithMe returns the SharedWithMe field if non-nil, zero value otherwise.

### GetSharedWithMeOk

`func (o *CampaignSharingChain) GetSharedWithMeOk() (*CampaignSharingStatus, bool)`

GetSharedWithMeOk returns a tuple with the SharedWithMe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedWithMe

`func (o *CampaignSharingChain) SetSharedWithMe(v CampaignSharingStatus)`

SetSharedWithMe sets SharedWithMe field to given value.

### HasSharedWithMe

`func (o *CampaignSharingChain) HasSharedWithMe() bool`

HasSharedWithMe returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


