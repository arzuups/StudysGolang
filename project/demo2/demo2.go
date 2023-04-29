package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func AddProduct() {
	product := Product{Id: 4, ProductName: "Charging", CategoryId: 1, UnitPrice: 500.0}
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	}

	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	fmt.Println("Added!")
}

func AddProduct2() {
	product := Product{Id: 5, ProductName: "Microphone", CategoryId: 1, UnitPrice: 2500.0}
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	}
	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	fmt.Println("Added!")

}

func AddProduct3() {
	product := Product{Id: 6, ProductName: "Telephone", CategoryId: 1, UnitPrice: 17000.0}
	jsonProduct, err := json.Marshal(product)
	if err != nil {
		fmt.Println(err)
	}
	response, err := http.Post("http://localhost:3000/products", "application/json;charset=utf-8", bytes.NewBuffer(jsonProduct))
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	fmt.Println("Added!")
}

