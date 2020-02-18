package status

type DoTransitionObject struct {
	Transition InnerTransition `json:"transition"`
}

type InnerTransition struct {
	Id string `json:"id"`
}
