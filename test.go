// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Printf("We are open %d days per %s", 7, "week")
// }

package main

import "fmt"

func printmessage(s string) {
	fmt.Println(s)
}

func main() {
	// a goroutine
	go printmessage("Hello World!")
	// another go routine
	go printmessage("Welcome to Golang Goroutines.")

	fmt.Println("End of the main goroutine.")
}
