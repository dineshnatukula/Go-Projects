package models

type Config struct {
	DMs                      map[string]Derivatives
	SportNames               map[string]struct{}
	RSKPlayerPropsTaxonomies map[string]struct{}
}

type Derivatives struct {
	Templates []string
}

type STL string
