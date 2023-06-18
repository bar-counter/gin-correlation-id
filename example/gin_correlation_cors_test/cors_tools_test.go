package gin_correlation_cors_test

import (
	"github.com/bar-counter/gin-correlation-id/gin_correlation_cors"
	"github.com/bar-counter/gin-correlation-id/gin_correlation_id_uuidv4"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func ginPingRouter() *gin.Engine {
	router := gin.New()

	router.Use(CorrelationIDMiddleware())
	gin_correlation_cors.SetIsSupportCorsHeader(true)
	router.Use(gin_correlation_id_uuidv4.Middleware())

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	return router
}

func CorrelationIDMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Header(gin_correlation_cors.HeaderAccessControlExposeHeaders, "x-user, x-api")

		context.Next()
	}
}

func TestAllowCorrelationID(t *testing.T) {
	// mock AllowCorrelationID

	t.Logf("~> mock AllowCorrelationID")
	router := ginPingRouter()
	// do AllowCorrelationID

	t.Logf("~> do AllowCorrelationID")
	w := performRequest(router, "GET", "/ping")
	// verify AllowCorrelationID
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Header().Get(gin_correlation_cors.CorrelationIDHeaderDefault))
	t.Logf("HeaderAccessControlExposeHeaders val: %s", w.Header().Get(gin_correlation_cors.HeaderAccessControlExposeHeaders))
	assert.Equal(t, "x-api, x-request-id, x-user", w.Header().Get(gin_correlation_cors.HeaderAccessControlExposeHeaders))
}
