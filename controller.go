package servant

import (
	"appengine"
	"appengine/datastore"
	"github.com/martini-contrib/render"
	"log"
	"net/http"
	"net/url"
)

func GetCardList(r render.Render, req *http.Request) []Card {
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
		c.Criticalf(err.Error())
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
