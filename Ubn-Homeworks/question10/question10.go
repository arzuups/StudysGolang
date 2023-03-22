/*When an array is given, prepare the code and flowchart that gives the number of elements in the array.
For example :
Input : arr[] = {3, 3, 3, 2, 3, 4, 1} 
Output : There are 3 "3" elements, 1 "2" element, 1 "4" element, 1 "1" element.
Input : arr[] = {100, 100, 100, 100, 100, 30, 20, 20} 
Output : There are 4 "100" elements, 2 "20" elements, 1 "30" element.*/

package main

import "fmt"

func main() {
	arr := []int{30, 60, 40, 70, 30, 50, 10, 80, 40, 50, 70, 20}
	freq := make(map[int]int)
	for _, num := range arr {
		freq[num] = freq[num] + 1
	}
	fmt.Println("Frequency of the Array is : ", freq)
}
