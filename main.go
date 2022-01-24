package main

import (
	"fmt"
)

func main() {
	message := "Хеллоу ворлд"
	fmt.Println(&message)
	fmt.Println(message)
	changeMessage(&message)

	fmt.Println(message)

}

func changeMessage(message *string) {

	*message += " от функции изменить"
}
