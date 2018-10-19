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
	buf := make([]byte, 1000)
	sc.Buffer(buf, 100000000)

	sc.Scan()
	sc.Scan()
	listA := strings.Split(sc.Text(), " ")

	sc.Scan()
	sc.Scan()
	listB := strings.Split(sc.Text(), " ")

	i := 0
	j := 0
	for {
		a, _ := strconv.Atoi(listA[i])
		b, _ := strconv.Atoi(listB[j])
		if a == b {
			fmt.Println(listA[i])
			i++
			j++
		} else if a < b {
			fmt.Println(listA[i])
			i++
		} else {
			fmt.Println(listB[j])
			j++
		}

		if i == len(listA) {
			for k := j; k < len(listB); k++ {
				fmt.Println(listB[k])
			}
			break
		} else if j == len(listB) {
			for k := i; k < len(listA); k++ {
				fmt.Println(listA[k])
			}
			break
		}
	}

}
