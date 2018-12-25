package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

	mergeSort(A, 0, len(A)-1)
	fmt.Println(A)
}

func merge(A []int, left int, mid int, right int) {
	n1 := mid - left
	n2 := right - mid

	L := []int{}
	R := []int{}
	for i := 0; i <= n1-1; i++ {
		L[i] = A[left+i]
	}
	for i := 0; i <= n2-1; i++ {
		R[i] = A[mid+i]
	}

	i := 0
	j := 0

	for k := left; k <= right-1; k++ {
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
