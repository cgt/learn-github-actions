package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test2Plus2Equals4(t *testing.T) {
	actual := add(2, 2)
	assert.Equal(t, 4, actual)
}
