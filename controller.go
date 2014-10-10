package servant

import (
	"appengine"
	"appengine/datastore"
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
	c.Infof("%s", keyStr)
	key := datastore.NewKey(c, "Card", keyStr, 0, nil)
	var card Card
	if err := datastore.Get(c, key, &card); err != nil {
		c.Criticalf(err.Error())
		r.JSON(400, "")
	} else {
		r.JSON(200, card)
	}
}

func GetDeck(r render.Render, params martini.Params, req *http.Request) {
	c := appengine.NewContext(req)
	id, _ := strconv.Atoi(params["id"])
	key := datastore.NewKey(c, "Deck", "", int64(id), nil)
	var e2 Deck
	if err := datastore.Get(c, key, &e2); err != nil {
		c.Criticalf(err.Error())
	}
	r.JSON(200, e2)
}
