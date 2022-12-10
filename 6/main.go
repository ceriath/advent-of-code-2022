package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var stacks = map[int][]string{}

func main() {
	p1()
	p2()
}

func p1() {
	f, _ := os.Open("./input.txt")
	fileScanner := bufio.NewScanner(f)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		window := make([]rune, 4)
		txt := strings.TrimSpace(fileScanner.Text())
		for i := 0; i < len(window); i++ {
			window[i] = rune(txt[i])
		}
		for i := 3; i < len(txt); i++ {
			shiftWindow(window, rune(txt[i]))
			if checkWindow(window) {
				fmt.Printf("Found unique sequence %s after %d", string(window), i+1)
				return
			}
		}
	}
}

func p2() {
	f, _ := os.Open("./input.txt")
	fileScanner := bufio.NewScanner(f)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		window := make([]rune, 14)
		txt := strings.TrimSpace(fileScanner.Text())
		for i := 0; i < len(window); i++ {
			window[i] = rune(txt[i])
		}
		for i := 13; i < len(txt); i++ {
			shiftWindow(window, rune(txt[i]))
			if checkWindow(window) {
				fmt.Printf("Found unique sequence %s after %d", string(window), i+1)
				return
			}
		}
	}
}

func shiftWindow(window []rune, c rune) {
	for i := 0; i < len(window)-1; i++ {
		window[i] = window[i+1]
	}
	window[len(window)-1] = c
}

func checkWindow(window []rune) bool {
	for i := 0; i < len(window); i++ {
		counter := 0
		for j := 0; j < len(window); j++ {
			if window[i] == window[j] {
				counter++
			}
		}
		if counter > 1 {
			return false
		}
	}
	return true
}
