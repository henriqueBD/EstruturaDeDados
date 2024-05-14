package main

import (
	"fmt"
	"math"
)

//Problema #1
func alarme(){

	const dia = 1440
	
	var H1, M1, H2, M2 uint16
	var res_i, tmp1, tmp2 int

	res := make([]int, 0)

	for{
		fmt.Scan(&H1)
		fmt.Scan(&M1)
		fmt.Scan(&H2)
		fmt.Scan(&M2)

		if H1 == 0 && M1 == 0 && H2 == 0 && M2 == 0{
			break
		}

		tmp1 = int((H1 * 60) + M1)
		tmp2 = int((H2 * 60) + M2)

		res_i = tmp2 - tmp1

		if res_i<0 {
			res_i = (tmp2 + dia) - tmp1
		}

		res = append(res, res_i)
	}

	for i := range(res){
		fmt.Println(res[i])
	}
}

//Problema #2
func topN(){

	K := []uint8{1, 3, 5, 10, 25, 50, 100}
	
	var N uint8
	fmt.Scan(&N)
	N--

	for i := 0; i < 7; i++{
		if N < K[i]{
			fmt.Println("Top", K[i])
			break
		}
	}
}

//Problema #3
func escada(){
	var degraus, H, C, L uint16

	res := make([]float64, 0)
	var res_i float64

	for{
		_, err := fmt.Scan(&degraus)
		if err != nil { break }
		fmt.Scan(&H)
		fmt.Scan(&C)
		fmt.Scan(&L)

		res_i = math.Sqrt(float64(H*H) + float64(C*C))
		res_i *= float64(degraus)
		res_i *= float64(L)
		res_i /= 10000

		res = append(res, res_i)
	}

	for i := range(res){
		fmt.Printf("%.4f\n",res[i])
	}
}
