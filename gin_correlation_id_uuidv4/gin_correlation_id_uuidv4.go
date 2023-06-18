package gin_correlation_id_uuidv4

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"strings"
)

const (

	// CorrelationIDHeaderDefault
	// this is default header key
	CorrelationIDHeaderDefault = "x-request-id"

	headerAccessControlExposeHeaders = "Access-Control-Expose-Headers"
)

// Middleware
// adds correlationID if it's not specified in HTTP request.
// this will add `Access-Control-Expose-Headers` with default `x-request-id`.
// If you are using CORS, you also have to include the Access-Control-Allow-Origin.
//
// For more details, see the https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
func Middleware() gin.HandlerFunc {
	return addCorrelationID
}

var correlationIDUuidV4Key = CorrelationIDHeaderDefault

// SetCorrelationIDUuidV4Key
// default will use `x-request-id`
func SetCorrelationIDUuidV4Key(key string) {
	if key == "" {
		panic(fmt.Errorf("can not SetCorrelationIDUuidV4Key set by empty"))
	}
	correlationIDUuidV4Key = key
}

func addCorrelationID(c *gin.Context) {

	c.Header(headerAccessControlExposeHeaders, correlationIDUuidV4Key)

	correlationID := c.Request.Header.Get(correlationIDUuidV4Key)

	if strings.TrimSpace(correlationID) == "" {
		key, err := uuid.NewV4()
		if err != nil {
			panic(fmt.Errorf("can not generate uuid v4: %w", err))
		}
		id := key.String()
		c.Request.Header.Add(correlationIDUuidV4Key, id)
		c.Header(correlationIDUuidV4Key, id)
	}
	c.Next()
}

// GetCorrelationID
// can get uuid from gin.Context
func GetCorrelationID(c *gin.Context) string {
	return c.Request.Header.Get(correlationIDUuidV4Key)
}
