/*
	The for loop loops through a block of code a specified number of times 

	the for loop is the only available loop in Go
	
	for statement1; statement2; statement3 {
		// code to be executed for each iteration 


	}


*/ 

package main 
import ("fmt")


/*
func main(){
	for i:=0; i < 5; i++{
		fmt.Println(i)
	}

}
*/

func main(){
	/*
	for i:=0; i <= 100; i+=10{
		fmt.Println(i)
	}
	*/

/*	
	for i:=0; i < 5; i++{
		if i == 3{
			continue
		}
		fmt.Println(i)
	}

	*/

// break statement 
/*
for i:=0; i < 5; i++{
	if i == 3{
		break
	}
	fmt.Println(i)

}
*/


/*
	nested loops 

	
*/

/*
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i:=0; i < len(adj); i++{
		for j:=0; j < len(fruits); j++{
			fmt.Println(adj[i], fruits[j])

		}
	}
*/

/*
	Range keyword 
	 - the range keyword is used to more easily iterate through the elements of the array, slice or map. It returns both the index and the value. 
	
	 for index, value := range range array|slice|map {
		// code to be executed for each iteration 
	
		}
	


*/


/*
fruits := [3]string{"apple", "orange", "banana"}
for idx, val := range fruits{
	fmt.Printf("%v\t%v\n", idx, val)
	
}
*/
/*
fruits := [3]string{"apple", "orange", "banana"}
for _, val := range fruits { // _ helps to avoid printing the indexes 
	fmt.Printf("%v\n", val)

}
*/

fruits := [3]string{"apple", "orange", "banana"}
for idx, _ := range fruits{ // helps to avoid printing the values 
	fmt.Printf("%v\n",idx)		
}

}






