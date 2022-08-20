package main

import "fmt"

type Student struct {
	name string
	// math, english float64
}

type User struct {
	name string
}

// (s Student)レシーバという
// func (s Student) avg() {
// 	fmt.Println(s.name, (s.math+s.english)/2)
// }

// func main() {
// 	a001 := Student{"kojima", 80, 70}
// 	a001.avg()
// }

// func (s Student) avg(math, english float64) {
// 	fmt.Println(s.name, (math+english)/2)
// }

// func main() {
// 	a001 := Student{"kojima"}
// 	a001.avg(80, 70)
// }

// func (s Student) avg(math, english float64) float64 {
// 	return (math + english) / 2
// }

// func main() {
// 	a001 := Student{"kojima"}
// 	fmt.Println(a001.avg(80, 70))
// }

func (s Student) avg(math, english float64) (avgResult float64) {
	avgResult = (math + english) / 2
	return
}

func (u User) cal(weight, height float64) (calAvg float64) {
	calAvg = (weight / height / height) * 10000
	return
}

func main() {
	a001 := Student{"kojima"}
	fmt.Println(a001.avg(70, 80))

	bmi := User{"tanaka"}
	fmt.Println(bmi.name, "のBMIは、", bmi.cal(62, 169))
}
