package ui

import (
	"godaba/internal/db"
)

type Model struct {
	explorer db.Explorer
}

func NewModel(exp db.Explorer) Model {
	return Model{explorer: exp}
}
