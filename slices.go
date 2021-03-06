package main

import "fmt"

func main() {
	// Unlike arrays, slices are typed only by the elements they contain 
	// (not the number of elements). To create an empty slice with non-zero length, 
	// use the builtin make. Here we make a slice 
	// of strings of length 3 (initially zero-valued).

	s := make([]string, 3)
	fmt.Println("emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"		

	// We can set and get just like with arrays.
	fmt.Println("set:", s)
	fmt.Println("get:",s[2])


	fmt.Println("length of slice: ", len(s))

	 // One is the builtin append, which returns a slice containing 
	 // one or more new values. Note that 
	 // we need to accept a return value from append as we may get a new slice value.
	s = append(s,"d")
	s = append(s,"e","f")
	fmt.Println("apd:",s)


	// Slices can also be copy’d. Here we create an empty slice 
	// c of the same length as s and copy into c from s.
	c := make([]string,len(s))
	copy(c,s)

	fmt.Println("c:",c)

	// Slices support a “slice” operator with the syntax slice[low:high]. 
	// For example, this gets a slice of the elements s[2], s[3], and s[4].

	l := s[2:5]
	fmt.Println("l:",l)

	l = s[:5]

	fmt.Println("l:",l)

	l = s[2:]
	fmt.Println("l:",l)

	// We can declare and initialize a variable for slice in a single line as well.
	t := []string{"g","h","l"}

	fmt.Println("t:",t)

	// Slices can be composed into multi-dimensional data structures.
	// The length of the inner slices can vary, unlike with multi-dimensional arrays.

	twoD := make([][]int,3)

	for i := 0; i < 3; i++ {
		innerLength := i + 1
		twoD[i] = make([]int, innerLength)
		for j := 0; j < innerLength; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("twoD:",twoD)

	

}