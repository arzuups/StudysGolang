/*You need to run main.go to print to the terminal.*/
package restful

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// POST Request
func Demo2() {
	todo := Todo{1, 2, "Shopping will be done!", false}
	jsonTodo, err := json.Marshal(todo)
	if err != nil {
		fmt.Println(err)
	}

	response, err := http.Post("https://jsonplaceholder.typicode.com/todos", "application/json;charset=utf-8", bytes.NewBuffer(jsonTodo))
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var todoResponse Todo
	json.Unmarshal(bodyBytes, &todoResponse)
	fmt.Println(todoResponse)
}
