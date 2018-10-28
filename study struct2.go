package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToLower(p.lastName)
}

func main() {

	pesr1 := &Person{firstName: "asdf", lastName: "adsdddd"}

	upPerson(pesr1)

	fmt.Printf(pesr1.firstName, pesr1.lastName, pesr1)
}
