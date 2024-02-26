package main

import (
	"fmt"
)

type Contato struct {
  Nome string
  Email string
}

func (c *Contato) AlterarEmail(novoEmail string) {
  c.Email = novoEmail
}

func adicionarContato(arr *[5]Contato, contato Contato){
  for i:=0; i<5; i++{
    if(arr[i] == Contato{}){
      arr[i] = contato
      return
    }
  }
  fmt.Println("Lista de contatos cheia")
}

func removerContato(arr *[5]Contato){
  for i:=4; i>-1; i--{
    if!(arr[i] == Contato{}){
      arr[i] = Contato{}
      return
    }
  }
  fmt.Println("Lista de contatos vazia")
}

func main() {

  var lista [5]Contato

  for {
    fmt.Print("Digite 1 para adicionar um contato, 2 para excluir e 3 para sair: ")

    var res int
    fmt.Scan(&res)
    switch res{
      case 1:
        var nome string
        var email string
        print("Informe o nome: ")
        fmt.Scan(&nome)
        print("Informe o email: ")
        fmt.Scan(&email)
        adicionarContato(&lista, Contato{nome, email})
      case 2:
        removerContato(&lista)
      case 3:
        return
    }

    fmt.Println(lista)
  }
  
}
