# SharedCampaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BrandId** | **string** | Alphanumeric identifier of the brand associated with this campaign. | 
**CampaignId** | **string** | Alphanumeric identifier assigned by the registry for a campaign. This identifier is required by the NetNumber OSR SMS enabling process of 10DLC. | 
**CreateDate** | Pointer to **string** | Unix timestamp when campaign was created. | [optional] 
**Status** | Pointer to **string** | Current campaign status. Possible values: ACTIVE, EXPIRED. A newly created campaign defaults to ACTIVE status.  | [optional] 
**Usecase** | **string** | Campaign usecase. Must be of defined valid types. Use &#x60;/registry/enum/usecase&#x60; operation to retrieve usecases available for given brand. | 

## Methods

### NewSharedCampaign

`func NewSharedCampaign(brandId string, campaignId string, usecase string, ) *SharedCampaign`

NewSharedCampaign instantiates a new SharedCampaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSharedCampaignWithDefaults

`func NewSharedCampaignWithDefaults() *SharedCampaign`

NewSharedCampaignWithDefaults instantiates a new SharedCampaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBrandId

`func (o *SharedCampaign) GetBrandId() string`

GetBrandId returns the BrandId field if non-nil, zero value otherwise.

### GetBrandIdOk

`func (o *SharedCampaign) GetBrandIdOk() (*string, bool)`

GetBrandIdOk returns a tuple with the BrandId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandId

`func (o *SharedCampaign) SetBrandId(v string)`

SetBrandId sets BrandId field to given value.


### GetCampaignId

`func (o *SharedCampaign) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *SharedCampaign) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *SharedCampaign) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.


### GetCreateDate

`func (o *SharedCampaign) GetCreateDate() string`

GetCreateDate returns the CreateDate field if non-nil, zero value otherwise.

### GetCreateDateOk

`func (o *SharedCampaign) GetCreateDateOk() (*string, bool)`

GetCreateDateOk returns a tuple with the CreateDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateDate

`func (o *SharedCampaign) SetCreateDate(v string)`

SetCreateDate sets CreateDate field to given value.

### HasCreateDate

`func (o *SharedCampaign) HasCreateDate() bool`

HasCreateDate returns a boolean if a field has been set.

### GetStatus

`func (o *SharedCampaign) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SharedCampaign) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SharedCampaign) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SharedCampaign) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUsecase

`func (o *SharedCampaign) GetUsecase() string`

GetUsecase returns the Usecase field if non-nil, zero value otherwise.

### GetUsecaseOk

`func (o *SharedCampaign) GetUsecaseOk() (*string, bool)`

GetUsecaseOk returns a tuple with the Usecase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsecase

`func (o *SharedCampaign) SetUsecase(v string)`

SetUsecase sets Usecase field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


