package calc

import (
	"fmt"
	"math"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func stringToFloat64(str string) float64 {
	n, err := strconv.ParseFloat(str, 64)
	if err != nil {
		log.Fatal(err)
	}
	return n
}

// Calculate the result based on inputs
func Calculate(num1, num2, op string) float64 {

	log.Info("Calculating...")
	var firstNum float64 = stringToFloat64(num1)
	var secondNum float64 = stringToFloat64(num2)

startAgain:
	switch op {
	case "+":
		return math.Round((firstNum+secondNum)*100) / 100

	case "-":
		return math.Round((firstNum-secondNum)*100) / 100

	case "*":
		return math.Round((firstNum*secondNum)*100) / 100

	case "/":
		return math.Round((firstNum/secondNum)*100) / 100

	default:
		fmt.Println("Please choose an operator from below:\n +, -, *, /")
		fmt.Scanln(&op)
		goto startAgain
	}
}
