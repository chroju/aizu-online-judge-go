package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type Coordinate struct {
	x float64
	y float64
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	num, _ := strconv.Atoi(sc.Text())
	begin := Coordinate{0, 0}
	end := Coordinate{100, 0}

	fmt.Println(begin.x, begin.y)
	kock(begin, end, num)
	fmt.Println(end.x, end.y)
}

func kock(begin Coordinate, end Coordinate, n int) {
	if n == 0 {
		return
	}

	var s, t, u Coordinate
	s.x = (begin.x*2 + end.x) / 3
	s.y = (begin.y*2 + end.y) / 3
	t.x = s.x*2 - begin.x
	t.y = s.y*2 - begin.y
	u.x = (t.x-s.x)/2 - (t.y-s.y)*math.Sqrt(3)/2 + s.x
	u.y = (t.x-s.x)*math.Sqrt(3)/2 - (t.y-s.y)/2 + s.y

	kock(begin, s, n-1)
	fmt.Println(s.x, s.y)
	kock(s, u, n-1)
	fmt.Println(u.x, u.y)
	kock(u, t, n-1)
	fmt.Println(t.x, t.y)
	kock(t, end, n-1)
}
