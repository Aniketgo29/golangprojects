package main 

import "fmt"

func main () {
    
	var a [3] int = [3] int {3,4,5}
	fmt.Println(a[0] * a[1])
	fmt.Println(a[0] + a[2])
	fmt.Println(a[2] - a[0])
	fmt.Println(a[1] + a[1])
}