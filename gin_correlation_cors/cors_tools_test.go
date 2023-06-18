package gin_correlation_cors

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_deduplicateStringSlice(t *testing.T) {
	// mock deduplicateStringSlice

	t.Logf("~> mock deduplicateStringSlice")
	// do deduplicateStringSlice
	t.Logf("~> do deduplicateStringSlice")
	slice := deduplicateStringSlice([]string{})
	// verify deduplicateStringSlice
	assert.Equal(t, 0, len(slice))
}

func Test_removeWhitespaceAsStringSlice(t *testing.T) {
	// mock removeWhitespaceAsStringSlice

	t.Logf("~> mock removeWhitespaceAsStringSlice")
	// do removeWhitespaceAsStringSlice
	t.Logf("~> do removeWhitespaceAsStringSlice")
	slice := removeWhitespaceAsStringSlice([]string{})
	// verify removeWhitespaceAsStringSlice
	assert.Equal(t, 0, len(slice))
}
