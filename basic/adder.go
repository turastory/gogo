package basic

func Add(a, b int) int {
	return a + b
}

func Sum(numbers []int) (result int) {
	for _, number := range numbers {
		result += number
	}

	return
}

func SumAll(listOfNumbers ...[]int) []int {
	var sums []int

	for _, numbers := range listOfNumbers {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTrails(listOfNumbers ...[]int) []int {
	var sums []int

	for _, numbers := range listOfNumbers {
		lenOfNumbers := len(numbers)
		if lenOfNumbers == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[1:]
			sums = append(sums, Sum(tails))
		}
	}

	return sums
}
