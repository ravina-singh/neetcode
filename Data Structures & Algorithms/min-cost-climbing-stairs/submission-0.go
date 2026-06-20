func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	return min(minCost(cost, 0, dp), minCost(cost, 1, dp))
}

func minCost(cost []int, n int, dp []int) int{
	if n >= len(cost) {
		return 0
	}
	if dp[n] != 0 {
		return dp[n]
	}
	dp[n] = cost[n] + min(minCost(cost, n+1, dp), minCost(cost, n+2, dp))
	return dp[n]
}
