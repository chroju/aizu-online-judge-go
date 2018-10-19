package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 1000)
	sc.Buffer(buf, 100000000)

	for i := 0; i < 2; i++ {
		sc.Scan()
	}
	listA := strings.Split(sc.Text(), " ")
	for i := 0; i < 2; i++ {
		sc.Scan()
	}
	listB := strings.Split(sc.Text(), " ")

	for _, v := range listB {
		num := contains(v, listA)
		if num == -1 {
			fmt.Println(0)
			os.Exit(0)
		}
		listA = listA[num:]
	}
	fmt.Println(1)
}

func contains(s string, list []string) int {
	for i, v := range list {
		if s == v {
			return i
		}
	}
	return -1
}
