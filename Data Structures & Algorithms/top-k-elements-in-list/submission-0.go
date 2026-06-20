func topKFrequent(nums []int, k int) []int {
    count := map[int]int{}
    for _, num := range nums {
        count[num]++
    }

    freq := make([][]int,len(nums)+1)
    for key, val := range count {
        if freq[val] == nil {
            freq[val] = []int{}
        }
        freq[val] = append(freq[val], key)
    }

    ans := []int{}
    for i := len(freq)-1; i>=0; i-- {
        for _, val := range freq[i] {
            if len(ans) == k {
                break
            }
            ans = append(ans, val)
        }
    }
    return ans
}
