package main

import (
	"fmt"
	"strings"
)

func countConstruct(target string, words []string) int {

	if len(target) == 0 {
		return 1
	}

	totalCount := 0

	for _, v := range words {
		if strings.HasPrefix(target, v) {
			suffix := target[len(v):]
			numWays := countConstruct(suffix, words)
			totalCount += numWays
		}
	}
	return totalCount
}

func countConstructMemo(target string, words []string, memo map[string]int) int {

	if v, ok := memo[target]; ok {
		return v
	}

	if len(target) == 0 {
		return 1
	}

	totalCount := 0

	for _, v := range words {
		if strings.HasPrefix(target, v) {
			suffix := target[len(v):]
			numWays := countConstructMemo(suffix, words, memo)
			totalCount += numWays
		}
	}
	memo[target] = totalCount
	return totalCount
}

func main() {

	fmt.Println(countConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))                  // 2
	fmt.Println(countConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))                  // 1
	fmt.Println(countConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))   // 0
	fmt.Println(countConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"})) // 4
	//fmt.Println(countConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"})) // 0

	memo := make(map[string]int)

	fmt.Println(countConstructMemo("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}, memo)) // 4

	memo = make(map[string]int)

	fmt.Println(countConstructMemo("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eee", "eeee", "eeeee", "eeeeee"}, memo)) // 0
}
