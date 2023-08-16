package main
import "fmt"

func main()  {
	var a int;
	fmt.Print("enter age :");
	fmt.Scan(&a);
	fmt.Println(a);
}

// We use the Scanln() function to get input values up to the new line. When it encounters a new line, it stops taking input values

// The fmt.Scanf() function is similar to Scanln() function. The only difference is that Scanf() takes inputs using format specifiers

// Example Scanf("%s %d",&str,&age);

// Data Type	Format Specifier
// integer			%d
// float			%g
// string			%s
// bool				%t