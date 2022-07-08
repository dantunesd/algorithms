package bubblesort

func BubbleSort(numbers []int) []int {
	lenNumbers := len(numbers)

	for i := 0; i < lenNumbers; i++ {
		sorted := false

		for j := 0; j < lenNumbers-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
				sorted = true
			}
		}

		if !sorted {
			break
		}
	}

	return numbers
}
