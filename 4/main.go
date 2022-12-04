package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	for fileScanner.Scan() {
		txt := strings.TrimSpace(fileScanner.Text())
		elfs := strings.Split(txt, ",")
		areas0 := strings.Split(elfs[0], "-")
		areas1 := strings.Split(elfs[1], "-")
		iareas0 := fillArea(areas0)
		iareas1 := fillArea(areas1)
		if checkContainments2(iareas0, iareas1) {
			sum++
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
		txt := strings.TrimSpace(fileScanner.Text())
		elfs := strings.Split(txt, ",")
		areas0 := strings.Split(elfs[0], "-")
		areas1 := strings.Split(elfs[1], "-")
		iareas0 := fillArea(areas0)
		iareas1 := fillArea(areas1)
		if checkContainments(iareas0, iareas1) {
			sum++
		}

	}
	fmt.Println(sum)
}

func fillArea(a []string) []int {
	areas := make([]int, 0)
	start, _ := strconv.Atoi(a[0])
	end, _ := strconv.Atoi(a[1])
	for i := start; i <= end; i++ {
		areas = append(areas, i)
	}
	return areas
}

func checkContainments(l, r []int) bool {
	lcount := 0
	for _, v := range l {
		for _, w := range r {
			if v == w {
				lcount++
			}
		}
	}
	rcount := 0
	for _, v := range r {
		for _, w := range l {
			if v == w {
				rcount++
			}
		}
	}
	if lcount == len(l) || rcount == len(r) {
		return true
	}
	return false
}

func checkContainments2(l, r []int) bool {
	lcount := 0
	for _, v := range l {
		for _, w := range r {
			if v == w {
				lcount++
			}
		}
	}
	rcount := 0
	for _, v := range r {
		for _, w := range l {
			if v == w {
				rcount++
			}
		}
	}
	if lcount > 0 || rcount > 0 {
		return true
	}
	return false
}
