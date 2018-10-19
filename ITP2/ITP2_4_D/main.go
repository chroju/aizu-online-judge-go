package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var former string
	var i int
	sc.Split(bufio.ScanWords)
	sc.Scan()

	for sc.Scan() {
		v := sc.Text()
		if i == 0 {
			fmt.Printf("%s", v)
		} else if former != v {
			fmt.Printf(" %s", v)
		}
		former = v
		i++
	}
	fmt.Print("\n")
}
