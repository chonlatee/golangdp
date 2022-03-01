package main

import "fmt"

func howSum(targetSum int, nums []int) []int {
	var nilSlice []int
	if targetSum == 0 {
		return []int{}
	}

	if targetSum < 0 {
		return nilSlice
	}

	for _, v := range nums {
		rem := targetSum - v
		s := howSum(rem, nums)
		if s != nil {
			x := []int{}
			x = append(x, s...)
			x = append(x, v)
			return x
		}
	}

	return nilSlice
}

func howSumMemo(targetSum int, nums []int, memo map[int][]int) []int {

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

	for _, v := range nums {
		rem := targetSum - v
		s := howSumMemo(rem, nums, memo)
		if s != nil {
			x := []int{}
			x = append(x, s...)
			x = append(x, v)
			memo[targetSum] = x
			return x
		}
	}

	memo[targetSum] = nilSlice
	return nilSlice
}

func main() {
	r := howSum(7, []int{5, 3, 4, 7})

	fmt.Printf("%v\n", r)

	r = howSum(7, []int{2, 4})

	fmt.Printf("%v\n", r)

	memo := make(map[int][]int)

	r = howSumMemo(3000, []int{7, 14}, memo)

	fmt.Printf("%v\n", r)

	r = howSum(3000, []int{7, 14})

	fmt.Printf("%v\n", r)
}
