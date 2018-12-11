package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	sr := strings.Fields(s)
	m := make(map[string]int)
	for i := 0; i < len(sr); i++ {
		if _, ok := m[sr[i]]；ok {
			m[sr[i]] += 1
		} else {
			m[sr[i]] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
