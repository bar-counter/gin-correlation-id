
## 1.0.0

```log
pkg: github.com/bar-counter/gin-correlation-id/example/ginid_uuidv4_test
goos: linux
goarch: amd64
cpu: Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz
Benchmark_gin_correlation_id_uuidv4
Benchmark_gin_correlation_id_uuidv4-2           	  167029	      7124 ns/op	    2320 B/op	      32 allocs/op
BenchmarkParallel_gin_correlation_id_uuidv4
BenchmarkParallel_gin_correlation_id_uuidv4-2   	  218067	      5572 ns/op	    2320 B/op	      32 allocs/op
goos: windows
goarch: amd64
cpu: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz
Benchmark_gin_correlation_id_uuidv4
Benchmark_gin_correlation_id_uuidv4-2           	  215640	      5563 ns/op	    2344 B/op	      33 allocs/op
BenchmarkParallel_gin_correlation_id_uuidv4
BenchmarkParallel_gin_correlation_id_uuidv4-2   	  291274	      3558 ns/op	    2344 B/op	      33 allocs/op
goos: darwin
goarch: arm64
cpu: apple silicon M1 Max core 10
Benchmark_gin_correlation_id_uuidv4
Benchmark_gin_correlation_id_uuidv4-10            	  440373	      2712 ns/op	    2321 B/op	      32 allocs/op
BenchmarkParallel_gin_correlation_id_uuidv4
BenchmarkParallel_gin_correlation_id_uuidv4-10    	  513122	      2296 ns/op	    2324 B/op	      32 allocs/op

pkg: github.com/bar-counter/gin-correlation-id/example/ginid_snowflake_test
goos: linux
goarch: amd64
cpu: Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz
Benchmark_gin_correlation_id_snowflake
Benchmark_gin_correlation_id_snowflake-2           	  199422	      5922 ns/op	    2256 B/op	      31 allocs/op
BenchmarkParallel_gin_correlation_id_snowflake
BenchmarkParallel_gin_correlation_id_snowflake-2   	  260205	      4920 ns/op	    2256 B/op	      31 allocs/op
goos: windows
goarch: amd64
cpu: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz
Benchmark_gin_correlation_id_snowflake
Benchmark_gin_correlation_id_snowflake-2           	  228192	      5059 ns/op	    2256 B/op	      31 allocs/op
BenchmarkParallel_gin_correlation_id_snowflake
BenchmarkParallel_gin_correlation_id_snowflake-2   	  347394	      3327 ns/op	    2256 B/op	      31 allocs/op
goos: darwin
goarch: arm64
cpu: apple silicon M1 Max core 10
Benchmark_gin_correlation_id_snowflake
Benchmark_gin_correlation_id_snowflake-10            	  570752	      2095 ns/op	    2257 B/op	      31 allocs/op
BenchmarkParallel_gin_correlation_id_snowflake
BenchmarkParallel_gin_correlation_id_snowflake-10    	  557127	      2213 ns/op	    2260 B/op	      31 allocs/op


pkg: github.com/bar-counter/gin-correlation-id/example/ginid_shortuuid_test
goos: linux
goarch: amd64
cpu: Intel(R) Xeon(R) CPU E5-2673 v4 @ 2.30GHz
BenchmarkGinIdShortUuid
BenchmarkGinIdShortUuid-2           	   73371	     15670 ns/op	    4891 B/op	     128 allocs/op
BenchmarkParallelGinIdShortUuid
BenchmarkParallelGinIdShortUuid-2   	  103296	     10477 ns/op	    4891 B/op	     128 allocs/op
goos: windows
goarch: amd64
cpu: Intel(R) Xeon(R) Platinum 8272CL CPU @ 2.60GHz
BenchmarkGinIdShortUuid
BenchmarkGinIdShortUuid-2           	   95556	     12331 ns/op	    4915 B/op	     129 allocs/op
BenchmarkParallelGinIdShortUuid
BenchmarkParallelGinIdShortUuid-2   	  153523	      8331 ns/op	    4915 B/op	     129 allocs/op
goos: darwin
goarch: arm64
cpu: apple silicon M1 Max core 10
BenchmarkGinIdShortUuid
BenchmarkGinIdShortUuid-10            	  202786	      5920 ns/op	    4893 B/op	     128 allocs/op
BenchmarkParallelGinIdShortUuid
BenchmarkParallelGinIdShortUuid-10    	  330234	      3614 ns/op	    4898 B/op	     128 allocs/op
```