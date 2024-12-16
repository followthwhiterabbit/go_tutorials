package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)




func Hello(name string) (string, error) {
	// if no name was given, return an error with a message

	if name == ""{
		return "", errors.New("empty name")
	}
	// if a name was received, return a value that embeds the name 
	// in a greeting message 
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}


// randomFormat returns one of a set of greeting messages. The returned 
// message is selected at random 
func randomFormat() string {
	// a slice of message formats 
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!", 
		"Hail, %v  Well met!", 
}

	// return a randomly selected message format by specifying a random index for the slice of formats. 
	return formats[rand.Intn(len(formats))]
}




