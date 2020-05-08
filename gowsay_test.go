package gowsay_test

import (
	"github.com/stretchr/testify/assert"
	"gowsay"
	"testing"
)

func TestGowsay(t *testing.T) {
	output, err := gowsay.MakeCow("Hello there", gowsay.Apt, gowsay.Mooptions{})
	assert.Nil(t, err)
	assert.Contains(t, output, "Hello there")
	// The default cow is the apt one
	assert.Contains(t, output, "/------\\/")
}

func TestGowsayErrorOnInvalidCow(t *testing.T) {
	output, err := gowsay.MakeCow("Hello there", gowsay.CowType(1000), gowsay.Mooptions{Cowfile: "no-such-cow"})
	assert.NotNil(t, err)
	assert.Equal(t, output, "")
}
