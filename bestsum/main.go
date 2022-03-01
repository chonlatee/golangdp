package main

import "fmt"

func bestSum(targetSum int, nums []int) []int {
	var nilSlice []int
	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nilSlice
	}

	var shortCom []int

	for _, v := range nums {
		rem := targetSum - v
		r := bestSum(rem, nums)
		if r != nil {
			x := []int{}
			x = append(x, r...)
			x = append(x, v)

			if shortCom == nil || len(x) < len(shortCom) {
				shortCom = x
			}

		}
	}

	return shortCom
}

func bestSumMemo(targetSum int, nums []int, memo map[int][]int) []int {

	v, ok := memo[targetSum]
	if ok {
		return v
	}

	var nilSlice []int
	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nilSlice
	}

	var shortCom []int

	for _, v := range nums {
		rem := targetSum - v
		r := bestSumMemo(rem, nums, memo)
		if r != nil {
			x := []int{}
			x = append(x, r...)
			x = append(x, v)

			if shortCom == nil || len(x) < len(shortCom) {
				shortCom = x
			}
		}
	}

	memo[targetSum] = shortCom
	return shortCom
}

func main() {

	r := bestSum(7, []int{5, 3, 4, 7})
	fmt.Println(r)

	r = bestSum(8, []int{2, 3, 5})
	fmt.Println(r)

	memo1 := make(map[int][]int)
	r = bestSumMemo(8, []int{2, 3, 5}, memo1)
	fmt.Println(r)

	memo2 := make(map[int][]int)
	r = bestSumMemo(100, []int{1, 2, 5, 25}, memo2)
	fmt.Println(r)
}
