package main

import "fmt"

func main() {
	students := []struct {
		Name   string
		Height float64
	}{
		{"Kyle", 173.4},
		{"Ken", 164.5},
		{"Ryu", 178.8},
		{"Honda", 154.2},
		{"Hwarang", 188.8},
		{"Lebron", 209.8},
		{"Hodong", 197.7},
		{"Tom", 164.8},
		{"Kevin", 164.8},
	}

	// easy way
	// 160 ~ 170
	for i := 0; i < len(students); i++ {
		if students[i].Height >= 160 && students[i].Height < 170 {
			// fmt.Println(students[i].Name, students[i].Height)
		}
	}

	// 여러번 입력을 받아서 출력하는 경우 -> 정렬
	// 160 ~ 170, 170 ~ 180, 180 ~ 190, 190 ~ 200, 200 ~ 210
	// 최대 3m 키가 존재한다고 가정
	var heightMap [3000][]string
	for i := 0; i < len(students); i++ {
		idx := int(students[i].Height * 10)
		heightMap[idx] = append(heightMap[idx], students[i].Name)
	}

	for i := 1400; i < 1700; i++ {
		for _, name := range heightMap[i] {
			if name != "" {

				fmt.Println(name, float64(i)/10)
			}
		}
	}

	for i := 1800; i < 2100; i++ {
		for _, name := range heightMap[i] {
			if name != "" {

				fmt.Println(name, float64(i)/10)
			}
		}
	}

}
