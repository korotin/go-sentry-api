package sentry

import (
	"fmt"
	"time"
)

type ReleaseProject struct {
	ID            int      `json:"id,omitempty"`
	Slug          string   `json:"slug,omitempty"`
	Name          string   `json:"name,omitempty"`
	NewGroups     int      `json:"newGroups,omitempty"`
	Platform      string   `json:"platform,omitempty"`
	Platforms     []string `json:"platforms,omitempty"`
	HasHealthData bool     `json:"hasHealthData,omitempty"`
}

// Release is your release for a orgs teams project
type Release struct {
	ID                 int              `json:"id,omitempty"`
	Version            string           `json:"version,omitempty"`
	ShortVersion       string           `json:"shortVersion,omitempty"`
	VersionInfo        interface{}      `json:"versionInfo,omitempty"`
	Status             string           `json:"status,omitempty"`
	Ref                *string          `json:"ref,omitempty"`
	URL                *string          `json:"url,omitempty"`
	DateCreated        *time.Time       `json:"dateCreated,omitempty"`
	DateReleased       *time.Time       `json:"dateReleased,omitempty"`
	Data               interface{}      `json:"data,omitempty"`
	NewGroups          int              `json:"newGroups,omitempty"`
	Owner              interface{}      `json:"owner,omitempty"`
	CommitCount        int              `json:"commitCount,omitempty"`
	LastCommit         interface{}      `json:"lastCommit,omitempty"`
	DeployCount        int              `json:"deployCount,omitempty"`
	LastDeploy         interface{}      `json:"lastDeploy,omitempty"`
	Authors            []interface{}    `json:"authors,omitempty"`
	Projects           []ReleaseProject `json:"projects,omitempty"`
	FirstEvent         *time.Time       `json:"firstEvent,omitempty"`
	LastEvent          *time.Time       `json:"lastEvent,omitempty"`
	CurrentProjectMeta interface{}      `json:"currentProjectMeta,omitempty"`
	UserAgent          string           `json:"userAgent,omitempty"`
}

type Ref struct {
	Repository     string `json:"repository,omitempty"`
	Commit         string `json:"commit,omitempty"`         // HEAD commit SHA
	PreviousCommit string `json:"previousCommit,omitempty"` // commit SHA of previous release, optional
}

// NewRelease is used to create a new release
type NewRelease struct {
	// Required project slugs
	Projects []string `json:"projects,omitempty"`

	// Optional commit ref.
	Ref *string `json:"ref,omitempty"`

	// Optional URL to point to the online source code
	URL *string `json:"url,omitempty"`

	// Required for creating the release
	Version string `json:"version"`

	// Optional to set when it started
	DateStarted *time.Time `json:"dateStarted,omitempty"`

	// Optional to set when it was released to the public
	DateReleased *time.Time `json:"dateReleased,omitempty"`

	// An optional way to indicate the start and end commits for each repository included in a release
	Refs []Ref `json:"refs,omitempty"`
}

// GetRelease will fetch a release from your org and project this does need a version string
func (c *Client) GetRelease(o Organization, p Project, version string) (Release, error) {
	var rel Release
	err := c.do("GET", fmt.Sprintf("projects/%s/%s/releases/%s", *o.Slug, *p.Slug, version), &rel, nil)
	return rel, err
}

// GetReleases will fetch all releases from your org and project
func (c *Client) GetReleases(o Organization) ([]Release, *Link, error) {
	var rel []Release
	link, err := c.doWithPagination("GET", fmt.Sprintf("organizations/%s/releases", *o.Slug), &rel, nil)
	return rel, link, err
}

// CreateRelease will create a new release for a project in a org
func (c *Client) CreateRelease(o Organization, r NewRelease) (Release, error) {
	var rel Release
	err := c.do("POST", fmt.Sprintf("organizations/%s/releases", *o.Slug), &rel, &r)
	return rel, err
}

// UpdateRelease will update ref, url, started, released for a release.
// Version should not change.
func (c *Client) UpdateRelease(o Organization, r Release) error {
	return c.do("PUT", fmt.Sprintf("organizations/%s/releases/%s", *o.Slug, r.Version), &r, &r)
}

// DeleteRelease will delete the release from your project
func (c *Client) DeleteRelease(o Organization, r Release) error {
	return c.do("DELETE", fmt.Sprintf("organizations/%s/releases/%s", *o.Slug, r.Version), nil, nil)
}
