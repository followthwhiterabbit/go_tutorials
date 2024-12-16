package main 
import ("fmt")


func main() {

	/*
	prices := []int{10,20,30}


	fmt.Println(prices[0])
	fmt.Println(prices[2])
	*/ 

	/*
	Change elements of a slice


	*/


	/*
	prices := []int{10,20,30}
	prices[2] = 50
	fmt.Println(prices[0])
	fmt.Println(prices[2])
	*/ 


	/*
		Append elements to a slice
		
		slice_name = append(slice_name, element1, element2, ...)



	*/ 



	/*

	myslice1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

		
	myslice1 = append(myslice1, 20, 21)
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	*/


	/*
		Appending one slice to another slice 

		slice3 = append(slice1, slice2...) // ... ---> necessary to put after the last slice inside the function 


	*/ 


	/*
	
	myslice1 := []int{1, 2, 3}
	myslice2 := []int{4, 5, 6}
	myslice3 := append(myslice1, myslice2...)

	fmt.Printf("myslice3 = %v\n", myslice3)
	fmt.Printf("length=%d\n", len(myslice3))
	fmt.Printf("capacity=%d\n", cap(myslice3))


	*/


//-------------------------------------------------
/*
	Unlike arrays, it is possible to change the length of a slice

Index:     0    1    2    3    4    5
Array:     9   10   11   12   13   14
Slice:          10   11   12   13
Length:          <--- 4 elements --->
Capacity:        <-------- 5 elements --------->

*/


/*
arr1 := [6]int{9, 10, 11, 12, 13, 14} // an array
myslice1 := arr1[1:5] // slice an array
fmt.Printf("myslice1 = %v\n", myslice1)
fmt.Printf("length = %d\n", len(myslice1))
fmt.Printf("capacity = %d\n", cap(myslice1))


myslice1 = arr1[1:3] // change length by re-slicing the array 
fmt.Printf("myslice1 = %v\n", myslice1)
fmt.Printf("length = %d\n", len(myslice1))
fmt.Printf("capacity = %d\n", cap(myslice1))


myslice1 = append(myslice1, 20, 21, 22, 23) // change length by appending items 
fmt.Printf("myslice1 = %v\n", myslice1)
fmt.Printf("length = %d\n", len(myslice1))
fmt.Printf("capactity = %d\n", cap(myslice1))

*/


/*
	Memory efficiency 
		-- when using slices, go loads all the underlying elements into the memory 
		-- if the array is large and you only need a few elements, it is better to copy those elemnets using the copy() function 
		-- the copy function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the 
			program. 

		copy(dest, src)

		



*/


		numbers := []int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15}
		fmt.Printf("numbers = %v\n", numbers)
		fmt.Printf("length = %d\n", len(numbers))
		fmt.Printf("capacity = %d\n", cap(numbers))

		// create copy with only needed numbers /* the capacity of the new slice is now less than the capacity of the original slice */ 

		 neededNumbers := numbers[:len(numbers)-10]                          
		 numbersCopy := make([]int, len(neededNumbers))
		 copy(numbersCopy, neededNumbers)

		fmt.Printf("numbersCopy = %v\n", numbersCopy)
		fmt.Printf("length = %d\n", len(numbersCopy))
		fmt.Printf("capacity = %d\n", cap(numbersCopy))
		




		




}
