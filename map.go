package main

import "fmt"

func main() {

	//define a map and assign key values
	mapexample := map[string]string{
		"keith": "20/01/1991",
		"adam":  "18/05/1991",
		"kevin": "19/07/1997",
	}

	fmt.Println(mapexample) // map[adam:18/05/1991 keith:20/01/1991 kevin:19/07/1997]

	//define a map and assign key values later
	ages := map[string]int{}
	ages["keith"] = 28
	ages["adam"] = 29
	ages["kevin"] = 30

	fmt.Println(ages) // map[adam:29 keith:28 kevin:30]

	// defining a duplicate key with value will use the latest value

	ages["kevin"] = 10

	fmt.Println(ages["kevin"]) // 10
	fmt.Println(ages)          // map[adam:29 keith:28 kevin:10]

	// delete an element in a map

	delete(ages, "keith")

	fmt.Println(ages) // map[adam:29 kevin:10]

}
