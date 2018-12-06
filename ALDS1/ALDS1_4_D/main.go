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
	inputs := strings.Split(sc.Text(), " ")
	luggageNum, _ := strconv.Atoi(inputs[0])
	trackNum, _ := strconv.Atoi(inputs[1])

	luggages := make([]int, luggageNum)
	i := 0
	sum := 0
	max := 0
	for sc.Scan() {
		luggages[i], _ = strconv.Atoi(sc.Text())

		sum = sum + luggages[i]
		if max < luggages[i] {
			max = luggages[i]
		}

		i++
	}

	tmp := sum / trackNum
	if tmp > max {
		max = tmp
	}
	for {
		if carry(luggages, max, trackNum) {
			break
		} else {
			max++
		}
	}

	fmt.Println(max)
}

func carry(luggages []int, max int, trackNum int) bool {
	result := 1
	tmp := max
	for _, v := range luggages {
		if tmp < v {
			result++
			tmp = max - v
		} else {
			tmp = tmp - v
		}

		if result > trackNum {
			return false
		}
	}

	return true
}
