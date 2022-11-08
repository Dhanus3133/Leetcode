func longestConsecutive(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
	sort.Ints(nums)
	highest := 1
	tempHighest := 1
  prev := nums[0]
	for i := 1; i < len(nums); i++ {
    if (prev == nums[i]) {
      continue
    }
		if prev == nums[i]-1 {
			tempHighest++
			if tempHighest > highest {
				highest = tempHighest
			}
		} else {
			tempHighest = 1
		}
    prev = nums[i]
	}
	return highest
}