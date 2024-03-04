package main

import (
	"fmt"
	"geometria"
	"math"
)

type Ponto struct {
  X float64
  Y float64
}

func DistanciaOrigem(ponto Ponto) float64{
  return math.Sqrt(math.Pow(ponto.X, 2) + math.Pow(ponto.Y, 2))
}

func inverterString(input string) string{
  size := int(len(input))
  res := []rune(input)
  for i:=0; i<size; i++{
    res[i] = rune(input[size-i-1])
  }
  return string(res)
}

func main() {

  var arr [10]int

  for i:=0; i<10; i++ {
    arr[i] = i
  }

  for i:=0; i<10; i++ {
    fmt.Println(arr[i])
  }

  fmt.Println(inverterString("Ola"))
  
}
