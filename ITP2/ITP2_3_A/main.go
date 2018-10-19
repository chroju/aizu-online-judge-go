package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var inputs [3]int
	if sc.Scan() {
		for i, v := range strings.Split(sc.Text(), " ") {
			inputs[i], _ = strconv.Atoi(v)
		}

		min := inputs[0]
		max := inputs[0]
		for _, v := range inputs {
			if v < min {
				min = v
			} else if v > max {
				max = v
			}
		}

		fmt.Printf("%d %d\n", min, max)

	}
}
