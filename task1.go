package main // package 

import "fmt"

type me struct {
name string
age  int
}

func main() {
var x me
x.name = "hassan sohail"
x.age = 22
dis(x)

}
func dis(x me) {
fmt.Println("hello i am", x.name)
fmt.Println("hello my age is ", x.age)
}
