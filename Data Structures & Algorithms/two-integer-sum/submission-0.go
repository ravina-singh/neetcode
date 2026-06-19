func twoSum(nums []int, target int) []int {
   mp := map[int]int{}
   for i, num  := range nums{
	diff := target - num
	if ind, ok := mp[diff]; ok {
		return []int{ind, i}
	}
	mp[num] = i
   }
   return []int{-1,-1}
}
