package gin_correlation_cors

import "fmt"

var correlationIDSnowflakeKey = CorrelationIDHeaderDefault

// SetCorrelationIDSnowflakeKey
// default will use gin_correlation_cors.CorrelationIDHeaderDefault
func SetCorrelationIDSnowflakeKey(key string) {
	if key == "" {
		panic(fmt.Errorf("can not SetCorrelationIDSnowflakeKey set by empty"))
	}
	correlationIDSnowflakeKey = key
}

// GetCorrelationIDSnowflakeKey
// return now correlation ID snowflake Key
func GetCorrelationIDSnowflakeKey() string {
	return correlationIDSnowflakeKey
}
