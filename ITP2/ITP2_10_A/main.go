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
	input, _ := strconv.Atoi(sc.Text())

	bits := toBitSet(input)
	for i := 31; i >= 0; i-- {
		fmt.Printf("%d", bits[i])
	}
	fmt.Println("")

	for i := 31; i >= 0; i-- {
		if bits[i] == 1 {
			fmt.Printf("%d", 0)
		} else {
			fmt.Printf("%d", 1)
		}
	}
	fmt.Println("")

	for i := 30; i >= 0; i-- {
		fmt.Printf("%d", bits[i])
	}
	fmt.Println(0)

	fmt.Print(0)
	for i := 31; i > 0; i-- {
		fmt.Printf("%d", bits[i])
	}
	fmt.Println("")

}

func toBitSet(d int) [32]int {
	result := [32]int{}
	for i := 31; i >= 0; i-- {
		j := int(math.Pow(2, float64(i)))
		if d >= j {
			d = d - j
			result[i] = 1
		}
	}
	return result
}
