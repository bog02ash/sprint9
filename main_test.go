package main

// Пишите тесты в этом файле
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomElements(t *testing.T) {
	cases := []struct {
		input    int
		expected int
	}{
		{10, 10},
		{-11, 0},
		{0, 0},
	}
	for _, v := range cases {
		result := generateRandomElements(v.input)
		assert.Len(t, result, v.expected)
	}
}
func TestMaximum(t *testing.T) {
	cases := []struct {
		input    []int
		expected int
	}{
		{[]int{}, 0},
		{[]int{42}, 0},
		{[]int{8, 3, 2, 14, 4}, 14},
	}
	for _, v := range cases {
		result := maximum(v.input)
		assert.Equal(t, result, v.expected)
	}
}
