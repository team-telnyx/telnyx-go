# \PhoneNumberCampaignsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePhoneNumberCampaign**](PhoneNumberCampaignsAPI.md#CreatePhoneNumberCampaign) | **Post** /phone_number_campaigns | Create New Phone Number Campaign
[**DeletePhoneNumberCampaign**](PhoneNumberCampaignsAPI.md#DeletePhoneNumberCampaign) | **Delete** /phone_number_campaigns/{phoneNumber} | Delete Phone Number Campaign
[**GetAllPhoneNumberCampaigns**](PhoneNumberCampaignsAPI.md#GetAllPhoneNumberCampaigns) | **Get** /phone_number_campaigns | Retrieve All Phone Number Campaigns
[**GetSinglePhoneNumberCampaign**](PhoneNumberCampaignsAPI.md#GetSinglePhoneNumberCampaign) | **Get** /phone_number_campaigns/{phoneNumber} | Get Single Phone Number Campaign
[**PutPhoneNumberCampaign**](PhoneNumberCampaignsAPI.md#PutPhoneNumberCampaign) | **Put** /phone_number_campaigns/{phoneNumber} | Create New Phone Number Campaign



## CreatePhoneNumberCampaign

> PhoneNumberCampaign CreatePhoneNumberCampaign(ctx).PhoneNumberCampaignCreate(phoneNumberCampaignCreate).Execute()

Create New Phone Number Campaign

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
	phoneNumberCampaignCreate := *openapiclient.NewPhoneNumberCampaignCreate("+18005550199", "4b300178-131c-d902-d54e-72d90ba1620j") // PhoneNumberCampaignCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberCampaignsAPI.CreatePhoneNumberCampaign(context.Background()).PhoneNumberCampaignCreate(phoneNumberCampaignCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberCampaignsAPI.CreatePhoneNumberCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePhoneNumberCampaign`: PhoneNumberCampaign
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberCampaignsAPI.CreatePhoneNumberCampaign`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePhoneNumberCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **phoneNumberCampaignCreate** | [**PhoneNumberCampaignCreate**](PhoneNumberCampaignCreate.md) |  | 

### Return type

[**PhoneNumberCampaign**](PhoneNumberCampaign.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePhoneNumberCampaign

> PhoneNumberCampaign DeletePhoneNumberCampaign(ctx, phoneNumber).Execute()

Delete Phone Number Campaign



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
	phoneNumber := "phoneNumber_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberCampaignsAPI.DeletePhoneNumberCampaign(context.Background(), phoneNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberCampaignsAPI.DeletePhoneNumberCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePhoneNumberCampaign`: PhoneNumberCampaign
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberCampaignsAPI.DeletePhoneNumberCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumber** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePhoneNumberCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PhoneNumberCampaign**](PhoneNumberCampaign.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllPhoneNumberCampaigns

> PhoneNumberCampaignPaginated GetAllPhoneNumberCampaigns(ctx).RecordsPerPage(recordsPerPage).Page(page).FilterTelnyxCampaignId(filterTelnyxCampaignId).FilterTelnyxBrandId(filterTelnyxBrandId).FilterTcrCampaignId(filterTcrCampaignId).FilterTcrBrandId(filterTcrBrandId).Sort(sort).Execute()

Retrieve All Phone Number Campaigns

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
	recordsPerPage := TODO // interface{} |  (optional) (default to 20)
	page := TODO // interface{} |  (optional) (default to 1)
	filterTelnyxCampaignId := "f3575e15-32ce-400e-a4c0-dd78800c20b0" // string | Filter results by the Telnyx Campaign id (optional)
	filterTelnyxBrandId := "f3575e15-32ce-400e-a4c0-dd78800c20b0" // string | Filter results by the Telnyx Brand id (optional)
	filterTcrCampaignId := "CAMPID3" // string | Filter results by the TCR Campaign id (optional)
	filterTcrBrandId := "BRANDID" // string | Filter results by the TCR Brand id (optional)
	sort := "-phoneNumber" // string | Specifies the sort order for results. If not given, results are sorted by createdAt in descending order. (optional) (default to "-createdAt")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberCampaignsAPI.GetAllPhoneNumberCampaigns(context.Background()).RecordsPerPage(recordsPerPage).Page(page).FilterTelnyxCampaignId(filterTelnyxCampaignId).FilterTelnyxBrandId(filterTelnyxBrandId).FilterTcrCampaignId(filterTcrCampaignId).FilterTcrBrandId(filterTcrBrandId).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberCampaignsAPI.GetAllPhoneNumberCampaigns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllPhoneNumberCampaigns`: PhoneNumberCampaignPaginated
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberCampaignsAPI.GetAllPhoneNumberCampaigns`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPhoneNumberCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **recordsPerPage** | [**interface{}**](interface{}.md) |  | [default to 20]
 **page** | [**interface{}**](interface{}.md) |  | [default to 1]
 **filterTelnyxCampaignId** | **string** | Filter results by the Telnyx Campaign id | 
 **filterTelnyxBrandId** | **string** | Filter results by the Telnyx Brand id | 
 **filterTcrCampaignId** | **string** | Filter results by the TCR Campaign id | 
 **filterTcrBrandId** | **string** | Filter results by the TCR Brand id | 
 **sort** | **string** | Specifies the sort order for results. If not given, results are sorted by createdAt in descending order. | [default to &quot;-createdAt&quot;]

### Return type

[**PhoneNumberCampaignPaginated**](PhoneNumberCampaignPaginated.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSinglePhoneNumberCampaign

> PhoneNumberCampaign GetSinglePhoneNumberCampaign(ctx, phoneNumber).Execute()

Get Single Phone Number Campaign



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
	phoneNumber := "phoneNumber_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberCampaignsAPI.GetSinglePhoneNumberCampaign(context.Background(), phoneNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberCampaignsAPI.GetSinglePhoneNumberCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSinglePhoneNumberCampaign`: PhoneNumberCampaign
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberCampaignsAPI.GetSinglePhoneNumberCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumber** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSinglePhoneNumberCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PhoneNumberCampaign**](PhoneNumberCampaign.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutPhoneNumberCampaign

> PhoneNumberCampaign PutPhoneNumberCampaign(ctx, phoneNumber).PhoneNumberCampaignCreate(phoneNumberCampaignCreate).Execute()

Create New Phone Number Campaign

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
	phoneNumber := "phoneNumber_example" // string | 
	phoneNumberCampaignCreate := *openapiclient.NewPhoneNumberCampaignCreate("+18005550199", "4b300178-131c-d902-d54e-72d90ba1620j") // PhoneNumberCampaignCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PhoneNumberCampaignsAPI.PutPhoneNumberCampaign(context.Background(), phoneNumber).PhoneNumberCampaignCreate(phoneNumberCampaignCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PhoneNumberCampaignsAPI.PutPhoneNumberCampaign``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutPhoneNumberCampaign`: PhoneNumberCampaign
	fmt.Fprintf(os.Stdout, "Response from `PhoneNumberCampaignsAPI.PutPhoneNumberCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**phoneNumber** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutPhoneNumberCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **phoneNumberCampaignCreate** | [**PhoneNumberCampaignCreate**](PhoneNumberCampaignCreate.md) |  | 

### Return type

[**PhoneNumberCampaign**](PhoneNumberCampaign.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

