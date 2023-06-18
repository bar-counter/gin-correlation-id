package ginid_uuidv4_test

import (
	"github.com/bar-counter/gin-correlation-id/gin_correlation_cors"
	"github.com/bar-counter/gin-correlation-id/gin_correlation_id_snowflake"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"testing"
)

func ginPingRouter(mode uint) *gin.Engine {
	router := gin.New()

	gin_correlation_cors.SetIsSupportCorsHeader(true)
	gin_correlation_cors.SetCorrelationIDSnowflakeKey(gin_correlation_cors.CorrelationIDHeaderDefault)
	gin_correlation_id_snowflake.SetSnowflakeNameSpace(1)
	gin_correlation_id_snowflake.SetSnowflakeMode(mode)
	router.Use(gin_correlation_id_snowflake.Middleware())

	router.GET("/ping", func(c *gin.Context) {
		var b strings.Builder
		b.Grow(0)
		b.WriteString("pong as correlation ID: ")
		b.WriteString(gin_correlation_id_snowflake.GetCorrelationID(c))
		c.String(http.StatusOK, b.String())
	})
	return router
}

func Benchmark_gin_correlation_id_snowflake(b *testing.B) {
	// mock _gin_correlation_id_snowflake
	router := ginPingRouter(gin_correlation_id_snowflake.SnowflakeModeBase58)
	// reset counter
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// do _gin_correlation_id_snowflake
		w := performRequest(router, "GET", "/ping")
		// verify _gin_correlation_id_snowflake
		assert.Equal(b, http.StatusOK, w.Code)
		assert.NotNilf(b, w.Header().Get(gin_correlation_cors.CorrelationIDHeaderDefault), "want header [ %v ] not empty", gin_correlation_cors.CorrelationIDHeaderDefault)
	}
	b.StopTimer()
}

func BenchmarkParallel_gin_correlation_id_snowflake(b *testing.B) {
	// mock Parallel_gin_correlation_id_snowflake
	router := ginPingRouter(gin_correlation_id_snowflake.SnowflakeModeBase58)
	// reset counter
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// do Parallel_gin_correlation_id_snowflake
			w := performRequest(router, "GET", "/ping")
			// verify Parallel_gin_correlation_id_snowflake
			assert.Equal(b, http.StatusOK, w.Code)
			assert.NotNilf(b, w.Header().Get(gin_correlation_cors.CorrelationIDHeaderDefault), "want header [ %v ] not empty", gin_correlation_cors.CorrelationIDHeaderDefault)
		}
	})
	b.StopTimer()
}
