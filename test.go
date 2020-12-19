/**
* @author Jee
* @date 2020/12/5 0:28
 */
package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 34, 5, 5, 6, 67, 77, 87}
	fmt.Println(data)
	s1 := data[:0:0]
	s1 = append(s1, 123)
	s2 := data[:0]
	s2 = append(s2, 10)
	data = []int{}
	data = append(data, 99)
	fmt.Printf("%v %p\n", data, &data[0])
	fmt.Printf("%v %p\n", s2, &s2[0])
	fmt.Printf("%v %p\n", s1, &s1[0])
}
