package structs

func Demo1() {
  fmt.Println(product{"Telephone" , 5.500 , "Iphone"})
}

type product struct{
  name string
  unitPrice float64
  
}
