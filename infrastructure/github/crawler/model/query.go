package model

import "time"

type Query struct {
	Repository struct {
		DatabaseID *int   `json:"databaseId"`
		Url        string `json:"url"`
		Issues     struct {
			Nodes    []*Issue `json:"nodes"`
			PageInfo struct {
				EndCursor   *string `json:"endCursor"`
				HasNextPage bool    `json:"hasNextPage"`
			} `json:"pageInfo"`
			TotalCount int `json:"totalCount"`
		} `json:"issues"`
		Refs struct {
			Nodes []*struct {
				Name *string `json:"name"`
			} `json:"nodes"`
			PageInfo struct {
				EndCursor   *string `json:"endCursor"`
				HasNextPage bool    `json:"hasNextPage"`
			} `json:"pageInfo"`
			TotalCount int `json:"totalCount"`
		} `json:"refs"`
		CreatedAt time.Time `json:"createdAt"`
	} `json:"repository"`
	RateLimit *struct {
		Limit     int       `json:"limit"`
		Cost      int       `json:"cost"`
		Remaining int       `json:"remaining"`
		ResetAt   time.Time `json:"reset_at"`
	} `json:"rateLimit"`
}

type Issue struct {
	DatabaseID *int `json:"databaseId"`
	Number     int  `json:"number"`
	Author     *struct {
		Login string `json:"login"`
	} `json:"author"`
	Closed    bool       `json:"closed"`
	ClosedAt  *time.Time `json:"closedAt"`
	CreatedAt time.Time  `json:"createdAt"`
	Labels    *struct {
		Nodes []*struct {
			Name string `json:"name"`
		} `json:"nodes"`
	} `json:"labels"`
	Assignees struct {
		Nodes []*struct {
			Login string `json:"login"`
			Email string `json:"email"`
		} `json:"nodes"`
	} `json:"assignees"`
	Title         string `json:"title"`
	Body          string `json:"body"`
	Url           string `json:"url"`
	TimelineItems struct {
		Nodes []*struct {
			Typename             *string `json:"__typename"`
			CrossReferencedEvent struct {
				Actor *struct {
					Login string `json:"login"`
				} `json:"actor"`
				CreatedAt time.Time `json:"createdAt"`
			}
			AssignedEvent struct {
				Actor *struct {
					Login string `json:"login"`
				} `json:"actor"`
				Assignee *struct {
					Login string `json:"login"`
					Email string `json:"email"`
				} `json:"assignee"`
				CreatedAt time.Time `json:"createdAt"`
			}
			UnassignedEvent struct {
				Actor *struct {
					Login string `json:"login"`
				} `json:"actor"`
				Assignee *struct {
					Login string `json:"login"`
					Email string `json:"email"`
				} `json:"assignee"`
				CreatedAt time.Time `json:"createdAt"`
			}
		} `json:"nodes"`
	} `json:"timelineItems"`
	Comments struct {
		Nodes []*struct {
			DatabaseID     *int   `json:"databaseId"`
			Body           string `json:"body"`
			ViewerCanReact bool   `json:"viewerCanReact"`
			Author         *struct {
				Login string `json:"login"`
			} `json:"author"`
		} `json:"nodes"`
		PageInfo struct {
			EndCursor   *string `json:"endCursor"`
			HasNextPage bool    `json:"hasNextPage"`
		} `json:"pageInfo"`
		TotalCount int `json:"totalCount"`
	} `json:"comments"`
}
