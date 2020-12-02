package model

type IssueCommentConnection struct {
	Nodes      []*IssueComment `json:"nodes"`
	PageInfo   *PageInfo       `json:"pageInfo"`
	TotalCount int             `json:"totalCount"`
}

type IssueComment struct {
	DatabaseID int    `json:"databaseId"`
	Body       string `json:"body"`
	Author     *Actor `json:"author"`
}
