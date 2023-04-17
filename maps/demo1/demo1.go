package main

import (
	"fmt"
)

func Demo1() {
      //key-value
	Dictionary := make(map[string]string)
	
	Dictionary["Table"] = "Masa"
	Dictionary["Book"] = "Kitap"
	Dictionary["Pencil"] = "Kalem"
	Dictionary["Notebook"] = "Not Defteri"
	
	fmt.Println(Dictionary)
  
}
