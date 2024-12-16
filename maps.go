package main 
import ("fmt")

/*
maps are used to store data values in key:value pairs.

each element in a map is a key value pair.

A map is an unordered and changeable collection that does not allow duplicates

The length of a map is the number of its elements. You can find it using the len() function.

The default value of a map is nil.

Maps hold references to an underlying hash table

Go has multiple ways for creating maps

	
Creating maps using var and := 

var a = map[KeyType]ValueType{key1:value1, key2:value2,...)
b := map[KeyType]ValueType{key1:value1, key2:value2,...)


*/
func main(){

	/*
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year": "1964"} // mapping from string to string 
	b := map[string]int{"Oslo": 1, "Bergen": 2, "Trondheim": 3, "Stavanger": 4}

	fmt.Printf("a\t%v\n",a) // \t ---> tab 
	fmt.Printf("b\t%v\n",b)
	*/

	// creating maps with make() function 

	/*
	var a = make(map[KeyType]ValueType)
	b := make(map[KeyType]ValueType)

	*/


	/*
	var a = make(map[string]string) // the map is empty now 

	a["brand"] = "Ford"
	a["model"] = "Mustang"
	a["year"] = "1964"
										// a is no loger empty  
	b := make(map[string]int)
	b["Oslo"] = 1
	b["Bergen"] = 2
	b["Trondheim"] = 3
	b["Stavanger"] = 4
	
	fmt.Printf("a\t%v\n", a)
	fmt.Printf("b\t%v\n", b)
	*/

	// create an empty map 
	//var a map[KeyType][ValueType]

/*
Note: The make()function is the right way to create an empty map.
If you make an empty map in a different way and write to it, it will causes a runtime panic.

*/

/*
var a = make(map[string]string)
var b map[string]string

fmt.Println(a == nil)
fmt.Println(b == nil)
*/


/*
	// valid key types: - booleans 
			- numbers 
			- strings
			- arrays
			- pointers
			- structs 
			- interfaces ( as long as the dynamic type supports equality)

	// invalid key types:  // these types are invalid because the equality operetor (==) is not defined for them .
			- slices
			- maps
			- functions 
	
*/

/*
	Allowed Value Types
		- The map values can be any type 

*/

// Access map elements 

/*
	value = map_name[key]



*/
/*
var a = make(map[string]string) 
a["brand"] = "Ford"
a["model"] = "Mustang"
a["year"] = "1964"

fmt.Printf(a["brand"])
*/


/*
	Update & Add new map elements 
	map_name[key] = value 


*/

/*

var a = make(map[string]string)
a["brand"] = "Ford"
a["model"] = "Mustang"
a["year"] = "1964"

fmt.Println(a)
// updating an element
a["year"] = "1970" 
a["color"] = "red" // adding an element

fmt.Println(a)
*/

// remove element from map

// removing elements is done using the delete() function. 

// delete(map_name, key)

/*
var a = make(map[string]string)

a["brand"] = "Ford"
a["model"] = "Mustang"
a["year"] = "1964"


fmt.Println(a)

delete(a, "year")

fmt.Println(a)
*/

/*
	Check for specific elements in a map
		- you can check if a certain key exists in a map 
	
		val, ok := map_name[key]
		
	- If you only want to check the existence of a certain key, you can use the blank identifier(_) in place of val. 


*/

/*
	var a = map[string]string{"brand": "Ford", "model": "Mustang", "year":"1964", "day":""}
	
	val1, ok1 := a["brand"] // checking for existing key and its value
	val2, ok2 := a["color"] // checking for non-existing key and its value
	val3, ok3 := a["day"] // checking for existing key and its value
	
	_, ok4 := a["model"] // only checking for existing key and not its value 
	
	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
	fmt.Println(val3, ok3)
	fmt.Println(ok4)

*/

/*
	Maps are References
		- Maps are references to hash tables.
		- If two map variables refer to the same hash table, changing the content of one variable affect the content of the other.

	

*/
/*
	var a = map[string]string{"brand":"Ford", "model": "Mustang", "year":"1964"}
	b := a

	fmt.Println(a)
	fmt.Println(b)

	b["year"] = "1970"
	fmt.Println("After change to b:")

	fmt.Println(a)
	fmt.Println(b)
*/

/*
	Iterate over maps 
		- you can use range to iterate over maps 

*/

/*
a := map[string]int{"one":1, "two":2, "three":3, "four":4} // unordered of course, because it's a map 

for k, v := range a {
	fmt.Printf("%v : %v, ", k, v)	
}
*/

/*
	Iterate over maps in a specific order
		- Maps are unordered data structures. If you need to iterate over a map in a specific order, you must have a seperate data structure that 
		specifies that order.


*/


/*
a := map[string]int{"one":1, "two":2, "three":3, "four":4}

var b []string // defining the order 

b = append(b, "one", "two", "three", "four")


for k, v := range a{ // loop with no order
	fmt.Printf("%v : %v, ", k, v) 
	
}

fmt.Println()

for _, element := range b{ // loop with the defined order 
	fmt.Printf("%v : %v, ", element, a[element])
}
*/







}
