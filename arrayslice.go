package main 

import "fmt"

func main () {

	array1 := [6]int{12,4,2,5,1,40}
	myslice := array1[2:5]

	fmt.Println(myslice)
	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
}