package servant

import (
	"appengine"
	"appengine/datastore"
	"appengine/memcache"
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
		return
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

	for _, lrig := range formDeck.OriginalLrigs {
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
	for _, main := range formDeck.OriginalMains {
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
	deck.Lrig01 = formDeck.OriginalLrigs[0].KeyName
	deck.Lrig02 = formDeck.OriginalLrigs[1].KeyName
	deck.Lrig03 = formDeck.OriginalLrigs[2].KeyName
	deck.Lrig04 = formDeck.OriginalLrigs[3].KeyName
	deck.Lrig05 = formDeck.OriginalLrigs[4].KeyName
	deck.Lrig06 = formDeck.OriginalLrigs[5].KeyName
	deck.Lrig07 = formDeck.OriginalLrigs[6].KeyName
	deck.Lrig08 = formDeck.OriginalLrigs[7].KeyName
	deck.Lrig09 = formDeck.OriginalLrigs[8].KeyName
	deck.Lrig10 = formDeck.OriginalLrigs[9].KeyName

	deck.Main01 = formDeck.OriginalMains[0].KeyName
	deck.Main02 = formDeck.OriginalMains[1].KeyName
	deck.Main03 = formDeck.OriginalMains[2].KeyName
	deck.Main04 = formDeck.OriginalMains[3].KeyName
	deck.Main05 = formDeck.OriginalMains[4].KeyName
	deck.Main06 = formDeck.OriginalMains[5].KeyName
	deck.Main07 = formDeck.OriginalMains[6].KeyName
	deck.Main08 = formDeck.OriginalMains[7].KeyName
	deck.Main09 = formDeck.OriginalMains[8].KeyName
	deck.Main10 = formDeck.OriginalMains[9].KeyName
	deck.Main11 = formDeck.OriginalMains[10].KeyName
	deck.Main12 = formDeck.OriginalMains[11].KeyName
	deck.Main13 = formDeck.OriginalMains[12].KeyName
	deck.Main14 = formDeck.OriginalMains[13].KeyName
	deck.Main15 = formDeck.OriginalMains[14].KeyName
	deck.Main16 = formDeck.OriginalMains[15].KeyName
	deck.Main17 = formDeck.OriginalMains[16].KeyName
	deck.Main18 = formDeck.OriginalMains[17].KeyName
	deck.Main19 = formDeck.OriginalMains[18].KeyName
	deck.Main20 = formDeck.OriginalMains[19].KeyName
	deck.Main21 = formDeck.OriginalMains[20].KeyName
	deck.Main22 = formDeck.OriginalMains[21].KeyName
	deck.Main23 = formDeck.OriginalMains[22].KeyName
	deck.Main24 = formDeck.OriginalMains[23].KeyName
	deck.Main25 = formDeck.OriginalMains[24].KeyName
	deck.Main26 = formDeck.OriginalMains[25].KeyName
	deck.Main27 = formDeck.OriginalMains[26].KeyName
	deck.Main28 = formDeck.OriginalMains[27].KeyName
	deck.Main29 = formDeck.OriginalMains[28].KeyName
	deck.Main30 = formDeck.OriginalMains[29].KeyName
	deck.Main31 = formDeck.OriginalMains[30].KeyName
	deck.Main32 = formDeck.OriginalMains[31].KeyName
	deck.Main33 = formDeck.OriginalMains[32].KeyName
	deck.Main34 = formDeck.OriginalMains[33].KeyName
	deck.Main35 = formDeck.OriginalMains[34].KeyName
	deck.Main36 = formDeck.OriginalMains[35].KeyName
	deck.Main37 = formDeck.OriginalMains[36].KeyName
	deck.Main38 = formDeck.OriginalMains[37].KeyName
	deck.Main39 = formDeck.OriginalMains[38].KeyName
	deck.Main40 = formDeck.OriginalMains[39].KeyName

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
	card, _ := GetCardByKey(keyStr, c)
	r.JSON(200, card)
}

func GetUserDeckList(r render.Render, params martini.Params, req *http.Request) {
	c := appengine.NewContext(req)
	q := datastore.NewQuery("Deck")
	q = q.Filter("Owner=", params["userId"])
	decks := make([]Deck, 0, 10)
	keys, err := q.GetAll(c, &decks)
	if err != nil {
		c.Criticalf(err.Error())
		r.JSON(200, err)
		return
	}
	for i := range decks {
		decks[i].Id = keys[i].IntID()
	}
	r.JSON(200, decks)
}

func GetPublicDeckList(r render.Render, params martini.Params, req *http.Request) {
	c := appengine.NewContext(req)
	q := datastore.NewQuery("Deck")
	q = q.Filter("Scope=", "PUBLIC")
	q = q.Limit(ToInt(params["limit"]))
	q = q.Offset(ToInt(params["offset"]))
	decks := make([]Deck, 0, 10)
	keys, err := q.GetAll(c, &decks)
	if err != nil {
		c.Criticalf(err.Error())
		r.JSON(200, err)
		return
	}
	for i := range decks {
		decks[i].Id = keys[i].IntID()
	}
	r.JSON(200, decks)
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
	viewDeck.Id = int64(id)
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
	for k := range lrigs {
		card, _ := GetCardByKey(k, c)
		card.Num = lrigs[k]
		viewDeck.Lrig = append(viewDeck.Lrig, card)
	}
	for k := range mains {
		card, _ := GetCardByKey(k, c)
		card.Num = mains[k]
		viewDeck.Main = append(viewDeck.Main, card)
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
