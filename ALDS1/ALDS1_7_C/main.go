package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	isNotRoot bool
	left      int
	right     int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	length, _ := strconv.Atoi(sc.Text())
	tree := make([]node, length)

	for sc.Scan() {
		v := strings.Split(sc.Text(), " ")
		id, _ := strconv.Atoi(v[0])

		tree[id].left, _ = strconv.Atoi(v[1])
		if tree[id].left >= 0 {
			tree[tree[id].left].isNotRoot = true
		}

		tree[id].right, _ = strconv.Atoi(v[2])
		if tree[id].right >= 0 {
			tree[tree[id].right].isNotRoot = true
		}
	}

	var root int
	for i, v := range tree {
		if !v.isNotRoot {
			root = i
			break
		}
	}

	fmt.Println("Preorder")
	preorder(tree, root)
	fmt.Println("\nInorder")
	inorder(tree, root)
	fmt.Println("\nPostorder")
	postorder(tree, root)
	fmt.Println("")
}

func preorder(tree []node, point int) {
	if point == -1 {
		return
	}
	fmt.Printf(" %d", point)
	preorder(tree, tree[point].left)
	preorder(tree, tree[point].right)
}

func inorder(tree []node, point int) {
	if point == -1 {
		return
	}
	inorder(tree, tree[point].left)
	fmt.Printf(" %d", point)
	inorder(tree, tree[point].right)
}

func postorder(tree []node, point int) {
	if point == -1 {
		return
	}
	postorder(tree, tree[point].left)
	postorder(tree, tree[point].right)
	fmt.Printf(" %d", point)
}
