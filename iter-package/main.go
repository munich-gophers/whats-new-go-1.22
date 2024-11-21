// package main
//
// import (
// 	"fmt"
// )
//
// func Backward(s []string) func(func(int, string) bool) {
// 	return func(yield func(int, string) bool) {
// 		for i := len(s) - 1; i >= 0; i-- {
// 			if !yield(i, s[i]) {
// 				return
// 			}
// 		}
// 	}
// }
//
// func main() {
// 	s := []string{"hello", "munich", "gophers"}
// 	for i, x := range Backward(s) {
// 		fmt.Println(i, x)
// 	}
// }
