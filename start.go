package main

import "fmt"

func main() {
	/*  var age int = 44
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

	  fmt.Println("True && False = ", true && false)
	  fmt.Println("True || False = ", true || false)
	  fmt.Println("!True = ", !true)

	  //First for loop
	  fmt.Println("First for loop")
	  for j:=0; j<=3; j++{
	    fmt.Println(j)
	  }
	  //Second for loop
	fmt.Println("Second for loop")
	  j:=0;
	  for j<=3{
	    fmt.Println(j)
	    j++
	  }

	  if age >= 16 {
	    fmt.Printf("Enjoy your driving license. Your driving license id is %b\n", age)
	  } else {
	    fmt.Println("You are not allowed to drive")
	  }*/

	//Forloop on array
	var favNum = [5]float64{4, 5, 6, 7, 8}
	for _, value := range favNum {
		fmt.Println(value)
	}
}
