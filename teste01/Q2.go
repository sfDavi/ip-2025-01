package main
import "fmt"

func main() {
  var num float64
  fmt.Scanln(&num)
  
  if(num > 20 && num < 90){
      fmt.Println("NUMERO COMPREENDIDO ENTRE 20 E 90")
  }else{
      fmt.Println("NUMERO NAO COMPREENDIDO ENTRE 20 E 90")
  }
}
