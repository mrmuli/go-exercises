package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

const Hello = "hello"
const world = "world"

const Summer, time = "summer", "time"

const (
	Blue   = "blue"
	lagoon = "lagoon"
)

func main() {
	fmt.Println("Hello world!")
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))

	fmt.Println(Summer, time)
	fmt.Println(Blue, lagoon)
	fmt.Printf("%v %v\n", Hello, world)
}
