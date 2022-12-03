package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	p1()
	p2()
}

func p2() {
	f, _ := os.Open("./input.txt")
	fileScanner := bufio.NewScanner(f)

	fileScanner.Split(bufio.ScanLines)

	sum := 0
	fileScanner.Scan()
	for {
		group := make([]string, 3)
		eof := false
		for i := 0; i < 3; i++ {
			txt := fileScanner.Text()
			if !fileScanner.Scan() {
				eof = true
			}
			group[i] = strings.TrimSpace(txt)
		}
		sum += prio(diffGroup(group))
		if eof {
			break
		}
	}
	fmt.Println(sum)
}

func p1() {
	f, _ := os.Open("./input.txt")
	fileScanner := bufio.NewScanner(f)

	fileScanner.Split(bufio.ScanLines)

	sum := 0
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		txt = strings.TrimSpace(txt)
		l := txt[:len(txt)/2]
		r := txt[len(txt)/2:]
		sum += prio(diff(l, r))
	}
	fmt.Println(sum)
}

func diff(l, r string) rune {
	for _, c := range l {
		if strings.Contains(r, string(c)) {
			return c
		}
	}
	return '0' // never happens
}

func prio(c rune) int {
	if unicode.IsUpper(c) {
		return int(c - 64 + 26)
	} else {
		return int(c - 96)
	}
}

func diffGroup(group []string) rune {
	for _, c := range group[0] {
		if strings.Contains(group[1], string(c)) {
			if strings.Contains(group[2], string(c)) {
				return c
			}
		}
	}
	return '0' // never happens
}
