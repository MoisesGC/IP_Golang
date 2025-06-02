package main
  
import (
    "fmt"
    "math/rand"
)

 func imprimeMatriz(mat [][]int){ 
          var contI int
          var contJ int   
  	  for contI = 0; contI < len(mat); contI++ {
            	for contJ = 0; contJ < len(mat[0]); contJ++ {
            		fmt.Print(mat[contI][contJ]," ")
	  	}
	  	fmt.Println()
	  }
 }
 
func iniciaMatrix(mat [][]int, valor int){ 
          var contI int
          var contJ int   
  	  for contI = 0; contI < len(mat); contI++ {
            	for contJ = 0; contJ < len(mat[0]); contJ++ {
            		mat[contI][contJ] = valor
	  	}
	  }
 }
 
 func iniciaMatrixOrdenado(mat [][]int, valorInicial int){ 
          var contI int
          var contJ int 
  	  for contI = 0; contI < len(mat); contI++ {
            	for contJ = 0; contJ < len(mat[0]); contJ++ {
            		mat[contI][contJ] = valorInicial
            		valorInicial = valorInicial + 1
            		
	  	}
	  }
 }
 
 func somaComMatriz(mat1 [][]int, mat2 [][]int){ 
          var contI int
          var contJ int   
  	  for contI = 0; contI < len(mat1); contI++ {
            	for contJ = 0; contJ < len(mat1[0]); contJ++ {
            		mat1[contI][contJ] = mat1[contI][contJ] + mat2[contI][contJ]
	  	}
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


func main(){

	
        //slices em go
        var numLinhas int
        var numColunas int
       
        numLinhas = 3
        numColunas = 7
        
        var cont int
        
        matrix := make([][]int, numLinhas)
	for cont = 0; cont < numLinhas; cont++ {
	    matrix[cont] = make([]int, numColunas)
	}
         
       /*cubo := make([][][]int, numLinhas)
	for cont = 0; cont < numLinhas; cont++ {
	
	    cubo[cont] = make([][]int, numColunas)
	    for contS = 0; contS < numColunas; contS++ {
	    	cubo[cont][contS] = make([]int, numSlots)
	    }
	} 
	*/
	
        imprimeMatriz(matrix)  
        iniciaMatrixOrdenado(matrix,0)
        imprimeMatriz(matrix)       
}
