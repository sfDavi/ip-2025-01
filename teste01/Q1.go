package main
import "fmt"

func main() {
  var num float64
  fmt.Scanln(&num)
  
  if(num >0){
      fmt.Println("NUMERO POSITIVO")
  }else if(num < 0){
      fmt.Println("NUMERO NEGATIVO")
  }else{
      fmt.Println("NUMERO NULO")
  }
}
