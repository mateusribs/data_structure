package algorithms

func TwoSumHashMap(nums []int, target int) []int { 
	indexMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		curr := nums[i]
		diff := target - curr
		
		if j, ok := indexMap[diff]; ok {
			sum := []int{i, j}
			return sum
		} else {
			indexMap[curr] = i
		}
	}

	return nil
}

func TwoSumBruteForce(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			sum := nums[i] + nums[j]

			if sum == target {
				return []int{i, j}
			}
		}
	}

	return nil
}