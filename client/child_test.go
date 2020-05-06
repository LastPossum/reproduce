package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublicFunc(t *testing.T) {
	assert.Equal(t, PublicFunc(), "foo", "it should be foo")
}
