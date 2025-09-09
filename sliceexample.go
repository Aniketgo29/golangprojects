package main 

import "fmt"

func main () {

	myslice := []int{}

	fmt.Println(len(myslice))
	fmt.Println(cap(myslice))
	fmt.Println(myslice)

	myslice2 := []string{"go","slices","are","powerfull"}

	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
	fmt.Println(myslice)

}