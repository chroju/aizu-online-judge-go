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
	buf := make([]byte, 100)
	sc.Buffer(buf, 1000000)
	sc.Scan()
	length, _ := strconv.Atoi(sc.Text())

	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")
	lastValue, _ := strconv.Atoi(inputs[length-1])

	sc.Scan()
	for sc.Scan() {
		q, _ := strconv.Atoi(sc.Text())
		if q > lastValue {
			fmt.Println(0)
			continue
		}

		low := 0
		high := length - 1
		for {
			mid := (low + high) / 2
			midValue, _ := strconv.Atoi(inputs[mid])
			// fmt.Println(q, midValue)
			if q > midValue {
				low = mid + 1
			} else if q < midValue {
				high = mid - 1
			} else {
				fmt.Println(1)
				break
			}

			if low > high {
				fmt.Println(0)
				break
			}
		}
	}
}
