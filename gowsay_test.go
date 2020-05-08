package gowsay_test

import (
	"github.com/stretchr/testify/assert"
	"gowsay"
	"testing"
)

func TestGowsay(t *testing.T) {
	output, err := gowsay.MakeCow("Hello there", gowsay.Mooptions{})
	assert.Nil(t, err)
	assert.Contains(t, output, "Hello there")
	// The default cow is the apt one
	assert.Contains(t, output, "/------\\/")
}
