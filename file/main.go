package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	err := os.WriteFile("hello.txt", []byte("Hello Go!"), 0644)
	if err != nil {
		fmt.Println("error writing:", err)
		return
	}
	fmt.Println("file written!")

	data, err := os.ReadFile("hello.txt")
	if err != nil {
		fmt.Println("error reading:", err)
		return
	}
	fmt.Println("file contents:", string(data))

	p := Person{Name: "John", Age: 25}
	jsonData, _ := json.Marshal(p)
	fmt.Println(string(jsonData))

	jsonStr := `{"name":"Jane","age":30}`
	var p2 Person
	json.Unmarshal([]byte(jsonStr), &p2)
	fmt.Println(p2.Name, p2.Age)
}
