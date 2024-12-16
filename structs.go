/*
	structs in go language 

	type struct_name struct {
		member1 datatype;
		member2 datatype;
		member3 datatype;
		...
	}



*/

package main
import ("fmt")

type Person struct { // struct variable named Person with the type struct
	name string
	age int
	job string
	salary int
	
}


func main(){

	var pers1 Person 
	var pers2 Person 


	// pers1 sepecification 
	pers1.name = "Hege"
	pers1.age = 45
	pers1.job = "Teacher"
	pers1.salary = 6000
	

	// pers2 sepecification 
	pers2.name = "Cecilie"
	pers2.age = 24
	pers2.job = "Marketing"
	pers2.salary  = 4500
	

	// access and print Pers1 info
	fmt.Println("Name: ", pers1.name)
	fmt.Println("Age: ", pers1.age)
	fmt.Println("Job: ", pers1.job)
	fmt.Println("Salary: ", pers1.salary)

	fmt.Println("Name: ", pers2.name)
	fmt.Println("Age: ", pers2.age)
	fmt.Println("Job: ", pers2.job)
	fmt.Println("Salary: ", pers2.salary)

	
	// pass struct as function arguments 
	
	printPerson(pers1)  // printPerson takes a struct now 
	
	printPerson(pers2) // printPerson takes a struct now 




}


func printPerson(pers Person){ // takes a struct of Person type with the variable named pers
	fmt.Println("Name: ", pers.name)
	fmt.Println("Age: ", pers.age)
	fmt.Println("Job: ", pers.job)
	fmt.Println("Salary: ", pers.salary)
}


