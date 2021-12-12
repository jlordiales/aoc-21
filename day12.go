package aoc_2021

import (
	"strings"
)

type caveSystem struct {
	visited map[string]int
	caves   map[string][]string
}

func (c *caveSystem) addLink(orig, dest string) {
	c.caves[orig] = append(c.caves[orig], dest)
	c.caves[dest] = append(c.caves[dest], orig)
}

func buildCaveSystem(input []string) caveSystem {
	res := caveSystem{caves: make(map[string][]string), visited: make(map[string]int)}
	for _, l := range input {
		split := strings.Split(l, "-")
		res.addLink(split[0], split[1])
	}

	return res
}

func day12FirstHalf(input []string) int {
	caves := buildCaveSystem(input)

	caves.visited["start"] = 1
	paths := allPaths(caves, "start", "end", "start", shouldVisit)
	return len(paths)
}

func day12SecondHalf(input []string) int {
	caves := buildCaveSystem(input)

	caves.visited["start"] = 1
	paths := allPaths(caves, "start", "end", "start", shouldVisit2)
	return len(paths)
}

func allPaths(system caveSystem, start, end, currentPath string, visitFn func(caveSystem, string) bool) []string {
	var res []string
	for _, neighbor := range system.caves[start] {
		if neighbor == end {
			return append(res, currentPath+","+end)
		}

		if !visitFn(system, neighbor) {
			continue
		}

		system.visited[neighbor]++

		pathsFromNeighbor := allPaths(system, neighbor, end, currentPath+","+neighbor, visitFn)
		res = append(res, pathsFromNeighbor...)

		system.visited[neighbor]--
	}

	return res
}

func shouldVisit(system caveSystem, cave string) bool {
	return system.visited[cave] == 0 || strings.ToUpper(cave) == cave
}

func shouldVisit2(system caveSystem, cave string) bool {
	if system.visited[cave] == 0 || strings.ToUpper(cave) == cave {
		return true
	}

	if cave == "start" {
		return false
	}

	// if any of the small caves was already visited more than once
	// (including the current one we're evaluating) we can't visit again
	for k, v := range system.visited {
		if v > 1 && strings.ToLower(k) == k {
			return false
		}
	}

	return true
}
