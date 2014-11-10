package servant

import (
	"appengine"
	"appengine/datastore"
	"container/list"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func GetCardList(r render.Render, req *http.Request) []Card {
	u, _ := url.Parse(req.URL.String())
	params := u.Query()
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
	keys, err := q.GetAll(c, &entities)
	if err != nil {
		c.Criticalf(err.Error())
	}
	for i := range entities {
		entities[i].KeyName = keys[i].StringID()
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
	return entities
}

func CreateModels(r render.Render, req *http.Request) map[string][]string {
	c := appengine.NewContext(req)
	q := datastore.NewQuery("Card")
	entities := make([]Card, 0, 0)
	if _, err := q.GetAll(c, &entities); err != nil {
		c.Criticalf(err.Error())
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
	return response
}

func GetProductList(r render.Render, req *http.Request) []Product {
	c := appengine.NewContext(req)
	q := datastore.NewQuery("Product")
	res := make([]Product, 0, 0)
	if _, err := q.GetAll(c, &res); err != nil {
		c.Criticalf(err.Error())
	}
	return res
}

func GetIllustratorList(r render.Render, req *http.Request) []Illustrator {
	c := appengine.NewContext(req)
	q := datastore.NewQuery("Illustrator")
	res := make([]Illustrator, 0, 0)
	if _, err := q.GetAll(c, &res); err != nil {
		c.Criticalf(err.Error())
	}
	return res
}

func GetConstraintList(r render.Render, req *http.Request) []Constraint {
	c := appengine.NewContext(req)
	q := datastore.NewQuery("Constraint")
	res := make([]Constraint, 0, 0)
	if _, err := q.GetAll(c, &res); err != nil {
		c.Criticalf(err.Error())
	}
	return res
}

func GetTypeList(r render.Render, req *http.Request) []Type {
	c := appengine.NewContext(req)
	q := datastore.NewQuery("Type")
	res := make([]Type, 0, 0)
	if _, err := q.GetAll(c, &res); err != nil {
		c.Criticalf(err.Error())
	}
	return res
}

func CreateDeck(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)

	q := datastore.NewQuery("Card")
	entities := make([]Card, 0, 0)
	if _, err := q.GetAll(c, &entities); err != nil {
		c.Criticalf(err.Error())
	}
	deck := &Deck{}
	deck.Title = "ほげほげ"
	deck.Introduction = "intoro"
	deck.Description = "desc"
	deck.CreatedAt = time.Now()
	deck.UpdatedAt = time.Now()
	deck.Main01 = entities[0].KeyName
	key := datastore.NewKey(c, "Deck", "", 0, nil)
	c.Infof("%s", key)
	key, err := datastore.Put(c, key, deck)
	if err != nil {
		c.Criticalf("%s", err)
	} else {
		c.Infof("success. IntID: %s", key.IntID())
		c.Infof("success. StringID: %s", key.StringID())
	}
	r.JSON(200, deck)
}

func GetCardByExpansion(r render.Render, params martini.Params, req *http.Request) {
	c := appengine.NewContext(req)
	q := datastore.NewQuery("Card")
	q = q.Filter("Expansion=", params["expansion"])
	cards := make([]Card, 0, 10)
	keys, err := q.GetAll(c, &cards)
	if err != nil {
		c.Criticalf(err.Error())
	}
	for i := range cards {
		cards[i].KeyName = keys[i].StringID()
	}
	r.JSON(200, cards)
}

func GetCard(r render.Render, params martini.Params, req *http.Request) {
	c := appengine.NewContext(req)
	keyStr := params["expansion"] + "-" + fmt.Sprintf("%03d", ToInt(params["no"]))
	card := GetCardByKey(keyStr, c)
	r.JSON(200, card)
}

func GetDeck(r render.Render, params martini.Params, req *http.Request) {
	c := appengine.NewContext(req)
	id, _ := strconv.Atoi(params["id"])
	key := datastore.NewKey(c, "Deck", "", int64(id), nil)
	var deck Deck
	if err := datastore.Get(c, key, &deck); err != nil {
		c.Criticalf(err.Error())
	}
	viewDeck := &ViewDeck{}
	viewDeck.Id = deck.Id
	viewDeck.Owner = deck.Owner
	viewDeck.Title = deck.Title
	viewDeck.Introduction = deck.Introduction
	viewDeck.Description = deck.Description
	viewDeck.White = deck.White
	viewDeck.Red = deck.Red
	viewDeck.Blue = deck.Blue
	viewDeck.Green = deck.Green
	viewDeck.Black = deck.Black
	lrigs := make(map[string]int)
	addUnique(lrigs, deck.Lrig01)
	addUnique(lrigs, deck.Lrig02)
	addUnique(lrigs, deck.Lrig03)
	addUnique(lrigs, deck.Lrig04)
	addUnique(lrigs, deck.Lrig05)
	addUnique(lrigs, deck.Lrig06)
	addUnique(lrigs, deck.Lrig07)
	addUnique(lrigs, deck.Lrig08)
	addUnique(lrigs, deck.Lrig09)
	addUnique(lrigs, deck.Lrig10)
	mains := make(map[string]int)
	addUnique(mains, deck.Main01)
	addUnique(mains, deck.Main02)
	addUnique(mains, deck.Main03)
	addUnique(mains, deck.Main04)
	addUnique(mains, deck.Main05)
	addUnique(mains, deck.Main06)
	addUnique(mains, deck.Main07)
	addUnique(mains, deck.Main08)
	addUnique(mains, deck.Main09)
	addUnique(mains, deck.Main10)
	addUnique(mains, deck.Main11)
	addUnique(mains, deck.Main12)
	addUnique(mains, deck.Main13)
	addUnique(mains, deck.Main14)
	addUnique(mains, deck.Main15)
	addUnique(mains, deck.Main16)
	addUnique(mains, deck.Main17)
	addUnique(mains, deck.Main18)
	addUnique(mains, deck.Main19)
	addUnique(mains, deck.Main20)
	addUnique(mains, deck.Main21)
	addUnique(mains, deck.Main22)
	addUnique(mains, deck.Main23)
	addUnique(mains, deck.Main24)
	addUnique(mains, deck.Main25)
	addUnique(mains, deck.Main26)
	addUnique(mains, deck.Main27)
	addUnique(mains, deck.Main28)
	addUnique(mains, deck.Main29)
	addUnique(mains, deck.Main30)
	addUnique(mains, deck.Main31)
	addUnique(mains, deck.Main32)
	addUnique(mains, deck.Main33)
	addUnique(mains, deck.Main34)
	addUnique(mains, deck.Main35)
	addUnique(mains, deck.Main36)
	addUnique(mains, deck.Main37)
	addUnique(mains, deck.Main38)
	addUnique(mains, deck.Main39)
	addUnique(mains, deck.Main40)
	viewDeck.Lrig = list.New()
	for k := range lrigs {
		viewDeck.Lrig.PushBack(GetCardByKey(k, c))
	}
	viewDeck.Main = list.New()
	for k := range mains {
		viewDeck.Main.PushBack(GetCardByKey(k, c))
	}
	viewDeck.Scope = deck.Scope
	viewDeck.CreatedAt = deck.CreatedAt
	viewDeck.UpdatedAt = deck.UpdatedAt
	r.JSON(200, viewDeck)
}

func addUnique(m map[string]int, key string) {
	if key == "" {
		return
	}
	if m[key] == 0 {
		m[key] = 1
	} else {
		m[key]++
	}
}
