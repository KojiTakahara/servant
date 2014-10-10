package servant

import (
	"appengine"
	"appengine/datastore"
	"appengine/urlfetch"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

func CreateCard(id int, req *http.Request) *Card {
	card := &Card{}
	card.CostBlack = -1
	card.CostBlue = -1
	card.CostColorless = -1
	card.CostGreen = -1
	card.CostRed = -1
	card.CostWhite = -1
	card.Id = id

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
					SetColor(card, s.Find("td").Eq(0).Text())
					SetColor(card, s.Find("td").Eq(1).Text())
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

	cards := GetCardByNameAndId(card.Name, id, req)
	if len(cards) != 0 {
		card.ParentKeyName = cards[0].KeyName
	}

	keyStr := card.Expansion + "-" + fmt.Sprintf("%03d", card.No)
	key := datastore.NewKey(c, "Card", keyStr, 0, nil)
	key, err := datastore.Put(c, key, card)
	if err != nil {
		c.Criticalf("save error. id: %d", id)
		c.Criticalf(err.Error())
	} else {
		c.Infof("success. id: %d", id)
	}
	return card
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

func SetColor(card *Card, str string) {
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
