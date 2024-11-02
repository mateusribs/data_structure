package algorithms

func CountNumberOccurence(nums []int, target int) int {
	firstOccur := findFirst(nums, target)
	lastOccur := findLast(nums, target)

	occurences := lastOccur - firstOccur + 1
	
	if firstOccur != -1 {
		return occurences
	} else {
		return 0
	}
}

func findFirst(nums []int, target int) int {
	lower := 0
	higher := len(nums)

	firstOcurr := -1

	for {
		mid := lower + (higher - lower) / 2
		currValue := nums[mid]
		if currValue == target {
			higher = mid
			firstOcurr = mid
		} else if target > currValue {
			lower = mid + 1
		} else {
			higher = mid
		}
		
		if lower >= higher {
			break
		}
	}

	return firstOcurr
}

func findLast(nums []int, target int) int {
	lower := 0
	higher := len(nums)

	lastOcurr := -1

	for {
		mid := lower + (higher - lower) / 2
		currValue := nums[mid]
		if currValue == target {
			lower = mid + 1
			lastOcurr = mid
		} else if target > currValue {
			lower = mid + 1
		} else {
			higher = mid
		}
		
		if lower >= higher {
			break
		}
	}

	return lastOcurr
}