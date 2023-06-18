[![golang-ci](https://github.com/bar-counter/gin-correlation-id/actions/workflows/golang-ci.yml/badge.svg)](https://github.com/bar-counter/gin-correlation-id/actions/workflows/golang-ci.yml)
[![go mod version](https://img.shields.io/github/go-mod/go-version/bar-counter/gin-correlation-id?label=go.mod)](https://github.com/bar-counter/gin-correlation-id)
[![GoDoc](https://godoc.org/github.com/bar-counter/gin-correlation-id?status.png)](https://godoc.org/github.com/bar-counter/gin-correlation-id/)
[![GoReportCard](https://goreportcard.com/badge/github.com/bar-counter/gin-correlation-id)](https://goreportcard.com/report/github.com/bar-counter/gin-correlation-id)
[![codecov](https://codecov.io/gh/bar-counter/gin-correlation-id/branch/main/graph/badge.svg)](https://codecov.io/gh/bar-counter/gin-correlation-id)
[![github release](https://img.shields.io/github/v/release/bar-counter/gin-correlation-id?style=social)](https://github.com/bar-counter/gin-correlation-id/releases)

## for what

- this project used to github golang lib project

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

| lib                                 | version |
|:------------------------------------|:--------|
| https://github.com/stretchr/testify | v1.8.4  |
| https://github.com/gin-gonic/gin    | v1.9.1  |
| https://github.com/gofrs/uuid/v5    | v5.0.0  |

## Feature

- id creation method
  - [x] uuid-v4 by [https://github.com/gofrs/uuid](https://github.com/gofrs/uuid)
  - [ ] snowflake
- [CORS](#CORS) cross-origin resource sharing

## usage

### uuid-v4

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

### integrationTest

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
    gin_correlation_id_uuidv4.GetCorrelationID(*gin.Context)
```

# CORS

If you are using cross-origin resource sharing (CORS)

> e.g. you are making requests to an API from a frontend JavaScript code served from a different origin, you have to ensure two things:

- permit correlation ID header in the incoming requests' HTTP headers so the value can be reused by the middleware,
- add the correlation ID header to the allowlist in responses' HTTP headers so it can be accessed by the browser.

have to include the `Access-Control-Allow-Origin` and `Access-Control-Expose-Headers`

- use api will add `Access-Control-Expose-Headers` auto
- so only add `Access-Control-Allow-Origin` and [https://github.com/gin-contrib/cors](https://github.com/gin-contrib/cors) 

# dev

```bash
$ make init dep
```

- test code

```bash
$ make test testBenchmark
```

add main.go file and run

```bash
# run at env dev
$ make dev

# run at env ordinary
$ make run
```

- ci to fast check

```bash
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

## use

- use to replace
  `bar-counter/gin-correlation-id` to you code

