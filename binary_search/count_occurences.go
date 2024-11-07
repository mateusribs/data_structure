package algorithms

func CountNumberOccurence(nums []int, target int) int {
	firstOccur := findLimitRange(nums, target, true)
	lastOccur := findLimitRange(nums, target, false)

	occurences := lastOccur - firstOccur + 1
	
	if firstOccur != -1 {
		return occurences
	} else {
		return 0
	}
}

func findLimitRange(nums []int, target int, findFirst bool) int {
	lower := 0
	higher := len(nums)

	limitIndex := -1

	for higher > lower {
		mid := lower + (higher - lower) / 2
		currValue := nums[mid]
		if currValue == target {
			if findFirst {
				higher = mid
			} else {
				lower = mid + 1
			}
			limitIndex = mid
		} else if target > currValue {
			lower = mid + 1
		} else {
			higher = mid
		}
	}

	return limitIndex
}