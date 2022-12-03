package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	f, _ := os.Open("./input.txt")
	fileScanner := bufio.NewScanner(f)

	fileScanner.Split(bufio.ScanLines)

	sums := make([]int, 0)

	tmp := make([]int, 0)
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		if txt != "" {
			i, _ := strconv.Atoi(txt)
			tmp = append(tmp, i)
		} else {
			sum := 0
			for _, v := range tmp {
				sum += v
			}
			sums = append(sums, sum)
			tmp = make([]int, 0)
		}
	}
	max := make([]int, 3)
	for n := 0; n < 3; n++ {
		for i, e := range sums {
			if i == 0 || e > max[n] {
				max[n] = e
			}
		}
		for i, e := range sums {
			if e == max[n] {
				sums = append(sums[:i], sums[i+1:]...)
			}
		}
	}

	total := 0
	for _, e := range max {
		total += e
	}

	fmt.Println(max)
	fmt.Println(total)
}
