package main

import "fmt"


func main(){
	// this program contains a lot of funny stuff!!
	fmt.Println("Hello, Git!!!! ahahah")
	numberSlice := []int{1,2,3,4,5,6,7,8}
	for i,_ := range numberSlice{
		fmt.Printf("%d", i)
	}
}
