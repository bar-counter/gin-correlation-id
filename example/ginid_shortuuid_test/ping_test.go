package ginid_uuidv4_test

import (
	"encoding/json"
	"fmt"
	"github.com/bar-counter/gin-correlation-id/gin_correlation_cors"
	"github.com/bar-counter/gin-correlation-id/gin_correlation_id_shortuuid"
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

func ginPingJsonRouter(correlationKey string, isSupportCorsHeader bool) *gin.Engine {
	gin_correlation_cors.SetIsSupportCorsHeader(isSupportCorsHeader)
	gin_correlation_cors.SetCorrelationIDShortUuidKey(correlationKey)
	router := gin.New()
	router.Use(gin_correlation_id_shortuuid.Middleware())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			gin_correlation_cors.CorrelationIDHeaderDefault: gin_correlation_id_shortuuid.GetCorrelationID(c),
		})
	})
	return router
}

func TestPanicSetCorrelationIDShortUuidKey(t *testing.T) {
	// mock TestPanicSetCorrelationIDShortUuidKey

	errString := "can not SetCorrelationIDShortUuidKey set by empty"
	if !assert.PanicsWithError(t, errString, func() {
		// do TestPanicSetCorrelationIDShortUuidKey
		gin_correlation_cors.SetCorrelationIDShortUuidKey("")
	}) {
		// verify TestPanicSetCorrelationIDShortUuidKey
		t.Fatalf("TestPanicSetCorrelationIDShortUuidKey should panic")
	}
}

func TestPingOnce(t *testing.T) {
	// mock PingOnce

	t.Logf("~> mock PingOnce")
	router := ginPingJsonRouter(gin_correlation_cors.CorrelationIDHeaderDefault, true)
	// do PingOnce
	t.Logf("~> do PingOnce")
	w := performRequest(router, "GET", "/ping")
	// verify PingOnce
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Header().Get(gin_correlation_cors.CorrelationIDHeaderDefault))

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Nil(t, err)

	value, exists := response[gin_correlation_cors.CorrelationIDHeaderDefault]

	t.Logf("~> verify PingOnce %s", value)

	assert.True(t, exists)
	assert.NotNil(t, value)
}

func TestTestCloseCors(t *testing.T) {
	// mock CloseCors
	t.Logf("~> mock CloseCors")

	// do CloseCors
	router := ginPingJsonRouter(gin_correlation_cors.CorrelationIDHeaderDefault, false)
	// do CloseCors
	t.Logf("~> do CloseCors")
	w := performRequest(router, "GET", "/ping")
	// verify PingOnce
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Header().Get(gin_correlation_cors.CorrelationIDHeaderDefault))
	// verify CloseCors
	assert.Equal(t, "", w.Header().Get(gin_correlation_cors.HeaderAccessControlExposeHeaders))
}

func TestSetCorrelationIDUuidV4Key(t *testing.T) {
	// mock SetCorrelationIDUuidV4Key
	t.Logf("~> mock SetCorrelationIDUuidV4Key")
	router := ginPingJsonRouter(gin_correlation_cors.KeyCorrelationIDHeaderId, true)
	// do SetCorrelationIDUuidV4Key
	t.Logf("~> do SetCorrelationIDUuidV4Key")
	w := performRequest(router, "GET", "/ping")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Header().Get(gin_correlation_cors.KeyCorrelationIDHeaderId))
	// verify SetCorrelationIDUuidV4Key
	assert.Equal(t, fmt.Sprintf("%s, %s",
		gin_correlation_cors.KeyCorrelationIDHeaderId, gin_correlation_cors.CorrelationIDHeaderDefault),
		w.Header().Get(gin_correlation_cors.HeaderAccessControlExposeHeaders))
}
