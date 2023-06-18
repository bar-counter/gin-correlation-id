package gin_correlation_id_snowflake

import (
	"fmt"
	"github.com/bar-counter/gin-correlation-id/gin_correlation_cors"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	SnowflakeModeBase58 = iota + 1
	SnowflakeModeBase36
	SnowflakeModeBase32
	SnowflakeModeBase2
	SnowflakeModeBase64
	SnowflakeModeInt64
)

var snowflakeMode = SnowflakeModeBase58

// SetSnowflakeMode
// default mode is SnowflakeModeBase58
func SetSnowflakeMode(mode uint) {
	if mode < SnowflakeModeBase58 || mode > SnowflakeModeInt64 {
		panic(fmt.Errorf("mode must be between 1 and 6"))
	}
	snowflakeMode = int(mode)
}

var snowflakeNode *snowflake.Node

var snowflakeNameSpace = int64(1)

// SetSnowflakeNameSpace
// default namespace is 1
func SetSnowflakeNameSpace(nameSpace int64) {
	snowflakeNameSpace = nameSpace
}

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

		var id string
		switch snowflakeMode {
		default:
			id = snowflakeNode.Generate().Base58()
		case SnowflakeModeBase58:
			id = snowflakeNode.Generate().Base58()
		case SnowflakeModeBase36:
			id = snowflakeNode.Generate().Base36()
		case SnowflakeModeBase32:
			id = snowflakeNode.Generate().Base32()
		case SnowflakeModeBase2:
			id = snowflakeNode.Generate().Base2()
		case SnowflakeModeBase64:
			id = snowflakeNode.Generate().Base64()
		case SnowflakeModeInt64:
			id = snowflakeNode.Generate().String()
		}

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
