package mains

import (
	"fmt"
	"math"
)

func plusPlus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

//Embedding Struct

type base struct {
	num int
}

type container struct {
	base
	str string
}

// Generics
type List[T any] struct {
	head, tail *element[T]
}

type element[T any] struct {
	next *element[T]
	val  T
}

// Use New for structures and make foe maps, slices and channels
// new returns a pointer and make returns a nin zerod value.

func main() {
	//Simple hello world
	fmt.Println("hello world")

	//Couple of prints
	fmt.Println("go" + "lang")
	//Integers and floats.
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)
	//Boolean
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	//Constant
	const s string = "constant"
	fmt.Println(s)
	const n = 500000000
	const d = 3e20 / n
	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))

	//For loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println("range", i)
	}

	//if/else
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	//Switch
	ni := 2
	fmt.Print("Write ", ni, " as ")
	switch ni {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	//Array
	var a [5]int
	fmt.Println("emp:", a)
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	fmt.Println("len:", len(a))
	// Declaracion de un mapa
	mapita := make(map[string]int)

	mapita["k1"] = 7
	mapita["k2"] = 12

	fmt.Println("map:", mapita)

	//One line comments
	/*
		Multiple line comments - same as C
	*/

	baseMessage := "Hello, world!"
	fmt.Println(baseMessage)

	con := container{
		base: base{
			num: 1,
		},
		str: "some name",
	}

	fmt.Printf("co={num: %v, str: %v}\n", con.num, con.str)

}
