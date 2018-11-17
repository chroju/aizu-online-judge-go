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
	lenA, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	A := strings.Split(sc.Text(), " ")

	sc.Scan()
	lenB, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	B := strings.Split(sc.Text(), " ")

	shorter := lenA
	short := 1
	if lenB < lenA {
		shorter = lenB
		short = 0
	}

	for i := 0; i < shorter; i++ {
		if A[i] > B[i] {
			fmt.Println(0)
			os.Exit(0)
		} else if A[i] < B[i] {
			fmt.Println(1)
			os.Exit(0)
		}
	}

	fmt.Println(short)
}
