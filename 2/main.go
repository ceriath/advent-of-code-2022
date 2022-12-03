package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("./input.txt")
	fileScanner := bufio.NewScanner(f)

	fileScanner.Split(bufio.ScanLines)

	sum1 := 0
	sum2 := 0
	for fileScanner.Scan() {
		txt := fileScanner.Text()
		in := strings.Split(txt, " ")
		sum1 += calcResult(in[0], in[1])
		sum2 += calcResult2(in[0], in[1])
	}
	fmt.Println(sum1)
	fmt.Println(sum2)
}

// A X Rock
// B Y Paper
// C Z Scissors
func calcResult(o, s string) int {
	points := 0
	if s == "X" {
		points += 1
	} else if s == "Y" {
		points += 2
	} else if s == "Z" {
		points += 3
	}

	if o == "A" && s == "X" || o == "B" && s == "Y" || o == "C" && s == "Z" {
		// draw
		points += 3
	} else if (o == "A" && s == "Y") || (o == "B" && s == "Z") || (o == "C" && s == "X") {
		// win
		points += 6
	} else {
		// loss
		points += 0
	}
	return points
}

func calcResult2(o, s string) int {
	points := 0

	switch o {
	case "A":
		switch s {
		case "X":
			points += calcResult("A", "Z")
		case "Y":
			points += calcResult("A", "X")
		case "Z":
			points += calcResult("A", "Y")
		}
	case "B":
		switch s {
		case "X":
			points += calcResult("B", "X")
		case "Y":
			points += calcResult("B", "Y")
		case "Z":
			points += calcResult("B", "Z")
		}
	case "C":
		switch s {
		case "X":
			points += calcResult("C", "Y")
		case "Y":
			points += calcResult("C", "Z")
		case "Z":
			points += calcResult("C", "X")
		}
	}
	return points
}
