# \AuditLogsAPI

All URIs are relative to *https://api.telnyx.com/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAuditLogs**](AuditLogsAPI.md#ListAuditLogs) | **Get** /audit_events | List Audit Logs



## ListAuditLogs

> AuditLogList ListAuditLogs(ctx).PageNumber(pageNumber).PageSize(pageSize).FilterCreatedBefore(filterCreatedBefore).FilterCreatedAfter(filterCreatedAfter).Sort(sort).Execute()

List Audit Logs



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
	pageNumber := int32(56) // int32 | The page number to load. (optional) (default to 1)
	pageSize := int32(56) // int32 | The size of the page. (optional) (default to 20)
	filterCreatedBefore := time.Now() // time.Time | Filter for audit events created before a specific date. (optional)
	filterCreatedAfter := time.Now() // time.Time | Filter for audit events created after a specific date. (optional)
	sort := "email" // string | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the <code> -</code> prefix.<br/><br/> That is: <ul>   <li>     <code>email</code>: sorts the result by the     <code>email</code> field in ascending order.   </li>    <li>     <code>-email</code>: sorts the result by the     <code>email</code> field in descending order.   </li> </ul> <br/> If not given, results are sorted by <code>created_at</code> in descending order. (optional) (default to "created_at")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditLogsAPI.ListAuditLogs(context.Background()).PageNumber(pageNumber).PageSize(pageSize).FilterCreatedBefore(filterCreatedBefore).FilterCreatedAfter(filterCreatedAfter).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditLogsAPI.ListAuditLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListAuditLogs`: AuditLogList
	fmt.Fprintf(os.Stdout, "Response from `AuditLogsAPI.ListAuditLogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAuditLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **pageNumber** | **int32** | The page number to load. | [default to 1]
 **pageSize** | **int32** | The size of the page. | [default to 20]
 **filterCreatedBefore** | **time.Time** | Filter for audit events created before a specific date. | 
 **filterCreatedAfter** | **time.Time** | Filter for audit events created after a specific date. | 
 **sort** | **string** | Specifies the sort order for results. By default sorting direction is ascending. To have the results sorted in descending order add the &lt;code&gt; -&lt;/code&gt; prefix.&lt;br/&gt;&lt;br/&gt; That is: &lt;ul&gt;   &lt;li&gt;     &lt;code&gt;email&lt;/code&gt;: sorts the result by the     &lt;code&gt;email&lt;/code&gt; field in ascending order.   &lt;/li&gt;    &lt;li&gt;     &lt;code&gt;-email&lt;/code&gt;: sorts the result by the     &lt;code&gt;email&lt;/code&gt; field in descending order.   &lt;/li&gt; &lt;/ul&gt; &lt;br/&gt; If not given, results are sorted by &lt;code&gt;created_at&lt;/code&gt; in descending order. | [default to &quot;created_at&quot;]

### Return type

[**AuditLogList**](AuditLogList.md)

### Authorization

[bearerAuth](../README.md#bearerAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

