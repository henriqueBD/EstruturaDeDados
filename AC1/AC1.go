package main

import (
	"fmt"
)

//Aluno: Henrique Barbosa Dantas Rolan
//Matricula: 202208818609

func calculaMedia(nums ...float64) float64{
  soma := 0.0
  for _, num := range nums {
    soma += num
  }
  return soma / float64(len(nums))
}

func verificaParidade(num int) bool{
  return num % 2 == 0
}

func minhaPotencia(base, expoente int) int{
  res := 1
  for i := 0; i < expoente; i++ {
    res *= base
  }
  return res
}

func converteCelsiusParaFahrenheit(celsius float64) float64{
  return celsius * 9/5 + 32
}

func main() {
  // Teste das funções
  
  fmt.Println(calculaMedia(0,128))
  fmt.Println(verificaParidade(64))
	fmt.Println(minhaPotencia(2,6))
  fmt.Printf("%.2f\n", converteCelsiusParaFahrenheit(17.7778))
}
