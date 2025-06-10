package main
  
import (
    "fmt"
)


 func fatorial(numero int) int{ 
 	var resposta int
	if numero == 0 || numero == 1 {
		resposta = 1
	} else {
		resposta = fatorial(numero -1) * numero
	}
	fmt.Println(resposta)
	return resposta
 }


 func multiplica(numero int, multiplicador int) int{ 
 	var resposta int
	if multiplicador == 1 {
		resposta = numero
	} else {
		resposta = multiplica(numero,multiplicador-1) +  numero
	}
	fmt.Println(resposta)
	return resposta
 }


func main(){

        var numFatorial int
        var multi int
   
        numFatorial = fatorial(5)
        fmt.Println(numFatorial)
        multi = multiplica(5,3)     
        fmt.Println(multi)     
}
