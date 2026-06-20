func climbStairs(n int) int {
	dp := make([]int, n+1)
	return ways(n, dp)
}

func ways(n int, dp []int) int{
	if n <= 2 {
		return n
	}
	if dp[n] != 0 {
		return dp[n]
	}
	dp[n] = ways(n-1, dp) + ways(n-2, dp)
	return dp[n]
}
