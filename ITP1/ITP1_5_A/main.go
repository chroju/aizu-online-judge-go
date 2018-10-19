package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func rectangle(height int, width int) {
	for i := 0; i < height; i++ {
		fmt.Println(strings.Repeat("#", width))
	}
}

func main() {
	for sc.Scan() {
		inputs := strings.Split(sc.Text(), " ")
		if inputs[0] == "0" && inputs[1] == "0" {
			break
		}
		h, _ := strconv.Atoi(inputs[0])
		w, _ := strconv.Atoi(inputs[1])
		rectangle(h, w)
		fmt.Println("")
	}

}
