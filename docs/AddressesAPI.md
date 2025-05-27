# \AddressesAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AcceptAddressSuggestions**](AddressesAPI.md#AcceptAddressSuggestions) | **Post** /addresses/{id}/actions/accept_suggestions | Accepts this address suggestion as a new emergency address for Operator Connect and finishes the uploads of the numbers associated with it to Microsoft.
[**CreateAddress**](AddressesAPI.md#CreateAddress) | **Post** /addresses | Creates an address
[**DeleteAddress**](AddressesAPI.md#DeleteAddress) | **Delete** /addresses/{id} | Deletes an address
[**FindAddresses**](AddressesAPI.md#FindAddresses) | **Get** /addresses | List all addresses
[**GetAddress**](AddressesAPI.md#GetAddress) | **Get** /addresses/{id} | Retrieve an address
[**ValidateAddress**](AddressesAPI.md#ValidateAddress) | **Post** /addresses/actions/validate | Validate an address



## AcceptAddressSuggestions

> AddressSuggestionResponse AcceptAddressSuggestions(ctx, id).AcceptSuggestionsRequest(acceptSuggestionsRequest).Execute()

Accepts this address suggestion as a new emergency address for Operator Connect and finishes the uploads of the numbers associated with it to Microsoft.

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
	id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The UUID of the address that should be accepted.
	acceptSuggestionsRequest := *openapiclient.NewAcceptSuggestionsRequest() // AcceptSuggestionsRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.AcceptAddressSuggestions(context.Background(), id).AcceptSuggestionsRequest(acceptSuggestionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.AcceptAddressSuggestions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AcceptAddressSuggestions`: AddressSuggestionResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.AcceptAddressSuggestions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The UUID of the address that should be accepted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAcceptAddressSuggestionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **acceptSuggestionsRequest** | [**AcceptSuggestionsRequest**](AcceptSuggestionsRequest.md) |  | 

### Return type

[**AddressSuggestionResponse**](AddressSuggestionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAddress

> CreateAddress200Response CreateAddress(ctx).AddressCreate(addressCreate).Execute()

Creates an address



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
	addressCreate := *openapiclient.NewAddressCreate("Alfred", "Foster", "Toy-O'Kon", "600 Congress Avenue", "Austin", "US") // AddressCreate | Parameters that can be defined during address creation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.CreateAddress(context.Background()).AddressCreate(addressCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.CreateAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAddress`: CreateAddress200Response
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.CreateAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addressCreate** | [**AddressCreate**](AddressCreate.md) | Parameters that can be defined during address creation | 

### Return type

[**CreateAddress200Response**](CreateAddress200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAddress

> CreateAddress200Response DeleteAddress(ctx, id).Execute()

Deletes an address



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
	id := "id_example" // string | address ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.DeleteAddress(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.DeleteAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAddress`: CreateAddress200Response
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.DeleteAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | address ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateAddress200Response**](CreateAddress200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAddresses

> FindAddresses200Response FindAddresses(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterCustomerReferenceEq(filterCustomerReferenceEq).FilterCustomerReferenceContains(filterCustomerReferenceContains).FilterUsedAsEmergency(filterUsedAsEmergency).FilterStreetAddressContains(filterStreetAddressContains).FilterAddressBookEq(filterAddressBookEq).Sort(sort).Execute()

List all addresses



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
	filterUsedAsEmergency := "filterUsedAsEmergency_example" // string | If set as 'true', only addresses used as the emergency address for at least one active phone-number will be returned. When set to 'false', the opposite happens: only addresses not used as the emergency address from phone-numbers will be returned. (optional) (default to "null")
	filterStreetAddressContains := "filterStreetAddressContains_example" // string | If present, addresses with <code>street_address</code> containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. (optional) (default to "null")
	filterAddressBookEq := "filterAddressBookEq_example" // string | If present, only returns results with the <code>address_book</code> flag set to the given value. (optional) (default to "null")
	sort := "street_address" // string | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the <code> -</code> prefix.<br/><br/> That is: <ul>   <li>     <code>street_address</code>: sorts the result by the     <code>street_address</code> field in ascending order.   </li>    <li>     <code>-street_address</code>: sorts the result by the     <code>street_address</code> field in descending order.   </li> </ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order. (optional) (default to "created_at")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.FindAddresses(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterCustomerReferenceEq(filterCustomerReferenceEq).FilterCustomerReferenceContains(filterCustomerReferenceContains).FilterUsedAsEmergency(filterUsedAsEmergency).FilterStreetAddressContains(filterStreetAddressContains).FilterAddressBookEq(filterAddressBookEq).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.FindAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindAddresses`: FindAddresses200Response
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.FindAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterCustomerReferenceEq** | **string** | Filter addresses via the customer reference set. Matching is not case-sensitive. | 
 **filterCustomerReferenceContains** | **string** | If present, addresses with &lt;code&gt;customer_reference&lt;/code&gt; containing the given value will be returned. Matching is not case-sensitive. | 
 **filterUsedAsEmergency** | **string** | If set as &#39;true&#39;, only addresses used as the emergency address for at least one active phone-number will be returned. When set to &#39;false&#39;, the opposite happens: only addresses not used as the emergency address from phone-numbers will be returned. | [default to &quot;null&quot;]
 **filterStreetAddressContains** | **string** | If present, addresses with &lt;code&gt;street_address&lt;/code&gt; containing the given value will be returned. Matching is not case-sensitive. Requires at least three characters. | [default to &quot;null&quot;]
 **filterAddressBookEq** | **string** | If present, only returns results with the &lt;code&gt;address_book&lt;/code&gt; flag set to the given value. | [default to &quot;null&quot;]
 **sort** | **string** | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the &lt;code&gt; -&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;street_address&lt;/code&gt;: sorts the result by the     &lt;code&gt;street_address&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-street_address&lt;/code&gt;: sorts the result by the     &lt;code&gt;street_address&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; If not given, results are sorted by &lt;code&gt;created_at&lt;/code&gt; in descending order. | [default to &quot;created_at&quot;]

### Return type

[**FindAddresses200Response**](FindAddresses200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAddress

> CreateAddress200Response GetAddress(ctx, id).Execute()

Retrieve an address



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
	id := "id_example" // string | address ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.GetAddress(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.GetAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAddress`: CreateAddress200Response
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.GetAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | address ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateAddress200Response**](CreateAddress200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateAddress

> ValidateAddressActionResponse ValidateAddress(ctx).ValidateAddressRequest(validateAddressRequest).Execute()

Validate an address



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
	validateAddressRequest := *openapiclient.NewValidateAddressRequest("600 Congress Avenue", "78701", "US") // ValidateAddressRequest | Parameters that can be defined during address validation

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AddressesAPI.ValidateAddress(context.Background()).ValidateAddressRequest(validateAddressRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AddressesAPI.ValidateAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateAddress`: ValidateAddressActionResponse
	fmt.Fprintf(os.Stdout, "Response from `AddressesAPI.ValidateAddress`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validateAddressRequest** | [**ValidateAddressRequest**](ValidateAddressRequest.md) | Parameters that can be defined during address validation | 

### Return type

[**ValidateAddressActionResponse**](ValidateAddressActionResponse.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

