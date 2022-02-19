package hamt

type HAMT struct {
	root *node
	size uint32
}

func NewHAMT() *HAMT {
	return &HAMT{
		root: newNode(),
		size: 0,
	}
}

// func main() {
// 	// create a new HAMT
// 	h := NewHAMT()
// 	// insert 13 at index 2
// 	h2, _ := Append(h, 2, 13)
// 	// insert 912 at index 45
// 	h3, _ := Append(h2, 45, 912)
// 	// get value at index 45
// 	d, err := Get[int](h3, 45)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Println(d)
// 	}
// 	// update value at index 45 to value 34
// 	h4, err := Update(h3, 45, 34)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	// get value at index 45
// 	d, err = Get[int](h4, 45)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	fmt.Println(d)
// 	// pop value at index 45
// 	h5, d, err := Pop[int](h4, 45)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	fmt.Println(d)
// 	// get value at index 45
// 	d, err = Get[int](h5, 45)
// 	if err != nil {
// 		// since the value doesn't exist into "h5" the error != nil
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	fmt.Println(d)
// }
