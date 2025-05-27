# \DocumentsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDocument**](DocumentsAPI.md#CreateDocument) | **Post** /documents | Upload a document
[**DeleteDocument**](DocumentsAPI.md#DeleteDocument) | **Delete** /documents/{id} | Delete a document
[**DownloadDocument**](DocumentsAPI.md#DownloadDocument) | **Get** /documents/{id}/download | Download a document
[**ListDocumentLinks**](DocumentsAPI.md#ListDocumentLinks) | **Get** /document_links | List all document links
[**ListDocuments**](DocumentsAPI.md#ListDocuments) | **Get** /documents | List all documents
[**RetrieveDocument**](DocumentsAPI.md#RetrieveDocument) | **Get** /documents/{id} | Retrieve a document
[**UpdateDocument**](DocumentsAPI.md#UpdateDocument) | **Patch** /documents/{id} | Update a document



## CreateDocument

> CreateDocument200Response CreateDocument(ctx).CreateDocumentRequest(createDocumentRequest).Execute()

Upload a document



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
	createDocumentRequest := openapiclient.CreateDocument_request{CreateDocumentRequestOneOf: openapiclient.NewCreateDocumentRequestOneOf("https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf")} // CreateDocumentRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.CreateDocument(context.Background()).CreateDocumentRequest(createDocumentRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.CreateDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDocument`: CreateDocument200Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.CreateDocument`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDocumentRequest** | [**CreateDocumentRequest**](CreateDocumentRequest.md) |  | 

### Return type

[**CreateDocument200Response**](CreateDocument200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json, multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDocument

> CreateDocument200Response DeleteDocument(ctx, id).Execute()

Delete a document



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.DeleteDocument(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DeleteDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDocument`: CreateDocument200Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.DeleteDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateDocument200Response**](CreateDocument200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DownloadDocument

> map[string]interface{} DownloadDocument(ctx, id).Execute()

Download a document



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.DownloadDocument(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.DownloadDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DownloadDocument`: map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.DownloadDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDownloadDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]interface{}**

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocumentLinks

> ListDocumentLinks200Response ListDocumentLinks(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterDocumentId(filterDocumentId).FilterLinkedRecordType(filterLinkedRecordType).FilterLinkedResourceId(filterLinkedResourceId).Execute()

List all document links



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
	pageNumber := int32(56) // int32 | The page number to load (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page (optional) (default to 20)
	filterDocumentId := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the associated document to filter on. (optional)
	filterLinkedRecordType := "porting_order" // string | The `linked_record_type` of the document to filter on. (optional)
	filterLinkedResourceId := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | The `linked_resource_id` of the document to filter on. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.ListDocumentLinks(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterDocumentId(filterDocumentId).FilterLinkedRecordType(filterLinkedRecordType).FilterLinkedResourceId(filterLinkedResourceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.ListDocumentLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDocumentLinks`: ListDocumentLinks200Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.ListDocumentLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDocumentLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load | [default to 1]
 **pageSize** | **int32** | The size of the page | [default to 20]
 **filterDocumentId** | **string** | Identifies the associated document to filter on. | 
 **filterLinkedRecordType** | **string** | The &#x60;linked_record_type&#x60; of the document to filter on. | 
 **filterLinkedResourceId** | **string** | The &#x60;linked_resource_id&#x60; of the document to filter on. | 

### Return type

[**ListDocumentLinks200Response**](ListDocumentLinks200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDocuments

> ListDocuments200Response ListDocuments(ctx).FilterFilenameContains(filterFilenameContains).FilterCustomerReferenceEq(filterCustomerReferenceEq).FilterCustomerReferenceIn(filterCustomerReferenceIn).FilterCreatedAtGt(filterCreatedAtGt).FilterCreatedAtLt(filterCreatedAtLt).Sort(sort).PageNumber(pageNumber).PageSize(pageSize).Execute()

List all documents



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
	filterFilenameContains := "invoice" // string | Filter by string matching part of filename. (optional)
	filterCustomerReferenceEq := "MY REF 001" // string | Filter documents by a customer references. (optional)
	filterCustomerReferenceIn := "MY REF 001" // string | Filter documents by a list of customer references. (optional)
	filterCreatedAtGt := "2021-04-09T22:25:27.521Z" // string | Filter by created at greater than provided value. (optional)
	filterCreatedAtLt := "2021-04-09T22:25:27.521Z" // string | Filter by created at less than provided value. (optional)
	sort := "filename" // string | Specifies the sort order for results. If you want to sort by a field in ascending order, include it as a sort parameter. If you want to sort in descending order, prepend a `-` in front of the field name. (optional)
	pageNumber := int32(56) // int32 | The page number to load (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.ListDocuments(context.Background()).FilterFilenameContains(filterFilenameContains).FilterCustomerReferenceEq(filterCustomerReferenceEq).FilterCustomerReferenceIn(filterCustomerReferenceIn).FilterCreatedAtGt(filterCreatedAtGt).FilterCreatedAtLt(filterCreatedAtLt).Sort(sort).PageNumber(pageNumber).PageSize(pageSize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.ListDocuments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListDocuments`: ListDocuments200Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.ListDocuments`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDocumentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filterFilenameContains** | **string** | Filter by string matching part of filename. | 
 **filterCustomerReferenceEq** | **string** | Filter documents by a customer references. | 
 **filterCustomerReferenceIn** | **string** | Filter documents by a list of customer references. | 
 **filterCreatedAtGt** | **string** | Filter by created at greater than provided value. | 
 **filterCreatedAtLt** | **string** | Filter by created at less than provided value. | 
 **sort** | **string** | Specifies the sort order for results. If you want to sort by a field in ascending order, include it as a sort parameter. If you want to sort in descending order, prepend a &#x60;-&#x60; in front of the field name. | 
 **pageNumber** | **int32** | The page number to load | [default to 1]
 **pageSize** | **int32** | The size of the page | [default to 20]

### Return type

[**ListDocuments200Response**](ListDocuments200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDocument

> CreateDocument200Response RetrieveDocument(ctx, id).Execute()

Retrieve a document



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the resource.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.RetrieveDocument(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.RetrieveDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RetrieveDocument`: CreateDocument200Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.RetrieveDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateDocument200Response**](CreateDocument200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDocument

> CreateDocument200Response UpdateDocument(ctx, id).DocServiceDocument(docServiceDocument).Execute()

Update a document



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
	id := "6a09cdc3-8948-47f0-aa62-74ac943d6c58" // string | Identifies the resource.
	docServiceDocument := *openapiclient.NewDocServiceDocument() // DocServiceDocument |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DocumentsAPI.UpdateDocument(context.Background(), id).DocServiceDocument(docServiceDocument).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DocumentsAPI.UpdateDocument``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDocument`: CreateDocument200Response
	fmt.Fprintf(os.Stdout, "Response from `DocumentsAPI.UpdateDocument`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Identifies the resource. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDocumentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **docServiceDocument** | [**DocServiceDocument**](DocServiceDocument.md) |  | 

### Return type

[**CreateDocument200Response**](CreateDocument200Response.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

