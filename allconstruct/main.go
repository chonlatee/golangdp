package main

import (
	"fmt"
	"strings"
)

func allConstruct(target string, words []string) [][]string {

	if len(target) == 0 {
		s := [][]string{{}}
		return s
	}

	result := [][]string{}

	for _, v := range words {
		if strings.HasPrefix(target, v) {
			suffix := target[len(v):]
			suffixWays := allConstruct(suffix, words)
			targetWays := make([][]string, len(suffixWays))
			for i, s := range suffixWays {
				targetWays[i] = append(targetWays[i], v)
				targetWays[i] = append(targetWays[i], s...)
			}
			for _, vv := range targetWays {
				result = append(result, vv)
			}
		}
	}

	return result
}

func allConstructMemo(target string, words []string, memo map[string][][]string) [][]string {

	if v, ok := memo[target]; ok {
		return v
	}

	if len(target) == 0 {
		s := [][]string{{}}
		return s
	}

	result := [][]string{}

	for _, v := range words {
		if strings.HasPrefix(target, v) {
			suffix := target[len(v):]
			suffixWays := allConstructMemo(suffix, words, memo)
			targetWays := make([][]string, len(suffixWays))
			for i, s := range suffixWays {
				targetWays[i] = append(targetWays[i], v)
				targetWays[i] = append(targetWays[i], s...)
			}
			for _, vv := range targetWays {
				result = append(result, vv)
			}
		}
	}

	memo[target] = result
	return result
}

func main() {

	//[[purp le] [p ur p le]]
	fmt.Printf("purple = %v\n", allConstruct("purple", []string{"purp", "p", "ur", "le", "purpl"}))

	// [[ab cd ef] [ab c def] [abc def] [abcd ef]]
	fmt.Printf("abcdef = %v\n", allConstruct("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}))

	// []
	fmt.Printf("skateboard = %v\n", allConstruct("skateboard", []string{"bo", "rd", "ate", "t", "ska", "boar"}))

	// []

	memo := make(map[string][][]string)
	fmt.Printf("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaz = %v\n", allConstructMemo("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaz", []string{"a", "aa", "aaa", "aaa", "aaaa", "aaaaa"}, memo))

	// fmt.Printf("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaz = %v\n", allConstruct("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaz", []string{"a", "aa", "aaa", "aaa", "aaaa", "aaaaa"}))
}
