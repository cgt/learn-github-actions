package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test2Plus2Equals4(t *testing.T) {
	assert.Equal(t, 4, add(2, 2))
}
