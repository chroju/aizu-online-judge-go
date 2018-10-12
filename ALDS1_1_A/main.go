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
	sc.Buffer([]byte{}, 10000000000000000)
	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	inputs := strings.Split(sc.Text(), " ")

	var raw = make([]int, num)
	for i := range raw {
		raw[i], _ = strconv.Atoi(inputs[i])
	}

	for i := range raw {
		key := raw[i]
		for j := i - 1; j >= 0; j-- {
			if raw[j] > key {
				raw[j+1] = raw[j]
				if j == 0 {
					raw[0] = key
				}
			} else {
				raw[j+1] = key
				break
			}
		}
		fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(raw)), " "), "[]"))
	}

}
