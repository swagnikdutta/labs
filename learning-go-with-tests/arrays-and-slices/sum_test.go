package arrays_and_slices

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		got := Sum(nums)
		want := 55

		if got != want {
			t.Errorf("want %d, got %d, given, %v\n", want, got, nums)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{4, 5})
	want := []int{6, 9}

	// if !reflect.DeepEqual(got, want) {
	// 	t.Errorf("want %v, got %v\n", want, got)
	// }

	if !slices.Equal(got, want) {
		t.Errorf("want %v, got %v\n", want, got)
	}
}

func TestSumAllTails(t *testing.T) {
	// Here we are assigning the helper fucntion to a variable within the
	// scope of the test function instead of writing the helper function
	// outside.
	// This technique is helpful when you want to reduce the surface area
	// of your API by binding the function to the other local variables
	// in the scope (test functions' scope).
	// By defining this function inside the test, it cannot be used by other
	// functions in the package. Hiding variables and function that don't
	// need to be exported is an important design consideration.
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("want %v, got %v", want, got)
		}
	}

	t.Run("calculate sum of some slices", func(t *testing.T) {
		got := SumAllTails([]int{10, 20, 30}, []int{15, 25}, []int{1, 2, 3, 4, 5, 6, 7})
		want := []int{50, 25, 27}
		checkSums(t, got, want)
	})
	t.Run("calculate sum of slices including empty ones", func(t *testing.T) {
		got := SumAllTails([]int{10, 20}, []int{1, 2, 3, 4, 5}, []int{})
		want := []int{20, 14, 0}
		checkSums(t, got, want)
	})
}
