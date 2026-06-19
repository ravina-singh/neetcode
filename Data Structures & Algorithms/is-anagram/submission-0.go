func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	mp := [26]int{}
	for i:=0; i<len(s); i++{
		mp[s[i]-'a']++
		mp[t[i]-'a']--
	}
	for _, v := range mp{
		if v != 0 {
			return false
		}
	}
	return true
}
