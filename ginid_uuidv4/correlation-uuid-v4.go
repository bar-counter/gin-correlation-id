package ginid_uuidv4

import (
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"strings"
)

const (

	// CorrelationIDUUidV4HeaderDefault
	// this is default header key
	CorrelationIDUUidV4HeaderDefault = "x-request-id"

	headerAccessControlExposeHeaders = "Access-Control-Expose-Headers"
)

// CorrelationIDUUidV4Middleware
// adds correlationID if it's not specified in HTTP request.
// this will add `Access-Control-Expose-Headers` with default `x-request-id`.
// If you are using CORS, you also have to include the Access-Control-Allow-Origin.
//
// For more details, see the https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
func CorrelationIDUUidV4Middleware() gin.HandlerFunc {
	return addCorrelationID
}

var correlationIDUUidV4Key = CorrelationIDUUidV4HeaderDefault

// SetCorrelationIDUUidV4Key
// default will use `x-request-id`
func SetCorrelationIDUUidV4Key(key string) {
	if key == "" {
		panic("can not SetCorrelationIDUUidV4Key set by empty")
	}
	correlationIDUUidV4Key = key
}

func addCorrelationID(c *gin.Context) {

	c.Header(headerAccessControlExposeHeaders, correlationIDUUidV4Key)

	correlationID := c.Request.Header.Get(correlationIDUUidV4Key)

	if strings.TrimSpace(correlationID) == "" {
		key := uuid.NewV4()
		id := key.String()
		c.Request.Header.Add(correlationIDUUidV4Key, id)
		c.Header(correlationIDUUidV4Key, id)
	}
	c.Next()
}

// GetCorrelationIDUUidV4
// can get uuid from gin.Context
func GetCorrelationIDUUidV4(c *gin.Context) string {
	return c.Request.Header.Get(correlationIDUUidV4Key)
}
