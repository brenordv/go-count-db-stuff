package models

import (
	"fmt"
	"strings"
)

type CountQuery struct {
	Name          string `json:"queryName"`
	Sql           string `json:"sql"`
	ExpectedCount int    `json:"expectedCount"`
}

type RuntimeConfig struct {
	ConnectionString string       `json:"connectionString"`
	Queries          []CountQuery `json:"queries"`
}

func (c CountQuery) String() string {
	return fmt.Sprintf(" - Name: '%s'\n - Sql: %s\n - Expected Count Result: %d\n", c.Name, c.Sql, c.ExpectedCount)
}

func (r RuntimeConfig) String() string {
	var lines []string

	lines = append(lines, fmt.Sprintf("Connection string: %s\n", r.ConnectionString))
	lines = append(lines, fmt.Sprintf("Queries (%d):\n", len(r.Queries)))
	for _, query := range r.Queries {
		lines = append(lines, query.String())
	}

	return strings.Join(lines, "")
}