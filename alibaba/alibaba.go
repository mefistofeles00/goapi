package alibaba

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

var pl = fmt.Println

func Alibaba() {
	res, err := http.Get("https://www.aliexpress.com/w/wholesale-alibaba-best-sellers.html")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find(".list--gallery--C2f2tvm").Each(func(i int, s *goquery.Selection) {
		// Çeşitli sınıflardan veya etiketlerden verileri çekme
		title := s.Find(".multi--outWrapper--SeJ8lrF").Text()

		// Verileri ekrana yazdırma
		pl("\n")
		fmt.Printf("Ürün #%d:\n", i+5)

		fmt.Printf("Başlık: %s\n", title)
		pl("\n")
	})

}
