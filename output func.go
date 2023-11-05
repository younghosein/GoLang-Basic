package main

import "fmt"

func main() {
	var l, j string = "fuck", "you"
	fmt.Print(l, "\n", j, "\n-----------\n")

	var a, s string = "rock", "metal"
	fmt.Println(a, s, "\n----------")

	var z, x int = 69, 85
	fmt.Printf("value is %v and type is %T\n", z, z)
	fmt.Printf("value is %v and type is %T\n------------\n", x, x)

	var i = 12
	fmt.Printf("%b\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%+d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%O\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%#x\n", i)
	fmt.Printf("%4d\n", i)
	fmt.Printf("%-4d\n", i)
	fmt.Printf("%04d\n----------------\n", i)
}
