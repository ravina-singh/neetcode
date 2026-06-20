func rob(nums []int) int {
    dp := make([]int , len(nums) + 1)
	return robMoney(nums, 0, dp)
}

func robMoney(nums []int, n int, dp []int) int{
	if n >= len(nums) {
		return 0
	}
	if dp[n] != 0 {
		return dp[n]
	}
	dp[n] = max(nums[n] + robMoney(nums, n+2, dp), robMoney(nums, n+1, dp))
	return dp[n]
}