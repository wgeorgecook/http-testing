# Resource Server for HTTP Testing
This repository contains an API server AND the client SDK
for consuming it. Both the server and SDK are very simple,
they serve and get `Resource` (defined in `/internal/pkg/resources/resources.go`)
provided by ID. The examples are a bit contrived but they get the job done.

## Testing
### Server
The server is tested using the `net/http/httptest"` tools provided by the 
standard library. The tests are found in `internal/pkg/server/handlers_test.go`
and ensure that calls to the handler return expected results. 

### Client
The client library provides a factor for creating new clients and a method to
get resources by ID. A custom HTTP client interface is embedded in the Client 
that the default HTTP client implements, but this also allows embedding a mockable
HTTP client that also satisfies the interface. The custom mock-HTTP client is 
located at `/internal/pkg/utils/mocks/httpClientMock.go`. This dependency is injected
when instantiating the Client to run the tests and hijacks the calls to `c.HttpClient.Get(url)`
in the client method `GetResourceByID`. 

## Running Tests
At the root of the repo, run `go test ./...` to run the entire suite at once. 
```shell 
 go test -v ./...
       github.com/wgeorgecook/testing-http/cmd [no test files]
       github.com/wgeorgecook/testing-http/internal/pkg/resources      [no test files]
       github.com/wgeorgecook/testing-http/internal/pkg/utils/errs     [no test files]
       github.com/wgeorgecook/testing-http/internal/pkg/utils/mocks    [no test files]
=== RUN   TestClient_GetResourceByID
--- PASS: TestClient_GetResourceByID (0.00s)
PASS
ok      github.com/wgeorgecook/testing-http/internal/pkg/api    0.005s
=== RUN   Test_getResourceHandler
=== RUN   Test_getResourceHandler/No_ID_is_Bad_Request
=== RUN   Test_getResourceHandler/ID_1_is_Returned
=== RUN   Test_getResourceHandler/ID_3_is_Not_Found
--- PASS: Test_getResourceHandler (0.00s)
    --- PASS: Test_getResourceHandler/No_ID_is_Bad_Request (0.00s)
    --- PASS: Test_getResourceHandler/ID_1_is_Returned (0.00s)
    --- PASS: Test_getResourceHandler/ID_3_is_Not_Found (0.00s)
PASS
ok      github.com/wgeorgecook/testing-http/internal/pkg/server 0.005s

```