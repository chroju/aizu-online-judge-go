package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func rotateDice(dice []string, direction string) []string {
	var result []string
	if direction == "N" {
		result = []string{dice[1], dice[5], dice[2], dice[3], dice[0], dice[4]}
	} else if direction == "E" {
		result = []string{dice[3], dice[1], dice[0], dice[5], dice[4], dice[2]}
	} else if direction == "W" {
		result = []string{dice[2], dice[1], dice[5], dice[0], dice[4], dice[3]}
	} else if direction == "S" {
		result = []string{dice[4], dice[0], dice[2], dice[3], dice[5], dice[1]}
	}
	return result
}

func main() {
	var input string
	var dice []string
	i := 0
	for sc.Scan() {
		if i == 0 {
			dice = strings.Split(sc.Text(), " ")
		} else if i == 1 {
			input = sc.Text()
		}
		i++
	}

	for _, i := range strings.Split(input, "") {
		dice = rotateDice(dice, i)
	}

	fmt.Printf("%v\n", dice)

}
