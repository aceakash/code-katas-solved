package rotate_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-katas/rotate"
	"testing"
)

func TestRotate(t *testing.T) {
	t.Run("rotate by 0 gives the same list", func(t *testing.T) {
		list := []int{1, 2, 3}
		newList := rotate.Rotate(list, 0)
		assert.Equal(t, []int{1, 2, 3}, newList)
	})

	t.Run("rotate by 1 gives the rotated list", func(t *testing.T) {
		type TestDatum struct {
			list []int
			want []int
		}
		testData := []TestDatum{
			{
				list: []int{},
				want: []int{},
			}, {
				list: []int{1},
				want: []int{1},
			}, {
				list: []int{1, 2, 3},
				want: []int{2, 3, 1},
			},
			{
				list: []int{1, 2},
				want: []int{2, 1},
			},
		}

		for _, td := range testData {
			t.Run(fmt.Sprintf("%v should give %v", td.list, td.want), func(t *testing.T) {
				got := rotate.Rotate(td.list, 1)
				assert.Equal(t, td.want, got)
			})
		}
	})

	t.Run("rotate by 2 gives the rotated list", func(t *testing.T) {
		type TestDatum struct {
			list []int
			want []int
		}
		testData := []TestDatum{
			{list: []int{}, want: []int{}},
			{list: []int{1}, want: []int{1}},
			{list: []int{1, 2, 3}, want: []int{3, 1, 2}},
			{list: []int{1, 2}, want: []int{1, 2}},
		}

		for _, td := range testData {
			t.Run(fmt.Sprintf("%v should give %v", td.list, td.want), func(t *testing.T) {
				got := rotate.Rotate(td.list, 2)
				assert.Equal(t, td.want, got)
			})
		}
	})

	t.Run("rotate by more the number of the items in the list", func(t *testing.T) {
		list := []int{1, 2, 3}
		newList := rotate.Rotate(list, 4)
		assert.Equal(t, []int{2, 3, 1}, newList)
	})
}
