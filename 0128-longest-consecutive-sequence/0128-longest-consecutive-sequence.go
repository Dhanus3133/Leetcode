func longestConsecutive(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}
  m := make(map[int]bool)
  fNums := []int{}
  for _, num := range nums {
    if _, pres := m[num]; !pres {
      m[num] = true
      fNums = append(fNums, num)
    }
  }
  nums = fNums
  fNums = nil
	sort.Ints(nums)
	highest := 1
	tempHighest := 1
	for i := 1; i < len(nums); i++ {
		if nums[i-1] == nums[i]-1 {
			tempHighest++
			if tempHighest > highest {
				highest = tempHighest
			}
		} else {
			tempHighest = 1
		}
	}
	return highest
}