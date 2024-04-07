package non_primitive

import "fmt"

// Struct คือชนิดข้อมูลที่ใช้สำหรับสร้างโครงสร้างข้อมูลที่มีฟิลด์หลายๆ ตัว
type Person struct {
	Name string
	Age  int
}

func Non_Primitive() {
	// Array เป็นชนิดข้อมูลที่เก็บข้อมูลที่มีขนาดคงที่
	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 2
	numbers[2] = 3
	fmt.Println("Array:", numbers)

	// Slice เป็นโครงสร้างข้อมูลที่สามารถขยายขนาดได้
	var colors []string
	colors = append(colors, "red")
	colors = append(colors, "green")
	colors = append(colors, "blue")
	fmt.Println("Slice:", colors)

	// Map เป็นโครงสร้างข้อมูลที่เก็บคีย์-ค่าแบบไม่มีลำดับ
	var grades map[string]int
	grades = make(map[string]int)
	grades["John"] = 90
	grades["Alice"] = 95
	grades["Bob"] = 85
	fmt.Println("Map:", grades)

	// Function เป็นชนิดข้อมูลที่สามารถถูกส่งเป็นอาร์กิวเมนต์หรือส่งคืนได้
	var add func(int, int) int
	add = func(a, b int) int {
		return a + b
	}
	fmt.Println("Function:", add(3, 5))

	// Struct เป็นชนิดข้อมูลที่สามารถรวมฟิลด์ต่างๆ เข้าไว้ในโครงสร้างเดียวกัน
	var person1 Person
	person1.Name = "Alice"
	person1.Age = 30
	fmt.Println("Struct:", person1)
}
