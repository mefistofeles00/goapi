package alibaba

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
)

var pl = fmt.Println

func Alibaba() {
	c := colly.NewCollector()
	c.OnHTML(".multi--outWrapper--SeJ8lrF card--out-wrapper", func(e *colly.HTMLElement) {
		pl(e.Text + "\n")

	})
	err := c.Visit("https://www.aliexpress.com/w/wholesale-alibaba-best-sellers.html")
	if err != nil {
		log.Fatal(err)
	}
}
