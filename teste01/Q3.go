package main
import "fmt"

func main() {
  soma := 0
  for i:=1; i<101; i++{
      fmt.Printf("%d\n",i)
      soma += i
  }
  fmt.Printf("%d\n",soma)
}
