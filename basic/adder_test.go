package basic

import (
	"fmt"
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
