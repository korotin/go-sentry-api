package sentry

import (
	"fmt"
	"time"
)

type Deploy struct {
	ID           string     `json:"id,omitempty"`
	Environment  string     `json:"environment"`
	Projects     []string   `json:"projects,omitempty"`
	Status       string     `json:"status,omitempty"`
	Name         *string    `json:"name,omitempty"`
	URL          *string    `json:"url,omitempty"`
	DateStarted  *time.Time `json:"dateStarted,omitempty"`
	DateFinished *time.Time `json:"dateFinished,omitempty"`
}

// CreateDeploy creates deploy for given releases
func (c *Client) CreateDeploy(o Organization, r Release, d Deploy) (Deploy, error) {
	var deploy Deploy
	err := c.do("POST", fmt.Sprintf("organizations/%s/releases/%s/deploys", *o.Slug, r.Version), &deploy, &d)
	return deploy, err
}
