package main

import (
	"assertT/assert"
	"fmt"
	"testing"
)

func TestAssert(t *testing.T) {
	a := assert.New(t)
	a.Equal(1, 1, "should be equal")
	a.NotEqual(1, 2, "should not be equal")
	a.True(true, "should be true")
	a.False(false, "should be false")
	a.Nil(nil, "should be nil")
	a.NotNil(1, "should not be nil")
	a.Contains([]int{1, 2, 3}, 2, "should contain 2")
	a.Contains([]int{1, 2, 3}, 4, "should not contain 4")
	a.NoError(nil, "should be nil")
	a.NoError(fmt.Errorf("error"), "should be nil")
}
