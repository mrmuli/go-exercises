package main

import (
	"fmt"
)

var jamo int = 42
var stingo string = "Jamo"
var zii bool = false

type potato int

var shira potato

var ro int

func main() {

	fmt.Println("#####")

	s := fmt.Sprintf("Mi naitwa %v, %v years old.", stingo, jamo)
	fmt.Println(s)

	fmt.Println(shira)
	fmt.Printf("%T\n", shira)
	shira = 42
	fmt.Printf("Shira is now %v.\n", shira)

	// type conversion from potato to underlying int
	ro = int(shira)

	fmt.Printf("ro now has %v\n", ro)
	fmt.Printf("%T\n", ro)

	fmt.Println("#####")
	Decclare()

	ConvertType()

	BitShifting()

	PrintTypes()

	AssignShift()

}

/*
Decclare is a way to declare Values and assign identifiers
*/
func Decclare() {
	x, y, z := 42, "James Bond", true

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
}

/*
StandardDeclare is a more typical way of assigning identifiers
*/
func StandardDeclare() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", y)
	fmt.Printf("%v\n", z)
}

/*
DeclareOnSingle just prints on a single line...
*/
func DeclareOnSingle() {
	bat, sup, indie := "star", 42, false

	fmt.Println(bat, sup, indie)
}

func ConvertType() {
	s := "Hello world"
	fmt.Printf("s started as a %T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Printf("%b\n", bs)
	fmt.Printf("%x\n", bs)

	fmt.Printf("s is now bs as a %T", bs)
}

func BitShifting() {

	bay := 4
	fmt.Printf("%d \t %b\n", bay, bay)

	you := bay << 2
	fmt.Printf("%d\t%b\n", you, you)
}

func PrintTypes() {
	stuff := 1024

	fmt.Printf("binary: %b\n", stuff)
	fmt.Printf("hex: %#x\n", stuff)
	fmt.Printf("decimal: %d\n", stuff)
}

/*
AssignShift assings an int to a var
Prints the int as a binary, hex and decimal val
Shifts the bits of that int over 1 left position
Assigns this to a var and prints as bin, hex and decimal
*/
func AssignShift() {
	tint := 1024
	fmt.Printf("%b\t%#x\t%d\n", tint, tint, tint)
	glen := tint << 1
	fmt.Printf("%b\t%#x\t%d\n", glen, glen, glen)
}

/*
Reminds you about the 10 types of people...

42 = 32 (1) 16 (0) 8 (1) 4 (0) 2 (1)

68 = 64 (1) 32 (0) 16 (0) 8 (0) 4 (1) 2 (0) 1 (0)

102 = 64 (1)  32 (1) 16 (0) 8 (0) 4 (1) 2 (1) 1 (0)

53 = 64 (0) 32 (1) 16 (1) 8 (0) 4 (1) 2 (0) 1 (1)

Binary scale (1,2,4,8,16,32,64,128,256...) {2 to power n}
*/
