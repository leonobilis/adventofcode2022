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

func MinMax(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func Mod(a, b int) int {
	return (a%b + b) % b
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

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](slice []T) Set[T] {
	set := make(Set[T])
	for _, s := range slice {
		set.Add(s)
	}
	return set
}

func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

func (s Set[T]) Intersect(s2 Set[T]) Set[T] {
	result := make(Set[T])
	for v := range s {
		if _, ok := s2[v]; ok {
			result.Add(v)
		}
	}
	return result
}

func (s Set[T]) Union(s2 Set[T]) Set[T] {
	for v := range s2 {
		s.Add(v)
	}
	return s
}
