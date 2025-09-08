package main

import "fmt"

func main () {

	a := [2] string {"aniket", "Jyoti"}
	b := [2] string {"Male", "female"}
	c := [2] int {30, 28}

	fmt.Println(a[0] + b[0] )//+ c[0])
	fmt.Println(a[1] + b[1] )//+ c[1])
	fmt.Println(c[0])
}