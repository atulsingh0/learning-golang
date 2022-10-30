package org

import (
	"fmt"
	"strings"
)

type Identifiable interface {
	ID() string
}

type Name struct {
	firstName string
	lastNmae  string
}

type Person struct {
	Name
	TwitterHandler
	Identifiable
}

type Employee struct {
	Name
}

type sin string

func NewSIN(val string) Identifiable {
	return sin(val)
}

func (s sin) ID() string {
	return string(s)
}

func (s *sin) Country() string {
	return "Canada"
}

type eusin struct {
	id      string
	country string
}

func NewEUSIN(val, country string) Identifiable {
	return eusin{
		id:      val,
		country: country,
	}
}

func (e eusin) ID() string {
	return e.id
}

func (e *eusin) Country() string {
	return e.country
}

type TwitterHandler string

func (th TwitterHandler) RedirectUrl() string {

	cleanHandler := strings.TrimPrefix(string(th), "@")
	return fmt.Sprintf("https://www.twitter.com/%s", cleanHandler)
}

func NewPerson(firstName, lastNmae string, id Identifiable) Person {
	return Person{
		Name: Name{
			firstName: firstName,
			lastNmae:  lastNmae,
		},
		Identifiable: id,
	}
}

func (p *Person) ID() string {
	return fmt.Sprintf("%s", p.Identifiable.ID())
}

func (n *Name) FullName() string {
	return fmt.Sprintf("%s %s", n.firstName, n.lastNmae)
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
