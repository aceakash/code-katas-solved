package stack_test

import (
	"github.com/stretchr/testify/assert"
	"go-katas/stack"
	"testing"
)

func TestStackIsEmptyWhenCreated(t *testing.T) {
	s := stack.New[int](5)
	assert.True(t, s.IsEmpty())
}

func TestPushingAnItemThenPoppingGivesTheItem(t *testing.T) {
	t.Run("with an empty stack", func(t *testing.T) {
		s := stack.New[int](5)
		ok := s.Push(1)
		assert.True(t, ok)

		got, ok := s.Pop()
		assert.True(t, ok)
		assert.Equal(t, 1, got)
	})

	t.Run("with a stack with a few items", func(t *testing.T) {
		s := stack.New[int](5)
		s.Push(1)
		s.Push(2)
		s.Push(3)
		got, ok := s.Pop()
		assert.True(t, ok)
		assert.Equal(t, 3, got)
	})
}

func TestPoppingAnEmptyStackReturnsFalse(t *testing.T) {
	s := stack.New[int](1)
	_, popped := s.Pop()
	assert.False(t, popped)
}

func Test_Pushing_Three_Items_Then_Popping_Thrice_Returns_Them_In_Reverse_Order(t *testing.T) {
	s := stack.New[string](10)
	s.Push("danny")
	s.Push("jon")
	s.Push("ned")

	n, _ := s.Pop()
	j, _ := s.Pop()
	d, _ := s.Pop()

	assert.Equal(t, "ned", n)
	assert.Equal(t, "jon", j)
	assert.Equal(t, "danny", d)
}

func Test_Pushing_Items_Then_Popping_Them_All_Leaves_An_Empty_Stack_That_Cannot_Be_Popped_Further(t *testing.T) {
	s := stack.New[int](10)

	s.Push(1)
	s.Push(2)
	s.Push(3)

	s.Pop()
	s.Pop()
	s.Pop()

	assert.True(t, s.IsEmpty())

	_, ok := s.Pop()
	assert.False(t, ok)
}

func Test_Stack_Is_Not_Empty_When_All_Pushed_Item_Have_Not_Been_Popped(t *testing.T) {
	s := stack.New[int](10)
	s.Push(1)
	s.Push(2)
	s.Push(3)

	s.Pop()

	assert.False(t, s.IsEmpty())
}

func Test_Stack_Is_Not_Full_When_Enough_Items_Have_Not_Been_Pushed(t *testing.T) {
	s := stack.New[int](3)

	s.Push(1)
	s.Push(2)

	assert.False(t, s.IsFull())
}

func Test_Stack_Is_Full_When_Enough_Items_Have_Been_Pushed_To_Fill_Capacity(t *testing.T) {
	s := stack.New[int](3)

	s.Push(1)
	s.Push(2)
	s.Push(3)

	assert.True(t, s.IsFull())
}

func Test_Pushing_To_A_Full_Stack_Returns_False(t *testing.T) {
	s := stack.New[int](3)
	s.Push(1)
	s.Push(2)
	s.Push(3)

	ok := s.Push(4)

	assert.False(t, ok)
}

func Test_Peek_Returns_Top_Of_The_Stack_Without_Popping_It(t *testing.T) {
	s := stack.New[int](3)
	s.Push(1)

	item, ok := s.Peek()
	assert.True(t, ok)
	assert.Equal(t, 1, item)

	// stack still contains the item
	assert.False(t, s.IsEmpty())
}

func Test_Peeking_An_Empty_Test_Returns_False(t *testing.T) {
	s := stack.New[int](10)

	_, ok := s.Peek()
	assert.False(t, ok)
}
