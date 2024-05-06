package main

import (
	"fmt"
)

//Problema #1
func distancia(){
	var val int
	fmt.Scan(&val)
	fmt.Println(val*2,"minutos")
}

//Problema #2
type Produto struct {
	Nome  string
	Preco float32
}

func feira(){	

	var numCasos int
	fmt.Scan(&numCasos)

	res := make([]float32, numCasos)

	for i := 0; i<numCasos; i++{
		
		var numProdutos int
		fmt.Scan(&numProdutos)
		produto := make([]Produto, numProdutos)
		
		for j := 0; j<numProdutos; j++{
			var nome string
			var preco float32
			fmt.Scan(&nome)
			fmt.Scan(&preco)
			produto[j] = Produto{Nome: nome, Preco: preco}
		}

		produtoMap := make(map[string]Produto)

		for j := range(produto){
			produtoMap[produto[j].Nome] = produto[j]
		}

		var numCompras int
		fmt.Scan(&numCompras)

		res[i] = 0

		for j := 0; j<numCompras; j++{
			var compra string
			var quantidade int
			fmt.Scan(&compra)
			fmt.Scan(&quantidade)
			if tmp, ok := produtoMap[compra]; ok {
				res[i] += tmp.Preco * float32(quantidade)
			}
		}
		
	}

	for i := range(res){
		fmt.Printf("R$ %.2f\n", res[i])
	}
}

//Problema #3
func avioesDePapel(){
	var competidores, folhasTot, folhasInd int
	fmt.Scan(&competidores)
	fmt.Scan(&folhasTot)
	fmt.Scan(&folhasInd)
	if folhasTot >= folhasInd * competidores {
		fmt.Println("S")
	}else{
		fmt.Println("N")	
	}
}

//Problema #4
func frequencia(){

	var numDigitos int
	fmt.Scan(&numDigitos)
	seq := make([]int8, numDigitos)

	for i := 0; i<numDigitos; i++{
		fmt.Scan(&seq[i])
	}
	
	res := 1

	curr := seq[0]

	for i := 1; i<numDigitos; i++{
		if seq[i] != curr{
			res++
			curr = seq[i]
		}
	}

	fmt.Println(res)
	
}
