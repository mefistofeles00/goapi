package main

import (
	"fmt"
	"goapi/alibaba"
	"goapi/amazon"
	"goapi/trendyol"
	"os"
)

func main() {
	var choice int
	fmt.Println("1. Trendyol Bestseller")
	fmt.Println("2, Amazon bestsellers")
	fmt.Println("3. Alibaba bestsellers (beta)")
	fmt.Scan(&choice)

	switch choice {
	case 1:
		trendyol.Trendyol()
	case 2:
		amazon.Amazon()
	case 3:
		alibaba.Alibaba()

	default:
		os.Exit(1)

	}
}
