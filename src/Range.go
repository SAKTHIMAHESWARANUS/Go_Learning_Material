package main
import ("fmt")

func main()  {
	letters := [4]string{"a","b","c","d"};
	for idx, val := range letters{ // if dont want idx then for _,val 
		fmt.Printf("%v %v\n",idx,val);       // if we dont want val then for idx,_
	}
}
	