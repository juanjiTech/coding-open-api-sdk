package team

import "github.com/juanjiTech/coding-open-api-sdk/config"

type Team struct {
	c *config.Basic
}

func New(c *config.Basic) Team {
	return Team{c}
}
