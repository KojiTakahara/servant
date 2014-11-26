package servant

import (
	"appengine"
	"appengine/datastore"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

func GetCardByKey(keyStr string, c appengine.Context) (card Card, err error) {
	k := datastore.NewKey(c, "Card", keyStr, 0, nil)
	e := Card{}
	return e, datastore.Get(c, k, &e)
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
