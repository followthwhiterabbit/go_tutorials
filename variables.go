package main 
import "fmt"

/*
var a int 
var b int = 2
var c = 3
*/

//a := 1 





func main() {

	/*
	var student1 string = "John" // type is string 
	var student2 = "Jane" // type is inferred 

	x := 2 // type is inferred too 


	fmt.Println(student1)
	fmt.Println(student2)
	
	fmt.Println(x)
	*/ 

/*

	var a string
	var b int 
	var c bool

	fmt.Println(a)  // default value is ""
	fmt.Println(b) // default value is 0 
	fmt.Println(c) // default value is false 
*/ 

/*
var student1 string 
student1 = "John"


fmt.Println(student1)
*/

/*
a = 1
fmt.Println(a)
fmt.Println(b)
fmt.Println(c)
*/


// fmt.Println(a)


/*
var a, b, c, d int = 1, 3, 5, 7

fmt.Println(a)
fmt.Println(b)
fmt.Println(c)
fmt.Println(d)
*/


/*
var a, b = 6, "Hello" // here the 
c, d := 7, "World"  // types are inferred here 


fmt.Println(a)
fmt.Println(b)
fmt.Println(c)
fmt.Println(d)
*/

/*
// Multiple variable declaration 

var (
	a int 
	b int = 1
	c string = "hello"
)

fmt.Println(a)
fmt.Println(b)
fmt.Println(c)

*/


// untyped constant 
const PI = 3.14 // the type of the variable is inferred 

const A int = 1 // type of the variable is specified 

const B = 5 // type of the variable is inferred 


// B = 2; // you cannot make this , redeclaration is not possible once a const variable is declared 


// multiple constants 

const (
	C int = 1  // typed constant 
	D = 3.14 // untyped constant 
	E = "Hi!" // untyped constant 
)

fmt.Println(C)
fmt.Println(D)
fmt.Println(E)

}


