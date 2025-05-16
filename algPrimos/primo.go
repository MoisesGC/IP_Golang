package main
  
import "fmt"
  
func echo(name string) string {
    return name
}
 
func soma(num1 int, num2 int) int {
	var cont int
	var acum int
	
	acum = 0
	cont = 0
	for(cont < num1){
		acum = acum + num2
		//fmt.Println(acum)
		cont = cont + 1
	}
	return acum
} 
  
func expon(base int, exponte int) int {
	var cont int
	var acum int
	
	acum = 1
	cont = 0
	for(cont < exponte){
		acum = soma(acum,base)
		//fmt.Println(acum)
		cont = cont + 1
	}
	return acum
}   
  
  
func  main(){
  	   
	   var numero1 int
	   var numero2 int
	   var numero3 int
	   
           numero1 = 3
           numero2 = 5
           numero3 = expon(numero1,numero2)
           fmt.Println(numero3)
           
           
           
  }
