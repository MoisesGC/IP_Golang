package main
  
import (
    "errors"
    "fmt"
)
  
 
func ehPrimo(numero int) bool {
	var cont int
	var ePrimo bool
	ePrimo = true
	for cont = 2; cont < numero; cont++ {
		if (numero % cont)  == 0 {
			fmt.Println(cont)
			ePrimo = false
		}
	}
	return ePrimo
}   

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
    
func main(){
  	   
	  // var numero1 int
	  // var numero2 int
	  // var numero3 int
	  var cont int
	   
         //  numero1 = 3
         //  numero2 = 5
          
          // fmt.Println(ehPrimo(1000002))
          
           var numerosTeste []int
           numerosTeste = make([]int, 7)
           numerosTeste[0] = -1
           numerosTeste[1] = 0
           numerosTeste[2] = 1
           numerosTeste[3] = 2
           numerosTeste[4] = 5
           numerosTeste[5] = 12
           numerosTeste[6] = 37
          
           for cont = 0; cont < 7; cont++ {
            	fmt.Println(ehPrimoV2(numerosTeste[cont]))
	   }
          
          
         /*  fmt.Println(ehPrimoV2(-1))   
           fmt.Println(ehPrimoV2(0))
           fmt.Println(ehPrimoV2(1))
           fmt.Println(ehPrimoV2(2))
 	   fmt.Println(ehPrimoV2(5))
 	   fmt.Println(ehPrimoV2(12))
 	   fmt.Println(ehPrimoV2(37)) */
}
