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

	var stacks = make([][]string, num)
	for i := range stacks {
		stacks[i] = []string{}
	}

	for sc.Scan() {
		inputs := strings.Split(sc.Text(), " ")
		i, _ := strconv.Atoi(inputs[1])

		switch inputs[0] {
		case "0":
			stacks[i] = append(stacks[i], inputs[2])
		case "1":
			if len(stacks[i]) > 0 {
				fmt.Println(stacks[i][len(stacks[i])-1])
			}
		case "2":
			if len(stacks[i]) > 0 {
				stacks[i] = stacks[i][:len(stacks[i])-1]
			}
		}
	}
}
