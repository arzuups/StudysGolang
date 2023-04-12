package slices

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
  
}
