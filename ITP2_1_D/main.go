package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	num, _ := strconv.Atoi(strings.Split(sc.Text(), " ")[0])
	lists := make([][]string, num)

	for sc.Scan() {
		inputs := strings.Split(sc.Text(), " ")
		num, _ := strconv.Atoi(inputs[1])

		if inputs[0] == "0" {
			lists[num] = append(lists[num], inputs[2])
		} else if inputs[0] == "1" {
			fmt.Printf("%s\n", strings.Join(lists[num], " "))
		} else if inputs[0] == "2" {
			lists[num] = []string{}
		}
	}
}
