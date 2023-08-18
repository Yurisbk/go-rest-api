package models

type Pesonalidade struct {
	Nome     string `json:"nome"`
	Historia string `json:"historia"`
}

var Personalidades []Pesonalidade
