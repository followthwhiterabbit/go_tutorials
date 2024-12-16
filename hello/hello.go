package main 

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"example.com/greetings"



)


func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// prompting the user to enter their name 
	fmt.Print("Enter your name: ")
	
	// reading the input from the terminal 
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')


	// Request a greeting message 
	// If an error was returned, print it to the console and 
	// exit the program 
	if err != nil{
		log.Fatal("Error reading input:", err)
	}
	
	// trim any extra whitespace or newline characters
	name = strings.TrimSpace(name)

	// request a greeting message 
	message, err := greetings.Hello(name)

	
	if err != nil {
		log.Fatal(err)

	}
	// if no error was returned, print the returned message 
	fmt.Println(message)



}


