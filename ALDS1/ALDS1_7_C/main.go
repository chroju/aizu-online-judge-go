package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	left  int
	right int
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
		tree[id].right, _ = strconv.Atoi(v[2])
	}

	fmt.Println(tree)
	fmt.Println("Preorder")
	preorder(tree, 0)
	fmt.Println("\nInorder")
	inorder(tree, 0)
	fmt.Println("\nPostorder")
	postorder(tree, 0)
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
	preorder(tree, tree[point].left)
	fmt.Printf(" %d", point)
	preorder(tree, tree[point].right)
}

func postorder(tree []node, point int) {
	if point == -1 {
		return
	}
	preorder(tree, tree[point].left)
	preorder(tree, tree[point].right)
	fmt.Printf(" %d", point)
}
