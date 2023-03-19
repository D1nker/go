package main

import (
	"fmt"
	"math"
	"math/rand"
	"reflect"
)

func main() {
	str := "Hello Q"
	float := float64(100)
	array := [3]int{1, 2, 3}

	fmt.Println(str)
	fmt.Println(reflect.TypeOf(str))

	fmt.Println(float)
	fmt.Println(reflect.TypeOf(float))

	fmt.Println(array)
	fmt.Println(reflect.TypeOf(array))

	for i := 1; i < 3; i++ {
		fmt.Println(array[i])
		array[i] += i
	}

	fmt.Println("My fav number is", rand.Intn(100))
	fmt.Println(math.Pi)

	fmt.Println(add(5, 5))
	fmt.Println(shortenAdd(5, 5))

	fmt.Println(multipleReturn("Hello", "World"))

	a, b := multipleReturn("var1", "var2")
	fmt.Println(a, b)

	fmt.Println(split(100))

	fmt.Println(quentin, alex, joseph)

	getType(quentin)
	getType(array)

	/* same but faster */
	/* var x int = 42 */
	/* Go is statically-typed so we cant re assign to a new type */
	x := 42
	x = 500
	fmt.Println(x)
}

func add(a int, b int) int {
	return a + b
}

func shortenAdd(a, b int) int {
	return a + b
}

func multipleReturn(a, b string) (string, string) {
	return a, b
}

func split(sum int) (x, y int) {
	x = sum * 100
	y = sum / 2
	return
}

const quentin, alex, joseph = "quentin", "alex", "joseph"

// the var declaration `:=` can only be used in a function scope

func getType(someVar interface{}) interface{} {
	fmt.Printf("value: (%v), type: (%T)\n", someVar, someVar)
	return someVar
}

const big = 1 << 100
const small = big >> 99

const (
	biggie  = 1 << 100
	smallie = biggie >> 99
)
