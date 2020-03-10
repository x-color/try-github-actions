package main

import "fmt"

func main() {
	fmt.Println(greet("World"))
	// fmt.Println(greet("GitHub"))
}

func greet(target string) string {
	return "Hello " + target + "!!"
}
