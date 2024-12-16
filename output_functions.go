package main 
import ("fmt")

func main(){
	
	//var i, j string = "Hello", "World"

	/*
	fmt.Print(i, "\n") // we can use these seperately with new lines but also 
	fmt.Print(j, "\n")
	*/ 

	// fmt.Print(i, "\n", j, "\n") // we can multiple variables and the characters in the same PRint func 
	
	// fmt.Print(i, " ", j, "\n")

	// print puts a space between them if they are not strings 

	/*
	var a, b = 10, 20 // if you try to compile while the privous variabls are not used, it gives the following errors 
	# command-line-arguments
	./output_functions.go:6:6: declared and not used: i
	./output_functions.go:6:9: declared and not used: j
	


	fmt.Print(a, b)
	*/

	// Printf() function first formats its argument based on the given formatting verb and then prints them 

	// 2 formatting verbs, %v ---> printing the value of the argumetns, %T ----> is used to print the type of the arguments 


	/*
	var i string = "Hello"
	var j int = 15

	fmt.Printf("i has value : %v and type : %T\n", i, i)
	fmt.Printf("j has value : %v and type : %T\n", j, j)
	*/ 


	/*
	var i = 15.5 
	var txt = "Hello World!"

	fmt.Printf("%v\n", i)
	fmt.Printf("%#v\n", i)
	fmt.Printf("%v%%\n", i)
	fmt.Printf("%T\n", i)
	
	fmt.Printf("%v\n", txt)
	fmt.Printf("%#v\n", txt)
	fmt.Printf("%T\n", txt)
	*/ 

	/*
	i := 32

	fmt.Printf("%X\n", i)
	*/

	/*
		
	  var i = 16
 
  fmt.Printf("%b\n", i)
  fmt.Printf("%d\n", i)
  fmt.Printf("%+d\n", i)
  fmt.Printf("%o\n", i)
  fmt.Printf("%O\n", i)
  fmt.Printf("%x\n", i)
  fmt.Printf("%X\n", i)
  fmt.Printf("%#x\n", i)
  fmt.Printf("%4d\n", i)
  fmt.Printf("%-4d\n", i)
  fmt.Printf("%04d\n", i)
	*/ 


	/*
	var txt = "Hello. My name is KAAN."

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt) // hex dump  of the given string 

	fmt.Printf("% x\n", txt) // hex dump with spaces 
	*/ 

	// boolean formatting verbs 
	//var i = true 
	//var j = false 

	
	// Float formatting Verbs 
	/*
		%e -- scientific notation with 'e' as exponent 
		%f -- decimal point, no exponent 
		%.2f -- default width, precision 2 
		%6.2f -- width 6, precision 2
		%g -- exponent as needed , only necessary digits 
	*/

	var i = 3.141

	fmt.Printf("%e\n", i)
	fmt.Printf("%f\n", i)
	fmt.Printf(".2f\n", i)
	fmt.Printf("6.2f\n", i)
	fmt.Printf("%g\n", i)






}
