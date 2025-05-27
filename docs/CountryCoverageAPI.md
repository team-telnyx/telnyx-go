# \CountryCoverageAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RetreiveCountryCoverage**](CountryCoverageAPI.md#RetreiveCountryCoverage) | **Get** /country_coverage | Get country coverage
[**RetreiveSpecificCountryCoverage**](CountryCoverageAPI.md#RetreiveSpecificCountryCoverage) | **Get** /country_coverage/countries/{country_code} | Get coverage for a specific country



## RetreiveCountryCoverage

> RetreiveCountryCoverage200Response RetreiveCountryCoverage(ctx).Execute()

Get country coverage



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountryCoverageAPI.RetreiveCountryCoverage(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountryCoverageAPI.RetreiveCountryCoverage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetreiveCountryCoverage`: RetreiveCountryCoverage200Response
	fmt.Fprintf(os.Stdout, "Response from `CountryCoverageAPI.RetreiveCountryCoverage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiRetreiveCountryCoverageRequest struct via the builder pattern


### Return type

[**RetreiveCountryCoverage200Response**](RetreiveCountryCoverage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetreiveSpecificCountryCoverage

> RetreiveSpecificCountryCoverage200Response RetreiveSpecificCountryCoverage(ctx, countryCode).Execute()

Get coverage for a specific country



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
	countryCode := "US" // string | Country ISO code.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.CountryCoverageAPI.RetreiveSpecificCountryCoverage(context.Background(), countryCode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `CountryCoverageAPI.RetreiveSpecificCountryCoverage``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetreiveSpecificCountryCoverage`: RetreiveSpecificCountryCoverage200Response
	fmt.Fprintf(os.Stdout, "Response from `CountryCoverageAPI.RetreiveSpecificCountryCoverage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**countryCode** | **string** | Country ISO code. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetreiveSpecificCountryCoverageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RetreiveSpecificCountryCoverage200Response**](RetreiveSpecificCountryCoverage200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

