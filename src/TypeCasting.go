package main
import "fmt"

func main()  {
	
	var age float32 = 9.2;
	fmt.Println(int(age)); // converting decimal to int; display number before decimal point

	var score int = 9;
	fmt.Printf("%.2f",float32(score)); // converting int to decimal; display number with decimal points
}

// Go does not support implicit typecasting (i.e) declaring a variable with int and using printf along with %g to convert to float32