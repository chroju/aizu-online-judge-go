package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sc = bufio.NewScanner(os.Stdin)
	sc.Scan()
	length, _ := strconv.Atoi(sc.Text())

	var list = make([]string, length*2)
	head := length
	// tail := length
	cursor := length + 1
	for sc.Scan() {
		// fmt.Println(list)
		inputs := strings.Split(sc.Text(), " ")

		switch inputs[0] {
		case "0":
			cursor--
			for i := head; i < cursor; i++ {
				list[i] = list[i+1]
			}
			head--
			list[cursor] = inputs[1]
		case "1":
			num, _ := strconv.Atoi(inputs[1])
			cursor = cursor + num
		case "2":
			for i := cursor; i > head+1; i-- {
				list[i] = list[i-1]
			}
			cursor++
			head++
			list[head] = ""
		}
	}

	for _, v := range list {
		if v != "" {
			fmt.Println(v)
		}
	}
}
