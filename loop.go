package main

import "fmt"

func main() {
        
	i := 1
	for i <=3 {
		fmt.Println(i)
		i = i + 1
	}
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}
	
	for i := 0; i < 5 ; i++ {
		fmt.Println(i)
		
	}
	for {
		fmt.Println(i)
		break
	}
	for n := 0; n < 10 ; n++ {
      if n%2 == 0 {
		continue
	  }
		fmt.Println(n)
	}
	
}

