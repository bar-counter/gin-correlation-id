package gin_correlation_cors

const (
	// CorrelationIDHeaderDefault
	// this is default header key
	CorrelationIDHeaderDefault = "x-request-id"

	// KeyCorrelationIDHeaderId
	// this is header key
	KeyCorrelationIDHeaderId = "x-correlation-id"

	// HeaderAccessControlExposeHeaders
	// for CORS
	HeaderAccessControlExposeHeaders = "Access-Control-Expose-Headers"
)

var isSupportCorsHeader = true

// SetIsSupportCorsHeader
// default will use true
func SetIsSupportCorsHeader(support bool) {
	isSupportCorsHeader = support
}

// IsSupportCorsHeader
// now is support cors header
func IsSupportCorsHeader() bool {
	return isSupportCorsHeader
}
