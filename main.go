package main

import (
	"errors"
	"fmt"
)

func main() {
	// messages := []string{"1", "2", "3"}
	messages := make([]string, 5, 15)
	messages[0] = "1"
	// printMessage(messages)

	fmt.Println(messages)
}
func printMessage(messages []string) error {

	if len(messages) == 0 {
		return errors.New("empty array")
	}
	messages[1] = "5"
	fmt.Println(messages)

	return nil
}
