package servant

import (
	"appengine/datastore"
	"github.com/mrjones/oauth"
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

type User struct {
	Key   *datastore.Key
	Id    string
	Token *oauth.AccessToken
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

type Deck struct {
	Key          *datastore.key
	Owner        *datastore.key
	Title        string
	Introduction string
	Description  string
	White        bool
	Red          bool
	Blue         bool
	Green        bool
	Black        bool
	Lrig01       *datastore.key
	Lrig02       *datastore.key
	Lrig03       *datastore.key
	Lrig04       *datastore.key
	Lrig05       *datastore.key
	Lrig06       *datastore.key
	Lrig07       *datastore.key
	Lrig08       *datastore.key
	Lrig09       *datastore.key
	Lrig10       *datastore.key
	Main01       *datastore.key
	Main02       *datastore.key
	Main03       *datastore.key
	Main04       *datastore.key
	Main05       *datastore.key
	Main06       *datastore.key
	Main07       *datastore.key
	Main08       *datastore.key
	Main09       *datastore.key
	Main10       *datastore.key
	Main11       *datastore.key
	Main12       *datastore.key
	Main13       *datastore.key
	Main14       *datastore.key
	Main15       *datastore.key
	Main16       *datastore.key
	Main17       *datastore.key
	Main18       *datastore.key
	Main19       *datastore.key
	Main20       *datastore.key
	Main21       *datastore.key
	Main22       *datastore.key
	Main23       *datastore.key
	Main24       *datastore.key
	Main25       *datastore.key
	Main26       *datastore.key
	Main27       *datastore.key
	Main28       *datastore.key
	Main29       *datastore.key
	Main30       *datastore.key
	Main31       *datastore.key
	Main32       *datastore.key
	Main33       *datastore.key
	Main34       *datastore.key
	Main35       *datastore.key
	Main36       *datastore.key
	Main37       *datastore.key
	Main38       *datastore.key
	Main39       *datastore.key
	Main40       *datastore.key
	Scope        string
	Use0500      string
	Use1000      string
	Use1500      string
	Use2000      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
