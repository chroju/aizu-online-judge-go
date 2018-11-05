package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	var stack []string
	for sc.Scan() {
		i := sc.Text()
		if i == "+" || i == "-" || i == "*" {
			a, _ := strconv.Atoi(stack[len(stack)-2])
			b, _ := strconv.Atoi(stack[len(stack)-1])
			if len(stack) > 2 {
				stack = stack[:len(stack)-2]
			} else {
				stack = []string{}
			}
			if i == "+" {
				stack = append(stack, strconv.Itoa(a+b))
			} else if i == "-" {
				stack = append(stack, strconv.Itoa(a-b))
			} else if i == "*" {
				stack = append(stack, strconv.Itoa(a*b))
			}
		} else {
			stack = append(stack, i)
		}
	}
	fmt.Println(stack[0])
}
