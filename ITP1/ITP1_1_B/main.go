package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var i int
	if sc.Scan() {
		i, _ = strconv.Atoi(sc.Text())
	}

	if i < 0 || i > 100 {
		panic("input is wrong")
	}

	result := math.Pow(float64(i), 3)
	fmt.Println(result)
}
