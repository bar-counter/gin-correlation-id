package gin_correlation_cors

import "fmt"

var correlationIDUuidV4Key = CorrelationIDHeaderDefault

// SetCorrelationIDUuidV4Key
// default will use gin_correlation_cors.CorrelationIDHeaderDefault
func SetCorrelationIDUuidV4Key(key string) {
	if key == "" {
		panic(fmt.Errorf("can not SetCorrelationIDUuidV4Key set by empty"))
	}
	correlationIDUuidV4Key = key
}

// GetCorrelationIDUuidV4Key
// return now correlation ID uuid v4 Key
func GetCorrelationIDUuidV4Key() string {
	return correlationIDUuidV4Key
}
