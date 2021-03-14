package app

import (
	"figurer/app/calc"
	"fmt"
	log "github.com/sirupsen/logrus"
)

type Calculations struct{
	Num1 string
	Operator string
	Num2 string
	Result float64
}

func Run() *Calculations{

	log.Info("Starting the Figurer app...\n")
	var cal Calculations

	fmt.Println("Enter the operator:\nAdd: +\nSubtract: -\nMultiply: *\nDivide: /")
	fmt.Scanln(&cal.Operator)

	fmt.Println("Enter your first number:")
	fmt.Scanln(&cal.Num1)

	fmt.Println("Enter your second number:")
	fmt.Scanln(&cal.Num2)

	cal.Result = calc.Calculate(cal.Num1, cal.Num2, cal.Operator)
	fmt.Println("Your result is", cal.Result)

	return &cal
}
