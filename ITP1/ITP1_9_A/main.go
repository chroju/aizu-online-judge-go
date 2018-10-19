package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)
var buf = make([]byte, 100000)

func main() {
	sc.Buffer(buf, 100000000)

	var searchTxt string
	var i, result int
	for sc.Scan() {
		input := strings.Split(sc.Text(), " ")
		if i == 0 {
			searchTxt = strings.ToLower(input[0])
		} else if input[0] == "END_OF_TEXT" {
			break
		} else {
			for _, word := range input {
				if strings.ToLower(word) == searchTxt {
					result++
				}
			}
		}
		i++
	}
	fmt.Println(result)
}
