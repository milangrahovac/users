package web

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckName(t *testing.T) {
	testCases := []struct {
		name     string
		expected bool
	}{
		{"Abc", true},
		{"123", false},
		{"Abc123", false},
	}

	for _, c := range testCases {
		res := checkName(c.name)
		assert.Equal(t, c.expected, res)
	}
}

func TestCheckPassword(t *testing.T) {
	testCases := []struct {
		password string
		expected bool
	}{
		{"qwerty", false},
		{"12345", false},
		{"asdfghjk", false},
		{"1234abcd", false},
		{"#dnkn2332", false},
		{"A123456b", true},
	}

	for _, c := range testCases {
		res := checkPassword(c.password)
		assert.Equal(t, c.expected, res, "Failed case is: "+c.password)
	}

}
