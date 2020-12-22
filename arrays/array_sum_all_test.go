package arrays

import (
	"reflect"
	"testing"
)

func TestArraySumAll(t *testing.T) {
	numbers1st := []int{1, 2, 3, 4, 5}
	numbers2nd := []int{5, 5}

	got := ArraySumAll(numbers1st, numbers2nd)
	want := []int{15, 10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
