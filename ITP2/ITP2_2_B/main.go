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

	var queues = make([][]string, num)
	for i := range queues {
		queues[i] = []string{}
	}

	for sc.Scan() {
		inputs := strings.Split(sc.Text(), " ")
		i, _ := strconv.Atoi(inputs[1])

		switch inputs[0] {
		case "0":
			queues[i] = append(queues[i], inputs[2])
		case "1":
			if len(queues[i]) > 0 {
				fmt.Println(queues[i][0])
			}
		case "2":
			if len(queues[i]) > 0 {
				queues[i] = queues[i][1:]
			}
		}
	}
}
