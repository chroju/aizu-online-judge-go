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

	var data []string

	for sc.Scan() {
		inputs := strings.Split(sc.Text(), " ")

		switch inputs[0] {
		case "0":
			data = append(data, inputs[1])
		case "1":
			i, _ := strconv.Atoi(inputs[1])
			fmt.Println(data[i])
		case "2":
			data = data[:len(data)-1]
		}
	}
}
