package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	// "net/http"
	// "regexp"
	"strings"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/api", func(r render.Render) {
		r.JSON(200, map[string]interface{}{"hello": "world"})
	})

	m.Get("/:id", func(r render.Render, params martini.Params) {
		result := map[string]string{}
		url := "http://www.takaratomy.co.jp/products/wixoss/card/card_detail.php?id=" + params["id"]

		result["url"] = url
		doc, _ := goquery.NewDocument(url)
		doc.Find(".card_detail").Each(func(_ int, s *goquery.Selection) {

			s.Find(".card_detail_title").Each(func(_ int, s *goquery.Selection) {
				result["no"] = s.Find("p").Text()
				result["name"] = s.Find("h3").Text()
			})
			result["reality"] = s.Find(".card_rarity").Text()

			s.Find(".card_date_box").Each(func(_ int, s *goquery.Selection) {
				s.Find(".card_img").Each(func(_ int, s *goquery.Selection) {
					s.Find("img").Each(func(_ int, s *goquery.Selection) {
						result["image"], _ = s.Attr("src")
					})
				})
				result["illus"] = strings.TrimLeft(s.Find(".card_img").Text(), "Illust ")
				s.Find("tbody").Find("tr").Each(func(_ int, s *goquery.Selection) {
					switch {
					case 0 == s.Index():
						result["category"] = s.Find("td").Eq(0).Text()
						result["type"] = s.Find("td").Eq(1).Text()
					case 1 == s.Index():
						result["color"] = s.Find("td").Eq(0).Text()
						result["level"] = s.Find("td").Eq(1).Text()
					case 2 == s.Index():
						result["grow"] = s.Find("td").Eq(0).Text()
						result["cost"] = s.Find("td").Eq(1).Text()
					case 3 == s.Index():
						result["limit"] = s.Find("td").Eq(0).Text()
						result["power"] = s.Find("td").Eq(1).Text()
					case 4 == s.Index():
						result["constraint"] = s.Find("td").Eq(0).Text()
						result["guard"] = s.Find("td").Eq(1).Text()
					case 5 == s.Index():
						result["text"], _ = s.Find("td").Html()
					case 6 == s.Index():
						result["flavor"] = s.Find("td").Text()
					}
				})
			})
		})
		fmt.Println("finish!")
		r.JSON(200, result)
	})

	m.Run()
}
