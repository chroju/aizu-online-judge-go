package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	sc.Scan()
	var inputSuits []string
	var inputNums []int
	for _, v := range strings.Split(sc.Text(), " ") {
		inputSuits = append(inputSuits, string(v[0]))
		tmp, _ := strconv.Atoi(string(v[1]))
		inputNums = append(inputNums, tmp)
	}

	sortedSuits, sortedNums := bubbleSort(inputSuits, inputNums)
	fmt.Println(sortedSuits)
	fmt.Println(sortedNums)
	fmt.Println(judgeStable(inputSuits, inputNums, sortedSuits, sortedNums))

	sortedSuits, sortedNums = selectionSort(inputSuits, inputNums)
	fmt.Println(sortedSuits)
	fmt.Println(sortedNums)
	fmt.Println(judgeStable(inputSuits, inputNums, sortedSuits, sortedNums))
}

func bubbleSort(strs []string, nums []int) ([]string, []int) {
	var flag bool
	var count int
	for {
		if flag {
			break
		}
		flag = true
		for i := range nums[:len(nums)-1] {
			if nums[i] > nums[i+1] {
				tmp := nums[i]
				nums[i] = nums[i+1]
				nums[i+1] = tmp

				tmp2 := strs[i]
				strs[i] = strs[i+1]
				strs[i+1] = tmp2

				flag = false
				count++
			}
		}
	}
	return strs, nums
}

func selectionSort(strs []string, nums []int) ([]string, []int) {
	var count int
	for i := range nums {
		minj := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[minj] {
				minj = j
			}
		}
		if i != minj {
			tmp := nums[i]
			nums[i] = nums[minj]
			nums[minj] = tmp

			tmp2 := strs[i]
			strs[i] = strs[minj]
			strs[minj] = tmp2

			count++
		}
	}

	return strs, nums
}

func judgeStable(orgStrs []string, orgNums []int, sortedStrs []string, sortedNums []int) bool {
	var sortedSuits [9][]string
	for i, v := range sortedNums {
		sortedSuits[v-1] = append(sortedSuits[v-1], sortedStrs[i])
	}

	var orgSuits [9][]string
	for i, v := range orgNums {
		orgSuits[v-1] = append(orgSuits[v-1], orgStrs[i])
	}

	return reflect.DeepEqual(sortedSuits, orgSuits)
}
