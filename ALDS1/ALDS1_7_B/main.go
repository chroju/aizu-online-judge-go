package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	parent   int
	depth    int
	degree   int
	height   int
	left     int
	right    int
	nodeType string
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	tree := make([]node, num)

	for sc.Scan() {
		rawInputs := strings.Split(sc.Text(), " ")

		id, _ := strconv.Atoi(rawInputs[0])
		tree[id].left, _ = strconv.Atoi(rawInputs[1])
		tree[id].right, _ = strconv.Atoi(rawInputs[2])
	}

	for i, v := range tree {
		if v.left != -1 || v.right != -1 {
			if v.left == -1 || v.right == -1 {
				tree[i].degree = 1
			} else {
				tree[i].degree = 2
			}
			if v.depth == 0 {
				getParent(tree, i, -1, 0)
			}
		} else if v.parent == 0 {
			tree[i].parent = -1
		}
	}

	for i, v := range tree {
		if v.parent == -1 {
			tree[i].nodeType = "root"
		} else if v.left == -1 && v.right == -1 {
			tree[i].nodeType = "leaf"
		} else {
			tree[i].nodeType = "internal node"
		}
	}

	for i, v := range tree {
		if v.nodeType == "leaf" {
			getHeight(tree, i, 0)
		}
	}

	for i, v := range tree {
		sibling := -1

		if (v.left != -1 || v.right != -1) && v.depth == 0 {
			v.parent = -1
			v.nodeType = "root"
		}

		if v.parent != -1 {
			if tree[v.parent].left != i {
				sibling = tree[v.parent].left
			} else {
				sibling = tree[v.parent].right
			}
		}

		fmt.Printf("node %d: parent = %d, sibling = %d, degree = %d, depth = %d, height = %d, %s\n", i, v.parent, sibling, v.degree, v.depth, v.height, v.nodeType)
	}
}

func getParent(tree []node, id int, parent int, depth int) {
	tree[id].parent = parent
	tree[id].depth = depth
	if tree[id].left != -1 {
		getParent(tree, tree[id].left, id, depth+1)
	}
	if tree[id].right != -1 {
		getParent(tree, tree[id].right, id, depth+1)
	}
}

func getHeight(tree []node, id int, height int) {
	if tree[id].height <= height {
		tree[id].height = height
		if tree[id].parent != -1 {
			getHeight(tree, tree[id].parent, height+1)
		}
	}
}
