package main

import (
	"fmt"
)

func main() {
	//var name string = "theerapong" // สร้างตัวแปรแบบใช้ "var" ต้องระบุ data type ด้วย *สามารถใช้ประกาศตัวแปรแบบ global ได้*
	var age int = 20

	email := "poonkwan_t@su.ac.th" // สร้างตัวแปรแบบใช้ ":=" *ไม่สามารถใช้ประกาศตัวแปรแบบ global ได้*
	gpa := 3.60
	firstname, lastname := "theerapong", "poonkwan"

	fmt.Printf("Name %s %s, age %d, email %s, gpa %.2f\n", firstname, lastname, age, email, gpa)
}
