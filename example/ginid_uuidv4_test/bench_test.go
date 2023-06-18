package ginid_uuidv4_test

import (
	"github.com/bar-counter/gin-correlation-id/gin_correlation_cors"
	"github.com/bar-counter/gin-correlation-id/gin_correlation_id_uuidv4"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func ginPingRouter() *gin.Engine {
	router := gin.New()

	gin_correlation_cors.SetIsSupportCorsHeader(true)
	router.Use(gin_correlation_id_uuidv4.Middleware())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return router
}

func BenchmarkGinIdUUuidV4(b *testing.B) {
	// mock GinIdUUidV4
	router := ginPingRouter()
	// reset counter
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// do GinIdUUidV4
		w := performRequest(router, "GET", "/ping")
		// verify GinIdUUidV4
		assert.Equal(b, http.StatusOK, w.Code)
		assert.NotNilf(b, w.Header().Get(gin_correlation_cors.CorrelationIDHeaderDefault), "want header [ %v ] not empty", gin_correlation_cors.CorrelationIDHeaderDefault)
	}
	b.StopTimer()
}

func BenchmarkParallelGinIdUuidV4(b *testing.B) {
	// mock lGinIdUuidV4
	router := ginPingRouter()
	// reset counter
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			// do GinIdUUidV4
			w := performRequest(router, "GET", "/ping")
			// verify GinIdUUidV4
			assert.Equal(b, http.StatusOK, w.Code)
			assert.NotNilf(b, w.Header().Get(gin_correlation_cors.CorrelationIDHeaderDefault), "want header [ %v ] not empty", gin_correlation_cors.CorrelationIDHeaderDefault)
		}
	})
	b.StopTimer()
}
