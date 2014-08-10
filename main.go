package main

import (
	"appengine"
	"appengine/datastore"
	"appengine/urlfetch"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

var m *martini.Martini

func init() {
	m := martini.Classic()
	m.Use(render.Renderer())

	m.Get("/api/search", func(r render.Render, req *http.Request) {
		u, _ := url.Parse(req.URL.String())
		params := u.Query()
		log.Println(params)
		c := appengine.NewContext(req)
		q := datastore.NewQuery("Card")
		q = EqualQuery(q, params, "bursted")
		q = EqualQuery(q, params, "category")
		q = EqualQuery(q, params, "color")
		q = EqualQuery(q, params, "constraint")
		q = EqualQuery(q, params, "expansion")
		q = EqualQuery(q, params, "guard")
		q = EqualQuery(q, params, "illus")
		q = EqualQuery(q, params, "reality")
		q = EqualQuery(q, params, "type")
		entities := make([]Card, 0, 10)
		if _, err := q.GetAll(c, &entities); err != nil {
			log.Println(err)
		}
		entities = Removenequality(entities, params, "costBlack")
		entities = Removenequality(entities, params, "costBlue")
		entities = Removenequality(entities, params, "costColorless")
		entities = Removenequality(entities, params, "costGreen")
		entities = Removenequality(entities, params, "costRed")
		entities = Removenequality(entities, params, "costWhite")
		entities = Removenequality(entities, params, "level")
		entities = Removenequality(entities, params, "limit")
		entities = Removenequality(entities, params, "power")
		r.JSON(200, entities)
	})

	m.Get("/api/model", func(r render.Render, req *http.Request) {
		c := appengine.NewContext(req)
		q := datastore.NewQuery("Card")
		entities := make([]Card, 0, 0)
		if _, err := q.GetAll(c, &entities); err != nil {
			log.Println(err)
		}
		products := make([]string, 0, 0)
		illus := make([]string, 0, 0)
		constraints := make([]string, 0, 0)
		types := make([]string, 0, 0)
		for _, card := range entities {
			if Contains(card.Expansion, products) == false {
				products = append(products, card.Expansion)
			}
			if Contains(card.Illus, illus) == false {
				illus = append(illus, card.Illus)
			}
			if Contains(card.Constraint, constraints) == false && card.Constraint != "" {
				constraints = append(constraints, card.Constraint)
			}
			if Contains(card.Type, types) == false && card.Type != "" {
				types = append(types, card.Type)
			}
		}
		for _, product := range products {
			p := &Product{}
			p.Id = product
			key := datastore.NewKey(c, "Product", product, 0, nil)
			key, err := datastore.Put(c, key, p)
			if err != nil {
				c.Criticalf("save error. ID: %v", product)
			} else {
				c.Infof("success. ID: %v", product)
			}
		}
		for _, i := range illus {
			illustrator := &Illustrator{}
			illustrator.Name = i
			key := datastore.NewKey(c, "Illustrator", i, 0, nil)
			key, err := datastore.Put(c, key, illustrator)
			if err != nil {
				c.Criticalf("save error. Name: %v", i)
			} else {
				c.Infof("success. Name: %v", i)
			}
		}
		for _, constraint := range constraints {
			con := &Constraint{}
			con.Type = constraint
			key := datastore.NewKey(c, "Constraint", constraint, 0, nil)
			key, err := datastore.Put(c, key, con)
			if err != nil {
				c.Criticalf("save error. Type: %v", constraint)
			} else {
				c.Infof("success. Type: %v", constraint)
			}
		}
		for _, ty := range types {
			t := &Type{}
			t.Name = ty
			key := datastore.NewKey(c, "Type", ty, 0, nil)
			key, err := datastore.Put(c, key, t)
			if err != nil {
				c.Criticalf("save error. Type: %v", ty)
			} else {
				c.Infof("success. Type: %v", ty)
			}
		}
		response := map[string][]string{}
		response["product"] = products
		response["illustrator"] = illus
		response["constraint"] = constraints
		response["type"] = types
		r.JSON(200, response)
	})

	m.Get("/api/product", func(r render.Render, req *http.Request) {
		c := appengine.NewContext(req)
		q := datastore.NewQuery("Product")
		res := make([]Product, 0, 0)
		if _, err := q.GetAll(c, &res); err != nil {
			log.Println(err)
		}
		r.JSON(200, res)
	})

	m.Get("/api/illustrator", func(r render.Render, req *http.Request) {
		c := appengine.NewContext(req)
		q := datastore.NewQuery("Illustrator")
		res := make([]Illustrator, 0, 0)
		if _, err := q.GetAll(c, &res); err != nil {
			log.Println(err)
		}
		r.JSON(200, res)
	})

	m.Get("/api/constraint", func(r render.Render, req *http.Request) {
		c := appengine.NewContext(req)
		q := datastore.NewQuery("Constraint")
		res := make([]Constraint, 0, 0)
		if _, err := q.GetAll(c, &res); err != nil {
			log.Println(err)
		}
		r.JSON(200, res)
	})

	m.Get("/api/type", func(r render.Render, req *http.Request) {
		c := appengine.NewContext(req)
		q := datastore.NewQuery("Type")
		res := make([]Type, 0, 0)
		if _, err := q.GetAll(c, &res); err != nil {
			log.Println(err)
		}
		r.JSON(200, res)
	})

	m.Get("/api/:from/:to", func(r render.Render, params martini.Params, req *http.Request) {
		for i := ToInt(params["from"]); i <= ToInt(params["to"]); i++ {
			CreateCard(i, req)
		}
		r.JSON(200, "finish!")
	})

	http.ListenAndServe(":8080", m)
	http.Handle("/", m)
}

func Contains(str string, list []string) bool {
	for _, s := range list {
		if s == str {
			return true
		}
	}
	return false
}

func EqualQuery(query *datastore.Query, values url.Values, s string) *datastore.Query {
	if len(values[s]) != 0 {
		if values[s][0] == "true" || values[s][0] == "false" {
			value, _ := strconv.ParseBool(values[s][0])
			query = query.Filter(strings.ToUpper(s[:1])+s[1:]+"=", value)
		} else {
			query = query.Filter(strings.ToUpper(s[:1])+s[1:]+"=", values[s][0])
		}
	}
	return query
}

func GreaterThanEqualQuery(query *datastore.Query, values url.Values, s string) *datastore.Query {
	if len(values[s]) != 0 {
		columnName := strings.TrimRight(strings.ToUpper(s[:1])+s[1:], "From")
		query = query.Filter(columnName+">=", ToInt(values[s][0])).Order(columnName)
	}
	return query
}

func Removenequality(cards []Card, values url.Values, s string) []Card {
	fromName := s + "From"
	toName := s + "To"
	if len(values[fromName]) != 0 || len(values[toName]) != 0 {
		fromValue := -1
		toValue := 100
		if len(values[fromName]) != 0 {
			fromValue = ToInt(values[fromName][0])
		}
		if len(values[toName]) != 0 {
			toValue = ToInt(values[toName][0])
		}
		columnName := strings.ToUpper(s[:1]) + s[1:]
		resultList := make([]Card, 0, 0)
		for _, card := range cards {
			columnValue := reflect.ValueOf(card).FieldByName(columnName).Int()
			if int64(fromValue) <= columnValue && columnValue <= int64(toValue) {
				resultList = append(resultList, card)
			}
		}
		return resultList
	}
	return cards
}

func CreateCard(id int, req *http.Request) *Card {

	card := &Card{}
	card.CostBlack = -1
	card.CostBlue = -1
	card.CostColorless = -1
	card.CostGreen = -1
	card.CostRed = -1
	card.CostWhite = -1

	url := "http://www.takaratomy.co.jp/products/wixoss/card/card_detail.php?id=" + strconv.Itoa(id)
	card.Url = url
	c := appengine.NewContext(req)
	client := urlfetch.Client(c)
	resp, _ := client.Get(url)

	doc, _ := goquery.NewDocumentFromResponse(resp)
	if doc.Find(".card_detail_title p").Text() == "" || strings.TrimLeft(doc.Find(".card_img").Text(), "Illust ") == "" {
		c.Warningf("data not found. id: %d", id)
		return card
	}
	doc.Find(".card_detail").Each(func(_ int, s *goquery.Selection) {

		s.Find(".card_detail_title").Each(func(_ int, s *goquery.Selection) {
			cardNo := strings.Split(s.Find("p").Text(), "-")
			card.Expansion, card.No = cardNo[0], ToInt(cardNo[1])
			cardName := strings.Split(s.Find("h3").Text(), "＜")
			card.Name = GetName(cardName[0], card)
			card.NameKana = strings.Trim(cardName[1], "＞")
		})
		card.Reality = strings.TrimRight(strings.TrimLeft(strings.TrimSpace(s.Find(".card_rarity").Text()), "\n\u0009"), "\n")

		s.Find(".card_date_box").Each(func(_ int, s *goquery.Selection) {
			s.Find(".card_img").Each(func(_ int, s *goquery.Selection) {
				s.Find("img").Each(func(_ int, s *goquery.Selection) {
					card.Image, _ = s.Attr("src")
				})
			})
			card.Illus = strings.TrimLeft(s.Find(".card_img").Text(), "Illust ")
			count := 0
			s.Find("tbody").Find("tr").Each(func(_ int, s *goquery.Selection) {
				count++
			})
			s.Find("tbody").Find("tr").Each(func(_ int, s *goquery.Selection) {
				if 0 == s.Index() {
					card.Category = s.Find("td").Eq(0).Text()
					card.Type = TrimHyphen(s.Find("td").Eq(1).Text())
				} else if 1 == s.Index() {
					card.Color = ToEng(s.Find("td").Eq(0).Text())
					card.Level = ToInt(s.Find("td").Eq(1).Text())
				} else if 2 == s.Index() {
					setColor(card, s.Find("td").Eq(0).Text())
					setColor(card, s.Find("td").Eq(1).Text())
				} else if 3 == s.Index() {
					card.Limit = ToInt(s.Find("td").Eq(0).Text())
					card.Power = ToInt(s.Find("td").Eq(1).Text())
				} else if 4 == s.Index() {
					card.Constraint = TrimHyphen(s.Find("td").Eq(0).Text())
					card.Guard = TrimHyphen(s.Find("td").Eq(1).Text())
				} else if 8 == count { // すべてあり
					if 5 == s.Index() {
						card.Text = GetText(s)
						card.SearchText = GetSearchText(s)
					} else if 6 == s.Index() {
						card.Burst = GetBurst(s)
						card.Bursted = (card.Burst != "")
					} else if 7 == s.Index() {
						card.Flavor = GetFlavor(s)
					}
				} else if 7 == count {
					if card.Category == "ルリグ" {
						if 5 == s.Index() {
							card.Text = GetText(s)
							card.SearchText = GetSearchText(s)
						} else if 6 == s.Index() {
							card.Flavor = GetFlavor(s)
						}
					} else if card.Category == "アーツ" {
						if 5 == s.Index() {
							card.Text = GetText(s)
							card.SearchText = GetSearchText(s)
						} else if 6 == s.Index() {
							card.Flavor = GetFlavor(s)
						}
					} else if card.Category == "スペル" {
						if 5 == s.Index() && IsLifeCloth(s) {
							card.Burst = GetBurst(s)
							card.Bursted = (card.Burst != "")
						} else if 5 == s.Index() {
							card.Text = GetText(s)
							card.SearchText = GetSearchText(s)
						} else if 6 == s.Index() && IsLifeCloth(s) {
							card.Burst = GetBurst(s)
							card.Bursted = (card.Burst != "")
						} else {
							card.Flavor = GetFlavor(s)
						}
					} else {
						if 5 == s.Index() && IsLifeCloth(s) == false {
							card.Text = GetText(s)
							card.SearchText = GetSearchText(s)
						} else if 5 == s.Index() {
							card.Burst = GetBurst(s)
							card.Bursted = (card.Burst != "")
						}
						if 6 == s.Index() && IsLifeCloth(s) {
							card.Burst = GetBurst(s)
							card.Bursted = (card.Burst != "")
						} else if 6 == s.Index() {
							card.Flavor = GetFlavor(s)
						}
					}
				} else if 6 == count {
					if card.Category == "ルリグ" {
						if IsCardText(s) {
							card.Text = GetText(s)
							card.SearchText = GetSearchText(s)
						} else {
							card.Flavor = GetFlavor(s)
						}
					} else if card.Category == "アーツ" {
						card.Text = GetText(s)
						card.SearchText = GetSearchText(s)
					} else if card.Category == "スペル" {
						card.Text = GetText(s)
						card.SearchText = GetSearchText(s)
					} else {
						if IsLifeCloth(s) {
							card.Burst = GetBurst(s)
							card.Bursted = (card.Burst != "")
						} else if IsCardText(s) {
							card.Text = GetText(s)
							card.SearchText = GetSearchText(s)
						} else {
							card.Flavor = GetFlavor(s)
						}
					}
				}
			})
		})
	})
	keyStr := card.Expansion + "-" + fmt.Sprintf("%03d", card.No)
	key := datastore.NewKey(c, "Card", keyStr, 0, nil)
	key, err := datastore.Put(c, key, card)
	if err != nil {
		c.Criticalf("save error. id: %d", id)
		log.Println(err)
	} else {
		c.Infof("success. id: %d", id)
	}
	return card
}

func ToInt(v string) int {
	if v == "-" {
		return -1
	}
	i, _ := strconv.Atoi(v)
	return i
}

func GetName(s string, card *Card) string {
	if card.Expansion == "PR" {
		return strings.Split(TrimLinefeed(s), "(")[0]
	}
	return s
}

func GetFlavor(s *goquery.Selection) string {
	flavor := TrimLinefeed(s.Find("td").Text())
	return strings.Replace(flavor, "デッキレベル０再録", "", -1)
}

func GetBurst(s *goquery.Selection) string {
	return strings.TrimLeft(TrimLinefeed(s.Find("td").Text()), "：")
}

func GetText(s *goquery.Selection) string {
	texts, _ := s.Find("td").Html()
	texts = TrimLinefeed(texts)
	texts = strings.Replace(texts, "<br/>", ",", -1)
	texts = strings.Replace(texts, "デッキレベル０再録", "", -1)
	return ReplaceIcon(texts)
}

func GetSearchText(s *goquery.Selection) string {
	return "" // TrimLinefeed(s.Find("td").Text())
}

func TrimLinefeed(v string) string {
	return strings.TrimRight(strings.TrimLeft(strings.TrimSpace(v), "\n"), "\n")
}

func ToEng(v string) string {
	switch {
	case "黒" == v:
		return "black"
	case "青" == v:
		return "blue"
	case "無" == v:
		return "colorless"
	case "赤" == v:
		return "red"
	case "緑" == v:
		return "green"
	case "白" == v:
		return "white"
	}
	return ""
}

func TrimHyphen(v string) string {
	if v == "-" {
		return ""
	}
	return v
}

func setColor(card *Card, str string) {
	costs := strings.Split(str, "、")
	for _, val := range costs {
		cost := strings.Split(val, "×")
		switch {
		case "黒" == cost[0]:
			card.CostBlack = ToInt(cost[1])
		case "青" == cost[0]:
			card.CostBlue = ToInt(cost[1])
		case "無" == cost[0]:
			card.CostColorless = ToInt(cost[1])
		case "赤" == cost[0]:
			card.CostRed = ToInt(cost[1])
		case "緑" == cost[0]:
			card.CostGreen = ToInt(cost[1])
		case "白" == cost[0]:
			card.CostWhite = ToInt(cost[1])
		case "0" == cost[0]:
			switch {
			case card.Color == "black":
				card.CostBlack = 0
			case card.Color == "blue":
				card.CostBlue = 0
			case card.Color == "colorless":
				card.CostColorless = 0
			case card.Color == "red":
				card.CostRed = 0
			case card.Color == "green":
				card.CostGreen = 0
			case card.Color == "white":
				card.CostWhite = 0
			}
		}
	}
}

func ReplaceIcon(str string) string {
	str = strings.Replace(str, "<img src=\"../images/icon_txt_null_01.png\" width=\"26\" height=\"23\" alt=\"無×1\"/>", "(無)", -1)
	str = strings.Replace(str, "<img src=\"/products/wixoss/images/icon_txt_starting.png\" width=\"26\" height=\"23\" alt=\"起動能力\"/>", "[起]", -1)
	str = strings.Replace(str, "<img src=\"/products/wixoss/images/icon_txt_regular.png\" width=\"26\" height=\"23\" alt=\"常\"/>", "[常]", -1)
	str = strings.Replace(str, "<img src=\"/products/wixoss/images/icon_txt_regular.png\" width=\"26\" height=\"23\" alt=\"常時能力\"/>", "[常]", -1)
	str = strings.Replace(str, "<img src=\"/products/wixoss/images/icon_txt_arrival.png\" width=\"26\" height=\"23\" alt=\"出現時能力\"/>", "[出]", -1)
	str = strings.Replace(str, "<img src=\"/products/wixoss/images/icon_txt_white.png\" width=\"26\" height=\"23\" alt=\"白\"/>", "(白)", -1)
	str = strings.Replace(str, "<img src=\"/products/wixoss/images/icon_txt_red.png\" width=\"26\" height=\"23\" alt=\"赤\"/>", "(赤)", -1)
	str = strings.Replace(str, "<img src=\"/products/wixoss/images/icon_txt_blue.png\" width=\"26\" height=\"23\" alt=\"青\"/>", "(青)", -1)
	str = strings.Replace(str, "<img src=\"/products/wixoss/images/icon_txt_green.png\" width=\"26\" height=\"23\" alt=\"緑\"/>", "(緑)", -1)
	str = strings.Replace(str, "<img src=\"/products/wixoss/images/icon_txt_black.png\" width=\"26\" height=\"23\" alt=\"黒\"/>", "(黒)", -1)
	str = strings.Replace(str, "<img src=\"/products/wixoss/images/icon_txt_null.png\" width=\"26\" height=\"23\" alt=\"無\"/>", "(無)", -1)
	str = strings.Replace(str, "<img src=\"/products/wixoss/images/icon_txt_down.png\" width=\"26\" height=\"23\" alt=\"タップ\"/>", "(Ｔ)", -1)
	return str
}

func IsLifeCloth(s *goquery.Selection) bool {
	html, _ := s.Find("td").Html()
	return -1 != strings.Index(html, "icon_txt_burst.png")
}

func IsCardText(s *goquery.Selection) bool {
	html, _ := s.Find("td").Html()
	return -1 != strings.Index(html, "icon_txt_") && -1 == strings.Index(html, "icon_txt_burst.png")
}
