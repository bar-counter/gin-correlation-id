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

## libs

| lib                                 | version |
|:------------------------------------|:--------|
| https://github.com/stretchr/testify | v1.8.4  |
| https://github.com/gin-gonic/gin    | v1.9.1  |
| https://github.com/gofrs/uuid/v5    | v5.0.0  |

## evn

- golang sdk 1.17+

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

