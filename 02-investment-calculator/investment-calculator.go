package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	outputText("Please, enter the investment amount:")
	fmt.Scan(&investmentAmount)

	outputText("Expected return rate:")
	fmt.Scan(&expectedReturnRate)

	outputText("Years:")
	fmt.Scan(&years)

    futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years, inflationRate)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future value: %v", futureValue)
	fmt.Println("Future real value: %f", futureRealValue)
    
}

func outputText(text string) {
    fmt.Println(text)
}

func calculateFutureValues(investmentAmount float64, expectedReturnRate float64, years float64, inflationRate float64) (float64, float64) {
    fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
    rfv := fv / math.Pow(1+inflationRate/100, years)
    return fv, rfv
}
