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
