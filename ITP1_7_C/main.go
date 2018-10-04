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
	var column int
	if sc.Scan() {
		tmp := strings.Split(sc.Text(), " ")
		column, _ = strconv.Atoi(tmp[1])
	}

	var lastRaw = make([]int, column+1)
	for sc.Scan() {
		var newRaw = make([]int, column+1)
		tmp := strings.Split(sc.Text(), " ")
		for i, v := range tmp {
			newRaw[i], _ = strconv.Atoi(v)
		}
		for i, v := range newRaw {
			if i != column {
				newRaw[column] += v
				lastRaw[i] += v
			}
		}
		lastRaw[column] += newRaw[column]
		fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(newRaw)), " "), "[]"))
	}
	fmt.Println(strings.Trim(strings.Join(strings.Fields(fmt.Sprint(lastRaw)), " "), "[]"))
}
