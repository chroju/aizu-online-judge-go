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
	list := strings.Split(sc.Text(), " ")

	sc.Scan()

	for sc.Scan() {
		inputs := strings.Split(sc.Text(), " ")
		begin, _ := strconv.Atoi(inputs[0])
		end, _ := strconv.Atoi(inputs[1])
		end--

		for {
			if begin == end || begin > end {
				break
			}
			tmp := list[begin]
			list[begin] = list[end]
			list[end] = tmp
			begin++
			end--
		}
	}

	fmt.Println(strings.Trim(fmt.Sprint(list), "[]"))
}
