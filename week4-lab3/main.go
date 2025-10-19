package main

import (
	"errors"
	"fmt"
)

type Student struct{
	ID string `json:"id"` // แปลง ID เป็น id สำหรับ json
	Name string `json:"name"`
	Email string `json:"email"`
	Year int `json:"year"`
	GPA float64 `json:"gpa"`
}

func (s *Student) IsHonor() bool{
	return s.GPA >= 3.50
}

func (s *Student) Validate() error{
	if s.Name == ""{
		return errors.New("name is required")
	}
	if s.Year <1 || s.Year > 4{
		return errors.New("year must be between 1-4")
	}
	if s.GPA <0 || s.GPA >4{
		return errors.New("gpa must be between 0-4")
	}
	return nil // nil = ว่าง(ไม่มี error) || ถ้า error ให้ return errors
}

func main(){
	//var st Student = Student {ID:"1", Name:"theerapong", Email:"poonkwan_t@su.ac.th", Year:4, GPA: 3.75}
	
	students := []Student{
		{ID:"1", Name:"theerapong", Email:"poonkwan_t@su.ac.th", Year:4, GPA: 3.75},
		{ID:"2", Name:"alice", Email:"alice_t@su.ac.th", Year:4, GPA: 2.75},
	}
	
	newStuent := Student {ID:"3", Name:"tpo", Email:"tpo_t@su.ac.th", Year:4, GPA: 3.50}
	students = append(students, newStuent)

	
	for i, student := range students { // ถ้าไม่อยากใช้ตัวแปร ให้เปลี่ยน i เป็น _ แทน
	fmt.Printf("%d Honor = %v\n",i, student.IsHonor())
	fmt.Printf("%d Validation = %v\n",i, student.Validate())
	}
}