package slices

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var testSlices Slices

func TestContains(t *testing.T) {
	testArr := []int{123, 321, 234, 342, 111}
	testEle := 234
	found, err := testSlices.Contains(testArr, testEle)
	if err != nil {
		t.Errorf("found unexpected error %s", err)
	}
	assert.True(t, found)

}
