package reading

type Issue struct {
	Id     string `json:"id"`
	Key    string `json:"key"`
	Fields Fields `json:"fields"`
	Names  Names  `json:"names"`
}
