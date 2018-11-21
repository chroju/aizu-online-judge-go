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

	for sc.Scan() {
		inputs := strings.Split(sc.Text(), " ")
		mid, _ := strconv.Atoi(inputs[0])
		end, _ := strconv.Atoi(inputs[1])
		re, _ := strconv.Atoi(inputs[2])

		if mid == -1 && end == -1 && re == -1 {
			break
		} else if mid == -1 || end == -1 {
			fmt.Println("F")
		} else if mid+end >= 80 {
			fmt.Println("A")
		} else if mid+end >= 65 {
			fmt.Println("B")
		} else if mid+end >= 50 {
			fmt.Println("C")
		} else if mid+end >= 30 {
			if re >= 50 {
				fmt.Println("C")
			} else {
				fmt.Println("D")
			}
		} else {
			fmt.Println("F")
		}
	}
}
