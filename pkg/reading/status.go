package reading

type Status struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IconUrl     string `json:"iconUrl"`
}

type Statuses []Status
