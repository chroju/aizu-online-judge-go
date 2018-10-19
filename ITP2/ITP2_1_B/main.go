package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func push(list []string, args []string) []string {
	if args[0] == "0" {
		length := len(list)
		if length == 0 {
			list = []string{args[1]}
		} else {
			list = append(list, list[length-1])
			for i := length; i > 0; i-- {
				list[i] = list[i-1]
			}
			list[0] = args[1]
		}
	} else if args[0] == "1" {
		list = append(list, args[1])
	}

	return list
}

func pop(list []string, arg string) []string {
	if arg == "0" {
		length := len(list)
		for i := 0; i < length-1; i++ {
			list[i] = list[i+1]
		}
	} else if arg == "1" {
		list = list[:len(list)-1]
	}
	return list
}

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	var list = []string{}

	var i = 0
	for sc.Scan() {
		i++
		if i == 1 {
			continue
		}
		inputs := strings.Split(sc.Text(), " ")

		if inputs[0] == "0" {
			list = push(list, inputs[1:])
		} else if inputs[0] == "1" {
			if list != nil {
				num, _ := strconv.Atoi(inputs[1])
				fmt.Println(list)
				fmt.Printf("%s\n", list[num])
			}
		} else if inputs[0] == "2" {
			list = pop(list, inputs[1])
		}
	}
}
