package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	i := 0
	for sc.Scan() {
		v, _ := strconv.Atoi(sc.Text())
		if v%2 == 0 && v != 2 {
			continue
		}

		if judgePrime(v) {
			i++
		}
	}
	fmt.Println(i)
}

func judgePrime(v int) bool {
	root := math.Sqrt(float64(v))
	for j := 3; float64(j) <= root; j = j + 2 {
		if v%j == 0 {
			return false
		}
	}
	return true
}
