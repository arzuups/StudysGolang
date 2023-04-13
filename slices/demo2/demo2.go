/*You need to run main.go to print to the terminal.*/
package slices

import (
        "fmt"
)
      
func Demo2() {
  
  names := []string{"Jack", "Olivia", "Robert"}
  fmt.Println(names)
  namesCopy := make([]string, len(names))
  fmt.Println(namesCopy)
  copy(namesCopy, names)
  fmt.Println(namesCopy)
  names = append(names, "Berrica")
  fmt.Println(names)
  fmt.Println(namesCopy)
  
  fmt.Println(names[1:3])
  fmt.Println(names[1:])
  fmt.Println(names[:2])
  
}
