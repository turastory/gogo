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
