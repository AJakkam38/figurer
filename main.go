package main

import (
	"figurer/calc"
	"fmt"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
	log.SetLevel(log.InfoLevel)
}

func main() {

	log.Info("Starting the Figurer app...\n")

	var first, second, operator string
	var result float64

	fmt.Println("Enter the operator:\nAdd: +\nSubtract: -\nMultiply: *\nDivide: /")
	fmt.Scanln(&operator)

	fmt.Println("Enter your first number:")
	fmt.Scanln(&first)

	fmt.Println("Enter your second number:")
	fmt.Scanln(&second)

	result = calc.Calculate(first, second, operator)
	fmt.Println("Your result is", result)
}
