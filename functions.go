/*
	-functions 
	
	
Naming Rules for Go Functions

    A function name must start with a letter
    A function name can only contain alpha-numeric characters and underscores (A-z, 0-9, and _ )
    Function names are case-sensitive
    A function name cannot contain spaces
    If the function name consists of multiple words, techniques introduced for multi-word variable naming can be used

Tip: Give the function a name that reflects what the function does!

*/

package main 
import ("fmt")

/*

func myMessage(){
	fmt.Println("I just go executed")
}




func main(){
	myMessage()	
	myMessage()	
	myMessage()	
}

*/



/*
	Go Function Parameters and Arguments 


Information can be passed to functions as a parameter. Parameters act as variables inside the function.

Parameters and their types are specified after the function name, inside the parentheses. You can add as many parameters as you want, just separate them with a comma:


func FunctionName(param1 type, param2 type, param3 type){


}



*/


/*
func familyName(fname string){ // takes only one parameter
	fmt.Println("Hello", fname, "Refsnes")
}


func main(){
	familyName("Liam")
	familyName("Jenny")
	familyName("Anja")

}
*/


/*
func familyName(fname string, age int){
	fmt.Println("Hello", age, "year old", fname, "Refsnes")
}

func main(){
	familyName("Liam", 3)
	familyName("Jenny", 3)
	familyName("Anja", 3)
}
*/

/*

	If you want the function to return a value, you need to define the data type of the return value(such as int, string, etc), and also use the return 
	keyword inside the function. 

	func FunctionName(param1 type, param2 type) type {
		// code to be executed 
		return output ---> whatever 
	}




*/


/*
func myFunction(x int, y int) int {
	return x + y
}

func main(){
	fmt.Println(myFunction(1, 2))
}
*/

/*
func myFunction(x int, y int) (result int){
	result = x + y
	return 
}

func main(){
	fmt.Println(myFunction(1, 2))
}
*/


/*
func myFunction(x int, y int) (result int){
	result = x + y	
	return result 

}



func main(){
	fmt.Println(myFunction(1, 2))
}
*/


// storing the return value in a variable 

/*
func myFunction(x int, y int) (result int){
	result = x + y
	return

}


func main(){
	total := myFunction(1, 2) // variable type is inferred 
	fmt.Println(total)

}
*/

// multiple return values 

// go functions can return multiple values 

/*
func myFunction(x int, y string) (result int, txt1 string){
	result = x + x
	txt1 = y + " World!"
	return 
}


func main(){
	fmt.Println(myFunction(5, "Hello"))

}
*/


/*
	If we (for some reason) do not want to use some of the returned values, we can add an underscore(_) to omit this value.



*/

/*
func myFunction(x int, y string) (result int, txt1 string){
	result = x + x
	txt1 = y + " World!"
	return 

}


func main(){
	/*
	_, b := myFunction(5, "Hello") // we omit the first returned value result 
	fmt.Println(b)
	*/ 
	
	// we can omit the txt as well like 
	/*
	a, _ := myFunction(5, "Hello")
	fmt.Println(a)
	*/

//}











