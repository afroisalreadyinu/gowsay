package gowsay_test

import (
	"github.com/stretchr/testify/assert"
	"gowsay"
	"testing"
)

func TestGowsay(t *testing.T) {
	output, err := gowsay.MakeCow("Hello there", gowsay.Mooptions{})
	assert.Nil(t, err)
	assert.NotEqual(t, "", output)
}
