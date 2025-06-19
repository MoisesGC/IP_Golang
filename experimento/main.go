package main
  
import (
    "fmt"
    "time"
)

func imprimeMatrizBaseline(mat [][]int){ 
          var contI int
          var contJ int   
  	  for contI = 0; contI < len(mat); contI++ {
            	for contJ = 0; contJ < len(mat[0]); contJ++ {
            		//fmt.Print(mat[contI][contJ]," ")
            		fmt.Print("")
	  	}
	  	fmt.Println()
	  }
}

func imprimeMatrizHacked(mat [][]int){ 
          var contI int
          var contJ int   
  	  for contI = 0; contI < len(mat); contI++ {
            	for contJ = 0; contJ < len(mat[0]); contJ++ {
            		//fmt.Print(mat[contI][contJ]," ")
            		fmt.Print("")
	  	}
	  	fmt.Println()
	  }
}


func fatorialI(num int) int{
	var cont int
	var resposta int
	resposta = 1
	for cont= 2;cont<num; cont++{
		resposta = resposta * cont
	}
	
	return resposta
}

func fatorialR(num int) int{
	var resposta int
	
        if num < 2 {
        	resposta = 1	
        } else {
        	resposta = resposta * fatorialR(num-1)
        }
        
	return resposta
}

func main(){

        var contOrdem int
        var numRepeticoes int
        var contRepeticoes int
        var inicio time.Time
 	var fim time.Time
        // parametros do experimento
	var ordens []int
	//ordens = []int{3, 5, 7, 9, 11}
	ordens = []int{8, 20, 30, 40, 50}
	numRepeticoes = 20
	//estruturas para processamento do tempo dos experimentos
	var tempoBaseline []int64
	var tempoHacked []int64
        tempoBaseline = make([]int64, len(ordens))
        tempoHacked   = make([]int64, len(ordens))  
        var tempoExperimento int64
    
       // para o exemplo usando matriz
        var numLinhas int
        var numColunas int  
        numLinhas = 50
        numColunas = 50
        var cont int 
        var matrix [][]int
        matrix = make([][]int, numLinhas)
	for cont = 0; cont < numLinhas; cont++ {
	    matrix[cont] = make([]int, numColunas)
	}
        //fim da inicalizacao da matriz
      
       for contOrdem = 0; contOrdem < len(ordens); contOrdem++ {
		for contRepeticoes = 0; contRepeticoes < numRepeticoes; contRepeticoes++{		
			//fmt.Print("Ordem: ", ordens[contOrdem])
			//fmt.Println(" - Repeticao: ", contRepeticoes)			
	       		// criar a matriz com numeros aleatorio
	       		// inicia a matriz
	       		
	       		//medir o tempo do baseline com a matriz
			inicio = time.Now()
			fatorialR(ordens[contOrdem])
			fim = time.Now()
			tempoExperimento = fim.UnixNano() - inicio.UnixNano()
			//fmt.Println(tempoExperimento)
			tempoBaseline[contOrdem] = tempoBaseline[contOrdem] + tempoExperimento
			
			//medir o tempo da hacked com a mesma matriz anterior
			inicio = time.Now()
			fatorialI(ordens[contOrdem])
			fim = time.Now()
			tempoExperimento = fim.UnixNano() - inicio.UnixNano()
			//fmt.Println(tempoExperimento)
			tempoHacked[contOrdem] = tempoHacked[contOrdem] +  tempoExperimento
		}
		//fmt.Print()
		//fmt.Println(tempoBaseline)
		//fmt.Println(tempoHacked)
		
	}
	
	for contOrdem = 0; contOrdem < len(ordens); contOrdem++ {
		tempoBaseline[contOrdem] = tempoBaseline[contOrdem]/3
		tempoHacked[contOrdem] = tempoHacked[contOrdem]/3
	}
	fmt.Println("Tempo no Total")
	fmt.Println(tempoBaseline)
	fmt.Println(tempoHacked)
}
