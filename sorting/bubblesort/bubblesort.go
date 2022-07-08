package bubblesort

func BubbleSort(numbers []int) []int {
	lenNumbers := len(numbers)

	for i := 0; i < lenNumbers; i++ {
		for j := 0; j < lenNumbers-1; j++ {
			if numbers[j] > numbers[j+1] {
				temporary := numbers[j]
				numbers[j] = numbers[j+1]
				numbers[j+1] = temporary
			}
		}
	}

	return numbers
}
