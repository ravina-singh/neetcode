func productExceptSelf(nums []int) []int {
	n := len(nums)
	lt := make([]int, n)
	rt := make([]int, n)

	lt[0], rt[n-1] = 1, 1

	for i := 1; i < n; i++ {
		lt[i] = nums[i-1]*lt[i-1]
	}

	for i := n-2; i >= 0 ; i-- {
		rt[i] = nums[i+1]*rt[i+1]
	}

	prd := make([]int, n)
	for i := 0; i < n; i++ {
		prd[i] = lt[i] * rt[i]
	}

	return prd
}
