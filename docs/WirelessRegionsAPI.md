# \WirelessRegionsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**WirelessRegionsGetAll**](WirelessRegionsAPI.md#WirelessRegionsGetAll) | **Get** /wireless/regions | Get all wireless regions



## WirelessRegionsGetAll

> ListRegions200Response WirelessRegionsGetAll(ctx).Product(product).Execute()

Get all wireless regions



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
	product := "public_ips" // string | The product for which to list regions (e.g., 'public_ips', 'private_wireless_gateways').

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WirelessRegionsAPI.WirelessRegionsGetAll(context.Background()).Product(product).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WirelessRegionsAPI.WirelessRegionsGetAll``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WirelessRegionsGetAll`: ListRegions200Response
	fmt.Fprintf(os.Stdout, "Response from `WirelessRegionsAPI.WirelessRegionsGetAll`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWirelessRegionsGetAllRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **product** | **string** | The product for which to list regions (e.g., &#39;public_ips&#39;, &#39;private_wireless_gateways&#39;). | 

### Return type

[**ListRegions200Response**](ListRegions200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

