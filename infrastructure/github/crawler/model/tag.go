package model

type RefConnection struct {
	Nodes      []*Ref    `json:"nodes"`
	PageInfo   *PageInfo `json:"pageInfo"`
	TotalCount int       `json:"totalCount"`
}

type Ref struct {
	Name string `json:"name"`
}
