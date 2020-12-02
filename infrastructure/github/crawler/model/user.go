package model

type UserConnection struct {
	Nodes      []*User   `json:"nodes"`
	PageInfo   *PageInfo `json:"pageInfo"`
	TotalCount int       `json:"totalCount"`
}

type User struct {
	DatabaseID *int   `json:"databaseId"`
	Login      string `json:"login"`
	Email      string `json:"email"`
}

type Actor struct {
	Typename   *string `json:"__typename"`
	DatabaseID *int    `json:"databaseId"`
	Login      string  `json:"login"`
	Email      string  `json:"email"`
}
