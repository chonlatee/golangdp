package main

import (
	"fmt"
	"strings"
)

func canConstruct(target string, words []string) bool {

	if len(target) == 0 {
		return true
	}

	for _, v := range words {
		if strings.HasPrefix(target, v) {
			suffix := target[len(v):]
			if canConstruct(suffix, words) {
				return true
			}
		}
	}

	return false
}

func canConstructMemo(target string, words []string, memo map[string]bool) bool {

	if v, ok := memo[target]; ok {
		return v
	}

	if len(target) == 0 {
		return true
	}

	for _, v := range words {
		if strings.HasPrefix(target, v) {
			suffix := target[len(v):]
			if canConstructMemo(suffix, words, memo) {
				memo[target] = true
				return true
			}
		}
	}

	memo[target] = false
	return false
}

func main() {
	fmt.Println(canConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))                  // true
	fmt.Println(canConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))   // false
	fmt.Println(canConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"})) // true

	//fmt.Println(canConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eeee", "eeeee", "eeeeee"})) // false

	memo := make(map[string]bool)
	fmt.Println(canConstructMemo("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"e", "ee", "eeee", "eeeee", "eeeeee"}, memo)) // false
}
