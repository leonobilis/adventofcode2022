package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/leonobilis/adventofcode2022/utils"
)

type Resource int

const (
	Ore = iota
	Clay
	Obsidian
	Geode
)

type Resources struct {
	ore, clay, obsidian, geode int
}

type OreRobot struct {
	oreCost int
}

type ClayRobot struct {
	oreCost int
}

type ObsidianRobot struct {
	oreCost  int
	clayCost int
}

type GeodeRobot struct {
	oreCost      int
	obsidianCost int
}

type Blueprint struct {
	oreRobot      OreRobot
	clayRobot     ClayRobot
	obsidianRobot ObsidianRobot
	geodeRobot    GeodeRobot
}

func parseInput(input string) (output []Blueprint) {
	for _, s := range strings.Split(input, "\n") {
		re := regexp.MustCompile(`[-\d]+`)
		s2 := re.FindAllString(s, 7)
		output = append(output, Blueprint{
			oreRobot:      OreRobot{utils.Atoi(s2[1])},
			clayRobot:     ClayRobot{utils.Atoi(s2[2])},
			obsidianRobot: ObsidianRobot{utils.Atoi(s2[3]), utils.Atoi(s2[4])},
			geodeRobot:    GeodeRobot{utils.Atoi(s2[5]), utils.Atoi(s2[6])},
		})
	}
	return
}

type CacheItem struct {
	minute           int
	resources, robot Resources
}

func factorial(number int) int {
	result := 1
	if number > 0 {
		for i := 1; i <= number; i++ {
			result *= i

		}
	}
	return result
}

func startPprocess(minute int, blueprint Blueprint, _resources, _robots Resources) int {
	var process func(minute int, resources, robots Resources)
	cache := make(utils.Set[CacheItem])
	best := 0
	process = func(minute int, resources, robots Resources) {
		if cache.Contains(CacheItem{minute, resources, robots}) {
			return
		}

		if minute == 0 {
			if resources.geode > best {
				best = resources.geode
			}
			return
		}

		if minute < 10 && minute*robots.geode+resources.geode+factorial(minute) < best {
			return
		}

		if resources.ore >= blueprint.geodeRobot.oreCost && resources.obsidian >= blueprint.geodeRobot.obsidianCost {
			newResources := Resources{
				resources.ore + robots.ore - blueprint.geodeRobot.oreCost,
				resources.clay + robots.clay,
				resources.obsidian + robots.obsidian - blueprint.geodeRobot.obsidianCost,
				resources.geode + robots.geode,
			}

			newRobots := Resources{
				robots.ore,
				robots.clay,
				robots.obsidian,
				robots.geode + 1,
			}
			process(minute-1, newResources, newRobots)
			cache.Add(CacheItem{minute, newResources, newRobots})
		} else if resources.ore >= blueprint.obsidianRobot.oreCost && resources.clay >= blueprint.obsidianRobot.clayCost {
			newResources := Resources{
				resources.ore + robots.ore - blueprint.obsidianRobot.oreCost,
				resources.clay + robots.clay - blueprint.obsidianRobot.clayCost,
				resources.obsidian + robots.obsidian,
				resources.geode + robots.geode,
			}
			newRobots := Resources{
				robots.ore,
				robots.clay,
				robots.obsidian + 1,
				robots.geode,
			}
			process(minute-1, newResources, newRobots)
			cache.Add(CacheItem{minute, newResources, newRobots})
		} else {
			if resources.ore >= blueprint.clayRobot.oreCost {
				newResources := Resources{
					resources.ore + robots.ore - blueprint.clayRobot.oreCost,
					resources.clay + robots.clay,
					resources.obsidian + robots.obsidian,
					resources.geode + robots.geode,
				}
				newRobots := Resources{
					robots.ore,
					robots.clay + 1,
					robots.obsidian,
					robots.geode,
				}
				process(minute-1, newResources, newRobots)
				cache.Add(CacheItem{minute, newResources, newRobots})
			}
			if resources.ore >= blueprint.oreRobot.oreCost {
				newResources := Resources{
					resources.ore + robots.ore - blueprint.oreRobot.oreCost,
					resources.clay + robots.clay,
					resources.obsidian + robots.obsidian,
					resources.geode + robots.geode,
				}
				newRobots := Resources{
					robots.ore + 1,
					robots.clay,
					robots.obsidian,
					robots.geode,
				}
				process(minute-1, newResources, newRobots)
				cache.Add(CacheItem{minute, newResources, newRobots})
			}

			newResources := Resources{
				resources.ore + robots.ore,
				resources.clay + robots.clay,
				resources.obsidian + robots.obsidian,
				resources.geode + robots.geode,
			}
			process(minute-1, newResources, robots)
			cache.Add(CacheItem{minute, newResources, robots})
		}
	}

	process(minute, _resources, _robots)
	return best
}

func p1(inp []Blueprint) (sum int) {
	sumChan := make(chan int)
	for i, b := range inp {
		go func(i int, b Blueprint) {
			sumChan <- (i + 1) * startPprocess(24, b, Resources{0, 0, 0, 0}, Resources{1, 0, 0, 0})
		}(i, b)
	}
	for i := 0; i < len(inp); i++ {
		sum += <-sumChan
	}
	return
}

func p2(inp []Blueprint) int {
	inp = inp[:3]
	product := 1
	productChan := make(chan int)
	for i, b := range inp {
		go func(i int, b Blueprint) {
			productChan <- startPprocess(32, b, Resources{0, 0, 0, 0}, Resources{1, 0, 0, 0})
		}(i, b)
	}
	for i := 0; i < len(inp); i++ {
		product *= <-productChan
	}
	return product
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	inp := parseInput(string(input))
	fmt.Printf("Part 1: %v\n", p1(inp))
	fmt.Printf("Part 2: %v\n", p2(inp))
}
