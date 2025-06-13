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
 
 func iniciaMatrizShuffle(mat [][]int, valorInicial int){ 
          var contI int
          var contJ int 
  	  for contI = 0; contI < len(mat); contI++ {
            	for contJ = 0; contJ < len(mat[0]); contJ++ {
            		mat[contI][contJ] = valorInicial
            		valorInicial = valorInicial + 1
            		
	  	}
	  }
	  for contI = 0; contI < len(mat)*len(mat); contI++ {
	  	troca(mat, rand.Intn(len(mat)), rand.Intn(len(mat[0])) , rand.Intn(len(mat)), rand.Intn(len(mat[0]))) 
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
 
 
 
func troca(mat [][]int, indAi int, indAj int, indBi int, indBj int){ 
	var temp int   
	temp = mat[indAi][indAj]
	mat[indAi][indAj] = mat[indBi][indBj]
	mat[indBi][indBj] = temp
          
}


func verificaQuadradaOrdem(mat [][]int) (bool, int) {
	var numLinhas int
	var numColunas int
	var ehQuadrada bool
	
	numLinhas = len(mat) 
	numColunas = len(mat[0])
	
	ehQuadrada = false
	if (numLinhas == numColunas){
		ehQuadrada = true
	}
	
	return ehQuadrada,numLinhas
}

func  calculaSinal(indiceL int, indiceC int) int{
	var sinal int

	sinal = -1
	if ((indiceL + indiceC)% 2) == 0 {
		sinal = 1
	}

	return sinal		
}

func copiaMatrizMaiorParaMenor(maior [][]int, menor [][]int, isqn int, jsqn int){
	var contAi,contAj,contBi,contBj,temp,numL,numC int;
	numL = len(menor) 
	numC = len(menor[0])

	contAi = 0
	for contBi = 0; contBi < numL; contBi++ {
		if(contAi == isqn){
			contAi++
		}
		contAj = 0
		for contBj = 0; contBj < numC; contBj++ {
			if(contAj == jsqn){
				contAj++ 
			}
			temp = maior[contAi][contAj]
			menor[contBi][contBj] = temp
			contAj++
		}
		contAi++
	}
}


/*func criaMatrizVazia(mat [][]int, numLinhas int, numColunas int){
        var cont int
	mat = make([][]int, numLinhas)
	for cont = 0; cont < numLinhas; cont++ {
	    mat[cont] = make([]int, numColunas)
	}
	fmt.Println(mat)
}
*/

func detOrdem1(mat [][]int) int{
	return mat[0][0]
}

func detOrdem2(mat [][]int) int{
	var diagonalP int
	var diagonalI int

	diagonalP = mat[0][0] * mat[1][1]		
	diagonalI = mat[1][0] * mat[0][1]		
	return (diagonalP - diagonalI)
}


func detOrdemN(mat [][]int) int{
	var sinal,cofator,detTemp,resposta,contL,contC,numL,numC,cont int
	var matMenor [][]int
	numL = len(mat) 
	numC = len(mat[0])
	
	resposta = 0;
	contL = 0;
	for contC = 0; contC < numC; contC++ {
		cofator = mat[contL][contC]
		sinal = calculaSinal(contL,contC)
		//criando a matriz menor
		matMenor = make([][]int, numL-1)
		for cont = 0; cont < (numL-1); cont++ {
    			matMenor[cont] = make([]int, numC-1)
		}
		
		copiaMatrizMaiorParaMenor(mat,matMenor,contL,contC)
		detTemp = determinante(matMenor);
		//fmt.Println("DetTemp ",detTemp)
		//fmt.Println("resposta ",resposta)
		//fmt.Println("cofator ",cofator)
		//fmt.Println("sinal ",sinal)
		resposta = resposta + (cofator * sinal * detTemp)
		//fmt.Println("resposta dps",resposta)
	}
	
	return resposta
}
	

func determinante(mat [][]int) int{
	var ordem int
	var ehQuadrada bool
	var det int

	ehQuadrada, ordem = verificaQuadradaOrdem(mat)
	det = 0
	if(ehQuadrada){		
		switch (ordem) {
		    case 1:
		        fmt.Println("Ordem 1")
		    	det = detOrdem1(mat)
		    case 2:
		    	fmt.Println("Ordem 2")
		    	det = detOrdem2(mat)
		    default: 
		        fmt.Println("Ordem ", ordem)
		    	det = detOrdemN(mat)
			
		}
		imprimeMatriz(mat)
		fmt.Println("Det ", det)
		
	} else {
		fmt.Println("Matriz nao eh quadrada!! retornando 0")
	}
	return det
}


func main(){

        var numLinhas int
        var numColunas int
       
        numLinhas = 5
        numColunas = 5
        
        var cont int
        
        var matrix [][]int
        matrix = make([][]int, numLinhas)
	for cont = 0; cont < numLinhas; cont++ {
	    matrix[cont] = make([]int, numColunas)
	}
	
        iniciaMatrizShuffle(matrix,1)
        imprimeMatriz(matrix)  
        fmt.Println(determinante(matrix))    
}
