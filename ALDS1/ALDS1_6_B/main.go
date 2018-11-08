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
	sc.Split(bufio.ScanWords)

	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())

	A := make([]int, num)
	i := 0
	for sc.Scan() {
		A[i], _ = strconv.Atoi(sc.Text())
		i++
	}

	criterion := A[len(A)-1]
	i = 0
	for j := 0; j < len(A)-1; j++ {
		if A[j] <= criterion {
			if j != 0 {
				tmp := A[i]
				A[i] = A[j]
				A[j] = tmp
			}
			i++
		}
	}
	A[len(A)-1] = A[i]

	fmt.Print(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(A[:i])), " "), "[]"))
	fmt.Printf(" [%d] ", criterion)
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(A[i+1:])), " "), "[]"))
}
