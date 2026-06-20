type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	ans := ""
	for _, str := range strs {
		ans += strconv.Itoa(len(str)) + "#" + str
	}
	return ans
}

func (s *Solution) Decode(encoded string) []string {
	ans := []string{}
	i := 0
	for i < len(encoded) {
		j := i
		for encoded[j] != '#'{
			j++
		}
		l, _ := strconv.Atoi(encoded[i:j])
		i = j+1
		str := encoded[i:i+l]
		ans = append(ans, str)
		i = i + l
	}
	return ans
}
