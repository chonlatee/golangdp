package main

import "fmt"

func canSum(targetSum int, nums []int) bool {
	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	for _, v := range nums {
		rem := targetSum - v
		if canSum(rem, nums) {
			return true
		}
	}

	return false

}

func canSumMemo(targetSum int, nums []int, memo map[int]bool) bool {

	v, ok := memo[targetSum]
	if ok {
		return v
	}

	if targetSum == 0 {
		return true
	}

	if targetSum < 0 {
		return false
	}

	for _, v := range nums {
		rem := targetSum - v
		if canSumMemo(rem, nums, memo) {
			memo[targetSum] = true
			return true
		}
	}

	memo[targetSum] = false
	return false

}

func cansumTabu(targetSum int, numbers []int) bool {

	table := make([]bool, targetSum+1, targetSum+1)
	table[0] = true
	for i := 0; i <= targetSum; i++ {
		if table[i] == true {
			for _, num := range numbers {
				if num+i <= targetSum {
					table[num+i] = true
				}
			}
		}
	}

	return table[targetSum]
}

func main() {

	r := canSum(8, []int{3})

	fmt.Println(r)

	memo := make(map[int]bool)

	r = canSumMemo(300, []int{7, 14}, memo)

	fmt.Println(r)

	r = cansumTabu(300, []int{7, 14})
	fmt.Println(r)

	r = cansumTabu(7, []int{2, 3})
	fmt.Println(r)
}
