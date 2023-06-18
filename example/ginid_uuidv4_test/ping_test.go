package ginid_uuidv4_test

import (
	"encoding/json"
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

func ginPingJsonRouter() *gin.Engine {
	gin_correlation_id_uuidv4.SetCorrelationIDUuidV4Key(gin_correlation_id_uuidv4.CorrelationIDHeaderDefault)
	router := gin.New()
	router.Use(gin_correlation_id_uuidv4.Middleware())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			gin_correlation_id_uuidv4.CorrelationIDHeaderDefault: gin_correlation_id_uuidv4.GetCorrelationID(c),
		})
	})
	return router
}

func TestPingOnce(t *testing.T) {
	// mock PingOnce

	t.Logf("~> mock PingOnce")
	router := ginPingJsonRouter()
	// do PingOnce
	t.Logf("~> do PingOnce")
	w := performRequest(router, "GET", "/ping")
	// verify PingOnce
	assert.Equal(t, http.StatusOK, w.Code)

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Nil(t, err)

	value, exists := response[gin_correlation_id_uuidv4.CorrelationIDHeaderDefault]

	t.Logf("~> verify PingOnce %s", value)

	assert.True(t, exists)
	assert.NotNil(t, value)
}

func TestPanicSetCorrelationIdUuidV4Key(t *testing.T) {
	// mock TestPanicSetCorrelationIdUuidV4Key

	errString := "can not SetCorrelationIDUuidV4Key set by empty"
	if !assert.PanicsWithError(t, errString, func() {
		// do TestPanicSetCorrelationIdUuidV4Key
		gin_correlation_id_uuidv4.SetCorrelationIDUuidV4Key("")
	}) {
		// verify TestPanicSetCorrelationIdUuidV4Key
		t.Fatalf("TestPanicSetCorrelationIdUuidV4Key should panic")
	}
}
