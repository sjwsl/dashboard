package model

import "time"

type Repository struct {
	DatabaseID      int              `json:"databaseId"`
	Url             string           `json:"url"`
	Name            string           `json:"name"`
	AssignableUsers *UserConnection  `json:"assignableUsers"`
	Issues          *IssueConnection `json:"issues"`
	Labels          *LabelConnection `json:"labels"`
	Refs            *RefConnection   `json:"refs"`
	CreatedAt       time.Time        `json:"createdAt"`
}
