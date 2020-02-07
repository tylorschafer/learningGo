package db

type Item struct {
  Price float64
}

// functions starting with Upper Case Letter are visable on imports. Lower case functions are not
func LoadItem(id int) *Item {
  return &Item{
    Price: 9.001,
  }
}
