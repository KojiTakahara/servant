package servant

import (
	"appengine"
	"appengine/datastore"
	"bytes"
	"net/http"
	"strconv"
)

/**
 * Nameに一致する、id以下のCard一覧を取得する
 */
func GetCardByNameAndId(name string, id int, req *http.Request) []Card {
	c := appengine.NewContext(req)
	q := datastore.NewQuery("Card")
	q = q.Filter("Name=", name)
	q = q.Filter("Id<", id)
	q = q.Order("Id")
	cards := make([]Card, 0, 10)
	keys, err := q.GetAll(c, &cards)
	if err != nil {
		c.Criticalf(err.Error())
	}
	for i := range cards {
		cards[i].KeyName = keys[i].StringID()
	}
	return cards
}

/**
 * デッキに使われているカードを判定する為の文字列を生成する
 * 含まれていれば 1, そうでなければ 0 をつける
 */
func CreateUseDeckStr(ids []string, start int, end int) string {
	var buffer bytes.Buffer
	for i := start; i < end; i++ {
		if Contains(strconv.Itoa(i), ids) {
			buffer.WriteString("1")
		} else {
			buffer.WriteString("0")
		}
	}
	return buffer.String()
}
