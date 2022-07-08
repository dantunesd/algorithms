package insertionsort

func Sort(numbers []int) []int {
	lenNumbers := len(numbers)

	for i := 1; i < lenNumbers; i++ {
		for j := i; j > 0 && numbers[j] < numbers[j-1]; j-- {
			numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
		}
	}

	return numbers
}
