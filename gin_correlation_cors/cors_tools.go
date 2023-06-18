package gin_correlation_cors

import (
	"github.com/gin-gonic/gin"
	"sort"
	"strings"
)

func AllowCorrelationID(c *gin.Context) {
	keys := []string{
		correlationIDUuidV4Key,
		correlationIDSnowflakeKey,
		correlationIDShortUuidKey,
	}
	keys = deduplicateStringSlice(keys)
	keys = removeWhitespaceAsStringSlice(keys)

	accessControlSet := c.Writer.Header().Get(HeaderAccessControlExposeHeaders)
	if accessControlSet == "" {
		val := strings.Join(keys, ", ")
		val = strings.TrimRight(val, ", ")
		c.Header(HeaderAccessControlExposeHeaders, val)
	} else {
		accessControlSlice := strings.Split(accessControlSet, ",")
		accessControlSlice = deduplicateStringSlice(accessControlSlice)
		accessControlSlice = append(accessControlSlice, keys...)
		accessControlSlice = deduplicateStringSlice(accessControlSlice)
		accessControlSlice = removeWhitespaceAsStringSlice(accessControlSlice)
		val := strings.Join(accessControlSlice, ", ")
		val = strings.TrimRight(val, ", ")
		c.Header(HeaderAccessControlExposeHeaders, val)
	}
}

func deduplicateStringSlice(stringSlice []string) []string {
	if len(stringSlice) == 0 {
		return stringSlice
	}
	sort.Strings(stringSlice)
	var deduplicatedStringSlice []string
	for i := 0; i < len(stringSlice)-1; i++ {
		if stringSlice[i] != stringSlice[i+1] {
			deduplicatedStringSlice = append(deduplicatedStringSlice, stringSlice[i])
		}
	}
	return append(deduplicatedStringSlice, stringSlice[len(stringSlice)-1])
}

func removeWhitespaceAsStringSlice(stringSlice []string) []string {
	if len(stringSlice) == 0 {
		return stringSlice
	}
	var deduplicatedStringSlice []string
	for _, str := range stringSlice {
		str = strings.TrimSpace(str)
		deduplicatedStringSlice = append(deduplicatedStringSlice, str)
	}
	return deduplicatedStringSlice
}
