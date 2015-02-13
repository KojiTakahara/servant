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

func GetUser(r render.Render, params martini.Params, req *http.Request) {
	c := appengine.NewContext(req)
	key := datastore.NewKey(c, "User", params["userId"], 0, nil)
	var user User
	if err := datastore.Get(c, key, &user); err != nil {
		c.Criticalf(err.Error())
		r.JSON(400, "")
		return
	}
	r.JSON(200, user)
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
	user := GetTwitterUser(req, accessToken)

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
			deck.LrigWhite = true
		}
		if lrig.Color == "red" {
			deck.LrigRed = true
		}
		if lrig.Color == "blue" {
			deck.LrigBlue = true
		}
		if lrig.Color == "green" {
			deck.LrigGreen = true
		}
		if lrig.Color == "black" {
			deck.LrigBlack = true
		}
	}
	for _, main := range formDeck.OriginalMains {
		if main.Color == "white" {
			deck.MainWhite = true
		}
		if main.Color == "red" {
			deck.MainRed = true
		}
		if main.Color == "blue" {
			deck.MainBlue = true
		}
		if main.Color == "green" {
			deck.MainGreen = true
		}
		if main.Color == "black" {
			deck.MainBlack = true
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

	deck.UseCard = CreateUseDeckStr(ids, 1, 1500)

	deck.Owner = fmt.Sprintf("%v", user["screen_name"])
	deck.CreatedAt = time.Now()
	deck.UpdatedAt = time.Now()
	key := datastore.NewKey(c, "Deck", "", 0, nil)
	if formDeck.Id != 0 {
		key = datastore.NewKey(c, "Deck", "", int64(formDeck.Id), nil)
	}
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
	u, _ := url.Parse(req.URL.String())
	query := u.Query()
	q := datastore.NewQuery("Deck")
	q = q.Filter("Owner=", params["userId"])
	q = EqualQuery(q, query, "scope")
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

func GetPublicDeckList(r render.Render, req *http.Request) {
	c := appengine.NewContext(req)
	u, _ := url.Parse(req.URL.String())
	params := u.Query()
	q := datastore.NewQuery("Deck")
	q = q.Filter("Scope=", "PUBLIC")
	q = EqualQuery(q, params, "lrigWhite")
	q = EqualQuery(q, params, "lrigRed")
	q = EqualQuery(q, params, "lrigBlue")
	q = EqualQuery(q, params, "lrigGreen")
	q = EqualQuery(q, params, "lrigBlack")
	q = EqualQuery(q, params, "mainWhite")
	q = EqualQuery(q, params, "mainRed")
	q = EqualQuery(q, params, "mainBlue")
	q = EqualQuery(q, params, "mainGreen")
	q = EqualQuery(q, params, "mainBlack")
	q = q.Order("-CreatedAt")
	if params["cards"] == nil {
		q = q.Limit(ToInt(params["limit"][0]))
		q = q.Offset(ToInt(params["offset"][0]))
	}
	decks := make([]Deck, 0, 10)
	keys, err := q.GetAll(c, &decks)
	if err != nil {
		c.Criticalf(err.Error())
		r.JSON(400, err)
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
	viewDeck.LrigWhite = deck.LrigWhite
	viewDeck.LrigRed = deck.LrigRed
	viewDeck.LrigBlue = deck.LrigBlue
	viewDeck.LrigGreen = deck.LrigGreen
	viewDeck.LrigBlack = deck.LrigBlack
	viewDeck.MainWhite = deck.MainWhite
	viewDeck.MainRed = deck.MainRed
	viewDeck.MainBlue = deck.MainBlue
	viewDeck.MainGreen = deck.MainGreen
	viewDeck.MainBlack = deck.MainBlack
	lrigs := make(map[string]int)
	viewDeck.Lrig = addUnique(viewDeck.Lrig, lrigs, deck.Lrig01, c)
	viewDeck.Lrig = addUnique(viewDeck.Lrig, lrigs, deck.Lrig02, c)
	viewDeck.Lrig = addUnique(viewDeck.Lrig, lrigs, deck.Lrig03, c)
	viewDeck.Lrig = addUnique(viewDeck.Lrig, lrigs, deck.Lrig04, c)
	viewDeck.Lrig = addUnique(viewDeck.Lrig, lrigs, deck.Lrig05, c)
	viewDeck.Lrig = addUnique(viewDeck.Lrig, lrigs, deck.Lrig06, c)
	viewDeck.Lrig = addUnique(viewDeck.Lrig, lrigs, deck.Lrig07, c)
	viewDeck.Lrig = addUnique(viewDeck.Lrig, lrigs, deck.Lrig08, c)
	viewDeck.Lrig = addUnique(viewDeck.Lrig, lrigs, deck.Lrig09, c)
	viewDeck.Lrig = addUnique(viewDeck.Lrig, lrigs, deck.Lrig10, c)
	mains := make(map[string]int)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main01, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main02, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main03, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main04, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main05, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main06, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main07, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main08, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main09, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main10, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main11, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main12, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main13, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main14, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main15, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main16, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main17, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main18, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main19, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main20, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main21, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main22, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main23, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main24, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main25, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main26, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main27, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main28, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main29, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main30, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main31, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main32, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main33, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main34, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main35, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main36, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main37, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main38, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main39, c)
	viewDeck.Main = addUnique(viewDeck.Main, mains, deck.Main40, c)
	for k := range lrigs {
		for i := range viewDeck.Lrig {
			if viewDeck.Lrig[i].KeyName == k {
				viewDeck.Lrig[i].Num = lrigs[k]
			}
		}
	}
	for k := range mains {
		for i := range viewDeck.Main {
			if viewDeck.Main[i].KeyName == k {
				viewDeck.Main[i].Num = mains[k]
			}
		}
	}
	viewDeck.Scope = deck.Scope
	viewDeck.CreatedAt = deck.CreatedAt
	viewDeck.UpdatedAt = deck.UpdatedAt
	r.JSON(200, viewDeck)
}

func DeleteDeck(r render.Render, req *http.Request, params martini.Params, session sessions.Session) {
	c := appengine.NewContext(req)
	accessToken := GetAccessToken(session)
	if accessToken == nil {
		r.JSON(400, "") // TODO
		return
	}
	user := GetTwitterUser(req, accessToken)
	if params["owner"] != fmt.Sprintf("%v", user["screen_name"]) {
		r.JSON(400, "削除に失敗しました")
		return
	}
	key := datastore.NewKey(c, "Deck", "", int64(ToInt(params["id"])), nil)
	err := datastore.Delete(c, key)
	if err != nil {
		r.JSON(400, "削除に失敗しました")
		return
	}
	r.JSON(200, "成功")
}

/**
 *
 */
func addUnique(deck []Card, m map[string]int, key string, c appengine.Context) []Card {
	if addUniqueMap(m, key) == true {
		card, _ := GetCardByKey(key, c)
		deck = append(deck, card)
	}
	return deck
}

/**
 * mapに指定したkeyがなければ追加する
 * ある場合は、valueに +1 する
 * @param m マップ
 * @param key キー
 * @return 追加した場合はtrue, それ以外はfalseを返す.
 */
func addUniqueMap(m map[string]int, key string) bool {
	if key == "" {
		return false
	}
	if m[key] == 0 {
		m[key] = 1
		return true
	} else {
		m[key]++
		return false
	}
}
