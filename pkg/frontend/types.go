package frontend

type InfoBlock struct {
	Building string `json:"building"`
	Num      string `json:"num"`
	Abbrev   string `json:"abbrev"`
}

type Room struct {
	Room  string
	Times []string
}
