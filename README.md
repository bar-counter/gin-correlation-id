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

## evn

- golang sdk 1.17+

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

```log
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i9-10910 CPU @ 3.60GHz

pkg: github.com/bar-counter/gin-correlation-id/example/ginid_uuidv4_test
Benchmark_gin_correlation_id_uuidv4
Benchmark_gin_correlation_id_uuidv4-20                    244521              4945 ns/op            2323 B/op         32 allocs/op
BenchmarkParallel_gin_correlation_id_uuidv4
BenchmarkParallel_gin_correlation_id_uuidv4-20            480244              2454 ns/op            2330 B/op         32 allocs/op

pkg: github.com/bar-counter/gin-correlation-id/example/ginid_snowflake_test
Benchmark_gin_correlation_id_snowflake
Benchmark_gin_correlation_id_snowflake-20                 344289              3269 ns/op            2259 B/op         31 allocs/op
BenchmarkParallel_gin_correlation_id_snowflake
BenchmarkParallel_gin_correlation_id_snowflake-20         492865              2339 ns/op            2265 B/op         31 allocs/op

pkg: github.com/bar-counter/gin-correlation-id/example/ginid_shortuuid_test
BenchmarkGinIdShortUuid
BenchmarkGinIdShortUuid-20                122689              9716 ns/op            4896 B/op        128 allocs/op
BenchmarkParallelGinIdShortUuid
BenchmarkParallelGinIdShortUuid-20        272950              4430 ns/op            4914 B/op        128 allocs/op
```

## usage

- see full api usage at unit test case
  - [gin_correlation_id_uuidv4](https://github.com/bar-counter/gin-correlation-id/blob/main/example/ginid_uuidv4_test/ping_test.go)
  - [gin_correlation_id_snowflake](https://github.com/bar-counter/gin-correlation-id/blob/main/example/ginid_snowflake_test/ping_test.go)
  - [gin_correlation_id_shortuuid](https://github.com/bar-counter/gin-correlation-id/blob/main/example/ginid_shortuuid_test/ping_test.go)

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
