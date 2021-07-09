package Classes

import (
	"fmt"
)

type Student struct{
	Name, Surname string
	Age int16
	CourseGrade []int
}

func (s Student) PrintLengthOfCourseGrades() {
	fmt.Println("Arratyin Eleman Sayisi: %d",len(s.CourseGrade))
}

func (s Student) CalculateAverateCourseGrades() float32{
	return 0
}