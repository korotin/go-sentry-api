package sentry

import (
	"fmt"
	"net/http"
)

type IntegrationProvider struct {
	Key        string                 `json:"key,omitempty"`
	Slug       string                 `json:"slug,omitempty"`
	Name       string                 `json:"name,omitempty"`
	CanAdd     bool                   `json:"canAdd,omitempty"`
	CanDisable bool                   `json:"canDisable,omitempty"`
	Features   []string               `json:"features,omitempty"`
	Aspects    map[string]interface{} `json:"aspects,omitempty"`
}

type Integration struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Icon string `json:"icon,omitempty"`

	Provider IntegrationProvider `json:"provider,omitempty"`

	DomainName                    string                 `json:"domainName,omitempty"`
	AccountType                   interface{}            `json:"accountType,omitempty"` // actual type is unknown
	Scopes                        []string               `json:"scopes,omitempty"`
	Status                        string                 `json:"status,omitempty"`
	ConfigOrganization            []interface{}          `json:"configOrganization,omitempty"`
	ConfigData                    map[string]interface{} `json:"configData,omitempty"`
	ExternalID                    string                 `json:"externalId,omitempty"`
	OrganizationID                int                    `json:"organizationId,omitempty"`
	OrganizationIntegrationStatus string                 `json:"organizationIntegrationStatus,omitempty"`
	GracePeriodEnd                interface{}            `json:"gracePeriodEnd,omitempty"` // actual type is unknown
}

func (c *Client) GetIntegrations(o Organization) ([]Integration, error) {
	var integrations []Integration

	err := c.do(
		http.MethodGet,
		fmt.Sprintf("organizations/%s/integrations/", *o.Slug),
		&integrations,
		nil,
	)

	return integrations, err
}
