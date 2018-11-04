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

	S := map[string]bool{}
	for i := 0; i < n; i++ {
		sc.Scan()
		key := sc.Text()
		if _, ok := S[key]; !ok {
			S[key] = true
		}
	}

	sc.Scan()
	result := 0
	for sc.Scan() {
		key := sc.Text()
		if _, ok := S[key]; ok {
			result++
		}
	}
	fmt.Println(result)
}
