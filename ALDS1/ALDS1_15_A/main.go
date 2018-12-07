package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())

	count := 0
	coins := []int{25, 10, 5, 1}

	for _, v := range coins {
		countTmp := num / v
		count = count + countTmp
		num = num - countTmp*v

		if num == 0 {
			break
		}
	}

	fmt.Println(count)

}
