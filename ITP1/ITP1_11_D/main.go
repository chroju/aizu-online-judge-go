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
	var num int 
	dices := []
	if sc.Scan() {
		input := sc.Text()
		for j, v := range strings.Split(input, " ") {
			i[j], _ = strconv.Atoi(v)
		}
	}

	fmt.Printf("%d %d\n", i[0]*i[1], (i[0]+i[1])*2)
}
