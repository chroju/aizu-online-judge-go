package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func push(queue []string, args []string, head int, tail int) ([]string, int, int) {
	if args[0] == "0" {
		head--
		queue[head] = args[1]
	} else {
		queue[tail] = args[1]
		tail++
	}
	return queue, head, tail
}

func pop(arg string, head int, tail int) (int, int) {
	if arg == "0" {
		head++
	} else if arg == "1" {
		tail--
	}

	return head, tail
}

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	length, _ := strconv.Atoi(sc.Text())

	var queue = make([]string, length*2)
	head := length
	tail := head
	for sc.Scan() {
		inputs := strings.Split(sc.Text(), " ")

		switch inputs[0] {
		case "0":
			queue, head, tail = push(queue, inputs[1:], head, tail)
		case "1":
			num, _ := strconv.Atoi(inputs[1])
			fmt.Println(queue[head+num])
		case "2":
			head, tail = pop(inputs[1], head, tail)
		}
	}
}
