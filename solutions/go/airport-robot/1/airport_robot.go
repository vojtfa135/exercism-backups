package airportrobot

// Write your code here.
// This exercise does not have tests for each individual task.
// Try to solve all the tasks first before running the tests.

import "fmt"

type Greeter interface {
	LanguageName() string
	Greet(visitorName string) string
}

type Italian struct {
	langStr, greeting string
}

func NewItalian() *Italian {
	return &Italian{
		langStr:  "Italian",
		greeting: "Ciao",
	}
}

func (it Italian) LanguageName() string {
	return it.langStr
}

func (it Italian) Greet(visitorName string) string {
	return fmt.Sprintf("I can speak %s: %s %s!", it.langStr, it.greeting, visitorName)
}

type Portuguese struct {
	langStr, greeting string
}

func NewPortuguese() *Portuguese {
	return &Portuguese{
		langStr:  "Portuguese",
		greeting: "Olá",
	}
}

func (prt Portuguese) LanguageName() string {
	return prt.langStr
}

func (prt Portuguese) Greet(visitorName string) string {
	return fmt.Sprintf("I can speak %s: %s %s!", prt.langStr, prt.greeting, visitorName)
}

func SayHello(visitorName string, langGreeter Greeter) string {
	return langGreeter.Greet(visitorName)
}
