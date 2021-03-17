package main

import (
	"fmt"
)

const Hello = "heya"

var world string = "world"

var e error

type hotdog int

var hot hotdog

func main() {

	s := "jamaa"
	// *T 			pviointer type
	// [5]Tee 		array type
	// []Tesa 		slice type
	// map[Tkey]T	map type

	// basic types
	type (
		MyInt int
		Age   int
		Text  string
	)

	// composite types
	type IntPtr *int
	type Book struct {
		author, title  string
		pages, columns int
	}
	type Convert func(in0 int, in1 bool) (out0 int, out1 string)
	type StringArray [5]string
	type StringSlice []string

	fmt.Println("Heya")
	fmt.Printf("%v\n", world)
	fmt.Println(s)

	fmt.Println(15 == 017)
	fmt.Println(15 == 0xF)
	Slam()
	Pear()
	Sarge()
	Loopa()

	fmt.Printf("%T", hot)
	fmt.Printf("%v", hot)
}

func f() {
	// names of the 3 defined types
	// can only be used within this func.
	type PersonAge map[string]int
	type MessageQueue chan string
	type Reader interface{ Read([]byte) int }
}
