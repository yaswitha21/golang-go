package main

import "fmt"

func main() {
	var flag bool = true
	var str string = "some text"
	var i int = 1234
	var ui uint = 1234
	var r rune = 1234 // rune = int32
	var b byte = 255  //	 byte = uint8(value from 0 to 255)
	var f float32 = 456.67
	fmt.Println(flag, str, i, ui, b, f, r)
}
