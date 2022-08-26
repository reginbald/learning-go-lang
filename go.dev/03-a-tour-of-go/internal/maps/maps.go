package maps

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, w := range strings.Fields(s) {
		m[w] += 1
	}

	return m
}

func Run() {
	wc.Test(WordCount)
}
