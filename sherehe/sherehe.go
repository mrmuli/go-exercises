package main

import "fmt"

/*
Zusha simply manipulates an array
*/
func Zusha() {
	noma := [5]int{}
	fmt.Println(noma)

	// A basic way to add values in there
	// manually this would be tedious
	// as noma[0-4] = digit as noma[0] = 23
	for i := 0; i <= 4; i++ {
		noma[i] = i
	}
	// fmt.Println(noma)
	// Range and Print content
	for k, v := range noma {
		fmt.Println(k, v)
	}

	// format printing: print type of the array
	fmt.Printf("Type is %T\n", noma)
}

// Kelele does the same as Zusha, only so much more noise :P
func Kelele() {
	kelele := make([]int, 10)

	for i := 0; i <= 9; i++ {
		kelele[i] = i
	}
	// fmt.Println(kelele)
	for i, j := range kelele {
		fmt.Println(i, j)
	}
	fmt.Printf("Type is %T", kelele)

}

func Kamati() {
	kamati := make([]int, 10)

	for i := 0; i <= 9; i++ {
		kamati[i] = i
	}
	// fmt.Println(kamati) [0 1 2 3 4 5 6 7 8 9]
	chizi := []int{10, 11, 12, 13, 14, 15, 16, 17, 18, 19}
	mkubwa := append(kamati, chizi...)
	// fmt.Println(mkubwa) [0 1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19]

	fmt.Printf("noma started here: %v\n", mkubwa[:6])
	fmt.Printf("Kelele got here: %v\n", mkubwa[5:10])
	fmt.Printf("Kamati brought the noise: %v\n", mkubwa[10:20])
}
