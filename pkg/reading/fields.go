package reading

type Fields struct {
	IssueType *IssueType `json:"issueType"`
	Parent    *Issue     `json:"parent"`
	Status    *Status    `json:"status"`
}
