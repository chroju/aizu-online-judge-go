package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	ponds := []int{0}
	currentPond := 0
	queue := []int{}

	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	inputs := strings.Split(sc.Text(), "")

	for _, v := range inputs {

		if v == "\\" {
			queue = append(queue, 0)
			for i := range queue {
				queue[i]++
			}
		} else if v == "_" {
			for i := range queue {
				queue[i]++
			}
		} else if v == "/" {
			if len(queue) != 0 {
				ponds[currentPond] = ponds[currentPond] + queue[len(queue)-1]
				queue = queue[:len(queue)-1]
				if len(queue) == 0 {
					ponds = append(ponds, 0)
					currentPond++
				} else {
					for i := range queue {
						queue[i]++
					}
				}
			}
		}
		fmt.Println(queue)
		fmt.Println(ponds)
	}

	firstRow := 0
	secondRow := strconv.Itoa(len(ponds) - 1)
	for _, v := range ponds[:len(ponds)-1] {
		firstRow = firstRow + v
		secondRow = secondRow + " " + strconv.Itoa(v)
	}

	fmt.Println(firstRow)
	fmt.Println(secondRow)
}
