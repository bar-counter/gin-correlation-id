package gin_correlation_id_uuidv4

import (
	"fmt"
	"github.com/bar-counter/gin-correlation-id/gin_correlation_cors"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid/v5"
	"strings"
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

func addCorrelationID(c *gin.Context) {

	if gin_correlation_cors.IsSupportCorsHeader() {
		gin_correlation_cors.AllowCorrelationID(c)
	}

	uuidKey := gin_correlation_cors.GetCorrelationIDUuidV4Key()

	correlationID := c.Request.Header.Get(uuidKey)
	if strings.TrimSpace(correlationID) == "" {
		key, err := uuid.NewV4()
		if err != nil {
			panic(fmt.Errorf("can not generate uuid v4: %w", err))
		}
		id := key.String()
		c.Request.Header.Add(uuidKey, id)
		c.Header(uuidKey, id)
	}
	c.Next()
}

// GetCorrelationID
// can get uuid from gin.Context
func GetCorrelationID(c *gin.Context) string {
	return c.Request.Header.Get(gin_correlation_cors.GetCorrelationIDUuidV4Key())
}
