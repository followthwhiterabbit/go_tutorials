/* - syntax of the slices - 
	slice_name := []datatype{values}

	-- common way to declare a slice --- myslice := []int{} // 0 length and 0 capacity slice 


	myslice := []int{1, 2, 3}
		- len function returns the length of the slice (the number of elements in the slice)
		- cap() function returns the capacity of the slice (the number of elements the slice can grow or shrink to)






*/


package main 
import ("fmt")


func main(){
/*
	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)
*/ 

// create a slice from an array 

/*
arr1 := [6]int{10, 11, 12, 13, 14, 15}
myslice := arr1[2:6] // slicing an array  , going more than 6 gives out of bounds error 
					// if myslice started from element 0, the slice capacity would be 6

fmt.Printf("myslice = %v\n", myslice)
fmt.Printf("length  = %v\n", len(myslice)) // length returns 2, 12 & 13 
fmt.Printf("capacity = %v\n", cap(myslice)) // capacity returns 4 because this slice can grow until the end of the array

*/


// create a slice with the make function 
/*
slice_name := make([]type, length, capacity) 
	- if the capacity is not defined, it will be equal to length.	

*/

myslice1 := make([]int, 5, 10) 
fmt.Printf("myslice1 = %v\n", myslice1)
fmt.Printf("length = %d\n", len(myslice1))
fmt.Printf("capacity = %d\n", cap(myslice1))

// with omitted capacity 
myslice2 := make([]int, 5) // with omitted capacity
fmt.Printf("myslice2 = %v\n", myslice2)
fmt.Printf("length = %d\n", len(myslice2))
fmt.Printf("capacity = %d\n", cap(myslice2))

}



