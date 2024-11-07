package algorithms


func BinarySearch(array []int, target int) int {
	lower_index := 0
	higher_index := len(array)

	for higher_index > lower_index {
		m := lower_index + (higher_index - lower_index) / 2

		currValue := array[m]

		if currValue == target {
			return m
		} else if target > currValue {
			lower_index = m + 1
		} else {
			higher_index = m
		}
	}

	return -1
}