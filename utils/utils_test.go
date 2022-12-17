package utils_test

import (
	"aoc2022/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIn(t *testing.T) {
	list := []int{1, 2, 3, 4, 5}
	assert.True(t, utils.In(list, 3))
	assert.False(t, utils.In(list, 10))
}

func TestMaxIn(t *testing.T) {
	list := []int{1, 12, 3, 24, 15}
	assert.Equal(t, utils.MaxIn(list), 24)
}

func TestMinIn(t *testing.T) {
	list := []int{11, 12, 3, 24, 15}
	assert.Equal(t, utils.MinIn(list), 3)
}

func TestCopyMap(t *testing.T) {
	m1 := map[string]interface{}{
		"a": "bbb",
		"b": map[string]interface{}{
			"c": 123,
		},
	}

	m2 := utils.CopyMap(m1)

	m1["a"] = "zzz"
	delete(m1, "b")

	assert.Equal(t, map[string]interface{}{"a": "zzz"}, m1)
	assert.Equal(t, map[string]interface{}{
		"a": "bbb",
		"b": map[string]interface{}{
			"c": 123,
		},
	}, m2)
}
