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

	dna := []string{"A", "C", "G", "T"}
	dict := map[string]bool{}
	for _, v1 := range dna {
		for _, v2 := range dna {
			for _, v3 := range dna {
				key := v1 + v2 + v3
				dict[key] = false
			}
		}
	}

	for sc.Scan() {
		inputs := strings.Split(sc.Text(), " ")

		switch inputs[0] {
		case "insert":
			dict[inputs[1]] = true
		case "find":
			if dict[inputs[1]] {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		}
	}
}
