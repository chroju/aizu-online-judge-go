package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var buf = make([]byte, 1000000)

func main() {
	sc.Buffer(buf, 1000000)

	var num int
	if sc.Scan() {
		num, _ = strconv.Atoi(sc.Text())
	}
	var seq = make([]int, num)

	if sc.Scan() {
		for i, v := range strings.Split(sc.Text(), " ") {
			seq[i], _ = strconv.Atoi(v)
		}
	}

	var min, max, sum int
	for _, i := range seq {
		sum += i
		if min == 0 || min > i {
			min = i
		}
		if max == 0 || max < i {
			max = i
		}
	}

	fmt.Printf("%d %d %d\n", min, max, sum)

}
