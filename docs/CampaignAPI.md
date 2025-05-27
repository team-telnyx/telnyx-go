# \CampaignAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptCampaign**](CampaignAPI.md#AcceptCampaign) | **Post** /campaign/acceptSharing/{campaignId} | Accept Shared Campaign
[**DeactivateCampaign**](CampaignAPI.md#DeactivateCampaign) | **Delete** /campaign/{campaignId} | Deactivate My Campaign
[**GetCampaign**](CampaignAPI.md#GetCampaign) | **Get** /campaign/{campaignId} | Get My Campaign
[**GetCampaignCost**](CampaignAPI.md#GetCampaignCost) | **Get** /campaign/usecase/cost | Get Campaign Cost
[**GetCampaignMnoMetadata**](CampaignAPI.md#GetCampaignMnoMetadata) | **Get** /campaign/{campaignId}/mnoMetadata | Get Campaign Mno Metadata
[**GetCampaignOperationStatus**](CampaignAPI.md#GetCampaignOperationStatus) | **Get** /campaign/{campaignId}/operationStatus | Get My Campaign Operation Status
[**GetCampaignOsrAttributes**](CampaignAPI.md#GetCampaignOsrAttributes) | **Get** /campaign/{campaignId}/osr/attributes | Get My Osr Campaign Attributes
[**GetCampaignSharingStatus**](CampaignAPI.md#GetCampaignSharingStatus) | **Get** /campaign/{campaignId}/sharing | Get Sharing Status
[**GetCampaigns**](CampaignAPI.md#GetCampaigns) | **Get** /campaign | List Campaigns
[**GetUsecaseQualification**](CampaignAPI.md#GetUsecaseQualification) | **Get** /campaignBuilder/brand/{brandId}/usecase/{usecase} | Qualify By Usecase
[**PostCampaign**](CampaignAPI.md#PostCampaign) | **Post** /campaignBuilder | Submit Campaign
[**UpdateCampaign**](CampaignAPI.md#UpdateCampaign) | **Put** /campaign/{campaignId} | Update My Campaign



## AcceptCampaign

> interface{} AcceptCampaign(ctx, campaignId).Execute()

Accept Shared Campaign



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
	campaignId := "campaignId_example" // string | TCR's ID for the campaign to import

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAPI.AcceptCampaign(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.AcceptCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcceptCampaign`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.AcceptCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | TCR&#39;s ID for the campaign to import | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivateCampaign

> CampaignDeletionResponse DeactivateCampaign(ctx, campaignId).Execute()

Deactivate My Campaign



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
	resp, r, err := apiClient.CampaignAPI.DeactivateCampaign(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.DeactivateCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeactivateCampaign`: CampaignDeletionResponse
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.DeactivateCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CampaignDeletionResponse**](CampaignDeletionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaign

> TelnyxCampaignCSP GetCampaign(ctx, campaignId).Execute()

Get My Campaign



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
	resp, r, err := apiClient.CampaignAPI.GetCampaign(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.GetCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaign`: TelnyxCampaignCSP
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.GetCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TelnyxCampaignCSP**](TelnyxCampaignCSP.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignCost

> CampaignCost GetCampaignCost(ctx).Usecase(usecase).Execute()

Get Campaign Cost

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
	usecase := "usecase_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAPI.GetCampaignCost(context.Background()).Usecase(usecase).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.GetCampaignCost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignCost`: CampaignCost
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.GetCampaignCost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **usecase** | **string** |  | 

### Return type

[**CampaignCost**](CampaignCost.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignMnoMetadata

> MnoMetadata GetCampaignMnoMetadata(ctx, campaignId).Execute()

Get Campaign Mno Metadata



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
	resp, r, err := apiClient.CampaignAPI.GetCampaignMnoMetadata(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.GetCampaignMnoMetadata``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignMnoMetadata`: MnoMetadata
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.GetCampaignMnoMetadata`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | ID of the campaign in question | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignMnoMetadataRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**MnoMetadata**](MnoMetadata.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignOperationStatus

> map[string]interface{} GetCampaignOperationStatus(ctx, campaignId).Execute()

Get My Campaign Operation Status



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
	resp, r, err := apiClient.CampaignAPI.GetCampaignOperationStatus(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.GetCampaignOperationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignOperationStatus`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.GetCampaignOperationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignOperationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignOsrAttributes

> map[string]interface{} GetCampaignOsrAttributes(ctx, campaignId).Execute()

Get My Osr Campaign Attributes

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
	resp, r, err := apiClient.CampaignAPI.GetCampaignOsrAttributes(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.GetCampaignOsrAttributes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignOsrAttributes`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.GetCampaignOsrAttributes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignOsrAttributesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignSharingStatus

> CampaignSharingChain GetCampaignSharingStatus(ctx, campaignId).Execute()

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
	resp, r, err := apiClient.CampaignAPI.GetCampaignSharingStatus(context.Background(), campaignId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.GetCampaignSharingStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaignSharingStatus`: CampaignSharingChain
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.GetCampaignSharingStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | ID of the campaign in question | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignSharingStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CampaignSharingChain**](CampaignSharingChain.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaigns

> CampaignRecordSetCSP GetCampaigns(ctx).BrandId(brandId).Page(page).RecordsPerPage(recordsPerPage).Sort(sort).Execute()

List Campaigns



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
	brandId := "brandId_example" // string | 
	page := int32(56) // int32 | The 1-indexed page number to get. The default value is `1`. (optional) (default to 1)
	recordsPerPage := int32(56) // int32 | The amount of records per page, limited to between 1 and 500 inclusive. The default value is `10`. (optional) (default to 10)
	sort := "-assignedPhoneNumbersCount" // string | Specifies the sort order for results. If not given, results are sorted by createdAt in descending order. (optional) (default to "-createdAt")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAPI.GetCampaigns(context.Background()).BrandId(brandId).Page(page).RecordsPerPage(recordsPerPage).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.GetCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCampaigns`: CampaignRecordSetCSP
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.GetCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **brandId** | **string** |  | 
 **page** | **int32** | The 1-indexed page number to get. The default value is &#x60;1&#x60;. | [default to 1]
 **recordsPerPage** | **int32** | The amount of records per page, limited to between 1 and 500 inclusive. The default value is &#x60;10&#x60;. | [default to 10]
 **sort** | **string** | Specifies the sort order for results. If not given, results are sorted by createdAt in descending order. | [default to &quot;-createdAt&quot;]

### Return type

[**CampaignRecordSetCSP**](CampaignRecordSetCSP.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsecaseQualification

> UsecaseMetadata GetUsecaseQualification(ctx, usecase, brandId).Execute()

Qualify By Usecase



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
	usecase := "usecase_example" // string | 
	brandId := "brandId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAPI.GetUsecaseQualification(context.Background(), usecase, brandId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.GetUsecaseQualification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsecaseQualification`: UsecaseMetadata
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.GetUsecaseQualification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**usecase** | **string** |  | 
**brandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsecaseQualificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**UsecaseMetadata**](UsecaseMetadata.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostCampaign

> ResponseSubmitCampaignPublicCampaignbuilderPost PostCampaign(ctx).CampaignRequest(campaignRequest).Execute()

Submit Campaign



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
	campaignRequest := *openapiclient.NewCampaignRequest("BrandId_example", "Description_example", "Usecase_example") // CampaignRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAPI.PostCampaign(context.Background()).CampaignRequest(campaignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.PostCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostCampaign`: ResponseSubmitCampaignPublicCampaignbuilderPost
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.PostCampaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **campaignRequest** | [**CampaignRequest**](CampaignRequest.md) |  | 

### Return type

[**ResponseSubmitCampaignPublicCampaignbuilderPost**](ResponseSubmitCampaignPublicCampaignbuilderPost.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaign

> TelnyxCampaignCSP UpdateCampaign(ctx, campaignId).UpdateCampaignRequest(updateCampaignRequest).Execute()

Update My Campaign



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
	updateCampaignRequest := *openapiclient.NewUpdateCampaignRequest() // UpdateCampaignRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CampaignAPI.UpdateCampaign(context.Background(), campaignId).UpdateCampaignRequest(updateCampaignRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CampaignAPI.UpdateCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCampaign`: TelnyxCampaignCSP
	fmt.Fprintf(os.Stdout, "Response from `CampaignAPI.UpdateCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCampaignRequest** | [**UpdateCampaignRequest**](UpdateCampaignRequest.md) |  | 

### Return type

[**TelnyxCampaignCSP**](TelnyxCampaignCSP.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

