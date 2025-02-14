package main

import (
	"fmt"
)

func main() {
	var revenue float32
	var expences float32
	var taxRate float32

	fmt.Print("Enter your revenue, $: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter your expences, $: ")
	fmt.Scan(&expences)

	fmt.Print("Enter your taxRate, %: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expences
	profit := ebt - ebt*(taxRate/100)

	fmt.Println(`Earnings before taxes - `, ebt)
	fmt.Println(`Profit - `, profit)
	fmt.Println(`Ratio - `, ebt/profit)
}
