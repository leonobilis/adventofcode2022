package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}

func decode(snafu string) (number int) {
	for i, place := len(snafu)-1, 1; i >= 0; i, place = i-1, place*5 {
		switch snafu[i] {
		case '-':
			number -= place
		case '=':
			number -= 2 * place
		case '1', '2':
			number += int(snafu[i]-48) * place
		}
	}
	return
}

func encode(number int) string {
	s := make([]byte, 0)

	for base := 1; number > 0; base *= 5 {
		switch (number / base) % 5 {
		case 0:
			s = append(s, '0')
		case 1:
			s = append(s, '1')
			number -= base
		case 2:
			s = append(s, '2')
			number -= 2 * base
		case 3:
			s = append(s, '=')
			number += base * 2
		case 4:
			s = append(s, '-')
			number += base
		}
	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}

func p1(inp []string) string {
	sum := 0
	for _, number := range inp {
		sum += decode(number)
	}
	return encode(sum)
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
}
