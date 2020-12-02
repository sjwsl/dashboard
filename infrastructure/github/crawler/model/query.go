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
