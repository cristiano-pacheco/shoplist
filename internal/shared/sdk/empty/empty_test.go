package empty

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected bool
	}{
		{"zero int", 0, true},
		{"non-zero int", 42, false},
		{"empty string", "", true},
		{"non-empty string", "test", false},
		{"nil", nil, true},
		{"empty slice", []int{}, true},
		{"non-empty slice", []int{1, 2, 3}, false},
		{"empty map", map[string]int{}, true},
		{"non-empty map", map[string]int{"a": 1}, false},
		{"zero float", 0.0, true},
		{"non-zero float", 3.14, false},
		{"false bool", false, true},
		{"true bool", true, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsEmpty(tt.input), "IsEmpty(%v) should be %v", tt.input, tt.expected)
		})
	}
}

func TestIsEmpty_AdditionalCases(t *testing.T) {
	// Struct tests
	type person struct {
		Name string
	}
	assert.False(t, IsEmpty(person{}), "empty struct should not be considered empty")
	assert.False(t, IsEmpty(person{Name: "John"}), "struct with values should not be considered empty")

	// Pointer tests
	var nilPtr *string
	str := "hello"
	strPtr := &str
	assert.True(t, IsEmpty(nilPtr), "nil pointer should be empty")
	assert.False(t, IsEmpty(strPtr), "pointer to non-empty value should not be empty")

	// Channel tests
	ch := make(chan int)
	assert.False(t, IsEmpty(ch), "channel should not be considered empty")

	// Interface tests
	var i interface{}
	assert.True(t, IsEmpty(i), "nil interface should be empty")
	i = 42
	assert.False(t, IsEmpty(i), "interface with non-zero value should not be empty")
}
