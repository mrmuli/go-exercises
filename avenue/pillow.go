package main

import "fmt"

type rectangle struct {
	length, width int

	geometry struct {
		area, perimeter int
	}
}

type square struct {
	length  float32
	breadth float32
	color   string
}

func Slam() {
	// fmt.Println(rectangle{105, 2510})
	fmt.Println("surprise mgf")
}

func Pear() {
	var sup rectangle
	sup.length = 20
	sup.width = 50

	sup.geometry.area = sup.length * sup.width
	sup.geometry.perimeter = 2 * (sup.length + sup.width)

	fmt.Println("Yoyo: ", sup)
	fmt.Printf("Area: %v\n", sup.geometry.area)
	fmt.Printf("Perimeter: %v\n", sup.geometry.perimeter)

}

func Sarge() {

	var sq = square{10.5, 20.3, "blue"}
	fmt.Println(sq)

	// skip values by assinging values to fileds of the struct
	var subb = square{length: 50.33, color: "matte"}
	fmt.Println(subb)

	wild := square{34.43, 22.32, "black"}
	fmt.Println(wild)
}

func functionName() int {
	fmt.Println("This one returns integers")
	return 42 // seems like this needs to be the last thing
}

func returnMulti() (int, string) {
	return 42, "wazzup" // can return more than one in order apparently
}

var x, str = returnMulti()

func returnMultiNamed() (n int, s string) {
	n = 42 // can't use := because n and s already exist in context with zero values
	s = "foo"
	return
}

var k, j = returnMultiNamed()

func sum(args ...int) int { // sample variadic function
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func Loopa() {
	count := 4
	for i := 0; i < count; i++ {
		fmt.Println("Hello")
	}
}

func Structz() {
	var x [5]int // arrays have size specified on declaration
	fmt.Println(x)
}
