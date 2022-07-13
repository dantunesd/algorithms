package binary

func Search(value int, list []int) (int, int) {
	return binarySearch(value, 0, len(list)-1, list, 0)
}

func binarySearch(value, left, right int, list []int, loops int) (int, int) {
	if right >= left {
		loops++

		middle := left + (right-left)/2

		if value == list[middle] {
			return middle, loops
		}

		if value < list[middle] {
			return binarySearch(value, left, middle-1, list, loops)
		}

		return binarySearch(value, middle+1, right, list, loops)
	}

	return -1, loops
}
