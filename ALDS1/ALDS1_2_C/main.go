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
	length, _ := strconv.Atoi(sc.Text())
	var inputSuits = make([]string, length)
	var inputNums = make([]int, length)
	var inputSuits2 = make([]string, length)
	var inputNums2 = make([]int, length)

	sc.Scan()
	for i, v := range strings.Split(sc.Text(), " ") {
		inputSuits[i] = string(v[0])
		inputSuits2[i] = string(v[0])
		inputNums[i], _ = strconv.Atoi(string(v[1]))
		inputNums2[i], _ = strconv.Atoi(string(v[1]))
	}

	bubbleSuits, bubbleNums := bubbleSort(inputSuits, inputNums)
	selectionSuits, selectionNums := selectionSort(inputSuits2, inputNums2)

	print(inputSuits, inputNums, bubbleSuits, bubbleNums)
	print(inputSuits, inputNums, selectionSuits, selectionNums)
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

func print(orgStrs []string, orgNums []int, sortedStrs []string, sortedNums []int) {
	for i, v := range sortedNums {
		if i == 0 {
			fmt.Printf("%s%d", sortedStrs[i], v)
		} else {
			fmt.Printf(" %s%d", sortedStrs[i], v)
		}
	}
	if judgeStable(orgStrs, orgNums, sortedStrs, sortedNums) {
		fmt.Println("\nStable")
	} else {
		fmt.Println("\nNot stable")
	}
}
