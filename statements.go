// the if statement



package main 
import ("fmt")



func main(){
	/*
	if 20 > 18{
		fmt.Println("20 is greater than 18")
	}
	*/


	/*
	x := 20
	y := 18

	if x > y{ // it gives unexpected newline error, if you put a newline next to the condition 
		fmt.Println("x is greater than y")
	}
	*/


	// if- else 
	/*
	time := 20
	if (time < 18) {
		fmt.Println("good day")
	}else{ // here the else statement should follow the closing curly brackets, and we should not leave a newline after the if statements without a curly bracket

		fmt.Println("good evening ")

	}
	*/ 

	// if - else if 
	/*

	*/
/*
	time := 22
	if time < 10{
	fmt.Println("Good morning.")
	} else if time < 20 {
		fmt.Println("Good day") 
	} else {
		fmt.Println("good evening.")
	}

*/

/*
	a := 14
	b := 14
	if a < b{
		fmt.Println("a is less than b")
	} else if a > b {	
		fmt.Println("a is greater than b")	
	} else {
		fmt.Println("a and b are equal")
	}
*/


/*
	x := 30
	if x >= 10{ 
		fmt.Println("x is larger than or equal to 10")  // if condition 1 and condition 2 are both true, only the code for condition1 are executed. 
	} else if x >20 {	
		fmt.Println("x is larger than 20")
	} else {	
		fmt.Println("x is less than 10")
	}
*/

/*
 nested if statement 
if condition1 {
	//  code to be executed if condition1 is true
	if condition2 { // code to be executed if both condition1 and condition2 are true 

		}
	}

} 
*/

num := 20 
if num >= 10 {
	fmt.Println("Num is more than 10.")
	if num > 15{
		fmt.Println("Num is also more than 15.")
	}

}else {
	fmt.Println("Num is less than 10.")
}

		










}
