package model

import "time"

type RateLimit struct {
	Limit     int       `json:"limit"`
	Cost      int       `json:"cost"`
	Remaining int       `json:"remaining"`
	ResetAt   time.Time `json:"reset_at"`
}

type PageInfo struct {
	EndCursor   *string `json:"endCursor"`
	HasNextPage bool    `json:"hasNextPage"`
}

type Query struct {
	Repository *Repository `json:"repository"`
	RateLimit  *RateLimit  `json:"rateLimit"`
}

type Repository struct {
	DatabaseID int              `json:"databaseId"`
	Url        string           `json:"url"`
	Name       string           `json:"name"`
	Issues     *IssueConnection `json:"issues"`
	Refs       *RefConnection   `json:"refs"`
	CreatedAt  time.Time        `json:"createdAt"`
}

type RefConnection struct {
	Nodes      []*Ref    `json:"nodes"`
	PageInfo   *PageInfo `json:"pageInfo"`
	TotalCount int       `json:"totalCount"`
}

type Ref struct {
	Name string `json:"name"`
}

type IssueConnection struct {
	Nodes      []*Issue  `json:"nodes"`
	PageInfo   *PageInfo `json:"pageInfo"`
	TotalCount int       `json:"totalCount"`
}

type Issue struct {
	DatabaseID int `json:"databaseId"`
	Number     int `json:"number"`
	Author     *struct {
		Login string `json:"login"`
	} `json:"author"`
	Closed    bool             `json:"closed"`
	ClosedAt  *time.Time       `json:"closedAt"`
	CreatedAt time.Time        `json:"createdAt"`
	Labels    *LabelConnection `json:"labels"`
	Assignees struct {
		Nodes []*struct {
			Login string  `json:"login"`
			Email *string `json:"email"`
		} `json:"nodes"`
	} `json:"assignees"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	Url           string `json:"url"`
	TimelineItems struct {
		Nodes []*struct {
			Typename *string `json:"__typename"`
			Actor    *struct {
				Login *string `json:"login"`
			} `json:"actor"`
			Assignee *struct {
				Login *string `json:"login"`
				Email *string `json:"email"`
			} `json:"assignee"`
			CreatedAt *time.Time `json:"createdAt"`
		} `json:"nodes"`
	} `json:"timelineItems"`
	Comments struct {
		Nodes      []*IssueComment `json:"nodes"`
		PageInfo   PageInfo        `json:"pageInfo"`
		TotalCount int             `json:"totalCount"`
	} `json:"comments"`
}

type LabelConnection struct {
	Nodes []*Label `json:"nodes"`
}

type Label struct {
	Name string `json:"name"`
}

type IssueComment struct {
	DatabaseID     int    `json:"databaseId"`
	Body           string `json:"body"`
	ViewerCanReact bool   `json:"viewerCanReact"`
	Author         *struct {
		Login string `json:"login"`
	} `json:"author"`
}
