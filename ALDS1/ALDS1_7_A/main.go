package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	id       int
	parent   int
	depth    int
	nodeType string
	children []int
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 1000)
	sc.Buffer(buf, 10000000)

	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())
	tree := make([]node, num)

	for sc.Scan() {
		rawInputs := strings.Split(sc.Text(), " ")
		inputs := make([]int, len(rawInputs))
		for j, v := range rawInputs {
			inputs[j], _ = strconv.Atoi(v)
		}

		id := inputs[0]
		if inputs[1] != 0 {
			tree[id].children = inputs[2:]
			tree[id].nodeType = "internal node"
		} else {
			tree[id].nodeType = "leaf"
		}
	}

	for i, v := range tree {
		if v.nodeType == "internal node" {
			registerParent(tree, i, v.children, 0)
		}
	}

	for i, v := range tree {
		children := strings.Join(strings.Split(strings.Trim(fmt.Sprint(v.children), "[]"), " "), ", ")
		if v.nodeType == "internal node" && v.depth == 0 || v.depth == 0 && len(v.children) == 0 {
			v.nodeType = "root"
			v.parent = -1
		}
		fmt.Printf("node %d: parent = %d, depth = %d, %s, [%s]\n", i, v.parent, v.depth, v.nodeType, children)
	}
}

func registerParent(tree []node, id int, children []int, depth int) {
	for _, v := range children {
		if tree[v].depth != 0 && tree[v].depth > depth+1 {
			continue
		}
		tree[v].parent = id
		tree[v].depth = depth + 1
		if len(tree[v].children) > 0 {
			registerParent(tree, v, tree[v].children, depth+1)
		}
	}
}
