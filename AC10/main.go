package main

import (
	"fmt"
  "math"
)

func figurinhas(){
  var NumCasos int
  fmt.Scan(&NumCasos)

  res := make([]int, NumCasos)
  for i := range res {
    res[i] = 1

    var num1 int
    var num2 int
    fmt.Scan(&num1)
    fmt.Scan(&num2)

    //Algoritmo de Euclides para MDC
    for num2 != 0 {
        num1, num2 = num2, num1%num2
    }

    res[i] = num1

  }

  for i := range res{
    fmt.Println(res[i])
  }
}

func dama(){
  res := make([]int, 0)

  var x1, y1, x2, y2 int

  for{
    fmt.Scan(&x1)
    if x1 == 0 {break}
    fmt.Scan(&y1)
    fmt.Scan(&x2)
    fmt.Scan(&y2)

    if x1 == x2{
      if y1 == y2{
        res = append(res, 0)
        continue
      }
      res = append(res, 1)
      continue
    }

    if y1 == y2{
      res = append(res, 1)
      continue
    }

    dx := x1 - x2
    dy := y1 - y2

    if dx < 0 {dx *= -1}
    if dy < 0 {dy *= -1}

    if dx == dy{
      res = append(res, 1)
      continue
    }

    res = append(res, 2)
  }

  for _, i := range(res){fmt.Println(i)}
}

func somaDeFatoriais(){
  res := make([]uint64, 0)

  var M, N int8

  for {
    _, err := fmt.Scan(&M)
    if err != nil {break}
    fmt.Scan(&N)
    
    num := fatorial(M) + fatorial(N)
    res = append(res, num)
  }
  
  for _, i := range(res){fmt.Println(i)}
}

func fatorial(val int8) uint64{
  val++
  res := uint64(1)
  for i := int8(2); i < val; i++{
    res *= uint64(i)
  }
  return res
}

func blob(){
  var numCasos int
  fmt.Scan(&numCasos)

  res := make([]int, numCasos)

  for i := 0; i < numCasos; i++{

    var C float32
    fmt.Scan(&C)

    for C > 1{
      C /= 2
      res[i]++
    }
    
  }

  for _, i := range(res){fmt.Println(i, "dias")}
}

func binarySearch(arr []int, val int) (int, bool){
  low, high := 0, len(arr)-1
  var mid int

  for low <= high {
    mid = low + (high-low)/2
    if arr[mid] == val {
      return mid, true
    }
    if arr[mid] > val {
      high = mid - 1
    } else {
      low = mid + 1
    }
  }
  
  return mid, false
}

func frequenciaDeNumeros(){
  var numCasos int
  fmt.Scan(&numCasos)

  numeros := make([]int, 1, numCasos)
  numerosFreq := make([]int, numCasos)

  if numCasos == 0 {return}

  fmt.Scan(&numeros[0])
  numerosFreq[0] = 1

  if numCasos == 1 {return}

  var tmp int
  fmt.Scan(&tmp)

  if tmp == numeros[0]{
    numerosFreq[0] = 2
  }else{
    if tmp > numeros[0]{
      numeros = append(numeros, tmp)
      numerosFreq[1] = 1
    }else{
      numeros = append(numeros, numeros[0])
      numeros[0] = tmp
      numerosFreq[1] = 1
    }
  }

  numCasos -= 2
  
  for i := 0; i < numCasos; i++{
    var num int
    fmt.Scan(&num)
    i, existe := binarySearch(numeros, num)
    if existe{
      numerosFreq[i]++
    }else{
      numeros = append(numeros, 0)
      copy(numeros[i+1:], numeros[i:])
      copy(numerosFreq[i+1:], numerosFreq[i:])
      numeros[i] = num
      numerosFreq[i] = 1
    }
  }

  for i := range(numeros){
    fmt.Println(numeros[i], "aparece", numerosFreq[i], "vez(es)")
  }
  
}

func primoRapido(){

  var num, numCasos int
  fmt.Scan(&numCasos)
  
  res := make([]bool, numCasos)

  for n := 0; n < numCasos; n++{
    
    fmt.Scan(&num)

    val := int(math.Sqrt(float64(num)))
    val++

    res[n] = true

    for i := 2; i < val; i++{
      if num % i == 0 {
        res[n] = false
        break
      }
    }
    
  }

  for i := range(res){
    if res[i]{
      fmt.Println("Prime")
    }else{
      fmt.Println("Not Prime")
    }
  }
  
}
