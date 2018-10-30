package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	var linkedList = make([]string, 1000000)
	var length int

	for sc.Scan() {
		inputs := strings.Split(sc.Text(), " ")

		switch inputs[0] {
		case "insert":
			linkedList[length] = inputs[1]
			length++
		case "delete":
			for i := length - 1; i >= 0; i-- {
				if linkedList[i] == inputs[1] {
					linkedList = append(linkedList[:i], linkedList[i+1:]...)
					length--
					break
				}
			}
		case "deleteFirst":
			linkedList[length-1] = ""
			length--
		case "deleteLast":
			linkedList = linkedList[1:]
			length--
		}
	}

	fmt.Printf("%s", linkedList[length-1])
	for i := length - 2; i >= 0; i-- {
		fmt.Printf(" %s", linkedList[i])
	}
	fmt.Println("")
}
