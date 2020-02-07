package main

import (
  "fmt"
  "strings"
  "math/rand"
  "sort"
)

// Arrays in Go are fixed, must declare size, cannot grow
func main () {
  scores := [4]int{4,3,2,1}

  // Iteration example
  for index, value := range scores {
    fmt.Printf("Index %d holds the integer %d\n", index, value)
  }

  // It is very uncommon to use arrays directly, instead we use slices
  // Slices don't require declaration of length
  newScores := []int{1,4,293,4,9}
  fmt.Println(newScores)

  // Create a slice with length 0 and capacity of 10, then add a value to a specific index
  newestScores := make([]int, 0, 10)
  newestScores = newestScores[0:6]
  newestScores[5] = 9033
  fmt.Println(newestScores)

  // Go slices alter orignal array, unlike Ruby or Python
  orgScores := []int{1,2,3,4,5}
  slice := orgScores[2:4]
  slice[0] = 999
  fmt.Println(orgScores)

  // Indexing and slicing strings
  testString := "This is my test string"

  // find the index of the first space
  fmt.Println(strings.Index(testString, " "))

  // No negative indexes in go
  newString := "Hi!, my name is Tylor"
  stringLength := len(newString)
  fmt.Printf("The newString: '%s', has a length of %d\n", newString, stringLength)

  // Subtracting from string length to find enpoint of slice
  fmt.Println(newString[:(stringLength - 5)] + "...\nSlim Shady!")

  // Copying slices

  finalScores := make([]int, 100)

  // creating random nums for each index in array
  for i := 0; i < 100; i++ {
    finalScores[i] = int(rand.Int31n(100))
  }
  // sort those random nums
  sort.Ints(finalScores)
  fmt.Println(finalScores)

  // copy original array, changes will not reflect old array
  worstScores := make([]int, 5)
  copy(worstScores, finalScores[:5])
  fmt.Println(worstScores)

  // maps are Go's: { Ruby: Hashes, Javascript: Objects, Python: Dictionaries }
  lookup := make(map[string]int)
  lookup["Tylor"] = 28
  fmt.Println(lookup)

  // Get number of keys using len
  total := len(lookup)
  fmt.Println(total)

  // Delete a value based on key
  delete(lookup, "Tylor")

  // Define an inital size for map to increase performance
  newMap := make(map[string]int, 100)
  fmt.Println(newMap)

  testPerson()
}

// Defining a map as a field of a struct
  type Person struct {
    Name string
    Friends map[string]*Person
  }

  func testPerson () {
    me := &Person{
      Name: "Tylor",
      Friends: make(map[string]*Person),
    }
    jon := Person{Name: "Jon"}
    me.Friends[jon.Name] = &jon
    fmt.Println(me)
  }
