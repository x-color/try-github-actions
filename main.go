package main

import "fmt"

func main() {
	fmt.Println(greet("World"))
}

func greet(target string) string {
	return "Hello " + target + "!!"
}
