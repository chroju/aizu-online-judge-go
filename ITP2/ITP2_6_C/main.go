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
	tmp := strings.Split(sc.Text(), " ")
	inputs := make([]int, length)
	for i, v := range tmp {
		inputs[i], _ = strconv.Atoi(v)
	}

	sc.Scan()
	for sc.Scan() {
		q, _ := strconv.Atoi(sc.Text())
		if inputs[len(inputs)-1] < q {
			fmt.Println(length)
			continue
		} else if inputs[0] >= q {
			fmt.Println(0)
			continue
		}
		low := 0
		high := length - 1
		target := 0
		for {
			mid := (low + high) / 2
			midValue := inputs[mid]
			// fmt.Println(q, midValue)
			if q > midValue {
				low = mid + 1
			} else if q < midValue {
				high = mid - 1
			} else {
				target = mid
				break
			}

			if low > high {
				target = high
				break
			}

		}

		for {
			if inputs[target] < q {
				break
			}
			target--
		}

		fmt.Println(target + 1)
	}
}
