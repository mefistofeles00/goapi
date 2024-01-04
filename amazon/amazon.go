package amazon

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
)

var p1 = fmt.Println

func Amazon() {
	c := colly.NewCollector()
	c.OnHTML(".a-carousel-card", func(e *colly.HTMLElement) {
		p1(e.Text + "\n")
	})
	err := c.Visit("https://www.amazon.com/Best-Sellers/zgbs")

	if err != nil {
		log.Fatal(err)
	}]

}
