func twoSum(nums []int, target int) []int {
    m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		j, pres := m[target-nums[i]]
		if pres {
			return []int{j, i}
		}
        m[nums[i]] = i
	}
	return []int{}

}