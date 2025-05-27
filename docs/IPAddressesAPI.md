# \IPAddressesAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessIpAddress**](IPAddressesAPI.md#CreateAccessIpAddress) | **Post** /access_ip_address | Create new Access IP Address
[**DeleteAccessIpAddress**](IPAddressesAPI.md#DeleteAccessIpAddress) | **Delete** /access_ip_address/{access_ip_address_id} | Delete access IP address
[**GetAccessIpAddress**](IPAddressesAPI.md#GetAccessIpAddress) | **Get** /access_ip_address/{access_ip_address_id} | Retrieve an access IP address
[**ListAccessIpAddresses**](IPAddressesAPI.md#ListAccessIpAddresses) | **Get** /access_ip_address | List all Access IP Addresses



## CreateAccessIpAddress

> AccessIPAddressResponseSchema CreateAccessIpAddress(ctx).AccessIPAddressPOST(accessIPAddressPOST).Execute()

Create new Access IP Address

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
	accessIPAddressPOST := *openapiclient.NewAccessIPAddressPOST("IpAddress_example") // AccessIPAddressPOST | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPAddressesAPI.CreateAccessIpAddress(context.Background()).AccessIPAddressPOST(accessIPAddressPOST).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesAPI.CreateAccessIpAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAccessIpAddress`: AccessIPAddressResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `IPAddressesAPI.CreateAccessIpAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessIpAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accessIPAddressPOST** | [**AccessIPAddressPOST**](AccessIPAddressPOST.md) |  | 

### Return type

[**AccessIPAddressResponseSchema**](AccessIPAddressResponseSchema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAccessIpAddress

> AccessIPAddressResponseSchema DeleteAccessIpAddress(ctx, accessIpAddressId).Execute()

Delete access IP address

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
	accessIpAddressId := "accessIpAddressId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPAddressesAPI.DeleteAccessIpAddress(context.Background(), accessIpAddressId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesAPI.DeleteAccessIpAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAccessIpAddress`: AccessIPAddressResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `IPAddressesAPI.DeleteAccessIpAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessIpAddressId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessIpAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessIPAddressResponseSchema**](AccessIPAddressResponseSchema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAccessIpAddress

> AccessIPAddressResponseSchema GetAccessIpAddress(ctx, accessIpAddressId).Execute()

Retrieve an access IP address

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
	accessIpAddressId := "accessIpAddressId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPAddressesAPI.GetAccessIpAddress(context.Background(), accessIpAddressId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesAPI.GetAccessIpAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessIpAddress`: AccessIPAddressResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `IPAddressesAPI.GetAccessIpAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessIpAddressId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessIpAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AccessIPAddressResponseSchema**](AccessIPAddressResponseSchema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAccessIpAddresses

> AccessIPAddressListResponseSchema ListAccessIpAddresses(ctx).FilterIpSource(filterIpSource).FilterIpAddress(filterIpAddress).FilterCreatedAtGt(filterCreatedAtGt).FilterCreatedAtLt(filterCreatedAtLt).PageNumber(pageNumber).PageSize(pageSize).Execute()

List all Access IP Addresses

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/telnyx/telnyx-go"
)

func main() {
	filterIpSource := "filterIpSource_example" // string |  (optional)
	filterIpAddress := "filterIpAddress_example" // string |  (optional)
	filterCreatedAtGt := time.Now() // time.Time |  (optional)
	filterCreatedAtLt := time.Now() // time.Time |  (optional)
	pageNumber := int32(56) // int32 |  (optional) (default to 1)
	pageSize := int32(56) // int32 |  (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IPAddressesAPI.ListAccessIpAddresses(context.Background()).FilterIpSource(filterIpSource).FilterIpAddress(filterIpAddress).FilterCreatedAtGt(filterCreatedAtGt).FilterCreatedAtLt(filterCreatedAtLt).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPAddressesAPI.ListAccessIpAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAccessIpAddresses`: AccessIPAddressListResponseSchema
	fmt.Fprintf(os.Stdout, "Response from `IPAddressesAPI.ListAccessIpAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAccessIpAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterIpSource** | **string** |  | 
 **filterIpAddress** | **string** |  | 
 **filterCreatedAtGt** | **time.Time** |  | 
 **filterCreatedAtLt** | **time.Time** |  | 
 **pageNumber** | **int32** |  | [default to 1]
 **pageSize** | **int32** |  | [default to 20]

### Return type

[**AccessIPAddressListResponseSchema**](AccessIPAddressListResponseSchema.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

