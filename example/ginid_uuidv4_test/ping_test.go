package ginid_uuidv4_test

import (
	"encoding/json"
	"github.com/bar-counter/gin-correlation-id/ginid_uuidv4"
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
	ginid_uuidv4.SetCorrelationIDUUidV4Key(ginid_uuidv4.CorrelationIDUUidV4HeaderDefault)
	router := gin.New()
	router.Use(ginid_uuidv4.CorrelationIDUUidV4Middleware())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			ginid_uuidv4.CorrelationIDUUidV4HeaderDefault: ginid_uuidv4.GetCorrelationIDUUidV4(c),
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

	value, exists := response[ginid_uuidv4.CorrelationIDUUidV4HeaderDefault]

	assert.True(t, exists)
	assert.NotNil(t, value)
}
