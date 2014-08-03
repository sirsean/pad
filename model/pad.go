package model

import (
	"strings"
	"code.google.com/p/go-uuid/uuid"
)

type Pad struct {
	Id string
	Consumable int
	ExpiresInSeconds int
	Callback string
}

func (p *Pad) GenerateId() {
	p.Id = uuid.New()
}

/*
A pad cannot expire in the past, and cannot expire more than
an hour into the future
*/
func (p *Pad) EnsureExpiration() {
	if p.ExpiresInSeconds < 0 {
		p.ExpiresInSeconds = 0
	} else if p.ExpiresInSeconds > 3600 {
		p.ExpiresInSeconds = 3600
	}
}

func (p *Pad) Use() {
	p.Consumable -= 1
}

func (p *Pad) IsConsumed() bool {
	return p.Consumable == 0
}

func (p *Pad) CallbackUrl() string {
	return strings.Replace(p.Callback, "{pad}", p.Id, -1)
}
