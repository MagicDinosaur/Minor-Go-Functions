package crawler

import (
	"crawdata/database"
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Craw(url string) {
	var ListTivi = make([]database.Tivi, 1)

	resp, err := http.Get(url)
	if err != nil || (resp != nil && (resp.StatusCode > 299 || resp.StatusCode < 200)) {
		fmt.Println("failed to fetch URL")
		return
	}
	defer resp.Body.Close()

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		fmt.Println("Error loading HTTP response body.", err)
		return
	}

	var i int64 = 0
	document.Find(".gallery-item-inner").Each(func(index int, element *goquery.Selection) {
		id := i
		i++
		name := element.Find(".product-name-wrap strong").Text()
		series := element.Find(".product-name-wrap product-model l3").Text()
		label, _ := element.Find(".product-logo-overlay hidden").Attr("src")
		img, _ := element.Find(".ghost-center > img").Attr("src")
		var description []string
		element.Find(".p4 > span").Each(func(index int, element1 *goquery.Selection) {
			description = append(description, element1.Text())
		})
		var sizes []string
		element.Find(".prime-differentiator-list span").Each(func(index int, element2 *goquery.Selection) {
			sizes = append(sizes, element2.Text())
		})
		var prices []string
		element.Find(".product-price strong").Each(func(index int, element3 *goquery.Selection) {
			prices = append(prices, element3.Text())
		})

		newTivi := database.Tivi{
			ID:          id,
			Name:        name,
			Series:      series,
			Label:       label,
			Img:         img,
			Description: description,
			Sizes:       sizes,
			Prices:      prices,
		}
		ListTivi = append(ListTivi, newTivi)
		fmt.Println(newTivi)
		fmt.Println(i)
	})
	fmt.Println("ok")
}
