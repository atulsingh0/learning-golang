package org

import (
	"fmt"
	"strings"
)

type Identifiable interface {
	ID() string
}

type Person struct {
	firstName string
	lastNmae  string
	TwitterHandler
}

type Handler struct {
	handle TwitterHandler
	name   string
}

type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {

	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

func NewPerson(firstName, lastNmae string) *Person {
	return &Person{
		firstName: firstName,
		lastNmae:  lastNmae,
	}
}

func (p *Person) ID() string {
	return "1234"
}

func (p *Person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastNmae)
}

func (p *Person) SetTwitter(handler TwitterHandler) error {

	if handler == "" {
		return fmt.Errorf("handle should not be empty")
	} else if !strings.HasPrefix(string(handler), "@") {
		return fmt.Errorf("handle should start with @")
	}
	p.TwitterHandler = handler
	return nil
}

func (p *Person) GetTwitter() TwitterHandler {
	return p.TwitterHandler
}
