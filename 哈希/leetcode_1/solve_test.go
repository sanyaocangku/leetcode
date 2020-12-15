package leetcode_1

func twoSum(nums []int, target int) []int {
	mp := make(map[int]int)
	for i, num := range nums {
		if pos, ok := mp[target-num]; ok {
			return []int{pos, i}
		}
		mp[num] = i
	}
	return []int{}
}
