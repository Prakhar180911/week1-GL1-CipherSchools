package main

import "fmt"

// func main() {
// 	// a()
// 	b()
// }
// func a() {
// 	fmt.Print(12)
// }
// func b() {
// 	c()
// }
// func c() {

// }
func main() {
	w := new(int)
	// name := new(string)
	t := "hjbjk"
	_, x := c(*w, &t)
	// fmt.Println(result)
	fmt.Println(x)
	// fmt.Println(*name)
	fmt.Println(w)
	// r, _ := b(1, 2, 3, 4, 5, 6, 6)
	// fmt.Print(r)
}

// func a() (int, int) {
// 	return 122, 19
func a() (int, string) {
	return 122, "weqrr3"

}
func b(args ...int) (bool, int) {
	// fmt.Println(args)
	for _, v := range args {
		fmt.Println(v)
	}
	return true, 11
}
func c(w int, name *string) (i int, j string) {
	i = 10
	j = "prakhar"
	w = 100
	*name = "code"
	return
}
