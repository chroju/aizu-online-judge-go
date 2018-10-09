package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var inputs [4]float64
	if sc.Scan() {
		for i, v := range strings.Split(sc.Text(), " ") {
			inputs[i], _ = strconv.ParseFloat(v, 64)
		}
	}

	x := inputs[2] - inputs[0]
	y := inputs[3] - inputs[1]
	fmt.Println(math.Sqrt(x*x + y*y))
}
