package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	input := strings.Split(sc.Text(), " ")
	a, _ := strconv.Atoi(input[0])
	b, _ := strconv.Atoi(input[1])
	bitsA := toBitSet(a)
	bitsB := toBitSet(b)

	for i, v := range bitsA {
		if v == 1 && bitsB[i] == 1 {
			fmt.Printf("%d", 1)
		} else {
			fmt.Printf("%d", 0)
		}
	}
	fmt.Println("")

	for i, v := range bitsA {
		if v == 1 || bitsB[i] == 1 {
			fmt.Printf("%d", 1)
		} else {
			fmt.Printf("%d", 0)
		}
	}
	fmt.Println("")

	for i, v := range bitsA {
		if v == bitsB[i] {
			fmt.Printf("%d", 0)
		} else {
			fmt.Printf("%d", 1)
		}
	}
	fmt.Println("")

}

func toBitSet(d int) [32]int {
	result := [32]int{}
	for i := 0; i < 32; i++ {
		j := int(math.Pow(2, float64(31-i)))
		if d >= j {
			d = d - j
			result[i] = 1
		}
	}
	return result
}
