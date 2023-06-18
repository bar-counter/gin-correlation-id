package gin_correlation_cors

import "fmt"

var correlationIDShortUuidKey = CorrelationIDHeaderDefault

// SetCorrelationIDShortUuidKey
// default will use gin_correlation_cors.CorrelationIDHeaderDefault
func SetCorrelationIDShortUuidKey(key string) {
	if key == "" {
		panic(fmt.Errorf("can not SetCorrelationIDShortUuidKey set by empty"))
	}
	correlationIDShortUuidKey = key
}

// GetCorrelationIDShortUuidKey
// return now correlation ID short uuid Key
func GetCorrelationIDShortUuidKey() string {
	return correlationIDShortUuidKey
}
