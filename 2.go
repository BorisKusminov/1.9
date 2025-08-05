package main

import (
	"fmt"
)

type Student struct {
	Name   string
	Grades []float64
}

func (s Student) AverageGrade() float64 {
	if len(s.Grades) == 0 {
		return 0.0 // Возвращаем 0, если нет оценок
	}
	sum := 0.0
	for _, grade := range s.Grades {
		sum += grade
	}
	return sum / float64(len(s.Grades))
}
func main() {
	Student1 := Student{
		Name:   "Petrov Ivan",
		Grades: []float64{4.5, 5.0, 3.8, 4.2},
	}
	fmt.Println("Средний балл студента:", Student1.Name, "=", Student1.AverageGrade())
}
