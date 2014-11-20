package servant

type FormDeck struct {
	Title        string     `json:"Title"`
	Id           int        `json:"Id"`
	Introduction string     `json:"Introduction"`
	Description  string     `json:"Description"`
	White        bool       `json:"White"`
	Red          bool       `json:"Red"`
	Blue         bool       `json:"Blue"`
	Green        bool       `json:"Green"`
	Black        bool       `json:"Black"`
	Lrig         []FormCard `json:"lrig"`
	Main         []FormCard `json:"main"`
	Scope        string     `json:"Scope"`
}

type FormCard struct {
	Category      string `json:"Category"`
	Color         string `json:"Color"`
	Id            int    `json:"Id"`
	KeyName       string `json:"KeyName"`
	No            int    `json:"No"`
	ParentKeyName string `json:"ParentKeyName"`
	num           int    `json:"num"`
}
