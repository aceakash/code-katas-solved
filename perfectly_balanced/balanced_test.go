package balanced_test

import (
	"github.com/stretchr/testify/assert"
	. "go-katas/perfectly_balanced"
	"testing"
)

func TestIsBalanced(t *testing.T) {
	assert.False(t, IsBalanced("("))
	assert.False(t, IsBalanced(")"))
	assert.False(t, IsBalanced("(("))
	assert.False(t, IsBalanced("))"))
	assert.False(t, IsBalanced(")("))
	assert.False(t, IsBalanced("()("))
	assert.False(t, IsBalanced("(()"))
	assert.False(t, IsBalanced(")()"))

	assert.True(t, IsBalanced("()"))
	assert.True(t, IsBalanced("()()"))
	assert.True(t, IsBalanced("(())"))
	assert.True(t, IsBalanced("(()())"))
	assert.True(t, IsBalanced("((()()()))"))
	assert.True(t, IsBalanced("()(())((()()()))"))

	assert.True(t, IsBalanced("{}"))
	assert.True(t, IsBalanced("{}{}"))

	assert.False(t, IsBalanced("((}}"))

	assert.True(t, IsBalanced("({})"))
	assert.True(t, IsBalanced("{}({(){}}({}))"))

	assert.True(t, IsBalanced("{[{}{}]}[()]"))
	assert.True(t, IsBalanced("{{}{}}"))
	assert.True(t, IsBalanced("[]{}()"))

	assert.False(t, IsBalanced("{()}[)"))
	assert.False(t, IsBalanced("{(})"))
}
