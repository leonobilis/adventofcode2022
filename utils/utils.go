package utils

import (
	"sort"
	"strconv"
	"strings"
)

func Atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		panic(err)
	}
	return i
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GetIntArray(input string) (output []int) {
	for _, s := range strings.Split(input, "\n") {
		output = append(output, Atoi(s))
	}
	return
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
