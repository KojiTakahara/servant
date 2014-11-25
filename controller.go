package servant

import (
	"appengine"
	"appengine/datastore"
	"appengine/memcache"
	"container/list"
	"fmt"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/sessions"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func GetCardList(r render.Render, req *http.Request) {
	u, _ := url.Parse(req.URL.String())
	params := u.Query()

	c := appengine.NewContext(req)
	entities := make([]Card, 0, 10)
	memcache.Gob.Get(c, fmt.Sprintf("%s", params), &entities)
	if len(entities) == 0 {
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
		q = q.Filter("ParentKeyName=", "")
		keys, err := q.GetAll(c, &entities)
		if err != nil {
			c.Criticalf(err.Error())
			r.JSON(400, err)
			return
		}
		for i := range entities {
			entities[i].KeyName = keys[i].StringID()
		}
		mem_item := &memcache.Item{
			Key:    fmt.Sprintf("%s", params),
			Object: entities,
		}
		memcache.Gob.Add(c, mem_item)
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
}

func CreateModels(r render.Render, req *http.Request) {
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
	r.JSON(200, response)
}

func GetProductList(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)
	res := make([]Product, 0, 0)
	memcache.Gob.Get(c, "Product", &res)
	if len(res) == 0 {
		q := datastore.NewQuery("Product")
		if _, err := q.GetAll(c, &res); err != nil {
			c.Criticalf(err.Error())
			r.JSON(400, err)
		}
		mem_item := &memcache.Item{
			Key:    "Product",
			Object: res,
		}
		memcache.Gob.Add(c, mem_item)
	}
	r.JSON(200, res)
}

func GetIllustratorList(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)
	res := make([]Illustrator, 0, 0)
	memcache.Gob.Get(c, "Illustrator", &res)
	if len(res) == 0 {
		q := datastore.NewQuery("Illustrator")
		if _, err := q.GetAll(c, &res); err != nil {
			c.Criticalf(err.Error())
			r.JSON(400, err)
		}
		mem_item := &memcache.Item{
			Key:    "Illustrator",
			Object: res,
		}
		memcache.Gob.Add(c, mem_item)
	}
	r.JSON(200, res)
}

func GetConstraintList(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)
	res := make([]Constraint, 0, 0)
	memcache.Gob.Get(c, "Constraint", &res)
	if len(res) == 0 {
		q := datastore.NewQuery("Constraint")
		if _, err := q.GetAll(c, &res); err != nil {
			c.Criticalf(err.Error())
			r.JSON(400, err)
		}
		mem_item := &memcache.Item{
			Key:    "Constraint",
			Object: res,
		}
		memcache.Gob.Add(c, mem_item)
	}
	r.JSON(200, res)
}

func GetTypeList(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)
	res := make([]Type, 0, 0)
	memcache.Gob.Get(c, "Type", &res)
	if len(res) == 0 {
		q := datastore.NewQuery("Type")
		if _, err := q.GetAll(c, &res); err != nil {
			c.Criticalf(err.Error())
			r.JSON(400, err)
		}
		mem_item := &memcache.Item{
			Key:    "Type",
			Object: res,
		}
		memcache.Gob.Add(c, mem_item)
	}
	r.JSON(200, res)

}

func CreateDeck(r render.Render, req *http.Request, formDeck FormDeck, session sessions.Session) {
	c := appengine.NewContext(req)
	accessToken := GetAccessToken(session)
	if accessToken == nil {
		r.JSON(400, "") // TODO
	}
	user := GetUser(req, accessToken)

	deck := &Deck{}
	ids := make([]string, 0, 0)

	for _, lrig := range formDeck.UniqueLrigs {
		ids = append(ids, strconv.Itoa(lrig.Id))
	}
	for _, main := range formDeck.UniqueMains {
		ids = append(ids, strconv.Itoa(main.Id))
	}

	for _, lrig := range formDeck.Lrig {
		if lrig.Color == "white" {
			deck.White = true
		}
		if lrig.Color == "red" {
			deck.Red = true
		}
		if lrig.Color == "blue" {
			deck.Blue = true
		}
		if lrig.Color == "green" {
			deck.Green = true
		}
		if lrig.Color == "black" {
			deck.Black = true
		}
	}
	for _, main := range formDeck.Main {
		if main.Color == "white" {
			deck.White = true
		}
		if main.Color == "red" {
			deck.Red = true
		}
		if main.Color == "blue" {
			deck.Blue = true
		}
		if main.Color == "green" {
			deck.Green = true
		}
		if main.Color == "black" {
			deck.Black = true
		}
	}
	deck.Lrig01 = formDeck.Lrig[0].KeyName
	deck.Lrig02 = formDeck.Lrig[1].KeyName
	deck.Lrig03 = formDeck.Lrig[2].KeyName
	deck.Lrig04 = formDeck.Lrig[3].KeyName
	deck.Lrig05 = formDeck.Lrig[4].KeyName
	deck.Lrig06 = formDeck.Lrig[5].KeyName
	deck.Lrig07 = formDeck.Lrig[6].KeyName
	deck.Lrig08 = formDeck.Lrig[7].KeyName
	deck.Lrig09 = formDeck.Lrig[8].KeyName
	deck.Lrig10 = formDeck.Lrig[9].KeyName

	deck.Main01 = formDeck.Main[0].KeyName
	deck.Main02 = formDeck.Main[1].KeyName
	deck.Main03 = formDeck.Main[2].KeyName
	deck.Main04 = formDeck.Main[3].KeyName
	deck.Main05 = formDeck.Main[4].KeyName
	deck.Main06 = formDeck.Main[5].KeyName
	deck.Main07 = formDeck.Main[6].KeyName
	deck.Main08 = formDeck.Main[7].KeyName
	deck.Main09 = formDeck.Main[8].KeyName
	deck.Main10 = formDeck.Main[9].KeyName
	deck.Main11 = formDeck.Main[10].KeyName
	deck.Main12 = formDeck.Main[11].KeyName
	deck.Main13 = formDeck.Main[12].KeyName
	deck.Main14 = formDeck.Main[13].KeyName
	deck.Main15 = formDeck.Main[14].KeyName
	deck.Main16 = formDeck.Main[15].KeyName
	deck.Main17 = formDeck.Main[16].KeyName
	deck.Main18 = formDeck.Main[17].KeyName
	deck.Main19 = formDeck.Main[18].KeyName
	deck.Main20 = formDeck.Main[19].KeyName
	deck.Main21 = formDeck.Main[20].KeyName
	deck.Main22 = formDeck.Main[21].KeyName
	deck.Main23 = formDeck.Main[22].KeyName
	deck.Main24 = formDeck.Main[23].KeyName
	deck.Main25 = formDeck.Main[24].KeyName
	deck.Main26 = formDeck.Main[25].KeyName
	deck.Main27 = formDeck.Main[26].KeyName
	deck.Main28 = formDeck.Main[27].KeyName
	deck.Main29 = formDeck.Main[28].KeyName
	deck.Main30 = formDeck.Main[29].KeyName
	deck.Main31 = formDeck.Main[30].KeyName
	deck.Main32 = formDeck.Main[31].KeyName
	deck.Main33 = formDeck.Main[32].KeyName
	deck.Main34 = formDeck.Main[33].KeyName
	deck.Main35 = formDeck.Main[34].KeyName
	deck.Main36 = formDeck.Main[35].KeyName
	deck.Main37 = formDeck.Main[36].KeyName
	deck.Main38 = formDeck.Main[37].KeyName
	deck.Main39 = formDeck.Main[38].KeyName
	deck.Main40 = formDeck.Main[39].KeyName

	deck.Title = formDeck.Title
	deck.Introduction = formDeck.Introduction
	deck.Description = formDeck.Description
	deck.Scope = formDeck.Scope

	deck.Use0500 = CreateUseDeckStr(ids, 1, 500)
	deck.Use1000 = CreateUseDeckStr(ids, 501, 1000)
	deck.Use1500 = CreateUseDeckStr(ids, 1001, 1500)
	deck.Use2000 = CreateUseDeckStr(ids, 1501, 2000)

	deck.Owner = fmt.Sprintf("%v", user["screen_name"])
	deck.CreatedAt = time.Now()
	deck.UpdatedAt = time.Now()
	key := datastore.NewKey(c, "Deck", "", 0, nil)
	key, err := datastore.Put(c, key, deck)
	if err != nil {
		c.Criticalf("%s", err)
		r.JSON(400, err)
	} else {
		r.JSON(200, deck)
	}
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
