package main

import (
  "fmt"
)
type Person struct {
  Name string
}

func (p *Person) Introduce() {
  fmt.Printf("Hi, I'm %s\n", p.Name)
}

type Saiyan struct {
  *Person
  Power int
  Father *Saiyan
}

func NewSaiyan(person *Person, power int, father *Saiyan) *Saiyan {
  return &Saiyan {
    Person: person,
    Power: power,
    Father: father,
  }
}

// *X means pointer to value of type X
func Super(s *Saiyan) {
  s.Power += 10000
}

func main() {
  // Using a pointer so we don't make changes to a copy, but the actual instance in memory
  // & operator gets the address of value
  goku := NewSaiyan(&Person{"Goku"}, 9000, nil)
  goku.Introduce()
  fmt.Printf("%s's power is %d!\n", goku.Name, goku.Power)
  Super(goku)
  fmt.Printf("%s has went Super!\n", goku.Name)
  fmt.Printf("%s's power is now %d!\n", goku.Name, goku.Power)
  gohan := NewSaiyan(&Person{"Gohan"}, 1000, goku)
  gohan.Introduce()
  fmt.Printf("%s's Father is %s\n", gohan.Name, gohan.Father.Name)
}
