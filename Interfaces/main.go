package main

import "fmt"

type bot interface {
	getGreeting() string
	// we can add multiple arg types and return types
	//we can have as many function as we need
	//we dont need to impletement all functions
	//Interface does not implement type Generics in Java
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func (englishBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
