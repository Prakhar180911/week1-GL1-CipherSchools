package main

import "fmt"

func main() {
	// 	for i := 0; i < 3; i++ {
	// 		for i := 0; i < 3; i++ {
	// 			fmt.Println(i)
	// 		}
	// 	}
	// }

	// 	for i := 0; i < 3; i++ {
	// 		if i == 1 {
	// 			continue
	// 		}
	// 		fmt.Println(i)
	// 	}
	// }
	// 	nums := []int{1, 3, 2, 4, 0}
	// 	for index, value := range nums {
	// 		fmt.Println(index)
	// 		fmt.Println(value)
	// 	}
	// }
	//  nums := []int{1, 3, 2, 4, 0}
	nums := []string{"prakhar", "tripathi"}
	for _, value := range nums {
		// if value == "h" {
		// 	break
		// }

		fmt.Println(value)
	}
}
