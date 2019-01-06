package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	value int
	left  int
	right int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	length, _ := strconv.Atoi(sc.Text())
	tree := make([]node, length)

	c := 0
	for sc.Scan() {
		inputs := strings.Split(sc.Text(), " ")

		switch inputs[0] {
		case "insert":
			num, _ := strconv.Atoi(inputs[1])
			tree[c] = node{value: num, left: -1, right: -1}
			if c != 0 {
				insert(tree, 0, c)
			}
			c++
		case "find":
			num, _ := strconv.Atoi(inputs[1])
			result, _ := find(tree, num, 0, -1)
			if result >= 0 {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		case "delete":
			num, _ := strconv.Atoi(inputs[1])
			point, parent := find(tree, num, 0, -1)
			if point < 0 {
				continue
			}

			if tree[point].left == -1 && tree[point].right == -1 {
				if tree[parent].left == current {
					tree[parent].left = -1
				else if tree[parent].right == current {
					tree[parent].right = -1
				}
			} else if tree[point].left != -1 && tree[point].right == -1 {
				if tree[parent].left == current {
					tree[parent].left = tree[point].left
				else if tree[parent].right == current {
					tree[parent].right = tree[point].left
				}
			} else if tree[point].left == -1 && tree[point].right != -1 {
				if tree[parent].left == current {
					tree[parent].left = tree[point].right
				else if tree[parent].right == current {
					tree[parent].right = tree[point].right
				}
			} else {
				next := findNextInOrder(tree, point)
				tree[point].value = tree[next].value
			}

		case "print":
			printInorder(tree, 0)
			fmt.Println()
			printPreorder(tree, 0)
			fmt.Println()
		}
	}
}

func insert(tree []node, diffPoint int, nodePoint int) {
	if tree[diffPoint].value <= tree[nodePoint].value {
		if tree[diffPoint].right == -1 {
			tree[diffPoint].right = nodePoint
		} else {
			insert(tree, tree[diffPoint].right, nodePoint)
		}
	} else {
		if tree[diffPoint].left == -1 {
			tree[diffPoint].left = nodePoint
		} else {
			insert(tree, tree[diffPoint].left, nodePoint)
		}
	}
}

func find(tree []node, value int, current int, parent int) (int, int) {
	if tree[current].value == value {
		return current, parent
	} else if tree[current].value > value {
		if tree[current].left == -1 {
			return -1, -1
		}
		return find(tree, value, tree[current].left, current)
	} else if tree[current].value < value {
		if tree[current].right == -1 {
			return -1, -1
		}
		return find(tree, value, tree[current].right, current)
	}
	return -2, -2
}

func findNextInOrder(tree []node, current int) int {
	if tree[current].right == -1 {
		return current
	}
	findNextInOrder(tree, tree[current].right)
}

func printInorder(tree []node, current int) {
	if current == -1 {
		return
	}
	printInorder(tree, tree[current].left)
	fmt.Printf(" %d", tree[current].value)
	printInorder(tree, tree[current].right)
}

func printPreorder(tree []node, current int) {
	if current == -1 {
		return
	}
	fmt.Printf(" %d", tree[current].value)
	printPreorder(tree, tree[current].left)
	printPreorder(tree, tree[current].right)
}
