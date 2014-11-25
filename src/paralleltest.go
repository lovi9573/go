package main

import (
   	"fmt"
	"math/rand"
)


func main() {
	//Create a random array to be sorted.
	size := 10 //20000000  
	var ar = make([] directory, size)  
	for i := range ar {  
		v := rand.Intn(2000)
		ar[i] = directory{size:v , path:"Dummydir" }
	}
	fmt.Print("original:")
	fmt.Println(ar)

	pmergesort(ar)
	
	fmt.Print("sorted:") 
    	fmt.Println(ar)
}
