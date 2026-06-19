func hasDuplicate(nums []int) bool {
	mp := map[int]struct{}{}
    for _, num := range nums {
		if _, ok := mp[num]; ok{
			return true
		}
		mp[num]=struct{}{}
	}
	return false
}
