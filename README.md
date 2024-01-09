# Metronome Go API Library

<a href="https://pkg.go.dev/github.com/Metronome-Industries/metronome-go"><img src="https://pkg.go.dev/badge/github.com/Metronome-Industries/metronome-go.svg" alt="Go Reference"></a>

The Metronome Go library provides convenient access to [the Metronome REST
API](https://docs.metronome.com) from applications written in Go.

## Installation

```go
import (
	"github.com/Metronome-Industries/metronome-go" // imported as metronome
)
```

Or to pin the version:

```sh
go get -u 'github.com/Metronome-Industries/metronome-go@v0.0.1'
```

## Requirements

This library requires Go 1.18+.

## Usage

The full API of this library can be found in [api.md](https://www.github.com/Metronome-Industries/metronome-go/blob/main/api.md).

```go
package main

import (
	"context"
	"github.com/Metronome-Industries/metronome-go"
	"github.com/Metronome-Industries/metronome-go/option"
)

func main() {
	client := metronome.NewClient(
		option.WithBearerToken("My Bearer Token"), // defaults to os.LookupEnv("METRONOME_BEARER_TOKEN")
	)
	err := client.Ingest(context.TODO(), metronome.IngestParams{
		Body: metronome.F([]metronome.IngestParamsBody{{
			TransactionID: metronome.F("2021-01-01T00:00:00+00:00_cluster42"),
			CustomerID:    metronome.F("team@example.com"),
			EventType:     metronome.F("heartbeat"),
			Timestamp:     metronome.F("2021-01-01T00:00:00+00:00"),
		}}),
	})
	if err != nil {
		panic(err.Error())
	}
}

```

### Request Fields

All request parameters are wrapped in a generic `Field` type,
which we use to distinguish zero values from null or omitted fields.

This prevents accidentally sending a zero value if you forget a required parameter,
and enables explicitly sending `null`, `false`, `''`, or `0` on optional parameters.
Any field not specified is not sent.

To construct fields with values, use the helpers `String()`, `Int()`, `Float()`, or most commonly, the generic `F[T]()`.
To send a null, use `Null[T]()`, and to send a nonconforming value, use `Raw[T](any)`. For example:

```go
params := FooParams{
	Name: metronome.F("hello"),

	// Explicitly send `"description": null`
	Description: metronome.Null[string](),

	Point: metronome.F(metronome.Point{
		X: metronome.Int(0),
		Y: metronome.Int(1),

		// In cases where the API specifies a given type,
		// but you want to send something else, use `Raw`:
		Z: metronome.Raw[int64](0.01), // sends a float
	}),
}
```

### Response Objects

All fields in response structs are value types (not pointers or wrappers).

If a given field is `null`, not present, or invalid, the corresponding field
will simply be its zero value.

All response structs also include a special `JSON` field, containing more detailed
information about each property, which you can use like so:

```go
if res.Name == "" {
	// true if `"name"` is either not present or explicitly null
	res.JSON.Name.IsNull()

	// true if the `"name"` key was not present in the repsonse JSON at all
	res.JSON.Name.IsMissing()

	// When the API returns data that cannot be coerced to the expected type:
	if res.JSON.Name.IsInvalid() {
		raw := res.JSON.Name.Raw()

		legacyName := struct{
			First string `json:"first"`
			Last  string `json:"last"`
		}{}
		json.Unmarshal([]byte(raw), &legacyName)
		name = legacyName.First + " " + legacyName.Last
	}
}
```

These `.JSON` structs also include an `Extras` map containing
any properties in the json response that were not specified
in the struct. This can be useful for API features not yet
present in the SDK.

```go
body := res.JSON.ExtraFields["my_unexpected_field"].Raw()
```

### RequestOptions

This library uses the functional options pattern. Functions defined in the
`option` package return a `RequestOption`, which is a closure that mutates a
`RequestConfig`. These options can be supplied to the client or at individual
requests. For example:

```go
client := metronome.NewClient(
	// Adds a header to every request made by the client
	option.WithHeader("X-Some-Header", "custom_header_info"),
)

client.Alerts.New(context.TODO(), ...,
	// Override the header
	option.WithHeader("X-Some-Header", "some_other_custom_header_info"),
	// Add an undocumented field to the request body, using sjson syntax
	option.WithJSONSet("some.json.path", map[string]string{"my": "object"}),
)
```

The full list of request options is [here](https://pkg.go.dev/github.com/Metronome-Industries/metronome-go/option).

### Pagination

This library provides some conveniences for working with paginated list endpoints.

You can use `.ListAutoPaging()` methods to iterate through items across all pages:

```go
// TODO
```

Or you can use simple `.List()` methods to fetch a single page and receive a standard response object
with additional helper methods like `.GetNextPage()`, e.g.:

```go
// TODO
```

### Errors

When the API returns a non-success status code, we return an error with type
`*metronome.Error`. This contains the `StatusCode`, `*http.Request`, and
`*http.Response` values of the request, as well as the JSON of the error body
(much like other response objects in the SDK).

To handle errors, we recommend that you use the `errors.As` pattern:

```go
_, err := client.Alerts.New(context.TODO(), metronome.AlertNewParams{
	AlertType: metronome.F(metronome.AlertNewParamsAlertTypeLowCreditBalanceReached),
	Name:      metronome.F("$100 credit balance alert for single customer"),
	Threshold: metronome.F(10000.000000),
})
if err != nil {
	var apierr *metronome.Error
	if errors.As(err, &apierr) {
		println(string(apierr.DumpRequest(true)))  // Prints the serialized HTTP request
		println(string(apierr.DumpResponse(true))) // Prints the serialized HTTP response
	}
	panic(err.Error()) // GET "/alerts/create": 400 Bad Request { ... }
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
client.Alerts.New(
	ctx,
	metronome.AlertNewParams{
		AlertType: metronome.F(metronome.AlertNewParamsAlertTypeLowCreditBalanceReached),
		Name:      metronome.F("$100 credit balance alert for single customer"),
		Threshold: metronome.F(10000.000000),
	},
	// This sets the per-retry timeout
	option.WithRequestTimeout(20*time.Second),
)
```

## Retries

Certain errors will be automatically retried 2 times by default, with a short exponential backoff.
We retry by default all connection errors, 408 Request Timeout, 409 Conflict, 429 Rate Limit,
and >=500 Internal errors.

You can use the `WithMaxRetries` option to configure or disable this:

```go
// Configure the default for all requests:
client := metronome.NewClient(
	option.WithMaxRetries(0), // default is 2
)

// Override per-request:
client.Alerts.New(
	context.TODO(),
	metronome.AlertNewParams{
		AlertType: metronome.F(metronome.AlertNewParamsAlertTypeLowCreditBalanceReached),
		Name:      metronome.F("$100 credit balance alert for single customer"),
		Threshold: metronome.F(10000.000000),
	},
	option.WithMaxRetries(5),
)
```

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

client := metronome.NewClient(
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

## Semantic Versioning

This package generally follows [SemVer](https://semver.org/spec/v2.0.0.html) conventions, though certain backwards-incompatible changes may be released as minor versions:

1. Changes to library internals which are technically public but not intended or documented for external use. _(Please open a GitHub issue to let us know if you are relying on such internals)_.
2. Changes that we do not expect to impact the vast majority of users in practice.

We take backwards-compatibility seriously and work hard to ensure you can rely on a smooth upgrade experience.

We are keen for your feedback; please open an [issue](https://www.github.com/Metronome-Industries/metronome-go/issues) with questions, bugs, or suggestions.
