package main

import (
	"fmt"
	"io/ioutil"

	"github.com/leonobilis/adventofcode2022/utils"
)

func getStartOfPacket(packet string, markerLen int) int {
	for i := 0; i < len(packet)-markerLen; i++ {
		s := utils.NewSet([]byte(packet[i : i+markerLen]))
		if len(s) == markerLen {
			return i + markerLen
		}
	}
	return 0
}

func p1(inp string) int {
	return getStartOfPacket(inp, 4)
}

func p2(inp string) int {
	return getStartOfPacket(inp, 14)
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	inp := string(input)
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
