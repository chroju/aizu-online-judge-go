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

	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	var A = make([]int, n)
	for i := range A {
		sc.Scan()
		A[i], _ = strconv.Atoi(sc.Text())
	}

	sc.Scan()
	for sc.Scan() {
		m, _ := strconv.Atoi(sc.Text())
		if solve(0, m, A) {
			fmt.Println("yes")
		} else {
			fmt.Println("no")
		}
	}
}

func solve(i int, m int, A []int) bool {
	if m == 0 {
		return true
	}
	if i >= len(A) {
		return false
	}
	return solve(i+1, m, A) || solve(i+1, m-A[i], A)
}
