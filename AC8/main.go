package main

import (
  "fmt"
)

//Problema #1
func RevisaoDeContrato(){
  var res []string

  for{
    var erro rune
    var num string
    fmt.Scan(&erro)
    if erro == 0 {break}
    fmt.Scanln(&num)

    erro += '0'

    var tmp string

    for _, char := range num{
      if char != erro {
        tmp += string(char)
      }
    }

    i := 0

    for i < len(tmp) && tmp[i] == '0'{i++}

    if i != len(tmp) {
      tmp = tmp[i:]
      res = append(res, tmp)
    } else if len(tmp) == 0{
      res = append(res, "0")
    }else if tmp[0] == '0'{
      res = append(res, "0")
    }else{
      res = append(res, tmp)
    }
  }

  for _, i := range(res){fmt.Println(i)}
  
}

//Problema #2
func leds(){
  val := []int{6, 2, 5, 5, 4, 5, 6, 3, 7, 6}
  var Num int
  fmt.Scan(&Num)
  res := make([]int, Num)


  for i := 0; i<Num; i++{
    var N string

    fmt.Scanln(&N)

    for _, char := range N {
        res[i] += val[(int(char) - '0')]
    }
  }

  for i := 0; i<Num; i++ {fmt.Println(res[i], "leds")}
}

//Problema #3
func cesar(){
  letras := [26]byte{
      'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
      'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
  }

  var numCasos int
  fmt.Scan(&numCasos)

  inputStr := make([]string, numCasos)
  inputN := make([]int, numCasos)

  for i := 0; i< numCasos; i++{
    fmt.Scanln(&inputStr[i])
    fmt.Scan(&inputN[i])
  }

  for i := 0; i< numCasos; i++{
    S := inputStr[i]
    num := inputN[i]

    for _, char := range S {
      index := int(char) - num - 'A'
      index = (index + 26) % 26
      fmt.Printf("%c", letras[index])
    }

    fmt.Printf("\n")
  }
}

//Problema #4
func natal(){
  var pais []string
  
  for{
    var tmp string
    fmt.Scanln(&tmp)
    if tmp == "" {break}
    pais = append(pais, tmp)
  }

  for i := 0; i<len(pais); i++{
    switch pais[i] {
      case "brasil":
        fmt.Println("Feliz Natal!")
      case "alemanha":
        fmt.Println("Frohliche Weihnachten!")
      case "austria":
        fmt.Println("Frohe Weihnacht!")
      case "coreia":
        fmt.Println("Chuk Sung Tan!")
      case "espanha":
        fmt.Println("Feliz Navidad!")
      case "grecia":
        fmt.Println("Kala Christougena!")
      case "estados-unidos", "inglaterra", "australia", "canada", "antardida":
        fmt.Println("Merry Christmas!")
      case "portugal":
        fmt.Println("Feliz Natal!")
      case "suecia":
        fmt.Println("God Jul!")
      case "turquia":
        fmt.Println("Mutlu Noeller")
      case "argentina", "chile", "mexico":
        fmt.Println("Feliz Navidad!")
      case "irlanda":
        fmt.Println("Nollaig Shona Dhuit!")
      case "belgica":
        fmt.Println("Zalig Kerstfeest!")
      case "italia", "libia":
        fmt.Println("Buon Natale!")
      case "siria", "marrocos":
        fmt.Println("Milad Mubarak!")
      case "japao":
        fmt.Println("Merii Kurisumasu!")
      default:
        fmt.Println("--- NOT FOUND ---")
    }
  }
}
