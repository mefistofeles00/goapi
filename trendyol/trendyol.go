package trendyol

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
)

func Trendyol() {
	c := colly.NewCollector()
	c.OnHTML(".product-card", func(e *colly.HTMLElement) {
		fmt.Println(e.Text + "\n")
	})
	err := c.Visit("https://www.trendyol.com/cok-satanlar?type=bestSeller&webGenderId=1")
	if err != nil {
		log.Fatal(err)
	}
}
