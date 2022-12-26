package utils_test

import (
	"aoc2022/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPow(t *testing.T) {
	assert.Equal(t, 0, utils.Pow(5, -1))
	assert.Equal(t, 1, utils.Pow(5, 0))
	assert.Equal(t, 5, utils.Pow(5, 1))
	assert.Equal(t, 25, utils.Pow(5, 2))
	assert.Equal(t, 125, utils.Pow(5, 3))
}

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

func TestEqualSlices(t *testing.T) {
	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	assert.True(t, utils.EqualSlices(s1, s2))

	s3 := []int{1, 2, 4}
	assert.False(t, utils.EqualSlices(s1, s3))

	s4 := []int{1, 2, 3, 4}
	assert.False(t, utils.EqualSlices(s1, s4))
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
