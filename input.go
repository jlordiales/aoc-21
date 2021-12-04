package aoc_2021

import (
	"io/ioutil"
	"strings"
)

func ReadLines(inputPath string) []string {
	data, _ := ioutil.ReadFile(inputPath)
	return strings.Split(string(data), "\n")
}
