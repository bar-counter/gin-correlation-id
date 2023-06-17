package ginid_uuidv4_test

import (
	"github.com/bar-counter/gin-correlation-id/ginid_uuidv4"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func ginPingRouter() *gin.Engine {
	router := gin.New()

	router.Use(ginid_uuidv4.CorrelationIDUUidV4Middleware())

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
		assert.NotNilf(b, w.Header().Get(ginid_uuidv4.CorrelationIDUUidV4HeaderDefault), "want header [ %v ] not empty", ginid_uuidv4.CorrelationIDUUidV4HeaderDefault)
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
			assert.NotNilf(b, w.Header().Get(ginid_uuidv4.CorrelationIDUUidV4HeaderDefault), "want header [ %v ] not empty", ginid_uuidv4.CorrelationIDUUidV4HeaderDefault)
		}
	})
	b.StopTimer()
}
