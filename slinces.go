package main

import "fmt"

func main() {
	s := make([]string, 4)
	fmt.Println(s)

    s [0] = "j"
    s [1] = "q"
	s [2] = "k"
	s [3] = "o"
	fmt.Println("set:" ,s)
	fmt.Println("get:", s[2])
	fmt.Println("len:", len(s))

c := make([]string, len(s))
copy(c,s)
fmt.Println("cpy:", s)
l := s[2:5]
fmt.Println("sl1:", l)


l = s[:6]
fmt.Println("sl2:", l)
l = s[2:]
fmt.Println("sl3:", l)

t := []string{"a", "b", "c", "d", "e", "f", "q"}
fmt.Println("dcl",t)

twoD := make([][]int, 10)
for i := 0; i <10 ; i++ {
	innerLen := i + 1
	twoD[i]=make([]int, innerLen)
	for  j := 0;  j < innerLen;  j++ {
		twoD[i][j] = i + j

		
	}
	
}
fmt.Println("2d", twoD)
}

