package main

import (
	"fmt"
)

func main() {
	var myNames = [3]string{"Dexter", "Cyber", "Tafadzwa"}
	myFriends := [2]string{"Kudzai", "Philipa"}
	mySlice := []string{"Cyber"}
	mySlicedArray := myNames[1:3]
	slice_name := make([]int, 5, 20)
	var (
		a int
		b int    = 1
		c string = "Dexter"
	)
	const PI float32 = 3.14

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(PI)
	fmt.Println(myNames)
	fmt.Printf("The array %v is of type: %#v \n", myFriends, myFriends)
	fmt.Println(myFriends[1])
	fmt.Println(len(myNames))
	fmt.Println(len(mySlice), "and the values are ", mySlice)
	fmt.Println(mySlicedArray)
	fmt.Println(slice_name)

}
