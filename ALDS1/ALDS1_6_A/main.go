package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	max = 10000
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	length, _ := strconv.Atoi(sc.Text())

	A := make([]int, length)
	B := make([]int, length)
	i := 0
	for sc.Scan() {
		A[i], _ = strconv.Atoi(sc.Text())
		i++
	}

	C := make([]int, max)
	for i = 0; i < len(A); i++ {
		C[A[i]]++
	}

	for i = 1; i < max; i++ {
		C[i] = C[i] + C[i-1]
	}

	for i = length - 1; i >= 0; i-- {
		B[C[A[i]]-1] = A[i]
		C[A[i]]--
	}

	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(B)), " "), "[]"))
}
