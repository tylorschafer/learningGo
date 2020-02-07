package main

import (
  "testing"
  "shopping"
)

func TestPriceCheck(t *testing.T) {
  price, _ := shopping.PriceCheck(1)

  if price != 9.001 {
    t.Error("Expected Price to equal 9.001")
  } else {
    t.Log("TestPriceCheck Passed")
  }
}
