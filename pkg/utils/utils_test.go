package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumUtilValidValue(t *testing.T) {
	assert.Equal(t, Sum(10, 10), 20)
}

func TestSumUtilInvalidValue(t *testing.T) {
	assert.NotEqual(t, Sum(2, 4), 10)
}
