package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type proc struct {
	name string
	time int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	quantum, _ := strconv.Atoi(strings.Split(sc.Text(), " ")[1])

	var queue []proc
	for sc.Scan() {
		inputs := strings.Split(sc.Text(), " ")
		time, _ := strconv.Atoi(inputs[1])
		queue = append(queue, proc{inputs[0], time})
	}

	var processedTime int
	for {
		restTime := queue[0].time - quantum
		if restTime <= 0 {
			processedTime = processedTime + queue[0].time
			fmt.Printf("%s %d\n", queue[0].name, processedTime)
			queue = queue[1:]
		} else {
			queue = append(queue[1:], queue[0])
			queue[len(queue)-1].time = restTime
			processedTime = processedTime + quantum
		}

		if len(queue) == 0 {
			break
		}
	}
}
