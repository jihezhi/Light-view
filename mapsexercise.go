package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	sr := strings.Fields(s)
	m := make(map[string]int)
	for i := 0; i < len(sr); i++ {
		if _, ok := m[sr[i]]；ok { //纱雾酱：很棒，这才是优雅的Go写法。这样ok变量不会被泄露到main中，十分安全。
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
