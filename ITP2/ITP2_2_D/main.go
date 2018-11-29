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
	buf := make([]byte, 1000)
	sc.Buffer(buf, 10000000)

	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")
	num, _ := strconv.Atoi(inputs[0])
	queues := make([][]string, num)

	for sc.Scan() {
		inputs = strings.Split(sc.Text(), " ")
		n, _ := strconv.Atoi(inputs[1])

		switch inputs[0] {
		case "0":
			queues[n] = append(queues[n], inputs[2])
		case "1":
			fmt.Println(strings.Trim(fmt.Sprint(queues[n]), "[ ]"))
		case "2":
			t, _ := strconv.Atoi(inputs[2])
			newLength := len(queues[t]) + len(queues[n])
			tmp := make([]string, newLength)
			for i := range tmp {
				if i < len(queues[t]) {
					tmp[i] = queues[t][i]
				} else {
					tmp[i] = queues[n][i-len(queues[t])]
				}
			}
			queues[t] = tmp
			queues[n] = make([]string, 0)
		}
	}
}
