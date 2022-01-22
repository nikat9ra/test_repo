package main

import (
	"fmt"
)

func main() {
	printMessage("Vasya")
	printMessage("Petya")
	printMessage("Ivan")
}
func printMessage(message string) {
	fmt.Println(message)
}
