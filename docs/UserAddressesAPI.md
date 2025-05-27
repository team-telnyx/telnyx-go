# \UserAddressesAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUserAddress**](UserAddressesAPI.md#CreateUserAddress) | **Post** /user_addresses | Creates a user address
[**FindUserAddress**](UserAddressesAPI.md#FindUserAddress) | **Get** /user_addresses | List all user addresses
[**GetUserAddress**](UserAddressesAPI.md#GetUserAddress) | **Get** /user_addresses/{id} | Retrieve a user address



## CreateUserAddress

> CreateUserAddress200Response CreateUserAddress(ctx).UserAddressCreate(userAddressCreate).Execute()

Creates a user address



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
	userAddressCreate := *openapiclient.NewUserAddressCreate("Alfred", "Foster", "Toy-O'Kon", "600 Congress Avenue", "Austin", "US") // UserAddressCreate | Parameters that can be defined during user address creation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAddressesAPI.CreateUserAddress(context.Background()).UserAddressCreate(userAddressCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAddressesAPI.CreateUserAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUserAddress`: CreateUserAddress200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAddressesAPI.CreateUserAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userAddressCreate** | [**UserAddressCreate**](UserAddressCreate.md) | Parameters that can be defined during user address creation | 

### Return type

[**CreateUserAddress200Response**](CreateUserAddress200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindUserAddress

> FindUserAddress200Response FindUserAddress(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterCustomerReferenceEq(filterCustomerReferenceEq).FilterCustomerReferenceContains(filterCustomerReferenceContains).FilterStreetAddressContains(filterStreetAddressContains).Sort(sort).Execute()

List all user addresses



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
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	filterCustomerReferenceEq := "filterCustomerReferenceEq_example" // string | Filter addresses via the customer reference set. Matching is not case-sensitive. (optional)
	filterCustomerReferenceContains := "filterCustomerReferenceContains_example" // string | If present, addresses with <code>customer_reference</code> containing the given value will be returned. Matching is not case-sensitive. (optional)
	filterStreetAddressContains := "filterStreetAddressContains_example" // string | If present, addresses with <code>street_address</code> containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. (optional) (default to "null")
	sort := "street_address" // string | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the <code> -</code> prefix.<br/><br/> That is: <ul>   <li>     <code>street_address</code>: sorts the result by the     <code>street_address</code> field in ascending order.   </li>    <li>     <code>-street_address</code>: sorts the result by the     <code>street_address</code> field in descending order.   </li> </ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order. (optional) (default to "created_at")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAddressesAPI.FindUserAddress(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterCustomerReferenceEq(filterCustomerReferenceEq).FilterCustomerReferenceContains(filterCustomerReferenceContains).FilterStreetAddressContains(filterStreetAddressContains).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAddressesAPI.FindUserAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindUserAddress`: FindUserAddress200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAddressesAPI.FindUserAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindUserAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterCustomerReferenceEq** | **string** | Filter addresses via the customer reference set. Matching is not case-sensitive. | 
 **filterCustomerReferenceContains** | **string** | If present, addresses with &lt;code&gt;customer_reference&lt;/code&gt; containing the given value will be returned. Matching is not case-sensitive. | 
 **filterStreetAddressContains** | **string** | If present, addresses with &lt;code&gt;street_address&lt;/code&gt; containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. | [default to &quot;null&quot;]
 **sort** | **string** | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the &lt;code&gt; -&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;street_address&lt;/code&gt;: sorts the result by the     &lt;code&gt;street_address&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-street_address&lt;/code&gt;: sorts the result by the     &lt;code&gt;street_address&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; If not given, results are sorted by &lt;code&gt;created_at&lt;/code&gt; in descending order. | [default to &quot;created_at&quot;]

### Return type

[**FindUserAddress200Response**](FindUserAddress200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserAddress

> CreateUserAddress200Response GetUserAddress(ctx, id).Execute()

Retrieve a user address



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
	id := "id_example" // string | user address ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserAddressesAPI.GetUserAddress(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserAddressesAPI.GetUserAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUserAddress`: CreateUserAddress200Response
	fmt.Fprintf(os.Stdout, "Response from `UserAddressesAPI.GetUserAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | user address ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateUserAddress200Response**](CreateUserAddress200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

