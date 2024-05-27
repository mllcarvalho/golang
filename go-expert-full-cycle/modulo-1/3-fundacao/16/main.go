package main

import "fmt"

func main() {

	var minhaVar interface{} = "Matheus"

	res, ok := minhaVar.(string)

	fmt.Println(res, ok)
}
