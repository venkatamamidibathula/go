package main

import (
	"fmt"
	"math"
)

func main() {

	const inflation float64 = 6.5
	investment := 10000.0
	roi := 5.5
	years := 10.0

	returoninvestment := investment * math.Pow(1+roi/100, years)
	fmt.Println(returoninvestment)
	adjustedreturnoninvestment := returoninvestment / (math.Pow(1+inflation/100, years))
	fmt.Println(adjustedreturnoninvestment)

}
