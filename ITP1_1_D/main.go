package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var i int
	if sc.Scan() {
		i, _ = strconv.Atoi(sc.Text())
	}

	hour := i / 3600
	min := (i % 3600) / 60
	sec := i - hour*3600 - min*60

	fmt.Printf("%d:%d:%d\n", hour, min, sec)
}
