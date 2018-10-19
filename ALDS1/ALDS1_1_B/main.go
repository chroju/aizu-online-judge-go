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
	inputs := strings.Split(sc.Text(), " ")
	x, _ := strconv.Atoi(inputs[0])
	y, _ := strconv.Atoi(inputs[1])

	var r int
	for {
		r = x % y
		if r == 0 {
			fmt.Println(y)
			break
		}
		x = y
		y = r
	}
}
