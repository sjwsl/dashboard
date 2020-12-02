package model

type LabelConnection struct {
	Nodes      []*Label  `json:"nodes"`
	PageInfo   *PageInfo `json:"pageInfo"`
	TotalCount int       `json:"totalCount"`
}

type Label struct {
	Name string `json:"name"`
}
