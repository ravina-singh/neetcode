func rob(nums []int) int {
	if len(nums) <= 0{
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	n := len(nums)
	dp1 := make([]int, len(nums)+1)
	dp2 := make([]int, len(nums)+1)

	return max(robMoney(nums[0:n-1], 0, dp1),
	robMoney(nums[1:n], 0, dp2))
}

func robMoney(nums[]int, n int, dp []int) int {
	if n >= len(nums) {
		return 0
	}
	if dp[n] != 0 {
		return dp[n]
	}
	dp[n] = max(nums[n]+robMoney(nums, n+2, dp), robMoney(nums, n+1, dp))
	return dp[n]
}
