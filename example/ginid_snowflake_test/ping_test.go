package ginid_uuidv4_test

import (
	"encoding/json"
	"fmt"
	"github.com/bar-counter/gin-correlation-id/gin_correlation_cors"
	"github.com/bar-counter/gin-correlation-id/gin_correlation_id_snowflake"
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

func ginPingJsonRouter(correlationKey string, isSupportCorsHeader bool, namespace int64) *gin.Engine {
	gin_correlation_cors.SetIsSupportCorsHeader(isSupportCorsHeader)
	gin_correlation_cors.SetCorrelationIDSnowflakeKey(correlationKey)
	gin_correlation_id_snowflake.SetSnowflakeNameSpace(namespace)
	router := gin.New()
	router.Use(gin_correlation_id_snowflake.Middleware())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			gin_correlation_cors.CorrelationIDHeaderDefault: gin_correlation_id_snowflake.GetCorrelationID(c),
		})
	})
	return router
}

func TestPanicSetCorrelationIDSnowflakeKey(t *testing.T) {
	// mock TestPanicSetCorrelationIdUuidV4Key

	errString := "can not SetCorrelationIDSnowflakeKey set by empty"
	if !assert.PanicsWithError(t, errString, func() {
		// do TestPanicSetCorrelationIdUuidV4Key
		gin_correlation_cors.SetCorrelationIDSnowflakeKey("")
	}) {
		// verify TestPanicSetCorrelationIdUuidV4Key
		t.Fatalf("TestPanicSetCorrelationIdUuidV4Key should panic")
	}
}

func TestPingOnce(t *testing.T) {
	// mock PingOnce

	t.Logf("~> mock PingOnce")
	router := ginPingJsonRouter(gin_correlation_cors.CorrelationIDHeaderDefault, true, 1)
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
	router := ginPingJsonRouter(gin_correlation_cors.CorrelationIDHeaderDefault, false, 1)
	// do CloseCors
	t.Logf("~> do CloseCors")
	w := performRequest(router, "GET", "/ping")
	// verify PingOnce
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Header().Get(gin_correlation_cors.CorrelationIDHeaderDefault))
	// verify CloseCors
	assert.Equal(t, "", w.Header().Get(gin_correlation_cors.HeaderAccessControlExposeHeaders))
}

func TestSetCorrelationIDSnowflakeKey(t *testing.T) {
	// mock SetCorrelationIDSnowflakeKey
	t.Logf("~> mock SetCorrelationIDSnowflakeKey")
	router := ginPingJsonRouter(gin_correlation_cors.KeyCorrelationIDHeaderId, true, 1)
	// do SetCorrelationIDSnowflakeKey
	t.Logf("~> do SetCorrelationIDSnowflakeKey")
	w := performRequest(router, "GET", "/ping")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Header().Get(gin_correlation_cors.KeyCorrelationIDHeaderId))
	// verify SetCorrelationIDSnowflakeKey
	assert.Equal(t, fmt.Sprintf("%s, %s",
		gin_correlation_cors.KeyCorrelationIDHeaderId, gin_correlation_cors.CorrelationIDHeaderDefault),
		w.Header().Get(gin_correlation_cors.HeaderAccessControlExposeHeaders))
}

func TestSetSnowflakeNameSpace(t *testing.T) {
	// mock SetSnowflakeNameSpace
	t.Logf("~> mock SetSnowflakeNameSpace")

	newNs := int64(2)
	router := ginPingJsonRouter(gin_correlation_cors.CorrelationIDHeaderDefault, true, newNs)

	// do SetSnowflakeNameSpace
	t.Logf("~> do SetSnowflakeNameSpace")
	w := performRequest(router, "GET", "/ping")

	// verify SetSnowflakeNameSpace
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNil(t, w.Header().Get(gin_correlation_cors.CorrelationIDHeaderDefault))

	var response map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &response)

	assert.Nil(t, err)

	value, exists := response[gin_correlation_cors.CorrelationIDHeaderDefault]

	t.Logf("~> verify PingOnce %s", value)

	assert.True(t, exists)
	assert.NotNil(t, value)
	t.Logf("~> verify SetSnowflakeNameSpace [ %d ] and [ %s ] value: %s", newNs, gin_correlation_cors.CorrelationIDHeaderDefault, value)
}

func TestPanicSetSnowflakeMode(t *testing.T) {
	// mock TestPanicSetSnowflakeMode

	errString := "mode must be between 1 and 6"
	if !assert.PanicsWithError(t, errString, func() {
		// do TestPanicSetSnowflakeMode
		gin_correlation_id_snowflake.SetSnowflakeMode(1024)
	}) {
		// verify TestPanicSetSnowflakeMode
		t.Fatalf("TestPanicSetSnowflakeMode should panic")
	}
	if !assert.PanicsWithError(t, errString, func() {
		// do TestPanicSetSnowflakeMode
		gin_correlation_id_snowflake.SetSnowflakeMode(0)
	}) {
		// verify TestPanicSetSnowflakeMode
		t.Fatalf("TestPanicSetSnowflakeMode should panic")
	}
}

func TestSetSnowflakeModeBase58(t *testing.T) {
	t.Logf("~> mock SetSnowflakeModeBase58")
	// mock SetSnowflakeModeBase58
	router := ginPingRouter(gin_correlation_id_snowflake.SnowflakeModeBase58)

	t.Logf("~> do SetSnowflakeModeBase58")
	// do SetSnowflakeModeBase58
	w := performRequest(router, "GET", "/ping")
	// verify SetSnowflakeModeBase58
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNilf(t, w.Header().Get(gin_correlation_cors.CorrelationIDHeaderDefault), "want header [ %v ] not empty", gin_correlation_cors.CorrelationIDHeaderDefault)
	content := w.Body.String()
	t.Logf("~> verify SetSnowflakeModeBase58 [ %s ]", content)
	assert.Equal(t, 35,
		len(content))
}

func TestSetSnowflakeModeBase36(t *testing.T) {
	t.Logf("~> mock SetSnowflakeModeBase36")
	// mock SetSnowflakeModeBase36
	router := ginPingRouter(gin_correlation_id_snowflake.SnowflakeModeBase36)

	t.Logf("~> do SetSnowflakeModeBase36")
	// do SetSnowflakeModeBase36
	w := performRequest(router, "GET", "/ping")

	// verify SetSnowflakeModeBase36
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNilf(t, w.Header().Get(gin_correlation_cors.CorrelationIDHeaderDefault), "want header [ %v ] not empty", gin_correlation_cors.CorrelationIDHeaderDefault)
	content := w.Body.String()
	t.Logf("~> verify SetSnowflakeModeBase36 [ %s ]", content)
	assert.Equal(t, 36,
		len(content))
}

func TestSetSnowflakeModeBase32(t *testing.T) {
	t.Logf("~> mock SetSnowflakeModeBase32")
	// mock SetSnowflakeModeBase32
	router := ginPingRouter(gin_correlation_id_snowflake.SnowflakeModeBase32)

	t.Logf("~> do SetSnowflakeModeBase32")
	// do SetSnowflakeModeBase32
	w := performRequest(router, "GET", "/ping")

	// verify SetSnowflakeModeBase32
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNilf(t, w.Header().Get(gin_correlation_cors.CorrelationIDHeaderDefault), "want header [ %v ] not empty", gin_correlation_cors.CorrelationIDHeaderDefault)
	content := w.Body.String()
	t.Logf("~> verify SetSnowflakeModeBase32 [ %s ]", content)
	assert.Equal(t, 37,
		len(content))
}

func TestSetSnowflakeModeBase2(t *testing.T) {
	t.Logf("~> mock SetSnowflakeModeBase2")
	// mock SetSnowflakeModeBase2
	router := ginPingRouter(gin_correlation_id_snowflake.SnowflakeModeBase2)

	t.Logf("~> do SetSnowflakeModeBase2")
	// do SetSnowflakeModeBase2
	w := performRequest(router, "GET", "/ping")

	// verify SetSnowflakeModeBase2
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNilf(t, w.Header().Get(gin_correlation_cors.CorrelationIDHeaderDefault), "want header [ %v ] not empty", gin_correlation_cors.CorrelationIDHeaderDefault)
	content := w.Body.String()
	t.Logf("~> verify SetSnowflakeModeBase2 [ %s ]", content)
	assert.Equal(t, 85,
		len(content))
}

func TestSetSnowflakeModeBase64(t *testing.T) {
	t.Logf("~> mock SetSnowflakeModeBase64")
	// mock SetSnowflakeModeBase64
	router := ginPingRouter(gin_correlation_id_snowflake.SnowflakeModeBase64)

	t.Logf("~> do SetSnowflakeModeBase64")
	// do SetSnowflakeModeBase64
	w := performRequest(router, "GET", "/ping")

	// verify SetSnowflakeModeBase64
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNilf(t, w.Header().Get(gin_correlation_cors.CorrelationIDHeaderDefault), "want header [ %v ] not empty", gin_correlation_cors.CorrelationIDHeaderDefault)
	content := w.Body.String()
	t.Logf("~> verify SetSnowflakeModeBase64 [ %s ]", content)
	assert.NotEqual(t, "",
		content)
}

func TestSetSnowflakeModeInt64(t *testing.T) {
	t.Logf("~> mock SetSnowflakeModeInt64")
	// mock SetSnowflakeModeInt64
	router := ginPingRouter(gin_correlation_id_snowflake.SnowflakeModeInt64)

	t.Logf("~> do SetSnowflakeModeInt64")
	// do SetSnowflakeModeInt64
	w := performRequest(router, "GET", "/ping")

	// verify SetSnowflakeModeInt64
	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotNilf(t, w.Header().Get(gin_correlation_cors.CorrelationIDHeaderDefault), "want header [ %v ] not empty", gin_correlation_cors.CorrelationIDHeaderDefault)
	content := w.Body.String()
	t.Logf("~> verify SetSnowflakeModeInt64 [ %s ]", content)
	assert.NotEqual(t, "",
		content)
}
