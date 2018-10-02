package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	cards := map[string]*[13]bool{
		"S": &[13]bool{},
		"H": &[13]bool{},
		"C": &[13]bool{},
		"D": &[13]bool{},
	}
	i := 0
	for sc.Scan() {
		if i != 0 {
			card := strings.Split(sc.Text(), " ")
			num, _ := strconv.Atoi(card[1])
			cards[card[0]][num-1] = true
		}
		i++
	}

	printCards("S", cards["S"])
	printCards("H", cards["H"])
	printCards("C", cards["C"])
	printCards("D", cards["D"])
}

func printCards(suit string, cards *[13]bool) {
	for i, v := range cards {
		if !v {
			fmt.Printf("%s %d\n", suit, i+1)
		}
	}
}
