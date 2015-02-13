package servant

import (
	"time"
)

/**
 * カード
 */
type Card struct {
	Id            int
	KeyName       string
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
	Num           int
	Power         int
	Reality       string
	SearchText    string `datastore:",noindex"`
	Text          string `datastore:",noindex"`
	Type          string
	Url           string
	ParentKeyName string
}

/**
 * 登録ユーザ
 */
type User struct {
	Key   string
	Id    string
	Token string
}

/**
 * エキスパンション
 */
type Product struct {
	Id   string
	Name string
}

/**
 * 種族またはルリグ
 */
type Type struct {
	Name string
}

/**
 * イラストレーター
 */
type Illustrator struct {
	Name string
}

/**
 * 限定条件
 */
type Constraint struct {
	Type string
}

/**
 * デッキ
 */
type Deck struct {
	Id           int64
	Owner        string
	Title        string
	Introduction string `datastore:",noindex"`
	Description  string `datastore:",noindex"`
	LrigWhite    bool
	LrigRed      bool
	LrigBlue     bool
	LrigGreen    bool
	LrigBlack    bool
	MainWhite    bool
	MainRed      bool
	MainBlue     bool
	MainGreen    bool
	MainBlack    bool
	Lrig01       string
	Lrig02       string
	Lrig03       string
	Lrig04       string
	Lrig05       string
	Lrig06       string
	Lrig07       string
	Lrig08       string
	Lrig09       string
	Lrig10       string
	Main01       string
	Main02       string
	Main03       string
	Main04       string
	Main05       string
	Main06       string
	Main07       string
	Main08       string
	Main09       string
	Main10       string
	Main11       string
	Main12       string
	Main13       string
	Main14       string
	Main15       string
	Main16       string
	Main17       string
	Main18       string
	Main19       string
	Main20       string
	Main21       string
	Main22       string
	Main23       string
	Main24       string
	Main25       string
	Main26       string
	Main27       string
	Main28       string
	Main29       string
	Main30       string
	Main31       string
	Main32       string
	Main33       string
	Main34       string
	Main35       string
	Main36       string
	Main37       string
	Main38       string
	Main39       string
	Main40       string
	Scope        string
	UseCard      string `datastore:",noindex"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type SearchDeck struct {
	Id           int64
	Owner        string
	Title        string
	Introduction string `datastore:",noindex"`
	LrigWhite    bool
	LrigRed      bool
	LrigBlue     bool
	LrigGreen    bool
	LrigBlack    bool
	MainWhite    bool
	MainRed      bool
	MainBlue     bool
	MainGreen    bool
	MainBlack    bool
	Scope        string
	UpdatedAt    time.Time
}

type ViewDeck struct {
	Id           int64
	Owner        string
	Title        string
	Introduction string
	Description  string
	// LrigWhite    bool
	// LrigRed      bool
	// LrigBlue     bool
	// LrigGreen    bool
	// LrigBlack    bool
	// MainWhite    bool
	// MainRed      bool
	// MainBlue     bool
	// MainGreen    bool
	// MainBlack    bool
	Lrig      []Card
	Main      []Card
	Scope     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Amazon struct {
	Name    string
	Weight  int
	Html    string `datastore:",noindex"`
	Enabled bool
}
