package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var count int

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	length, _ := strconv.Atoi(sc.Text())
	A := make([]int, length)
	i := 0
	for sc.Scan() {
		A[i], _ = strconv.Atoi(sc.Text())
		i++
	}

	mergeSort(A, 0, length)
	fmt.Println(strings.Trim(fmt.Sprint(A), "[]"))
	fmt.Println(count)
}

func merge(A []int, left int, mid int, right int) {
	n1 := mid - left
	n2 := right - mid

	L := make([]int, n1+1)
	R := make([]int, n2+1)
	for i := 0; i <= n1-1; i++ {
		L[i] = A[left+i]
	}
	for i := 0; i <= n2-1; i++ {
		R[i] = A[mid+i]
	}
	L[n1] = int(math.Pow(10, 9)) + 1
	R[n2] = int(math.Pow(10, 9)) + 1

	i := 0
	j := 0

	for k := left; k <= right-1; k++ {
		count++
		if L[i] <= R[j] {
			A[k] = L[i]
			i++
		} else {
			A[k] = R[j]
			j++
		}
	}

}

func mergeSort(A []int, left int, right int) {
	if left+1 < right {
		mid := (left + right) / 2
		mergeSort(A, left, mid)
		mergeSort(A, mid, right)
		merge(A, left, mid, right)
	}
}
