package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type card struct {
	suit   string
	number int
}

type trump struct {
	cards []card
}

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	num, _ := strconv.Atoi(sc.Text())

	c := make([]card, num)
	t := trump{cards: c}
	org := []card{}
	i := 0
	for sc.Scan() {
		input := strings.Split(sc.Text(), " ")
		t.cards[i].suit = input[0]
		t.cards[i].number, _ = strconv.Atoi(input[1])
		n, _ := strconv.Atoi(input[1])
		org = append(org, card{suit: input[0], number: n})
		i++
	}

	t.quickSort(0, len(t.cards)-1)

	if judgeStable(org, t.cards) {
		fmt.Println("Stable")
	} else {
		fmt.Println("Not stable")
	}
	for _, v := range t.cards {
		fmt.Println(strings.Trim(fmt.Sprint(v), "{}"))
	}
}

func (t *trump) quickSort(p int, r int) {
	if p < r {
		q := t.partiton(p, r)
		t.quickSort(p, q-1)
		t.quickSort(q+1, r)
	}
}

func (t *trump) partiton(p int, r int) int {
	criterion := t.cards[r].number
	i := p - 1
	for j := p; j < r; j++ {
		if t.cards[j].number <= criterion {
			i++
			tmp := t.cards[i]
			t.cards[i] = t.cards[j]
			t.cards[j] = tmp
		}
	}
	tmp := t.cards[r]
	t.cards[r] = t.cards[i+1]
	t.cards[i+1] = tmp

	return i + 1
}

func judgeStable(org []card, sorted []card) bool {
	sortedSuits := map[int]string{}
	for _, v := range sorted {
		if _, ok := sortedSuits[v.number]; ok {
			sortedSuits[v.number] = sortedSuits[v.number] + v.suit
		} else {
			sortedSuits[v.number] = v.suit
		}
	}

	orgSuits := map[int]string{}
	for _, v := range org {
		if _, ok := orgSuits[v.number]; ok {
			orgSuits[v.number] = orgSuits[v.number] + v.suit
		} else {
			orgSuits[v.number] = v.suit
		}
	}

	return reflect.DeepEqual(sortedSuits, orgSuits)
}
