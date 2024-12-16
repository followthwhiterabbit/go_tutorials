package main 
import ("fmt")
//import ("math")

func main(){

/*

	var a bool = true  // boolean 
	var b int = 5 // integer 
	var c float32 = 3.14 // floating point number 
	var d string = "Hi!" // string 



	fmt.Println("Boolean: ", a)
	fmt.Println("Integer: ", b)
	fmt.Println("float :", c)
	fmt.Println("String : ", d)

*/ 


// boolean data types 
/*
var b1  bool = true // type declaration with initial value 
var b2  = true // untyped declaration with initial value 
var b3 bool // typed declaration without initial value 
b4 := true // untyped declaration with initial value // type is inferred from the context 


fmt.Println(b1)
fmt.Println(b2)
fmt.Println(b3)
fmt.Println(b4)
*/

// signed integers 
/*
var x int = 500
var y int = -4500
fmt.Printf("Type: %T, value: %v\n", x, x)
fmt.Printf("Type: %T, value: %v", y, y)
*/

// unsigned int 
/*
var x uint = 500 
var y uint = 4500
fmt.Printf("Type: %T, value: %v\n", x, x)
fmt.Printf("Type: %T, value: %v", y, y)
*/


/*
var x int8 = 1000 

fmt.Printf("Type: %T, value: %v", x, x)

# command-line-arguments
./go_data_types.go:54:14: cannot use 1000 (untyped int constant) as int8 value in variable declaration (overflows)
*/

/*
	float data types are used store positive and negative numbers with a decimal point, like 35.3, -2.34, 3597.34987
	
	float32 --> 32 bit 
	float64 --> 64 bit 




*/ 

/*

var y float64 = 3.4e+38
x := math.Pow(2, 16)


if (x > y){
	fmt.Println("2 to the power 16 is greater than 3.4e+38\n")
} else {

	fmt.Println("2 to the power 16 is less than 3.4e+38\n")

}


	fmt.Println(math.Pow(2, 15))
*/

// string data type 
var txt1 string = "Hello!"
var txt2 string 
txt3 := "World 1"

fmt.Printf("Type: %T, value: %v\n", txt1, txt1)
fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
fmt.Printf("Type: %T, value: %v\n", txt3, txt3)


}
