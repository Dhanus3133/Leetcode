func containsDuplicate(nums []int) bool {
    m := make(map[int]bool)
	for _, num := range nums {
		_, prs := m[num]
		if prs {
			return true
		}
		m[num] = true
	}
	return false

}