package main

import (
	"fmt"
	"goapi/amazon"
	"goapi/trendyol"
	"os"
)

func main() {
	var choice int
	fmt.Println("1. Trendyol Bestseller")
	fmt.Println("2, Amazon bestsellers")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		trendyol.Trendyol()
	case 2:
		amazon.Amazon()
	default:
		os.Exit(1)

	}
}
