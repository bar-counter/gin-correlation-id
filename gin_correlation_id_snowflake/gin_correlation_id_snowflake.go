package gin_correlation_id_snowflake

import (
	"fmt"
	"github.com/bar-counter/gin-correlation-id/gin_correlation_cors"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
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

var snowflakeNode *snowflake.Node

var snowflakeNameSpace = int64(1)

// SetSnowflakeNameSpace
// default namespace is 1
func SetSnowflakeNameSpace(nameSpace int64) {
	snowflakeNameSpace = nameSpace
}

func addCorrelationID(c *gin.Context) {

	if gin_correlation_cors.IsSupportCorsHeader() {
		gin_correlation_cors.AllowCorrelationID(c)
	}

	uuidKey := gin_correlation_cors.GetCorrelationIDSnowflakeKey()

	correlationID := c.Request.Header.Get(uuidKey)
	if strings.TrimSpace(correlationID) == "" {

		if snowflakeNode == nil {
			sNode, errSnowflake := snowflake.NewNode(snowflakeNameSpace)
			if errSnowflake != nil {
				panic(fmt.Errorf("can not new snowflake node by count %d err: %w", snowflakeNameSpace, errSnowflake))
			}
			snowflakeNode = sNode
		}

		id := snowflakeNode.Generate().String()
		c.Request.Header.Add(uuidKey, id)
		c.Header(uuidKey, id)
	}
	c.Next()
}

// GetCorrelationID
// can get uuid from gin.Context
func GetCorrelationID(c *gin.Context) string {
	return c.Request.Header.Get(gin_correlation_cors.GetCorrelationIDSnowflakeKey())
}
