package main
import "fmt"

func main() {
	// To create an empty map,
	// use the builtin make: make(map[key-type]val-type).
	m := make(map[string]int) 

	// Set key/value pairs using typical name[key] = val syntax.

	m["a"] = 1
	m["b"] = 2

	fmt.Println("map:",m)
}