package greeting

import "fmt"

func Hello() {
	fmt.Println("Hello!")
}

func Hi() { // Capitalized function names to be used outside of package
	fmt.Println("Hi!")
}