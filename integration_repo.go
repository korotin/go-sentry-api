package sentry

import (
	"fmt"
	"net/http"
)

type IntegrationRepo struct {
	ID   string `json:"identifier,omitempty"`
	Name string `json:"name,omitempty"`
	URL  string `json:"url,omitempty""`
}

func (c *Client) GetIntegrationRepos(o Organization) ([]IntegrationRepo, *Link, error) {
	var repos []IntegrationRepo

	link, err := c.doWithPagination(
		http.MethodGet,
		fmt.Sprintf("organizations/%s/repos/", *o.Slug),
		&repos,
		nil,
	)
	return repos, link, err
}

func (c *Client) AddRepoToIntegration(o Organization, i Integration, repoID string) (interface{}, error) {
	request := struct {
		Installation string `json:"installation,omitempty"`
		Identifier   string `json:"identifier,omitempty"`
		Provider     string `json:"provider,omitempty"`
	}{
		Installation: i.ID,
		Identifier:   repoID,
		Provider:     "integrations:" + i.Provider.Slug,
	}

	var response interface{}

	err := c.do(http.MethodPost, fmt.Sprintf("organizations/%s/repos/", *o.Slug), &response, request)

	return response, err
}
