# \BrandsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBrandPost**](BrandsAPI.md#CreateBrandPost) | **Post** /brand | Create Brand
[**DeleteBrand**](BrandsAPI.md#DeleteBrand) | **Delete** /brand/{brandId} | Delete Brand
[**GetBrand**](BrandsAPI.md#GetBrand) | **Get** /brand/{brandId} | Get Brand
[**GetBrandFeedbackById**](BrandsAPI.md#GetBrandFeedbackById) | **Get** /brand/feedback/{brandId} | Get Brand Feedback By Id
[**GetBrands**](BrandsAPI.md#GetBrands) | **Get** /brand | List Brands
[**ListExternalVettings**](BrandsAPI.md#ListExternalVettings) | **Get** /brand/{brandId}/externalVetting | List External Vettings
[**PostOrderExternalVetting**](BrandsAPI.md#PostOrderExternalVetting) | **Post** /brand/{brandId}/externalVetting | Order Brand External Vetting
[**PutExternalVettingRecord**](BrandsAPI.md#PutExternalVettingRecord) | **Put** /brand/{brandId}/externalVetting | Import External Vetting Record
[**ResendBrand2faEmail**](BrandsAPI.md#ResendBrand2faEmail) | **Post** /brand/{brandId}/2faEmail | Resend brand 2FA email
[**RevetBrand**](BrandsAPI.md#RevetBrand) | **Put** /brand/{brandId}/revet | Revet Brand
[**UpdateBrand**](BrandsAPI.md#UpdateBrand) | **Put** /brand/{brandId} | Update Brand



## CreateBrandPost

> TelnyxBrand CreateBrandPost(ctx).CreateBrand(createBrand).Execute()

Create Brand



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
	createBrand := *openapiclient.NewCreateBrand(openapiclient.EntityType("PRIVATE_PROFIT"), "ABC Mobile", "US", "Email_example", openapiclient.Vertical("REAL_ESTATE")) // CreateBrand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandsAPI.CreateBrandPost(context.Background()).CreateBrand(createBrand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandsAPI.CreateBrandPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBrandPost`: TelnyxBrand
	fmt.Fprintf(os.Stdout, "Response from `BrandsAPI.CreateBrandPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBrandPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBrand** | [**CreateBrand**](CreateBrand.md) |  | 

### Return type

[**TelnyxBrand**](TelnyxBrand.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBrand

> interface{} DeleteBrand(ctx, brandId).Execute()

Delete Brand



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandsAPI.DeleteBrand(context.Background(), brandId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandsAPI.DeleteBrand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteBrand`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BrandsAPI.DeleteBrand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBrandRequest struct via the builder pattern


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


## GetBrand

> TelnyxBrandWithCampaignsCount GetBrand(ctx, brandId).Execute()

Get Brand



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandsAPI.GetBrand(context.Background(), brandId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandsAPI.GetBrand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBrand`: TelnyxBrandWithCampaignsCount
	fmt.Fprintf(os.Stdout, "Response from `BrandsAPI.GetBrand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**TelnyxBrandWithCampaignsCount**](TelnyxBrandWithCampaignsCount.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrandFeedbackById

> BrandFeedback GetBrandFeedbackById(ctx, brandId).Execute()

Get Brand Feedback By Id



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandsAPI.GetBrandFeedbackById(context.Background(), brandId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandsAPI.GetBrandFeedbackById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBrandFeedbackById`: BrandFeedback
	fmt.Fprintf(os.Stdout, "Response from `BrandsAPI.GetBrandFeedbackById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandFeedbackByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BrandFeedback**](BrandFeedback.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBrands

> BrandRecordSetCSP GetBrands(ctx).Page(page).RecordsPerPage(recordsPerPage).Sort(sort).DisplayName(displayName).EntityType(entityType).State(state).Country(country).BrandId(brandId).TcrBrandId(tcrBrandId).Execute()

List Brands



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
	page := int32(56) // int32 |  (optional) (default to 1)
	recordsPerPage := int32(56) // int32 | number of records per page. maximum of 500 (optional) (default to 10)
	sort := "-identityStatus" // string | Specifies the sort order for results. If not given, results are sorted by createdAt in descending order. (optional) (default to "-createdAt")
	displayName := "displayName_example" // string |  (optional)
	entityType := "entityType_example" // string |  (optional)
	state := "state_example" // string |  (optional)
	country := "country_example" // string |  (optional)
	brandId := "826ef77a-348c-445b-81a5-a9b13c68fbfe" // string | Filter results by the Telnyx Brand id (optional)
	tcrBrandId := "BBAND1" // string | Filter results by the TCR Brand id (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandsAPI.GetBrands(context.Background()).Page(page).RecordsPerPage(recordsPerPage).Sort(sort).DisplayName(displayName).EntityType(entityType).State(state).Country(country).BrandId(brandId).TcrBrandId(tcrBrandId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandsAPI.GetBrands``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBrands`: BrandRecordSetCSP
	fmt.Fprintf(os.Stdout, "Response from `BrandsAPI.GetBrands`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **recordsPerPage** | **int32** | number of records per page. maximum of 500 | [default to 10]
 **sort** | **string** | Specifies the sort order for results. If not given, results are sorted by createdAt in descending order. | [default to &quot;-createdAt&quot;]
 **displayName** | **string** |  | 
 **entityType** | **string** |  | 
 **state** | **string** |  | 
 **country** | **string** |  | 
 **brandId** | **string** | Filter results by the Telnyx Brand id | 
 **tcrBrandId** | **string** | Filter results by the TCR Brand id | 

### Return type

[**BrandRecordSetCSP**](BrandRecordSetCSP.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListExternalVettings

> interface{} ListExternalVettings(ctx, brandId).Execute()

List External Vettings



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandsAPI.ListExternalVettings(context.Background(), brandId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandsAPI.ListExternalVettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListExternalVettings`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BrandsAPI.ListExternalVettings`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListExternalVettingsRequest struct via the builder pattern


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


## PostOrderExternalVetting

> interface{} PostOrderExternalVetting(ctx, brandId).OrderExternalVetting(orderExternalVetting).Execute()

Order Brand External Vetting



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
	orderExternalVetting := *openapiclient.NewOrderExternalVetting("EvpId_example", "VettingClass_example") // OrderExternalVetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandsAPI.PostOrderExternalVetting(context.Background(), brandId).OrderExternalVetting(orderExternalVetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandsAPI.PostOrderExternalVetting``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostOrderExternalVetting`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BrandsAPI.PostOrderExternalVetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPostOrderExternalVettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderExternalVetting** | [**OrderExternalVetting**](OrderExternalVetting.md) |  | 

### Return type

**interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutExternalVettingRecord

> ExternalVetting PutExternalVettingRecord(ctx, brandId).ImportExternalVetting(importExternalVetting).Execute()

Import External Vetting Record



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
	importExternalVetting := *openapiclient.NewImportExternalVetting("EvpId_example", "VettingId_example") // ImportExternalVetting | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandsAPI.PutExternalVettingRecord(context.Background(), brandId).ImportExternalVetting(importExternalVetting).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandsAPI.PutExternalVettingRecord``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutExternalVettingRecord`: ExternalVetting
	fmt.Fprintf(os.Stdout, "Response from `BrandsAPI.PutExternalVettingRecord`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutExternalVettingRecordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **importExternalVetting** | [**ImportExternalVetting**](ImportExternalVetting.md) |  | 

### Return type

[**ExternalVetting**](ExternalVetting.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResendBrand2faEmail

> ResendBrand2faEmail(ctx, brandId).Execute()

Resend brand 2FA email

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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BrandsAPI.ResendBrand2faEmail(context.Background(), brandId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandsAPI.ResendBrand2faEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiResendBrand2faEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RevetBrand

> interface{} RevetBrand(ctx, brandId).Execute()

Revet Brand



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandsAPI.RevetBrand(context.Background(), brandId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandsAPI.RevetBrand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RevetBrand`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BrandsAPI.RevetBrand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiRevetBrandRequest struct via the builder pattern


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


## UpdateBrand

> TelnyxBrand UpdateBrand(ctx, brandId).UpdateBrand(updateBrand).Execute()

Update Brand



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
	updateBrand := *openapiclient.NewUpdateBrand(openapiclient.EntityType("PRIVATE_PROFIT"), "ABC Mobile", "US", "Email_example", openapiclient.Vertical("REAL_ESTATE")) // UpdateBrand | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BrandsAPI.UpdateBrand(context.Background(), brandId).UpdateBrand(updateBrand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BrandsAPI.UpdateBrand``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateBrand`: TelnyxBrand
	fmt.Fprintf(os.Stdout, "Response from `BrandsAPI.UpdateBrand`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**brandId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBrandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateBrand** | [**UpdateBrand**](UpdateBrand.md) |  | 

### Return type

[**TelnyxBrand**](TelnyxBrand.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

