package status

type Transition struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	To   To     `json:"to"`
}

type To struct {
	Name string `json:"name"`
}

type Transitions struct {
	Expand      string       `json:"expand"`
	Transitions []Transition `json:"transitions"`
}

var TransitionsSlice Transitions
