package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())

	fmt.Println(fibonacci(num))
}

func fibonacci(num int) int {
	fibs := make([]int, num+1)

	for i := range fibs {
		if i <= 1 {
			fibs[i] = 1
		} else {
			fibs[i] = fibs[i-1] + fibs[i-2]
		}
	}

	return fibs[num]
}
