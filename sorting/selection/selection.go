package selection

func Sort(numbers []int) []int {
	lenNumbers := len(numbers)

	for i := 0; i < lenNumbers; i++ {
		min := i

		for j := i + 1; j < lenNumbers; j++ {
			if numbers[j] < numbers[min] {
				min = j
			}
		}

		if numbers[i] > numbers[min] {
			temporary := numbers[i]
			numbers[i] = numbers[min]
			numbers[min] = temporary
		}
	}

	return numbers
}
