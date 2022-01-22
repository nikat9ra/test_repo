package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	message, err := prediction("вава")
	if err != nil {
		log.Print(err)
		return
	}
	fmt.Println(message)

}
func sayHello(name string, age int) string {
	return fmt.Sprintf("Hello, %s! You're %d years old", name, age)
}

func enterTheClub(age int) (string, error) {
	if age >= 18 && age < 45 {
		return "Проходи, только аккуратно", nil
	} else if age >= 45 && age < 65 {
		return "Вам точно понраиться эта музыка?", nil
	} else if age >= 65 {
		return "Это уже слишком для Вас", errors.New("you are too old")
	}
	return "Ты не пройдешь", errors.New("you are too young")
}

func prediction(dayOfWeek string) (string, error) {
	switch dayOfWeek {
	case "пн":
		return "Хорошего понедельника", nil
	case "вт":
		return "Хорошего вторника", nil
	default:
		return "Такого дня недели не существует", errors.New("invalid day of week")
	}

}
