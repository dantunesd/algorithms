package linear

func Search(value int, list []int) (int, int) {
	for i := 0; i < len(list); i++ {
		if value == list[i] {
			return i, i + 1
		}
	}

	return -1, 0
}
