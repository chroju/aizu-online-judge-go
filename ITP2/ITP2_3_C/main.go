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
	num, _ := strconv.Atoi(sc.Text())
	var A = make([]int, num)

	sc.Scan()
	for i, v := range strings.Split(sc.Text(), " ") {
		A[i], _ = strconv.Atoi(v)
	}

	sc.Scan()

	for sc.Scan() {
		inputs := strings.Split(sc.Text(), " ")
		start, _ := strconv.Atoi(inputs[0])
		end, _ := strconv.Atoi(inputs[1])
		k, _ := strconv.Atoi(inputs[2])

		result := 0
		for _, v := range A[start:end] {
			if v == k {
				result++
			}
		}
		fmt.Println(result)
	}
}
