package main
import (
	"fmt"
)

func main(){

	var revenue float64
	var expenses float64
	var tax float64
	fmt.Print("Enter the revenue of your company:")
	fmt.Scan(&revenue)
	fmt.Print("Enter the expenses of your company:")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate:")
	fmt.Scan(&tax)
	EBT := (revenue - expenses)
	fmt.Println("Earnings before tax:",EBT)
	PROFIT:=(1-tax/100) * EBT
	fmt.Println("Get Profit:",PROFIT)
	PROFITRATIO:=EBT/PROFIT
	fmt.Println("Profit ratio:",PROFITRATIO)

}