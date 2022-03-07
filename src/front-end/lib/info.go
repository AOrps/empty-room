package lib

import "encoding/json"

type InfoBlock struct {
	Building string `json:"building"`
	Num      string `json:"num"`
	Abbrev   string `json:"abbrev"`
}

func jsonToHTML() []InfoBlock {
	var data []InfoBlock
	jsonString := `[{"building":"Campus Center","num":"12","abbrev":"CTR | CC"},{"building":"Central Avenue Building","num":"6","abbrev":"CAB"},{"building":"Central King Building","num":"7","abbrev":"CKB"},{"building":"Colton Hall","num":"5","abbrev":"COLT"},{"building":"Cullimore Hall","num":"10","abbrev":"CULM"},{"building":"Electrical and Computer Eng. Center","num":"16","abbrev":"ECEC"},{"building":"Faculty Memorial Hall","num":"15","abbrev":"FMH"},{"building":"Fenster Hall","num":"8","abbrev":"FENS"},{"building":"Guttenberg Information Technology Center","num":"31","abbrev":"GITC"},{"building":"Kupfrian Hall","num":"13","abbrev":"KUPF"},{"building":"Mechanical Engineering Center","num":"30","abbrev":"ME"},{"building":"Tiernan Hall","num":"26","abbrev":"TIER"},{"building":"Weston Hall","num":"4","abbrev":"WEST"}]`

	json.Unmarshal([]byte(jsonString), &data)

	return data
}
