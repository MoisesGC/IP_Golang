package main
  
import (
    "fmt"
    "math/rand"
)
  
 
/*
func ehPrimoV2(numero int) (bool, error) {
	var cont int
	var ePrimo bool
	var primoNeg error
	
	primoNeg = nil
	ePrimo = false
	if (numero > 0) {
		if (numero >= 2) {
			ePrimo = true
			cont = 2
			for  (cont < numero) && ePrimo {
				if (numero % cont)  == 0 {
					ePrimo = false
				}
				cont++
			}
		}
	} else {
		primoNeg = errors.New("Nao se calcula primo negativo")
	}
	return ePrimo, primoNeg
}   
 */
 
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
	     
	   // fmt.Println(vetor[indMenor])
	  //  fmt.Println(indMenor)
	    troca(vetor, contInicio, indMenor )
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
          
          fmt.Print("Digite o tamanho do vetor:")
          fmt.Scanln(&tamVet)

          numerosTeste = make([]int, tamVet)  
          iniciaVetorShuffle(numerosTeste)
	  imprimeVetor(numerosTeste)
	  selectionSort(numerosTeste)
          imprimeVetor(numerosTeste)
}
