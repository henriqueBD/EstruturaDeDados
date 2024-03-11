package main

import (
	"fmt"
)

func main() {

  //teste dos algoritimos
  
  arr := []int{1,2,6,8,54,80}
  buscaArr(arr, 10)

  hanoi(4, 1, 3)
  
}

func buscaArr(arr []int, target int) bool{
  n := len(arr)
  i := 0
  j := n-1

  for j>i{
    tmp := arr[i] + arr[j]
    if(tmp == target){
      fmt.Print(arr[i], arr[j])
      return true
    } else if tmp > target{
      j--
    } else {
      i++
    }
  }
  
  return false
}

func hanoi(n, inicio, final int){
  if(inicio < 1 || inicio > 3) { return }
  if(final < 1  || final > 3)  { return }

  resolver(n, inicio, final)
}

func resolver(n, inicio, final int){
  trabalho := 6 - (inicio + final)

  if(n == 1){
    mover(inicio, final)
  }else{
    resolver(n-1, inicio, trabalho)
    mover(inicio, final)
    resolver(n-1, trabalho, final)
  }
  
}

func mover(inicio, final int){
  fmt.Println("Mover disco da torre", inicio, "para", final)
}
