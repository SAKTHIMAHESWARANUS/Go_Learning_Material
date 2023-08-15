package main
import "fmt"

func main()  {

	var a = func(){ 		// declaring Anonymous function
		fmt.Println("hello");
	}
	
	a(); // caling the Anonymous function
}

// We can perform all the operations we do with normal functions such as we can pass parameters, return single or multiple values