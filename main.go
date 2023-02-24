package main

import (
	"fmt"
	"math/rand"
	"time"
)

var rnd *rand.Rand

func main() {

	seed := time.Now().UnixNano()
	rnd = rand.New(rand.NewSource(seed))

	gg := rnd.Uint32()
	fmt.Println("My favorite number is", gg)

	//rand.Seed(time.Now().UnixNano())

	weight := gg & 0x0000ffff
	ww := (float64(weight) - 32768.0) / 8192.0
	source := gg & 0xff000000 >> 24
	sink := gg & 0x00ff0000 >> 16
	so := (source & 0x80 >> 7) % 2
	newsource := (source & 0x7f) % 8
	newsink := (sink & 0x7f) % 8
	si := (sink & 0x80 >> 7) % 2

	fmt.Println("weight", weight)
	fmt.Println("ww", ww)
	fmt.Println("source", source)
	fmt.Println("sink", sink)
	fmt.Println("so", so)
	fmt.Println("newsource", newsource)
	fmt.Println("newsink", newsink)
	fmt.Println("si", si)

}
