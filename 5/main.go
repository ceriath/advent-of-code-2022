package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var stacks = map[int][]string{}

func main() {
	resetStacks()
	p1()
	resetStacks()
	p2()
}

func resetStacks() {
	stacks = map[int][]string{
		1: {"F", "H", "M", "T", "V", "L", "D"},
		2: {"P", "N", "T", "C", "J", "G", "Q", "H"},
		3: {"H", "P", "M", "D", "S", "R"},
		4: {"F", "V", "B", "L"},
		5: {"Q", "L", "G", "H", "N", "B"},
		6: {"P", "M", "R", "G", "D", "B", "W"},
		7: {"Q", "L", "H", "C", "R", "N", "M", "G"},
		8: {"W", "L", "C"},
		9: {"T", "M", "Z", "J", "Q", "L", "D", "R"},
	}
}

func p1() {
	f, _ := os.Open("./input.txt")
	fileScanner := bufio.NewScanner(f)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		txt := strings.TrimSpace(fileScanner.Text())
		procedure := strings.Split(txt, " ")
		count, _ := strconv.Atoi(procedure[1])
		sourceStack, _ := strconv.Atoi(procedure[3])
		targetStack, _ := strconv.Atoi(procedure[5])
		moveOneByOne(sourceStack, count, targetStack)
	}
	fmt.Println("Resulting stacks:")
	printStacks()
	fmt.Printf("result p1: ")
	for i := 1; i <= len(stacks); i++ {
		fmt.Printf("%v", stacks[i][0])
	}
	fmt.Println()
}

func moveOneByOne(sourceStack, count, targetStack int) {
	// fmt.Printf("move %d from %d to %d\n", count, sourceStack, targetStack)
	moveables := stacks[sourceStack][0:count]
	stacks[sourceStack] = stacks[sourceStack][count:]
	stacks[targetStack] = prependOnByOne(stacks[targetStack], moveables)
}

func prependOnByOne(x []string, y []string) []string {
	for i := 0; i < len(y); i++ {
		x = append(x, "")
		copy(x[1:], x)
		x[0] = y[i]
	}
	return x
}

func p2() {
	f, _ := os.Open("./input.txt")
	fileScanner := bufio.NewScanner(f)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		txt := strings.TrimSpace(fileScanner.Text())
		procedure := strings.Split(txt, " ")
		count, _ := strconv.Atoi(procedure[1])
		sourceStack, _ := strconv.Atoi(procedure[3])
		targetStack, _ := strconv.Atoi(procedure[5])
		moveAll(sourceStack, count, targetStack)
	}
	fmt.Println("Resulting stacks:")
	printStacks()
	fmt.Printf("result p2: ")
	for i := 1; i <= len(stacks); i++ {
		fmt.Printf("%v", stacks[i][0])
	}
	fmt.Println()
}

func moveAll(sourceStack, count, targetStack int) {
	moveables := stacks[sourceStack][0:count]
	stacks[sourceStack] = stacks[sourceStack][count:]
	stacks[targetStack] = prependAll(stacks[targetStack], moveables)
}

func prependAll(x []string, y []string) []string {
	for i := len(y); i > 0; i-- {
		x = append(x, "")
		copy(x[1:], x)
		x[0] = y[i-1]
	}
	return x
}

func printStacks() {
	for i := 1; i <= len(stacks); i++ {
		fmt.Printf("%d : %v\n", i, stacks[i])
	}
}
