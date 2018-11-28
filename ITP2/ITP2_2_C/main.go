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
		if n == 23 {
			fmt.Println(inputs)
			fmt.Println(queues[n])
		}

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
			if tail[n] != 0 {
				tail[n]--
				queues[n][0] = queues[n][tail[n]]
				queues[n][tail[n]] = 0

				i := 0
				for {
					child1 := (i + 1) * 2
					child2 := child1 - 1
					cf := 0
					if child2 >= tail[n] {
						break
					} else if child1 == tail[n] || queues[n][child2] > queues[n][child1] {
						cf = child2
					} else {
						cf = child1
					}

					if queues[n][cf] <= queues[n][i] {
						break
					} else {
						tmp := queues[n][cf]
						queues[n][cf] = queues[n][i]
						queues[n][i] = tmp
						i = cf
					}
				}
			}
		}
	}
}
