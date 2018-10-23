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

	var flag bool
	var count int
	for {
		if flag {
			break
		}
		flag = true
		for i := range inputs[:len(inputs)-1] {
			if inputs[i] > inputs[i+1] {
				tmp := inputs[i]
				inputs[i] = inputs[i+1]
				inputs[i+1] = tmp
				flag = false
				count++
			}
		}
	}

	fmt.Printf("%d", inputs[0])
	for _, v := range inputs[1:] {
		fmt.Printf(" %d", v)
	}
	fmt.Printf("\n%d\n", count)

}
