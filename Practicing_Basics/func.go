
// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type Product struct {
// 	Id      int    `json:"id"`
// 	Name    string `json:"name"`
// 	Price   int    `json:"price"`
// 	Instock bool   `json:"instock"`
// }
// type data struct {
// 	Product Product `json:"product"`
// }

// func main() {
// 	jsondata := `{
//   "product": {
//     "id": 1001,
//     "name": "Laptop",
//     "price": 75000,
//     "inStock": true
//   }
// }`
// 	var Data data
// 	err := json.Unmarshal([]byte(jsondata), &Data)
// 	if err != nil {
// 		fmt.Println("error found")
// 	}
// 	fmt.Println(Data)
// }

// package main

// import "fmt"

// var name string

// func main() {
// 	fmt.Println("give a name")
// 	fmt.Scan(&name)
// 	call()
// }
// func call() {
// 	fmt.Println(name + " Hello World")

// }

// // package main

// // import "fmt"

// // func main() {

// // 	// a, b := omg()
// // 	// fmt.Println(a + b + xxx())
// // 	// msg()

// // 	func() {
// // 		msg()
// // 	}()

// // }

// // func omg() (int, int) {
// // 	return 7, 15
// // }

// // func xxx() int {
// // 	return 9
// // }

// // func msg() {
// // 	fmt.Println("Hello from Golang !!")
