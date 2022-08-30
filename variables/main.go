package main

import "fmt"

var (
	t = 100
	a = 50
)

func main() {

	//var name string = "text"
	var foo int = 42
	var foo1 int = 42
	var bar int
	var i, j int = 23, 44
	var x, y int

	name := "text"

	name = "textbook"

	foo = 46 // changing the variable value and observe the output

	fmt.Println(name, foo, foo1, bar, i, j, x, y)
	fmt.Println("main()", a)
	a = 12
	fmt.Println("main()", a)

	printSomething()
}

func printSomething() {

	a := 12

	fmt.Println("printSomething()", a)

	//fmt.Println(name, foo, foo1, bar, i, j, x, y)// this will give error as foo,etc are declared in main scope
	//and these cant be visible outside main scope

	// whenever we create a variable the var scope will be inside that func only.

}
