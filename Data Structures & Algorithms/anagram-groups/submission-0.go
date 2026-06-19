func groupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, str := range strs{
		count := [26]int{}
		for _, c := range str{
			count[c-'a']++
		}
		mp[count] = append(mp[count], str)
	}
	res := [][]string{}
	for _, v := range mp {
		res = append(res, v)
	}
	return res
}
