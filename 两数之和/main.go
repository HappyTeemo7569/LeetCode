package main

import "fmt"

/**
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。

*/

func main() {
	fmt.Println(11)

	test := []int{3, 2, 4}
	res := twoSumHash(test, 6)

	fmt.Print(res)
}

//两层遍历
func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := len(nums) - 1; j > 0; j-- {
			if i != j && nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

//利用哈希
func twoSumHash(nums []int, target int) []int {

	length := len(nums)

	hashTmp := make(map[int]int)
	for i := 0; i < length; i++ {
		hashTmp[nums[i]] = i
	}

	for j := 0; j < length; j++ {
		if _, ok := hashTmp[target-nums[j]]; ok && hashTmp[target-nums[j]] != j {
			return []int{hashTmp[target-nums[j]], j}
		}
	}

	return []int{}
}

//答案
func answer(nums []int, target int) []int {

	hashTable := map[int]int{}
	for i, x := range nums {
		if p, ok := hashTable[target-x]; ok {
			return []int{p, i}
		}
		hashTable[x] = i
	}
	return nil

}
