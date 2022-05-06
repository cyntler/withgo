package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumUtil(t *testing.T) {
	assert.Equal(t, Sum(10, 10), 21)
}
