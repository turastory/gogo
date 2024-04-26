package basic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 3)
	expected := 5

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	sum := Sum(numbers)
	expected := 15

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{4, 5}, []int{3, 7})
	want := []int{9, 10}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTrails(t *testing.T) {
	checkSums := func(got []int, want []int, t *testing.T) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTrails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(got, want, t)
	})

	t.Run("safely handle empty slices", func(t *testing.T) {
		got := SumAllTrails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}

		checkSums(got, want, t)
	})
}
