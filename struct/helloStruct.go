package main

import "fmt"

type person struct {
    name string
    age  int
}

type plane struct {
    modelo string
    serial  string
    motorEsquerdo bool
    motorDireito bool
    velocidade int
}


func ignicao(aviao *plane){ 
	aviao.motorEsquerdo = true
	aviao.motorDireito = true
        fmt.Println("Vrurmmmmmmm")  
}

func acelera(aviao *plane){ 
        if(aviao.motorEsquerdo && aviao.motorDireito){
		aviao.velocidade = aviao.velocidade + 10
		fmt.Println("humMMMMMM") 
        } else {
        	fmt.Println("Motores Desligados!")
        }
}


func main() {

    var professor person

    professor = person{name: "Moises", age: 50}
    fmt.Println(professor.name)
    
    var baraoVermelho plane
    
    baraoVermelho = plane{
    	    modelo: "Albatros D",
	    serial:  "789/16",
	    motorEsquerdo: false,
	    motorDireito: false,
	    velocidade: 0 }
    
    fmt.Println(baraoVermelho.modelo)
    ignicao(&baraoVermelho)
    acelera(&baraoVermelho)
    fmt.Println(baraoVermelho)

}
