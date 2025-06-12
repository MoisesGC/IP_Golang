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

func fibo(numero int) int{ 
 	var resposta int
	if numero == 0 || numero == 1 {
		resposta = 1
	} else {
		resposta = fibo(numero -1)  + fibo(numero - 2)
	}
	//fmt.Println(resposta)
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

func selectionSortRecursivo(indexInicio int, vetor []int) { 
        var cont int
        var indMenor int
        
	if(indexInicio < len(vetor) -1){
	    indMenor = indexInicio    
	    for cont = indMenor + 1; cont < len(vetor); cont++ {
	  	 if (vetor[indMenor] > vetor[cont]){
	  	 	indMenor = cont
	  	 }	 
	    }
	    troca(vetor, indexInicio, indMenor )
	    //fmt.Println(vetor)    
	    selectionSortRecursivo(indexInicio+1, vetor)
	    fmt.Println(vetor)
        }	    
 }


//func sortRecursivo() 
	//if (criterio de parada) {
		//faz algo sem chamar a funcao sortRecursivo
	//} else  { //passo recursivo
		//alguma coisa aqui chama de novo a funcao sortRecursivo
	//}
 //}

 
func troca(vetor []int, indA int, indB int){ 
          var temp int   
          temp = vetor[indA]
          vetor[indA] = vetor[indB]
          vetor[indB] = temp
          
}

func selectionSort(vetor []int){ 
    var cont int
    var contInicio int
    var indMenor int
    
    for contInicio = 0; contInicio < len(vetor) -1 ; contInicio++ {
	    indMenor = contInicio    
	    for cont = indMenor + 1; cont < len(vetor); cont++ {
	  	 if (vetor[indMenor] > vetor[cont]){
	  	 	indMenor = cont
	  	 }	 
	    }
	     
	   // fmt.Println(vetor[indMenor])
	  //  fmt.Println(indMenor)
	    troca(vetor, contInicio, indMenor )
     }	    
}

func main(){

        var numFatorial int
        var multi int
        var cont int

        var vetorX []int
        vetorX = []int{5}
        
        numFatorial = fatorial(1)
        fmt.Println(numFatorial)
        multi = multiplica(1,1)     
        fmt.Println(multi)     
        
        selectionSortRecursivo(0,vetorX)
        fmt.Println(vetorX)    
        
        for cont = 0; cont < 12; cont++ {
 		fmt.Print(fibo(cont), " ")
	}
	fmt.Println()
}

