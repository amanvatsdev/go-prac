package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func main() {

	jsonData:=[]byte(`[{"name":"aman","age":25},{"name":"Raju","age":34}]`)

	// jsonData := `{"name":"Aman","Age":25}`

	var p []person

	err := json.Unmarshal(jsonData,&p)
	if err !=nil {
	fmt.Println(err)
	}
	fmt.Println(string(jsonData))
	fmt.Println(p)

}
