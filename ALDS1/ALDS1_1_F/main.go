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
	sc.Scan()
	var inputs []int
	for _, v := range strings.Split(sc.Text(), " ") {
		tmp, _ := strconv.Atoi(v)
		inputs = append(inputs, tmp)
	}

	var count int
	for i := range inputs {
		minj := i
		for j := i; j < len(inputs); j++ {
			if inputs[j] < inputs[minj] {
				minj = j
			}
		}
		if i != minj {
			tmp := inputs[i]
			inputs[i] = inputs[minj]
			inputs[minj] = tmp
			count++
		}
	}

	fmt.Printf("%d", inputs[0])
	for _, v := range inputs[1:] {
		fmt.Printf(" %d", v)
	}
	fmt.Printf("\n%d\n", count)
}
