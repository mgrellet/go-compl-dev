package main

import "fmt"

//Interface with 1 method
type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {

	//eb and sb are "members" of type bot
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

//One method, but 2 differents implementation
func printGreeting(b bot){
	fmt.Print(b)
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	//Very custom logic for generate english greeting
	return "Hi"
}

func (sb spanishBot) getGreeting() string {
	//Very custom logic for generate spanish greeting
	return "Hola"
}
