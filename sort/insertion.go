package main
  
import (
    "fmt"
    "math/rand"
)
 
 func imprimeVetor(vetor []int){ 
          var cont int   
  	  for cont = 0; cont < len(vetor); cont++ {
            	fmt.Print(vetor[cont]," ")
	  }
	  fmt.Println()
 }
 
func iniciaVetor(vetor []int, valorInicial int){ 
          var cont int   
  	  for cont = 0; cont < len(vetor); cont++ {
            	vetor[cont] = valorInicial
	  }
 }
 
 func iniciaVetorRandomico(vetor []int){ 
          var cont int   
          var valor int
  	  for cont = 0; cont < len(vetor); cont++ {
  	  	valor = rand.Intn(len(vetor))
            	vetor[cont] = valor
	  }
 }
 
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
	    troca(vetor, contInicio, indMenor )
	    imprimeVetor(vetor)
     }	    
}

  /*

    baralho = 5   
    para contInsercao = 1 até tamanho(baralho)
         anterior = contInsercao-1
         enquanto (baralho[contInsercao] <  baralho[anterior])
         	troca(baralho, contInsercao, anterior)
         	contInsercao = anterior
         fim-enquanto
    
    fim-para
            	
 
  */
 
 func insertionSort(vetor []int){ 
    var contInsercao int
    var anterior int
   // contInsercao = len(vetor) - 1
    for contInsercao = 1; contInsercao < len(vetor) ; contInsercao++ {
	    anterior = contInsercao - 1    
	    for (anterior > -1) && (vetor[contInsercao] < vetor[anterior]) {
	  	 troca(vetor, contInsercao, anterior)
	  	 contInsercao = anterior
	  	 anterior = contInsercao -1
	    }	     
	    imprimeVetor(vetor)
     }	    
}

 
 func iniciaVetorShuffle(vetor []int){ 
          var cont int   
          var indiceA int
          var indiceB int
          
  	  for cont = 0; cont < len(vetor); cont++ {
            	vetor[cont] = cont
	  }	
	  
	  for cont = 0; cont < len(vetor); cont++ {
            	indiceA = rand.Intn(len(vetor))
            	indiceB = rand.Intn(len(vetor))
            	troca(vetor, indiceA, indiceB)
	  }
  
 }
 
 
 
func main(){

	  var tamVet int
          var numerosTeste []int
        //  numerosTeste = []int{5,1,4,3,2,0}
          
          fmt.Print("Digite o tamanho do vetor:")
          fmt.Scanln(&tamVet)

          numerosTeste = make([]int, tamVet)  
          iniciaVetorShuffle(numerosTeste)
	  imprimeVetor(numerosTeste) 
	  insertionSort(numerosTeste)
          //imprimeVetor(numerosTeste)
}
