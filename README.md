[![golang-ci](https://github.com/bar-counter/gin-correlation-id/actions/workflows/golang-ci.yml/badge.svg)](https://github.com/bar-counter/gin-correlation-id/actions/workflows/golang-ci.yml)
[![license](https://img.shields.io/github/license/bar-counter/gin-correlation-id)](https://github.com/bar-counter/gin-correlation-id)
[![go mod version](https://img.shields.io/github/go-mod/go-version/bar-counter/gin-correlation-id?label=go.mod)](https://github.com/bar-counter/gin-correlation-id)
[![GoDoc](https://godoc.org/github.com/bar-counter/gin-correlation-id?status.png)](https://godoc.org/github.com/bar-counter/gin-correlation-id/)
[![GoReportCard](https://goreportcard.com/badge/github.com/bar-counter/gin-correlation-id)](https://goreportcard.com/report/github.com/bar-counter/gin-correlation-id)
[![codecov](https://codecov.io/gh/bar-counter/gin-correlation-id/branch/main/graph/badge.svg)](https://codecov.io/gh/bar-counter/gin-correlation-id)
[![github release](https://img.shields.io/github/v/release/bar-counter/gin-correlation-id?style=social)](https://github.com/bar-counter/gin-correlation-id/releases)

## for what

Middleware for reading or generating correlation IDs for each incoming request.

Correlation IDs can then be added to your logs, making it simple to retrieve all logs generated from a single HTTP
request.

The middleware checks for the `x-request-id` header by default, but can be set to any key，like `x-correlation-id`

## Contributing

[![Contributor Covenant](https://img.shields.io/badge/contributor%20covenant-v1.4-ff69b4.svg)](.github/CONTRIBUTING_DOC/CODE_OF_CONDUCT.md)
[![GitHub contributors](https://img.shields.io/github/contributors/bar-counter/gin-correlation-id)](https://github.com/bar-counter/gin-correlation-id/graphs/contributors)

We welcome community contributions to this project.

Please read [Contributor Guide](.github/CONTRIBUTING_DOC/CONTRIBUTING.md) for more information on how to get started.

## depends

in go mod project

```bash
# before above global settings
# test version info
$ git ls-remote -q https://github.com/bar-counter/gin-correlation-id.git

# test depends see full version
$ go list -mod readonly -v -m -versions github.com/bar-counter/gin-correlation-id
# or use last version add go.mod by script
$ echo "go mod edit -require=$(go list -mod=readonly -m -versions github.com/bar-counter/gin-correlation-id | awk '{print $1 "@" $NF}')"
$ echo "go mod vendor"
```

## env

- minimum go version: go 1.18
- change `go 1.18`, `^1.18`, `1.18.10` to new go version

### libs

| lib                                      | version |
|:-----------------------------------------|:--------|
| https://github.com/stretchr/testify      | v1.8.4  |
| https://github.com/gin-gonic/gin         | v1.9.1  |
| https://github.com/gofrs/uuid/v5         | v5.0.0  |
| https://github.com/bwmarrin/snowflake    | v0.3.0  |
| https://github.com/lithammer/shortuuid/4 | v4.0.0  |

## Feature

- id creation support
    - [x] uuid v4 by [https://github.com/gofrs/uuid](https://github.com/gofrs/uuid)
    - [x] snowflake [https://github.com/bwmarrin/snowflake](https://github.com/bwmarrin/snowflake)
    - [x] shortuuid by [https://github.com/lithammer/shortuuid](https://github.com/lithammer/shortuuid)
- [x] change CorrelationID Header key at each support
- [X] [CORS](#CORS) cross-origin resource sharing
- [X] more perfect test case coverage
- [X] more perfect benchmark case

## Performance

- Performance log
  see [PerformanceLog.md](https://github.com/bar-counter/gin-correlation-id/blob/main/doc/PerformanceLog.md)

| platform | arch  | method                                            | times  | ns/op       | B/op      | allocs/op     | cpu                                            |
|:---------|:------|:--------------------------------------------------|:-------|:------------|:----------|:--------------|:-----------------------------------------------|
| linux    | amd64 | Benchmark_gin_correlation_id_uuidv4-2             | 167029 | 7124 ns/op  | 2320 B/op | 32 allocs/op  | Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz      |
| linux    | amd64 | BenchmarkParallel_gin_correlation_id_uuidv4-2     | 218067 | 5572 ns/op  | 2320 B/op | 32 allocs/op  | Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz      |
| linux    | amd64 | BenchmarkParallel_gin_correlation_id_snowflake-2  | 260205 | 4920 ns/op  | 2256 B/op | 31 allocs/op  | Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz      |
| linux    | amd64 | Benchmark_gin_correlation_id_snowflake-2          | 199422 | 5922 ns/op  | 2256 B/op | 31 allocs/op  | Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz      |
| linux    | amd64 | BenchmarkGinIdShortUuid-2                         | 73371  | 15670 ns/op | 4891 B/op | 128 allocs/op | Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz      |
| linux    | amd64 | BenchmarkParallelGinIdShortUuid-2                 | 103296 | 10477 ns/op | 4891 B/op | 128 allocs/op | Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz      |
| windows  | amd64 | Benchmark_gin_correlation_id_uuidv4-2             | 215640 | 5563 ns/op  | 2344 B/op | 33 allocs/op  | Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz |
| windows  | amd64 | BenchmarkParallel_gin_correlation_id_uuidv4-2     | 291274 | 3558 ns/op  | 2344 B/op | 33 allocs/op  | Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz |
| windows  | amd64 | Benchmark_gin_correlation_id_snowflake-2          | 228192 | 5059 ns/op  | 2256 B/op | 31 allocs/op  | Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz |
| windows  | amd64 | BenchmarkParallel_gin_correlation_id_snowflake-2  | 347394 | 3327 ns/op  | 2256 B/op | 31 allocs/op  | Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz |
| windows  | amd64 | BenchmarkGinIdShortUuid-2                         | 95556  | 12331 ns/op | 4915 B/op | 129 allocs/op | Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz |
| windows  | amd64 | BenchmarkParallelGinIdShortUuid-2                 | 153523 | 8331 ns/op  | 4915 B/op | 129 allocs/op | Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz |
| darwin   | arm64 | Benchmark_gin_correlation_id_uuidv4-2             | 440373 | 2712 ns/op  | 2321 B/op | 32 allocs/op  | apple silicon M1 Max core 10                   |
| darwin   | arm64 | BenchmarkParallel_gin_correlation_id_uuidv4-10    | 513122 | 2296 ns/op  | 2324 B/op | 32 allocs/op  | apple silicon M1 Max core 10                   |
| darwin   | arm64 | Benchmark_gin_correlation_id_snowflake-2          | 570752 | 2095 ns/op  | 2257 B/op | 31 allocs/op  | apple silicon M1 Max core 10                   |
| darwin   | arm64 | BenchmarkParallel_gin_correlation_id_snowflake-10 | 557127 | 2213 ns/op  | 2260 B/op | 31 allocs/op  | apple silicon M1 Max core 10                   |
| darwin   | arm64 | BenchmarkGinIdShortUuid-2                         | 202786 | 5920 ns/op  | 4893 B/op | 128 allocs/op | apple silicon M1 Max core 10                   |
| darwin   | arm64 | BenchmarkParallelGinIdShortUuid-10                | 330234 | 3614 ns/op  | 4898 B/op | 128 allocs/op | apple silicon M1 Max core 10                   |

- `ns/op` is the average time per operation, lower latency better performance
- `B/op` is the number of bytes allocated per operation, memory allocation means getting more memory from the operating system, less occupied better
- `allocs/op` is the number of allocations per operation, memory allocation means getting more memory from the operating system, less occupied better

## usage

- see full api usage at unit test case
    - [gin_correlation_id_uuidv4](https://github.com/bar-counter/gin-correlation-id/blob/main/example/ginid_uuidv4_test/ping_test.go)
    - [gin_correlation_id_snowflake](https://github.com/bar-counter/gin-correlation-id/blob/main/example/ginid_snowflake_test/ping_test.go)
    - [gin_correlation_id_shortuuid](https://github.com/bar-counter/gin-correlation-id/blob/main/example/ginid_shortuuid_test/ping_test.go)
    - `go 1.18`, `^1.18`, `1.18.10` to new go version for dev

### sample uuid-v4

```go
package main

import (
	"github.com/bar-counter/gin-correlation-id/gin_correlation_id_uuidv4"
	"github.com/gin-gonic/gin"
	"net/http"
	"log"
	"fmt"
)

func ginPingRouter() *gin.Engine {
	router := gin.New()

	// add Middleware
	router.Use(gin_correlation_id_uuidv4.Middleware())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong as correlation ID: %s", gin_correlation_id_uuidv4.GetCorrelationID(c))
	})
	return router
}

func main() {
	// Create the Gin engine. 
	g := ginPingRouter()
	serverPort := "49002"
	log.Printf("=> now server access as: http://127.0.0.1:%s", serverPort)
	err := http.ListenAndServe(fmt.Sprintf(":%v", serverPort), g)
	if err != nil {
		fmt.Printf("run gin server err %v\n", err)
		return
	}
}
```

## integrationTest

- run server then test as curl

```bash
$ curl -v http://127.0.0.1:49002/ping
*   Trying 127.0.0.1:49002...
* Connected to 127.0.0.1 (127.0.0.1) port 49002 (#0)
> GET /ping HTTP/1.1
> Host: 127.0.0.1:49002
> User-Agent: curl/7.80.0
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Access-Control-Expose-Headers: x-request-id
< Content-Type: text/plain; charset=utf-8
< X-Request-Id: 0b737fda-5439-487b-9f81-a7824b88532e
< Date: Sun, 18 Jun 2023 04:20:27 GMT
< Content-Length: 60
<
* Connection #0 to host 127.0.0.1 left intact
pong as correlation ID: 0b737fda-5439-487b-9f81-a7824b88532e
```

1. `response header` has `X-Request-Id` key as `correlation ID`
2. can get `correlation ID` as

```go
    // use uuid v4
gin_correlation_id_uuidv4.GetCorrelationID(*gin.Context)

// use snowflake
gin_correlation_id_snowflake.GetCorrelationID(*gin.Context)

// use shortuuid
gin_correlation_id_shortuuid.GetCorrelationID(*gin.Context)
```

# CORS

If you are using cross-origin resource sharing (CORS)

> e.g. you are making requests to an API from a frontend JavaScript code served from a different origin, you have to
> ensure two things:

- permit correlation ID header in the incoming requests' HTTP headers so the value can be reused by the middleware,
- add the correlation ID header to the allowlist in responses' HTTP headers so it can be accessed by the browser.

have to include the `Access-Control-Allow-Origin` and `Access-Control-Expose-Headers`

- use this lib will add `Access-Control-Expose-Headers` auto by middleware
- so only add `Access-Control-Allow-Origin`
  like use this lib [https://github.com/gin-contrib/cors](https://github.com/gin-contrib/cors)

# dev

```bash
# It needs to be executed after the first use or update of dependencies.
$ make init dep
```

- test code

```bash
$ make test testBenchmark
```

add main.go file and run

```bash
# run at env dev use cmd/main.go
$ make dev
```

- ci to fast check

```bash
# check style at local
$ make style

# run ci at local
$ make ci
```

## docker

```bash
# then test build as test/Dockerfile
$ make dockerTestRestartLatest
# clean test build
$ make dockerTestPruneLatest

# more info see
$ make helpDocker
```
