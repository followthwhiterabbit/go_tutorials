/*

1 - ) 

	declaration of arrays in go 
	var array_name = [length]datatype{values} // here length is defined 


	or 

	var array_name = [...]datatype{values} // here length is inferred 


2 - ) 

	with the := sign : 

	array_name := [length]datatype{values} // we dont put the var keyword 

	or 

	array_name := [...]datatype{values}   // here the length is inferred 





*/ 

package main 
import ("fmt")


func main() {
	//var arr1 = [3]int{1, 2, 3} // or var arr1 = [...]int{1, 2, 3}
	//arr2 := [5]int{4, 5, 6, 7, 8} // or arr2 := [...]int{4, 5, 6, 7, 8}

/*
	var arr1 = [...]int{1, 2, 3}
	arr2 := [...]int{4, 5, 6, 7, 8}

	fmt.Println(arr1)
	fmt.Println(arr2)
	
*/ 

/*

var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}



// fmt.Printf(cars)  // printf doesn't print these values , it takes a string with any format specifier to output 

fmt.Println(cars)
*/

/*
prices := [3]int{10, 20, 30}

//fmt.Println(prices[0])
//fmt.Println(prices[2])

prices[2] = 50 


fmt.Println(prices)
*/

/*

arr1 := [5]int{} // not initialized 
arr2 := [5]int{1, 2} // partially initialized 
arr3 := [5]int{1, 2, 3, 4, 5} // fully initialized 



fmt.Println(arr1)
fmt.Println(arr2)
fmt.Println(arr3)
*/ 

// initializing only specific elements 
//arr1 := [5]int{1:10, 2:40} // initializing only second and third items in the array 

//fmt.Println(arr1)

arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
arr2 := [...]int{1, 2, 3, 4, 5, 6} // length is inferred

fmt.Println(len(arr1))
fmt.Println(len(arr2))










}











