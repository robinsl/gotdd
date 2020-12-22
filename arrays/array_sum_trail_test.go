package arrays

import (
	"reflect"
	"testing"
)

func TestArraySumTrail(t *testing.T) {
	numbers1st := []int{1, 2, 3, 4, 5}
	numbers2nd := []int{5, 5}

	got := ArraySumTrail(numbers1st, numbers2nd)
	want := []int{14, 5}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
