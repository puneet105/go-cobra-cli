/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// FibonacciSeriesCmd represents the FibonacciSeries command
var FibonacciSeriesCmd = &cobra.Command{
	Use:   "FibonacciSeries",
	Short: "Command to print a fibonacci series based on the number we provide as an argument",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("FibonacciSeries called....!!")
		printFibonacci(args[0])
	},
}

func printFibonacci(num string){
	x := 0
	y := 1
	numInt,_ := strconv.Atoi(num)
	fmt.Printf("Fibonacci Series is : %d %d", x,y)
	for i:=2;i<numInt;i++{
		z := x + y
		fmt.Printf(" %d", z)
		x = y
		y = z
	}
}

func init() {
	rootCmd.AddCommand(FibonacciSeriesCmd)

}
