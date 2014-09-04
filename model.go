package servant

import (
	"appengine/datastore"
)

type Card struct {
	Key           *datastore.Key
	Burst         string
	Bursted       bool
	Category      string
	Color         string
	Constraint    string
	CostBlack     int
	CostBlue      int
	CostColorless int
	CostGreen     int
	CostRed       int
	CostWhite     int
	Expansion     string
	Flavor        string
	Guard         string
	Illus         string
	Image         string
	Level         int
	Limit         int
	Name          string
	NameKana      string
	No            int
	Power         int
	Reality       string
	SearchText    string
	Text          string
	Type          string
	Url           string
}

type Product struct {
	Key  *datastore.Key
	Id   string
	Name string
}

type Type struct {
	Key  *datastore.Key
	Name string
}

type Illustrator struct {
	Key  *datastore.Key
	Name string
}

type Constraint struct {
	Key  *datastore.Key
	Type string
}
