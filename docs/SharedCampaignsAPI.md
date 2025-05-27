# \SharedCampaignsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPartnerCampaignSharingStatus**](SharedCampaignsAPI.md#GetPartnerCampaignSharingStatus) | **Get** /partnerCampaign/{campaignId}/sharing | Get Sharing Status
[**GetPartnerCampaignsSharedByUser**](SharedCampaignsAPI.md#GetPartnerCampaignsSharedByUser) | **Get** /partnerCampaign/sharedByMe | Get Partner Campaigns Shared By User
[**GetSharedCampaign**](SharedCampaignsAPI.md#GetSharedCampaign) | **Get** /partner_campaigns/{campaignId} | Get Single Shared Campaign
[**GetSharedCampaigns**](SharedCampaignsAPI.md#GetSharedCampaigns) | **Get** /partner_campaigns | List Shared Campaigns
[**UpdateSharedCampaign**](SharedCampaignsAPI.md#UpdateSharedCampaign) | **Patch** /partner_campaigns/{campaignId} | Update Single Shared Campaign



## GetPartnerCampaignSharingStatus

> map[string]CampaignSharingStatus GetPartnerCampaignSharingStatus(ctx, campaignId).Execute()

Get Sharing Status

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	campaignId := "campaignId_example" // string | ID of the campaign in question

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedCampaignsAPI.GetPartnerCampaignSharingStatus(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedCampaignsAPI.GetPartnerCampaignSharingStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPartnerCampaignSharingStatus`: map[string]CampaignSharingStatus
	fmt.Fprintf(os.Stdout, "Response from `SharedCampaignsAPI.GetPartnerCampaignSharingStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | ID of the campaign in question | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPartnerCampaignSharingStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**map[string]CampaignSharingStatus**](CampaignSharingStatus.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPartnerCampaignsSharedByUser

> SharedCampaignRecordSet GetPartnerCampaignsSharedByUser(ctx).Page(page).RecordsPerPage(recordsPerPage).Execute()

Get Partner Campaigns Shared By User



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	page := int32(56) // int32 | The 1-indexed page number to get. The default value is `1`. (optional) (default to 1)
	recordsPerPage := int32(56) // int32 | The amount of records per page, limited to between 1 and 500 inclusive. The default value is `10`. (optional) (default to 10)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedCampaignsAPI.GetPartnerCampaignsSharedByUser(context.Background()).Page(page).RecordsPerPage(recordsPerPage).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedCampaignsAPI.GetPartnerCampaignsSharedByUser``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPartnerCampaignsSharedByUser`: SharedCampaignRecordSet
	fmt.Fprintf(os.Stdout, "Response from `SharedCampaignsAPI.GetPartnerCampaignsSharedByUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPartnerCampaignsSharedByUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The 1-indexed page number to get. The default value is &#x60;1&#x60;. | [default to 1]
 **recordsPerPage** | **int32** | The amount of records per page, limited to between 1 and 500 inclusive. The default value is &#x60;10&#x60;. | [default to 10]

### Return type

[**SharedCampaignRecordSet**](SharedCampaignRecordSet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSharedCampaign

> TelnyxDownstreamCampaign GetSharedCampaign(ctx, campaignId).Execute()

Get Single Shared Campaign



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	campaignId := "campaignId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedCampaignsAPI.GetSharedCampaign(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedCampaignsAPI.GetSharedCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSharedCampaign`: TelnyxDownstreamCampaign
	fmt.Fprintf(os.Stdout, "Response from `SharedCampaignsAPI.GetSharedCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSharedCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TelnyxDownstreamCampaign**](TelnyxDownstreamCampaign.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSharedCampaigns

> TelnyxDownstreamCampaignRecordSet GetSharedCampaigns(ctx).Page(page).RecordsPerPage(recordsPerPage).Sort(sort).Execute()

List Shared Campaigns



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	page := int32(56) // int32 | The 1-indexed page number to get. The default value is `1`. (optional) (default to 1)
	recordsPerPage := int32(56) // int32 | The amount of records per page, limited to between 1 and 500 inclusive. The default value is `10`. (optional) (default to 10)
	sort := "-assignedPhoneNumbersCount" // string | Specifies the sort order for results. If not given, results are sorted by createdAt in descending order. (optional) (default to "-createdAt")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedCampaignsAPI.GetSharedCampaigns(context.Background()).Page(page).RecordsPerPage(recordsPerPage).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedCampaignsAPI.GetSharedCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSharedCampaigns`: TelnyxDownstreamCampaignRecordSet
	fmt.Fprintf(os.Stdout, "Response from `SharedCampaignsAPI.GetSharedCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSharedCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** | The 1-indexed page number to get. The default value is &#x60;1&#x60;. | [default to 1]
 **recordsPerPage** | **int32** | The amount of records per page, limited to between 1 and 500 inclusive. The default value is &#x60;10&#x60;. | [default to 10]
 **sort** | **string** | Specifies the sort order for results. If not given, results are sorted by createdAt in descending order. | [default to &quot;-createdAt&quot;]

### Return type

[**TelnyxDownstreamCampaignRecordSet**](TelnyxDownstreamCampaignRecordSet.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSharedCampaign

> TelnyxDownstreamCampaign UpdateSharedCampaign(ctx, campaignId).UpdatePartnerCampaignRequest(updatePartnerCampaignRequest).Execute()

Update Single Shared Campaign



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	campaignId := "campaignId_example" // string | 
	updatePartnerCampaignRequest := *openapiclient.NewUpdatePartnerCampaignRequest() // UpdatePartnerCampaignRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SharedCampaignsAPI.UpdateSharedCampaign(context.Background(), campaignId).UpdatePartnerCampaignRequest(updatePartnerCampaignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SharedCampaignsAPI.UpdateSharedCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSharedCampaign`: TelnyxDownstreamCampaign
	fmt.Fprintf(os.Stdout, "Response from `SharedCampaignsAPI.UpdateSharedCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSharedCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePartnerCampaignRequest** | [**UpdatePartnerCampaignRequest**](UpdatePartnerCampaignRequest.md) |  | 

### Return type

[**TelnyxDownstreamCampaign**](TelnyxDownstreamCampaign.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

