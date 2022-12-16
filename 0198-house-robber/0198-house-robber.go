func max(a, b int) int {
    if a < b {
      return b
    }
    return a
}

func rob(nums []int) int {
  if (len(nums) == 1) {
    return nums[0]
  }
  if (len(nums) == 2) {
    return max(nums[0], nums[1])
  }
  
	dp := make([]int, len(nums))
  
  dp[0] = nums[0]
  dp[1] = nums[1]
  dp[2] = nums[0] + nums[2]
  
  for i := 3; i < len(nums); i++ {
    dp[i] = max(dp[i-3], dp[i-2]) + nums[i]
  }
  fmt.Println(dp)
  return max(dp[len(nums)-1], dp[len(nums)-2])
}