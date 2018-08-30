package main
import "fmt"

func main(){
  var age int = 8
  const year float64 = 2.9840;
  name := "Rajat"
  fmt.Printf("Up to 3 decimals %.3f \n", year)
  fmt.Printf("Type: %T\n", year)
  fmt.Printf("Age:. %d\n", age)
  fmt.Printf("Binary notation %b\n", age)
  fmt.Printf("Character notation %c\n", 44)
  fmt.Printf("hexa decimal notation %x\n", age)
  fmt.Printf("Scientific notation %e\n", age)
  fmt.Println(age,year,name)
}
