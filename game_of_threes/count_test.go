package game_of_threes_test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go-katas/game_of_threes"
	"testing"
)

func TestCount(t *testing.T) {
	type testDatum struct {
		num  int
		want []int
	}
	testdata := []testDatum{
		{1, []int{1}},
		{2, []int{2, 3, 1}},
		{3, []int{3, 1}},
		{5, []int{5, 6, 2, 3, 1}},
		{10, []int{10, 9, 3, 1}},
	}

	for _, td := range testdata {
		t.Run(fmt.Sprintf("For %d", td.num), func(t *testing.T) {
			assert.Equal(t, td.want, game_of_threes.Count(td.num))
		})
	}

}
