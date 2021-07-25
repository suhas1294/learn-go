// decliaring a array
var x [5]int // integer array of length 5, zero-based indexing


// slices
var x = []int{1,2,3,4}

// loop over slice using 'range'
for i,v  := range x{
	fmt.Println(i, v);
}

// slicing a slice : upto but not including 
fmt.Println(x[:]) // {1,2,3,4}
fmt.Println(x[1:]) // {2,3,4}
fmt.Println(x[1:3]) // {2,3}

// appending to a slice
x = append(x, 23,45,67);

y := []int{99,88,77}
x = append(x, y...) // unfurling

// ...a int: unlimited values of type int
// a... : unfurling a slice

// make([]T, length, capacity)
make([]int, 10, 100)