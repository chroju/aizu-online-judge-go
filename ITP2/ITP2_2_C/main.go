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
	inputs := strings.Split(sc.Text(), " ")
	num, _ := strconv.Atoi(inputs[0])
	queries, _ := strconv.Atoi(inputs[1])
	queues := make([][]int, num)
	tail := make([]int, num)
	for i := range queues {
		queues[i] = make([]int, queries)
		tail[i] = 0
	}

	for sc.Scan() {
		inputs = strings.Split(sc.Text(), " ")
		n, _ := strconv.Atoi(inputs[1])
		fmt.Println(queues[999])

		switch inputs[0] {
		case "0":
			v, _ := strconv.Atoi(inputs[2])
			i := tail[n]
			for {
				if i == 0 || queues[n][i/2] > v {
					break
				} else {
					queues[n][i] = queues[n][i/2]
					i = i / 2
				}
			}
			queues[n][i] = v
			tail[n]++
		case "1":
			if tail[n] != 0 {
				fmt.Println(queues[n][0])
			}
		case "2":
			queues[n][0] = queues[n][tail-1]
			queues[n][tail-1] = 0
			i := 0
			for {
				if (i+1)*2-1 >= tail || queues[n][i] > queues[n]
			}
			for i := 0; i < tail[n]-1; i++ {
				queues[n][i] = queues[n][i+1]
			}
			tail[n]--
		}
	}
}
