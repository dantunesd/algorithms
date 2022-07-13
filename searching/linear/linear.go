package linear

func Search(value int, list []int) (int, int) {
	loops := 0

	for i := 0; i < len(list); i++ {
		loops++

		if value == list[i] {
			return i, loops
		}
	}

	return -1, loops
}
