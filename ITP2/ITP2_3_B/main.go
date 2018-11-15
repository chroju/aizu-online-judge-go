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
		start, _ := strconv.Atoi(inputs[1])
		end, _ := strconv.Atoi(inputs[2])

		result := A[start]
		switch inputs[0] {
		case "0":
			for i := start + 1; i < end; i++ {
				if result > A[i] {
					result = A[i]
				}
			}
		case "1":
			for i := start + 1; i < end; i++ {
				if result < A[i] {
					result = A[i]
				}
			}
		}
		fmt.Println(result)
	}
}
