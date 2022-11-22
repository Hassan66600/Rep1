package main //comment

import (
	"fmt"
	"strings"
	"crypto/sha256"
)

type Student struct { //comment2
	rollnumber int
	name       string
	address    string
	course     []string
}

func NewStudent(rollno int, name string, address string, crs []string) *Student { //new student
	s := new(Student)
	s.rollnumber = rollno
	s.name = name
	s.address = address
	s.course = crs
	return s
}

type StudentList struct { // comment3
	list []*Student
}

func (ls *StudentList) CreateStudent(rollno int, name string, address string, crs []string) *Student { //comment
	st := NewStudent(rollno, name, address, crs)
	ls.list = append(ls.list, st)
	return st
}
func PrintStudentList(ls *StudentList) { //comemtn3
	for p := range ls.list {
		fmt.Printf("%s List %d %s\n", strings.Repeat("=", 25), p+1, strings.Repeat("=", 25))
		fmt.Printf("student rollno         %d\n", ls.list[p].rollnumber)
		fmt.Printf("student name           %s\n", ls.list[p].name)
		fmt.Printf("student address        %s\n", ls.list[p].address)
		fmt.Printf("student course         %s\n", ls.list[p].course)
	}
}
func main() {
	student := new(StudentList)
	var x = []string{"maths", "physics", "cnet"}
	var y = []string{"InfoSec", "crypto", "english"}
	student.CreateStudent(17, "Hassan", "DHA", x)
	student.CreateStudent(85, "Sohail", "DHA", y)
	PrintStudentList(student)
}
