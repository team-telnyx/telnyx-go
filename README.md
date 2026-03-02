# Telnyx Go API Library

<!-- x-release-please-start-version -->

<a href="https://pkg.go.dev/github.com/team-telnyx/telnyx-go/v4"><img src="https://pkg.go.dev/badge/github.com/team-telnyx/telnyx-go.svg" alt="Go Reference"></a>

<!-- x-release-please-end -->

The official Telnyx Go SDK — providing APIs for Voice, SMS, MMS, WhatsApp, Fax, SIP Trunking, Wireless IoT, Call Control, Number Management, and more.  
Build global communications applications on Telnyx’s private, carrier-grade network with simple Go bindings to the Telnyx REST API.

It is generated with [Stainless](https://www.stainless.com/).

## MCP Server

Use the Telnyx MCP Server to enable AI assistants to interact with this API, allowing them to explore endpoints, make test requests, and use documentation to help integrate this SDK into your application.

[![Add to Cursor](https://cursor.com/deeplink/mcp-install-dark.svg)](https://cursor.com/en-US/install-mcp?name=telnyx-mcp&config=eyJjb21tYW5kIjoibnB4IiwiYXJncyI6WyIteSIsInRlbG55eC1tY3AiXSwiZW52Ijp7IlRFTE5ZWF9BUElfS0VZIjoiTXkgQVBJIEtleSIsIlRFTE5ZWF9QVUJMSUNfS0VZIjoiTXkgUHVibGljIEtleSIsIlRFTE5ZWF9DTElFTlRfSUQiOiJNeSBDbGllbnQgSUQiLCJURUxOWVhfQ0xJRU5UX1NFQ1JFVCI6Ik15IENsaWVudCBTZWNyZXQifX0)
[![Install in VS Code](https://img.shields.io/badge/_-Add_to_VS_Code-blue?style=for-the-badge&logo=data:image/svg%2bxml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIGZpbGw9Im5vbmUiIHZpZXdCb3g9IjAgMCA0MCA0MCI+PHBhdGggZmlsbD0iI0VFRSIgZmlsbC1ydWxlPSJldmVub2RkIiBkPSJNMzAuMjM1IDM5Ljg4NGEyLjQ5MSAyLjQ5MSAwIDAgMS0xLjc4MS0uNzNMMTIuNyAyNC43OGwtMy40NiAyLjYyNC0zLjQwNiAyLjU4MmExLjY2NSAxLjY2NSAwIDAgMS0xLjA4Mi4zMzggMS42NjQgMS42NjQgMCAwIDEtMS4wNDYtLjQzMWwtMi4yLTJhMS42NjYgMS42NjYgMCAwIDEgMC0yLjQ2M0w3LjQ1OCAyMCA0LjY3IDE3LjQ1MyAxLjUwNyAxNC41N2ExLjY2NSAxLjY2NSAwIDAgMSAwLTIuNDYzbDIuMi0yYTEuNjY1IDEuNjY1IDAgMCAxIDIuMTMtLjA5N2w2Ljg2MyA1LjIwOUwyOC40NTIuODQ0YTIuNDg4IDIuNDg4IDAgMCAxIDEuODQxLS43MjljLjM1MS4wMDkuNjk5LjA5MSAxLjAxOS4yNDVsOC4yMzYgMy45NjFhMi41IDIuNSAwIDAgMSAxLjQxNSAyLjI1M3YuMDk5LS4wNDVWMzMuMzd2LS4wNDUuMDk1YTIuNTAxIDIuNTAxIDAgMCAxLTEuNDE2IDIuMjU3bC04LjIzNSAzLjk2MWEyLjQ5MiAyLjQ5MiAwIDAgMS0xLjA3Ny4yNDZabS43MTYtMjguOTQ3LTExLjk0OCA5LjA2MiAxMS45NTIgOS4wNjUtLjAwNC0xOC4xMjdaIi8+PC9zdmc+)](https://vscode.stainless.com/mcp/%7B%22name%22%3A%22telnyx-mcp%22%2C%22command%22%3A%22npx%22%2C%22args%22%3A%5B%22-y%22%2C%22telnyx-mcp%22%5D%2C%22env%22%3A%7B%22TELNYX_API_KEY%22%3A%22My%20API%20Key%22%2C%22TELNYX_PUBLIC_KEY%22%3A%22My%20Public%20Key%22%2C%22TELNYX_CLIENT_ID%22%3A%22My%20Client%20ID%22%2C%22TELNYX_CLIENT_SECRET%22%3A%22My%20Client%20Secret%22%7D%7D)

> Note: You may need to set environment variables in your MCP client.

## Installation

<!-- x-release-please-start-version -->

```go
import (
	"github.com/team-telnyx/telnyx-go/v4" // imported as telnyx
)
```

<!-- x-release-please-end -->

Or to pin the version:

<!-- x-release-please-start-version -->

```sh
go get -u 'github.com/team-telnyx/telnyx-go@v4.36.1'
```

<!-- x-release-please-end -->

## Requirements

This library requires Go 1.22+.

## Usage

The full API of this library can be found in [api.md](api.md).

```go
package main

import (
	"context"
	"fmt"

	"github.com/team-telnyx/telnyx-go/v4"
	"github.com/team-telnyx/telnyx-go/v4/option"
)

func main() {
	client := telnyx.NewClient(
		option.WithAPIKey("My API Key"), // defaults to os.LookupEnv("TELNYX_API_KEY")
	)
	response, err := client.Calls.Dial(context.TODO(), telnyx.CallDialParams{
		ConnectionID: "conn12345",
		From:         "+15557654321",
		To: telnyx.CallDialParamsToUnion{
			OfString: telnyx.String("+15551234567"),
		},
		WebhookURL: telnyx.String("https://your-webhook.url/events"),
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", response.Data)
}

```

### Request fields

The telnyx library uses the [`omitzero`](https://tip.golang.org/doc/go1.24#encodingjsonpkgencodingjson)
semantics from the Go 1.24+ `encoding/json` release for request fields.

Required primitive fields (`int64`, `string`, etc.) feature the tag <code>\`json:"...,required"\`</code>. These
fields are always serialized, even their zero values.

Optional primitive types are wrapped in a `param.Opt[T]`. These fields can be set with the provided constructors, `telnyx.String(string)`, `telnyx.Int(int64)`, etc.

Any `param.Opt[T]`, map, slice, struct or string enum uses the
tag <code>\`json:"...,omitzero"\`</code>. Its zero value is considered omitted.

The `param.IsOmitted(any)` function can confirm the presence of any `omitzero` field.

```go
p := telnyx.ExampleParams{
	ID:   "id_xxx",             // required property
	Name: telnyx.String("..."), // optional property

	Point: telnyx.Point{
		X: 0,             // required field will serialize as 0
		Y: telnyx.Int(1), // optional field will serialize as 1
		// ... omitted non-required fields will not be serialized
	},

	Origin: telnyx.Origin{}, // the zero value of [Origin] is considered omitted
}
```

To send `null` instead of a `param.Opt[T]`, use `param.Null[T]()`.
To send `null` instead of a struct `T`, use `param.NullStruct[T]()`.

```go
p.Name = param.Null[string]()       // 'null' instead of string
p.Point = param.NullStruct[Point]() // 'null' instead of struct

param.IsNull(p.Name)  // true
param.IsNull(p.Point) // true
```

Request structs contain a `.SetExtraFields(map[string]any)` method which can send non-conforming
fields in the request body. Extra fields overwrite any struct fields with a matching
key. For security reasons, only use `SetExtraFields` with trusted data.

To send a custom value instead of a struct, use `param.Override[T](value)`.

```go
// In cases where the API specifies a given type,
// but you want to send something else, use [SetExtraFields]:
p.SetExtraFields(map[string]any{
	"x": 0.01, // send "x" as a float instead of int
})

// Send a number instead of an object
custom := param.Override[telnyx.FooParams](12)
```

### Request unions

Unions are represented as a struct with fields prefixed by "Of" for each of its variants,
only one field can be non-zero. The non-zero field will be serialized.

Sub-properties of the union can be accessed via methods on the union struct.
These methods return a mutable pointer to the underlying data, if present.

```go
// Only one field can be non-zero, use param.IsOmitted() to check if a field is set
type AnimalUnionParam struct {
	OfCat *Cat `json:",omitzero,inline`
	OfDog *Dog `json:",omitzero,inline`
}

animal := AnimalUnionParam{
	OfCat: &Cat{
		Name: "Whiskers",
		Owner: PersonParam{
			Address: AddressParam{Street: "3333 Coyote Hill Rd", Zip: 0},
		},
	},
}

// Mutating a field
if address := animal.GetOwner().GetAddress(); address != nil {
	address.ZipCode = 94304
}
```

### Response objects

All fields in response structs are ordinary value types (not pointers or wrappers).
Response structs also include a special `JSON` field containing metadata about
each property.

```go
type Animal struct {
	Name   string `json:"name,nullable"`
	Owners int    `json:"owners"`
	Age    int    `json:"age"`
	JSON   struct {
		Name        respjson.Field
		Owner       respjson.Field
		Age         respjson.Field
		ExtraFields map[string]respjson.Field
	} `json:"-"`
}
```

To handle optional data, use the `.Valid()` method on the JSON field.
`.Valid()` returns true if a field is not `null`, not present, or couldn't be marshaled.

If `.Valid()` is false, the corresponding field will simply be its zero value.

```go
raw := `{"owners": 1, "name": null}`

var res Animal
json.Unmarshal([]byte(raw), &res)

// Accessing regular fields

res.Owners // 1
res.Name   // ""
res.Age    // 0

// Optional field checks

res.JSON.Owners.Valid() // true
res.JSON.Name.Valid()   // false
res.JSON.Age.Valid()    // false

// Raw JSON values

res.JSON.Owners.Raw()                  // "1"
res.JSON.Name.Raw() == "null"          // true
res.JSON.Name.Raw() == respjson.Null   // true
res.JSON.Age.Raw() == ""               // true
res.JSON.Age.Raw() == respjson.Omitted // true
```

These `.JSON` structs also include an `ExtraFields` map containing
any properties in the json response that were not specified
in the struct. This can be useful for API features not yet
present in the SDK.

```go
body := res.JSON.ExtraFields["my_unexpected_field"].Raw()
```

### Response Unions

In responses, unions are represented by a flattened struct containing all possible fields from each of the
object variants.
To convert it to a variant use the `.AsFooVariant()` method or the `.AsAny()` method if present.

If a response value union contains primitive values, primitive fields will be alongside
the properties but prefixed with `Of` and feature the tag `json:"...,inline"`.

```go
type AnimalUnion struct {
	// From variants [Dog], [Cat]
	Owner Person `json:"owner"`
	// From variant [Dog]
	DogBreed string `json:"dog_breed"`
	// From variant [Cat]
	CatBreed string `json:"cat_breed"`
	// ...

	JSON struct {
		Owner respjson.Field
		// ...
	} `json:"-"`
}

// If animal variant
if animal.Owner.Address.ZipCode == "" {
	panic("missing zip code")
}

// Switch on the variant
switch variant := animal.AsAny().(type) {
case Dog:
case Cat:
default:
	panic("unexpected type")
}
```

### RequestOptions

This library uses the functional options pattern. Functions defined in the
`option` package return a `RequestOption`, which is a closure that mutates a
`RequestConfig`. These options can be supplied to the client or at individual
requests. For example:

```go
client := telnyx.NewClient(
	// Adds a header to every request made by the client
	option.WithHeader("X-Some-Header", "custom_header_info"),
)

client.NumberOrders.New(context.TODO(), ...,
	// Override the header
	option.WithHeader("X-Some-Header", "some_other_custom_header_info"),
	// Add an undocumented field to the request body, using sjson syntax
	option.WithJSONSet("some.json.path", map[string]string{"my": "object"}),
)
```

The request option `option.WithDebugLog(nil)` may be helpful while debugging.

See the [full list of request options](https://pkg.go.dev/github.com/team-telnyx/telnyx-go/option).

### Pagination

This library provides some conveniences for working with paginated list endpoints.

You can use `.ListAutoPaging()` methods to iterate through items across all pages:

```go
iter := client.AccessIPAddress.ListAutoPaging(context.TODO(), telnyx.AccessIPAddressListParams{
	PageNumber: telnyx.Int(1),
	PageSize:   telnyx.Int(50),
})
// Automatically fetches more pages as needed.
for iter.Next() {
	accessIPAddressResponse := iter.Current()
	fmt.Printf("%+v\n", accessIPAddressResponse)
}
if err := iter.Err(); err != nil {
	panic(err.Error())
}
```

Or you can use simple `.List()` methods to fetch a single page and receive a standard response object
with additional helper methods like `.GetNextPage()`, e.g.:

```go
page, err := client.AccessIPAddress.List(context.TODO(), telnyx.AccessIPAddressListParams{
	PageNumber: telnyx.Int(1),
	PageSize:   telnyx.Int(50),
})
for page != nil {
	for _, accessIPAddress := range page.Data {
		fmt.Printf("%+v\n", accessIPAddress)
	}
	page, err = page.GetNextPage()
}
if err != nil {
	panic(err.Error())
}
```

### Errors

When the API returns a non-success status code, we return an error with type
`*telnyx.Error`. This contains the `StatusCode`, `*http.Request`, and
`*http.Response` values of the request, as well as the JSON of the error body
(much like other response objects in the SDK).

To handle errors, we recommend that you use the `errors.As` pattern:

```go
_, err := client.NumberOrders.New(context.TODO(), telnyx.NumberOrderNewParams{
	PhoneNumbers: []telnyx.NumberOrderNewParamsPhoneNumber{{
		PhoneNumber: "+15558675309",
	}},
})
if err != nil {
	var apierr *telnyx.Error
	if errors.As(err, &apierr) {
		println(string(apierr.DumpRequest(true)))  // Prints the serialized HTTP request
		println(string(apierr.DumpResponse(true))) // Prints the serialized HTTP response
	}
	panic(err.Error()) // GET "/number_orders": 400 Bad Request { ... }
}
```

When other errors occur, they are returned unwrapped; for example,
if HTTP transport fails, you might receive `*url.Error` wrapping `*net.OpError`.

### Timeouts

Requests do not time out by default; use context to configure a timeout for a request lifecycle.

Note that if a request is [retried](#retries), the context timeout does not start over.
To set a per-retry timeout, use `option.WithRequestTimeout()`.

```go
// This sets the timeout for the request, including all the retries.
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()
client.NumberOrders.New(
	ctx,
	telnyx.NumberOrderNewParams{
		PhoneNumbers: []telnyx.NumberOrderNewParamsPhoneNumber{{
			PhoneNumber: "+15558675309",
		}},
	},
	// This sets the per-retry timeout
	option.WithRequestTimeout(20*time.Second),
)
```

### File uploads

Request parameters that correspond to file uploads in multipart requests are typed as
`io.Reader`. The contents of the `io.Reader` will by default be sent as a multipart form
part with the file name of "anonymous_file" and content-type of "application/octet-stream".

The file name and content-type can be customized by implementing `Name() string` or `ContentType()
string` on the run-time type of `io.Reader`. Note that `os.File` implements `Name() string`, so a
file returned by `os.Open` will be sent with the file name on disk.

We also provide a helper `telnyx.File(reader io.Reader, filename string, contentType string)`
which can be used to wrap any `io.Reader` with the appropriate file name and content type.

```go
// A file from the file system
file, err := os.Open("/path/to/file")
telnyx.AIAudioTranscribeParams{
	Model: telnyx.AIAudioTranscribeParamsModelDistilWhisperDistilLargeV2,
	File:  file,
}

// A file from a string
telnyx.AIAudioTranscribeParams{
	Model: telnyx.AIAudioTranscribeParamsModelDistilWhisperDistilLargeV2,
	File:  strings.NewReader("my file contents"),
}

// With a custom filename and contentType
telnyx.AIAudioTranscribeParams{
	Model: telnyx.AIAudioTranscribeParamsModelDistilWhisperDistilLargeV2,
	File:  telnyx.File(strings.NewReader(`{"hello": "foo"}`), "file.go", "application/json"),
}
```

### Retries

Certain errors will be automatically retried 2 times by default, with a short exponential backoff.
We retry by default all connection errors, 408 Request Timeout, 409 Conflict, 429 Rate Limit,
and >=500 Internal errors.

You can use the `WithMaxRetries` option to configure or disable this:

```go
// Configure the default for all requests:
client := telnyx.NewClient(
	option.WithMaxRetries(0), // default is 2
)

// Override per-request:
client.NumberOrders.New(
	context.TODO(),
	telnyx.NumberOrderNewParams{
		PhoneNumbers: []telnyx.NumberOrderNewParamsPhoneNumber{{
			PhoneNumber: "+15558675309",
		}},
	},
	option.WithMaxRetries(5),
)
```

### Accessing raw response data (e.g. response headers)

You can access the raw HTTP response data by using the `option.WithResponseInto()` request option. This is useful when
you need to examine response headers, status codes, or other details.

```go
// Create a variable to store the HTTP response
var response *http.Response
numberOrder, err := client.NumberOrders.New(
	context.TODO(),
	telnyx.NumberOrderNewParams{
		PhoneNumbers: []telnyx.NumberOrderNewParamsPhoneNumber{{
			PhoneNumber: "+15558675309",
		}},
	},
	option.WithResponseInto(&response),
)
if err != nil {
	// handle error
}
fmt.Printf("%+v\n", numberOrder)

fmt.Printf("Status Code: %d\n", response.StatusCode)
fmt.Printf("Headers: %+#v\n", response.Header)
```

### Making custom/undocumented requests

This library is typed for convenient access to the documented API. If you need to access undocumented
endpoints, params, or response properties, the library can still be used.

#### Undocumented endpoints

To make requests to undocumented endpoints, you can use `client.Get`, `client.Post`, and other HTTP verbs.
`RequestOptions` on the client, such as retries, will be respected when making these requests.

```go
var (
    // params can be an io.Reader, a []byte, an encoding/json serializable object,
    // or a "…Params" struct defined in this library.
    params map[string]any

    // result can be an []byte, *http.Response, a encoding/json deserializable object,
    // or a model defined in this library.
    result *http.Response
)
err := client.Post(context.Background(), "/unspecified", params, &result)
if err != nil {
    …
}
```

#### Undocumented request params

To make requests using undocumented parameters, you may use either the `option.WithQuerySet()`
or the `option.WithJSONSet()` methods.

```go
params := FooNewParams{
    ID:   "id_xxxx",
    Data: FooNewParamsData{
        FirstName: telnyx.String("John"),
    },
}
client.Foo.New(context.Background(), params, option.WithJSONSet("data.last_name", "Doe"))
```

#### Undocumented response properties

To access undocumented response properties, you may either access the raw JSON of the response as a string
with `result.JSON.RawJSON()`, or get the raw JSON of a particular field on the result with
`result.JSON.Foo.Raw()`.

Any fields that are not present on the response struct will be saved and can be accessed by `result.JSON.ExtraFields()` which returns the extra fields as a `map[string]Field`.

### Middleware

We provide `option.WithMiddleware` which applies the given
middleware to requests.

```go
func Logger(req *http.Request, next option.MiddlewareNext) (res *http.Response, err error) {
	// Before the request
	start := time.Now()
	LogReq(req)

	// Forward the request to the next handler
	res, err = next(req)

	// Handle stuff after the request
	end := time.Now()
	LogRes(res, err, start - end)

    return res, err
}

client := telnyx.NewClient(
	option.WithMiddleware(Logger),
)
```

When multiple middlewares are provided as variadic arguments, the middlewares
are applied left to right. If `option.WithMiddleware` is given
multiple times, for example first in the client then the method, the
middleware in the client will run first and the middleware given in the method
will run next.

You may also replace the default `http.Client` with
`option.WithHTTPClient(client)`. Only one http client is
accepted (this overwrites any previous client) and receives requests after any
middleware has been applied.

## Semantic versioning

This package generally follows [SemVer](https://semver.org/spec/v2.0.0.html) conventions, though certain backwards-incompatible changes may be released as minor versions:

1. Changes to library internals which are technically public but not intended or documented for external use. _(Please open a GitHub issue to let us know if you are relying on such internals.)_
2. Changes that we do not expect to impact the vast majority of users in practice.

We take backwards-compatibility seriously and work hard to ensure you can rely on a smooth upgrade experience.

We are keen for your feedback; please open an [issue](https://www.github.com/team-telnyx/telnyx-go/issues) with questions, bugs, or suggestions.

## Contributing

See [the contributing documentation](./CONTRIBUTING.md).
