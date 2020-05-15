package gowsay_test

import (
	"github.com/afroisalreadyinu/gowsay"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGowsayDefaultCow(t *testing.T) {
	output, err := gowsay.MakeCow("Hello there", gowsay.Mooptions{})
	assert.Nil(t, err)
	assert.Contains(t, output, "Hello there")
	// The default cow is the apt one
	assert.Contains(t, output, "/------\\/")
}

func TestGowsayErrorOnInvalidCow(t *testing.T) {
	output, err := gowsay.MakeCow("Hello there", gowsay.Mooptions{Cowfile: "no-such-cow"})
	assert.NotNil(t, err)
	assert.Equal(t, output, "")
}
