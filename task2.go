package main // pkg

import "fmt"

type employee struct {
name     string
salary   int
position string
}
type company struct {
companyName string
employees   []employee
}

func main() {
x1 := employee{"Amir", 80000, "Full-Stack Developer"}
x2 := employee{"hassan", 50000, "CEO"}
x3 := employee{"wajeeha", 6000, "writer"}
member := []employee{x1, x2, x3}
org := company{"magpiestarpvtltd", member}

fmt.Printf("comp details=  %v  ", org)
}
