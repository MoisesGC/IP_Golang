package main
  
import (
    "fmt"
    "math/rand"
    "math"
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
 
 func somaComEscalar(valor int, vetor []int){ 
          var cont int   
  	  for cont = 0; cont < len(vetor); cont++ {
            	vetor[cont] = valor + vetor[cont]
	  }
 }
 
 func subtraiComEscalar(valor int, vetor []int){ 
          var cont int   
  	  for cont = 0; cont < len(vetor); cont++ {
            	vetor[cont] = vetor[cont] - valor
	  }
 }
 
 func multiplicaComEscalar(valor int, vetor []int){ 
          var cont int   
  	  for cont = 0; cont < len(vetor); cont++ {
            	vetor[cont] = valor * vetor[cont]
	  }
 }
 
 func produtoInterno(vetorA []int, vetorB []int) int { 
          var cont int  
          var acumulador int
          acumulador = 0
  	  for cont = 0; cont < len(vetorA); cont++ {
            	acumulador =  acumulador + (vetorA[cont]* vetorB[cont]) 
	  } 
	  return acumulador
 }
 
  func norma(vetor []int) float32 { 
          
          var resposta float32
          var prodInter int
          
          prodInter =  produtoInterno(vetor,vetor)
          resposta = float32(math.Sqrt(float64(prodInter)))
         
	  return resposta
 }
 
  
  func similaridadeVetores(vetorA []int, vetorB []int) float32 { 
          
          var resposta float32        
          resposta = float32(produtoInterno(vetorA, vetorB))/( norma(vetorA) * norma(vetorB))
    
	  return resposta
 }
 
 /*
    
   a torta de morango contem morango e doce de leite
   
   1 1 2 2 1 1 1 1 0 0
   
   a torta de chocolate contem chocolate e creme de leite
   
   1 1 2 0 1 1 0 1 2 1
 
 0 a
 1 torta
 2 de 
 3 morango
 4 contem
 5 e
 6 doce
 7 leite
 8 chocolate
 9 creme
 
 
 
 */
 
 
func main(){

	//  var tamVet int
	  var resposta int
       //   var vetorX []int
         // var vetorY []int
          
        //  fmt.Print("Digite o tamanho do vetor:")
       //   fmt.Scanln(&tamVet)

        //  vetorX = make([]int, tamVet)  
        //  vetorY = make([]int, tamVet)  
          
          
          vetorX := []int{1, 1, 2, 2, 1, 1, 1, 1, 0, 0}
          vetorY := []int{1, 1, 2, 0, 1, 1, 0, 1, 2, 1}
          consultaMorango := []int{0, 0, 0, 1, 0, 0, 0, 0, 0, 0}
          consultaChocolate := []int{0, 0, 0, 0, 0, 0, 0, 0, 1, 0}
          
          //iniciaVetorShuffle(vetorX)
         // iniciaVetorShuffle(vetorY)
          imprimeVetor(vetorX)
          imprimeVetor(vetorY)
         
          
          /*
          somaComEscalar(3, numerosTeste)
          imprimeVetor(numerosTeste)
          subtraiComEscalar(3, numerosTeste)
          imprimeVetor(numerosTeste)
          multiplicaComEscalar(4, numerosTeste)
          imprimeVetor(numerosTeste)
          */
          resposta = produtoInterno(vetorX, vetorX)
          fmt.Println(resposta)
          resposta = produtoInterno(vetorY, vetorY)
          fmt.Println(resposta)
          resposta = produtoInterno(vetorX, vetorY)
          fmt.Println(resposta)
          fmt.Println(norma(vetorX))
          fmt.Println(norma(vetorY))
          
          fmt.Println(similaridadeVetores(vetorX, vetorY))
          
          fmt.Println(similaridadeVetores(vetorX, consultaMorango))
          fmt.Println(similaridadeVetores(vetorY, consultaMorango))
          
          fmt.Println(similaridadeVetores(vetorX, consultaChocolate))
          fmt.Println(similaridadeVetores(vetorY, consultaChocolate))
          
}
