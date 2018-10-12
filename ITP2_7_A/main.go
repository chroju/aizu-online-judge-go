package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	s := make(map[string]bool)
	sum := 0
	for sc.Scan() {
		inputs := strings.Split(sc.Text(), " ")

		if inputs[0] == "0" {
			if !s[inputs[1]] {
				s[inputs[1]] = true
				sum++
			}
			fmt.Printf("%d\n", sum)
		} else if inputs[0] == "1" {
			if s[inputs[1]] {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		}
	}
}
