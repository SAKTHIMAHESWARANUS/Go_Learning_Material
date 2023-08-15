package main
import "fmt"

func main()  {
  	a := 32;
	var p *int = &a ;
	fmt.Println(*p);
}

// pointers are similar to c
// stores address of other variable and dereference is required using accessing the value stored in it