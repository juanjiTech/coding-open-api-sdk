package openAPI

import (
	"github.com/juanjiTech/coding-open-api-sdk/config"
	"github.com/juanjiTech/coding-open-api-sdk/openAPI/team"
)

type OpenAPI struct {
	c    *config.Basic
	Team team.Team
}

func New(c *config.Basic) OpenAPI {
	return OpenAPI{c,
		team.New(c)}
}
