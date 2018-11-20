package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	s := map[int]bool{}
	sum := 0
	for sc.Scan() {
		inputs := strings.Split(sc.Text(), " ")
		value, _ := strconv.Atoi(inputs[1])

		if inputs[0] == "0" {
			if !s[value] {
				s[value] = true
				sum++
			}
			fmt.Printf("%d\n", sum)
		} else if inputs[0] == "1" {
			if s[value] {
				fmt.Println(1)
			} else {
				fmt.Println(0)
			}
		} else if inputs[0] == "2" {
			if s[value] {
				s[value] = false
				sum--
			}
		} else if inputs[0] == "3" {
			end, _ := strconv.Atoi(inputs[2])
			var nums []int
			for key := range s {
				if s[key] {
					nums = append(nums, key)
				}
			}
			sort.Ints(nums)
			for _, v := range nums {
				if v > end {
					break
				} else if v >= value {
					fmt.Println(v)
				}
			}
		}
	}
}
