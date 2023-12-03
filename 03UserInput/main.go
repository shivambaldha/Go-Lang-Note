package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	info_msg := "This is how to take a input from the user"
	fmt.Println(info_msg)

	reader := bufio.NewReader(os.Stdin) // so NewReader read from where, it's read  Read from stdin
	fmt.Println("Enter your rating here: ")

	// now syntax is comma ok or  error ok

	input, _ := reader.ReadString('\n') // this is read the till the next line that why we use '\n' , _ store the error 
	fmt.Println("thans for rating", input)
	fmt.Printf("data type is %T", input)

}
