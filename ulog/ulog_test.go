package ulog

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	SetLogger(DefaultLogger)
	assert.Equal(t, DefaultLogger, GetLogger())
}
